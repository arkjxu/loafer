package loafer

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// SlackApp - A simple slack app starter kit
type SlackApp struct {
	opts              SlackAppOptions
	server            *http.Server                                                                           // Slack App options
	distCB            func(installRes *SlackOauth2Response, res http.ResponseWriter, req *http.Request) bool // Handler for app distribution
	errorCB           func(res http.ResponseWriter, req *http.Request, err error)                            // Handler for error
	cmds              map[string]func(ctx *SlackContext)                                                     // List of command handlers
	shortcutListeners map[string]func(ctx *SlackContext)                                                     // List of shortcut handlers
	actionListeners   map[string]func(ctx *SlackContext)                                                     // List of action handlers
	submitListeners   map[string]func(ctx *SlackContext)                                                     // List of view submission handlers
	closeListeners    map[string]func(ctx *SlackContext)                                                     // List of view close handlers
	eventListeners    map[string]func(ctx *SlackContext)                                                     // List of slack event listeners
}

// SlackAuthToken - Slack App Auth Token
type SlackAuthToken struct {
	Workspace string
	Token     string
}

// SlackAppOptions - Slack App options
type SlackAppOptions struct {
	Name          string                                  // Slack App name
	Prefix        string                                  // Prefix of routes
	TokensCache   func(workspace string) []SlackAuthToken // List of available workspace tokens
	ClientSecret  string                                  // App client secret
	ClientID      string                                  // App client id
	SigningSecret string                                  // Signning secret
}

// SlackContext - Slack request context
type SlackContext struct {
	Body  []byte
	Token string
	Req   *http.Request
	Res   http.ResponseWriter
}

// SlackOauth2Team - Slack App Access Response Team
type SlackOauth2Team struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// SlackOauth2User - Slack App Access Response User
type SlackOauth2User struct {
	ID          string `json:"id"`
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// SlackOauth2Response - Slack App Access Response
type SlackOauth2Response struct {
	Ok          bool            `json:"ok"`
	AccessToken string          `json:"access_token"`
	TokenType   string          `json:"token_type"`
	Scope       string          `json:"scope"`
	BotUserID   string          `json:"bot_user_id"`
	AppID       string          `json:"app_id"`
	Team        SlackOauth2Team `json:"team"`
	Enterprise  SlackOauth2Team `json:"enterprise"`
	AuthedUser  SlackOauth2User `json:"authed_user"`
}

// SlackSubscriptionEvent - Slack Subscription event
type SlackSubscriptionEvent struct {
	Type    string `json:"type"`
	EventTS string `json:"event_ts"`
}

// SlackSubscriptionEventRequest - Slack subscription event
type SlackSubscriptionEventRequest struct {
	Token     string                 `json:"token"`
	TeamID    string                 `json:"team_id"`
	APIAppID  string                 `json:"api_app_id"`
	Event     SlackSubscriptionEvent `json:"event"`
	Type      string                 `json:"type"`
	EventID   string                 `json:"event_id"`
	EventTime uint32                 `json:"event_time"`
}

// OnCommand - Add handler to command
func (a *SlackApp) OnCommand(cmd string, handler func(ctx *SlackContext)) {
	if a.cmds == nil {
		a.cmds = make(map[string]func(ctx *SlackContext))
	}
	a.cmds[cmd] = handler
}

// RemoveCommand - Remove a command to the app base on command
func (a *SlackApp) RemoveCommand(cmd string) {
	if a.cmds != nil {
		delete(a.cmds, cmd)
	}
}

// OnAction - Add an action handler to the app base on action_id
func (a *SlackApp) OnAction(actionID string, handler func(ctx *SlackContext)) {
	if a.actionListeners == nil {
		a.actionListeners = make(map[string]func(ctx *SlackContext))
	}
	a.actionListeners[actionID] = handler
}

// OnShortcut - Add an shortcut handler to the app base on callback_id
func (a *SlackApp) OnShortcut(callbackID string, handler func(ctx *SlackContext)) {
	if a.shortcutListeners == nil {
		a.shortcutListeners = make(map[string]func(ctx *SlackContext))
	}
	a.shortcutListeners[callbackID] = handler
}

// OnViewSubmission - Add handler to view submission base on callback_id
func (a *SlackApp) OnViewSubmission(callbackID string, handler func(ctx *SlackContext)) {
	if a.submitListeners == nil {
		a.submitListeners = make(map[string]func(ctx *SlackContext))
	}
	a.submitListeners[callbackID] = handler
}

// OnViewClose - Add handler to view close base on callback_id
func (a *SlackApp) OnViewClose(callbackID string, handler func(ctx *SlackContext)) {
	if a.closeListeners == nil {
		a.closeListeners = make(map[string]func(ctx *SlackContext))
	}
	a.closeListeners[callbackID] = handler
}

// OnEvent - Add handler to events
func (a *SlackApp) OnEvent(eventType string, handler func(ctx *SlackContext)) {
	if a.eventListeners == nil {
		a.eventListeners = make(map[string]func(ctx *SlackContext))
	}
	a.eventListeners[eventType] = handler
}

// OnError - Add handler to errors
func (a *SlackApp) OnError(handler func(res http.ResponseWriter, req *http.Request, err error)) {
	a.errorCB = handler
}

// OnAppInstall - Add handler to app distribution after it's been successfully installed
func (a *SlackApp) OnAppInstall(cb func(installRes *SlackOauth2Response, res http.ResponseWriter, req *http.Request) bool) {
	a.distCB = cb
}

// appInstall - Handler for app distribution
func (a *SlackApp) appInstall(res http.ResponseWriter, req *http.Request) {
	var installResponse SlackOauth2Response
	form := url.Values{}
	form.Set("code", req.URL.Query().Get("code"))
	form.Set("client_id", a.opts.ClientID)
	form.Set("client_secret", a.opts.ClientSecret)
	resp, err := http.Post("https://slack.com/api/oauth.v2.access", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		Response(&SlackContext{Res: res}, http.StatusInternalServerError, []byte("Unable to authorize Slack App for workspace"), nil)
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&installResponse)
	if err != nil {
		Response(&SlackContext{Res: res}, http.StatusInternalServerError, []byte("Unable to get Slack OAuth2 Access Response for workspace"), nil)
		return
	}
	if installResponse.Ok {
		avoidDefaultPage := false
		if a.distCB != nil {
			avoidDefaultPage = a.distCB(&installResponse, res, req)
		}
		if !avoidDefaultPage {
			Response(&SlackContext{Res: res}, http.StatusOK, []byte(strings.Replace(INSTALLSUCCESSPAGE, "{{APP_NAME}}", a.opts.Name, -1)), map[string]string{
				"Content-Type": "text/html; charset=utf-8"})
			return
		}
	} else {
		Response(&SlackContext{Res: res}, http.StatusInternalServerError, []byte("Slack App Access Request is not Ok"), nil)
		return
	}
}

// checkSlackSecret - Checking the signing secret of slack request
func (a *SlackApp) checkSlackSecret(signing string, ts string, body string) bool {
	data := strings.Join([]string{"v0", ts, body}, ":")
	signed := []byte(a.opts.SigningSecret)
	tested := hmac.New(sha256.New, []byte(signed))
	tested.Write([]byte(data))
	own := strings.Join([]string{"v0", hex.EncodeToString(tested.Sum(nil))}, "=")
	if own == signing {
		return true
	}
	return false
}

// interaction - Slack App interactions handler
func (a *SlackApp) interactions(res http.ResponseWriter, req *http.Request) {
	var event SlackInteractionEvent
	bodyText, err := ioutil.ReadAll(req.Body)
	if err != nil {
		a.errorHandling(res, req, err)
		return
	}
	defer req.Body.Close()
	queries, err := url.ParseQuery(string(bodyText))
	if err != nil {
		a.errorHandling(res, req, err)
		return
	}
	isAuthorizedCaller := a.checkSlackSecret(req.Header.Get("X-Slack-Signature"), req.Header.Get("X-Slack-Request-TimeStamp"), string(bodyText))
	if isAuthorizedCaller {
		err = json.Unmarshal([]byte(queries.Get("payload")), &event)
		if err != nil {
			a.errorHandling(res, req, err)
			return
		}
		accessToken := a.GetTokenForWorkspace(event.Team.ID)
		if accessToken == nil {
			a.errorHandling(res, req, fmt.Errorf("App is not installed to workspace: %s", event.Team.ID))
			return
		}
		ctx := &SlackContext{
			Body:  bodyText,
			Token: accessToken.Token,
			Res:   res,
			Req:   req}
		switch Type := event.Type; Type {
		case "shortcut":
			callbackID := event.CallbackID
			if handler, ok := a.shortcutListeners[callbackID]; ok {
				handler(ctx)
			} else {
				a.errorHandling(res, req, fmt.Errorf("Unrecognized shortcut: %s", callbackID))
				return
			}
			break
		case "block_actions":
			action := event.Actions[0]
			if handler, ok := a.actionListeners[action.ActionID]; ok {
				handler(ctx)
			} else {
				a.errorHandling(res, req, fmt.Errorf("Unrecognized action: %s", action.ActionID))
				return
			}
			break
		case "view_submission":
			if handler, ok := a.submitListeners[event.View.CallbackID]; ok {
				handler(ctx)
			} else {
				a.errorHandling(res, req, fmt.Errorf("Unrecognized submission event from view: %s", event.View.CallbackID))
				return
			}
			break
		case "view_closed":
			if handler, ok := a.closeListeners[event.View.CallbackID]; ok {
				handler(ctx)
			} else {
				a.errorHandling(res, req, fmt.Errorf("Unrecognized closed event from view: %s", event.View.CallbackID))
				return
			}
			break
		default:
			a.errorHandling(res, req, fmt.Errorf("Unrecognized interaction type: %s", event.Type))
			return
		}
	} else {
		Response(&SlackContext{Res: res}, http.StatusUnauthorized, []byte("Unauthorized"), nil)
		return
	}
}

// events - Slack App events handler
func (a *SlackApp) events(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		a.errorHandling(res, req, err)
		return
	}
	defer req.Body.Close()
	var event SlackSubscriptionEventRequest
	err = json.Unmarshal(body, &event)
	if err != nil {
		a.errorHandling(res, req, err)
		return
	}
	if event.Type == "url_verification" {
		Response(&SlackContext{Res: res}, http.StatusOK, body, map[string]string{
			"Content-Type": "application/json"})
		return
	}
	isAuthorizedCaller := a.checkSlackSecret(req.Header.Get("X-Slack-Signature"), req.Header.Get("X-Slack-Request-TimeStamp"), string(body))
	if isAuthorizedCaller {
		accessToken := a.GetTokenForWorkspace(event.TeamID)
		if accessToken == nil {
			a.errorHandling(res, req, fmt.Errorf("App is not installed for workspace: %s", event.TeamID))
			return
		}
		ctx := &SlackContext{
			Body:  body,
			Token: accessToken.Token,
			Res:   res,
			Req:   req}
		if handler, ok := a.eventListeners[event.Event.Type]; ok {
			handler(ctx)
		} else {
			a.errorHandling(res, req, fmt.Errorf("Unrecognized event: %s", event.Event.Type))
			return
		}
	} else {
		Response(&SlackContext{Res: res}, http.StatusUnauthorized, []byte("Unauthorized"), nil)
		return
	}
}

// commands - Slack App commands handler
func (a *SlackApp) commands(res http.ResponseWriter, req *http.Request) {
	bodyText, err := ioutil.ReadAll(req.Body)
	if err != nil {
		a.errorHandling(res, req, err)
		return
	}
	defer req.Body.Close()
	queries, err := url.ParseQuery(string(bodyText))
	if err != nil {
		a.errorHandling(res, req, err)
		return
	}
	isAuthorizedCaller := a.checkSlackSecret(req.Header.Get("X-Slack-Signature"), req.Header.Get("X-Slack-Request-TimeStamp"), string(bodyText))
	if isAuthorizedCaller {
		accessToken := a.GetTokenForWorkspace(queries.Get("team_id"))
		if accessToken == nil {
			a.errorHandling(res, req, fmt.Errorf("App not installed for workspace: %s", queries.Get("team_id")))
			return
		}
		ctx := &SlackContext{
			Body:  bodyText,
			Token: accessToken.Token,
			Res:   res,
			Req:   req}
		if handler, ok := a.cmds[queries.Get("command")]; ok {
			handler(ctx)
		} else {
			a.errorHandling(res, req, fmt.Errorf("Unrecognized command: %s", queries.Get("command")))
			return
		}
	} else {
		Response(&SlackContext{Res: res}, http.StatusUnauthorized, []byte("Unauthorized"), nil)
		return
	}
}

// CustomRoute - Add custom route
func (a *SlackApp) CustomRoute(pattern string, handler func(res http.ResponseWriter, req *http.Request)) {
	http.HandleFunc("/"+a.opts.Prefix+"/"+strings.Trim(pattern, "/"), handler)
}

func notFound(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		Response(&SlackContext{Res: res}, http.StatusNotFound, nil, nil)
		return
	}
	Response(&SlackContext{Res: res}, http.StatusOK, nil, nil)
}

// ServeApp - Listen and Serve App on desired port, callback can be nil
func (a *SlackApp) ServeApp(port uint16, cb func()) {
	if len(a.opts.Prefix) == 0 {
		panic(fmt.Sprintf("\x1b[31m%s\x1b[0m\n", "Slack App Route Prefix Cannot Be Empty"))
	}
	a.server = &http.Server{Addr: fmt.Sprintf(":%d", port)}
	http.HandleFunc("/", notFound)
	http.HandleFunc(fmt.Sprintf("/%s/events", a.opts.Prefix), a.events)
	http.HandleFunc(fmt.Sprintf("/%s/install", a.opts.Prefix), a.appInstall)
	http.HandleFunc(fmt.Sprintf("/%s/commands", a.opts.Prefix), a.commands)
	http.HandleFunc(fmt.Sprintf("/%s/", a.opts.Prefix), a.interactions)
	if cb != nil {
		go cb()
	}
	a.server.ListenAndServe()
}

// Close - Shutting down the server
func (a *SlackApp) Close(ctx context.Context) {
	if a.server != nil {
		if err := a.server.Shutdown(ctx); err != nil {
			panic(err)
		}
	}
}

// InitializeSlackApp - Return an instance of SlackApp
func InitializeSlackApp(opts *SlackAppOptions) SlackApp {
	app := SlackApp{
		opts: SlackAppOptions{
			Name:          opts.Name,
			TokensCache:   opts.TokensCache,
			Prefix:        opts.Prefix,
			ClientSecret:  opts.ClientSecret,
			ClientID:      opts.ClientID,
			SigningSecret: opts.SigningSecret},
		distCB:          nil,
		cmds:            make(map[string]func(ctx *SlackContext)),
		actionListeners: make(map[string]func(ctx *SlackContext)),
		submitListeners: make(map[string]func(ctx *SlackContext)),
		closeListeners:  make(map[string]func(ctx *SlackContext)),
		eventListeners:  make(map[string]func(ctx *SlackContext))}
	return app
}

// GetTokenForWorkspace - Finding the token for the corresponding workspace
func (a *SlackApp) GetTokenForWorkspace(workspace string) *SlackAuthToken {
	availableTokens := a.opts.TokensCache(workspace)
	var token *SlackAuthToken
	for _, t := range availableTokens {
		if t.Workspace == workspace {
			token = &t
			break
		}
	}
	return token
}

// Response - Send response back to slack
func Response(ctx *SlackContext, code int, message []byte, headers map[string]string) {
	for k, v := range headers {
		ctx.Res.Header().Set(k, v)
	}
	ctx.Res.WriteHeader(code)
	ctx.Res.Write(message)
}

// ConvertState - Convert unknown state to struct
func ConvertState(state ISlackBlockKitUI, dst interface{}) error {
	jsonView, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonView, dst)
	if err != nil {
		return err
	}
	return nil
}

// errorHandling - Error handling
func (a *SlackApp) errorHandling(res http.ResponseWriter, req *http.Request, err error) {
	if a.errorCB != nil {
		a.errorCB(res, req, err)
	} else {
		log.Println(err.Error())
		Response(&SlackContext{Res: res}, http.StatusUnauthorized, []byte(err.Error()), nil)
	}
	return
}

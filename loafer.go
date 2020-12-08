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

// OnCommand - Add handler to command
func (a *SlackApp) OnCommand(cmd string, handler func(ctx *SlackContext)) {
	a.cmds[cmd] = handler
}

// RemoveCommand - Remove a command to the app base on command
func (a *SlackApp) RemoveCommand(cmd string) {
	delete(a.cmds, cmd)
}

// OnAction - Add an action handler to the app base on action_id
func (a *SlackApp) OnAction(actionID string, handler func(ctx *SlackContext)) {
	a.actionListeners[actionID] = handler
}

// OnShortcut - Add an shortcut handler to the app base on callback_id
func (a *SlackApp) OnShortcut(callbackID string, handler func(ctx *SlackContext)) {
	a.shortcutListeners[callbackID] = handler
}

// OnViewSubmission - Add handler to view submission base on callback_id
func (a *SlackApp) OnViewSubmission(callbackID string, handler func(ctx *SlackContext)) {
	a.submitListeners[callbackID] = handler
}

// OnViewClose - Add handler to view close base on callback_id
func (a *SlackApp) OnViewClose(callbackID string, handler func(ctx *SlackContext)) {
	a.closeListeners[callbackID] = handler
}

// OnEvent - Add handler to events
func (a *SlackApp) OnEvent(eventType string, handler func(ctx *SlackContext)) {
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
	form := url.Values{
		"code":          []string{req.URL.Query().Get("code")},
		"client_id":     []string{a.opts.ClientID},
		"client_secret": []string{a.opts.ClientSecret}}
	resp, err := http.Post("https://slack.com/api/oauth.v2.access", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		a.errorHandling(res, req, err)
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&installResponse)
	if err != nil {
		a.errorHandling(res, req, err)
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
		}
	} else {
		Response(&SlackContext{Res: res}, http.StatusInternalServerError, []byte("Slack App Access Request is not Ok"), nil)
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
		accessToken := a.opts.TokensCache.Get(event.Team.ID)
		if len(accessToken) == 0 {
			a.errorHandling(res, req, fmt.Errorf("App is not installed to workspace: %s", event.Team.ID))
			return
		}
		ctx := &SlackContext{
			Body:      bodyText,
			Token:     accessToken,
			Workspace: event.Team.ID,
			Res:       res,
			Req:       req}
		switch Type := event.Type; Type {
		case "shortcut":
			callbackID := event.CallbackID
			if handler, ok := a.shortcutListeners[callbackID]; ok {
				handler(ctx)
			} else {
				a.errorHandling(res, req, fmt.Errorf("Unrecognized shortcut: %s", callbackID))
			}
		case "block_actions":
			action := event.Actions[0]
			if handler, ok := a.actionListeners[action.ActionID]; ok {
				handler(ctx)
			} else {
				a.errorHandling(res, req, fmt.Errorf("Unrecognized action: %s", action.ActionID))
			}
		case "view_submission":
			if handler, ok := a.submitListeners[event.View.CallbackID]; ok {
				handler(ctx)
			} else {
				a.errorHandling(res, req, fmt.Errorf("Unrecognized submission event from view: %s", event.View.CallbackID))
			}
		case "view_closed":
			if handler, ok := a.closeListeners[event.View.CallbackID]; ok {
				handler(ctx)
			} else {
				a.errorHandling(res, req, fmt.Errorf("Unrecognized closed event from view: %s", event.View.CallbackID))
			}
		default:
			a.errorHandling(res, req, fmt.Errorf("Unrecognized interaction type: %s", event.Type))
		}
	} else {
		Response(&SlackContext{Res: res}, http.StatusUnauthorized, []byte("Unauthorized"), nil)
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
		accessToken := a.opts.TokensCache.Get(event.TeamID)
		if len(accessToken) == 0 {
			a.errorHandling(res, req, fmt.Errorf("App is not installed for workspace: %s", event.TeamID))
			return
		}
		ctx := &SlackContext{
			Body:      body,
			Token:     accessToken,
			Workspace: event.TeamID,
			Res:       res,
			Req:       req}
		if handler, ok := a.eventListeners[event.Event.Type]; ok {
			handler(ctx)
		} else {
			a.errorHandling(res, req, fmt.Errorf("Unrecognized event: %s", event.Event.Type))
		}
	} else {
		Response(&SlackContext{Res: res}, http.StatusUnauthorized, []byte("Unauthorized"), nil)
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
		accessToken := a.opts.TokensCache.Get(queries.Get("team_id"))
		if len(accessToken) == 0 {
			a.errorHandling(res, req, fmt.Errorf("App not installed for workspace: %s", queries.Get("team_id")))
			return
		}
		ctx := &SlackContext{
			Body:      bodyText,
			Token:     accessToken,
			Workspace: queries.Get("team_id"),
			Res:       res,
			Req:       req}
		if handler, ok := a.cmds[queries.Get("command")]; ok {
			handler(ctx)
		} else {
			a.errorHandling(res, req, fmt.Errorf("Unrecognized command: %s", queries.Get("command")))
		}
	} else {
		Response(&SlackContext{Res: res}, http.StatusUnauthorized, []byte("Unauthorized"), nil)
	}
}

// CustomRoute - Add custom route
func (a *SlackApp) CustomRoute(pattern string, handler func(res http.ResponseWriter, req *http.Request)) {
	http.HandleFunc("/"+a.opts.Prefix+"/"+strings.Trim(pattern, "/"), handler)
}

// defaultRoute - handling defualt root path /
func defaultRoute(res http.ResponseWriter, req *http.Request) {
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
	http.HandleFunc("/", defaultRoute)
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
		opts:              *opts,
		distCB:            nil,
		cmds:              make(map[string]func(ctx *SlackContext)),
		actionListeners:   make(map[string]func(ctx *SlackContext)),
		submitListeners:   make(map[string]func(ctx *SlackContext)),
		closeListeners:    make(map[string]func(ctx *SlackContext)),
		eventListeners:    make(map[string]func(ctx *SlackContext)),
		shortcutListeners: make(map[string]func(ctx *SlackContext))}
	return app
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

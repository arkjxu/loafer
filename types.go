package loafer

import "net/http"

// ISlackBlockKitUI - Slack Generic UI Kit
type ISlackBlockKitUI interface{}

// TokensCache - Token cache interface type
type TokensCache interface {
	Get(workspace string) string
	Set(workspace string, token string)
	Remove(workspace string)
}

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

// SlackBlockText - Slack Text
type SlackBlockText struct {
	Type  string `json:"type,omitempty"`
	Text  string `json:"text,omitempty"`
	Emoji *bool  `json:"emoji,omitempty"`
}

// SlackDivider - Slack divider
type SlackDivider struct {
	Type string `json:"type"`
}

// SlackBlockAccessory - Slack Accessory
type SlackBlockAccessory struct {
	Type                 string             `json:"type,omitempty"`
	Title                *SlackBlockText    `json:"title,omitempty"`
	AltText              string             `json:"alt_text,omitempty"`
	Text                 *SlackBlockText    `json:"text,omitempty"`
	Value                string             `json:"value,omitempty"`
	IsMultiline          bool               `json:"multiline,omitempty"`
	MaxLength            uint16             `json:"max_length,omitempty"`
	MinQueryLength       uint16             `json:"min_query_length,omitempty"`
	Placeholder          *SlackBlockText    `json:"placeholder,omitempty"`
	ImageURL             string             `json:"image_url,omitempty"`
	ActionID             string             `json:"action_id,omitempty"`
	Options              []SlackInputOption `json:"options,omitempty"`
	InitialDate          string             `json:"initial_date,omitempty"`
	InitialTime          string             `json:"initial_time,omitempty"`
	InitialOption        *SlackInputOption  `json:"initial_option,omitempty"`
	InitialOptions       []SlackInputOption `json:"initial_options,omitempty"`
	InitialConversations []string           `json:"initial_conversations,omitempty"`
	InitialUser          string             `json:"initial_user,omitempty"`
	InitialUsers         []string           `json:"initial_users,omitempty"`
}

// SlackBlockTextFields - Slack Text fields
type SlackBlockTextFields struct {
	Type   string           `json:"type,omitempty"`
	Fields []SlackBlockText `json:"fields,omitempty"`
}

// SlackBlockSection - Slack message section
type SlackBlockSection struct {
	Type      string               `json:"type,omitempty"`
	Text      *SlackBlockText      `json:"text,omitempty"`
	Accessory *SlackBlockAccessory `json:"accessory,omitempty"`
}

// SlackModalSelect - Slack modal select
type SlackModalSelect struct {
	Type     string               `json:"type,omitempty"`
	Optional bool                 `json:"optional,omitempty"`
	BlockID  string               `json:"block_id,omitempty"`
	Element  *SlackBlockAccessory `json:"element,omitempty"`
	Label    *SlackBlockText      `json:"label,omitempty"`
}

// SlackBlockButton - Slack Button action
type SlackBlockButton struct {
	Type     string          `json:"type,omitempty"`
	Text     *SlackBlockText `json:"text,omitempty"`
	Value    string          `json:"value,omitempty"`
	ActionID string          `json:"action_id,omitempty"`
}

// SlackInputOption - Slack Select option
type SlackInputOption struct {
	Text  *SlackBlockText `json:"text,omitempty"`
	Value string          `json:"value,omitempty"`
}

// SlackBlockActions - Slack Actions
type SlackBlockActions struct {
	Type     string           `json:"type,omitempty"`
	Elements ISlackBlockKitUI `json:"elements,omitempty"`
}

// SlackUI - Slack UI
type SlackUI struct {
	Blocks ISlackBlockKitUI `json:"blocks,omitempty"`
}

// SlackInteractionUser - Slack Interaction User
type SlackInteractionUser struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	TeamID   string `json:"team_id,omitempty"`
}

// SlackInteractionContainer - Slack Interaction Container
type SlackInteractionContainer struct {
	Type        string `json:"type,omitempty"`
	MessageTS   string `json:"message_ts,omitempty"`
	ChannelID   string `json:"channel_id,omitempty"`
	IsEphemeral string `json:"is_ephemeral,omitempty"`
}

// SlackInteractionTeam - Slack Interaction Team
type SlackInteractionTeam struct {
	ID     string `json:"id,omitempty"`
	Domain string `json:"domain,omitempty"`
}

// SlackInteractionAction - Slack Interaction Action
type SlackInteractionAction struct {
	ActionID string          `json:"action_id,omitempty"`
	BlockID  string          `json:"block_id,omitempty"`
	Text     *SlackBlockText `json:"text,omitempty"`
	Value    string          `json:"value,omitempty"`
	Type     string          `json:"type,omitempty"`
	ActionTS string          `json:"action_ts,omitempty"`
}

// SlackInteractionEvent - Slack Interaction Event
type SlackInteractionEvent struct {
	Type        string                      `json:"type,omitempty"`
	User        *SlackInteractionUser       `json:"user,omitempty"`
	APIAppID    string                      `json:"api_app_id,omitempty"`
	Token       string                      `json:"token,omitempty"`
	Container   *SlackInteractionContainer  `json:"blocks,omitempty"`
	TriggerID   string                      `json:"trigger_id,omitempty"`
	Team        *SlackInteractionTeam       `json:"team,omitempty"`
	Channel     *SlackOauth2Team            `json:"channel,omitempty"`
	ResponseURL string                      `json:"response_url,omitempty"`
	Actions     []SlackInteractionAction    `json:"actions,omitempty"`
	Value       string                      `json:"value,omitempty"`
	State       map[string]ISlackBlockKitUI `json:"state,omitempty"`
	View        *SlackInteractionView       `json:"view,omitempty"`
	CallbackID  string                      `json:"callback_id,omitempty"`
	ActionID    string                      `json:"action_id,omitempty"`
	BlockID     string                      `json:"block_id,omitempty"`
	ActionTS    string                      `json:"action_ts,omitempty"`
}

// SlackInteractionView - Slack Interaction View
type SlackInteractionView struct {
	ID                 string                      `json:"id,omitempty"`
	TeamID             string                      `json:"team_id,omitempty"`
	Type               string                      `json:"type,omitempty"`
	Blocks             ISlackBlockKitUI            `json:"blocks,omitempty"`
	PrivateMetadata    string                      `json:"private_metadata,omitempty"`
	CallbackID         string                      `json:"callback_id,omitempty"`
	State              map[string]ISlackBlockKitUI `json:"state,omitempty"`
	Hash               string                      `json:"hash,omitempty"`
	Title              *SlackBlockText             `json:"title,omitempty"`
	Close              *SlackBlockText             `json:"close,omitempty"`
	Submit             *SlackBlockText             `json:"submit,omitempty"`
	ClearOnClose       bool                        `json:"clear_on_close,omitempty"`
	NotifyOnClose      bool                        `json:"notify_on_close,omitempty"`
	PreviousViewID     string                      `json:"previous_view_id,omitempty"`
	RootViewID         string                      `json:"root_view_id,omitempty"`
	AppID              string                      `json:"app_id,omitempty"`
	ExternalID         string                      `json:"external_id,omitempty"`
	AppInstalledTeamID string                      `json:"app_installed_team_id,omitempty"`
	BotID              string                      `json:"bot_id,omitempty"`
}

// SlackModal - Slack Modal
type SlackModal struct {
	Type            string           `json:"type,omitempty"`
	Title           *SlackBlockText  `json:"title,omitempty"`
	Submit          *SlackBlockText  `json:"submit,omitempty"`
	Close           *SlackBlockText  `json:"close,omitempty"`
	Blocks          ISlackBlockKitUI `json:"blocks,omitempty"`
	CallbackID      string           `json:"callback_id,omitempty"`
	NotifyOnClose   bool             `json:"notify_on_close,omitempty"`
	PrivateMetadata string           `json:"private_metadata,omitempty"`
}

// SlackInputElement - Slack Modal Plain text input
type SlackInputElement struct {
	Type             string               `json:"type,omitempty"`
	BlockID          string               `json:"block_id,omitempty"`
	IsDispatchAction bool                 `json:"dispatch_action,omitempty"`
	Element          *SlackBlockAccessory `json:"element,omitempty"`
	Label            *SlackBlockText      `json:"label,omitempty"`
}

// SlackActionSelect - Slack action select list
type SlackActionSelect struct {
	Type      string              `json:"type,omitempty"`
	Text      SlackBlockText      `json:"text,omitempty"`
	BlockID   string              `json:"block_id,omitempty"`
	Accessory SlackBlockAccessory `json:"accessory,omitempty"`
}

// SlackUsersQuery - Slack User query
type SlackUsersQuery struct {
	Ok    bool      `json:"ok"`
	User  SlackUser `json:"user"`
	Error string    `json:"error"`
}

// SlackUser - Slack User
type SlackUser struct {
	ID       string `json:"id"`
	TeamID   string `json:"team_id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	Color    string `json:"color"`
	RealName string `json:"real_name"`
	TZ       string `json:"tz"`
	TZLabel  string `json:"tz_label"`
	TZOffset int32  `json:"tz_offset"`
	Profile  struct {
		AvatarHash            string `json:"avatar_hash"`
		StatusText            string `json:"status_text"`
		StatusEmoji           string `json:"status_emoji"`
		RealName              string `json:"real_name"`
		DisplayName           string `json:"display_name"`
		RealNameNormalized    string `json:"real_name_normalized"`
		DisplayNameNormalized string `json:"display_name_normalized"`
		Email                 string `json:"email"`
		Image24               string `json:"image_24"`
		Image32               string `json:"image_32"`
		Image48               string `json:"image_48"`
		Image72               string `json:"image_72"`
		Image192              string `json:"image_192"`
		Image512              string `json:"image_512"`
		Team                  string `json:"team"`
	} `json:"profile"`
	IsAdmin           bool   `json:"is_admin"`
	IsOwner           bool   `json:"is_owner"`
	IsPrimaryOwner    bool   `json:"is_primary_owner"`
	IsRestricted      bool   `json:"is_restricted"`
	IsUltraRestricted bool   `json:"is_ultra_restricted"`
	IsBot             bool   `json:"is_bot"`
	Updated           uint32 `json:"updated"`
	IsAppUser         bool   `json:"is_app_user"`
	Has2FA            bool   `json:"has_2fa"`
}

// SlackAppOptions - Slack App options
type SlackAppOptions struct {
	Name          string      // Slack App name
	Prefix        string      // Prefix of routes
	TokensCache   TokensCache // List of available workspace tokens
	ClientSecret  string      // App client secret
	ClientID      string      // App client id
	SigningSecret string      // Signning secret
}

// SlackContext - Slack request context
type SlackContext struct {
	Body      []byte
	Token     string
	Workspace string
	Req       *http.Request
	Res       http.ResponseWriter
}

// SlackOauth2Team - Slack App Access Response Team
type SlackOauth2Team struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
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

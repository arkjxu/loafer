package loafer

// ISlackBlockKitUI - Slack Generic UI Kit
type ISlackBlockKitUI interface{}

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

// SlackBlockTextSection - Slack Text section
type SlackBlockTextSection struct {
	Type string          `json:"type,omitempty"`
	Text *SlackBlockText `json:"text,omitempty"`
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

// SlackInteractionChannel - Slack Interaction Channel
type SlackInteractionChannel struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
	Channel     *SlackInteractionChannel    `json:"channel,omitempty"`
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

// MakeSlackButton - Make a slack button
func MakeSlackButton(text string, value string, actionID string) SlackBlockButton {
	isEmojiSupported := true
	return SlackBlockButton{
		Type: "button",
		Text: &SlackBlockText{
			Type:  "plain_text",
			Text:  text,
			Emoji: &isEmojiSupported},
		Value:    value,
		ActionID: actionID,
	}
}

// MakeSlackTextSection - Make a text section (markdown)
func MakeSlackTextSection(text string) SlackBlockTextSection {
	return SlackBlockTextSection{
		Type: "section",
		Text: &SlackBlockText{
			Type: "mrkdwn",
			Text: text}}
}

// MakeSlackTextFieldsSection - Make a text section (markdown)
func MakeSlackTextFieldsSection(texts []string) SlackBlockTextFields {
	textFields := []SlackBlockText{}
	for _, t := range texts {
		textFields = append(textFields, SlackBlockText{Type: "mrkdwn", Text: t})
	}
	return SlackBlockTextFields{
		Type:   "section",
		Fields: textFields}
}

// MakeSlackModal - Make a slack button
func MakeSlackModal(title string, callbackID string, blocks ISlackBlockKitUI, submitText string, closeText string, notifyOnClose bool) SlackModal {
	return SlackModal{
		Type: "modal",
		Title: &SlackBlockText{
			Type: "plain_text",
			Text: title},
		Submit: &SlackBlockText{
			Type: "plain_text",
			Text: submitText},
		Close: &SlackBlockText{
			Type: "plain_text",
			Text: closeText},
		Blocks:        blocks,
		CallbackID:    callbackID,
		NotifyOnClose: notifyOnClose}
}

// MakeSlackActions - Make slack actions
func MakeSlackActions(actions ISlackBlockKitUI) SlackBlockActions {
	return SlackBlockActions{
		Type:     "actions",
		Elements: actions}
}

// MakeSlackModalTextInput - Make slack modal input field
func MakeSlackModalTextInput(label string, placeholder string, actionID string, isMultiline bool, isDispatch bool, maxLength uint16) SlackInputElement {
	isEmojiSupported := true
	return SlackInputElement{
		Type:             "input",
		BlockID:          actionID,
		IsDispatchAction: isDispatch,
		Element: &SlackBlockAccessory{
			Type:        "plain_text_input",
			IsMultiline: isMultiline,
			MaxLength:   maxLength,
			Placeholder: &SlackBlockText{
				Type: "plain_text",
				Text: placeholder},
			ActionID: actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackModalStaticSelectInput - Make slack static select input field
func MakeSlackModalStaticSelectInput(label string, placeholder string, options []SlackInputOption, initialOption *SlackInputOption, actionID string, isMulti bool, isOptional bool) SlackModalSelect {
	selectType := "static_select"
	isEmojiSupported := true
	if isMulti {
		selectType = "multi_static_select"
	}
	return SlackModalSelect{
		Type:     "input",
		BlockID:  actionID,
		Optional: isOptional,
		Element: &SlackBlockAccessory{
			Type: selectType,
			Placeholder: &SlackBlockText{
				Type:  "plain_text",
				Text:  placeholder,
				Emoji: &isEmojiSupported},
			Options:       options,
			InitialOption: initialOption,
			ActionID:      actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackModalExternalStaticSelectInput - Make slack external static select input field
func MakeSlackModalExternalStaticSelectInput(label string, placeholder string, initialOption *SlackInputOption, actionID string, isMulti bool, minQueryLength uint16, isOptional bool) SlackModalSelect {
	selectType := "external_select"
	isEmojiSupported := true
	minQueryLen := 1
	if minQueryLength > 1 {
		minQueryLen = int(minQueryLength)
	}
	if isMulti {
		selectType = "multi_external_select"
	}
	return SlackModalSelect{
		Type:     "input",
		BlockID:  actionID,
		Optional: isOptional,
		Element: &SlackBlockAccessory{
			Type: selectType,
			Placeholder: &SlackBlockText{
				Type:  "plain_text",
				Text:  placeholder,
				Emoji: &isEmojiSupported},
			MinQueryLength: uint16(minQueryLen),
			InitialOption:  initialOption,
			ActionID:       actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackBlockExternalStaticSelectInput - Make slack external static select input field
func MakeSlackBlockExternalStaticSelectInput(label string, placeholder string, initialOption *SlackInputOption, actionID string, isMulti bool, minQueryLength uint16) SlackBlockSection {
	selectType := "external_select"
	isEmojiSupported := true
	if isMulti {
		selectType = "multi_external_select"
	}
	return SlackBlockSection{
		Type: "section",
		Text: &SlackBlockText{
			Type: "mrkdwn",
			Text: label},
		Accessory: &SlackBlockAccessory{
			Type: selectType,
			Placeholder: &SlackBlockText{
				Type:  "plain_text",
				Text:  placeholder,
				Emoji: &isEmojiSupported},
			MinQueryLength: minQueryLength,
			ActionID:       actionID}}
}

// MakeSlackBlockButton - Make a slack button
func MakeSlackBlockButton(label string, text string, value string, actionID string) SlackBlockSection {
	isEmojiSupported := true
	return SlackBlockSection{
		Type: "section",
		Text: &SlackBlockText{
			Type: "mrkdwn",
			Text: label},
		Accessory: &SlackBlockAccessory{
			Type:     "button",
			ActionID: actionID,
			Value:    value,
			Text: &SlackBlockText{
				Type:  "plain_text",
				Text:  text,
				Emoji: &isEmojiSupported}}}
}

// MakeSlackModalMultiConversationSelectInput - Make slack multi convo select input field
func MakeSlackModalMultiConversationSelectInput(label string, placeholder string, initialConversations []string, actionID string) SlackModalSelect {
	isEmojiSupported := true
	return SlackModalSelect{
		Type:    "input",
		BlockID: actionID,
		Element: &SlackBlockAccessory{
			Type: "multi_conversations_select",
			Placeholder: &SlackBlockText{
				Type:  "plain_text",
				Text:  placeholder,
				Emoji: &isEmojiSupported},
			InitialConversations: initialConversations,
			ActionID:             actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackActionExternalStaticSelectInput - Make slack external static select input field
func MakeSlackActionExternalStaticSelectInput(label string, placeholder string, initialOption *SlackInputOption, actionID string, isMulti bool, minQueryLength uint16) SlackActionSelect {
	selectType := "external_select"
	isEmojiSupported := true
	minQueryLen := 1
	if minQueryLength > 1 {
		minQueryLen = int(minQueryLength)
	}
	if isMulti {
		selectType = "multi_static_select"
	}
	return SlackActionSelect{
		Type:    "section",
		BlockID: actionID,
		Text: SlackBlockText{
			Type: "mrkdwn",
			Text: label},
		Accessory: SlackBlockAccessory{
			Type: selectType,
			Placeholder: &SlackBlockText{
				Type:  "plain_text",
				Text:  placeholder,
				Emoji: &isEmojiSupported},
			ActionID:       actionID,
			MinQueryLength: uint16(minQueryLen),
			InitialOption:  initialOption}}
}

// MakeSlackModalMultiUserSelectInput - Make slack multi user select input field
func MakeSlackModalMultiUserSelectInput(label string, placeholder string, initialUsers []string, actionID string) SlackModalSelect {
	isEmojiSupported := true
	return SlackModalSelect{
		Type:    "input",
		BlockID: actionID,
		Element: &SlackBlockAccessory{
			Type: "multi_users_select",
			Placeholder: &SlackBlockText{
				Type:  "plain_text",
				Text:  placeholder,
				Emoji: &isEmojiSupported},
			InitialUsers: initialUsers,
			ActionID:     actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackModalUserSelectInput - Make slack user select input field
func MakeSlackModalUserSelectInput(label string, placeholder string, initialUser string, actionID string) SlackModalSelect {
	isEmojiSupported := true
	return SlackModalSelect{
		Type:    "input",
		BlockID: actionID,
		Element: &SlackBlockAccessory{
			Type: "users_select",
			Placeholder: &SlackBlockText{
				Type:  "plain_text",
				Text:  placeholder,
				Emoji: &isEmojiSupported},
			InitialUser: initialUser,
			ActionID:    actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackModalDatePickerInput - Make slack modal datepicker input field
func MakeSlackModalDatePickerInput(label string, placeholder string, initialDate string, actionID string) SlackInputElement {
	isEmojiSupported := true
	return SlackInputElement{
		Type: "input",
		Element: &SlackBlockAccessory{
			Type:        "datepicker",
			InitialDate: initialDate,
			Placeholder: &SlackBlockText{
				Type:  "plain_text",
				Text:  placeholder,
				Emoji: &isEmojiSupported},
			ActionID: actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackModalCheckboxesInput - Make slack modal checkboxes input field
func MakeSlackModalCheckboxesInput(label string, placeholder string, options []SlackInputOption, initialOptions []SlackInputOption, actionID string) SlackInputElement {
	isEmojiSupported := true
	return SlackInputElement{
		Type: "input",
		Element: &SlackBlockAccessory{
			Type:           "checkboxes",
			Options:        options,
			InitialOptions: initialOptions,
			ActionID:       actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackModalRadioInput - Make slack modal Radio input field
func MakeSlackModalRadioInput(label string, placeholder string, options []SlackInputOption, actionID string) SlackInputElement {
	isEmojiSupported := true
	return SlackInputElement{
		Type: "input",
		Element: &SlackBlockAccessory{
			Type:     "radio_buttons",
			Options:  options,
			ActionID: actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackInputOption - Make Slack input option
func MakeSlackInputOption(text string, value string) SlackInputOption {
	isEmojiSupported := true
	return SlackInputOption{
		Text: &SlackBlockText{
			Type:  "plain_text",
			Text:  text,
			Emoji: &isEmojiSupported},
		Value: value}
}

// MakeSlackModalTimePickerInput - Make slack modal time picker input field
func MakeSlackModalTimePickerInput(label string, placeholder string, initialTime string, actionID string) SlackInputElement {
	isEmojiSupported := true
	return SlackInputElement{
		Type: "input",
		Element: &SlackBlockAccessory{
			Type:        "timepicker",
			InitialTime: initialTime,
			ActionID:    actionID},
		Label: &SlackBlockText{Type: "plain_text", Text: label, Emoji: &isEmojiSupported}}
}

// MakeSlackHeader - Make a slack header section
func MakeSlackHeader(text string) SlackBlockTextSection {
	return SlackBlockTextSection{
		Type: "header",
		Text: &SlackBlockText{
			Type: "mrkdwn",
			Text: text}}
}

// MakeSlackDivider -Make a slack divider
func MakeSlackDivider() SlackDivider {
	return SlackDivider{
		Type: "divider"}
}

// MakeSlackContext - Make a slack context
func MakeSlackContext(text string) SlackBlockActions {
	return SlackBlockActions{
		Type: "context",
		Elements: []ISlackBlockKitUI{
			SlackBlockText{
				Type: "mrkdwn",
				Text: text}}}
}

// MakeSlackImage - Make a slack image section
func MakeSlackImage(title string, imageURL string, altText string) SlackBlockAccessory {
	return SlackBlockAccessory{
		Type: "image",
		Title: &SlackBlockText{
			Type: "mrkdwn",
			Text: title},
		ImageURL: imageURL,
		AltText:  altText}
}

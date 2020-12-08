package loafer

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
func MakeSlackTextSection(text string) SlackBlockSection {
	return SlackBlockSection{
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
func MakeSlackBlockExternalStaticSelectInput(label string, placeholder string, initialOptions []SlackInputOption, actionID string, isMulti bool, minQueryLength uint16) SlackBlockSection {
	selectType := "external_select"
	isEmojiSupported := true
	if isMulti {
		selectType = "multi_external_select"
	}
	view := SlackBlockSection{
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
	if isMulti {
		view.Accessory.InitialOptions = initialOptions
	} else if len(initialOptions) > 0 {
		view.Accessory.InitialOption = &initialOptions[0]
	}
	return view
}

// MakeSlackBlockStaticSelectInput - Make slack static select input field
func MakeSlackBlockStaticSelectInput(label string, placeholder string, options []SlackInputOption, initialOptions []SlackInputOption, actionID string, isMulti bool, minQueryLength uint16) SlackBlockSection {
	selectType := "static_select"
	isEmojiSupported := true
	if isMulti {
		selectType = "multi_static_select"
	}
	view := SlackBlockSection{
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
			Options:  options,
			ActionID: actionID}}
	if isMulti {
		view.Accessory.InitialOptions = initialOptions
	} else if len(initialOptions) > 0 {
		view.Accessory.InitialOption = &initialOptions[0]
	}
	return view
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
func MakeSlackHeader(text string) SlackBlockSection {
	return SlackBlockSection{
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

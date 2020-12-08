# Loafer

Note: 
```
This is a quick library to help with my development work at Nike Inc, please check the supported features before you use.
Feel free to clone and add your features if needed
```

A Simple Slack App library to help you quickly spin up Slack Apps that are capable of app distribution and signature checking without the hassle.

* All command, interactions, events endpoints are secured with slack signature checking.
* App distribution is automatically enabled, howered, you need to handle the token storage by using the `OnAppInstall` function
* You can enable custom routes if needed, such as for external select inputs for slack, or for your other service needs

## Supported Features
* Block Kit UI:
  * Button - Action
  * Text Section
  * Text Fields
  * Modal
  * Actions
  * Plain Text Input
  * (multi/single) Static Select Input
  * (multi/single) External Select Input
  * Multi-Conversation Select Input
  * (multi/single) User Select Input
  * Date Picker Input
  * Time Picker Input
  * Radio Input
  * Checkbox Input
  * Header
  * Context
  * Image
  * Divider
* Slack API:
  * Open View
  * Update View
  * Find User by Email
  * Find User by Slack ID
  * Update Message
  * Post Message
  * File upload


## Slack Context

All handlers are functions with the `loafer.SlackContext` parameter passed to it, and the format is as followed:
```golang
type SlackContext struct {
	Body      []byte              // Body of the request
	Token     string              // Token of the corresponding workspace
	Workspace string              // Workspace where event is coming from
	Req       *http.Request       // http request
	Res       http.ResponseWriter // http response
}
```

## To Start an App
```golang
package main

// TokenCache - Implementation of the loafer.TokensCache interface, this allows you to control how you store/access your tokens
type TokenCache struct {
	tokens map[string]string
}

// Get - Getting the token for the corresponding workspace
func (t *TokenCache) Get(workspace string) string {
	return t.tokens[workspace]
}

// Set - Setting the token for the corresponding workspace
func (t *TokenCache) Set(workspace string, token string) {
	t.tokens[workspace] = token
}

// Remove - Removing the token for the corresponding workspace
func (t *TokenCache) Remove(workspace string) {
	delete(t.tokens, workspace)
}

// main - entrypoint of program
func main {
  // Initialize an instance of the TokenCache
	myTokenCache := TokenCache{
		tokens: make(map[string]string)}
  
  // Set up the options for the slack app, clientId & client secret is not needed if you don't need app distribution
	opts := loafer.SlackAppOptions{
		Name:          "Dev Bot",
		Prefix:        "dev",
		TokensCache:   &myTokenCache,
		SigningSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ClientID:      "xxxxxxxxxxxx.xxxxxxxxxx",
		ClientSecret:  "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
    
  // Initialize your Slack App with the options
	app := loafer.InitializeSlackApp(&opts)
  
  // Add handler to command /coaching
	app.OnCommand("/coaching", handleDevCommand)
  
  // Serve the app on PORT 8080, you can use callback here if needed
	app.ServeApp(8080, nil)
}
```

App will now serve on the route:

http://0.0.0.0:8080/{prefix}/install - For app distribution
http://0.0.0.0:8080/{prefix}/events - For event subscription
http://0.0.0.0:8080/{prefix}/commands - For app commands
http://0.0.0.0:8080/{prefix}/interactions - For app interactions
http://0.0.0.0:8080/{prefix}/${custom_route_pattern} - For app custom routes

# API Reference

## App

### InitializeSlackApp(opts *SlackAppOptions) SlackApp

Returns:
* `app` SlackApp

Get an instance of a slack app with options:
- `Name` - Name of slack app
- `Prefix` - Prefix of slack app route
- `TokensCache` - Token cache that implemented the TokensCache interface from loafer
- `ClientSecret` - Client secret of slack app, used for app distribution
- `ClientID` - Client ID of slack app, used for app distribution
- `SigningSecret` - Signning secret for slack app, used for slack request verification

### ServeApp(port uint16, cb func())

Serve Slack app on port and cb when server first starts

### Close(ctx context.Context)

Shutdown Slack app

### OnCommand(cmd string, handler func(ctx *SlackContext))

Add handler to command

### RemoveCommand(cmd string)

Remove handler to command

### OnAction(actionID string, handler func(ctx *SlackContext))

Add handler to action

### OnShortcut(callbackID string, handler func(ctx *SlackContext))

Add handler to shortcut

### OnViewSubmission(callbackID string, handler func(ctx *SlackContext))

Add handler to view submission

### OnViewClose(callbackID string, handler, handler func(ctx *SlackContext))

Add handler to view close

### OnEvent(eventType string, handler func(ctx *SlackContext))

Add handler to view close

### OnError(handler func(res http.ResponseWriter, req *http.Request, err error))

Add handler to errors

### OnAppInstall(cb func(installRes *SlackOauth2Response, res http.ResponseWriter, req *http.Request) bool

Returns:
* `avoidDefaultPage` bool

handle the app distribute, once app distribution is successfull, it will call the `cb` function,
cb function should return a boolean value.
- `true` - you want to use your own html/redirection
- `false` - show default installation html page

### CustomRoute(pattern string, handler func(res http.ResponseWriter, req *http.Request))

Add handler to a custom pattern

### Response(ctx *SlackContext, code int, message []byte, headers map[string]string)

Response back to Slack

### ConvertState(state ISlackBlockKitUI, dst interface{})

Response interaction event state to your form state struct type

## Slack APIs

### OpenView(view SlackModal, triggerID string, token string) error

Returns:
* `err` error

Open a modal within slack

### UpdateView(view SlackInteractionView, viewID string, token string) error

Returns:
* `err` error

Updates a modal within slack

### FindUserByEmail(email string, token string) (*SlackUser, error)

Returns:
* `user` SlackUser
* `err` error

Find a user within the slack workspace

### FindUserByID(id string, token string) (*SlackUser, error)

Returns:
* `user` SlackUser
* `err` error

Find a user within the slack workspace

### PostMessage(channel string, blocks ISlackBlockKitUI, text string, token string) error

Returns:
* `err` error

Post a message to a slack channel/user/conversation

### UpdateMessage(channel string, ts string, blocks ISlackBlockKitUI, text string, token string) error

Returns:
* `err` error

Update a message of a slack channel/user/conversation given a time stamp of the original message

### FileUpload(channels []string, filename string, content string, filetype string, token string) error

Returns:
* `err` error

Upload a file to the Slack workspace and share to the list of channels/users/conversations

## Block Kit UIs

### MakeSlackButton(text string, value string, actionID string) SlackBlockButton

Returns:
* `button` SlackBlockButton

Make a Block Kit Button

### MakeSlackTextSection(text string) SlackBlockSection

Returns:
* `section` SlackBlockSection

Make a Block Kit text section

### MakeSlackTextFieldsSection(texts []string) SlackBlockTextFields

Returns:
* `textField` SlackBlockTextFields

Make a Block Kit text field section

### MakeSlackModal(title string, callbackID string, blocks ISlackBlockKitUI, submitText string, closeText string, notifyOnClose bool) SlackModal

Returns:
* `modal` SlackModal

Make a Block Kit modal

### MakeSlackActions(actions ISlackBlockKitUI) SlackBlockActions

Returns:
* `actions` SlackBlockActions

Make a Block Kit actions section

### MakeSlackModalTextInput(label string, placeholder string, actionID string, isMultiline bool, isDispatch bool, maxLength uint16) SlackInputElement

Returns:
* `input` SlackInputElement

Make a Block Kit modal plain text input

### MakeSlackModalStaticSelectInput(label string, placeholder string, options []SlackInputOption, initialOption *SlackInputOption, actionID string, isMulti bool, isOptional bool) SlackModalSelect

Returns:
* `select` SlackModalSelect

Make a Block Kit modal static select input

### MakeSlackModalExternalStaticSelectInput(label string, placeholder string, initialOption *SlackInputOption, actionID string, isMulti bool, minQueryLength uint16, isOptional bool) SlackModalSelect

Returns:
* `externalSelect` SlackModalSelect

Make a Block Kit modal external static select input

### MakeSlackBlockExternalStaticSelectInput(label string, placeholder string, initialOptions []SlackInputOption, actionID string, isMulti bool, minQueryLength uint16) SlackBlockSection

Returns:
* `blockSection` SlackBlockSection

Make a Block Kit block section external section

### MakeSlackBlockExternalStaticSelectInput(label string, placeholder string, initialOptions []SlackInputOption, actionID string, isMulti bool) SlackBlockSection

Returns:
* `blockSection` SlackBlockSection

Make a Block Kit block section static select section

### MakeSlackBlockButton(label string, text string, value string, actionID string) SlackBlockSection

Returns:
* `button` SlackBlockSection

Make a Block Kit block section with button

### MakeSlackModalMultiConversationSelectInput(label string, placeholder string, initialConversations []string, actionID string) SlackModalSelect

Returns:
* `select` SlackModalSelect

Make a Block Kit modal multi conversation select list

### MakeSlackActionExternalStaticSelectInput(label string, placeholder string, initialOption *SlackInputOption, actionID string, isMulti bool, minQueryLength uint16) SlackActionSelect

Returns:
* `select` SlackActionSelect

Make a Block Kit modal external select list

### MakeSlackModalMultiUserSelectInput(label string, placeholder string, initialUsers []string, actionID string) SlackModalSelect

Returns:
* `select` SlackModalSelect

Make a Block Kit modal multi user select list

### MakeSlackModalUserSelectInput(label string, placeholder string, initialUser string, actionID string) SlackModalSelect

Returns:
* `select` SlackModalSelect

Make a Block Kit modal user select list

### MakeSlackModalDatePickerInput(label string, placeholder string, initialDate string, actionID string) SlackInputElement

Returns:
* `picker` SlackInputElement

Make a Block Kit modal date picker

### MakeSlackModalCheckboxesInput(label string, placeholder string, options []SlackInputOption, initialOptions []SlackInputOption, actionID string) SlackInputElement

Returns:
* `picker` SlackInputElement

Make a Block Kit modal checkbox picker

### MakeSlackModalRadioInput(label string, placeholder string, options []SlackInputOption, actionID string) SlackInputElement

Returns:
* `picker` SlackInputElement

Make a Block Kit modal radio picker

### MakeSlackInputOption(text string, value string) SlackInputOption

Returns:
* `option` SlackInputOption

Make a Block Kit select option

### MakeSlackModalTimePickerInput(label string, placeholder string, initialTime string, actionID string) SlackInputElement

Returns:
* `picker` SlackInputElement

Make a Block Kit time picker

### MakeSlackHeader(text string) SlackBlockSection

Returns:
* `header` SlackBlockSection

Make a Block Kit header

### MakeSlackDivider() SlackDivider

Returns:
* `divider` SlackDivider

Make a Block Kit divider

### MakeSlackContext(text string) SlackBlockActions

Returns:
* `ctx` SlackBlockActions

Make a Block Kit context section

### MakeSlackImage(title string, imageURL string, altText string) SlackBlockAccessory

Returns:
* `img` SlackBlockAccessory

Make a Block Kit image section

# Contributing
Kevin Xu


# License
Apache 2.0

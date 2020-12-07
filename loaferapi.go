package loafer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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

// slackUserQuery - Calls Slack Users.info or Users.lookupByEmail API
func slackUserCall(uri string, token string) SlackUser {
	var userQuery SlackUsersQuery
	client := &http.Client{}
	r, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return userQuery.User
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return userQuery.User
	}
	text, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return userQuery.User
	}
	if !strings.Contains(string(text), "\"ok\":true") {
		return userQuery.User
	}
	err = json.NewDecoder(strings.NewReader(string(text))).Decode(&userQuery)
	if err != nil {
		return userQuery.User
	}
	return userQuery.User
}

// slackMessageCall - Calls Slack Chat.PostMessage API
func slackMessageCall(uri string, form url.Values, token string) bool {
	client := &http.Client{}
	r, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		return false
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return false
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return false
	}
	if !strings.Contains(string(bodyText), "\"ok\":true") {
		return false
	}
	return true
}

// OpenView - Open view in slack
func OpenView(view SlackModal, triggerID string, token string) bool {
	jsonView, err := json.Marshal(view)
	if err != nil {
		return false
	}
	form := url.Values{}
	form.Set("view", string(jsonView))
	form.Set("trigger_id", triggerID)
	client := &http.Client{}
	r, err := http.NewRequest("POST", "https://slack.com/api/views.open", strings.NewReader(form.Encode()))
	if err != nil {
		return false
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return false
	}
	text, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return false
	}
	if !strings.Contains(string(text), "\"ok\":true") {
		fmt.Println(string(text))
		return false
	}
	return true
}

// UpdateView - Update a view in slack
func UpdateView(view SlackInteractionView, viewID string, token string) bool {
	jsonView, err := json.Marshal(view)
	if err != nil {
		log.Printf("Unable to get JSON from view")
		return false
	}
	form := url.Values{}
	form.Set("view", string(jsonView))
	form.Set("view_id", viewID)
	client := &http.Client{}
	r, err := http.NewRequest("POST", "https://slack.com/api/views.update", strings.NewReader(form.Encode()))
	if err != nil {
		return false
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return false
	}
	text, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return false
	}
	if !strings.Contains(string(text), "\"ok\":true") {
		return false
	}
	return true
}

// FindUserByEmail - Finding slack user by email
func FindUserByEmail(email string, token string) SlackUser {
	foundUser := slackUserCall(fmt.Sprintf("https://slack.com/api/users.lookupByEmail?email=%s&token=%s", email, token), token)
	return foundUser
}

// FindUserByID - Finding slack user by id
func FindUserByID(id string, token string) SlackUser {
	foundUser := slackUserCall(fmt.Sprintf("https://slack.com/api/users.info?user=%s&token=%s", id, token), token)
	return foundUser
}

// UpdateMessage - Update a slack message
func UpdateMessage(channel string, ts string, blocks ISlackBlockKitUI, text string, token string) bool {
	var jsonBlocks []byte
	var err error
	if blocks != nil {
		jsonBlocks, err = json.Marshal(blocks)
		if err != nil {
			panic("Invalid JSON Block Object passed to UpdateMessage")
		}
	}
	form := url.Values{}
	form.Set("channel", channel)
	form.Set("ts", ts)
	if blocks != nil {
		form.Set("blocks", string(jsonBlocks))
	}
	form.Set("text", text)
	res := slackMessageCall("https://slack.com/api/chat.update", form, token)
	return res
}

// PostMessage - Post a message
func PostMessage(channel string, blocks ISlackBlockKitUI, text string, token string) bool {
	var jsonBlocks []byte
	var err error
	if blocks != nil {
		jsonBlocks, err = json.Marshal(blocks)
		if err != nil {
			panic("Invalid JSON Block Object passed to UpdateMessage")
		}
	}
	form := url.Values{}
	form.Set("channel", channel)
	if blocks != nil {
		form.Set("blocks", string(jsonBlocks))
	}
	form.Set("text", text)
	res := slackMessageCall("https://slack.com/api/chat.postMessage", form, token)
	return res
}

// FileUpload - Upload a file
func FileUpload(channels []string, filename string, content string, filetype string, token string) error {
	form := url.Values{}
	form.Set("content", content)
	form.Set("filename", filename)
	form.Set("filetype", filetype)
	form.Set("channels", strings.Join(channels, ","))
	client := &http.Client{}
	r, err := http.NewRequest("POST", "https://slack.com/api/files.upload", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	text, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return err
	}
	if !strings.Contains(string(text), "\"ok\":true") {
		return err
	}
	return nil
}

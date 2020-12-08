package loafer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// slackUserQuery - Calls Slack Users.info or Users.lookupByEmail API
func slackUserCall(uri string, token string) (*SlackUser, error) {
	var userQuery SlackUsersQuery
	client := &http.Client{}
	r, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	text, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return nil, err
	}
	if !strings.Contains(string(text), "\"ok\":true") {
		return nil, err
	}
	err = json.NewDecoder(strings.NewReader(string(text))).Decode(&userQuery)
	if err != nil {
		return nil, err
	}
	return &userQuery.User, nil
}

// slackMessageCall - Calls Slack Chat.PostMessage API
func slackMessageCall(uri string, form url.Values, token string) error {
	client := &http.Client{}
	r, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return err
	}
	if !strings.Contains(string(bodyText), "\"ok\":true") {
		return errors.New(string(bodyText))
	}
	return nil
}

// OpenView - Open view in slack
func OpenView(view SlackModal, triggerID string, token string) error {
	jsonView, err := json.Marshal(view)
	if err != nil {
		return err
	}
	form := url.Values{
		"view":       []string{string(jsonView)},
		"trigger_id": []string{triggerID}}
	client := &http.Client{}
	r, err := http.NewRequest("POST", "https://slack.com/api/views.open", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return err
	}
	if !strings.Contains(string(bodyText), "\"ok\":true") {
		fmt.Println(string(bodyText))
		return errors.New(string(bodyText))
	}
	return nil
}

// UpdateView - Update a view in slack
func UpdateView(view SlackInteractionView, viewID string, token string) error {
	jsonView, err := json.Marshal(view)
	if err != nil {
		return err
	}
	form := url.Values{
		"view":    []string{string(jsonView)},
		"view_id": []string{viewID}}
	client := &http.Client{}
	r, err := http.NewRequest("POST", "https://slack.com/api/views.update", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	if err != nil {
		return err
	}
	if !strings.Contains(string(bodyText), "\"ok\":true") {
		return errors.New(string(bodyText))
	}
	return nil
}

// FindUserByEmail - Finding slack user by email
func FindUserByEmail(email string, token string) (*SlackUser, error) {
	return slackUserCall(fmt.Sprintf("https://slack.com/api/users.lookupByEmail?email=%s&token=%s", email, token), token)
}

// FindUserByID - Finding slack user by id
func FindUserByID(id string, token string) (*SlackUser, error) {
	return slackUserCall(fmt.Sprintf("https://slack.com/api/users.info?user=%s&token=%s", id, token), token)
}

// UpdateMessage - Update a slack message
func UpdateMessage(channel string, ts string, blocks ISlackBlockKitUI, text string, token string) error {
	var jsonBlocks []byte
	var err error
	if blocks != nil {
		jsonBlocks, err = json.Marshal(blocks)
		if err != nil {
			return err
		}
	}
	form := url.Values{
		"channel": []string{channel},
		"ts":      []string{ts},
		"text":    []string{text}}
	if blocks != nil {
		form.Set("blocks", string(jsonBlocks))
	}
	return slackMessageCall("https://slack.com/api/chat.update", form, token)
}

// PostMessage - Post a message
func PostMessage(channel string, blocks ISlackBlockKitUI, text string, token string) error {
	var jsonBlocks []byte
	var err error
	if blocks != nil {
		jsonBlocks, err = json.Marshal(blocks)
		if err != nil {
			return err
		}
	}
	form := url.Values{
		"channel": []string{channel},
		"text":    []string{text}}
	if blocks != nil {
		form.Set("blocks", string(jsonBlocks))
	}
	return slackMessageCall("https://slack.com/api/chat.postMessage", form, token)
}

// FileUpload - Upload a file
func FileUpload(channels []string, filename string, content string, filetype string, token string) error {
	form := url.Values{
		"content":  []string{content},
		"filename": []string{filename},
		"filetype": []string{filetype},
		"channels": []string{strings.Join(channels, ",")}}
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

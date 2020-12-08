package main

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	loafer "github.com/arkjxu/loafer"
)

func handleDevCommand(ctx *loafer.SlackContext) {
	isEmojiSupported := true
	buttons := []loafer.SlackBlockButton{}
	buttons = append(buttons, loafer.SlackBlockButton{
		Type: "button",
		Text: &loafer.SlackBlockText{
			Type:  "plain_text",
			Text:  "Click Me",
			Emoji: &isEmojiSupported,
		},
		Value:    "Click",
		ActionID: "clicked_me",
	})
	actions := loafer.SlackBlockActions{
		Type:     "actions",
		Elements: buttons,
	}
	blocks := []loafer.ISlackBlockKitUI{}
	blocks = append(blocks, actions)
	ctx.Res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.Res).Encode(loafer.SlackUI{
		Blocks: blocks,
	})
}

type TokenCache struct {
	tokens map[string]string
}

func (t *TokenCache) Get(workspace string) string {
	return t.tokens[workspace]
}

func (t *TokenCache) Set(workspace string, token string) {
	t.tokens[workspace] = token
}

func (t *TokenCache) Remove(workspace string) {
	delete(t.tokens, workspace)
}

func TestRun(t *testing.T) {
	myTokenCache := TokenCache{
		tokens: make(map[string]string)}
	opts := loafer.SlackAppOptions{
		Name:          "Dev Bot",
		Prefix:        "dev",
		TokensCache:   &myTokenCache,
		SigningSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ClientID:      "xxxxxxxxxxxx.xxxxxxxxxx",
		ClientSecret:  "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
	app := loafer.InitializeSlackApp(&opts)
	app.OnCommand("/coaching", handleDevCommand)
	app.ServeApp(8080, func() {
		time.Sleep(3 * time.Second)
		timeOut, cancel := context.WithTimeout(context.Background(), 1000)
		defer cancel()
		app.Close(timeOut)
	})
}

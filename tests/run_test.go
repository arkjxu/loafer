package main

import (
	"context"
	"encoding/json"
	"testing"

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

func tokenCache(_ string) []loafer.SlackAuthToken {
	tokens := []loafer.SlackAuthToken{}
	return tokens
}

func TestRun(t *testing.T) {
	opts := loafer.SlackAppOptions{
		Name:          "Dev Bot",
		Prefix:        "dev",
		TokensCache:   tokenCache,
		SigningSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ClientID:      "xxxxxxxxxxxx.xxxxxxxxxx",
		ClientSecret:  "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
	app := loafer.InitializeSlackApp(&opts)
	app.OnCommand("/coaching", handleDevCommand)
	app.ServeApp(8080, func() {
		timeOut, cancel := context.WithTimeout(context.Background(), 1000)
		defer cancel()
		app.Close(timeOut)
	})
}

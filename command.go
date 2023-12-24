package main

import (
	"strings"
	"encoding/json"
	"fmt"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

const (
	cmdTrigger = "createtask"
	webhookUrl = "https://influencersu1.uforce.pro/test_webhook"
)

// OnActivate register the plugin command
func (p *Plugin) OnActivate() error {
	return p.API.RegisterCommand(&model.Command{
		Trigger:          cmdTrigger,
		Description:      "Description",
		DisplayName:      "DisplayName",
		AutoComplete:     true,
		AutoCompleteDesc: "AutoCompleteDesc",
		AutoCompleteHint: "AutoCompleteHint",
	})
}

// ExecuteCommand post a custom-type spoiler post, the webapp part of the plugin will display it right
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	rawText := strings.TrimSpace((strings.Replace(args.Command, "/"+cmdTrigger, "", 1)))

	response := sendPostRequest(webhookUrl, map[string]string{
		"root_id":args.RootId,
		"raw_text":rawText,
	})
	jsonStr, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return nil, nil
	}

	_, appErr := p.API.CreatePost(&model.Post{
		UserId:    args.UserId,
		ChannelId: args.ChannelId,
		RootId:    args.RootId,
		Message: string(jsonStr),
	})
	if appErr != nil {
		fmt.Println("Error creating a new Post:", err)
		return nil, appErr
	}

	return &model.CommandResponse{}, nil
}

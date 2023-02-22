package main

import (
	"fmt"
	"path"
	"runtime"
	"slack-file-bot/config"

	"github.com/slack-go/slack"
)

func main() {
	config, err := config.LoadConfig(getSourcePath())
	if err != nil {
		panic(err)
	}
	fmt.Println(config)

	api := slack.New(config.Slack.Bot)

	channelArr := []string{config.Slack.ChannelId}
	fileArr := []string{"text.txt"}
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Println("err in upload :", err)
			continue
		}
		fmt.Printf("Uploaded file %s and url is %s\n", file.Name, file.URLPrivate)
	}
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

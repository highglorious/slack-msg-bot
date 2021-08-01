package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	type Channel struct {
		Id   string `json:"channel"`
		Text string `json:"text"`
	}

	type Channels struct {
		Channels []Channel `json:"channels"`
	}

	type Token struct {
		Token string `json:"bot_token"`
	}

	url := "https://slack.com/api/chat.postMessage"

	jsonMessage, err := os.Open("example.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonMessage.Close()

	byteMessage, err := ioutil.ReadAll(jsonMessage)
	if err != nil {
		fmt.Println(err)
	}

	var channels Channels
	var token Token

	json.Unmarshal(byteMessage, &token)

	fmt.Println("Bot token: " + token.Token)

	json.Unmarshal(byteMessage, &channels)

	// Publishing a message
	client := http.Client{Timeout: time.Second * 10}

	for i := 0; i < len(channels.Channels); i++ {

		if !((channels.Channels[i].Id == "") || (channels.Channels[i].Text == "")) {

			fmt.Println("Channel: " + channels.Channels[i].Id)
			fmt.Println("Text: " + channels.Channels[i].Text)

			chanMessage, _ := json.Marshal(channels.Channels[i])
			request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(chanMessage))
			request.Header.Set("Content-type", "application/json")
			request.Header.Set("Authorization", "Bearer "+token.Token)
			response, err := client.Do(request)

			if err != nil {
				panic(err)
			}
			//fmt.Println(response)
			//fmt.Println("Message sended ok")
			defer response.Body.Close()

		} else {
			fmt.Println("Bad JSON on channel: ", i, ", message not sended")
		}

	}

}

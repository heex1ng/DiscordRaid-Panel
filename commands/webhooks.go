package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SpamWebhooks() {
	var webhook string
	var message string

	fmt.Print("Webhook: ")
	fmt.Scan(&webhook)

	fmt.Print("Message: ")
	fmt.Scan(&message)

	message_json := map[string]interface{}{
		"content": message,
	}
	toBytes, err := json.Marshal(message_json)
	if err != nil {
		panic(err)
	}
	for {
		site, err := http.Post(webhook, "application/json", bytes.NewBuffer(toBytes))
		if err != nil {
			panic(err)
		}

		var result map[string]interface{}
		json.NewDecoder(site.Body).Decode(&result)
		time.Sleep(3 * time.Millisecond)
	}
}

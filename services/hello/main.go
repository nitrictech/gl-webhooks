package main

import (
	"encoding/json"
	"fmt"

	"github.com/nitrictech/go-sdk/nitric"
	"github.com/nitrictech/go-sdk/nitric/apis"
)

func main() {
	api := nitric.NewApi("main")

	api.Post("/webhook", func(ctx *apis.Ctx) {
		data := ctx.Request.Data()
		payload := map[string]interface{}{}

		err := json.Unmarshal(data, &payload)
		if err != nil {
			fmt.Printf("Error parsing JSON payload: %v", err)
		} else {
			fmt.Printf("Received webhook payload: %+v\n", payload)
		}

	})

	nitric.Run()
}

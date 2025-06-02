package controllers

import (
	"bufio"
	"fmt"
	"log"
	"pub-sub-go/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func Notification(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		publSub := configs.RDS.Subscribe(configs.RDS_CTX, configs.REDIS_CHANNEL_NOTIFICATION)
		defer publSub.Close()

		for {
			msg, err := publSub.ReceiveMessage(configs.RDS_CTX)
			if err != nil {
				log.Println("pubSub.ReceiveMessage(CTX):", err)
			} else {
				log.Println("Message:", msg.Payload)
			}

			fmt.Fprintf(w, "data: %s\n\n", msg.Payload)

			err = w.Flush()
			if err != nil {
				fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
				break
			}
		}
	}))

	return nil // if no error, return nil to Fiber to end the reques
}

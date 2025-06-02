package controllers

import (
	"fmt"
	"log"
	"pub-sub-go/configs"

	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	msg := fmt.Sprintf("User dengan IP %v mengakses halaman home", c.IP())
	err := configs.RDS.Publish(
		configs.RDS_CTX,
		configs.REDIS_CHANNEL_NOTIFICATION,
		msg).Err()

	if err != nil {
		log.Println("Error:", err)
	}

	return c.Render("index", fiber.Map{
		"Title": "Home Page",
	})
}

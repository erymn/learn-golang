package main

import (
	"log"
	"pub-sub-go/configs"
	"pub-sub-go/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// init redis client, menghubungkan aplikasi ke redis server
	configs.InitRedisClient()

	// cek apakan ada error ketika connect ke redis server
	if err := configs.RDS.Ping(configs.RDS_CTX).Err(); err != nil {
		log.Println("Failed to connect to redis server")
		panic(err)
	} else {
		log.Println("Connected to redis server")
		println("Connected to redis server")
	}

	// inisialisasi variable yang berguna untuk mengkonfirmasi direktori untuk render halaman web
	// pada kasus ini, direktori views dan render html
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", controllers.HomePage)
	app.Get("/notification", controllers.Notification)

	app.Listen(":9090")
}

package main

import (
	"github.com/joho/godotenv"
	"log"

	"rtm/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"rtm/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_ = godotenv.Load(".env")
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}

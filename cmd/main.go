package main

import (
	"log"
	postsRepo "refound-introspection/modules/posts/repo"
	"refound-introspection/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := router.New()

	api := app.Group("/api/v1")

	// ping
	api.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("all good")
	})

	api.Get("/posts", func(ctx *fiber.Ctx) error {
		repo, repoErr := postsRepo.NewLensPostsRepo()
		if repoErr != nil {
			return ctx.SendStatus(500)
		}

		posts, getErr := repo.Get(10, 0)
		if getErr != nil {
			return ctx.SendStatus(500)
		}

		return ctx.Status(200).JSON(posts)
	})

	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(404)
	})

	log.Fatal(app.Listen(":80"))
}

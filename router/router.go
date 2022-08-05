package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func New() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	app.Use(logger.New(logger.Config{
		Next: func(ctx *fiber.Ctx) bool {
			return ctx.Path() == "/metrics" || ctx.Path() == "/favicon.ico"
		},
		Format:     "${time} ${latency} [${ip}]:${port} ${pid} ${locals:requestid}\n${status}-${method}${path}\nsent:${bytesSent} received:${bytesReceived}\n>>${error}",
		TimeFormat: "2006-01-02_15:04:05",
		TimeZone:   "America/Los_Angeles",
	}))
	app.Use(requestid.New())
	app.Use(recover.New())

	app.Get("/metrics", monitor.New())

	return app
}

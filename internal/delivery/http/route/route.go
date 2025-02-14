package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/delivery/http/controller"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *controller.UserController
	PostController *controller.PostController
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthenticatedRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/v1/user/register", c.UserController.Register)
	c.App.Post("/v1/user/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthenticatedRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Post("/v1/post", c.PostController.CreatePost)
	// c.App.Put("/api/users", c.UserController.Update)

	// c.App.Get("/api/balance", c.BalanceController.BalanceInquiry)
}

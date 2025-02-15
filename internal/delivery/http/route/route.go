package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkirmdhnnn/segokucing-be/internal/delivery/http/controller"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *controller.UserController
	PostController    *controller.PostController
	CommentController *controller.CommentController
	FriendController  *controller.FriendController
	AuthMiddleware    fiber.Handler
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

	c.App.Post("/v1/post/comment", c.CommentController.CreateComment)

	c.App.Post("/v1/friend", c.FriendController.AddFriend)
}

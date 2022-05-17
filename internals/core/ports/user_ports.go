package ports

import "github.com/gofiber/fiber/v2"

type UserService interface {
	Login(email string, password string) error
	Register(email string, password string, passwordConfirmation string) error
}

type UserRepository interface {
	Login(email string, password string) error
	Register(emil string, password string) error
}

type UserHandlers interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

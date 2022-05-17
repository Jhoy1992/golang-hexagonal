package handlers

import (
	"hexagonal/internals/core/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService ports.UserService
}

var _ ports.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) Login(c *fiber.Ctx) error {
	user := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	log.Println(user)

	err := h.userService.Login(user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (h *UserHandlers) Register(c *fiber.Ctx) error {
	newUser := struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}

	if err := c.BodyParser(&newUser); err != nil {
		return err
	}

	err := h.userService.Register(newUser.Email, newUser.Password, newUser.ConfirmPassword)

	if err != nil {
		return err
	}

	return nil
}

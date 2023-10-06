package handlers

import (
	"github.com/MerlinEmris/go-fiber-first/config"
	"github.com/MerlinEmris/go-fiber-first/entities"
	"github.com/gofiber/fiber/v2"
)

func GetDogs(c *fiber.Ctx) error {
	var dogs []entities.Dog

	config.Database.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

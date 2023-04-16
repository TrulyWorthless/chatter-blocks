package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trulyworthless/chatter-blocks/pkg/web3"
)

func GetBalance(c *fiber.Ctx) error {
	balance, err := web3.GetBalance(c.Params("holder"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "balance not found", "data": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "ETH balance", "data": balance})
}

// func GetTokenBalance(c *fiber.Ctx) error {
// 	balance, err := web3.GetTokenBalance(c.Params("address"), c.Params("holder"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "balance not found", "data": err})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "erc20 balance", "data": balance})
// }

// func SubmitTransaction(c *fiber.Ctx) error {}

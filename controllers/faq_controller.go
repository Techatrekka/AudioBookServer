package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/responses"

	"github.com/gofiber/fiber/v2"
)

func GetFaq(c *fiber.Ctx) error {
	file, err := os.Open("FAQ.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err}})
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err}})
	}

	var FAQ []map[string]interface{}
	err = json.Unmarshal(data, &FAQ)
	if err != nil {
		fmt.Println("Error unmarshelling file:", err)
		// Handle error
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": FAQ}})
}

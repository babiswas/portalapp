package RoleController

import (
	database "portalapp/Database"
	testCaseModel "portalapp/Model/TestCaseModel"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllTestCases(ctx *fiber.Ctx) error {
	var testCases []testCaseModel.TestCase
	database.DB.Find(&testCases)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"test_cases": testCases})
}

func CreateTestCase(ctx *fiber.Ctx) error {

	var body struct {
		TestCase  string `json:"testcase"`
		Feature   string `json:"feature"`
		Status    bool   `json:"status"`
		ProjectID int    `json:"project_id"`
	}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "InvalidData"})
	}

	testCaseObj := testCaseModel.TestCase{TestCase: body.TestCase, Feature: body.Feature, Status: body.Status, ProjectID: body.ProjectID, CreatedAt: time.Time{}, UpdatedAt: time.Time{}}
	result := database.DB.Create(&testCaseObj)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create user."})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Sucessfully created a testcase."})
}

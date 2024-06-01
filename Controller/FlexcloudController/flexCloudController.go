package FlexCloudController

import (
	database "portalapp/Database"
	flexmodel "portalapp/Model/FlexCloudModel"

	"github.com/gofiber/fiber/v2"
)

func CreateFlexCloud(ctx *fiber.Ctx) error {

	var body struct {
		FlexCloudName string `json:"flexcloud_name"`
		Status        bool   `json:"availability_status"`
		Feature       string `json:"feature_deployed"`
	}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "InvalidData"})
	}

	flexcloudObj := flexmodel.FlexCloud{
		FlexCloudName: body.FlexCloudName,
		Status:        body.Status,
		Feature:       body.Feature,
	}

	result := database.DB.Create(&flexcloudObj)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create user."})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Sucessfully created a testcase."})
}

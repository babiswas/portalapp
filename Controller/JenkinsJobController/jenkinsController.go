package JenkinsJobController

import (
	"log"
	database "portalapp/Database"
	jenkinsModel "portalapp/Model/JenkinsJobModel"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateJenkinsJob(ctx *fiber.Ctx) error {

	var body struct {
		JobName          string `json:"jobname"`
		BuildNumber      string `json:"build_number"`
		CompletionStatus string `json:"completion_status"`
		ReportLink       string `json:"report_link"`
		ProjectName      string `json:"project_name"`
		FeatureName      string `json:"feature_name"`
		FlexCloudID      string `json:"flex_cloud_id"`
	}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "InvalidData"})
	}

	jenkinsJobObj := jenkinsModel.JenkinsJobStatus{
		JobName:          body.JobName,
		BuildNumber:      body.BuildNumber,
		CompletionStatus: body.CompletionStatus,
		ReportLink:       body.ReportLink,
		ProjectName:      body.ProjectName,
		FeatureName:      body.FeatureName,
		FlexCloudID:      body.FlexCloudID,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}

	result := database.DB.Create(&jenkinsJobObj)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create user."})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Sucessfully created a testcase."})
}

func GetAllJobs(ctx *fiber.Ctx) error {
	var jobs []jenkinsModel.JenkinsJobStatus
	database.DB.Find(&jobs)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"jobs": jobs})
}

func GetAllJenkinsJob(ctx *fiber.Ctx) error {
	log.Println("Providing all jenkins job.")
	return ctx.Render("jenkinsjob", fiber.Map{})
}

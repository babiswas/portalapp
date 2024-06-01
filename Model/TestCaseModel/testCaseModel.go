package TestCaseModel

import (
	"time"

	project_model "portalapp/Model"

	"gorm.io/gorm"
)

type TestCase struct {
	gorm.Model
	TestCase  string    `json:"testcase"`
	Feature   string    `json:"feature"`
	Status    bool      `json:"status"`
	ProjectID int       `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Project   project_model.Project
}

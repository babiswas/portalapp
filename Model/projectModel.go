package Model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
}

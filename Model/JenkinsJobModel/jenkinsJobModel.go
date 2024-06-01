package JenkinsJobModel

import (
	flex "portalapp/Model/FlexCloudModel"
	"time"

	"gorm.io/gorm"
)

type JenkinsJobStatus struct {
	gorm.Model
	JobName          string         `json:"jobname"`
	BuildNumber      string         `json:"build_number"`
	CompletionStatus string         `json:"completion_status"`
	ReportLink       string         `json:"report_link"`
	ProjectName      string         `json:"project_name"`
	FeatureName      string         `json:"feature_name"`
	FlexCloudID      string         `json:"flex_cloud_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	FlexCloud        flex.FlexCloud `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

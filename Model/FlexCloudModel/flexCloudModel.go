package FlexCloudModel

import (
	"gorm.io/gorm"
)

type FlexCloud struct {
	gorm.Model
	FlexCloudName string `json:"flexcloud_name"`
	Status        bool   `json:"availability_status"`
	Feature       string `json:"feature_deployed"`
}

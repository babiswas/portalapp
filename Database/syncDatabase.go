package Database

import usermodel "portalapp/Model/UserModel"

func SyncDatabase() {
	DB.AutoMigrate(&usermodel.User{})
}

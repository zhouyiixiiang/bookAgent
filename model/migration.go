package model

func migration() {
	DB.AutoMigrate(&BookInLocal{})
}

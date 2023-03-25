package initializers

import (
	"VekterBackend/src/models"
	"fmt"
)

func Migrate() {
	err := MigrateUser()
	if err != nil {
		fmt.Print(err)
	}
}

func MigrateUser() error {
	err := DB.AutoMigrate(&models.Users{})
	return err
}

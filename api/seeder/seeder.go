
package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/NEHA20-1992/task_1/api/model"
)

var users = []model.UserData{
	model.UserData{
		FirstName: "Steven",
		LastName: "victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	model.UserData{
		FirstName: "Martin",
		LastName: "Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
	model.UserData{
		FirstName: "Siva",
		LastName:"Mathur",
		Email:    "nihu.y.20@gmail.com",
		Password: "hello123",
	},
}


func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&model.UserData{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&model.UserData{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	
	for i, _ := range users {
		err = db.Debug().Model(&model.UserData{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}
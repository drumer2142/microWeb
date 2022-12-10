package migrations

import (
	"log"

	"github.com/drumer2142/microWeb/src/api/database"
	"github.com/drumer2142/microWeb/src/api/models"
)

var err error

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Println("\nFailed to connect to db....\n", err.Error())
	}
	defer db.Close()

	// Drop table if exits
	err = db.Debug().DropTableIfExists(&models.Website{}).Error
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	err = db.Debug().AutoMigrate(&models.Website{}).Error
	if err != nil {
		log.Fatal(err)
	}

	// insert demo values to the db
	for i, _ := range websites {

		err = db.Debug().Model(&models.Website{}).Create(&websites[i]).Error
		if err != nil {
			log.Fatal(err)
		}

	}

}

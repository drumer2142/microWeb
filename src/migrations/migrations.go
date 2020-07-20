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
  err = db.Debug().DropTableIfExists(&models.Site{}).Error
  if err != nil {
    log.Fatal(err)
  }

  // Migrate the schema
  err = db.Debug().AutoMigrate(&models.Site{}).Error
  if err != nil {
    log.Fatal(err)
  }

  // insert demo values to the db
  for i, _ := range sites {

    err = db.Debug().Model(&models.Site{}).Create(&sites[i]).Error
    if err != nil {
      log.Fatal(err)
    }

  }

}

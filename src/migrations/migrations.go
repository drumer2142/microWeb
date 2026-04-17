package migrations

import (
	"log"
	"os"

	"github.com/drumer2142/microWeb/src/api/database"
	"github.com/drumer2142/microWeb/src/api/models"
	"gorm.io/gorm"
)

// Run applies schema changes and optional demo seeding. Uses the shared DB from database.Init().
func Run() {
	db := database.Get()
	if db == nil {
		log.Fatal("migrations: database not initialized")
	}

	reset := os.Getenv("RESET_DB_ON_STARTUP") == "true"
	if reset {
		if err := db.Migrator().DropTable(&models.Site{}); err != nil {
			log.Fatal(err)
		}
	}

	if err := db.AutoMigrate(&models.Site{}); err != nil {
		log.Fatal(err)
	}

	if reset {
		seed(db)
		return
	}

	if os.Getenv("SEED_DEMO_DATA") == "true" {
		var n int64
		if err := db.Model(&models.Site{}).Count(&n).Error; err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			seed(db)
		}
	}
}

func seed(db *gorm.DB) {
	for i := range sites {
		if err := db.Create(&sites[i]).Error; err != nil {
			log.Fatal(err)
		}
	}
}

package connection

import (
	"expense-tracker-api/domain"
	"expense-tracker-api/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase(conf config.Database) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Name,
		conf.Port,
		conf.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	err = db.AutoMigrate(&domain.Expense{})
	if err != nil {
		log.Fatal("failed to auto migrate database: ", err)
	}

	return db
}

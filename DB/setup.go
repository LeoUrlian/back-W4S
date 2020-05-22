//Database connections with gorm
package DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"w4s/models"
)

func SetupModels() *gorm.DB {
	/*db, err := gorm.Open("mysql",
		""+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/w4s?charset=utf8&parseTime=True&loc=Local")*/

	db, err := gorm.Open("postgres", "host="+os.Getenv("DATABASE_URL")+" port="+os.Getenv("DB_PORT")+
		" user="+os.Getenv("DB_USER")+" dbname="+os.Getenv("DB_NAME") +" sslmode=disable"+" password="+os.Getenv("DB_PASSWORD")+"")

	​



	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.UserAccountBadListToken{})
	db.AutoMigrate(&models.Table{})
	db.AutoMigrate(&models.Picture{})
	db.AutoMigrate(&models.OtherLinks{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Profile{})
	db.AutoMigrate(&models.TypeofTable{})
	db.AutoMigrate(&models.RpgSystem{})
	db.AutoMigrate(&models.LogoffListTokens{})
	db.AutoMigrate(&models.User{})

	return db
}
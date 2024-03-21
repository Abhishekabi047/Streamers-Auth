package db

import (
	"fmt"
	"log"
	"service/pkg/config"
	models "service/pkg/models/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(c *config.Config) (*gorm.DB,error) {
	psqlinfo:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",c.DB_host,c.DB_username,c.DB_password,c.DB_name,c.DB_port)
	db,err:=gorm.Open(postgres.Open(psqlinfo),&gorm.Config{})
	if err != nil{
		log.Fatalln(err)
	}
	db.AutoMigrate(models.Channel{},models.Signup{},models.OtpKey{})
	return db,err
}
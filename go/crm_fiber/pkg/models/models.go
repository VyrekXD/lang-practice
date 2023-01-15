package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func init() {
	var err error

	db, err = gorm.Open(sqlite.Open("database/lead.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to load database: %v", err)
	}

	db.AutoMigrate(&Lead{})

	log.Println("Database connected and migrated")
}

func (l *Lead) Create() Lead {
	db.Create(l)

	return *l
}

func GetAllLeads() []Lead {
	leads := []Lead{}
	db.Find(&leads)
	return leads
}

func GetLead(id string) *Lead {
	lead := &Lead{}
	db.Find(&lead, id)
	return lead
}

func DeleteLead(id string) {
	db.Delete(&Lead{}, id)
}

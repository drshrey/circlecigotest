package main

import (
	"fmt"
	"github.com/drshrey/circlecigotest/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 4, add(2, 2), "2 +2 = 4")
}

func TestMakeOther(t *testing.T) {
	err := migrate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		"localhost",
		"5432",
		"drshrey",
		"testdb",
		"disable",
		"",
	))
	if err != nil {
		fmt.Println("main: error could not connect to postgres")
		return
	}

	// create entry
	entry := models.Entry{
		Property: false,
	}

	db.Create(&entry)
	assert.Equal(t, false, entry.Property)
	makeOther(db, &entry)
	assert.Equal(t, true, entry.Property)
}

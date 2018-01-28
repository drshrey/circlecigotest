package main

import (
	"fmt"
	"github.com/drshrey/circlecigotest/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func add(x int, y int) int {
	return x + y
}

func migrate() error {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		"localhost",
		"5432",
		"drshrey",
		"testdb",
		"disable",
		"",
	))
	if err != nil {
		return err
	}

	db.AutoMigrate(
		&models.Entry{},
	)
	return nil
}

func makeOther(db *gorm.DB, entry *models.Entry) {
	entry.Property = !entry.Property
	db.Save(entry)
}

func subtract(x int, y int) int {
	return x - y
}

func main() {
	fmt.Printf("2 + 2 = %d\n", add(2, 2))
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

	fmt.Println(entry.Property)
	makeOther(db, &entry)
	fmt.Println("changed", entry.Property)

	// rows := []models.Entry{}
	// if db.Find(&rows).RecordNotFound() {
	// 	fmt.Println("no records found")
	// } else {
	// 	fmt.Println(rows)
	// }

	fmt.Printf("2 - 2 = %d\n", subtract(2, 2))
}

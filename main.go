package main

import (
	"database/sql"
	"main/app"
	"main/config"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := config.GetConfig()
	
	// initDatabase(*config)
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}

func initDatabase(c config.Config) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.DB.Username,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.Name,
		c.DB.Charset)

	db, err := sql.Open(c.DB.Dialect, dbUri)
	// db, err := sql.Open(c.DB.Dialect, dbUri)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + c.DB.Name + " CHARACTER SET " + c.DB.Charset)

	if err != nil {
		panic(fmt.Sprintf("Could not create database: %v", err))
	}

	fmt.Println("Database initialized successfully")
}

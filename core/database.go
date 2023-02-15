package core

import (
	"fmt"
	"os"
)

func Database() {

	switch os.Getenv("DB_DRIVER") {

	case "mysql":
		// GORM format
		fmt.Printf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	case "sqlite":

	case "mongodb":
		// Expand a string containing environment variables in the form of $var or ${var}
		dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME")
		fmt.Println("DB URL: ", dbURL)

		// Output : DB URL:  mongodb://admin:password@localhost:27017/testdb
	}

}

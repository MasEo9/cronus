/*
Copyright © 2026 MasEo9 <ajmason09@gmail.com>
*/
package main

import (
	"cronus/cmd"
	"cronus/internal/db"
	"os"
	"path/filepath"
	"fmt"

)

func main() {
	value, _ := os.UserConfigDir()
	dbPath := filepath.Join(value, "cronus.db")
	dbInstance, err := db.New(dbPath)
	if err != nil {
		fmt.Println("Database init failed:", err)
		os.Exit(1)
	}
	dbInstance.SessionTableCreate()
	dbInstance.ProjectTableCreate()
	cmd.SetDB(dbInstance)
	cmd.Execute()
}

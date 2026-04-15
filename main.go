/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

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
	cmd.SetDB(dbInstance)
	cmd.Execute()
}

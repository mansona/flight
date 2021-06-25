package main

import (
	"flight/internal/core"
	"flight/internal/figma"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

const (
	iconPath          = "dist" // The folder to export `.svg`s to
	requestsPerMinute = 17     // Amount of requests to Figma to allow per minute
	waitPeriod        = 45     // Wait period in seconds between batch requests
)

func main() {
	log.Println("Taking flight...")

	// Load a .env file, for local development
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fill an array with all the icons from our Figma doc
	icons := figma.FindIconsInDoc()
	count := 0

	for _, i := range icons {
		// TODO: Diff icons on disk and Figma to see which ones we should update

		// For each icon, request an export from Figma
		url, err := figma.ExportIconToURL(i)

		if err != nil {
			// If something goes wrong while exporting with Figma, log the error and continue...
			log.Fatal(err)
		} else {
			// If the export is successful, asynchronously download the given URL
			go core.DownloadFile(url, fmt.Sprintf("./%s/%s.svg", iconPath, i.Name))
		}

		// Increment our total count
		count++

		// Sleep for 45 seconds every X requests so we don't hit Figma's rate limiting
		if count%requestsPerMinute == 0 {
			time.Sleep(waitPeriod * time.Second)
		}
	}
}

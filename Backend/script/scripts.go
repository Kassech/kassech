package scripts

import (
	"log"
	"os"

	"kassech/backend/pkg/database"
)

// HandleScriptCommands handles the command line arguments and calls the appropriate functions
func HandleScriptCommands() {
	if len(os.Args) < 2 {
		log.Println("Usage: go run main.go [migrate|seed]")
		return
	}

	// Run migrations or seed data based on command-line argument
	switch os.Args[1] {
	case "migrate":
		database.Migrate("auto")
		log.Println("Migration completed. Stopping application.")
		os.Exit(0)
	case "fresh":
		database.Migrate("clean")
		log.Println("Clean migration completed. Stopping application.")
		os.Exit(0)
	case "seed":
		database.SeedDB()
		log.Println("Seeding completed. Stopping application.")
		os.Exit(0)

	default:
		log.Fatal("Unknown command. Use 'migrate' or 'seed'.")
	}
}

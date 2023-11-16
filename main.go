package main

import (
	"fmt"

	"github.com/KarkiAnmol/go-fiber-crm/database"
	"github.com/KarkiAnmol/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SetUpRoutes configures the routes for the CRM application.
func SetUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

// initDB initializes the database connection and performs migrations.
func initDB() {
	var err error

	// Open a connection to the SQLite database
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connection opened to database")

	// Auto-migrate the Lead model to create the necessary table
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrateds")
}

func main() {
	// Create a new Fiber application
	app := fiber.New()

	// Initialize the database connection and perform migrations
	initDB()

	// Configure routes for the CRM application
	SetUpRoutes(app)

	// Start the Fiber application on port 3000
	app.Listen(3000)

	// Defer closing the database connection until the main function exits
	defer database.DBConn.Close()
}

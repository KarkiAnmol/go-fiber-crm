package lead

import (
	"github.com/KarkiAnmol/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Lead represents the structure of a lead in the CRM system.
type Lead struct {
	gorm.Model
	Name    string `json:"name"`    // Name of the lead
	Company string `json:"company"` // Company associated with the lead
	Email   string `json:"email"`   // Email address of the lead
	Phone   int    `json:"phone"`   // Phone number of the lead
}

// GetLeads retrieves all leads from the database and responds with a JSON array.
func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

// GetLead retrieves a specific lead by ID from the database and responds with a JSON object.
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

// NewLead creates a new lead by parsing the request body and adding it to the database.
// Responds with the created lead as a JSON object.
func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

// DeleteLead deletes a lead by ID from the database and responds with a success message.
func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}

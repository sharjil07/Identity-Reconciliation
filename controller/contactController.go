package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sharjil07/service-request/database"
	"github.com/sharjil07/service-request/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Identify(c *gin.Context) {
	var request struct {
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	collection := database.Client.Database("contactsdb").Collection("contacts")

	// Find existing contacts
	filter := bson.M{
		"$or": []bson.M{
			{"email": request.Email},
			{"phoneNumber": request.PhoneNumber},
		},
	}

	var existingContacts []models.Contact
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var contact models.Contact
		if err := cursor.Decode(&contact); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding contact"})
			return
		}
		existingContacts = append(existingContacts, contact)
	}

	var primaryContactId int
	var emails []string
	var phoneNumbers []string
	var secondaryContactIds []int

	if len(existingContacts) == 0 {
		// Create a new primary contact
		newContact := models.Contact{
			ID:             generateNewID(),
			Email:          &request.Email,
			PhoneNumber:    &request.PhoneNumber,
			LinkPrecedence: "primary",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		_, err := collection.InsertOne(context.Background(), newContact)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating contact"})
			return
		}
		primaryContactId = newContact.ID
		emails = append(emails, request.Email)
		phoneNumbers = append(phoneNumbers, request.PhoneNumber)
	} else {
		// Handle existing contacts
		primaryContact := existingContacts[0]
		primaryContactId = primaryContact.ID
		for _, contact := range existingContacts {
			if contact.Email != nil {
				emails = append(emails, *contact.Email)
			}
			if contact.PhoneNumber != nil {
				phoneNumbers = append(phoneNumbers, *contact.PhoneNumber)
			}
			if contact.LinkPrecedence == "secondary" {
				secondaryContactIds = append(secondaryContactIds, contact.ID)
			}
		}

		// Add new secondary contact if needed
		if !contains(emails, request.Email) || !contains(phoneNumbers, request.PhoneNumber) {
			newContact := models.Contact{
				ID:             generateNewID(),
				Email:          &request.Email,
				PhoneNumber:    &request.PhoneNumber,
				LinkedID:       &primaryContactId,
				LinkPrecedence: "secondary",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}
			_, err := collection.InsertOne(context.Background(), newContact)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating secondary contact"})
				return
			}
			secondaryContactIds = append(secondaryContactIds, newContact.ID)
		}
	}

	response := gin.H{
		"primaryContactId":    primaryContactId,
		"emails":              emails,
		"phoneNumbers":        phoneNumbers,
		"secondaryContactIds": secondaryContactIds,
	}
	c.JSON(http.StatusOK, response)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func generateNewID() int {
	//generate a new unique ID
	return int(time.Now().UnixNano())
}

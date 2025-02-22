Service Request API
This is a simple service request API built using Go, MongoDB, and the Gin web framework. The API allows you to manage contacts with features to create new contacts and identify existing ones by email or phone number.

Project Structure
models/contact.go: Defines the Contact struct that represents a contact entity in the database.

database/connection.go: Handles the connection to the MongoDB database. It loads environment variables, sets up the MongoDB client, and connects to the database.

controller/contactController.go: Contains the controller logic that handles HTTP requests. The main function is Identify, which processes identifying contacts based on provided email and phone number.

main.go: Entry point of the application. It sets up the Gin router, initializes the database connection, and defines the API routes.
Setup Instructions
Clone the repository to your local machine.

SH

git clone <repository-url>
cd service-request
Install Dependencies

Make sure you have Go installed on your machine. You will also need to install the necessary Go modules:

SH

go mod tidy
Setup MongoDB

Set up a MongoDB instance locally or use a hosted MongoDB service. Create a database named contactsdb.

Environment Variables

Create a .env file in the root directory of your project with the following content, replacing your-mongodb-uri with your actual MongoDB connection URI:

PLAINTEXT

MONGO_URI=your-mongodb-uri
Run the Application

Start the application using the following command:

SH

go run main.go
The server will run on http://localhost:8080.

API Endpoints
POST /identify: Identifies and manages contacts based on email and phone number provided in the request. If no matching contact exists, a new contact will be created as the primary contact. If matching contacts are found, it organizes them as primary and secondary contacts.

Request:

JSON

{
  "email": "example@example.com",
  "phoneNumber": "1234567890"
}
Response:

JSON

{
  "primaryContactId": 1,
  "emails": ["example@example.com"],
  "phoneNumbers": ["1234567890"],
  "secondaryContactIds": [2, 3]
}
Dependencies
Gin Web Framework for building the API.
MongoDB Go Driver for interacting with MongoDB.
godotenv for loading environment variables.

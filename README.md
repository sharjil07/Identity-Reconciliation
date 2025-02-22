# Service Request API

## Overview

This document describes the Service Request API, a simple API built using Go, MongoDB, and the Gin web framework. The API allows you to manage contacts with features to create new contacts and identify existing ones by email or phone number.

## Project Structure

- **models/contact.go**: Defines the `Contact` struct that represents a contact entity in the database.
- **database/connection.go**: Handles the connection to the MongoDB database. It loads environment variables, sets up the MongoDB client, and connects to the database.
- **controller/contactController.go**: Contains the controller logic that handles HTTP requests. The main function is `Identify`, which processes identifying contacts based on provided email and phone number.
- **main.go**: Entry point of the application. It sets up the Gin router, initializes the database connection, and defines the API routes.

## Setup Instructions

1. **Clone the repository** to your local machine.

   ```bash
   git clone <repository-url>
   cd service-request
2. **Install Dependencies** Make sure you have Go installed on your machine. You will also need to install the necessary Go modules:

    ```bash
    go mod tidy
3.**Setup MongoDB**Set up a MongoDB instance locally or use a hosted MongoDB service. Create a database named `contactsdb.

4.**Environment Variables**Create a `.env` file in the root directory of your project with the following content, replacing `your-mongodb-uri` with your actual MongoDB connection URI:
    
    ```bash
      MONGO_URI=your-mongodb-uri
      
5.**Run the Application**Start the application using the following command:

  ```bash
  go run main.go


    

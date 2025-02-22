\documentclass{article}
\usepackage{hyperref}

\title{Service Request API}
\author{}
\date{}

\begin{document}

\maketitle

\section*{Overview}
This document describes the Service Request API, a simple API built using Go, MongoDB, and the Gin web framework. The API allows you to manage contacts with features to create new contacts and identify existing ones by email or phone number.

\section*{Project Structure}
\begin{itemize}
    \item \textbf{models/contact.go}: Defines the \texttt{Contact} struct that represents a contact entity in the database.
    \item \textbf{database/connection.go}: Handles the connection to the MongoDB database. It loads environment variables, sets up the MongoDB client, and connects to the database.
    \item \textbf{controller/contactController.go}: Contains the controller logic that handles HTTP requests. The main function is \texttt{Identify}, which processes identifying contacts based on provided email and phone number.
    \item \textbf{main.go}: Entry point of the application. It sets up the Gin router, initializes the database connection, and defines the API routes.
\end{itemize}

\section*{Setup Instructions}
\begin{enumerate}
    \item \textbf{Clone the repository} to your local machine.
    \begin{verbatim}
    git clone <repository-url>
    cd service-request
    \end{verbatim}

    \item \textbf{Install Dependencies}

    Make sure you have Go installed on your machine. You will also need to install the necessary Go modules:
    \begin{verbatim}
    go mod tidy
    \end{verbatim}

    \item \textbf{Setup MongoDB}

    Set up a MongoDB instance locally or use a hosted MongoDB service. Create a database named \texttt{contactsdb}.

    \item \textbf{Environment Variables}

    Create a \texttt{.env} file in the root directory of your project with the following content, replacing \texttt{your-mongodb-uri} with your actual MongoDB connection URI:
    \begin{verbatim}
    MONGO_URI=your-mongodb-uri
    \end{verbatim}

    \item \textbf{Run the Application}

    Start the application using the following command:
    \begin{verbatim}
    go run main.go
    \end{verbatim}

    The server will run on \texttt{http://localhost:8080}.
\end{enumerate}

\section*{API Endpoints}
\begin{itemize}
    \item \textbf{POST /identify}: Identifies and manages contacts based on email and phone number provided in the request. If no matching contact exists, a new contact will be created as the primary contact. If matching contacts are found, it organizes them as primary and secondary contacts.

    \textbf{Request}:
    \begin{verbatim}
    {
      "email": "example@example.com",
      "phoneNumber": "1234567890"
    }
    \end{verbatim}

    \textbf{Response}:
    \begin{verbatim}
    {
      "primaryContactId": 1,
      "emails": ["example@example.com"],
      "phoneNumbers": ["1234567890"],
      "secondaryContactIds": [2, 3]
    }
    \end{verbatim}
\end{itemize}

\section*{Dependencies}
\begin{itemize}
    \item \href{https://github.com/gin-gonic/gin}{Gin Web Framework} for building the API.
    \item \href{https://github.com/mongodb/mongo-go-driver}{MongoDB Go Driver} for interacting with MongoDB.
    \item \href{https://github.com/joho/godotenv}{godotenv} for loading environment variables.
\end{itemize}



\end{document}

# CRM Backend Project

## Project Overview

This project is a CRM (Customer Relationship Management) backend server built with Go. It provides an HTTP API to perform CRUD (Create, Read, Update, Delete) operations on customer data. The server handles HTTP requests, interacts with a mock database (a slice of structs), and returns appropriate responses.

## Features

- **Get a Single Customer**: Retrieve details of a customer by ID.
- **Get All Customers**: Retrieve a list of all customers.
- **Create a Customer**: Add a new customer to the database.
- **Update a Customer**: Modify details of an existing customer by ID.
- **Delete a Customer**: Remove a customer from the database by ID.

## Installation

To install and run the project, follow these steps:

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/eoaranda/go-crm-backend.git
    cd go-crm-backend
    ```

2. **Install Dependencies**:
    Make sure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).

## Launch

To launch the server, simply run:

```sh
go run main.go
```

The server will start on port 3000.

## Usage

Interact with the API using tools like Postman or cURL:

| Method | URL                 | Description                              |
|--------|---------------------|------------------------------------------|
| GET    | /                   | API v1 basic documentation               | 
| GET    | /customers/{id}     | Get a single customer by ID              |
| GET    | /customers          | Get all customers                        |
| POST   | /customers          | Create a new customer                    |
| PUT    | /customers/{id}     | Update a customer by ID                  |
| DELETE | /customers/{id}     | Delete a customer by ID                  |

## Contact data sample

```json
{
    "ID": "550e8400-e29b-41d4-a716-446655440000",
    "Name": "John Doe",
    "Role": "Software Engineer",
    "Email": "john.doe@example.com",
    "Phone": 5550101,
    "Contacted": true
}
```

## Career Benefits

Building this project will enhance my backend development skills, particularly in handling CRUD operations and HTTP requests.

## Final Note

Feel free to explore and expand on this project. Happy coding!





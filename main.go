package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string `json:"ID"`
	Name      string `json:"Name"`
	Role      string `json:"Role"`
	Email     string `json:"Email"`
	Phone     int    `json:"Phone"`
	Contacted bool   `json:"Contacted"`
}

var customers = []Customer{
	{
		ID:        "550e8400-e29b-41d4-a716-446655440000",
		Name:      "John Doe",
		Role:      "Software Engineer",
		Email:     "john.doe@example.com",
		Phone:     5550101,
		Contacted: true,
	},
	{
		ID:        "550e8400-e29b-41d4-a716-446655440001",
		Name:      "Jane Smith",
		Role:      "Project Manager",
		Email:     "jane.smith@example.com",
		Phone:     5550102,
		Contacted: false,
	},
	{
		ID:        "550e8400-e29b-41d4-a716-446655440002",
		Name:      "Alice Johnson",
		Role:      "Designer",
		Email:     "alice.johnson@example.com",
		Phone:     5550103,
		Contacted: true,
	},
	{
		ID:        "550e8400-e29b-41d4-a716-446655440003",
		Name:      "Bob Brown",
		Role:      "QA Engineer",
		Email:     "bob.brown@example.com",
		Phone:     5550104,
		Contacted: false,
	},
	{
		ID:        "550e8400-e29b-41d4-a716-446655440004",
		Name:      "Charlie Black",
		Role:      "DevOps Engineer",
		Email:     "charlie.black@example.com",
		Phone:     5550105,
		Contacted: true,
	},
	{
		ID:        "550e8400-e29b-41d4-a716-446655440005",
		Name:      "Diana Green",
		Role:      "Business Analyst",
		Email:     "diana.green@example.com",
		Phone:     5550106,
		Contacted: false,
	},
}

func getCustomer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the content type to json
	w.Header().Set("X-Powered-By", "Go")               // set the custom header

	// Receive the slug from the URL
	id := mux.Vars(request)["id"]

	// Check if the customer exists
	for _, customer := range customers {
		if customer.ID == id {
			w.WriteHeader(http.StatusOK) // set the status code to 200
			json.NewEncoder(w).Encode(customer)
			return // return to exit the function when found
		}
	}

	// Return a 404 if the customer does not exist
	w.WriteHeader(http.StatusNotFound) // set the status code to 404
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer not found"})
}

func getCustomers(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the content type to json
	w.Header().Set("X-Powered-By", "Go")               // set the custom header
	w.WriteHeader(http.StatusOK)                       // set the status code to 200
	json.NewEncoder(w).Encode(customers)               // encode the customers to json and write it to the response writer
}

func addCustomer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the content type to json
	w.Header().Set("X-Powered-By", "Go")               // set the custom header

	var newEntry map[string]interface{}
	reqBody, _ := io.ReadAll(request.Body)
	json.Unmarshal(reqBody, &newEntry)

	for _, customer := range customers {
		if customer.ID == newEntry["ID"] {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"message": "Customer already exists"})
			return
		}
	}

	// if the customer does not exist, add it to the slice
	id, _ := uuid.NewRandom()                      // Generate a new Version 4 UUID
	phoneFloat64, _ := newEntry["Phone"].(float64) // Convert Phone to int
	phoneNumber := int(phoneFloat64)
	contacted, _ := newEntry["Contacted"].(bool) // Convert Contacted to bool
	customers = append(customers, Customer{id.String(), newEntry["Name"].(string), newEntry["Role"].(string), newEntry["Email"].(string), phoneNumber, contacted})
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the content type to json
	w.Header().Set("X-Powered-By", "Go")               // set the custom header

	// Receive the slug from the URL
	id := mux.Vars(request)["id"]

	// Keep track of the new entry
	var newEntry map[string]string

	// Read the request
	reqBody, _ := io.ReadAll(request.Body)
	json.Unmarshal(reqBody, &newEntry)

	// Check if the customer exists
	for i, customer := range customers {
		if customer.ID == id {
			// update the customer in the slice if found
			phoneNumber, _ := strconv.Atoi(newEntry["Phone"])
			contacted := newEntry["Contacted"] == "true"
			customers[i] = Customer{id, newEntry["Name"], newEntry["Role"], newEntry["Email"], phoneNumber, contacted}
			w.WriteHeader(http.StatusOK) // set the status code to 200
			json.NewEncoder(w).Encode(customers)
			return // return to exit the function when found
		}
	}

	// Return a 404 if the customer does not exist
	w.WriteHeader(http.StatusNotFound) // set the status code to 404
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer not found"})
}

func deleteCustomer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the content type to json
	w.Header().Set("X-Powered-By", "Go")               // set the custom header

	// Receive the slug from the URL
	id := mux.Vars(request)["id"]

	// Check if the customer exists
	for i, customer := range customers {
		if customer.ID == id {
			// remove the customer from the slice
			customers = append(
				customers[:i],      //  produces a slice of all elements before index i.
				customers[i+1:]...) // produces a slice of all elements after index i.
			w.WriteHeader(http.StatusOK) // set the status code to 200
			json.NewEncoder(w).Encode(customers)
			return // return to exit the function when found
		}
	}

	// Return a 404 if the customer does not exist
	w.WriteHeader(http.StatusNotFound) // set the status code to 404
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer not found"})
}

func main() {
	port := "3000"
	router := mux.NewRouter()

	// Getting a single customer through a /customers/{id} path
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")

	// Getting all customers through a the /customers path
	router.HandleFunc("/customers", getCustomers).Methods("GET")

	// Creating a customer through a /customers path
	router.HandleFunc("/customers", addCustomer).Methods("POST")

	// Updating a customer through a /customers/{id} path
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")

	// Deleting a customer through a /customers/{id} path
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	filserServer := http.FileServer(http.Dir("./static")) // this will serve the files in the static directory
	router.Handle("/", filserServer)

	fmt.Println("Server started on port http://localhost:" + port)
	http.ListenAndServe(":"+port, router)

}

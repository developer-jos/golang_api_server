package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Item represents a shopping item
type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// items is a mock in-memory store for simplicity, replace with a database in production
var items []Item

func init() {
	items = []Item{
		{ID: uuid.New(), Name: "Airtime"},
		{ID: uuid.New(), Name: "Data"},
		{ID: uuid.New(), Name: "Mashup"},
		{ID: uuid.New(), Name: "Broadband"},
		{ID: uuid.New(), Name: "Just4U"},
		{ID: uuid.New(), Name: "Call Abroad"},
		{ID: uuid.New(), Name: "Caller Tunez"},
		{ID: uuid.New(), Name: "SME Plus"},
		// Add more demo items
		{ID: uuid.New(), Name: "Apple"},
		{ID: uuid.New(), Name: "Banana"},
		{ID: uuid.New(), Name: "Milk"},
		{ID: uuid.New(), Name: "Bread"},
		{ID: uuid.New(), Name: "Eggs"},
		{ID: uuid.New(), Name: "Cheese"},
		{ID: uuid.New(), Name: "Chicken"},
		{ID: uuid.New(), Name: "Rice"},
		{ID: uuid.New(), Name: "Pasta"},
		{ID: uuid.New(), Name: "Tomatoes"},
		{ID: uuid.New(), Name: "Potatoes"},
		{ID: uuid.New(), Name: "Onions"},
		{ID: uuid.New(), Name: "Carrots"},
		{ID: uuid.New(), Name: "Broccoli"},
		{ID: uuid.New(), Name: "Spinach"},
		{ID: uuid.New(), Name: "Salmon"},
		{ID: uuid.New(), Name: "Shrimp"},
		{ID: uuid.New(), Name: "Steak"},
		{ID: uuid.New(), Name: "Cereal"},
		{ID: uuid.New(), Name: "Yogurt"},
	}
}

// Database represents the data store
type Database interface {
	GetItems() ([]Item, error)
	CreateItem(newItem Item) (uuid.UUID, error)
	GetItem(id uuid.UUID) (Item, error)
	DeleteItem(id uuid.UUID) error
}

// InMemoryDatabase is a mock in-memory database
type InMemoryDatabase struct{}

// GetItems returns the list of shopping items
func (db *InMemoryDatabase) GetItems() ([]Item, error) {
	return items, nil
}

// CreateItem creates a new shopping item
func (db *InMemoryDatabase) CreateItem(newItem Item) (uuid.UUID, error) {
	newItem.ID = uuid.New()
	items = append(items, newItem)
	return newItem.ID, nil
}

// GetItem retrieves a specific shopping item by ID
func (db *InMemoryDatabase) GetItem(id uuid.UUID) (Item, error) {
	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}
	return Item{}, fmt.Errorf("item not found")
}

// DeleteItem removes a shopping item by ID
func (db *InMemoryDatabase) DeleteItem(id uuid.UUID) error {
	for i, item := range items {
		if item.ID == id {
			// Remove the item from the slice
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

// Server represents the API server
type Server struct {
	DB Database
}

// NewServer creates a new instance of the API server
func NewServer(db Database) *Server {
	return &Server{DB: db}
}

// GetItems returns the list of shopping items
func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := s.DB.GetItems()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving items: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// CreateItem creates a new shopping item
func (s *Server) CreateItem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newItem Item
	err := decoder.Decode(&newItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		return
	}

	newItem.ID = uuid.New()
	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
}

// HandlerFunc returns a handler function with logging
func HandlerFunc(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		handler(w, r)
	}
}

// logRequest logs information about the incoming request
func logRequest(r *http.Request) {
	// Log request details here
}

// GetItem retrieves a specific shopping item by ID
func (s *Server) GetItem(w http.ResponseWriter, r *http.Request) {
	// Implementation omitted for brevity
}

// DeleteItem removes a shopping item by ID
func (s *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Implementation omitted for brevity
}

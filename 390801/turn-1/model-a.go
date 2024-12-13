package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Product struct represents the data model
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ProductRepository defines the methods to access product data
type ProductRepository interface {
	Create(product *Product) error
	FindByID(id int) (*Product, error)
	GetAll() ([]Product, error)
	Update(product *Product) error
	Delete(id int) error
}

// InMemoryProductRepo is an in-memory implementation of ProductRepository
type InMemoryProductRepo struct {
	mu       sync.Mutex
	products map[int]Product
	nextID   int
}

// NewInMemoryProductRepo initializes a new InMemoryProductRepo
func NewInMemoryProductRepo() *InMemoryProductRepo {
	return &InMemoryProductRepo{
		products: make(map[int]Product),
		nextID:   1,
	}
}

// Create adds a new product to the repository
func (repo *InMemoryProductRepo) Create(product *Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	product.ID = repo.nextID
	repo.products[repo.nextID] = *product
	repo.nextID++
	return nil
}

// FindByID retrieves a product by ID
func (repo *InMemoryProductRepo) FindByID(id int) (*Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	product, exists := repo.products[id]
	if !exists {
		return nil, fmt.Errorf("product not found")
	}
	return &product, nil
}

// GetAll retrieves all products
func (repo *InMemoryProductRepo) GetAll() ([]Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	productsSlice := make([]Product, 0, len(repo.products))
	for _, product := range repo.products {
		productsSlice = append(productsSlice, product)
	}
	return productsSlice, nil
}

// Update modifies an existing product
func (repo *InMemoryProductRepo) Update(product *Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, exists := repo.products[product.ID]
	if !exists {
		return fmt.Errorf("product not found for update")
	}
	repo.products[product.ID] = *product
	return nil
}

// Delete removes a product from the repository
func (repo *InMemoryProductRepo) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, exists := repo.products[id]
	if !exists {
		return fmt.Errorf("product not found for deletion")
	}
	delete(repo.products, id)
	return nil
}

func main() {
	repo := NewInMemoryProductRepo()
	router := mux.NewRouter()

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			err := repo.Create(&product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
		} else {
			products, _ := repo.GetAll()
			json.NewEncoder(w).Encode(products)
		}
	}).Methods(http.MethodGet, http.MethodPost)

	router.HandleFunc("/products/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		// Handle GET, PUT, DELETE following the same structure.
	}).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	http.ListenAndServe(":8000", router)
}

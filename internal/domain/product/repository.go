package product

import (
	"errors"
	"fmt"

	domain "github.com/aldogayala/GoWeb/internal/domain"
)

var (
	ErrorNotFound = errors.New("item not found")
)

type Repository interface {
	//read
	Get() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	ExistProduct(name string) bool
	//write
	Create(domain.Product) (int, error)
}

type repository struct {
	db *[]domain.Product
}

func NewRepository(db *[]domain.Product) Repository {
	return &repository{db: db}
}

// Read
func (r *repository) Get() ([]domain.Product, error) {
	return *r.db, nil
}

func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, w := range *r.db {
		if w.Id == id {
			return w, nil
		}
	}

	return domain.Product{}, fmt.Errorf("%w. %s", ErrorNotFound, "product does not exist")
}

func (r *repository) ExistProduct(name string) bool {
	for _, w := range *r.db {
		if w.Name == name {
			return true
		}
	}
	return false
}

// Write
func (r *repository) Create(product domain.Product) (int, error) {
	*r.db = append(*r.db, product)
	return product.Id, nil
}

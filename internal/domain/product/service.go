package product

import (
	"errors"
	"fmt"

	domain "github.com/aldogayala/GoWeb/internal/domain"
)

var (
	ErrorAlreadyExist = errors.New("already exist")
)

type Service interface {
	//read
	Get() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	//write
	Create(name string, quantity int, codeValue string, isPublished bool, price float64) (domain.Product, error)
}

type service struct {
	//Repo inyection
	rp Repository
	//External APIS
	//..
}

func NewService(rp Repository) Service {
	return &service{rp: rp}
}

// read
func (sv *service) Get() ([]domain.Product, error) {
	return sv.rp.Get()
}

func (sv *service) GetByID(id int) (domain.Product, error) {
	return sv.rp.GetByID(id)
}

// write
func (sv *service) Create(name string, quantity int, codeValue string, isPublished bool, price float64) (domain.Product, error) {
	if sv.rp.ExistProduct(name) {
		return domain.Product{}, ErrorAlreadyExist
	}

	pr := domain.Product{
		Name:        name,
		Quantity:    quantity,
		CodeValue:   codeValue,
		IsPublished: isPublished,
		Price:       price,
	}

	lastID, err := sv.rp.Create(pr)
	if err != nil {
		return domain.Product{}, fmt.Errorf("%w. %s", err, "Create product error")
	}

	pr.Id = lastID

	return pr, nil
}

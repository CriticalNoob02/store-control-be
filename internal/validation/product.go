package validation

import (
	"errors"

	"github.com/CriticalNoob02/store-control-be/internal/model"
)

var ProductFilterList = []string{"name", "amount", "category", "price"}

func ProductValidation(product model.Product) error {
	if product.Name == "" {
		return errors.New("campo de Nome e de preenchimento obrigatorio")
	}
	if product.Category == "" {
		return errors.New("campo de Categoria e de preenchimento obrigatorio")
	}
	if product.Price == 0 {
		return errors.New("campo de Preco e de preenchimento obrigatorio")
	}
	return nil
}

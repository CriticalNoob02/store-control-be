package validation

import (
	"errors"

	"github.com/CriticalNoob02/store-control-be/internal/model"
)

var ProductFilterList = []string{"name", "amount", "category", "price"}

func ProductValidation(user model.Product) error {
	if user.Name == "" {
		return errors.New("campo de Nome e de preenchimento obrigatorio")
	}
	if user.Category == "" {
		return errors.New("campo de Categoria e de preenchimento obrigatorio")
	}
	if user.Price == 0 {
		return errors.New("campo de Preco e de preenchimento obrigatorio")
	}
	return nil
}

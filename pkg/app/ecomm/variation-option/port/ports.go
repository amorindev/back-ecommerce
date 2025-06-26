package port

import (
	"context"

	"com.fernando/pkg/app/ecomm/variation-option/model"
)

type VariationOptionRepo interface {
	Insert(ctx context.Context, varOption *model.VariationOption) error
	// * VER EL TEMA DE ARREGLOS MEJOR DESDE EL SERVICO
	// * GET SI EXISTE AHI MISMO RETORNAR EL ERROR, Y NO ESTAR BUSCANDO DENTRO DE CREATE FUNCÍN
	CreateMany(ctx context.Context, varOptions []*model.VariationOption) error
	// * o ExistsVaritionOption, verifica si un elemento existe,
	//? que debería retornar un erroglo o la primera coincidencia
	// * y asi imformarle al usario, esta retornando los que ya existen
	ExistOne(ctx context.Context, varOptions []*model.VariationOption) ([]*model.VariationOption, error)
	GetByName(ctx context.Context, name string) (*model.VariationOption, error)
}

type VariationOptionService interface {
}

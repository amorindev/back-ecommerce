package port

import (
	"context"

	"com.fernando/pkg/app/cupon/model"
)

// * no nos compliquemoos por que cupon es mas para saas  y no creo que se usen ambos ver
// * activar desde admin al final es calculable y se notifica al usuario
type CuponRepository interface {
	Insert(ctx context.Context, cupon *model.Cupon) error
	// GetAllpaginacion Get de pruebas rapidas
	//Get(ctx context.Context) ([]*model.Cupon, error) // no lo uso de momento admin me parece
	// * De momento Get By description - description de momento es unique
	GetByDescription(ctx context.Context, description string) (*model.Cupon, error)
}

// * ver lalogica para calcularo con todos sus propiedades si esta activo
// * y mencionarle al usuario
// * Get es para obtener mediante el ID ver onboarding ports
type CuponService interface {
	Create(ctx context.Context, cupon *model.Cupon) error
	Get(ctx context.Context) ([]*model.Cupon, error)
}

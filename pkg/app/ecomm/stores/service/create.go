package service

import (
	"context"
	"time"

	"com.fernando/pkg/app/ecomm/address/model"
	storeM "com.fernando/pkg/app/ecomm/stores/model"
)

func (s *Service) Create(ctx context.Context, store *storeM.Store, address *model.Address) error {
	now := time.Now().UTC()

	// ver del store su created
	address.CreatedAt = &now
	address.UpdatedAt = &now

	
	return s.StoreTx.Create(ctx, store, address)

	//return s.StoreRepo.Insert(ctx, store)
}

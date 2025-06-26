package service

import (
	authPort "com.fernando/pkg/app/auth/port"
	categoryPort "com.fernando/pkg/app/ecomm/category/port"
	productPort "com.fernando/pkg/app/ecomm/products/port"
	storePort "com.fernando/pkg/app/ecomm/stores/port"
	varOptionPort "com.fernando/pkg/app/ecomm/variation-option/port"
	variationPort "com.fernando/pkg/app/ecomm/variation/port"
	initPort "com.fernando/pkg/app/init/port"
	onboardingPort "com.fernando/pkg/app/onboarding/port"
	permissionPort "com.fernando/pkg/app/permission/port"
	rolePort "com.fernando/pkg/app/role/port"
	userRepo "com.fernando/pkg/app/user/port"
)

var _ initPort.InitService = &Service{}

type Service struct {
	AuthRepo       authPort.AuthRepo
	UserRepo       userRepo.UserRepo
	RoleRepo       rolePort.RoleRepo
	PermissionRepo permissionPort.PermissinRepo
	CategoryRepo   categoryPort.CategoryRepo
	VariationRepo  variationPort.VariationRepo
	VarOptionRepo  varOptionPort.VariationOptionRepo
	ProductRepo    productPort.ProductRepo
	OnboardingRepo onboardingPort.OnboardingRepo
	OnboardingSrv  onboardingPort.OnboardingSrv
	ProductSrv     productPort.ProductSrv
	StoreSrv       storePort.StoreSrv
	AuthTx         authPort.AuthTransaction
}

func NewService(
	authRepo authPort.AuthRepo, 
	userRepo userRepo.UserRepo, 
	roleRepo rolePort.RoleRepo, 
	permissionRepo permissionPort.PermissinRepo, 
	categoryRepo categoryPort.CategoryRepo, 
	variationRepo variationPort.VariationRepo, 
	varOptionRepo varOptionPort.VariationOptionRepo, 
	productRepo productPort.ProductRepo, 
	onboardingRepo onboardingPort.OnboardingRepo, 
	onboardingSrv onboardingPort.OnboardingSrv, 
	productSrv productPort.ProductSrv, 
	storeSrv storePort.StoreSrv, 
	authTx authPort.AuthTransaction,
) *Service {
	return &Service{
		AuthRepo:       authRepo,
		UserRepo:       userRepo,
		RoleRepo:       roleRepo,
		PermissionRepo: permissionRepo,
		CategoryRepo:   categoryRepo,
		VariationRepo:  variationRepo,
		VarOptionRepo:  varOptionRepo,
		ProductRepo:    productRepo,
		OnboardingRepo: onboardingRepo,
		OnboardingSrv:  onboardingSrv,
		ProductSrv:     productSrv,
		StoreSrv:       storeSrv,
		AuthTx:         authTx,
	}
}

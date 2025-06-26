package service

import (
	onboardingPort "com.fernando/pkg/app/onboarding/port"
	fileStoragePort "com.fernando/pkg/file-storage/port"
)

var _ onboardingPort.OnboardingSrv = &Service{}

type Service struct {
	OnboardingRepo onboardingPort.OnboardingRepo
	FileStorageSrv fileStoragePort.FileStorageSrv
}

// ! si vamos a guardar que el usuario vio o no vio el onbarding
// ! necesitaríamos tabla intermedia y trasnsaccion
// ! o llamarlo NewOnboarding service para buscarlo mas rápido
func NewService(onboardingRepo onboardingPort.OnboardingRepo, fileStorageSrv fileStoragePort.FileStorageSrv) *Service {
	return &Service{
		OnboardingRepo: onboardingRepo,
		FileStorageSrv: fileStorageSrv,
	}
}

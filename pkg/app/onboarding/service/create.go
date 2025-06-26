package service

import (
	"context"
	"time"

	"com.fernando/pkg/app/onboarding/model"
)

// antes de retornar el json ver que datos no se devben enviar
/* func print(onboarding *model.Onboarding) {
	fmt.Printf("FileName: %v\n", onboarding.FileName)
	fmt.Printf("File length: %v\n", len(onboarding.File))
	fmt.Printf("Title: %v\n", onboarding.Title)
	fmt.Printf("Text: %v\n", onboarding.Text)
	fmt.Printf("ExpireAt: %v\n", onboarding.ExpiresAt)
} */

func (s *Service) Create(ctx context.Context, onboarding *model.Onboarding) error {
	now := time.Now().UTC() // ver en todos que se esten usando UTC

	onboarding.CreatedAt = &now

	bucketFolderStruct := "onboardings/"

	onboarding.FileName = bucketFolderStruct + onboarding.FileName

	//fmt.Printf("Create Size: %v\n", len(onboarding.File))

	err := s.FileStorageSrv.UploadProduct(context.Background(), onboarding.FileName, onboarding.File, onboarding.ContentType)
	if err != nil {
		return err
	}

	// * como obtener la url, o quería por que al momento de crear un producto queiro retornar el onboarding
	// * creado asi para todos los DDD que usan filestorage

	err = s.OnboardingRepo.Insert(ctx, onboarding)
	if err != nil {
		return err
	}

	return nil
}

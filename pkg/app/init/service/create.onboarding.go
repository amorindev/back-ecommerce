package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	onboardingErr "com.fernando/pkg/app/onboarding/errors"
	"com.fernando/pkg/app/onboarding/model"
)

// * cambiar los demas a prvado por que solo se usara aqui y no cause confuciones
type OnboardingData struct {
	Data []*model.Onboarding `json:"onboarding"`
}

func (s *Service) CreateOnboarding() error {
	jsonFile, err := os.Open("pkg/app/init/files/onboarding/data/data.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var onboardingData OnboardingData
	err = json.Unmarshal(byteValue, &onboardingData)
	if err != nil {
		return err
	}

	/* for i, obd := range onboardingData.Data {
		//fmt.Printf("Onboarding %d\n", i)
		//fmt.Printf("Filename: %s\n", obd.FileName)
		//fmt.Printf("Title: %s\n", obd.Title)
		//fmt.Printf("Text: %s\n", obd.Text)
	} */

	// ! me parece mejorotra lista para agregar todos losque se van a crear
	var onboardingsInsert []*model.Onboarding
	for _, obd := range onboardingData.Data {

		// * Verificar que no exista el onboarding
		onboarding, err := s.OnboardingRepo.GetByTitle(context.Background(), obd.Title)
		// *ver esta condicional para el proyecto en general
		if err != nil && err != onboardingErr.ErrOnboardingNotFound {
			return err
		}
		if onboarding != nil {
			continue
		}
		// * validaciones
		if obd.FileName == "" || obd.Text == "" || obd.Title == "" {
			return errors.New("text, filename, tile are required")
		}
		// * esto no funciona ademas tendríamos que agregar otra variable
		//onboardingData.Data[i].ID = "test" en la siguiente capa su valor es nulo

		// * Ruta y nombre de la imagen
		// Make Payment -> make-payment, seria mejor usar title t no filename aunque es como aux
		// si esta bien asi entonces en CreateProducts mejorarlo y pasarle la ruta completa
		// para no hacer joins inecesarios que pueden ser causa de posibles errores
		// ! recuerda cambiarlo al final en elservico que lo va a utilizar
		// ! hoy refactorizar minio con DDD profile categories
		// ! el bug de productos
		//fmt.Printf("File path: %s", obd.FileName)
		imgName := strings.ToLower(strings.ReplaceAll(obd.FileName, " ", "-"))
		// ubicacion de la imagen
		imgPath := filepath.Join("pkg/app/init/files/onboarding/imgs", obd.FileName)

		//fmt.Printf("File path: %s", imgPath)

		// * Cargar la imagen
		imgData, err := os.ReadFile(imgPath)
		if err != nil {
			return fmt.Errorf("create onboarding error reading image file %s: %v", imgPath, err)
		}
		//fmt.Printf("Image data: %v\n", len(imgData))

		obd.File = imgData
		obd.FileName = imgName
		onboardingsInsert = append(onboardingsInsert, obd)
	}

	/* for i, obd := range onboardingData.Data {
		//fmt.Printf("Onboarding %d\n", i)
		//fmt.Printf("Filename: %s\n", obd.FileName)
		//fmt.Printf("File length: %d\n", len(obd.File))
		//fmt.Printf("Title: %s\n", obd.Title)
		//fmt.Printf("Text: %s\n", obd.Text)
	} */

	for _, obd := range onboardingsInsert {
		// ! no se me parece inadecuado hacer un for mejor crear otra lista con 
		// ! solo con los onoarding a insertar
		// * Devería pasar por el servicio?
		// ! verificar por que arriba ya lo estoy revisando
		// ! otro forma seria asignarlo null al campo que ya existe de la lista
		_, err := s.OnboardingRepo.GetByTitle(context.Background(), obd.Title)
		if err != nil && err != onboardingErr.ErrOnboardingNotFound {
			return err
		}
		/* if  o!= nil{
			continue
		} */
		err = s.OnboardingSrv.Create(context.Background(), obd)
		if err != nil {
			return err
		}
	}
	return nil
}

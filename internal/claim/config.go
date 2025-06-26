package claim

import (
	"errors"
	"os"
	"sync"
)

type ConfigJWT struct {
	AccessString string
	RefreshString string
	Issuer        string
	Audience      []string // ! de momento vacío
}

var (
	once      sync.Once
	configJwt *ConfigJWT
)

// es mejor crear esta funcionaidad o directamente desde os.Getenv funcion
// esta afectando refresh.token service
func GetConfig() (*ConfigJWT, error) {
	once.Do(func() {
		configJwt = &ConfigJWT{
			AccessString: os.Getenv("JWT_SIGNING_STRING"),
			Issuer:        os.Getenv("JWT_ISS"),
			RefreshString: os.Getenv("JWT_REFRESH_STRING"),
		}
	})
	if configJwt.AccessString == "" || configJwt.RefreshString == "" || configJwt.Issuer == "" {
		return nil, errors.New("singning or refresh or issuer string not set")
	}
	return configJwt, nil
}

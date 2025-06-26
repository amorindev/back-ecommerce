package service

import (
	"context"
	"encoding/json"
	"io"
	"os"

	addressM "com.fernando/pkg/app/ecomm/address/model"
	errStore "com.fernando/pkg/app/ecomm/stores/errors"
	storeM "com.fernando/pkg/app/ecomm/stores/model"
)

type store struct {
	Name        string  `json:"name"`
	AddressLine string  `json:"address_line"`
	Description string  `json:"description"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Country     string  `json:"country"`
	PostalCode  string  `json:"postal_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type storeData struct {
	Stores []*store `json:"stores"`
}

/*
	 type ShopsData struct {
		Shops []*model.Store `json:"stores"`
	}
*/
func (s *Service) CreateStores() error {
	jsonFile, err := os.Open("pkg/app/init/files/store/data/data.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var shopData storeData
	err = json.Unmarshal(byteValue, &shopData)
	if err != nil {
		return err
	}
	//
	var storesInsert []*storeM.Store
	var addressesInsert []*addressM.Address
	for _, store := range shopData.Stores {
		str, err := s.StoreSrv.GetByName(context.Background(), store.Name)
		if err != nil && err != errStore.ErrStoreNotFound {
			return err
		}
		if str != nil {
			continue
		}
		// * Como tenemos control desde el json despues lo vemos
		/* if str.Name != nil || str.Address != nil || *str.Descripcion == "" {
			if *str.Name == "" || *str.Address == "" || *str.Descripcion == "" {
				return errors.New("name, address, description are required")
			}
		} */
		storeI := &storeM.Store{
			Name:        &store.Name,
			Descripcion: &store.Description,
		}
		addressI := &addressM.Address{
			Label:       &store.Name,
			AddressLine: store.AddressLine,
			City:        store.City,
			State:       store.State,
			Country:     store.Country,
			PostalCode:  store.PostalCode,
			Latitude:    store.Latitude,
			Longitude:   store.Longitude,
		}

		storesInsert = append(storesInsert, storeI)
		addressesInsert = append(addressesInsert, addressI)
	}
	/* for i, s := range storesInsert {
		fmt.Printf("Storage num %d\n", i+1)
		fmt.Printf("Storage: %+v\n", s)

	} */
	// * ver por que se estacreando dos slices difernetes
	// * Lo mejor sería insertMany con trnasacciones de
	for i, store := range storesInsert {
		addss := addressesInsert[i]
		err = s.StoreSrv.Create(context.Background(), store, addss)
		if err != nil {
			return err
		}
	}

	return nil
}

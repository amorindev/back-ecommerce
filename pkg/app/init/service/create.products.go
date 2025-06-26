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

	errProduct "com.fernando/pkg/app/ecomm/products/errors"
	"com.fernando/pkg/app/ecomm/products/model"
)

// ! todavia no coliciona o se confunde de imagen minio por que estamos usando
// ! nombres diferentes blac39 y asi pero que pasa si hay una variacion black39
// ! mejor asignar un identificador único - por que agregarás mas productos
type ProductData struct {
	Products []*model.Product `json:"products"`
}

func (s *Service) CreateProducts() error {
	jsonFile, err := os.Open("pkg/app/init/files/data/insert_products.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var productsData ProductData

	err = json.Unmarshal(byteValue, &productsData)
	if err != nil {
		return err
	}

	var productsInsert []*model.Product

	for i, product := range productsData.Products {
		// * Verificar que no exista el producto
		// * verificar que no exista la variacion - eso despues
		// * ver cual vamos adoptar - no crear si por lo menos uno existe y reotrnar que existe
		// * o como este si existe no lo creo ver
		resp, err := s.ProductRepo.GetByName(context.Background(), product.Name)
		if err != nil {
			if err != errProduct.ErrProductNotFound {
				return err
			}
		}
		if resp != nil {
			continue
		}

		if product.CategoryName == "" {
			return errors.New("category no asignado")
		}

		category, err := s.CategoryRepo.GetByName(context.Background(), product.CategoryName)
		if err != nil {
			return err
		}

		// * porseacaso pero no debería pasae
		if category == nil {
			return errors.New("category is  nil")
		}

		//product.CategoryID = category.ID.(string)
		productsData.Products[i].CategoryID = category.ID

		// Determinar la ruta de las imágenes basado en el nombre del producto
		folderName := strings.ToLower(strings.ReplaceAll(product.Name, " ", "-"))
		imgPath := filepath.Join("pkg/app/init/files/imgs", folderName)
		//fmt.Printf("Folder name: %v\n", folderName)

		// * Procesar imágenes para cada producto
		for j, productItem := range product.ProductItems {
			var imgName string
			if len(product.Variations) > 0 {
				//if len(pg.Variations) > 0 {
				// Producto con variaciones (ej: "xl-10.png")
				var optionValues []string
				for _, opt := range productItem.Options {
					optionValues = append(optionValues, strings.ToLower(opt.Value))
					// * Otra cosa que voy hacer aqui es traerme el ID del variation option
					// * por que desde la ui se debe tener el id
					varOption, err := s.VarOptionRepo.GetByName(context.Background(), opt.Value)
					if err != nil {
						return err
					}
					opt.VarOptionID = varOption.ID
				}

				imgName = strings.Join(optionValues, "-") + ".jpg"

			} else {
				imgName = "img.jpg"

			}

			// cargar la imagen
			imgFullPath := filepath.Join(imgPath, imgName)
			imgData, err := os.ReadFile(imgFullPath)
			if err != nil {
				return fmt.Errorf("error reading image file %s: %v", imgFullPath, err)
			}
			// * Asignar la imagen al producto
			/* product.File = imgData
			product.FileName = imgName */

			productsData.Products[i].ProductItems[j].File = imgData
			productsData.Products[i].ProductItems[j].FileName = imgName

			// * Si es el primer producto, asignar como imagen principal del grupo
			// * o hacerlo desde servicio? me parece mejor
			if j == 0 {
				//pg.FileName = imgName
				// ? como hacer con el image url
				productsData.Products[i].FileName = imgName
			}
		}

		productsInsert = append(productsInsert, product)
	}

	/* for i, product := range productsInsert {
		fmt.Printf("********************************************** Product: %d************\n", i+1)

		fmt.Printf("Product %d\n", i+1)
		fmt.Printf("Name: %v\n", product.Name)
		fmt.Printf("Desc: %v\n", product.Description)
		fmt.Printf("ImgUrl: %v\n", product.ImgUrl)
		fmt.Printf("FilePath: %v\n", product.FileName)
		fmt.Printf("CategoryName: %v\n", product.CategoryName)
		fmt.Printf("CategoryID: %v\n", product.CategoryID)
		for _, proItem := range product.ProductItems {
			fmt.Printf("File name: %v\n", proItem.FileName)
			fmt.Printf("File lenght: %v\n", len(proItem.File))
			fmt.Printf("Name: %v\n", proItem.Price)
		}
	} */
	// * por que no usar productsData.Products en ves de crear otro slice
	//for i, proInsert := range productsData.Products {
	for _, proInsert := range productsInsert {
		//fmt.Printf("Creating product %d: %s\n", i+1, proInsert.Name)
		//fmt.Printf("ProductItems count: %d\n", len(proInsert.ProductItems))
		/* for j, item := range proInsert.ProductItems {
			fmt.Printf("  Item %d: Price=%v, FileName=%v, FileLen=%d\n",
				j+1, item.Price, item.FileName, len(item.File))
		} */
		err = s.ProductSrv.Create(context.Background(), proInsert)
		if err != nil {
			return err
		}
	}

	return nil
}

// pueden ser nulos pro que ya existen estas haciendo un get produuc getByname
// revisar si coinciden los nombres de las imagenes con el insert.json

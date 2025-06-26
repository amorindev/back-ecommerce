package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"com.fernando/cmd/api/message"
	productModel "com.fernando/pkg/app/ecomm/products/model"
	"com.fernando/pkg/app/ecomm/product-item/core"
	"com.fernando/pkg/app/ecomm/product-item/model"
	"com.fernando/pkg/app/ecomm/product-item/validate"
)

// ? Pasarlo a product handler? - hasta ahora sería mejor para separarlo
// ? la creacion de prodcutos con variantes y sin variantes
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// * Obtener la imagen ------------------------------------

	// * parser el formularion multipart (limite de 10mb)
	// * El límite de tamaño que estableces con ParseMultipartForm(10 << 20)
	// * (que equivale a 10 MB) se aplica a todo el cuerpo de la solicitud multipart,
	// * no a cada archivo individualmente.
	err := r.ParseMultipartForm(10 << 20) // * 10 * 2 ^ 20 =  10mb
	if err != nil {
		msg := fmt.Sprintf("ParseMultipartForm err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}

	// ? que pasa si es el agregglo de products como lo identifica a cada imagen
	// * son una lista de imagenes
	// * Obtener el archivo del formulario
	file, handler, err := r.FormFile("image")
	if err != nil {
		msg := fmt.Sprintf("FormFile err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}

	// *Otra forma
	// Specify the maximum file size (in bytes)
	//maxFileSize := int64(10 * 1024 * 1024) // 10 MB
	// * Verificar el tamaño por cada archivo
	maxFileSize := int64(5 << 20) // 5MB como máximo
	if handler.Size > maxFileSize {
		// como hacer el 5MB dinámico
		msg := fmt.Sprintf("el archivo %s excede el tamaño adecuado (5MB)", handler.Filename)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}

	defer file.Close()

	img, err := handler.Open()
	if err != nil {
		msg := fmt.Sprintf("Open err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}

	// Read the file content into a byte slice
	imgData, err := io.ReadAll(img)
	if err != nil {
		msg := fmt.Sprintf("ReadAll err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}

	// * Retrieve the file name from the file header
	fileName := handler.Filename
	// * Create a unique file name to avoid overwriting existing files
	fileName = fileName[:len(fileName)-len(filepath.Ext(fileName))]
	uniqueFileName := fmt.Sprintf("%s-%d%s", fileName, time.Now().UnixNano(), filepath.Ext(fileName))

	// * Obtener el producto --------------------------------------------
	var productReq core.CreateProductReq

	err = json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	defer r.Body.Close()

	// * Validate
	_, err = validate.Create()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	productItem := &model.ProductItem{
		//ID: , desde la base de datos
		Stock:    productReq.CreateProductReq.Stock,
		Discount: productReq.CreateProductReq.Discount,
		//Rating: ,
		Price:    productReq.CreateProductReq.Price,
		FileName: uniqueFileName,
		File:     imgData,
		//CreatedAt: , service
		//UpdatedAt: , service
	}

	product := &productModel.Product{
		// ID: , database
		Name:        productReq.Name,
		Description: productReq.Description,
		//ImgUrl: , despues de cargarlo en file-storage
		Status:     productReq.Status,
		CategoryID: productReq.CategoryID,
		// CreatedAt: , servicio
		// UpdatedAt: , servicio
		// * aqui el producto no mostrará mas de uno
		ProductItems: []*model.ProductItem{productItem},
	}

	err = h.ProductSrv.Create(context.Background(), product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

/*
// Crear un archivo local para guardar la imagen
	dst, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copiar el contenido del archivo subido al archivo local
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)


	func main() {
	// Crear directorio de uploads si no existe
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		panic(err)
	}

	http.HandleFunc("/upload", uploadHandler)
	//fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
*/

// * Medium https://yustinayasin.medium.com/mastering-object-storage-in-go-a-comprehensive-guide-to-minio-minio-series-part-3-3c4eadc27a21

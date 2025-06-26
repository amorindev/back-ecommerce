package handler

import "net/http"

func (h Handler) CreateWithVariations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// * Obtener la imagen ------------------------------------

}

/* func (h Handler) Create(w http.ResponseWriter, r *http.Request) {


	// parser el formularion multipart (limite de 10mb)
	err := r.ParseMultipartForm(10 << 20) // 10mb
	if err != nil {
		msg := fmt.Sprintf("ParseMultipartForm err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}

	// * Obtener el archivo del formulario
	file, handler, err := r.FormFile("image")
	if err != nil {
		msg := fmt.Sprintf("FormFile err: %v", err)
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

	imgData, err := io.ReadAll(img)
	if err != nil {
		msg := fmt.Sprintf("ReadAll err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: msg})
		return
	}

	fileName := handler.Filename
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

	product := &model.Product{
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

	product := &model.Product{
		// ID: , database
		Name:        productReq.Name,
		Description: productReq.Description,
		//ImgUrl: , despues de cargarlo en file-storage
		Status:     productReq.Status,
		CategoryID: productReq.CategoryID,
		// CreatedAt: , servicio
		// UpdatedAt: , servicio
		// * aqui el producto no mostrará mas de uno
		Products: []*model.Product{product},
	}

	err = h.ProductSrv.Create(context.Background(), product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
} */
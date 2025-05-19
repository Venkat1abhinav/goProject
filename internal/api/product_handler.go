package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/venkat1abhinav/goProject/internal/store"
)

type ProductInventoryHandler struct {
	ProductStore store.ProductStore
}

func NewProductHander(ProductStore store.ProductStore) *ProductInventoryHandler {
	return &ProductInventoryHandler{
		ProductStore: ProductStore,
	}
}

func (pi *ProductInventoryHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productParamsId := chi.URLParam(r, "id")
	if productParamsId == "" {
		http.NotFound(w, r)
		return
	}
	pID, err := strconv.ParseInt(productParamsId, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "the product has with ID: %d\n", pID)
}

func (pi *ProductInventoryHandler) HandleCreateProductInventory(w http.ResponseWriter, r *http.Request) {
	var product store.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "something problematic has occured check the log information", http.StatusInternalServerError)
		return
	}

	createProduct, err := pi.ProductStore.CreateProduct(&product)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "something problematic has occured check the log information", http.StatusInternalServerError)
		return

	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&createProduct)

}

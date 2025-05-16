package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductInventoryHandler struct {
}

func NewProductHander() *ProductInventoryHandler {
	return &ProductInventoryHandler{}
}

func (pi *ProductInventoryHandler) CreatProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the product has been invented\n")
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

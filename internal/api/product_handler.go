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

func (pi *ProductInventoryHandler) HandleGetProductById(w http.ResponseWriter, r *http.Request) {
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

	product, err := pi.ProductStore.GetProductById(pID)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to fetch the workout", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(product)

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

	json.NewEncoder(w).Encode(createProduct)

}

func (pi *ProductInventoryHandler) HandleUpdateProductInventory(w http.ResponseWriter, r *http.Request) {
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

	existingProduct, err := pi.ProductStore.GetProductById(pID)

	if err != nil {
		http.Error(w, "failed to fetch thw workout", http.StatusInternalServerError)
		return
	}

	if existingProduct == nil {
		http.NotFound(w, r)
		return
	}

	// at this point we are gonna assume that user was able to find exisiting product

	var updateProductRequest struct {
		ImageUrl    *string              `json:"image_url"`
		DisplayName *string              `json:"display_name"`
		Rating      *int                 `json:"rating"`
		Description *string              `json:"description"`
		Category    *string              `json:"category"`
		Activation  *bool                `json:"activation"`
		Entries     []store.ProductEntry `json:"entries"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateProductRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updateProductRequest.ImageUrl == nil {
		existingProduct.ImageUrl = updateProductRequest.ImageUrl
	}
	if updateProductRequest.DisplayName == nil {
		existingProduct.DisplayName = *updateProductRequest.DisplayName
	}
	if updateProductRequest.Rating == nil {
		existingProduct.Rating = updateProductRequest.Rating
	}
	if updateProductRequest.Description == nil {
		existingProduct.Description = *updateProductRequest.Description
	}
	if updateProductRequest.Category == nil {
		existingProduct.Category = *updateProductRequest.Category
	}
	if updateProductRequest.Activation == nil {
		existingProduct.Activation = updateProductRequest.Activation
	}
	if updateProductRequest.Entries == nil {
		existingProduct.Entries = updateProductRequest.Entries
	}

	err = pi.ProductStore.UpdateProduct(existingProduct)

	if err != nil {
		fmt.Println("the update error", err)
		http.Error(w, "failed to update the workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingProduct)
}

package products

import (
	"log"
	"net/http"

	"github.com/riyagarg2711/ecom-api-course/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// method
func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Call the service -> ListProduct
	products,err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// 2. Return JSON In a HTTP Response

	
	json.Write(w, http.StatusOK, products)
}

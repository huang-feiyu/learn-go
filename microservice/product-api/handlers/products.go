package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"huangblog.com/microservice/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

/**
 * https://docs.microsoft.com/zh-cn/azure/architecture/best-practices/api-design
 **/

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET: getProducts")

	// fetch from the database
	lp := data.GetProducts()

	// serialize the data to list
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST: addProducts")
	// get from the request(user)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	// add to the database
	data.AddProduct(&prod)

	p.l.Printf("Product: %#v\n", &prod)
}

// put the entire object, expect the id in the URI
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	// gorilla auto
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to decode id", http.StatusBadRequest)
		return
	}

	p.l.Println("PUT: updateProducts", id)
	// get from the request(user)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNOTFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

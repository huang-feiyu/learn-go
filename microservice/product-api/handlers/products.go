package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"huangblog.com/product-api/data"
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
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(rw, r)

	case http.MethodPost:
		p.addProducts(rw, r)

	case http.MethodPut:
		// expect the id in the uri
		rg := regexp.MustCompile("/([0-9]+)")
		g := rg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI not a number", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProducts(id, rw, r)

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET: getProducts")

	// fetch from the database
	lp := data.GetProducts()

	// serialize the data to list
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST: addProducts")

	// get from the request(user)
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
		return
	}

	// add to the database
	data.AddProduct(prod)

	p.l.Printf("Product: %#v\n", prod)
}

// put the entire object, expect the id in the URI
func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT: updateProducts")

	// get from the request(user)
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNOTFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/nduhiu17/go-micro-services-example/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)
		//EXPECT ID IN THE URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		group := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			p.l.Println("Invalid URL more than one id", group)
			http.Error(rw, "INVALID URL", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			p.l.Println("Invalid URL more than one capture group")
			http.Error(rw, "INVALID URL", http.StatusBadRequest)
			return
		}

		idstring := group[0][1]

		id, err := strconv.Atoi(idstring)

		if err != nil {
			p.l.Println("Invalid URL unable to convert to number", idstring)
			http.Error(rw, "INVALID URL", http.StatusBadRequest)
			return
		}

		p.l.Println("got id", id)

		p.updateProduct(id, rw, r)
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := &data.Product{}

	err := prod.FromJson(r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle UPDATE product")

	prod := &data.Product{}

	err := prod.FromJson(r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

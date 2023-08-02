package users

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	Endpoints  struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}
	CreateReq struct {
		XMLName   xml.Name `xml:"User"`
		FirstName string   `xml:"name"`
		LastName  string   `xml:"lastName"`
		Email     string   `xml:"email"`
		Phone     string   `xml:"phone"`
	}
	ErrorResp struct {
		XMLName xml.Name `xml:"Error"`
		Error   string   `xml:"mensaje"`
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(),
		Get:    makeGetEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}

}
func makeDeleteEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// logica
		err := xml.NewEncoder(w).Encode([]string{"Hola", "Mundo"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}
func makeUpdateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// logica
		fmt.Println("got /Update")
		err := xml.NewEncoder(w).Encode([]string{"Hola", "Mundo"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}
func makeGetEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// logica
		fmt.Println("got /Get")
		err := xml.NewEncoder(w).Encode([]string{"Hola", "Mundo"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}
func makeGetAllEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		// logica
		err := xml.NewEncoder(w).Encode([]string{"Hola", "Mundo"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}
func makeCreateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateReq
		if err := xml.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err := xml.NewEncoder(w).Encode(ErrorResp{xml.Name{}, "Invalid request format"}); err != nil {
				log.Fatal(err)
			}
			return
		}
		/*
			bytexml, err := xml.MarshalIndent(&req, "", "	")

			fmt.Println("Error:", err)
			fmt.Println(string(bytexml))
		*/
		if req.FirstName == "" {
			w.WriteHeader(406)
			if err := xml.NewEncoder(w).Encode(ErrorResp{xml.Name{}, "name is required"}); err != nil {
				log.Fatal(err)
			}
			return
		}
		if req.LastName == "" {
			w.WriteHeader(406)
			if err := xml.NewEncoder(w).Encode(ErrorResp{xml.Name{}, "lastName is required"}); err != nil {
				log.Fatal(err)
			}
			return
		}

		fmt.Println("got /Create")
		err := xml.NewEncoder(w).Encode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}

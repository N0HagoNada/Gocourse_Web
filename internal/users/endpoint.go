package users

import (
	"encoding/xml"
	"github.com/gorilla/mux"
	"gocourse/pkg/meta"
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
		XMLName   xml.Name `xml:"CreateUser"`
		FirstName string   `xml:"name"`
		LastName  string   `xml:"lastName"`
		Email     string   `xml:"email,omitempty"`
		Phone     string   `xml:"phone,omitempty"`
	}

	UpdateReq struct {
		XMLName   xml.Name `xml:"UpdateUser"`
		FirstName *string  `xml:"name,omitempty"`
		LastName  *string  `xml:"lastName,omitempty"`
		Email     *string  `xml:"email,omitempty"`
		Phone     *string  `xml:"phone,omitempty"`
	}
	Response struct {
		XMLName xml.Name    `xml:"Response"`
		Status  int         `xml:"status"`
		Data    interface{} `xml:"data,omitempty"`
		Err     string      `xml:"error,omitempty"`
		Meta    *meta.Meta  `xml:"meta,omitempty"`
	}
)

// Capa de Presentacion | endpoints, validaciones simples
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}

}
func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		path := mux.Vars(r)
		id := path["id"]
		if err := s.Delete(id); err != nil {
			w.WriteHeader(http.StatusNotFound)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 404, nil, "Not Found User", nil}); err != nil {
				log.Println(err)
			}
			return
		}
		w.WriteHeader(200)
		err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 200, nil, "", nil})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}
func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		var req UpdateReq
		path := mux.Vars(r)
		id := path["id"]
		if err := xml.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 400, nil, "Invalid request format", nil}); err != nil {
				log.Fatal(err)
			}
			return
		}
		user, err := s.Update(id, req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 404, nil, "Not Found User", nil}); err != nil {
				log.Println(err)
			}
			return
		}
		w.WriteHeader(200)
		err = xml.NewEncoder(w).Encode(Response{xml.Name{}, 200, user, "", nil})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}
func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		path := mux.Vars(r)
		id := path["id"]
		user, err := s.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 404, nil, "Not Found User", nil}); err != nil {
				log.Println(err)
			}
		}
		err = xml.NewEncoder(w).Encode(Response{xml.Name{}, 200, user, "", nil})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
}
func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")

		v := r.URL.Query()

		filters := Filters{
			FirstName: v.Get("first_name"),
			LastName:  v.Get("last_name"),
		}
		count, err := s.Count(filters)
		if err != nil {
			w.WriteHeader(500)
			_ = xml.NewEncoder(w).Encode(&Response{xml.Name{}, 500, nil, err.Error(), nil})
			return
		}
		data, _ := meta.New(count)
		users, _ := s.GetAll(filters)
		if err := xml.NewEncoder(w).Encode(&Response{xml.Name{}, 200, users, "", data}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
		}

	}
}
func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		var req CreateReq
		if err := xml.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 400, nil, "Invalid request format", nil}); err != nil {
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
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 406, nil, "name is required", nil}); err != nil {
				log.Println(err)
			}
			return
		}
		if req.LastName == "" {
			w.WriteHeader(406)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 406, nil, "name is required", nil}); err != nil {
				log.Println(err)
			}
			return
		}
		user, err := s.Create(req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 404, nil, "Not Found User", nil}); err != nil {
				log.Println(err)
			}
			return
		}
		w.WriteHeader(200)
		err = xml.NewEncoder(w).Encode(Response{xml.Name{}, 200, user, "", nil})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}

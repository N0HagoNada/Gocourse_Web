package courses

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
		XMLName   xml.Name `xml:"CreateCourse"`
		Name      string   `xml:"name"`
		StartDate string   `xml:"startDate"`
		EndDate   string   `xml:"endDate,omitempty"`
	}

	UpdateReq struct {
		XMLName   xml.Name `xml:"UpdateCourse"`
		Name      *string  `xml:"name"`
		StartDate *string  `xml:"start_date,omitempty"`
		EndDate   *string  `xml:"end_date,omitempty"`
	}
	Response struct {
		XMLName xml.Name    `xml:"Response"`
		Status  int         `xml:"status"`
		Data    interface{} `xml:"data,omitempty"`
		Err     string      `xml:"error,omitempty"`
		Meta    *meta.Meta  `xml:"meta,omitempty"`
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		path := mux.Vars(r)
		id := path["id"]
		course, err := s.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 404, nil, "Not Found Courses", nil}); err != nil {
				log.Println(err)
			}
		}
		err = xml.NewEncoder(w).Encode(Response{xml.Name{}, 200, course, "", nil})
		if err != nil {
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
		if req.Name == "" {
			w.WriteHeader(406)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 406, nil, "name is required", nil}); err != nil {
				log.Println(err)
			}
			return
		}
		course, err := s.Create(req.Name, req.StartDate, req.EndDate)
		if err != nil {
			w.WriteHeader(400)
			_ = xml.NewEncoder(w).Encode(&Response{XMLName: xml.Name{}, Status: 400, Err: err.Error()})
			return
		}
		w.WriteHeader(200)
		_ = xml.NewEncoder(w).Encode(&Response{Status: 200, Data: course})
	}
}
func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		path := mux.Vars(r)
		id := path["id"]
		if err := s.Delete(id); err != nil {
			w.WriteHeader(http.StatusNotFound)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 404, nil, "Not Found Course", nil}); err != nil {
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
		user, err := s.Update(id, req.Name, req.StartDate, req.EndDate)
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
func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")

		v := r.URL.Query()

		filters := Filters{
			Name: v.Get("name"),
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

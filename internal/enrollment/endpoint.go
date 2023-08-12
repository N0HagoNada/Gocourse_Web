package enrollment

import (
	"encoding/xml"
	"gocourse/pkg/meta"
	"log"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
	}
	CreateReq struct {
		XMLName  xml.Name `xml:"EnrollUser"`
		UserId   string   `xml:"user_id"`
		CourseID string   `xml:"course_id"`
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
		if req.UserId == "" || req.CourseID == "" {
			w.WriteHeader(406)
			if err := xml.NewEncoder(w).Encode(Response{xml.Name{}, 406, nil, "name is required", nil}); err != nil {
				log.Println(err)
			}
			return
		}
		course, err := s.Create(req.UserId, req.CourseID)
		if err != nil {
			w.WriteHeader(400)
			_ = xml.NewEncoder(w).Encode(&Response{XMLName: xml.Name{}, Status: 400, Err: err.Error()})
			return
		}
		w.WriteHeader(200)
		_ = xml.NewEncoder(w).Encode(&Response{Status: 200, Data: course})
	}
}

package responses

import (
	"afera-projects/internal/model"
	"encoding/json"
	"net/http"
)

type ProjectsResponse struct {
	Projects  []*model.Project `json:"projects"`
	FullCount int              `json:"full_count"`
	FullPage  int              `json:"full_page"`
}

type Response struct {
	Response interface{} `json:"response"`
}

// Responses 200

func ResponseProjects200(writer http.ResponseWriter, projects []*model.Project, fullCount int, fullPage int) {
	msg := Response{
		Response: ProjectsResponse{projects, fullCount, fullPage},
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(msg)
}

func Response200(writer http.ResponseWriter, response interface{}) {
	msg := Response{
		Response: response,
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(msg)
}

func Response201(writer http.ResponseWriter, response interface{}) {
	msg := Response{
		Response: response,
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(msg)
}

// Responses 400

func Response400(writer http.ResponseWriter, response string) {
	msg := Response{
		Response: response,
	}

	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(msg)
}

func Response404(writer http.ResponseWriter, response string) {
	msg := Response{
		Response: response,
	}

	writer.WriteHeader(http.StatusNotFound)
	json.NewEncoder(writer).Encode(msg)
}

// Responses 500

func Response500(writer http.ResponseWriter, response string) {
	msg := Response{
		Response: response,
	}

	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(msg)
}

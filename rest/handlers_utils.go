package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"gopkg.in/logex.v1"
)

func setJsonContentType(writer http.ResponseWriter) {
	writer.Header().Set(contentTypeHeaderKey, applicationJsonValue)
}

func setStatusOk(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
}

func setStatusCreated(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusCreated)
}

func setStatusAccepted(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusAccepted)
}

func setStatusNotFound(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusNotFound)
}

func setStatusInternalServerError(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusInternalServerError)
}

func setStatusBadRequest(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusBadRequest)
}

func buildErrorResponse(request string, errorMsg string) *errorResponse {
	return &errorResponse{Request: request, Message: errorMsg}
}

func returnErrorResponse(writer http.ResponseWriter, request, errorMsg string) {
	err := json.NewEncoder(writer).Encode(buildErrorResponse(request, errorMsg))
	if err != nil {
		logex.Errorf("Error on %s (encode ERROR response): %s - No response back to client",
			request, err.Error())
	}
}

func getProductIdFromUrl(request *http.Request) (int, error) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars[productIdKey])
	if err != nil {
		return -1, err
	}
	return id, nil
}

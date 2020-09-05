package rest

import (
	"encoding/json"
	"fmt"
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

func getProductIdFromRequest(request *http.Request) (int, error) {
	queryParams := request.URL.Query()
	idValues, idValuesFound := queryParams[productIdKey]
	idStr := idValues[0]
	if !idValuesFound || len(idStr) < 1 {
		return -1, fmt.Errorf("%s query param not found", productIdKey)
	}
	return strconv.Atoi(idStr)
}

package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bygui86/go-gorm/model"
	"gopkg.in/logex.v1"
)

const (
	errorMessageFormat         = "Error on %s (%s): %s"
	errorEncodeResponseMessage = "encode response"
	errorDecodeRequestMessage  = "decode request"
)

func (r *RestInterfaceImpl) getProducts(writer http.ResponseWriter, request *http.Request) {
	requestStr := "get products"
	logex.Info(requestStr)

	products, dbErr := r.dbInterface.GetProducts()
	if dbErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, "query to database", dbErr.Error())
		setStatusInternalServerError(writer)
		returnErrorResponse(writer, requestStr, dbErr.Error())
	}

	setJsonContentType(writer)
	setStatusOk(writer)
	jsonErr := json.NewEncoder(writer).Encode(products)
	if jsonErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, jsonErr.Error())
		setStatusInternalServerError(writer)
		returnErrorResponse(writer, requestStr, jsonErr.Error())
	}
}

func (r *RestInterfaceImpl) getProductById(writer http.ResponseWriter, request *http.Request) {
	requestStr := "get product by id"
	logex.Info(requestStr)

	id, paramErr := getProductIdFromUrl(request)
	if paramErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, "retrieve query param", paramErr.Error())
		setStatusBadRequest(writer)
		returnErrorResponse(writer, requestStr, paramErr.Error())
		return
	}

	logex.Infof("Search for product by id %d", id)
	product, dbErr := r.dbInterface.GetProductById(uint(id))
	if dbErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, "query to database", dbErr.Error())
		setStatusBadRequest(writer)
		returnErrorResponse(writer, requestStr, dbErr.Error())
		return
	}

	if product != nil {
		setJsonContentType(writer)
		setStatusOk(writer)
		jsonErr := json.NewEncoder(writer).Encode(product)
		if jsonErr != nil {
			logex.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, jsonErr.Error())
			setStatusInternalServerError(writer)
			returnErrorResponse(writer, requestStr, jsonErr.Error())
		}
		return
	}

	logex.Warn(fmt.Sprintf("Product with id %d not found", id))
	setStatusNotFound(writer)
	returnErrorResponse(writer, requestStr, "product not found")
}

func (r *RestInterfaceImpl) createProduct(writer http.ResponseWriter, request *http.Request) {
	requestStr := "create product"
	logex.Info(requestStr)

	var product *model.Product
	decErr := json.NewDecoder(request.Body).Decode(&product)
	if decErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, errorDecodeRequestMessage, decErr.Error())
		setStatusBadRequest(writer)
		returnErrorResponse(writer, requestStr, decErr.Error())
		return
	}

	product, dbErr := r.dbInterface.CreateProduct(product)
	if dbErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, "query to database", dbErr.Error())
		setStatusBadRequest(writer)
		returnErrorResponse(writer, requestStr, dbErr.Error())
		return
	}

	setJsonContentType(writer)
	setStatusCreated(writer)
	encErr := json.NewEncoder(writer).Encode(product)
	if encErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, encErr.Error())
		setStatusInternalServerError(writer)
		returnErrorResponse(writer, requestStr, encErr.Error())
	}
}

func (r *RestInterfaceImpl) updateProduct(writer http.ResponseWriter, request *http.Request) {
	requestStr := "update product"
	logex.Info(requestStr)

	var product *model.Product
	decErr := json.NewDecoder(request.Body).Decode(&product)
	if decErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, errorDecodeRequestMessage, decErr.Error())
		setStatusBadRequest(writer)
		returnErrorResponse(writer, requestStr, decErr.Error())
		return
	}

	product, dbErr := r.dbInterface.UpdateProduct(product)
	if dbErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, "query to database", dbErr.Error())
		setStatusBadRequest(writer)
		returnErrorResponse(writer, requestStr, dbErr.Error())
		return
	}

	setJsonContentType(writer)
	setStatusAccepted(writer)
	encErr := json.NewEncoder(writer).Encode(product)
	if encErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, errorEncodeResponseMessage, encErr.Error())
		setStatusInternalServerError(writer)
		returnErrorResponse(writer, requestStr, encErr.Error())
	}
}

func (r *RestInterfaceImpl) deleteProductById(writer http.ResponseWriter, request *http.Request) {
	requestStr := "delete product by id"
	logex.Info(requestStr)

	id, paramErr := getProductIdFromUrl(request)
	if paramErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, "retrieve query param", paramErr.Error())
		setStatusBadRequest(writer)
		returnErrorResponse(writer, requestStr, paramErr.Error())
		return
	}

	logex.Infof("Search for product by id %d", id)
	dbErr := r.dbInterface.DeleteProductById(uint(id))
	if dbErr != nil {
		logex.Errorf(errorMessageFormat, requestStr, "query to database", dbErr.Error())
		setStatusBadRequest(writer)
		returnErrorResponse(writer, requestStr, dbErr.Error())
		return
	}

	logex.Infof("Product with id %d deleted", id)
	setStatusOk(writer)
}

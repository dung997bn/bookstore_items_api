package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/dung997bn/bookstore_utils-go/resterrors"

	"github.com/dung997bn/bookstore_items_api/domain/items"
	"github.com/dung997bn/bookstore_items_api/domain/queries"
	"github.com/dung997bn/bookstore_items_api/services"
	"github.com/dung997bn/bookstore_items_api/utils/httputils"
	"github.com/dung997bn/bookstore_oauth-go/oauth"
)

var (
	//ItemController type
	ItemController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

//Create func
func (i *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if errOauth := oauth.AuthenticateRequest(r); errOauth != nil {
		httputils.ResponseError(w, errOauth)
		return
	}

	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		respErr := resterrors.NewUnauthorizedError("invalid access_token")
		httputils.ResponseError(w, respErr)
		return
	}

	requestBody, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		respErr := resterrors.NewBadRequestError("invalid request body")
		httputils.ResponseError(w, respErr)
		return
	}

	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := resterrors.NewBadRequestError("invalid reques body")
		httputils.ResponseError(w, respErr)
		return
	}

	itemRequest.Seller = sellerID

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		httputils.ResponseError(w, createErr)
		return
	}
	fmt.Println(result)
	httputils.ResponseJSON(w, http.StatusCreated, result)
}

func (i *itemsController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])
	item, err := services.ItemsService.GetByID(itemID)
	if err != nil {
		httputils.ResponseError(w, err)
	}
	httputils.ResponseJSON(w, http.StatusOK, item)
}

func (i *itemsController) Search(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := resterrors.NewBadRequestError("invalid json body")
		httputils.ResponseError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := resterrors.NewBadRequestError("invalid json body")
		httputils.ResponseError(w, apiErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		httputils.ResponseError(w, searchErr)
		return
	}
	httputils.ResponseJSON(w, http.StatusOK, items)
}

func (i *itemsController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])
	item, err := services.ItemsService.Delete(itemID)
	if err != nil {
		httputils.ResponseError(w, err)
	}
	httputils.ResponseJSON(w, http.StatusOK, item)
}

//Create func
func (i *itemsController) Update(w http.ResponseWriter, r *http.Request) {
	if errOauth := oauth.AuthenticateRequest(r); errOauth != nil {
		httputils.ResponseError(w, errOauth)
		return
	}

	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		respErr := resterrors.NewUnauthorizedError("invalid access_token")
		httputils.ResponseError(w, respErr)
		return
	}

	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])

	requestBody, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		respErr := resterrors.NewBadRequestError("invalid request body")
		httputils.ResponseError(w, respErr)
		return
	}

	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := resterrors.NewBadRequestError("invalid reques body")
		httputils.ResponseError(w, respErr)
		return
	}

	itemRequest.Seller = sellerID
	itemRequest.ID = itemID

	result, createErr := services.ItemsService.Update(&itemRequest)
	if createErr != nil {
		httputils.ResponseError(w, createErr)
		return
	}
	httputils.ResponseJSON(w, http.StatusCreated, result)
}

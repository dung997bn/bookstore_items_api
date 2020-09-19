package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dung997bn/bookstore_utils-go/resterrors"

	"github.com/dung997bn/bookstore_items_api/domain/items"
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
	Get(w http.ResponseWriter, r *http.Request)
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
		respErr := resterrors.NewUnauthorizedError("invalid reques body")
		httputils.ResponseError(w, respErr)
		return
	}

	requestBody, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		respErr := resterrors.NewBadRequestError("invalid reques body")
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

//Get func
func (i *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}

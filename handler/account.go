package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codepnw/hexagonal/errs"
	"github.com/codepnw/hexagonal/service"
	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewErrValidation("request body in correct format"))
		return
	}

	request := service.NewAccountRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handleError(w, errs.NewErrValidation("request body in correct format"))
		return
	}

	response, err := h.accSrv.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	responses, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/biskitsx/Hexagonal-Architecture-Go/errs"
	"github.com/biskitsx/Hexagonal-Architecture-Go/service"
	"github.com/gorilla/mux"
)

type accountHandler struct {
	accountService service.AccountService
}

func NewAccoutHandler(accountService service.AccountService) accountHandler {
	return accountHandler{accountService: accountService}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	response, err := h.accountService.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "applicaiton/json")
	json.NewEncoder(w).Encode(response)
}
func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	responses, err := h.accountService.GetAccounts(customerID)
	if err != nil {
		appErr, ok := err.(errs.AppError)
		if ok {
			w.WriteHeader(appErr.Code)
			fmt.Fprintln(w, appErr.Message)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode((responses))
}

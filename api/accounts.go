package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	utils "github.com/rd67/go-accounts/api/utils"
	db "github.com/rd67/go-accounts/db/sqlc"
)

type createAccountRequest struct {
	Name     string `json:"name" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD INR"`
}

type createAccountResponseData struct {
	Account db.Account `json:"account"`
}
type createAccountResponse struct {
	utils.ResponseCommonParameters
	Data createAccountResponseData `json:"data"`
}

func (server *Server) createAccount(context *gin.Context) {

	var data createAccountRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(utils.BAD_REQUEST_STATUS_CODE, utils.ValidationErrorResponseH(err))
		return
	}

	args := db.CreateAccountParams{
		Name:     data.Name,
		Currency: data.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(context, args)
	if err != nil {
		context.JSON(utils.ERROR_REQUEST_STATUS_CODE, utils.ErrorResponseH(err))
		return
	}

	result := createAccountResponse{
		Data: createAccountResponseData{
			Account: account,
		},
	}
	result.Message = "Account created successfully"
	result.Status_Code = utils.CREATED_REQUEST_STATUS_CODE

	context.JSON(http.StatusCreated, result)
}

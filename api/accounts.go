package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	utils "github.com/rd67/go-accounts/api/utils"
	db "github.com/rd67/go-accounts/db/sqlc"
)

/*
	Account creates function
*/

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

	context.JSON(utils.CREATED_REQUEST_STATUS_CODE, result)
}

/*
Account get details function
*/
type getAccountRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

type getAccountResponseData struct {
	Account db.Account `json:"account"`
}
type getAccountResponse struct {
	utils.ResponseCommonParameters
	Data getAccountResponseData `json:"data"`
}

func (server *Server) getAccount(context *gin.Context) {

	var data getAccountRequest

	if err := context.ShouldBindUri(&data); err != nil {
		context.JSON(utils.BAD_REQUEST_STATUS_CODE, utils.ValidationErrorResponseH(err))
		return
	}

	account, err := server.store.GetAccount(context, data.Id)
	if err != nil {

		if err == sql.ErrNoRows {
			context.JSON(utils.NOT_FOUND_REQUEST_STATUS_CODE, utils.NotFoundErrorResponseH(err))
			return
		}

		context.JSON(utils.ERROR_REQUEST_STATUS_CODE, utils.ErrorResponseH(err))
		return
	}

	result := getAccountResponse{
		Data: getAccountResponseData{
			Account: account,
		},
	}
	result.Message = "Account details"
	result.Status_Code = utils.CREATED_REQUEST_STATUS_CODE

	context.JSON(utils.SUCCESS_REQUEST_STATUS_CODE, result)

}

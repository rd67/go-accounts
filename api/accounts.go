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

	// Creating response
	response := createAccountResponse{
		Data: createAccountResponseData{
			Account: account,
		},
	}
	response.Message = "Account created successfully"
	response.Status_Code = utils.CREATED_REQUEST_STATUS_CODE

	context.JSON(response.Status_Code, response)
}

/*
Account get details function
*/
type getAccountRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
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

	// Creating response
	response := getAccountResponse{
		Data: getAccountResponseData{
			Account: account,
		},
	}
	response.Message = "Account details"
	response.Status_Code = utils.CREATED_REQUEST_STATUS_CODE

	context.JSON(response.Status_Code, response)

}

/*
Account get details function
*/
type listAccountsRequest struct {
	PageId int64 `form:"page_id" binding:"required,min=1"`
	Limit   int64 `form:"limit" binding:"required,min=10,max=100"`
}

type listAccountsResponseData struct {
	Count    int64        `json:"count"`
	Accounts []db.Account `json:"accounts"`
}
type listAccountsResponse struct {
	utils.ResponseCommonParameters
	Data listAccountsResponseData `json:"data"`
}

func (server *Server) listAccounts(context *gin.Context) {

	var data listAccountsRequest

	if err := context.ShouldBindQuery(&data); err != nil {
		context.JSON(utils.BAD_REQUEST_STATUS_CODE, utils.ValidationErrorResponseH(err))
		return
	}

	// Accounts listings
	args := db.ListAccountsParams{
		Limit:  int32(data.Limit),
		Offset: int32((data.PageId - 1) * data.Limit),
	}
	records, err := server.store.ListAccounts(context, args)
	if err != nil {
		context.JSON(utils.ERROR_REQUEST_STATUS_CODE, utils.ErrorResponseH(err))
		return
	}

	//	Accounts countings
	count, err := server.store.CountAccounts(context)
	if err != nil {
		context.JSON(utils.ERROR_REQUEST_STATUS_CODE, utils.ErrorResponseH(err))
		return
	}

	// Creating response
	response := listAccountsResponse{
		Data: listAccountsResponseData{
			Accounts: records,
			Count:    count,
		},
	}
	response.Status_Code = utils.SUCCESS_REQUEST_STATUS_CODE
	response.Message = "Listing"

	context.JSON(response.Status_Code, response)
}

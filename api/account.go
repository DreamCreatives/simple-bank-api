package api

import (
	"database/sql"
	"errors"
	db "github.com/DreamCreatives/simplebank/db/sqlc"
	"github.com/DreamCreatives/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req *CreateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.MakeErrorResponse(err))
		return
	}

	dbArg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}

	acc, err := server.store.CreateAccount(ctx, dbArg)

	if err != nil {

		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			switch pgErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, util.MakeErrorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, util.MakeErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)
	return
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req *GetAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.MakeErrorResponse(err))
		return
	}

	dbAccount, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, util.MakeErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, util.MakeErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, dbAccount)
	return
}

type GetAccountsRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Limit int32 `form:"limit" binding:"required,min=1,max=25"`
}

func (server *Server) getAccounts(ctx *gin.Context) {
	var req GetAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.MakeErrorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Page,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.MakeErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
	return
}

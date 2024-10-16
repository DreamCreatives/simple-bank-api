package api

import (
	"database/sql"
	"errors"
	"fmt"
	db "github.com/DreamCreatives/simplebank/db/sqlc"
	"github.com/DreamCreatives/simplebank/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("Cannot parse transfer request.Error: %s", err)
		ctx.JSON(http.StatusBadRequest, util.MakeErrorResponse(err))
		return
	}

	if server.validateAccounts(ctx, req.FromAccountID, req.Currency) == false {
		return
	}

	if server.validateAccounts(ctx, req.ToAccountID, req.Currency) == false {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.MakeErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validateAccounts(ctx *gin.Context, accountID int64, currency string) bool {
	acc, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, fmt.Sprintf("cannot find acccount with ID: %d", accountID))
			return false
		}
	}

	if acc.Currency != currency {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("cannot transfer money with currency: %s. Expected: %s", currency, acc.Currency))
		return false
	}

	return true
}

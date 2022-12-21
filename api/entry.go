package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/obasootom/ajo/db/sqlc"
)




type Entry struct {
   Owner string `json:"owner" form:"owner" binding:"required"`
   Ammount int  `json:"ammount" form:"ammount" binding:"required"`
}


func (server *Server) createEntry(ctx *gin.Context) {
	var req Entry

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,ErroResponse(err) )
	}

	arg := db.CreateEntryParams{
		Owner: req.Owner,
		Ammount: int64(req.Ammount),
	}

	entry, err := server.store.CreateEntry(ctx,arg)
	if err != nil {
		if pkErr, ok := err.(*pq.Error); ok {
			switch pkErr.Code {
			case "unique violation":
				ctx.JSON(http.StatusForbidden,ErroResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError,ErroResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, entry)
}
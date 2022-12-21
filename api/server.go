package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/obasootom/ajo/db/sqlc"
)

type Server struct {
	store *db.Store
	route *gin.Engine
}

func NewServer(store *db.Store) Server {
	server := Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/entry",server.createEntry)
	server.route = router
	return server
}

func ErroResponse(err error) gin.H {
	return gin.H{"err":err.Error() }
}

func (server *Server) Run(address string ) error{
   return server.route.Run(address)
}

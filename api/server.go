package api

import (
	db "github.com/amyunfei/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// 创建一个 http 服务
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.gerAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// 启动HTTP服务
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

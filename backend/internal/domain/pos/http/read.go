package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/unayip1710/unayip-pos/utils/response"
)

func (handler Handler) Ping(c *gin.Context) {
	response.EndWithStatus(c, http.StatusOK, "pong")
}

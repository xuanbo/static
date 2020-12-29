package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/static/index.html")
}

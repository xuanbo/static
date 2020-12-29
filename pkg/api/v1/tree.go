package v1

import (
	"net/http"
	"os"
	"strings"

	"github.com/xuanbo/static/pkg/api"
	"github.com/xuanbo/static/pkg/core/fs"

	"github.com/gin-gonic/gin"
)

// Tree 文件树
func Tree(c *gin.Context) {
	root := &fs.Tree{Title: "project"}
	if err := fs.Walk(projectPath, root, 0, 5); err != nil {
		c.JSON(http.StatusBadRequest, api.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, api.Success(root))
}

// Remove 删除
func Remove(c *gin.Context) {
	var data fs.Tree
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, api.Fail(err.Error()))
		return
	}
	path := data.Path
	// 确保只删除项目下的资源，防止恶意删除
	if strings.Contains(path, "..") || !strings.HasPrefix(path, "project") {
		c.JSON(http.StatusOK, api.Fail("不合法的路径"))
		return
	}
	if err := os.RemoveAll(path); err != nil {
		c.JSON(http.StatusOK, api.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, api.Success(nil))
}

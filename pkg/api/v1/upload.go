package v1

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/xuanbo/static/pkg/api"

	"github.com/b3log/gulu"
	"github.com/gin-gonic/gin"
)

var projectPath = "./project"

// Upload 上传
func Upload(c *gin.Context) {
	// 从请求中读取文件
	f, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, api.Fail(err.Error()))
		return
	}

	// 文件名包含相对路径，绝对路径
	if strings.Contains(f.Filename, "..") || strings.Contains(f.Filename, "/") {
		c.JSON(http.StatusBadRequest, api.Fail("文件名不合法"))
		return
	}

	// 将读取到的文件保存到本地(服务端)
	dst := path.Join(projectPath, f.Filename)
	if err := c.SaveUploadedFile(f, dst); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, api.Fail(err.Error()))
		return
	}

	unzipDir := strings.TrimSuffix(f.Filename, ".zip")
	unzipDst := path.Join(projectPath, unzipDir)
	// 移除旧文件
	if err := os.RemoveAll(unzipDst); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, api.Fail(err.Error()))
		return
	}
	// 解压
	if err := gulu.Zip.Unzip(dst, unzipDst); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, api.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, api.Success(unzipDir))
}

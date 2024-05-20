package file

import (
	"fmt"

	"github.com/changwei4869/wedding/modules/minio"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(400, "No file is received")
		return
	}

	src, err := file.Open()
	if err != nil {
		c.String(500, "Unable to open file")
		return
	}
	defer src.Close()

	client, err := minio.NewMinIOClient()
	if err != nil {
		c.String(500, "Failed to create MinIO client")
		return
	}

	contentType := file.Header.Get("Content-Type")
	fileName := uuid.NewString()
	result, err := minio.UploadFileToMinIO(client, "mybucket", fileName, src, file.Size, contentType)
	if err != nil {
		c.String(500, "Failed to upload file")
		return
	}
	c.JSON(200, gin.H{
		"path": fmt.Sprintf("%s/%s", "mybucket", result.Key),
	})
}

package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func PostUploadFile(c *gin.Context) {
	// Upload file đến CKfinder
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)

	// Save file to ckfinder
	//http://0.0.0.0:3000/ckfinder/core/connector/php/connector.php?command=FileUpload&type=Files&currentFolder=&hash=523cac8c2c3bbeda&responseType=json

}

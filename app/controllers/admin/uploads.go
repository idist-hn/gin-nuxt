package admin

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"idist-core/app/controllers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func PostUploadFile(c *gin.Context) {

	url := "https://ptitglobal.dtqt.edu.vn/ckfinder/core/connector/php/connector.php?command=FileUpload&type=Files&currentFolder=&hash=523cac8c2c3bbeda&responseType=json"
	method := "POST"

	data := bson.M{}
	fileHeader, _ := c.FormFile("upload")
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := fileHeader.Open()
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("upload", filepath.Base(fileHeader.Filename))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Upload file lỗi", errFile1)

	}
	err := writer.Close()
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Upload file lỗi", err)

	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Upload file lỗi", err)

	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Upload file lỗi", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Upload file lỗi", err)
	}

	// Parse Response from Ckfinder
	//@TODO: tao struct moi de convert format cua json sang Model.
	response := bson.M{}
	_ = json.Unmarshal(body, &response)
	data = response
	controllers.ResponseSuccess(c, http.StatusOK, "Upload file thành công", data)
}

package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/collections"
	"idist-core/app/const/mongo-constant"
	"idist-core/app/const/response"
	"idist-core/app/controllers"
	"net/http"
)

func ListDepartments(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Department{}
	entries := collections.Departments{}

	pagination := BindRequestTable(c, "priority")
	filter := pagination.CustomFilters(bson.M{})
	opts := pagination.CustomOptions(options.Find())

	if entries, err = entry.Find(filter, opts); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusNotFound, response.NOT_FOUND, err)
	}

	for i := 0; i < len(entries); i++ {
		entries[i].Preload("count_user", "parent")
	}
	pagination.Total, _ = entry.Count(filter)
	data["entries"] = entries
	data["pagination"] = pagination
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func CreateDepartment(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Department{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.CREATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.CREATE_SUCCESS, data)
}

func ReadDepartment(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Department{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}

	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, MongoConstant.NOT_FOUND, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func UpdateDepartment(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Department{}
	exist := collections.Department{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}

	if err = exist.First(filter); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusNotFound, MongoConstant.NOT_FOUND, err)
	}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil || entry.ID != entryId {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}

	if err = entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.UPDATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.UPDATE_SUCCESS, data)

}
func DeleteDepartment(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Department{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}

	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, MongoConstant.NOT_FOUND, err)
	}

	if err = entry.Delete(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.DELETE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.DELETE_SUCCESS, data)

}

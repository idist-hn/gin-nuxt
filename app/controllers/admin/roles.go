package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/collections"
	MongoConstant "idist-core/app/const/mongo-constant"
	"idist-core/app/const/response"
	"idist-core/app/controllers"
	"net/http"
)

func ListRoles(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}
	entries := collections.Roles{}
	pagination := BindRequestTable(c, "created_at")
	filter := pagination.CustomFilters(bson.M{"deleted_at": nil})
	opts := pagination.CustomOptions(options.Find())
	if entries, err = entry.Find(filter, opts); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, MongoConstant.QUERY_FAIL, err)
	}
	for i := 0; i < len(entries); i++ {
		entries[i].Preload("permissions", "department")
	}
	pagination.Total, _ = entry.Count(filter)
	data["entries"] = entries
	data["pagination"] = pagination
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func CreateRole(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.CREATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.CREATE_SUCCESS, data)
}

func ReadRole(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}
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

	entry.Preload("permissions", "department")
	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func UpdateRole(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}
	exist := collections.Role{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}

	if err = exist.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, MongoConstant.NOT_FOUND, err)
	}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}

	if err = entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.UPDATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.UPDATE_SUCCESS, data)

}
func DeleteRole(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Role{}
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

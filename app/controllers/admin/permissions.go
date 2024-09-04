package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"idist-core/app/collections"
	"idist-core/app/const/mongo-constant"
	"idist-core/app/const/response"
	"idist-core/app/controllers"
	"net/http"
)

func ListPermissions(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Permission{}
	entries := collections.Permissions{}
	filter := bson.M{}
	if entries, err = entry.Find(filter); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, MongoConstant.QUERY_FAIL, err)
	}
	data["entries"] = entries
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func ReadPermission(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Permission{}
	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"_id": entryId,
	}

	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, MongoConstant.NOT_FOUND, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

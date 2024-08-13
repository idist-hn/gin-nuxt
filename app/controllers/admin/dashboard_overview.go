package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/collections"
	"idist-core/app/const/response"
	"idist-core/app/controllers"
	"idist-core/helpers"
	"net/http"
)

func GetOverviewRegularPosts(c *gin.Context) {
	data := bson.M{}
	entry := collections.Article{}
	entries := collections.Articles{}
	var err error
	filter := bson.M{
		"status": "published",
		"published_at": bson.M{
			"$lte": helpers.Now(),
		},
		"deleted_at": nil,
	}
	opts := options.Find().SetSort(bson.M{"view": -1}).SetLimit(6)
	if entries, err = entry.Find(filter, opts); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusNotFound, response.NOT_FOUND, err)
	}

	data["entries"] = entries
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

// @TODO: GetOverviewRegularPages
func GetOverviewRegularPages(c *gin.Context) {
	data := bson.M{}
	entry := collections.Article{}
	entries := collections.Articles{}
	var err error
	filter := bson.M{
		"status": "published",
		"published_at": bson.M{
			"$lte": helpers.Now(),
		},
		"deleted_at": nil,
	}
	opts := options.Find().SetSort(bson.M{"view": -1}).SetLimit(6)
	if entries, err = entry.Find(filter, opts); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusNotFound, response.NOT_FOUND, err)
	}

	data["entries"] = entries
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

// @TODO: GetOverviewTraffics
func GetOverviewTraffics(c *gin.Context) {
	data := bson.M{}
	entry := collections.Article{}
	entries := collections.Articles{}
	var err error
	filter := bson.M{
		"status": "published",
		"published_at": bson.M{
			"$lte": helpers.Now(),
		},
		"deleted_at": nil,
	}
	opts := options.Find().SetSort(bson.M{"view": -1}).SetLimit(6)
	if entries, err = entry.Find(filter, opts); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusNotFound, response.NOT_FOUND, err)
	}

	data["entries"] = entries
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

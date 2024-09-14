package cie

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"idist-core/app/collections"
	"idist-core/app/const/response"
	"idist-core/app/controllers"
	"net/http"
)

func GetSubMenu(c *gin.Context) {
	data := bson.M{}
	var err error
	// Get data from database
	locationMenu := collections.MenuLocation{}
	filterLocation := bson.M{
		"alias":      "top-sub-menu",
		"language":   "vi",
		"is_used":    true,
		"deleted_at": nil,
	}
	if err = locationMenu.First(filterLocation); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err.Error())
		return
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusNotFound, response.NOT_FOUND, response.NOT_FOUND)
		return
	}

	filterMenu := bson.M{
		"menu_location_id": locationMenu.ID,
		"is_used":          true,
		"deleted_at":       nil,
	}
	menu := collections.Menu{}

	if err = menu.First(filterMenu); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err.Error())
		return
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusNotFound, response.NOT_FOUND, response.NOT_FOUND)
		return
	}

	data["entry"] = menu
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

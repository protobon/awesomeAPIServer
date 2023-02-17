package controller

import (
	"awesomeapiserver/database"
	"awesomeapiserver/httputil"
	"awesomeapiserver/model"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewDummy godoc
//
//	@Summary		Insert Dummy
//	@Description	save dummy in database
//	@Tags			dummy
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Dummy
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		401	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/dummy/new [post]
func (c *Controller) NewDummy(ctx *gin.Context) {
	var dummy model.Dummy
	if err := ctx.BindJSON(&dummy); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := dummy.QCreateDummy(database.DB); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, dummy)
}

// Dummies godoc
//
//	@Summary		Select all Dummies
//	@Description	Select all Dummies
//	@Tags			dummy
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Dummy
//	@Router			/dummy/dummies [get]
func (c *Controller) Dummies(ctx *gin.Context) {
	start, _ := strconv.Atoi(ctx.Query("start"))
	count, _ := strconv.Atoi(ctx.Query("count"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	var dummy = model.Dummy{}
	dummies, err := dummy.QGetDummies(database.DB, start, count)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, dummies)
}

// GetDummy godoc
//
//	@Summary		Select Dummy
//	@Description	Get one Dummy from ID
//	@Tags			dummy
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Dummy
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/dummy/{id} [get]
func (c *Controller) GetDummy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	dummy := model.Dummy{ID: id}
	if err = dummy.QGetDummy(database.DB); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			httputil.NewError(ctx, http.StatusNotFound, err)
		default:
			httputil.NewError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(200, dummy)
}

// UpdateDummy godoc
//
//	@Summary		Updates Dummy
//	@Description	Updates a dummy in database (id req)
//	@Tags			dummy
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Dummy
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/dummy/update [put]
func (c *Controller) UpdateDummy(ctx *gin.Context) {
	var dummy model.Dummy
	if err := ctx.BindJSON(&dummy); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := dummy.QUpdateDummy(database.DB); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			httputil.NewError(ctx, http.StatusNotFound, err)
		default:
			httputil.NewError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(200, dummy)
}

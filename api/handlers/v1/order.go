package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/FarrukhibnAkbar/bookshop-api/api/handlers/models"
	pb "github.com/FarrukhibnAkbar/bookshop-api/genproto/order_service"
	l "github.com/FarrukhibnAkbar/bookshop-api/pkg/logger"
	"github.com/FarrukhibnAkbar/bookshop-api/pkg/utils"
)

// CreateOrder ...
// @Summary CreateOrder
// @Description This API for creating a new order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order request body models.order true "orderCreateRequest"
// @Success 200 {object} models.Order
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/ [post]
func (h *handlerV1) CreateOrder(c *gin.Context) {
	var (
		body        pb.Order
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create order", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetOrder ...
// @Summary Getorder
// @Description This API for getting order detail
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.UpdateOrder
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [get]
func (h *handlerV1) GetOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().Get(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get order", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListOrders ...
// @Summary Listorders
// @Description This API for getting list of orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListOrders
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders [get]
func (h *handlerV1) Listorders(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().List(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list orders", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateOrder ...
// @Summary Updateorder
// @Description This API for updating order
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param order request body models.Order true "orderUpdateRequest"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [put]
func (h *handlerV1) UpdateOrder(c *gin.Context) {
	var (
		body        pb.order
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update order", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteOrder ...
// @Summary Deleteorder
// @Description This API for deleting order
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orders/{id} [delete]
func (h *handlerV1) DeleteOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.orderService().Delete(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete order", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

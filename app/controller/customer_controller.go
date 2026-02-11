package controller

import (
	"fmt"

	"ilicense-lite/library/http"
	"ilicense-lite/service"
	"ilicense-lite/type/input"

	"github.com/gin-gonic/gin"
)

var customerService = service.NewCustomerService()

// CustomerGet
// @Summary      获取客户信息
// @Description  通过ID查询信息详情
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "ID"
// @Success      200  {object}  http.BaseResponse[model.Customer]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]      "参数错误"
// @Failure      404  {object}  http.BaseResponse[any]      "不存在"
// @Router       /api/customer/get [get]
func CustomerGet(ctx *gin.Context) {
	var in input.CustomerGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := customerService.CustomerGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

// CustomerAdd
// @Summary      客户添加接口
// @Description  添加客户信息
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        input  body  input.CustomerAddInput  true  "添加参数"
// @Success      200  {object}  http.BaseResponse[model.Customer]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/customer/add [post]
func CustomerAdd(ctx *gin.Context) {
	var in input.CustomerAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := customerService.CustomerAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

// CustomerQuery
// @Summary      客户查询接口
// @Description  分页查询客户列表
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        input  body  input.CustomerQueryInput  true  "查询参数"
// @Success      200  {object}  http.BaseResponse[[]model.Customer]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/customer/query [post]
func CustomerQuery(ctx *gin.Context) {
	var in input.CustomerQueryInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	fmt.Println(ctx.Get("userID"))
	result, err := customerService.CustomerQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

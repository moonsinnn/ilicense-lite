package controller

import (
	"fmt"

	"ilicense-lite/library/http"
	"ilicense-lite/service"
	"ilicense-lite/type/input"

	"github.com/gin-gonic/gin"
)

var productService = service.NewProductService()

// ProductGet
// @Summary      获取产品信息
// @Description  通过ID查询信息详情
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "产品ID"
// @Success      200  {object}  http.BaseResponse[model.Product]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]      "参数错误"
// @Failure      404  {object}  http.BaseResponse[any]      "用户不存在"
// @Router       /api/product/get [get]
func ProductGet(ctx *gin.Context) {
	var in input.ProductGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := productService.ProductGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

// ProductAdd
// @Summary      产品添加接口
// @Description  添加产品信息
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        input  body  input.ProductAddInput  true  "添加参数"
// @Success      200  {object}  http.BaseResponse[model.Product]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/product/add [post]
func ProductAdd(ctx *gin.Context) {
	var in input.ProductAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := productService.ProductAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

// ProductQuery
// @Summary      产品查询接口
// @Description  分页查询产品列表
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        input  body  input.ProductQueryInput  true  "查询参数"
// @Success      200  {object}  http.BaseResponse[[]model.Product]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/product/query [post]
func ProductQuery(ctx *gin.Context) {
	var in input.ProductQueryInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	fmt.Println(ctx.Get("userID"))
	result, err := productService.ProductQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

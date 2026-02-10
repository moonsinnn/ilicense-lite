package controller

import (
	"fmt"
	"ilicense-lite/library/http"
	"ilicense-lite/service"
	"ilicense-lite/type/request"

	"github.com/gin-gonic/gin"
)

var productService = service.NewProductService()

// UserGet
// @Summary      获取用户信息
// @Description  通过用户ID查询用户详情
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "用户ID"
// @Success      200  {object}  http.BaseResponse[do.User]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]      "参数错误"
// @Failure      404  {object}  http.BaseResponse[any]      "用户不存在"
// @Router       /api/user/get [get]
func ProductGet(ctx *gin.Context) {
	var in request.ProductGetRequest
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

// UserAdd
// @Summary      用户添加接口
// @Description  添加用户信息
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body  request.UserAddRequest  true  "添加参数"
// @Success      200  {object}  http.BaseResponse[do.User]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/user/add [post]
func ProductAdd(ctx *gin.Context) {
	var in request.ProductAddRequest
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

// UserQuery
// @Summary      用户查询接口
// @Description  分页查询用户列表
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body  request.UserQueryRequest  true  "查询参数"
// @Success      200  {object}  http.BaseResponse[[]do.User]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/user/query [post]
func ProductQuery(ctx *gin.Context) {
	var in request.ProductQueryRequest
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

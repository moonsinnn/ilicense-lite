package controller

import (
	"fmt"

	"ilicense-lite/library/http"
	"ilicense-lite/service"
	"ilicense-lite/type/input"

	"github.com/gin-gonic/gin"
)

var issuerService = service.NewIssuerService()

// IssuerGet
// @Summary      获取签发机构信息
// @Description  通过ID查询信息详情
// @Tags         Issuer
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "产品ID"
// @Success      200  {object}  http.BaseResponse[model.Issuer]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]      "参数错误"
// @Failure      404  {object}  http.BaseResponse[any]      "用户不存在"
// @Router       /api/issuer/get [get]
func IssuerGet(ctx *gin.Context) {
	var in input.IssuerGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := issuerService.IssuerGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

// IssuerDeleteOne
// @Summary      删除签发机构信息
// @Description  通过ID删除数据
// @Tags         Issuer
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "产品ID"
// @Success      200  {object}  http.BaseResponse[model.Issuer]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]      "参数错误"
// @Failure      404  {object}  http.BaseResponse[any]      "用户不存在"
// @Router       /api/issuer/delete/:id [post]
func IssuerDeleteOne(ctx *gin.Context) {
	var in input.IssuerDeleteOneInput
	if err := ctx.ShouldBindUri(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	if err := issuerService.IssuerDeleteOne(ctx.Request.Context(), &in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, nil)
}

// IssuerDelete
// @Summary      删除签发机构信息
// @Description  通过ID列表删除数据
// @Tags         Issuer
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "产品ID"
// @Success      200  {object}  http.BaseResponse[model.Issuer]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]      "参数错误"
// @Failure      404  {object}  http.BaseResponse[any]      "用户不存在"
// @Router       /api/issuer/delete [post]
func IssuerDelete(ctx *gin.Context) {
	var in input.IssuerDeleteInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	if err := issuerService.IssuerDelete(ctx.Request.Context(), &in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, nil)
}

// IssuerAdd
// @Summary      签发机构添加接口
// @Description  添加签发机构信息
// @Tags         Issuer
// @Accept       json
// @Produce      json
// @Param        input  body  input.IssuerAddInput  true  "添加参数"
// @Success      200  {object}  http.BaseResponse[model.Issuer]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/issuer/add [post]
func IssuerAdd(ctx *gin.Context) {
	var in input.IssuerAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := issuerService.IssuerAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

// IssuerQuery
// @Summary      签发机构查询接口
// @Description  分页查询签发机构列表
// @Tags         Issuer
// @Accept       json
// @Produce      json
// @Param        input  body  input.IssuerQueryInput  true  "查询参数"
// @Success      200  {object}  http.BaseResponse[[]model.Issuer]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/issuer/query [post]
func IssuerQuery(ctx *gin.Context) {
	var in input.IssuerQueryInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	fmt.Println(ctx.Get("userID"))
	result, err := issuerService.IssuerQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

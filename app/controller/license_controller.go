package controller

import (
	"fmt"

	"ilicense-lite/library/http"
	"ilicense-lite/service"
	"ilicense-lite/type/input"

	"github.com/gin-gonic/gin"
)

var licenseService = service.NewLicenseService()

// LicenseGet
// @Summary      获取许可证信息
// @Description  通过ID查询信息详情
// @Tags         License
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "产品ID"
// @Success      200  {object}  http.BaseResponse[model.License]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]      "参数错误"
// @Failure      404  {object}  http.BaseResponse[any]      "用户不存在"
// @Router       /api/license/get [get]
func LicenseGet(ctx *gin.Context) {
	var in input.LicenseGetInput
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := licenseService.LicenseGet(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

// LicenseAdd
// @Summary      许可证添加接口
// @Description  添加许可证信息
// @Tags         License
// @Accept       json
// @Produce      json
// @Param        input  body  input.LicenseAddInput  true  "添加参数"
// @Success      200  {object}  http.BaseResponse[model.License]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/license/add [post]
func LicenseAdd(ctx *gin.Context) {
	var in input.LicenseAddInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := licenseService.LicenseAdd(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

// LicenseQuery
// @Summary      许可证查询接口
// @Description  分页查询许可证列表
// @Tags         License
// @Accept       json
// @Produce      json
// @Param        input  body  input.LicenseQueryInput  true  "查询参数"
// @Success      200  {object}  http.BaseResponse[[]model.License]  "成功响应"
// @Failure      400  {object}  http.BaseResponse[any]  "参数错误"
// @Failure      500  {object}  http.BaseResponse[any]  "内部错误"
// @Router       /api/license/query [post]
func LicenseQuery(ctx *gin.Context) {
	var in input.LicenseQueryInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	fmt.Println(ctx.Get("userID"))
	result, err := licenseService.LicenseQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func LicenseActivate(ctx *gin.Context) {
	var in input.LicenseActivateInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	fmt.Println(ctx.Get("userID"))
	result, err := licenseService.LicenseActivate(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func LicenseRenew(ctx *gin.Context) {
	var in input.LicenseRenewInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	fmt.Println(ctx.Get("userID"))
	result, err := licenseService.LicenseRenew(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

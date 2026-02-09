package controller

import (
	"fmt"
	url2 "net/url"
	"time"

	"ilicense-lite/library/http"
	"ilicense-lite/service"
	"ilicense-lite/type/request"

	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

func UserLogin(ctx *gin.Context) {
	var in request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserLogin(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

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
func UserGet(ctx *gin.Context) {
	var in request.UserGetRequest
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserGet(ctx.Request.Context(), &in)
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
func UserAdd(ctx *gin.Context) {
	var in request.UserAddRequest
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserAdd(ctx.Request.Context(), &in)
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
func UserQuery(ctx *gin.Context) {
	var in request.UserQueryRequest
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	fmt.Println(ctx.Get("userID"))
	result, err := userService.UserQuery(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserSignIn(ctx *gin.Context) {
	clientID := "flashcat.cloud"
	redirectURI := "http://localhost:8080/api/user/sign/back?source=sso"
	state := time.Now().Unix()
	url := "https://auth.oncallbox.com/realms/oncallbox/protocol/openid-connect/auth?client_id=%s&nonce=%d&redirect_uri=%s&response_type=code&scope=openid+profile+email+phone&state=%d"
	rawURL := fmt.Sprintf(url, clientID, state, url2.QueryEscape(redirectURI), state)
	http.JsonResponse(ctx, rawURL)
}

func UserSignBack(ctx *gin.Context) {
	var in request.UserSignBackRequest
	if err := ctx.ShouldBindQuery(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, in)
}

package controller

import (
	"errors"
	"strconv"

	"ilicense-lite/library/http"
	"ilicense-lite/service"
	"ilicense-lite/type/input"

	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

func getCurrentUserID(ctx *gin.Context) (uint64, error) {
	v, ok := ctx.Get("userID")
	if !ok {
		return 0, errors.New("未登录")
	}
	s, ok := v.(string)
	if !ok || s == "" {
		return 0, errors.New("未登录")
	}
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, errors.New("用户身份无效")
	}
	return id, nil
}

func UserRegister(ctx *gin.Context) {
	var in input.UserRegisterInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserRegister(ctx.Request.Context(), &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserLogin(ctx *gin.Context) {
	var in input.UserLoginInput
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

func UserProfileGet(ctx *gin.Context) {
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserProfileGet(ctx.Request.Context(), userID)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserProfileUpdate(ctx *gin.Context) {
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	var in input.UserProfileUpdateInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	result, err := userService.UserProfileUpdate(ctx.Request.Context(), userID, &in)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, result)
}

func UserPasswordUpdate(ctx *gin.Context) {
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	var in input.UserPasswordUpdateInput
	if err := ctx.ShouldBindJSON(&in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	if err := userService.UserPasswordUpdate(ctx.Request.Context(), userID, &in); err != nil {
		http.JsonResponse(ctx, err)
		return
	}
	http.JsonResponse(ctx, nil)
}

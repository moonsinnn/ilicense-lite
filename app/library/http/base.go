package http

import (
	"net/http"

	"ilicense-lite/library/code"

	"github.com/gin-gonic/gin"
)

type BaseResponse[T any] struct {
	Code    int    `json:"code" xml:"code" example:"0"`         // 状态码
	Message string `json:"message" xml:"message" example:"ok"`  // 消息
	Data    T      `json:"data,omitempty" xml:"data,omitempty"` // 实际数据
}

func JsonResponse(ctx *gin.Context, v any) {
	ctx.JSON(http.StatusOK, wrapBaseResponse(v))
}

func wrapBaseResponse(v any) BaseResponse[any] {
	var resp BaseResponse[any]
	switch data := v.(type) {
	case *code.Code:
		resp.Code = data.Code
		resp.Message = data.Message
	case code.Code:
		resp.Code = data.Code
		resp.Message = data.Message
	case error:
		resp.Code = BusinessCodeError
		resp.Message = data.Error()
	default:
		resp.Code = BusinessCodeOK
		resp.Message = BusinessMsgOk
		resp.Data = v
	}

	return resp
}

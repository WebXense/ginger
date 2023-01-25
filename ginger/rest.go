package ginger

import (
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

type Response struct {
	Success    bool        `json:"success"`
	Error      *Error      `json:"error,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func OK(ctx *gin.Context, data interface{}, p *Pagination) {
	resp := &Response{
		Success:    true,
		Data:       data,
		Pagination: p,
	}
	ctx.Set("response", resp) // write to ctx for perform tests
	ctx.JSON(200, resp)
}

func ERR(ctx *gin.Context, errCode string, errMsg string, data any) {
	resp := &Response{
		Success: false,
		Error: &Error{
			Code:    errCode,
			Message: errMsg,
		},
		Data: data,
	}
	ctx.Set("response", resp) // write to ctx for perform tests
	ctx.JSON(200, resp)
}

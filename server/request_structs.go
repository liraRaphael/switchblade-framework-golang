package server

import (
	"github.com/gin-gonic/gin"
)

type RestResponse[B any, H any] struct {
	StatusCode int
	Body       *B
	Headers    *H
	IsError    bool
}

type RestRequest[B any, H any, P any, Q any] struct {
	Context *gin.Context

	Body    *B
	Headers *H
	Queries *Q
	Path    *P
}

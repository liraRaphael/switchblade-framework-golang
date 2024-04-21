package server

import "github.com/gofiber/fiber/v2"

type RestResponse[B any, H any] struct {
	StatusCode int
	Body       B
	Headers    H
	IsError    bool
}

type RestRequest[B any, H any, Q any, P any] struct {
	Context *fiber.Ctx

	Body    B
	Headers H
	Queries Q
	Path    P
}

package server

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/liraRaphael/golang-api-lib/validator"
)

func NewRoute[BReq, BResp, HReq, HResp, P, Q any](s *Server) *Route[BReq, BResp, HReq, HResp, P, Q] {
	return &Route[BReq, BResp, HReq, HResp, P, Q]{Server: s}
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DefineDefaultStatusCode() int {
	if r.Response.Body == nil {
		return 204
	}

	if strings.EqualFold(r.Method, "POST") {
		return 201
	}

	return 200
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DefaultCallbackFiber(c *fiber.Ctx) error {
	start := time.Now()

	request := RestRequest[BReq, HReq, P, Q]{
		Context: c,
		Body:    r.Request.Body,
		Headers: r.Request.Headers,
		Queries: r.Request.Queries,
		Path:    r.Request.Path,
	}

	// ToDo: Serealziar os campos

	valitor := validator.Get()

	if r.Validator.RequestBody.Enable && request.Body != nil {
		err := valitor.Validate(request.Body)
		if err != nil {
			response := r.Validator.RequestBody.Handle(err)
			DefineResponseContextFromRestResponse(c, response)

			return err
		}
	}

	if r.Validator.RequestHeaders.Enable && request.Headers != nil {
		err := valitor.Validate(request.Headers)
		if err != nil {
			response := r.Validator.RequestHeaders.Handle(err)
			DefineResponseContextFromRestResponse(c, response)

			return err
		}
	}

	if r.Validator.Path.Enable && request.Path != nil {
		err := valitor.Validate(request.Path)
		if err != nil {
			response := r.Validator.Path.Handle(err)
			DefineResponseContextFromRestResponse(c, response)

			return err
		}
	}

	if r.Validator.Queries.Enable && request.Queries != nil {
		err := valitor.Validate(request.Queries)
		if err != nil {
			response := r.Validator.Queries.Handle(err)
			DefineResponseContextFromRestResponse(c, response)

			return err
		}
	}

	response, err := r.Handle(request)

	if err != nil {
		responseError, errHandle := r.ExceptionHandler[err](err)
		DefineResponseContextFromRestResponse(c, responseError)

		return errHandle
	}

	if response.StatusCode == 0 {
		response.StatusCode = r.DefineDefaultStatusCode()
	}

	DefineResponseContextFromRestResponse(c, response)

	end := time.Since(start)
	return nil
}

func DefineResponseContextFromRestResponse[B, H any](c *fiber.Ctx, response RestResponse[B, H]) {
	//ToDo
	c.Context().SetStatusCode(response.StatusCode)
	c.JSON(response.Body)
}

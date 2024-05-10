package server

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/liraRaphael/golang-api-lib/validator"
)

const defaultValidatorEnable = true
const defaultDocumentationEnable = true

func NewRoute[BReq, BResp, HReq, HResp, P, Q any](s *Server, bodyRequest BReq, bodyResponse BResp, headersRequest HReq, headersResponse HResp, Path P, query Q) *Route[BReq, BResp, HReq, HResp, P, Q] {
	return &Route[BReq, BResp, HReq, HResp, P, Q]{
		Server: s,
		Request: Request[BReq, HReq, P, Q]{
			BaseRequest: BaseRequest[BReq, HReq]{
				Body:    &bodyRequest,
				Headers: &headersRequest,
			},
			Queries: &query,
			Path:    &Path,
		},
		Response: Response[BResp, HResp]{
			BaseRequest: BaseRequest[BResp, HResp]{
				Body:    &bodyResponse,
				Headers: &headersResponse,
			},
		},
		Validator: Validator{
			RequestBody: ValidatorHandle{
				Enable: defaultValidatorEnable,
			},
			RequestHeaders: ValidatorHandle{
				Enable: defaultValidatorEnable,
			},
			Path: ValidatorHandle{
				Enable: defaultValidatorEnable,
			},
			Queries: ValidatorHandle{
				Enable: defaultValidatorEnable,
			},
		},
		Documentation: Documentation{
			Enable:      defaultDocumentationEnable,
			Summary:     "",
			Description: "",
			OperationId: "",
		},
	}
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
	//ToDo: logs
	start := time.Now()

	request := RestRequest[BReq, HReq, P, Q]{
		Context: c,
		Body:    new(BReq),
		Headers: new(HReq),
		Queries: new(Q),
		Path:    new(P),
	}

	reqBody, err := r.InputDefaultBodySerealizer(c.Body())
	// ToDo: Serealziar os campos

	valitor := validator.Get()

	if r.Validator.RequestBody.Enable {
		err := valitor.Validate(request.Body)
		if err != nil {
			return r.DefineErrorHandle(c, err)
		}
	}

	if r.Validator.RequestHeaders.Enable {
		err := valitor.Validate(request.Headers)
		if err != nil {
			return r.DefineErrorHandle(c, err)
		}
	}

	if r.Validator.Path.Enable {
		err := valitor.Validate(request.Path)
		if err != nil {
			return r.DefineErrorHandle(c, err)
		}
	}

	if r.Validator.Queries.Enable {
		err := valitor.Validate(request.Queries)
		if err != nil {
			return r.DefineErrorHandle(c, err)
		}
	}

	response, err := r.Handle(request)

	if err != nil {
		return r.DefineErrorHandle(c, err)
	}

	if response.StatusCode == 0 {
		response.StatusCode = r.DefineDefaultStatusCode()
	}

	DefineResponseContextFromRestResponse(c, response)

	end := time.Since(start)
	return nil
}

func DefineResponseContextFromRestResponse[B, H any](c *fiber.Ctx, response RestResponse[B, H]) {
	c.Context().SetStatusCode(response.StatusCode)
	c.JSON(response.Body)
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DefineErrorHandle(c *fiber.Ctx, err error) error {
	response, errHandle := r.ExceptionHandler[err](err)

	DefineResponseContextFromRestResponse(c, response)
	return errHandle
}

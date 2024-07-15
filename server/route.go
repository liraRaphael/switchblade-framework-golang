package server

import (
	"strings"

	"github.com/gin-gonic/gin"
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

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DefaultCallbackFiber(c *gin.Context) {
	request := RestRequest[BReq, HReq, P, Q]{
		Context: c,
		Body:    new(BReq),
		Headers: new(HReq),
		Queries: new(Q),
		Path:    new(P),
	}

	r.bindRequester(&request)

	isValid := r.findErrorByValidatorRequest(request)
	if isValid {
		return
	}

	response, err := r.Handle(request)

	if err != nil {
		r.DefineErrorHandle(c, err)
		return
	}

	if response.StatusCode == 0 {
		response.StatusCode = r.DefineDefaultStatusCode()
	}

	DefineResponseContextFromRestResponse(c, response)
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) findErrorByValidatorRequest(request RestRequest[BReq, HReq, P, Q]) bool {
	c := request.Context
	valitor := validator.Get()

	if r.Validator.RequestBody.Enable {
		if err := valitor.Validate(request.Body); err != nil {
			r.DefineErrorHandle(c, err)
			return true
		}
	}

	if r.Validator.RequestHeaders.Enable {
		if err := valitor.Validate(request.Headers); err != nil {
			r.DefineErrorHandle(c, err)
			return true
		}
	}

	if r.Validator.Path.Enable {
		if err := valitor.Validate(request.Path); err != nil {
			r.DefineErrorHandle(c, err)
			return true
		}
	}

	if r.Validator.Queries.Enable {

		if err := valitor.Validate(request.Queries); err != nil {
			r.DefineErrorHandle(c, err)
			return true
		}
	}
	return false
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) bindRequester(request *RestRequest[BReq, HReq, P, Q]) (*RestRequest[BReq, HReq, P, Q], error) {
	c := request.Context

	if err := c.Bind(request.Body); err != nil {
		r.DefineErrorHandle(c, err)
		return nil, err
	}

	if err := c.BindQuery(request.Queries); err != nil {
		r.DefineErrorHandle(c, err)
		return nil, err
	}

	if err := c.ShouldBindHeader(request.Headers); err != nil {
		r.DefineErrorHandle(c, err)
		return nil, err
	}

	if err := c.ShouldBindUri(request.Path); err != nil {
		r.DefineErrorHandle(c, err)
		return nil, err
	}

	return request, nil
}

func DefineResponseContextFromRestResponse[B, H any](c *gin.Context, response RestResponse[B, H]) {
	c.JSON(response.StatusCode, response.Body)
}

func (r *Route[BReq, BResp, HReq, HResp, P, Q]) DefineErrorHandle(c *gin.Context, err error) {
	response, _ := r.ExceptionHandler[err](err)

	DefineResponseContextFromRestResponse(c, response)
}

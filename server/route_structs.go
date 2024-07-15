package server

type BaseRequest[B, H any] struct {
	Body    *B
	Headers *H
}

type Response[B, H any] struct {
	BaseRequest[B, H]
}

type Request[B, H, P, Q any] struct {
	BaseRequest[B, H]

	Path    *P
	Queries *Q
}

type ValidatorHandle struct {
	Enable bool
}

type Validator struct {
	RequestBody    ValidatorHandle
	RequestHeaders ValidatorHandle
	Path           ValidatorHandle
	Queries        ValidatorHandle
}

type Documentation struct {
	Enable bool

	Summary     string
	Description string
	OperationId string
}

type Route[BReq, BResp, HReq, HResp, P, Q any] struct {
	Server *Server

	Endpoint string
	Method   string

	Handle func(request RestRequest[BReq, HReq, P, Q]) (RestResponse[BResp, HResp], error)

	ExceptionHandler map[error]func(error) (RestResponse[any, any], error)

	Request  Request[BReq, HReq, P, Q]
	Response Response[BResp, HResp]

	Validator Validator

	Documentation Documentation

	OutputDefaultBodyDeserealizer func(BResp) ([]byte, error)
}

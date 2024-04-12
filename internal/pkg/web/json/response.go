package json

type BadRequestResponse struct {
	Error string
}

type NotFoundResponse struct{}

type InternalServerErrorResponse struct {
	Error string
}

type ForbiddenResponse struct{}

type UnauthorizedResponse struct{}

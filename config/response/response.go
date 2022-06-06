package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	ResponseFactory IFactoryResponse
}

func NewResponse(responseFactory IFactoryResponse) Response {
	return Response{
		ResponseFactory: responseFactory,
	}
}

func (response Response) Send(ctx *fiber.Ctx) error {
	return response.ResponseFactory.Create(ctx)
}

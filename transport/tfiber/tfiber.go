package tfiber

import (
	"github.com/apexlang/api-go/errorz"
	"github.com/apexlang/api-go/transport/httpresponse"
	"github.com/gofiber/fiber/v2"
)

const Package = "tfiber"

type RegisterFn func(router fiber.Router)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		errz := errorz.New(errorz.FromStatus(e.Code), e.Message)
		return c.Status(e.Code).JSON(errz)
	}
	errz := errorz.From(err)
	return c.Status(errz.Status).JSON(errz)
}

func Register(router fiber.Router, services ...RegisterFn) {
	for _, s := range services {
		s(router)
	}
}

func Response(c *fiber.Ctx, resp *httpresponse.Response, val interface{}, err error) error {
	if err != nil {
		e := errorz.Translate(err)
		return c.Status(e.Status).JSON(e)
	}

	c.Status(resp.Status)
	response := c.Response()
	for k, v := range resp.Header {
		for _, val := range v {
			response.Header.Add(k, val)
		}
	}

	return c.JSON(val)
}

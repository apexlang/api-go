package tfiber

import (
	"github.com/apexlang/api-go/errorz"
	"github.com/apexlang/api-go/transport/httpresponse"
	"github.com/gofiber/fiber/v2"
)

type RegisterFn func(router fiber.Router)

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

package response

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

type IFactoryResponse interface {
	Create(ctx *fiber.Ctx) error
}

type FactoryBaseResponse struct {
	Data   interface{}           `json:"data"`
	Meta   interface{}           `json:"meta"`
	Status FactoryStatusResponse `json:"status"`
}

type FactoryStatusResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func ResponseSuccess(data, meta interface{}, message string) IFactoryResponse {
	return &FactoryBaseResponse{
		Data: data,
		Meta: meta,
		Status: FactoryStatusResponse{
			Success: true,
			Code:    200,
			Message: message,
		},
	}
}

func ResponseError(data, meta interface{}, code int, message string, errorMsg error) IFactoryResponse {
	return &FactoryBaseResponse{
		Data: data,
		Meta: meta,
		Status: FactoryStatusResponse{
			Success: false,
			Code:    code,
			Message: message,
			Error:   errorMsg.Error(),
		},
	}
}

func ResponseErrorValidation(data, meta interface{}, code int, message string, errorMsg validator.ValidationErrors) IFactoryResponse {
	return &FactoryBaseResponse{
		Data: data,
		Meta: meta,
		Status: FactoryStatusResponse{
			Success: false,
			Code:    code,
			Message: message,
			Error:   buildErrorValidation(errorMsg),
		},
	}
}

func (resp FactoryBaseResponse) Create(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(resp)
}

func buildErrorValidation(errorRes validator.ValidationErrors) interface{} {
	errorMessage := map[string][]string{}

	for _, err := range errorRes {
		errKey := Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			buildErrorValidationMessage(err),
		)
	}
	return errorMessage
}

func buildErrorValidationMessage(errorRes validator.FieldError) string {
	var sb strings.Builder

	sb.WriteString("Validation failed on field '" + Underscore(errorRes.StructField()) + "'")
	sb.WriteString(", condition: " + errorRes.ActualTag())

	// Print conidition parameters, e.g. min=4 -> { 4 }
	if errorRes.Param() != "" {
		sb.WriteString(" { " + errorRes.Param() + " }")
	}

	if errorRes.Value() != nil && errorRes.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", errorRes.Value()))
	}

	return sb.String()
}

func Underscore(s string) string {
	b := buffer{
		r: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.indent()
					w = true
				}
				b.write(m)
			}
			m = unicode.ToLower(ch)
		} else {
			if m != 0 {
				b.indent()
				b.write(m)
				m = 0
				w = false
			}
			b.write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.indent()
		}
		b.write(m)
	}

	return string(b.r)
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

func (b *buffer) indent() {
	if len(b.r) > 0 {
		b.r = append(b.r, '_')
	}
}

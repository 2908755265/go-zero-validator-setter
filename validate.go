package go_zero_validator_setter

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type v struct {
	httpx.Validator
}

func (*v) Validate(r *http.Request, data any) error {
	validate := validator.New()
	return validate.Struct(data)
}

package go_zero_validator_setter

import (
	"github.com/zeromicro/go-zero/rest/httpx"
)

func init() {
	httpx.SetValidator(&v{})
}

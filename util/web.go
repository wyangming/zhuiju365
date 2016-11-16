package util

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

//得到基础的url
func BaseUrl(ctx *context.Context) string {
	request := ctx.Request
	return fmt.Sprintf("http://%s", request.Host)
}

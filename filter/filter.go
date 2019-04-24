package filter

import (
	"fmt"
	"github.com/kataras/iris"
)

//前置过滤器
func Before(ctx iris.Context) {
	fmt.Println("前置过滤器", ctx.Path())

}

//后置过滤器
func After(ctx iris.Context) {
	fmt.Println("后置过滤器", ctx.Path())
}

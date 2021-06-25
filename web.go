package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"sggproject/contoller"
)

// MainController 定义 Controller，这里我们定义了一个 struct 为 MainController，
// 充分利用了 Go 语言的组合的概念，匿名包含了 web.Controller，
// 这样我们的 MainController 就拥有了 web.Controller 的所有方法。
type MainController struct {
	web.Controller
}

// Get 定义 RESTFUL 方法，通过匿名组合之后，其实目前的 MainController 已经拥有了 Get、Post、Delete、Put 等方法，
// 这些方法是分别用来对应用户请求的 Method 函数，如果用户发起的是 POST 请求，那么就执行 Post 函数。所以这里我们定义了 MainController 的 Get 方法用来重写继承的 Get 函数，这样当用户发起 GET 请求的时候就会执行该函数。
func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
	// 这是一个用来存储输出数据的 map，可以赋值任意类型的值，这里我们只是简单举例输出两个字符串。
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	// this.TplName 就是需要渲染的模板，这里指定了 index.tpl，如果用户不设置该参数，那么默认会去到模板目录的 Controller/<方法名>.tpl 查找，例如上面的方法会去 maincontroller/get.tpl (文件、文件夹必须小写)。
	// 用户设置了模板之后系统会自动的调用 Render 函数（这个函数是在 web.Controller 中实现的），所以无需用户自己来调用渲染。
	c.TplName = "index.tpl"
	fmt.Println("id = ",c.GetString("id"))
}
// GetName 自定义方法
func (c *MainController) GetName() {
	c.Ctx.WriteString("自定义方法")
}



func main() {
	//基本的get路由
	web.Get("/hello", func(ctx *context.Context) {
		ctx.Output.Body([]byte("基本的get路由"))
	})
	//Router 注册路由，路由就是告诉 beego，当用户来请求的时候，该如何去调用相应的 Controller，这里我们注册了请求 / 的时候，请求到 MainController。
	//这里我们需要知道，Router 函数的两个参数函数，第一个是路径，第二个是 Controller 的指针。
	web.Router("/api/id", &MainController{})
	//mappingMethods 请求方式:方法名
	web.Router("/api/getName", &MainController{},"POST:GetName")
	web.Router("/user", &contoller.UserController{})
	//自定义多个静态文件处理目录
	web.SetStaticPath("/static2", "static2")
	//Run 应用，最后一步就是把在步骤 1 中初始化的 BeeApp 开启起来，其实就是内部监听了 8080 端口：Go 默认情况会监听你本机所有的 IP 上面的 8080 端口。
	web.Run()
}

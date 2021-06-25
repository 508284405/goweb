package contoller

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

// Get 定义 RESTFUL 方法，通过匿名组合之后，其实目前的 MainController 已经拥有了 Get、Post、Delete、Put 等方法，
// 这些方法是分别用来对应用户请求的 Method 函数，如果用户发起的是 POST 请求，那么就执行 Post 函数。所以这里我们定义了 MainController 的 Get 方法用来重写继承的 Get 函数，这样当用户发起 GET 请求的时候就会执行该函数。
func (c *UserController) Get() {
	//c.Ctx.WriteString("user controller")
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "508284405@qq.com"
	c.TplName = "hello.html"
}

// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dough_go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.BConfig.CopyRequestBody = true

	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/object",
			&controllers.ObjectController{},
		),
		beego.NSRouter("/user",
			&controllers.UserController{},
		),

	//beego.NSNamespace("/user",
	//	beego.NSInclude(
	//		&controllers.UserController{},
	//	),
	//),
	)
	beego.AddNamespace(ns)
}

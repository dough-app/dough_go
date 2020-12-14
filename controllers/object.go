package controllers

import (
	"dough_go/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ObjectController struct {
	beego.Controller
}

func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

func (o *ObjectController) Get() {
	users := models.GetAllUsers()
	o.Data["json"] = users
	o.ServeJSON()
	//var u models.User
	//u.Id = "12"
	//u.Password = "深圳"
	//u.Username = "q1"
	//fmt.Printf("user::%v\n", u)
	//fmt.Printf("user::%#v\n", u)

}

func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

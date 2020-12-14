package controllers

import (
	"dough_go/models"
	"encoding/json"
	"fmt"
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
	//obs := models.GetAll()
	//o.Data["json"] = obs
	//o.ServeJSON()

	fmt.Println("test")

	objectId := o.Ctx.Input.Param(":objectId")
	fmt.Println("1")
	fmt.Println(objectId)
	fmt.Println("1")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
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

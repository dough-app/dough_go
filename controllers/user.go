package controllers

import (
	"dough_go/models"
	"dough_go/tools"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Post() {

	uid1, err := tools.ReqPostBody(u.Ctx.Request, "uid")

	if err != nil {
		fmt.Printf("err::%v\n", err)
	}
	fmt.Printf("uid::%v\n", uid1)

	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Id = uid1
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

func (u *UserController) Get() {
	uid := u.GetString("uid")
	fmt.Printf("uid为::%v\n", uid)

	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
		u.ServeJSON()
	} else if uid == "json" {
		uModel := models.User{Id: "id1", Username: "name1"}
		u.Data["json"] = &uModel
		u.ServeJSON()
	} else {
		users := models.GetAllUsers()
		u.Data["json"] = users
		u.ServeJSON()
	}
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
// @example http://localhost:8080/v1/user/login
// @example http://localhost:8080/v1/user/login?username=astaxie&password=11111
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	fmt.Printf("username::%v\n", username)
	fmt.Printf("password::%v\n", password)
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

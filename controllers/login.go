package controllers

import (
	"beegodemo/controllers/utils"
	"beegodemo/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type LoginController struct {
	beego.Controller
}



func (this *LoginController) Get () {

	this.TplName="login.tpl"
}

func (this *LoginController) Post () {

   username :=this.GetString("username")
   password :=this.GetString("password")

   if username=="" ||  password=="" {
	   this.Data["json"]=beego.M{
		   "status":false,"msg":"账号或密码错误1",
	   }
   }else{

		userinfo, err :=GetUserInfo(username)
		if err==nil {
			md5password :=strings.ToLower(utils.GetMd5String(password))

			if userinfo.PassWord==md5password {
				//验证成功后，设置缓存




				this.Data["json"]=beego.M{
					"status":true,"msg":"ok",
				}
			}else{
				this.Data["json"]=beego.M{
					"status":false,"msg":"账号或密码错误2",
				}
			}
		}else{
			this.Data["json"]=beego.M{
				"status":false,"msg":"账号或密码错误3",
			}
		}
   }
	this.ServeJSON()
}

func  GetUserInfo(username string) (u models.User,err error) {
	o :=orm.NewOrm()
	user :=models.User{UserName: username}
	errr :=o.Read(&user,"UserName")

	if errr == orm.ErrNoRows {
		return user, errors.New("No User Data")
	} else if errr == orm.ErrMissPK {
		return user , errors.New("No User Data")
	} else {
		return user,nil
	}

	return user, errors.New("No User Data")
}

/*
func  (this *RegisterController) Update() {
	o := orm.NewOrm()
    user :=models.User{Id:1}
    if o.Read(&user) ==nil {
    	fmt.Println(&user)
    	user.Name="Up测试名称"
    	user.UpdateTime=time.Now()
    	if num,err :=o.Update(&user);err ==nil{
    		fmt.Println(num)
		}
	}
	this.Data["json"]=beego.M{"errcode":"200","errmsg":"修改成功"}
	this.ServeJSON()
}/*

/*
func (this *RegisterController) AddRole () {
	o := orm.NewOrm()
	role :=new(models.Role)

	role.RoleName="管理员2"
	role.CreateTime=time.Now()
	role.CreateUser="admin"

	fmt.Println(o.Insert(role))
	this.Data["json"]=beego.M{"errcode":"200","errmsg":"添加角色成功"}
	this.ServeJSON()

}*/
/*
func (this *RegisterController) AddAuth () {
o := orm.NewOrm()
	auth :=new(models.Auth)

	auth.AuthCode="auth"
	auth.AuthName="权限管理"
	auth.AuthPath="root.sys.auth"
	auth.ParentPath="root.sys"
	auth.ParentId=2
	auth.CreateTime=time.Now()
	auth.CreateUser="admin"

	fmt.Println(o.Insert(auth))
	this.Data["json"]=beego.M{"errcode":"200","errmsg":"添加权限成功"}
	this.ServeJSON()
}
*/

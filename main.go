package main

import (
	"encoding/json"
	"fmt"
	"goflow/api"
	"goflow/model"
	"goflow/util"
	"strings"

	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	// config file
	var approvemap = make(map[string]*model.ApproveList)
	fmt.Println(approvemap)
	configfile, err := os.Open("config.json")

	if err != nil {
		log.Fatalln("can't find config.json!")
	}
	defer configfile.Close()
	decoder := json.NewDecoder(configfile)
	decoder.Decode(&util.Conf)
	port := util.Conf.Port

	// database
	util.Db, err = gorm.Open(mysql.Open(util.Conf.Dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Fatalln("can't connect to database!")
	}
	// load process from db to memory

	router := gin.Default()
	err = api.LoadAllProcess()
	if err != nil {
		log.Println(err.Error())
	}
	var user *model.User
	fmt.Println(user)
	router.LoadHTMLGlob("html/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/approve", func(ctx *gin.Context) {
		appid := ctx.PostForm("appid")
		appdata, ok := approvemap[appid]
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{"code": "1", "msg": "approve data not found"})
			return
		}
		// get form data
		fdata, err := api.FindFormData(appdata.Form_id, appdata.Procinst_id)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": "1", "msg": "form data not found"})
			return
		}
		// json unmarshall fdata
		var formdata map[string]interface{}
		err = json.Unmarshal([]byte(fdata), &formdata)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": "1", "msg": "form data not found"})
			return
		}
		// get form by form id
		form, err := api.FindFormById(appdata.Form_id)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
			return
		}
		// json unmarshall fdata
		var formdesc map[string]interface{}
		err = json.Unmarshal([]byte(form.Form_desc), &formdesc)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": "1", "msg": "form data not found"})
			return
		}
		
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(form.Form))
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
			return
		}
		for key := range formdesc {
			if formdesc[key] == "input" {
				doc.Find("#" + key).SetAttr("value",fmt.Sprintf("%v", formdata[key]))
				doc.Find("#" + key).SetAttr("readonly","readonly")
			}
		}
		html, err := doc.Find("body").Html()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
			return
		}
		// rawdata
		rawdata := `<form id="approveform" action="/movetoapprove" method="post">`
		rawdata += html
		rawdata += `<br><select id="approve" name="approve"><option value ="ok" selected="selected">OK</option><option value ="ng">NG</option></select>`
		rawdata += `<br><label>reason:</label>`
		rawdata += `<br><input type="text" id="reason" name="reason">`
		rawdata += `<br><input type="submit" value="submit">`
		rawdata += `<input type="hidden" name="appid" value="` + appdata.Id + `">`
		rawdata += `<input type="hidden" name="taskid" value="` + appdata.NodeId + `">`
		rawdata += `<input type="hidden" name="procdefid" value="` + appdata.Procdef_id + `">`
		rawdata += `<input type="hidden" name="procinstid" value="` + appdata.Procinst_id + `">`
		rawdata += `<input type="hidden" name="procflowid" value="` + appdata.Procflow_id + `">`
		rawdata += `</form>`

		ctx.HTML(http.StatusOK, "approve.html", gin.H{"rawdata": rawdata})
	})

	router.GET("/deal", func(ctx *gin.Context) {

		// get mylist
		procinst, err := api.FindMyStartList(user.Id)
		if err != nil {
			log.Println(err.Error())
		}
		mylist := ""
		for _, v := range procinst {
			mylist += "<label>"
			mylist += v.Name + " CURRENT=" + v.Current_name + "</label><br>"
		}

		// get alist
		alist, err := api.FindApproveList(user.Id)
		if err != nil {
			log.Println(err.Error())
		}
		aliststr := ""

		for _, v := range alist {
			aliststr += "<div>"
			aliststr += `<input type="button" value="Approve" onclick="postApprove('` + v.Id + `')">`
			aliststr += "<label>" + v.Name + "</label>"

			aliststr += "</div><hr/>"
			approvemap[v.Id] = v
		}

		ctx.HTML(http.StatusOK, "deal.html", gin.H{"mylist": mylist, "alist": aliststr})
	})

	router.POST("/login", func(ctx *gin.Context) {
		utype := ctx.PostForm("user")
		user = GetUser(utype)
		ctx.Request.URL.Path = "/deal"
		ctx.Request.Method = "GET"
		router.HandleContext(ctx)
	})

	router.POST("/movetoapprove", func(ctx *gin.Context) {
		//save form data
		taskid := ctx.PostForm("taskid")
		procdefid := ctx.PostForm("procdefid")
		procinstid := ctx.PostForm("procinstid")
		procflowid := ctx.PostForm("procflowid")
		appid := ctx.PostForm("appid")
		approve := ctx.PostForm("approve")
		reason := ctx.PostForm("reason")

		datamap := make(map[string]any)
		datamap["approve"] = approve
		datamap["reason"] = reason

		//move to next task
		err = api.MoveToNextNode(user, approve, procdefid, procinstid, procflowid, taskid, appid, nil, datamap)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println("success")
		ctx.Request.URL.Path = "/deal"
		ctx.Request.Method = "GET"
		router.HandleContext(ctx)
	})

	router.POST("/movetotask", func(ctx *gin.Context) {
		//save form data
		formid := ctx.PostForm("formid")
		taskid := ctx.PostForm("taskid")
		procdefid := ctx.PostForm("procdefid")
		procinstid := ctx.PostForm("procinstid")
		procflowid := ctx.PostForm("procflowid")
		datamap := make(map[string]any)

		ctx.Request.ParseForm()
		for k, v := range ctx.Request.PostForm {
			if k == "formid" || k == "taskid" || k == "procdefid" || k == "procinstid" || k == "procflowid" {
				continue
			}
			datamap[k] = v[0]
		}

		err := api.SaveFormData(user, formid, procinstid, procflowid, datamap)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Println(formid)

		//move to next task
		err = api.MoveToNextNode(user, "", procdefid, procinstid, procflowid, taskid, "", nil, datamap)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println("success")
		ctx.Request.URL.Path = "/deal"
		ctx.Request.Method = "GET"
		router.HandleContext(ctx)
	})
	router.POST("/startprocess", func(ctx *gin.Context) {
		pid := "1"
		if pid == "" {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": "procdefid is required!"})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
			return
		}

		instId, flowId, taskId, formId, err := api.StartProcessInstance(pid, user)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
			return
		}
		// get form by form id
		form, err := api.FindFormById(formId)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
			return
		}
		// rawdata
		rawdata := `<form id="` + formId + `" action="/movetotask" method="post">`
		rawdata += form.Form
		rawdata += `<input type="submit" value="submit">`
		rawdata += `<input type="hidden" name="formid" value="` + formId + `">`

		rawdata += `<input type="hidden" name="taskid" value="` + taskId + `">`
		rawdata += `<input type="hidden" name="procdefid" value="` + pid + `">`
		rawdata += `<input type="hidden" name="procinstid" value="` + instId + `">`
		rawdata += `<input type="hidden" name="procflowid" value="` + flowId + `">`
		rawdata += `</form>`

		ctx.HTML(http.StatusOK, "vacation.html", gin.H{"rawdata": rawdata})
	})

	router.Run(":" + port)
}

func GetUser(utype string) *model.User {
	user := &model.User{}
	if utype == "0" {
		user.Id = "1"
		user.Name = "mike"
		user.Email = "test@123.com"
		user.Phone = "123456789"
		user.DeptId = "1"
		user.DeptName = "test department"
		user.ManagerId = "2"
		user.ManagerName = "manager kelly"
	} else if utype == "1" {
		user.Id = "2"
		user.Name = "manager kelly"
		user.Email = "approver@123.com"
		user.Phone = "66666666"
		user.DeptId = "1"
		user.DeptName = "test department"
		user.ManagerId = "3111"
		user.ManagerName = "ceo"
	} else if utype == "2" {
		user.Id = "3111"
		user.Name = "ceo tony"
		user.Email = "tony@123.com"
		user.Phone = "777777"
		user.DeptId = "2"
		user.DeptName = "ceo department"
		user.ManagerId = "3111"
		user.ManagerName = "ceo tony"
	}
	return user
}

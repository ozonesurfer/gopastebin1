package controllers

import (
	"github.com/astaxie/beego"
	//	"html/template"
	"log"
	"model"
	"time"
)

type MainController struct {
	beego.Controller
}
type CreateController MainController
type ShowController MainController

func (this *MainController) Get() {
	/*	this.Data["Website"] = "beego.me"
		this.Data["Email"] = "astaxie@gmail.com"
	*/
	this.Data["Pastes"] = model.GetAll()
	this.Data["Languages"] = model.Languages
	this.TplNames = "index.tpl"
}

func (this *CreateController) Post() {
	paste := model.Paste{
		Title:   this.GetString("title"),
		Content: this.GetString("content"),
		// LanguageId: this.GetInt("language"),
		CreatedOn: time.Now(),
		Language:  this.GetString("language"),
	}
	log.Println("paste is", paste)

	paste.Add()

	/*	content, err := paste.HighlightKeywords()
		if err != nil {
			log.Fatalln("Display error:", err)
		}
	*/
	log.Println("Language =", paste.Language)
	this.Data["Paste"] = paste
	//	this.Data["Content"] = paste.Content
	//	this.Data["Language"] = model.Languages[paste.LanguageId].Name
	//	this.Data["Language"] = paste.Language
	this.Data["Prism"] = model.Languages[paste.Language]
	this.TplNames = "create.tpl"
}

func (this *ShowController) Get() {
	rawId := this.Ctx.Input.Param(":id")
	id := model.ToObjectId(rawId)
	paste := model.GetPaste(id)

	/*	content, err := paste.HighlightKeywords()
		if err != nil {
			log.Fatalln("Display error:", err)
		} */

	this.Data["Paste"] = paste
	//	this.Data["Content"] = paste.Content
	//	this.Data["Language"] = model.Languages[paste.LanguageId].Name
	//	this.Data["Language"] = paste.Language
	this.Data["Prism"] = model.Languages[paste.Language]

	this.TplNames = "create.tpl"
}

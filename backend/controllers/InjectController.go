package controllers

import (
	"encoding/xml"
	"fmt"

	"github.com/kataras/iris"
)

type InjectController struct {
	Ctx iris.Context
}

type AssetsFrom struct {
	Id              string `form:"id"`
	AssetsId        string `form:"assets_id"`
	AssetsVideoType int8   `form:"assets_video_type"`
	AssetsVideoInfo string `form:"asstes_video_info"`
}

type AssetsXml struct {
	XMLName   xml.Name `xml:"video"`           //元素名称
	Name      string   `xml:"name,attr"`       //，attr属性。
	CpId      int      `xml:"cp_id,attr"`      //，attr属性。
	SreachKey string   `xml:"sreach_key,attr"` //内部元素名称，值是它的主体。
}

func NewInjectController() *InjectController {
	return &InjectController{}
}

//媒资注入
func (this *InjectController) AssetsInject(ctx iris.Context) {

	asset := AssetsFrom{}
	err := ctx.ReadForm(&asset)
	if err != nil {
		// ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	inputReader := []byte(asset.AssetsVideoInfo)
	var AssetsXml AssetsXml
	xerr := xml.Unmarshal(inputReader, &AssetsXml)
	if xerr != nil {
		fmt.Println(xerr)
	} else {
		fmt.Printf("%v\n", AssetsXml)
		fmt.Printf("name=%s\n", AssetsXml.Name)
	}
	ctx.JSON(AssetsXml)

}

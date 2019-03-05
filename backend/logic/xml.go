package logic

import (
	"encoding/xml"
	"fmt"
)

type Result struct {
	AssetsXml []AssetsXml
}

type AssetsXml struct {
	XMLName   xml.Name `xml:"video"`           //元素名称
	Name      string   `xml:"name,attr"`       //，attr属性。
	CpId      string   `xml:"cp_id,attr"`      //，attr属性。
	SreachKey string   `xml:"sreach_key,attr"` //内部元素名称，值是它的主体。
	Number    int      `xml:"no,attr"`
}

func main() {
	data := `<video
	assets_id="来源媒资id（建议弃用）"
	name = "来源媒资id"
	cp_id="mgtv"
	cp_name="cp名称"
	no=""
	assets_category="电影"
	sreach_key="定价id"
	video_type="1为直播  0为点播"
	
></video>`
	inputReader := []byte(data)
	var AssetsXml AssetsXml
	err := xml.Unmarshal(inputReader, &AssetsXml)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", AssetsXml)
		fmt.Printf("name=%s\n", AssetsXml.Name)
	}

}

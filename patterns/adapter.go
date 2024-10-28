package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type XMLData struct {
	XMLName xml.Name `xml:"data"`
	Name    string   `xml:"name"`
	Value   string   `xml:"value"`
}

type Adapter struct {
	XMLData XMLData
}

func (a *Adapter) ToJSON() ([]byte, error) {
	data := Data{
		Name:  a.XMLData.Name,
		Value: a.XMLData.Value,
	}
	return json.Marshal(data)
}
func main() {
	xmlData := XMLData{Name: "Benny", Value: "Operator"}

	adapter := Adapter{XMLData: xmlData}

	jsonData, err := adapter.ToJSON()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(jsonData)

}

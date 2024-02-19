package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type KPU struct {
	TS    string              `json:"ts"`
	PSU   string              `json:"psu"`
	Mode  string              `json:"mode"`
	Chart ChartData           `json:"chart"`
	Table map[string]TableRow `json:"table"`
	//Progress ProgressData `json:"progres"`
}

type ChartData struct {
	Table1 int     `json:"100025"`
	Table2 int     `json:"100026"`
	Table3 int     `json:"100027"`
	Persen float32 `json:"persen"`
}

//
//type ProgressData string

type TableRow struct {
	Table1         int     `json:"100025"`
	Table2         int     `json:"100026"`
	Table3         int     `json:"100027"`
	PSU            string  `json:"psu"`
	Persen         float32 `json:"persen"`
	StatusProgress bool    `json:"status_progress"`
}

func main() {
	var kpu = KPU{}

	var url = "https://sirekap-obj-data.kpu.go.id/pemilu/hhcw/ppwp.json"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&kpu)

	fmt.Println(kpu)

	renderJson := func(w http.ResponseWriter, r *http.Request) {
		tmplt := template.Must(template.ParseFiles("index.html"))
		tmplt.Execute(w, struct {
			KPU KPU
		}{
			KPU: kpu,
		})
	}

	http.HandleFunc("/", renderJson)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

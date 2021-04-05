package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	Cuto "github.com/diiyw/cuto"
	"log"
	"net/http"
	"strings"
)

func main() {
	browser, err := Cuto.NewBrowser(
		Cuto.Debug())
	if err != nil {
		log.Println(err)
	}
	//defer browser.Close()
	// 打开百度首页
	tab, err := browser.Open("https://www.bet365.com/#/HO/")
	if err != nil {
		log.Println(err)
	}
	if err := tab.Wait(); err != nil {
		log.Println(err)
	}
	http.HandleFunc("/getHtml", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello world!___>")
		html, _ := tab.Js("document.body.outerHTML", 1000)
		s := fmt.Sprint(html.Value)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		if err != nil {
			log.Fatal(err)
		}
		doc.Find(".him-Fixture_Container").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Html())
		})
		writer.Write([]byte(s))
		//fmt.Fprint(writer,html.Value)
	})
	http.HandleFunc("/getText", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello world!___>")
		html, _ := tab.Js("document.body.outerText", 1000)
		fmt.Println(html.Value)
		s, _ := json.Marshal(html)
		writer.Write([]byte(s))
	})
	http.ListenAndServe(":8080", nil)
}

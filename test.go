package main

// Now have to change name of struct to proper to they can be exported.
//
// after that work on goimports

import (
	"fmt"
	"io"
	"log"

	"regexp"
	"strings"

	"github.com/ChimeraCoder/gojson"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const (
	idNames   = `//*[@id="hub-sidebar-content"]/div[3]/div/ul/li`
	names     = `//*[@id="page-{{id}}"]/div[1]/div[1]/header/h2`
	throttlez = `//*[@id="page-{{id}}"]/div[1]/div[1]/header/div/p`             //strong/a`
	post      = `//*[@id="page-{{id}}"]/form/div[2]/div[1]/div/div[1]/pre/span` // replace {{id}} with name
	returnz   = `//*[@id="page-{{id}}"]/form/div[2]/div[2]/div/div[2]/div/pre/span`
	filez     = `//*[@id="page-{{id}}"]/form/div[1]/div/div/span[2]/span[1]`
)

var (
	rgID = regexp.MustCompile(`[a-z]+`)
)

func getEndPointsTest() {
	// func main() {
	doc, err := htmlquery.LoadURL("https://dev.skuvault.com/reference")
	if err != nil {
		panic(err)
	}
	// Find all news item.
	for _, n := range htmlquery.Find(doc, idNames) {
		a := htmlquery.FindOne(n, "//a")
		name := getName(a, doc)
		fmt.Println("working on", name)
		file := getFile(a, doc)
		propCh <- strings.Title(name)
		println("propStr")
		nameCh <- `/` + name
		println("nameStr")
		throttle <- getThrottle(a, doc)[0]
		println("throttleStr")
		info <- getThrottle(a, doc)[1]
		println("infoStr")
		files <- file
		println("filesStr")
		apiStructs <- getPost(file, name, a, doc)
	}

}

func getName(n, doc *html.Node) string {
	xPath := getID(n, names)
	area := htmlquery.FindOne(doc, xPath)

	return strings.Replace(htmlquery.InnerText(area), `/`, ``, -1)
}

func getFile(n *html.Node, doc *html.Node) string {
	xPath := getID(n, filez)
	area := htmlquery.FindOne(doc, xPath)
	if area == nil {
		return "skuvault"
	}
	full := htmlquery.InnerText(area)
	splits := strings.Split(full, "/")
	if splits[len(splits)-1] == "api" {
		return "skuvault"
	}
	return splits[len(splits)-1]
}

func getID(n *html.Node, path string) string {
	id := rgID.FindString(htmlquery.SelectAttr(n, "href"))
	return strings.Replace(path, `{{id}}`, id, -1)
}

func getThrottle(n *html.Node, doc *html.Node) []string {
	xPath := getID(n, throttlez)
	// println(xPath)
	area := htmlquery.FindOne(doc, xPath)
	if area == nil {
		fmt.Println("No Throttle")
		return []string{"No Throttle", "No Info"}
	}

	firstChild := htmlquery.InnerText(area.FirstChild)
	secondChild := htmlquery.InnerText(area.LastChild)

	return []string{firstChild, secondChild}
}

func commentStructs(data string) string {
	typeStruct := regexp.MustCompile(`type ([a-zA-Z]+) struct {`)
	myStrucst := typeStruct.FindAllStringSubmatch(data, -1)
	for _, st := range myStrucst {
		repl := `// ` + st[1] + ` is a automatically generated struct from json provided by skuvault's api docs.` + "\n" + st[0]
		data = strings.Replace(data, st[0], repl, -1)
	}
	return data
}

func getPost(file, name string, n *html.Node, doc *html.Node) []apiStructsData {
	xPath := getID(n, post)
	areaP := htmlquery.FindOne(doc, xPath)
	posts := htmlquery.InnerText(areaP)
	pwj := washJSON(posts, name)
	poStuct, err := gojson.Generate(pwj, gojson.ParseJson, strings.Title(name), file, []string{"json"}, false)
	if err != nil {
		fmt.Println(pwj)
		log.Panicln(err)
	}

	xPath = getID(n, returnz)
	areaR := htmlquery.FindOne(doc, xPath)
	resps := htmlquery.InnerText(areaR)
	rwj := washJSON(resps, name)
	reStuct, err := gojson.Generate(rwj, gojson.ParseJson, strings.Title(name)+"Response", file, []string{"json"}, false)
	if err != nil {
		fmt.Println(resps)
		log.Panicln(err)
	}
	postStruct := apiStructsData{
		Spot:   1,
		Struct: commentStructs(string(poStuct)),
		File:   file,
		Name:   name,
	}
	repStruct := apiStructsData{
		Spot:   2,
		Struct: commentStructs(string(reStuct)),
		File:   file,
		Name:   name,
	}

	return []apiStructsData{postStruct, repStruct}
}

func washJSON(jsons, name string) io.Reader {
	ok := regexp.MustCompile(`"status2",`)
	ok2 := regexp.MustCompile(`}\s*]`)
	ok3 := regexp.MustCompile(`,(\n\s*})`)
	f1 := ok.ReplaceAllString(jsons, `"status2"`)
	f2 := ok3.ReplaceAllString(f1, `$1`)
	if name == "getTransactions" {
		f2 = ok2.ReplaceAllString(f1, `}}]`)
	}

	return strings.NewReader(f2)

}

package main

import (
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

// getEndPoints scrapes skuvault api docs to make package.
func getEndPoints() {
	doc, err := htmlquery.LoadURL("https://dev.skuvault.com/reference")
	if err != nil {
		panic(err)
	}
	// Find all news item.
	for _, n := range htmlquery.Find(doc, idNames) {
		a := htmlquery.FindOne(n, "//a")
		name := getName(a, doc)
		if name == "getTokens" {
			continue
		}
		file := getFile(a, doc)
		propCh <- strings.Title(name)
		nameCh <- `/` + name
		throttle <- getThrottle(a, doc)[0]
		info <- getThrottle(a, doc)[1]
		files <- file
		apiStructs <- getPost(file, name, a, doc)
	}

}

// getName scrapes name.
func getName(n, doc *html.Node) string {
	xPath := getID(n, names)
	area := htmlquery.FindOne(doc, xPath)

	return strings.Replace(htmlquery.InnerText(area), `/`, ``, -1)
}

// getFile scrapes the file name or category "inventory, shipping, etc".
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

// getID gets name then makes an xpath with the name as an id
func getID(n *html.Node, path string) string {
	id := rgID.FindString(htmlquery.SelectAttr(n, "href"))
	return strings.Replace(path, `{{id}}`, id, -1)
}

// getThrottle gets the throttle data from each endpoint some don't have any.
func getThrottle(n *html.Node, doc *html.Node) []string {
	xPath := getID(n, throttlez)
	area := htmlquery.FindOne(doc, xPath)
	if area == nil {
		return []string{"No Throttle", "No Info"}
	}

	firstChild := htmlquery.InnerText(area.FirstChild)
	secondChild := htmlquery.InnerText(area.LastChild)

	return []string{firstChild, secondChild}
}

// commentStructs adds comments to the structures that are made by the system.
func commentStructs(data string) string {
	typeStruct := regexp.MustCompile(`type ([a-zA-Z]+) struct {`)
	toks := regexp.MustCompile(`[TU][a-z]+[T][a-z]+\s+string\s+\x60json:"[TU][a-z]+[T][a-z]+"\x60\n`)
	data = toks.ReplaceAllString(data, "")
	myStrucs := typeStruct.FindAllStringSubmatch(data, -1)
	for _, st := range myStrucs {
		repl := `// ` + st[1] + ` is a automatically generated struct from json provided by sku vault's api docs.` + "\n" + st[0]
		data = strings.Replace(data, st[0], repl, -1)
	}
	return data
}

// getPost gets the structure of the post and response of the api from the docs.
func getPost(file, name string, n *html.Node, doc *html.Node) []apiStructsData {
	xPath := getID(n, post)
	areaP := htmlquery.FindOne(doc, xPath)
	posts := htmlquery.InnerText(areaP)
	pwj := washJSON(posts, name)
	poStruct, err := gojson.Generate(pwj, gojson.ParseJson, strings.Title(name), file, []string{"json"}, false, true)
	if err != nil {
		log.Panicln(err)
	}

	xPath = getID(n, returnz)
	areaR := htmlquery.FindOne(doc, xPath)
	resps := htmlquery.InnerText(areaR)
	rwj := washJSON(resps, name)
	reStruct, err := gojson.Generate(rwj, gojson.ParseJson, strings.Title(name)+"Response", file, []string{"json"}, false, true)
	if err != nil {
		log.Panicln(err)
	}
	postStruct := apiStructsData{
		Spot:   1,
		Struct: commentStructs(string(poStruct)),
		File:   file,
		Name:   name,
	}
	repStruct := apiStructsData{
		Spot:   2,
		Struct: commentStructs(string(reStruct)),
		File:   file,
		Name:   name,
	}

	return []apiStructsData{postStruct, repStruct}
}

// washJSON makes the json taken from the site correct.
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

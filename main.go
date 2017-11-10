package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ChimeraCoder/gojson"

	"golang.org/x/net/html"
)

var (
	nameCh     chan string
	propCh     chan string
	info       chan string
	throttle   chan string
	apiStructs chan []apiStructsData
	files      chan string
	allDone    chan bool

	infoDone bool

	rgx, _      = regexp.Compile(`\s?[tT]hrottling[.:]?\s?$`)
	th2, _      = regexp.Compile(`\s?[tT]hrottling[.:]?\s?`)
	nl, _       = regexp.Compile(`\n`)
	fileName, _ = regexp.Compile(`(inventory|products|sales|purchaseorders|integration)`)
)

type templateTags struct {
	Proper   string
	Throttle string
	Name     string
	Info     string
	File     string
	File0    string
}
type apiStructsData struct {
	Name   string
	File   string
	Spot   int
	Struct string
}

func init() {
	os.Mkdir(`skuvault`, os.ModePerm)
	os.Mkdir(`skuvault/inventory`, os.ModePerm)
	os.Mkdir(`skuvault/products`, os.ModePerm)
	os.Mkdir(`skuvault/sales`, os.ModePerm)
	os.Mkdir(`skuvault/purchaseorders`, os.ModePerm)
	os.Mkdir(`skuvault/integration`, os.ModePerm)
}

func main() {
	nameCh = make(chan string)
	propCh = make(chan string)
	info = make(chan string)
	throttle = make(chan string)
	apiStructs = make(chan []apiStructsData)
	files = make(chan string)
	allDone = make(chan bool)

	res, err := http.Get("https://dev.skuvault.com/reference")
	if err != nil {
		log.Panicln(err)
	}

	doc := html.NewTokenizer(res.Body)
	go makeFuncs()
	go makeStructs()

	getEndPoints(doc)
	<-allDone
}

func makeStructs() {
	for {
		apiStr := <-apiStructs
		var f *os.File
		var err error
		structStr := ""
		if apiStr[0].File != "skuvault" {
			f, err = os.Create("skuvault/" + apiStr[0].File + "/" + apiStr[0].Name + ".go")

			if err != nil {
				log.Panicln(err)
			}
		} else {
			f, err = os.Create("skuvault/" + apiStr[0].Name + ".go")
			if err != nil {
				log.Panicln(err)
			}
		}
		for _, stru := range apiStr {
			pakName, err := regexp.Compile(`package (inventory|products|sales|purchaseorders|integration|skuvault)`)
			if err != nil {
				log.Panicln(err)
			}

			cleand := pakName.ReplaceAllString(stru.Struct, "")
			structStr += cleand

		}
		f.WriteString(`package ` + apiStr[0].File)
		f.WriteString(structStr)
		f.Close()
	}

}

func makeFuncs() {
	inF, err := os.Create("skuvault/inventory.go")
	if err != nil {
		log.Panicln(err)
	}
	prF, err := os.Create("skuvault/products.go")
	if err != nil {
		log.Panicln(err)
	}
	saF, err := os.Create("skuvault/sales.go")
	if err != nil {
		log.Panicln(err)
	}
	poF, err := os.Create("skuvault/purchaseorders.go")
	if err != nil {
		log.Panicln(err)
	}
	noF, err := os.Create("skuvault/nocategory.go")
	if err != nil {
		log.Panicln(err)
	}
	funcTemp := `
	// {{.Proper}} creates http request for this SKU vault endpoint
	// {{.Throttle}} Throttle
	// {{.Info}}
	func (lc *{{.File0}}LoginCredentials) {{.Proper}}(pld *{{.File}}.{{.Proper}}) *{{.File}}.{{.Proper}}Response {
		credPld := &post{{.Proper}} {
			{{.Proper}}:       pld,
			{{.File0}}LoginCredentials: lc,
		}

		response := &{{.File}}.{{.Proper}}Response{}
		{{.File}}Call(credPld, response, {{.Name}})
		return response
	}`

	inTemp, err := template.New("inventory").Parse(funcTemp)
	if err != nil {
		log.Panicln(err)
	}
	proTemp, err := template.New("products").Parse(funcTemp)
	if err != nil {
		log.Panicln(err)
	}
	salTemp, err := template.New("sales").Parse(funcTemp)
	if err != nil {
		log.Panicln(err)
	}
	poTemp, err := template.New("purchaseorders").Parse(funcTemp)
	if err != nil {
		log.Panicln(err)
	}
	noTemp, err := template.New("noTemp").Parse(funcTemp)
	if err != nil {
		log.Panicln(err)
	}

	inF.WriteString("package skuvault")
	prF.WriteString("package skuvault")
	saF.WriteString("package skuvault")
	poF.WriteString("package skuvault")
	noF.WriteString("package skuvault")

	for {
		println("start")
		propStr := <-propCh
		println("propStr")
		nameStr := <-nameCh
		println("nameStr")
		throttleStr := <-throttle
		println("throttleStr")
		infoStr := <-info
		println("infoStr:::", infoStr)
		filesStr := <-files
		println("filesStr")

		file0 := ""
		if filesStr != "" {
			file0 = strings.ToUpper(string(filesStr[0]))
		}

		tags := templateTags{
			Proper:   propStr,
			Throttle: throttleStr,
			Name:     nameStr,
			Info:     infoStr,
			File:     filesStr,
			File0:    file0,
		}
		switch filesStr {
		case "inventory":
			err = inTemp.Execute(inF, tags)
			if err != nil {
				log.Panicln(err)
			}
		case "products":
			err = proTemp.Execute(prF, tags)
			if err != nil {
				log.Panicln(err)
			}
		case "sales":
			err = salTemp.Execute(saF, tags)
			if err != nil {
				log.Panicln(err)
			}
		case "purchaseorders":
			tags.File0 = "PO"
			err = poTemp.Execute(poF, tags)
			if err != nil {
				log.Panicln(err)
			}
		default:
			tags.File0 = ""
			err = noTemp.Execute(noF, tags)
			if err != nil {
				log.Panicln(err)
			}

		}
	}
}

func structRename(data string) {
	liens := strings.Split(data, "\n")
	for i, line := range liens {
		sub := regexp.MustCompile(`^\s[^\s]+\s+[^\s]+_[^\s]+` + strconv.Itoa(i+1) + `$`)
		if !sub.MatchString(line) {
			continue
		}
		splits := sub.

	}

}

func getEndPoints(doc *html.Tokenizer) {
	name := ""
	file := ""
	bothStructs := []apiStructsData{}
	apiCount := 0
	for {
		toks := doc.Next()

		switch toks {
		case html.ErrorToken:
			println("done")
			allDone <- true
			return
		case html.StartTagToken:
			tok := doc.Token()
			if tok.Data == "h2" {
				toks = doc.Next()
				full := doc.Token().Data
				name = strings.Replace(full, "/", "", -1)
				proper := strings.Title(name)

				if proper != "Developing For SkuVault" {
					propCh <- proper
					fmt.Println(proper)
					println("AFTER")
					nameCh <- name
					println("AFTER?")
				}

				continue
			}

			if tok.Data == "div" {
				for _, att := range tok.Attr {
					if att.Val == "excerpt" {
						count := 0
					out:
						for {
							toks = doc.Next()
							tok = doc.Token()
							switch toks {
							case html.TextToken:
								if count == 2 {
									continue
								}
								if !rgx.MatchString(tok.Data) {
									count++
									first := th2.ReplaceAllString(tok.Data, "")
									println(count)
									if count == 1 {
										infoDone = false
										println("throttle")
										throttle <- first
										println("throttle?")
										continue
									}
									println("Info")
									info <- first
									infoDone = true
									println("info?")

								}
							case html.EndTagToken:
								if tok.Data == "p" {
									break out
								}
							}
						}
					}
				}
			}

			if tok.Data == "span" {

				jsons := ""
				for _, att := range tok.Attr {
					if att.Val == "cm-s-tomorrow-night" {
					out2:
						for {
							toks = doc.Next()
							tok = doc.Token()
							switch toks {
							case html.TextToken:
								if !rgx.MatchString(tok.Data) {
									jsons += tok.Data
								}
							case html.EndTagToken:
								if tok.Data == "pre" {
									ok, _ := regexp.Compile(`"status2",`)
									ok2, _ := regexp.Compile(`}\s*]`)
									f1 := ok.ReplaceAllString(jsons, `"status2"`)
									f2 := f1
									if name == "getTransactions" {
										f2 = ok2.ReplaceAllString(f1, `}}]`)
									}
									jRead := strings.NewReader(f2)
									change := false
									if file == "" {
										file = "skuvault"
										change = true
									}
									StruName := name
									if apiCount == 1 {
										StruName += "Response"
									}
									b, err := gojson.Generate(jRead, gojson.ParseJson, StruName, file, nil, true)
									if err != nil {
										println(f2)
										log.Panicln(err)
									}
									apiCount++
									println("API COUNT>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", apiCount)
									st := apiStructsData{
										Spot:   apiCount,
										Struct: string(b),
										File:   file,
										Name:   name,
									}
									bothStructs = append(bothStructs, st)
									if apiCount == 2 {
										apiStructs <- bothStructs
										bothStructs = []apiStructsData{}
										apiCount = 0
									}
									if change {
										file = ""
									}

									println("^^^^^^^^^^")
									println("||||||||||")
									println("||||||||||")
									println("||||||||||")
									println("||||||||||")
									println("|________|")
									break out2
								}
							}
						}
					}

					if att.Val == "definition-url" {
						var done bool
					out3:
						for {
							toks = doc.Next()
							tok = doc.Token()
							switch {
							case toks == html.TextToken:
								url, _ := regexp.Compile(`app.skuvault.com/api/`)
								file = url.ReplaceAllString(tok.Data, "")
								if file == "app.skuvault.com/api" {
								urlSpan:
									for {
										toks = doc.Next()
										tok = doc.Token()
										switch {
										case toks == html.TextToken:
											if len(tok.Data) == 0 {
												continue urlSpan
											}
											otherURL := strings.Split(tok.Data, "/")
											file = otherURL[1]
											if !fileName.MatchString(file) {
												file = ""
											}
											break urlSpan
										}
									}
								}
								println(file)
								if !infoDone {
									info <- "NONE"
								}
								println("test2")
								files <- file
								println("test2?")
								done = true
							case done:
								break out3
							}
						}
					}

				}

			}
		}
	}

}

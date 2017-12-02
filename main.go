package main

/*
*
* Have to collect template data for SV main, and figue out goimports.
*
 */

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/ChimeraCoder/gojson"
	"github.com/OuttaLineNomad/skuvault-api-doc-scraper/templates"
	"golang.org/x/net/html"
)

const (
	folder = "../skuvault"
)

var (
	nameCh     chan string
	propCh     chan string
	info       chan string
	throttle   chan string
	apiStructs chan []apiStructsData
	files      chan string

	infoDone bool

	rgx, _      = regexp.Compile(`\s?[tT]hrottling[.:]?\s?$`)
	th2, _      = regexp.Compile(`\s?[tT]hrottling[.:]?\s?`)
	nl, _       = regexp.Compile(`\n`)
	fileName, _ = regexp.Compile(`(inventory|products|sales|purchaseorders|integration)`)
)

type creds struct {
	Name string
	Cred string
}

type svgTemplateTags struct {
	Company  string
	URL      string
	CtrBunch []creds
}

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
	os.Mkdir(folder, os.ModePerm)
	os.Mkdir(folder+"/inventory", os.ModePerm)
	os.Mkdir(folder+"/products", os.ModePerm)
	os.Mkdir(folder+"/sales", os.ModePerm)
	os.Mkdir(folder+"/purchaseorders", os.ModePerm)
	os.Mkdir(folder+"/integration", os.ModePerm)

}

func main() {
	nameCh = make(chan string)
	propCh = make(chan string)
	info = make(chan string)
	throttle = make(chan string)
	apiStructs = make(chan []apiStructsData)
	files = make(chan string)

	println("get api docs")
	res, err := http.Get("https://dev.skuvault.com/reference")
	if err != nil {
		log.Panicln(err)
	}

	doc := html.NewTokenizer(res.Body)
	go makeFuncs()
	go makeStructs()
	println("starting getEndPoints")
	getEndPoints(doc)
	println("go fmt...")
	exec.Command("go", "fmt", folder).Run()
	exec.Command("go", "fmt", folder+"/integration/").Run()
	exec.Command("go", "fmt", folder+"/inventory/").Run()
	exec.Command("go", "fmt", folder+"/products/").Run()
	exec.Command("go", "fmt", folder+"/purchaseorders/").Run()
	exec.Command("go", "fmt", folder+"/sales/").Run()
	println("goimports...")
	addImports()
}

func addImports() {
	cmd := exec.Command("goimports", folder+"/skuvault.go")
	src, err := os.Open(folder + "/skuvault.go")
	if err != nil {
		println(err.Error())
		return
	}
	b, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	src.Write(b)
	src.Close()
}

func makeSVControl(company, url string, crd []creds) {
	svg, err := os.Create(folder + "/skuvault.go")
	if err != nil {
		log.Panicln(err)
	}
	svgTemp, err := template.New("mainsvg").Parse(templates.SVMain)
	if err != nil {
		log.Panicln(err)
	}
	tags := svgTemplateTags{
		Company:  company,
		URL:      url,
		CtrBunch: crd,
	}
	err = svgTemp.Execute(svg, tags)
	if err != nil {
		log.Panicln(err)
	}
}

func makeStructs() {
	for {
		apiStr := <-apiStructs
		var f *os.File
		var err error
		structStr := ""
		if apiStr[0].File != "skuvault" {
			f, err = os.Create(folder + "/" + apiStr[0].File + "/" + apiStr[0].Name + ".go")

			if err != nil {
				log.Panicln(err)
			}
		} else {
			f, err = os.Create(folder + "/" + apiStr[0].Name + ".go")
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
		fileStr := structRename(structStr)
		f.WriteString(fileStr)
		f.Close()
	}

}

func makeFuncs() {
	inF, err := os.Create(folder + "/inventory.go")
	if err != nil {
		log.Panicln(err)
	}
	prF, err := os.Create(folder + "/products.go")
	if err != nil {
		log.Panicln(err)
	}
	saF, err := os.Create(folder + "/sales.go")
	if err != nil {
		log.Panicln(err)
	}
	poF, err := os.Create(folder + "/purchaseorders.go")
	if err != nil {
		log.Panicln(err)
	}
	intF, err := os.Create(folder + "/integration.go")
	if err != nil {
		log.Panicln(err)
	}
	noF, err := os.Create(folder + "/nocategory.go")
	if err != nil {
		log.Panicln(err)
	}
	funcTemp := templates.Funcs

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
	intTemp, err := template.New("integration").Parse(funcTemp)
	if err != nil {
		log.Panicln(err)
	}
	noTemp, err := template.New("noTemp").Parse(strings.Replace(funcTemp, "{{.File}}.", "", -1))
	if err != nil {
		log.Panicln(err)
	}

	inF.WriteString("package skuvault")
	prF.WriteString("package skuvault")
	saF.WriteString("package skuvault")
	poF.WriteString("package skuvault")
	noF.WriteString("package skuvault")
	intF.WriteString("package skuvault")

	for {
		propStr := <-propCh
		println("propStr")
		nameStr := <-nameCh
		println("nameStr")
		throttleStr := <-throttle
		println("throttleStr")
		infoStr := <-info
		println("infoStr")
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
		case "integration":
			tags.File0 = "IN"
			err = intTemp.Execute(intF, tags)
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

func structRename(data string) string {
	lines := strings.Split(data, "\n")
	newData := data
	for _, line := range lines {
		sub := regexp.MustCompile(`^\s+(?P<key>[^\s]+)\s+[\[\]]*(?P<val>[^\s]+_sub[1-9][0-9]*)`)
		if !sub.MatchString(line) {
			continue
		}
		key := sub.FindStringSubmatch(line)
		newData = strings.Replace(newData, key[2], key[1], -1)
	}
	return newData
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
			return
		case html.StartTagToken:
			tok := doc.Token()
			if tok.Data == "h2" {
				toks = doc.Next()
				full := doc.Token().Data
				name = strings.Replace(full, "/", "", -1)
				proper := strings.Title(name)

				if proper != "Developing For SkuVault" {
					println("Nameing")
					propCh <- proper
					nameCh <- name
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
									if count == 1 {
										infoDone = false
										println("Throttle")
										throttle <- first
										continue
									}
									info <- first
									infoDone = true

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
										log.Panicln(err)
									}
									apiCount++
									st := apiStructsData{
										Spot:   apiCount,
										Struct: string(b),
										File:   file,
										Name:   name,
									}
									bothStructs = append(bothStructs, st)
									if apiCount == 2 {
										println("adding structs")
										apiStructs <- bothStructs
										bothStructs = []apiStructsData{}
										apiCount = 0
									}
									if change {
										file = ""
									}

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
								if !infoDone {
									info <- "NONE"
								}
								println("get files")
								files <- file
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

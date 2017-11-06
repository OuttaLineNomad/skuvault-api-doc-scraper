package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/ChimeraCoder/gojson"

	"golang.org/x/net/html"
)

var (
	nameCh   chan string
	propCh   chan string
	info     chan string
	throttle chan string
	structs  chan string
	files    chan string
	allDone  chan bool

	wg     sync.WaitGroup
	rgx, _ = regexp.Compile(`\s?[tT]hrottling[.:]?\s?$`)
	th2, _ = regexp.Compile(`\s?[tT]hrottling[.:]?\s?`)
	nl, _  = regexp.Compile(`\n`)
)

type templateTags struct {
	Proper   string
	Throttle string
	Name     string
	Info     string
	File     string
	File0    string
}

func main() {
	nameCh = make(chan string)
	propCh = make(chan string)
	info = make(chan string)
	throttle = make(chan string)
	// structs = make(chan string, 2)
	files = make(chan string)
	allDone = make(chan bool)

	res, err := http.Get("https://dev.skuvault.com/reference")
	if err != nil {
		log.Panicln(err)
	}

	doc := html.NewTokenizer(res.Body)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-allDone:
				wg.Done()
				return
			default:
				println("start")
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
				// structsStr := <-structs

				tags := templateTags{
					Proper:   propStr,
					Throttle: throttleStr,
					Name:     nameStr,
					Info:     infoStr,
					File:     filesStr,
					File0:    string(filesStr[0]),
				}

				funcTemp := `
		// {{.Proper}} creates http request for this SKU vault endpoint
		// {{.Throttle}} Throttle
		// {{.Info}}
		func (lc *PLoginCredentials) {{.Proper}}(pld *{{.File}}.{{.Proper}}) *{{.File}}.{{.Proper}}Response {
			credPld := &post{{.Proper}} {
				{{.Proper}}:       pld,
				{{.File0}}LoginCredentials: lc,
			}

			response := &{{.File}}.{{.Proper}}Response{}
			{{.File}}Call(credPld, response, {{.Name}})
			return response
		}`

				f, err := os.Create("test.go")
				if err != nil {
					log.Panicln(err)
				}
				tmp, err := template.New("test").Parse(funcTemp)
				if err != nil {
					log.Panicln(err)
				}
				err = tmp.Execute(f, tags)
				if err != nil {
					log.Panicln(err)
				}
			}
		}
	}()

	getEndPoints(doc)
	wg.Wait()
}

func getEndPoints(doc *html.Tokenizer) {
	name := ""
	file := ""
	for {
		toks := doc.Next()

		switch toks {
		case html.ErrorToken:
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
									fmt.Println(first)
									println(count)
									if count == 1 {
										println("throttle")
										throttle <- first
										println("throttle?")
										continue
									}
									println(first)
									info <- first
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
									b, err := gojson.Generate(jRead, gojson.ParseJson, name, file, nil, true)
									if err != nil {
										println(f2)
										log.Panicln(err)
									}
									fmt.Println(string(b))
									break out2
								}
							}
						}
					}

					if att.Val == "definition-url" {
						// println("test2")
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
									toks = doc.Next()
									tok = doc.Token()
									otherURL := strings.Split(tok.Data, "/")
									file = otherURL[0]
								}
								fmt.Println(file)

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

package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/ChimeraCoder/gojson"

	"golang.org/x/net/html"
)

var (
	dCh    chan *html.Tokenizer
	wg     sync.WaitGroup
	rgx, _ = regexp.Compile(`\s?[tT]hrottling[.:]?\s?$`)
	th2, _ = regexp.Compile(`\s?[tT]hrottling[.:]?\s?`)
	nl, _  = regexp.Compile(`\n`)
)

func main() {
	res, err := http.Get("https://dev.skuvault.com/reference")
	if err != nil {
		log.Panicln(err)
	}

	doc := html.NewTokenizer(res.Body)
	getEndPoints(doc)
	// f, _ := os.Open("testing.json")

	// b, _ := gojson.Generate(f, gojson.ParseJson, "Test", "test", nil, true)
	// fmt.Println(string(b))
}

func getEndPoints(doc *html.Tokenizer) {
	name := ""
	for {
		toks := doc.Next()

		switch toks {
		case html.ErrorToken:
			// wg.Done()
			return
		case html.StartTagToken:
			tok := doc.Token()
			if tok.Data == "h2" {
				toks = doc.Next()
				full := doc.Token().Data
				name = strings.Replace(full, "/", "", -1)
				fmt.Println(full)
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
									b, err := gojson.Generate(jRead, gojson.ParseJson, name, "test", nil, true)
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
				}
			}
		}
	}

}

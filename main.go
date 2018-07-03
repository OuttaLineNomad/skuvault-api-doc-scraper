package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"

	"github.com/OuttaLineNomad/skuvault-api-doc-scraper/templates"
)

const (
	folder = "REVIEW_skuvault"
)

var (
	nameCh     chan string
	propCh     chan string
	info       chan string
	throttle   chan string
	apiStructs chan []apiStructsData
	files      chan string
	allDoneSon chan bool
	wg         sync.WaitGroup

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
	throttle = make(chan string, 1)
	apiStructs = make(chan []apiStructsData)
	files = make(chan string, 1)
	allDoneSon = make(chan bool, 1)

	go makeFuncs()
	go makeStructs()
	getEndPoints()
	allDoneSon <- true
	wg.Wait()
	exec.Command("go", "fmt", folder).Run()
	exec.Command("go", "fmt", folder+"/integration/").Run()
	exec.Command("go", "fmt", folder+"/inventory/").Run()
	exec.Command("go", "fmt", folder+"/products/").Run()
	exec.Command("go", "fmt", folder+"/purchaseorders/").Run()
	exec.Command("go", "fmt", folder+"/sales/").Run()
}

// makeSVControl makes the main control file from a template.
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

// makeStructs creates type struct from json scraped by scraper.
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
		fileStr := structStr
		f.WriteString(fileStr)
		f.Close()
	}

}

// makeFuncs is on a go runtime createing functions from the data it is given.
func makeFuncs() {
	wg.Add(1)

	// create all the go files seperated by functionality.
	names := []string{}
	inF, err := os.Create(folder + "/inventory.go")
	if err != nil {
		log.Panicln(err)
	}
	names = append(names, inF.Name())

	prF, err := os.Create(folder + "/products.go")
	if err != nil {
		log.Panicln(err)
	}
	names = append(names, prF.Name())

	saF, err := os.Create(folder + "/sales.go")
	if err != nil {
		log.Panicln(err)
	}
	names = append(names, saF.Name())

	poF, err := os.Create(folder + "/purchaseorders.go")
	if err != nil {
		log.Panicln(err)
	}
	names = append(names, poF.Name())

	intF, err := os.Create(folder + "/integration.go")
	if err != nil {
		log.Panicln(err)
	}
	names = append(names, intF.Name())

	noF, err := os.Create(folder + "/nocategory.go")
	if err != nil {
		log.Panicln(err)
	}
	names = append(names, noF.Name())

	// Create templates for all files
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

	// make creds for skuvault file to control all funcs
	theCreds := []creds{}
	fileNames := []string{"inventory", "products", "sales", "purchaseorders", "integration", "skuvault"}
	for _, name := range fileNames {
		start := strings.ToUpper(string(name[0]))
		if name == "skuvault" {
			start = ""
		}
		if name == "purchaseorders" {
			start = "PO"
		}
		if name == "integration" {
			start = "IN"
		}
		c := creds{
			Cred: start + "LoginCredentials",
			Name: name,
		}
		theCreds = append(theCreds, c)
	}
	makeSVControl("skuvault", "https://app.skuvault.com/api/", theCreds)

	// add pageage name to all files.
	inF.WriteString("package skuvault")
	prF.WriteString("package skuvault")
	saF.WriteString("package skuvault")
	poF.WriteString("package skuvault")
	noF.WriteString("package skuvault")
	intF.WriteString("package skuvault")

	// Go ruetine wait till all functions are created then run goimports
	go goImportFuncs(names)

	// for loop to wait for data from scrapper to create functions.
	for {
		propStr := <-propCh
		nameStr := <-nameCh
		throttleStr := <-throttle
		infoStr := <-info
		filesStr := <-files

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

// structRename renames subs to the actual name.
// func structRename(data string) string {
// 	lines := strings.Split(data, "\n")
// 	newData := data
// 	for _, line := range lines {
// 		sub := regexp.MustCompile(`^\s+(?P<key>[^\s]+)\s+[\[\]]*(?P<val>[^\s]+_sub[1-9][0-9]*)`)
// 		if !sub.MatchString(line) {
// 			continue
// 		}
// 		key := sub.FindStringSubmatch(line)
// 		newData = strings.Replace(newData, key[2], key[1], -1)
// 	}
// 	return newData
// }

// goImportFuncs uses addImports to add imports to each func file
func goImportFuncs(names []string) {
	for {
		myNames := names
		select {
		case <-allDoneSon:
			for _, name := range myNames {
				addImports(name)
			}
			wg.Done()
			return
		}
	}
}

// addImports uses goimports to add the import lines on top of each go file
func addImports(name string) {
	cmd := exec.Command("goimports", name)
	b, err := cmd.Output()
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(name, b, 0644)
	if err != nil {
		log.Panic(err)
	}
}

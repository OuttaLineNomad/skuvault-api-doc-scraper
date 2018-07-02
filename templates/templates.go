package templates

const (
	// SVMain constant template for skuvault main package.
	SVMain = `
package skuvault

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	// url is a constant for all {{.Company}} calls
	url = "{{.URL}}"
)

// Ctr is the controls the flow of endpoints.
type Ctr struct {
	{{range .CtrBunch}}
	{{.Name}} *{{.Cred}}
	{{end}}
}

{{range .CtrBunch}}
// {{.Cred}} hold credentials to sign into SKU Vault API for inventory endpoints.
type {{.Cred}} struct {
	TenantToken string
	UserToken   string
}
{{end}}

// NewEnvCredSession takes tokens from systems environmental varables.
// TENANT_TOKEN and USER_TOKEN
func NewEnvCredSession() *Ctr {
	return NewSession(os.Getenv("SV_TENANT_TOKEN"), os.Getenv("SV_USER_TOKEN"))
}

// NewSession creates a new session sets credentails to make call
func NewSession(tTok, uTok string) *Ctr {
	return &Ctr{
		{{range .CtrBunch}}
		{{.Name}}:  &{{.Cred}}{
			TenantToken: tTok,
			UserToken:   uTok,
		},
		{{end}}
	}
}

// postGetTokens payload sent to Sku Vault.
type postGetTokens struct {
	*GetTokens
}

// GetTokensResponse is a automatically generated struct from json provided by skuvault's api docs.
type GetTokensResponse struct {
	TenantToken string 
	UserToken   string
}

// GetTokens is a automatically generated struct from json provided by skuvault's api docs.
type GetTokens struct {
	Email    string
	Password string
}

// GetTokensCall creates http request for this SKU vault endpoint.
// Very Light Throttle.
// Use this call to retrieve your API tokens from SkuVault using your login email and password..
func GetTokensCall(pld *GetTokens) *GetTokensResponse {
	credPld := &postGetTokens{
		GetTokens: pld,
	}

	response := &GetTokensResponse{}
	do(credPld, response, "skuvault/getTokens")
	return response
}

// do internal makes calls based on information passed in from other Do calls for each endpoint
func do(pld interface{}, response interface{}, endPoint string) {
	fullURL := url + endPoint
	bt, err := json.Marshal(pld)
	if err != nil {
		log.Fatal(err)
	}
	payload := bytes.NewReader(bt)
	req, err := http.NewRequest(http.MethodPost, fullURL, payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, response)
	if err != nil {
		log.Fatal(err)
	}

}

// TimeString converts time to proper formated string for Sku Vault
func TimeString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.9999999Z")
}
`
	// Funcs func template for all SKU Vault api functions.
	Funcs = `
// post{{.Proper}} payload sent to Sku Vault.
type post{{.Proper}} struct {
	*{{.File}}.{{.Proper}}
	*{{.File0}}LoginCredentials
}

// {{.Proper}} creates http request for this SKU vault endpoint.
// {{.Throttle}} Throttle.
// {{.Info}}.
func (lc *{{.File0}}LoginCredentials) {{.Proper}}(pld *{{.File}}.{{.Proper}}) *{{.File}}.{{.Proper}}Response {
	credPld := &post{{.Proper}} {
		{{.Proper}}:       pld,
		{{.File0}}LoginCredentials: lc,
	}

	response := &{{.File}}.{{.Proper}}Response{}
	do(credPld, response, "{{.File}}{{.Name}}")
	return response
}`

	readme = `
# SkuVault 

## Overview

This library supports most if not all of the \60xapp.skuvault.com/api\60x REST calls in go. It was mostly been created automatically by a web scrapper scrapping SkuVault's api [reference](https://dev.skuvault.com/reference) then edited. Please refer to SkuVaults reference site for more info.

## Install

\60x\60x\60x
go get github.com/OuttaLineNomad/skuvault 
\60x\60x\60x

## Examples

### Using LeakyBucket
LeakyBucket sets size 4 bucket. Meaning only 10 calls can be made in a burst. After that the bucket is regulated at the rate of of 4 calls a minute. GetRate does the math for you.
\60x\60x\60xgo
func main() {
	bkt := throttle.LeakyBucket(10, throttle.GetRate(4, time.Minute))
	for i := 0; i < 100; i++ {
		bkt.Take()
		// after this you call you make your call.
		resp, err := http.Get("http://httpbin.org/get")
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(resp.StatusCode)
	}
}
\60x\60x\60x


## Author

* **Bryce Mullen** - *Project Manager*  @Wedgenix connect [here](https://www.linkedin.com/in/bryce-mullen).

## License

This project is licensed under the AP2 License - see the [LICENSE](LICENSE) file for details


`
)

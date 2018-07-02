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

// NewEnvCredSession takes tokens from systems enviomantal varables.
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
	do(credPld, response, "{{.File}}/{{.Name}}")
	return response
}`
)

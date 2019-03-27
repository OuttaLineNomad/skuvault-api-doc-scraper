package templates

const (
	// SVMain constant template for skuvault main package.
	SVMain = `
package skuvault

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	// url is a constant for all {{.Company}} calls
	url = "{{.URL}}"
)

// Error struct to store custom error message.
type Error struct {
	Func string
	Err  error
}

func (er *Error) Error() string {
	return "skuvualt: " + er.Func + ": " + er.Err.Error()
}

// Ctr is the controls the flow of endpoints.
type Ctr struct {
	{{- range .CtrBunch}}
	{{.Proper}} *{{.Cred}}
	{{- end}}
}

// LoginCredentials hold credentials to sign into SKU Vault API for endpoints.
type LoginCredentials struct {
	TenantToken string
	UserToken   string
}

{{- range .CtrBunch}}
// {{.Cred}} hold credentials to sign into SKU Vault API for {{.Name}} endpoints.
type {{.Cred}} struct {
	LoginCredentials
}
{{- end}}

// New takes tokens from systems environmental varables.
// TENANT_TOKEN and USER_TOKEN
func New() *Ctr {
	return NewSession(os.Getenv("SV_TENANT_TOKEN"), os.Getenv("SV_USER_TOKEN"))
}

// NewSession creates a new session sets credentails to make call
func NewSession(tTok, uTok string) *Ctr {
	svc := LoginCredentials{
		TenantToken: tTok,
		UserToken:   uTok,
	}

	return &Ctr{
		{{- range .CtrBunch}}
		{{.Proper}}:  &{{.Cred}}{svc},
		{{- end}}
	}
}

// postGetTokens payload sent to Sku Vault.
type postGetTokens struct {
	*GetTokens
}

// GetTokensResponse is a automatically generated struct from json provided by sku vault's api docs.
type GetTokensResponse struct {
	TenantToken string 
	UserToken   string
}

// GetTokens is a automatically generated struct from json provided by sku vault's api docs.
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
func do(pld interface{}, response interface{}, endPoint string) error {
	fullURL := url + endPoint
	bt, err := json.Marshal(pld)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(bt)
	req, err := http.NewRequest(http.MethodPost, fullURL, payload)
	if err != nil {
		return err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var respErr error
	if resp.StatusCode != 200 {
		respErr = errors.New(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, response)
	if err != nil {
		return err
	}

	return respErr
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
func (lc *{{.File0}}LoginCredentials) {{.Proper}}(pld *{{.File}}.{{.Proper}}) (*{{.File}}.{{.Proper}}Response, error) {
	credPld := &post{{.Proper}} {
		{{.Proper}}:       pld,
		{{.File0}}LoginCredentials: lc,
	}

	response := &{{.File}}.{{.Proper}}Response{}
	err := do(credPld, response, "{{.File}}{{.Name}}")
	if err != nil {
		return nil, &Error{"{{.Proper}}()", err}
	}

	return response, nil
}`
)

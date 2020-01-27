package crtsh

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// A Data struct for returning data.
type Data struct {
	Domain       string
	Timeout      time.Duration
	Certificates []Certificate
	Error        bool   `json:"error"`
	ErrorMessage string `json:"error_message,omitempty"`
}

// TODO: translating timestamps

// Certificate struct for crt.sh certificates
type Certificate struct {
	IssuerCAID        int    `json:"issuer_ca_id,omitempty"`        // "issuer_ca_id": 62131,
	IssuerName        string `json:"issuer_name,omitempty"`         //"issuer_name": "C=US, O=DigiCert Inc, OU=www.digicert.com, CN=Thawte RSA CA 2018",
	NameValue         string `json:"name_value,omitempty"`          // "name_value": "*.domain.eu",
	MinCertID         int    `json:"min_cert_id,omitempty"`         // "min_cert_id": 2141165848,
	MinEntryTimestamp string `json:"min_entry_timestamp,omitempty"` // "min_entry_timestamp": "2019-11-22T13:16:54.343",
	NoteBefore        string `json:"not_before,omitempty"`          // "not_before": "2019-11-22T00:00:00",
	NotAfter          string `json:"not_after,omitempty"`           // "not_after": "2020-11-21T12:00:00"
}

// https://crt.sh/?q=%25.domain.eu&output=json

// Get function for pulling certificates from crt.sh
func Get(domain string, timeout time.Duration) *Data {
	data := new(Data)
	data.Domain = domain
	data.Timeout = timeout

	url := "https://crt.sh/?q=%25." + domain + "&output=json"

	spaceClient := http.Client{
		Timeout: time.Second * timeout,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		data.Error = true
		data.ErrorMessage = err.Error()
		return data
	}

	req.Header.Set("User-Agent", "Mozilla")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		data.Error = true
		data.ErrorMessage = err.Error()
		return data
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		data.Error = true
		data.ErrorMessage = err.Error()
		return data
	}

	certificates := []Certificate{}
	jsonErr := json.Unmarshal(body, &certificates)
	if jsonErr != nil {
		data.Error = true
		data.ErrorMessage = jsonErr.Error()
		return data
	}
	data.Certificates = certificates

	return data
}

/*

TODO:

Passing string to sime:
	layout := "2006-01-02T15:04:05.000"
	str := "2019-11-22T13:16:54.343"
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
*/

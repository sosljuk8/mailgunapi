package client

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

type MgClient struct {
	MgAPIURL string
	MgAPIKey string
	Client   *http.Client
}

func NewMgClient() MgClient {
	return MgClient{
		os.Getenv("MG_DOMAIN"),
		os.Getenv("MG_API_KEY"),
		http.DefaultClient,
	}
}

func (mgc *MgClient) FormatEmailRequest(from, to, subject, body string) (r *http.Request, err error) {
	data := url.Values{}
	data.Add("from", from)
	data.Add("to", to)
	data.Add("subject", subject)
	data.Add("text", body)
	r, err = http.NewRequest(http.MethodPost, mgc.MgAPIURL+"/messages", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.SetBasicAuth("api", mgc.MgAPIKey)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return r, nil
}

package config

import (
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/docs/v1"
)

type Config struct {
	TimeToRepeat int         `json:"time_to_repeat"`
	URLFromParce string      `json:"url_from_parce"`
	DocumentID   string      `json:"document_id"`
	ServiceAuth  ServiceAuth `json:"service_auth"`
}

type ServiceAuth struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

func ReadConfig() (*Config, *http.Client, error) {
	var conf *Config

	data, err := os.ReadFile("./config.json")
	if err != nil {
		return nil, nil, err
	}

	if err := json.Unmarshal(data, &conf); err != nil {
		return nil, nil, err
	}

	if conf.TimeToRepeat == 0 {
		conf.TimeToRepeat = 5
	}

	c := &jwt.Config{
		Email:      conf.ServiceAuth.ClientEmail,
		PrivateKey: []byte(conf.ServiceAuth.PrivateKey),
		Scopes: []string{
			docs.DocumentsScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := c.Client(oauth2.NoContext)

	return conf, client, err
}

package infrastructure

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Director interface {
	GetDeployments() (BoshDeployments, error)
}

type BoshDirector struct {
	host     string
	username string
	password string
}

func NewDirector(host, username, password string) Director {
	return &BoshDirector{
		host:     host,
		username: username,
		password: password,
	}
}

func (d *BoshDirector) GetDeployments() (BoshDeployments, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Get(fmt.Sprintf("https://%s:%s@%s/deployments", d.username, d.password, d.host))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	deployments := []boshDeployment{}
	err = json.Unmarshal([]byte(data), &deployments)
	if err != nil {
		return nil, err
	}

	deps := BoshDeployments{}
	for _, d := range deployments {
		deps = append(deps, BoshDeployment{
			Name: d.Name,
		})
	}

	return deps, nil
}

type boshRelease struct {
	Name    string
	Version string
}
type boshDeployment struct {
	Name     string
	Releases []boshRelease
}

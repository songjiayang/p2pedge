package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	IpfsHost   = "http://localhost:5001"
	IpfsURL, _ = url.Parse(IpfsHost)
)

type IpfsIdentity struct {
	ID string `json:"ID"`
}

func GetIpfsNodeID() (*IpfsIdentity, error) {
	resp, err := http.DefaultClient.Get(IpfsHost + "/api/v0/id")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var id IpfsIdentity

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &id); err != nil {
		return nil, err
	}

	return &id, nil
}

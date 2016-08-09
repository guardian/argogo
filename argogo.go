package argogo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type ArgoResponse struct {
	Data  json.RawMessage `json:"data"`
	Links []string        `json:"links"`
	Uri   string          `json:"uri"`
}

type ArgoEntity interface {
	UnmarshalArgoData(*ArgoResponse) error
}

func GetArgoEntity(r *http.Response, v ArgoEntity) error {
	responseBody := r.Body

	defer responseBody.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(responseBody)
	s := buf.String()

	os.Stderr.WriteString("response.Body:" + s)

	responseBodyBytes, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return err
	}

	argoResponse := ArgoResponse{}

	if err := json.Unmarshal(responseBodyBytes, &argoResponse); err != nil {
		os.Stderr.WriteString("Cannot unmarshal response body to ArgoResponse!")
		return err
	}

	return v.UnmarshalArgoData(&argoResponse)
}

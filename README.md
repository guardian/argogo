# argogo
Golang wrapper for making requests to [argo-rest APIs](https://github.com/argo-rest).

## Usage

Provides the `ArgoResponse` type:
 
```go
type ArgoResponse struct {
	Data  json.RawMessage `json:"data"`
	Links []string        `json:"links"`
	Uri   string          `json:"uri"`
}
```

To unmarshall data to a known type:

```go
type SomeData struct {
        SomeValue string `json:"someValue"`
}

type MyArgoResponse struct {
        argo.ArgoResponse
        UnmarshalledData SomeData 
}

func (r MyArgoResponse) UnmarshalArgoData(argoResponse *argo.ArgoResponse) error {
        var someData SomeData 

        err := json.Unmarshal(argoResponse.Data, &someData)
        r.UnmarshalledData = storeAccess

        return err
}
```

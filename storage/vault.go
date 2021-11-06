package storage

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
)

var token = "s.5TpM91vpJMI8AYFGV0MWFD6J"
var vault_addr = "http://34.142.96.76:8200"

var VaultObj Vault

func InitVault() {
	client, err := getClient()
	if err != nil {
		log.Fatalf("Error in create vault client = %+v", err)
	}

	client.SetToken(token)

	VaultObj = Vault{Client: client}

}

func getClient() (*api.Client, error) {

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	config := &api.Config{
		Address:    vault_addr,
		HttpClient: httpClient,
	}

	return api.NewClient(config)

}

func convertStructToMap(obj interface{}) (map[string]interface{}, error) {
	writeObj := make(map[string]interface{})

	b, err := json.Marshal(obj)

	if err != nil {
		return writeObj, err
	}

	err = json.Unmarshal(b, &writeObj)

	return writeObj, err
}

func (v *Vault) Save(path string, obj interface{}) error {

	// writeObj, err := convertStructToMap(obj)
	// if err != nil {
	// 	log.Fatalf("Error in converting JSON: %+v", err)
	// 	return err
	// }

	log.Printf("obj = %+v", obj)
	s, err := v.Client.Logical().Write(path, map[string]interface{}{"data": obj})
	log.Printf("s = %+v", s)
	if err != nil {
		log.Fatalf("Error in Save vault: %+v", err)
		return err
	}

	log.Printf("Vault is stored in path %s", path)
	return nil

}

func (v *Vault) Update(path string, obj interface{}) error {
	// writeObj, err := convertStructToMap(obj)
	// if err != nil {
	// 	log.Fatalf("Error in converting JSON: %+v", err)
	// 	return err
	// }

	_, err := v.Client.Logical().Write(path, map[string]interface{}{"data": obj})

	if err != nil {
		log.Fatalf("Error in Update vault: %+v", err)
		return err
	}

	log.Printf("Vault is updated in path %s", path)
	return err
}

func (v *Vault) Get(path string) error {
	s, err := v.Client.Logical().Read(path)

	if err != nil {
		log.Fatalf("Error in get vault: %+v", err)
		return err
	}

	writeObj := V2APIKey{}

	b, err := json.Marshal(s.Data["data"])

	err = json.Unmarshal(b, &writeObj)

	log.Println("popoppppppppppp, ", writeObj)
	return err
}

func (v *Vault) delete() {

}

package main

import (
	"fmt"
	"log"

	"vault/storage"

	"github.com/google/uuid"
)

func main() {
	log.Println("Vault POC...")

	storage.InitVault()
	log.Println("------------------------------------")
	v2ApiKey()
	log.Println("------------------------------------")
	v3ApiKey()
	log.Println("------------------------------------")
}

func getPath(v string, key string) string {
	return fmt.Sprintf("secret/data/test-sky-%s-%s", v, key)
}

func v2ApiKey() {

	v2 := storage.V2APIKey{
		ApiKey:   uuid.New().String(),
		Email:    "s@s.com",
		Active:   true,
		Comments: "",
	}

	v21 := storage.V2APIKey{
		ApiKey:   uuid.New().String(),
		Email:    "s1@s1.com",
		Active:   true,
		Comments: "",
	}

	err := storage.VaultObj.Save(getPath("v2", v2.ApiKey), &v2)
	if err != nil {
		return
	}
	err = storage.VaultObj.Save(getPath("v2", v21.ApiKey), &v21)
	if err != nil {
		return
	}

	err = storage.VaultObj.Get(getPath("v2", v2.ApiKey))
	if err != nil {
		return
	}

	v2.Email = "SS@s.com"
	err = storage.VaultObj.Update(getPath("v2", v2.ApiKey), &v2)
	if err != nil {
		return
	}

	storage.VaultObj.Get(getPath("v2", v2.ApiKey))
}

func v3ApiKey() {

	v3 := storage.V3APIKey{
		ApiKey:   uuid.New().String(),
		Email:    "s3@s3.com",
		Active:   true,
		Comments: "",
		User:     "bharathi",
	}

	v31 := storage.V3APIKey{
		ApiKey:   uuid.New().String(),
		Email:    "s31@s31.com",
		Active:   true,
		Comments: "",
		User:     "bharathi",
	}

	err := storage.VaultObj.Save(getPath("v3", v3.ApiKey), &v3)
	if err != nil {
		return
	}
	err = storage.VaultObj.Save(getPath("v3", v31.ApiKey), &v31)
	if err != nil {
		return
	}

	err = storage.VaultObj.Get(getPath("v3", v3.ApiKey))
	if err != nil {
		return
	}

	v3.Email = "S3S@s3.com"
	err = storage.VaultObj.Update(getPath("v3", v3.ApiKey), &v3)
	if err != nil {
		return
	}

	storage.VaultObj.Get(getPath("v3", v3.ApiKey))
}

package main

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault/api"
)

const address string = "http://127.0.0.1:8200"
const token string = "hvs.GKTMH6CCwM4sojBflPADDOcE"

func main() {
	config := vault.DefaultConfig()
	config.Address = address

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	// Authenticate
	// WARNING: This quickstart uses the root token for our Vault dev server.
	// Don't do this in production!
	client.SetToken(token)

	ctx := context.Background()

	_, err = client.KVv2("secret").Delete(ctx, "my-secret-password")
	if err != nil {
		log.Fatalf("Unable to delete the latest version of the super secret password from the vault. Reason: %v", err)
	}
	log.Println("Deleted the latest version of the super secret password from the vault")
}

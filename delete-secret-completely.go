package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	// Authenticate
	// WARNING: This quickstart uses the root token for our Vault dev server.
	// Don't do this in production!
	client.SetToken(token)

	ctx := context.Background()

	err = client.KVv2("secret").DeleteMetadata(ctx, "my-secret-password")
	if err != nil {
		log.Fatalf("Unable to entirely delete the super secret password from the vault. Reason: %v", err)
	}
	log.Println("Deleted the latest version of the super secret password from the vault")
}

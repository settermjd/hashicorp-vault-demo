package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

const password string = "Hashi12345"

func main() {
	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	secretData := map[string]interface{}{
		"password": password,
	}

	ctx := context.Background()

	_, err = client.KVv2("secret").Delete(ctx, "my-secret-password")
	if err != nil {
		log.Fatalf("Unable to delete the latest version of the secret from the vault. Reason: %v", err)
	}
	log.Println("Delete the latest version of the secret from the vault")
}

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

	// Write a secret
	_, err = client.KVv2("secret").Put(ctx, "my-secret-password", secretData)
	if err != nil {
		log.Fatalf("Unable to write secret [%s] to the vault. Reason: %v.", secretData, err)
	}
	log.Println("Super secret password written successfully to the vault.")

	// Read a secret
	secret, err := client.KVv2("secret").Get(ctx, "my-secret-password")
	if err != nil {
		log.Fatalf("Unable to read the super secret password from the vault: %v", err)
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	if value != password {
		log.Fatalf("Unexpected super secret password value %q was retrieved from the vault", value)
	}

	log.Printf("Super secret password [%s] was retrieved.\n", value)
}

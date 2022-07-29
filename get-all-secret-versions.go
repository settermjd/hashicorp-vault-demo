package main

import (
	"context"
	"log"
	"os"
	"time"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	ctx := context.Background()

	// Get all versions of a secret, if it exists
	versions, err := client.KVv2("secret").GetVersionsAsList(ctx, "my-secret-password")
	if err != nil {
		log.Fatalf("Unable to retrieve all versions of the super secret password from the vault. Reason: %v", err)
	}

	for _, version := range versions {
		deleted := "Not deleted"
		if !version.DeletionTime.IsZero() {
			deleted = version.DeletionTime.Format(time.UnixDate)
		}

		secret, err := client.KVv2("secret").GetVersion(ctx, "my-secret-password", version.Version)
		if err != nil {
			log.Fatalf(
				"Unable to retrieve version %d of the super secret password from the vault. Reason: %v",
				err,
			)
		}
		value, ok := secret.Data["password"].(string)

		if ok {
			log.Printf(
				"Version: %d. Created at: %s. Deleted at: %s. Destroyed: %t. Value: '%s'.\n",
				version.Version,
				version.CreatedTime.Format(time.UnixDate),
				deleted,
				version.Destroyed,
				value,
			)
		}
	}
}

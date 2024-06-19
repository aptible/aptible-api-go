package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aptible/aptible-api-go/aptibleapi"
)

func main() {
	// Provide the API token through the context
	ctx := context.WithValue(context.Background(), aptibleapi.ContextAPIKeys, map[string]aptibleapi.APIKey{
		"token": {
			Prefix: "Bearer",
			Key:    os.Getenv("APTIBLE_ACCESS_TOKEN"),
		},
	})

	// Optionally, change the server URL
	if url := os.Getenv("APTIBLE_API_ROOT_URL"); url != "" {
		ctx = context.WithValue(ctx, aptibleapi.ContextServerVariables, map[string]string{"url": url})
	}

	client := aptibleapi.NewAPIClient(aptibleapi.NewAPIConfiguration())

	listApps(client, ctx)
}

func listApps(client *aptibleapi.APIClient, ctx context.Context) {
	// List all of the Apps the user has access to
	resp, r, err := client.AppsAPI.ListApps(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListApps`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %+v\n", r)
		os.Exit(1)
		return
	}

	// Print the results if there was no error
	for _, app := range resp.Embedded.Apps {
		fmt.Println(app.Handle)
	}
}

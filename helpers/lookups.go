package helpers

import (
	"context"
	"fmt"

	"github.com/aptible/aptible-api-go/aptibleapi"
)

func GetStackByName(ctx context.Context, client *aptibleapi.APIClient, name string) (*aptibleapi.Stack, error) {
	page := int32(1)
	for {
		resp, _, err := client.StacksAPI.ListStacks(ctx).Page(page).Execute()
		if err != nil {
			return nil, fmt.Errorf("error listing stacks: %w", err)
		}
		for i := range resp.Embedded.Stacks {
			if resp.Embedded.Stacks[i].Name == name {
				return &resp.Embedded.Stacks[i], nil
			}
		}
		if len(resp.Embedded.Stacks) == 0 || page*resp.PerPage >= resp.TotalCount {
			break
		}
		page++
	}
	return nil, fmt.Errorf("stack with name %q not found", name)
}

func GetDatabaseImageByTypeAndVersion(ctx context.Context, client *aptibleapi.APIClient, imageType, version string) (*aptibleapi.DatabaseImage, error) {
	page := int32(1)
	for {
		resp, _, err := client.ImagesAPI.ListDatabaseImages(ctx).Page(page).Execute()
		if err != nil {
			return nil, fmt.Errorf("error listing database images: %w", err)
		}
		for i := range resp.Embedded.DatabaseImages {
			img := &resp.Embedded.DatabaseImages[i]
			if img.Type == imageType && img.Version == version {
				return img, nil
			}
		}
		if len(resp.Embedded.DatabaseImages) == 0 || page*resp.PerPage >= resp.TotalCount {
			break
		}
		page++
	}
	return nil, fmt.Errorf("database image with type %q and version %q not found", imageType, version)
}

func GetServiceForAppByName(ctx context.Context, client *aptibleapi.APIClient, appID int32, processType string) (*aptibleapi.Service, error) {
	resp, _, err := client.ServicesAPI.ListServicesForApp(ctx, appID).Execute()
	if err != nil {
		return nil, fmt.Errorf("error listing services for app %d: %w", appID, err)
	}
	for i := range resp.Embedded.Services {
		if resp.Embedded.Services[i].ProcessType == processType {
			return &resp.Embedded.Services[i], nil
		}
	}
	return nil, fmt.Errorf("service with process type %q not found for app %d", processType, appID)
}

func GetReplicaByHandle(ctx context.Context, client *aptibleapi.APIClient, databaseID int32, handle string) (*aptibleapi.Database, error) {
	resp, _, err := client.DatabasesAPI.ListReplicasForDatabase(ctx, databaseID).Execute()
	if err != nil {
		return nil, fmt.Errorf("error listing replicas for database %d: %w", databaseID, err)
	}
	for i := range resp.Embedded.Databases {
		if resp.Embedded.Databases[i].Handle == handle {
			return &resp.Embedded.Databases[i], nil
		}
	}
	return nil, fmt.Errorf("replica with handle %q not found for database %d", handle, databaseID)
}

package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/aptible/aptible-api-go/aptibleapi"
)

// WaitForOperation polls an operation until it reaches a terminal state.
// Returns (true, nil) if the resource was deleted (404), (false, nil) on success,
// or (false, error) on failure or timeout.
func WaitForOperation(ctx context.Context, client *aptibleapi.APIClient, operationID int32) (bool, error) {
	for {
		select {
		case <-ctx.Done():
			return false, fmt.Errorf("timed out waiting for operation %d: %w", operationID, ctx.Err())
		default:
		}

		op, resp, err := client.OperationsAPI.GetOperation(ctx, operationID).Execute()
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				return true, nil
			}
			return false, fmt.Errorf("error getting operation %d: %w", operationID, err)
		}

		switch op.Status {
		case "succeeded":
			return false, nil
		case "failed":
			return false, fmt.Errorf("operation %d failed", operationID)
		}

		time.Sleep(5 * time.Second)
	}
}

func DeleteApp(ctx context.Context, client *aptibleapi.APIClient, appID int32) (bool, error) {
	op, _, err := client.OperationsAPI.CreateOperationForApp(ctx, appID).
		CreateOperationRequest(aptibleapi.CreateOperationRequest{Type: "deprovision"}).Execute()
	if err != nil {
		return false, err
	}
	return WaitForOperation(ctx, client, op.Id)
}

func DeleteDatabase(ctx context.Context, client *aptibleapi.APIClient, databaseID int32) (bool, error) {
	op, _, err := client.OperationsAPI.CreateOperationForDatabase(ctx, databaseID).
		CreateOperationRequest(aptibleapi.CreateOperationRequest{Type: "deprovision"}).Execute()
	if err != nil {
		return false, err
	}
	return WaitForOperation(ctx, client, op.Id)
}

func DeleteEndpoint(ctx context.Context, client *aptibleapi.APIClient, vhostID int32) (bool, error) {
	op, _, err := client.OperationsAPI.CreateOperationForVhost(ctx, vhostID).
		CreateOperationRequest(aptibleapi.CreateOperationRequest{Type: "deprovision"}).Execute()
	if err != nil {
		return false, err
	}
	return WaitForOperation(ctx, client, op.Id)
}

func DeleteLogDrain(ctx context.Context, client *aptibleapi.APIClient, logDrainID int32) (bool, error) {
	op, _, err := client.OperationsAPI.CreateOperationForLogDrain(ctx, logDrainID).
		CreateOperationRequest(aptibleapi.CreateOperationRequest{Type: "deprovision"}).Execute()
	if err != nil {
		return false, err
	}
	return WaitForOperation(ctx, client, op.Id)
}

func DeleteMetricDrain(ctx context.Context, client *aptibleapi.APIClient, metricDrainID int32) (bool, error) {
	op, _, err := client.OperationsAPI.CreateOperationForMetricDrain(ctx, metricDrainID).
		CreateOperationRequest(aptibleapi.CreateOperationRequest{Type: "deprovision"}).Execute()
	if err != nil {
		return false, err
	}
	return WaitForOperation(ctx, client, op.Id)
}

package common

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/waldur/terraform-provider-waldur/internal/client"
)

// ResolveResourceUUID extracts the resource UUID from a marketplace order response.
func ResolveResourceUUID(orderRes *OrderDetails) string {
	if orderRes == nil {
		return ""
	}
	if orderRes.ResourceUuid != nil && *orderRes.ResourceUuid != "" {
		return *orderRes.ResourceUuid
	}
	if orderRes.MarketplaceResourceUuid != nil && *orderRes.MarketplaceResourceUuid != "" {
		return *orderRes.MarketplaceResourceUuid
	}
	return ""
}

// WaitForOrder blocks until a marketplace order reaches the "done" state.
func WaitForOrder(ctx context.Context, c *client.Client, orderUUID string, timeout time.Duration) (*OrderDetails, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"pending", "pending-consumer", "pending-provider", "pending-project", "pending-start-date", "executing", "created"},
		Target:  []string{"done"},
		Refresh: func() (interface{}, string, error) {
			var res OrderDetails
			err := c.GetURL(ctx, fmt.Sprintf("/api/marketplace-orders/%s/", orderUUID), &res)
			if err != nil {
				return nil, "", err
			}

			state := ""
			if res.State != nil {
				state = *res.State
			}
			if state == "erred" || state == "rejected" || state == "canceled" {
				msg := ""
				if res.ErrorMessage != nil {
					msg = *res.ErrorMessage
				}
				return &res, "failed", fmt.Errorf("order failed: %s", msg)
			}
			return &res, state, nil
		},
		Timeout:    timeout,
		Delay:      DefaultPollDelay,
		MinTimeout: DefaultPollMinTimeout,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}

	return rawResult.(*OrderDetails), nil
}

// ResourceWithState defines the interface for resources that have a state and error message.
type ResourceWithState interface {
	GetState() string
	GetErrorMessage() string
}

// WaitForResource blocks until a resource reaches the "OK" state.
func WaitForResource[T ResourceWithState](ctx context.Context, getResource func(context.Context) (T, error), timeout time.Duration) (T, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATION_SCHEDULED", "CREATING", "UPDATE_SCHEDULED", "UPDATING", "DELETION_SCHEDULED", "DELETING"},
		Target:  []string{"OK"},
		Refresh: func() (interface{}, string, error) {
			res, err := getResource(ctx)
			if err != nil {
				return nil, "", err
			}

			// Use interface methods directly on res
			state := res.GetState()

			if state == "ERRED" {
				msg := res.GetErrorMessage()
				if msg == "" {
					msg = "unknown error"
				}
				return res, "failed", fmt.Errorf("%s", msg)
			}

			return res, state, nil
		},
		Timeout:    timeout,
		Delay:      DefaultPollDelay,
		MinTimeout: DefaultPollMinTimeout,
	}

	rawResult, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		var zero T
		return zero, err
	}

	return rawResult.(T), nil
}

// WaitForDeletion blocks until a resource is gone (404).
func WaitForDeletion[T ResourceWithState](ctx context.Context, getResource func(context.Context) (T, error), timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATION_SCHEDULED", "CREATING", "UPDATE_SCHEDULED", "UPDATING", "DELETION_SCHEDULED", "DELETING", "OK"},
		Refresh: func() (interface{}, string, error) {
			res, err := getResource(ctx)
			if err != nil {
				if IsNotFoundError(err) {
					return nil, "", nil
				}
				return nil, "", err
			}

			// Use interface methods
			state := res.GetState()

			if state == "ERRED" {
				msg := res.GetErrorMessage()
				if msg == "" {
					msg = "unknown error"
				}
				return res, "failed", fmt.Errorf("%s", msg)
			}

			return res, state, nil
		},
		Timeout:    timeout,
		Delay:      DefaultPollDelay,
		MinTimeout: DefaultPollMinTimeout,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

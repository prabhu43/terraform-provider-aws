package appstream

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

const (
	// stackOperationTimeout Maximum amount of time to wait for Stack operation eventual consistency
	stackOperationTimeout = 4 * time.Minute

	// fleetStateTimeout Maximum amount of time to wait for the statusFleetState to be RUNNING or STOPPED
	fleetStateTimeout = 180 * time.Minute
	// fleetOperationTimeout Maximum amount of time to wait for Fleet operation eventual consistency
	fleetOperationTimeout = 15 * time.Minute
	// imageBuilderStateTimeout Maximum amount of time to wait for the statusImageBuilderState to be RUNNING
	// or for the ImageBuilder to be deleted
	imageBuilderStateTimeout = 60 * time.Minute
)

// waitStackStateDeleted waits for a deleted stack
func waitStackStateDeleted(ctx context.Context, conn *appstream.AppStream, name string) (*appstream.Stack, error) {
	stateConf := &resource.StateChangeConf{
		Target:  []string{"NotFound", "Unknown"},
		Refresh: statusStackState(ctx, conn, name),
		Timeout: stackOperationTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*appstream.Stack); ok {
		if errors := output.StackErrors; len(errors) > 0 {
			var errs *multierror.Error

			for _, err := range errors {
				errs = multierror.Append(errs, fmt.Errorf("%s: %s", aws.StringValue(err.ErrorCode), aws.StringValue(err.ErrorMessage)))
			}

			tfresource.SetLastError(err, errs.ErrorOrNil())
		}

		return output, err
	}

	return nil, err
}

// waitFleetStateRunning waits for a fleet running
func waitFleetStateRunning(ctx context.Context, conn *appstream.AppStream, name string) (*appstream.Fleet, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{appstream.FleetStateStarting},
		Target:  []string{appstream.FleetStateRunning},
		Refresh: statusFleetState(ctx, conn, name),
		Timeout: fleetStateTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*appstream.Fleet); ok {
		if errors := output.FleetErrors; len(errors) > 0 {
			var errs *multierror.Error

			for _, err := range errors {
				errs = multierror.Append(errs, fmt.Errorf("%s: %s", aws.StringValue(err.ErrorCode), aws.StringValue(err.ErrorMessage)))
			}

			tfresource.SetLastError(err, errs.ErrorOrNil())
		}

		return output, err
	}

	return nil, err
}

// waitFleetStateStopped waits for a fleet stopped
func waitFleetStateStopped(ctx context.Context, conn *appstream.AppStream, name string) (*appstream.Fleet, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{appstream.FleetStateStopping},
		Target:  []string{appstream.FleetStateStopped},
		Refresh: statusFleetState(ctx, conn, name),
		Timeout: fleetStateTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*appstream.Fleet); ok {
		if errors := output.FleetErrors; len(errors) > 0 {
			var errs *multierror.Error

			for _, err := range errors {
				errs = multierror.Append(errs, fmt.Errorf("%s: %s", aws.StringValue(err.ErrorCode), aws.StringValue(err.ErrorMessage)))
			}

			tfresource.SetLastError(err, errs.ErrorOrNil())
		}

		return output, err
	}

	return nil, err
}

// waitImageBuilderStateRunning waits for a ImageBuilder running
func waitImageBuilderStateRunning(ctx context.Context, conn *appstream.AppStream, name string) (*appstream.ImageBuilder, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{appstream.ImageBuilderStatePending},
		Target:  []string{appstream.ImageBuilderStateRunning},
		Refresh: statusImageBuilderState(ctx, conn, name),
		Timeout: imageBuilderStateTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*appstream.ImageBuilder); ok {
		if state, errors := aws.StringValue(output.State), output.ImageBuilderErrors; state == appstream.ImageBuilderStateFailed && len(errors) > 0 {
			var errs *multierror.Error

			for _, err := range errors {
				errs = multierror.Append(errs, fmt.Errorf("%s: %s", aws.StringValue(err.ErrorCode), aws.StringValue(err.ErrorMessage)))
			}

			tfresource.SetLastError(err, errs.ErrorOrNil())
		}

		return output, err
	}

	return nil, err
}

// waitImageBuilderStateDeleted waits for a ImageBuilder deleted
func waitImageBuilderStateDeleted(ctx context.Context, conn *appstream.AppStream, name string) (*appstream.ImageBuilder, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{appstream.ImageBuilderStatePending, appstream.ImageBuilderStateDeleting},
		Target:  []string{},
		Refresh: statusImageBuilderState(ctx, conn, name),
		Timeout: imageBuilderStateTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*appstream.ImageBuilder); ok {
		if state, errors := aws.StringValue(output.State), output.ImageBuilderErrors; state == appstream.ImageBuilderStateFailed && len(errors) > 0 {
			var errs *multierror.Error

			for _, err := range errors {
				errs = multierror.Append(errs, fmt.Errorf("%s: %s", aws.StringValue(err.ErrorCode), aws.StringValue(err.ErrorMessage)))
			}

			tfresource.SetLastError(err, errs.ErrorOrNil())
		}

		return output, err
	}

	return nil, err
}

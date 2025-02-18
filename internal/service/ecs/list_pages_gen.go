// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeCapacityProviders -Export=yes"; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func DescribeCapacityProvidersPages(conn *ecs.ECS, input *ecs.DescribeCapacityProvidersInput, fn func(*ecs.DescribeCapacityProvidersOutput, bool) bool) error {
	return DescribeCapacityProvidersPagesWithContext(context.Background(), conn, input, fn)
}

func DescribeCapacityProvidersPagesWithContext(ctx context.Context, conn *ecs.ECS, input *ecs.DescribeCapacityProvidersInput, fn func(*ecs.DescribeCapacityProvidersOutput, bool) bool) error {
	for {
		output, err := conn.DescribeCapacityProvidersWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}

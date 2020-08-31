// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package migrationhubconfigiface provides an interface to enable mocking the AWS Migration Hub Config service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package migrationhubconfigiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
)

// MigrationHubConfigAPI provides an interface to enable mocking the
// migrationhubconfig.MigrationHubConfig service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Migration Hub Config.
//    func myFunc(svc migrationhubconfigiface.MigrationHubConfigAPI) bool {
//        // Make svc.CreateHomeRegionControl request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := migrationhubconfig.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMigrationHubConfigClient struct {
//        migrationhubconfigiface.MigrationHubConfigAPI
//    }
//    func (m *mockMigrationHubConfigClient) CreateHomeRegionControl(input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMigrationHubConfigClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MigrationHubConfigAPI interface {
	CreateHomeRegionControl(*migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error)
	CreateHomeRegionControlWithContext(aws.Context, *migrationhubconfig.CreateHomeRegionControlInput, ...request.Option) (*migrationhubconfig.CreateHomeRegionControlOutput, error)
	CreateHomeRegionControlRequest(*migrationhubconfig.CreateHomeRegionControlInput) (*request.Request, *migrationhubconfig.CreateHomeRegionControlOutput)

	DescribeHomeRegionControls(*migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error)
	DescribeHomeRegionControlsWithContext(aws.Context, *migrationhubconfig.DescribeHomeRegionControlsInput, ...request.Option) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error)
	DescribeHomeRegionControlsRequest(*migrationhubconfig.DescribeHomeRegionControlsInput) (*request.Request, *migrationhubconfig.DescribeHomeRegionControlsOutput)

	DescribeHomeRegionControlsPages(*migrationhubconfig.DescribeHomeRegionControlsInput, func(*migrationhubconfig.DescribeHomeRegionControlsOutput, bool) bool) error
	DescribeHomeRegionControlsPagesWithContext(aws.Context, *migrationhubconfig.DescribeHomeRegionControlsInput, func(*migrationhubconfig.DescribeHomeRegionControlsOutput, bool) bool, ...request.Option) error

	GetHomeRegion(*migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error)
	GetHomeRegionWithContext(aws.Context, *migrationhubconfig.GetHomeRegionInput, ...request.Option) (*migrationhubconfig.GetHomeRegionOutput, error)
	GetHomeRegionRequest(*migrationhubconfig.GetHomeRegionInput) (*request.Request, *migrationhubconfig.GetHomeRegionOutput)
}

var _ MigrationHubConfigAPI = (*migrationhubconfig.MigrationHubConfig)(nil)

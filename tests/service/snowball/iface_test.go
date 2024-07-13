package snowball_test

// tests for the snowball service interface for this ../../../service/snowball/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/snowball"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/snowball/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/snowball/snowball_iface"
	"github.com/stretchr/testify/assert"
)

func TestSnowballServiceCanBeMocked(t *testing.T) {
	var iface snowball_iface.IClient
	iface = &snowball.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := snowball.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelCluster", func(t *testing.T) {
        input := &snowball.CancelClusterInput{}
        output := &snowball.CancelClusterOutput{}

        mockClient.On("CancelCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CancelCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJob", func(t *testing.T) {
        input := &snowball.CancelJobInput{}
        output := &snowball.CancelJobOutput{}

        mockClient.On("CancelJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAddress", func(t *testing.T) {
        input := &snowball.CreateAddressInput{}
        output := &snowball.CreateAddressOutput{}

        mockClient.On("CreateAddress", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &snowball.CreateClusterInput{}
        output := &snowball.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &snowball.CreateJobInput{}
        output := &snowball.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLongTermPricing", func(t *testing.T) {
        input := &snowball.CreateLongTermPricingInput{}
        output := &snowball.CreateLongTermPricingOutput{}

        mockClient.On("CreateLongTermPricing", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLongTermPricing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReturnShippingLabel", func(t *testing.T) {
        input := &snowball.CreateReturnShippingLabelInput{}
        output := &snowball.CreateReturnShippingLabelOutput{}

        mockClient.On("CreateReturnShippingLabel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReturnShippingLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAddress", func(t *testing.T) {
        input := &snowball.DescribeAddressInput{}
        output := &snowball.DescribeAddressOutput{}

        mockClient.On("DescribeAddress", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAddresses", func(t *testing.T) {
        input := &snowball.DescribeAddressesInput{}
        output := &snowball.DescribeAddressesOutput{}

        mockClient.On("DescribeAddresses", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAddresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCluster", func(t *testing.T) {
        input := &snowball.DescribeClusterInput{}
        output := &snowball.DescribeClusterOutput{}

        mockClient.On("DescribeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJob", func(t *testing.T) {
        input := &snowball.DescribeJobInput{}
        output := &snowball.DescribeJobOutput{}

        mockClient.On("DescribeJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReturnShippingLabel", func(t *testing.T) {
        input := &snowball.DescribeReturnShippingLabelInput{}
        output := &snowball.DescribeReturnShippingLabelOutput{}

        mockClient.On("DescribeReturnShippingLabel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReturnShippingLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobManifest", func(t *testing.T) {
        input := &snowball.GetJobManifestInput{}
        output := &snowball.GetJobManifestOutput{}

        mockClient.On("GetJobManifest", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobUnlockCode", func(t *testing.T) {
        input := &snowball.GetJobUnlockCodeInput{}
        output := &snowball.GetJobUnlockCodeOutput{}

        mockClient.On("GetJobUnlockCode", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobUnlockCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSnowballUsage", func(t *testing.T) {
        input := &snowball.GetSnowballUsageInput{}
        output := &snowball.GetSnowballUsageOutput{}

        mockClient.On("GetSnowballUsage", ctx, input).Return(output, nil)

        result, err := mockClient.GetSnowballUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSoftwareUpdates", func(t *testing.T) {
        input := &snowball.GetSoftwareUpdatesInput{}
        output := &snowball.GetSoftwareUpdatesOutput{}

        mockClient.On("GetSoftwareUpdates", ctx, input).Return(output, nil)

        result, err := mockClient.GetSoftwareUpdates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusterJobs", func(t *testing.T) {
        input := &snowball.ListClusterJobsInput{}
        output := &snowball.ListClusterJobsOutput{}

        mockClient.On("ListClusterJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusterJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusters", func(t *testing.T) {
        input := &snowball.ListClustersInput{}
        output := &snowball.ListClustersOutput{}

        mockClient.On("ListClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCompatibleImages", func(t *testing.T) {
        input := &snowball.ListCompatibleImagesInput{}
        output := &snowball.ListCompatibleImagesOutput{}

        mockClient.On("ListCompatibleImages", ctx, input).Return(output, nil)

        result, err := mockClient.ListCompatibleImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &snowball.ListJobsInput{}
        output := &snowball.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLongTermPricing", func(t *testing.T) {
        input := &snowball.ListLongTermPricingInput{}
        output := &snowball.ListLongTermPricingOutput{}

        mockClient.On("ListLongTermPricing", ctx, input).Return(output, nil)

        result, err := mockClient.ListLongTermPricing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPickupLocations", func(t *testing.T) {
        input := &snowball.ListPickupLocationsInput{}
        output := &snowball.ListPickupLocationsOutput{}

        mockClient.On("ListPickupLocations", ctx, input).Return(output, nil)

        result, err := mockClient.ListPickupLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceVersions", func(t *testing.T) {
        input := &snowball.ListServiceVersionsInput{}
        output := &snowball.ListServiceVersionsOutput{}

        mockClient.On("ListServiceVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCluster", func(t *testing.T) {
        input := &snowball.UpdateClusterInput{}
        output := &snowball.UpdateClusterOutput{}

        mockClient.On("UpdateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJob", func(t *testing.T) {
        input := &snowball.UpdateJobInput{}
        output := &snowball.UpdateJobOutput{}

        mockClient.On("UpdateJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJobShipmentState", func(t *testing.T) {
        input := &snowball.UpdateJobShipmentStateInput{}
        output := &snowball.UpdateJobShipmentStateOutput{}

        mockClient.On("UpdateJobShipmentState", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJobShipmentState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLongTermPricing", func(t *testing.T) {
        input := &snowball.UpdateLongTermPricingInput{}
        output := &snowball.UpdateLongTermPricingOutput{}

        mockClient.On("UpdateLongTermPricing", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLongTermPricing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

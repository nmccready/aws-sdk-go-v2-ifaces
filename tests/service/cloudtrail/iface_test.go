package cloudtrail_test

// tests for the cloudtrail service interface for this ../../../service/cloudtrail/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudtrail/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudtrail/cloudtrail_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudtrailServiceCanBeMocked(t *testing.T) {
	var iface cloudtrail_iface.IClient
	iface = &cloudtrail.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudtrail.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &cloudtrail.AddTagsInput{}
        output := &cloudtrail.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelQuery", func(t *testing.T) {
        input := &cloudtrail.CancelQueryInput{}
        output := &cloudtrail.CancelQueryOutput{}

        mockClient.On("CancelQuery", ctx, input).Return(output, nil)

        result, err := mockClient.CancelQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &cloudtrail.CreateChannelInput{}
        output := &cloudtrail.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventDataStore", func(t *testing.T) {
        input := &cloudtrail.CreateEventDataStoreInput{}
        output := &cloudtrail.CreateEventDataStoreOutput{}

        mockClient.On("CreateEventDataStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventDataStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrail", func(t *testing.T) {
        input := &cloudtrail.CreateTrailInput{}
        output := &cloudtrail.CreateTrailOutput{}

        mockClient.On("CreateTrail", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &cloudtrail.DeleteChannelInput{}
        output := &cloudtrail.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventDataStore", func(t *testing.T) {
        input := &cloudtrail.DeleteEventDataStoreInput{}
        output := &cloudtrail.DeleteEventDataStoreOutput{}

        mockClient.On("DeleteEventDataStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventDataStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &cloudtrail.DeleteResourcePolicyInput{}
        output := &cloudtrail.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrail", func(t *testing.T) {
        input := &cloudtrail.DeleteTrailInput{}
        output := &cloudtrail.DeleteTrailOutput{}

        mockClient.On("DeleteTrail", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterOrganizationDelegatedAdmin", func(t *testing.T) {
        input := &cloudtrail.DeregisterOrganizationDelegatedAdminInput{}
        output := &cloudtrail.DeregisterOrganizationDelegatedAdminOutput{}

        mockClient.On("DeregisterOrganizationDelegatedAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterOrganizationDelegatedAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeQuery", func(t *testing.T) {
        input := &cloudtrail.DescribeQueryInput{}
        output := &cloudtrail.DescribeQueryOutput{}

        mockClient.On("DescribeQuery", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrails", func(t *testing.T) {
        input := &cloudtrail.DescribeTrailsInput{}
        output := &cloudtrail.DescribeTrailsOutput{}

        mockClient.On("DescribeTrails", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableFederation", func(t *testing.T) {
        input := &cloudtrail.DisableFederationInput{}
        output := &cloudtrail.DisableFederationOutput{}

        mockClient.On("DisableFederation", ctx, input).Return(output, nil)

        result, err := mockClient.DisableFederation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableFederation", func(t *testing.T) {
        input := &cloudtrail.EnableFederationInput{}
        output := &cloudtrail.EnableFederationOutput{}

        mockClient.On("EnableFederation", ctx, input).Return(output, nil)

        result, err := mockClient.EnableFederation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannel", func(t *testing.T) {
        input := &cloudtrail.GetChannelInput{}
        output := &cloudtrail.GetChannelOutput{}

        mockClient.On("GetChannel", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventDataStore", func(t *testing.T) {
        input := &cloudtrail.GetEventDataStoreInput{}
        output := &cloudtrail.GetEventDataStoreOutput{}

        mockClient.On("GetEventDataStore", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventDataStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventSelectors", func(t *testing.T) {
        input := &cloudtrail.GetEventSelectorsInput{}
        output := &cloudtrail.GetEventSelectorsOutput{}

        mockClient.On("GetEventSelectors", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventSelectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImport", func(t *testing.T) {
        input := &cloudtrail.GetImportInput{}
        output := &cloudtrail.GetImportOutput{}

        mockClient.On("GetImport", ctx, input).Return(output, nil)

        result, err := mockClient.GetImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsightSelectors", func(t *testing.T) {
        input := &cloudtrail.GetInsightSelectorsInput{}
        output := &cloudtrail.GetInsightSelectorsOutput{}

        mockClient.On("GetInsightSelectors", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsightSelectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryResults", func(t *testing.T) {
        input := &cloudtrail.GetQueryResultsInput{}
        output := &cloudtrail.GetQueryResultsOutput{}

        mockClient.On("GetQueryResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &cloudtrail.GetResourcePolicyInput{}
        output := &cloudtrail.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrail", func(t *testing.T) {
        input := &cloudtrail.GetTrailInput{}
        output := &cloudtrail.GetTrailOutput{}

        mockClient.On("GetTrail", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrailStatus", func(t *testing.T) {
        input := &cloudtrail.GetTrailStatusInput{}
        output := &cloudtrail.GetTrailStatusOutput{}

        mockClient.On("GetTrailStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrailStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &cloudtrail.ListChannelsInput{}
        output := &cloudtrail.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventDataStores", func(t *testing.T) {
        input := &cloudtrail.ListEventDataStoresInput{}
        output := &cloudtrail.ListEventDataStoresOutput{}

        mockClient.On("ListEventDataStores", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventDataStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImportFailures", func(t *testing.T) {
        input := &cloudtrail.ListImportFailuresInput{}
        output := &cloudtrail.ListImportFailuresOutput{}

        mockClient.On("ListImportFailures", ctx, input).Return(output, nil)

        result, err := mockClient.ListImportFailures(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImports", func(t *testing.T) {
        input := &cloudtrail.ListImportsInput{}
        output := &cloudtrail.ListImportsOutput{}

        mockClient.On("ListImports", ctx, input).Return(output, nil)

        result, err := mockClient.ListImports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInsightsMetricData", func(t *testing.T) {
        input := &cloudtrail.ListInsightsMetricDataInput{}
        output := &cloudtrail.ListInsightsMetricDataOutput{}

        mockClient.On("ListInsightsMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.ListInsightsMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPublicKeys", func(t *testing.T) {
        input := &cloudtrail.ListPublicKeysInput{}
        output := &cloudtrail.ListPublicKeysOutput{}

        mockClient.On("ListPublicKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListPublicKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueries", func(t *testing.T) {
        input := &cloudtrail.ListQueriesInput{}
        output := &cloudtrail.ListQueriesOutput{}

        mockClient.On("ListQueries", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &cloudtrail.ListTagsInput{}
        output := &cloudtrail.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrails", func(t *testing.T) {
        input := &cloudtrail.ListTrailsInput{}
        output := &cloudtrail.ListTrailsOutput{}

        mockClient.On("ListTrails", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLookupEvents", func(t *testing.T) {
        input := &cloudtrail.LookupEventsInput{}
        output := &cloudtrail.LookupEventsOutput{}

        mockClient.On("LookupEvents", ctx, input).Return(output, nil)

        result, err := mockClient.LookupEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEventSelectors", func(t *testing.T) {
        input := &cloudtrail.PutEventSelectorsInput{}
        output := &cloudtrail.PutEventSelectorsOutput{}

        mockClient.On("PutEventSelectors", ctx, input).Return(output, nil)

        result, err := mockClient.PutEventSelectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutInsightSelectors", func(t *testing.T) {
        input := &cloudtrail.PutInsightSelectorsInput{}
        output := &cloudtrail.PutInsightSelectorsOutput{}

        mockClient.On("PutInsightSelectors", ctx, input).Return(output, nil)

        result, err := mockClient.PutInsightSelectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &cloudtrail.PutResourcePolicyInput{}
        output := &cloudtrail.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterOrganizationDelegatedAdmin", func(t *testing.T) {
        input := &cloudtrail.RegisterOrganizationDelegatedAdminInput{}
        output := &cloudtrail.RegisterOrganizationDelegatedAdminOutput{}

        mockClient.On("RegisterOrganizationDelegatedAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterOrganizationDelegatedAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTags", func(t *testing.T) {
        input := &cloudtrail.RemoveTagsInput{}
        output := &cloudtrail.RemoveTagsOutput{}

        mockClient.On("RemoveTags", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreEventDataStore", func(t *testing.T) {
        input := &cloudtrail.RestoreEventDataStoreInput{}
        output := &cloudtrail.RestoreEventDataStoreOutput{}

        mockClient.On("RestoreEventDataStore", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreEventDataStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartEventDataStoreIngestion", func(t *testing.T) {
        input := &cloudtrail.StartEventDataStoreIngestionInput{}
        output := &cloudtrail.StartEventDataStoreIngestionOutput{}

        mockClient.On("StartEventDataStoreIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.StartEventDataStoreIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImport", func(t *testing.T) {
        input := &cloudtrail.StartImportInput{}
        output := &cloudtrail.StartImportOutput{}

        mockClient.On("StartImport", ctx, input).Return(output, nil)

        result, err := mockClient.StartImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartLogging", func(t *testing.T) {
        input := &cloudtrail.StartLoggingInput{}
        output := &cloudtrail.StartLoggingOutput{}

        mockClient.On("StartLogging", ctx, input).Return(output, nil)

        result, err := mockClient.StartLogging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQuery", func(t *testing.T) {
        input := &cloudtrail.StartQueryInput{}
        output := &cloudtrail.StartQueryOutput{}

        mockClient.On("StartQuery", ctx, input).Return(output, nil)

        result, err := mockClient.StartQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEventDataStoreIngestion", func(t *testing.T) {
        input := &cloudtrail.StopEventDataStoreIngestionInput{}
        output := &cloudtrail.StopEventDataStoreIngestionOutput{}

        mockClient.On("StopEventDataStoreIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.StopEventDataStoreIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopImport", func(t *testing.T) {
        input := &cloudtrail.StopImportInput{}
        output := &cloudtrail.StopImportOutput{}

        mockClient.On("StopImport", ctx, input).Return(output, nil)

        result, err := mockClient.StopImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopLogging", func(t *testing.T) {
        input := &cloudtrail.StopLoggingInput{}
        output := &cloudtrail.StopLoggingOutput{}

        mockClient.On("StopLogging", ctx, input).Return(output, nil)

        result, err := mockClient.StopLogging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &cloudtrail.UpdateChannelInput{}
        output := &cloudtrail.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventDataStore", func(t *testing.T) {
        input := &cloudtrail.UpdateEventDataStoreInput{}
        output := &cloudtrail.UpdateEventDataStoreOutput{}

        mockClient.On("UpdateEventDataStore", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventDataStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrail", func(t *testing.T) {
        input := &cloudtrail.UpdateTrailInput{}
        output := &cloudtrail.UpdateTrailOutput{}

        mockClient.On("UpdateTrail", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

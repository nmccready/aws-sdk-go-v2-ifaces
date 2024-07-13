package cloudwatchlogs_test

// tests for the cloudwatchlogs service interface for this ../../../service/cloudwatchlogs/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudwatchlogs/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudwatchlogs/cloudwatchlogs_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudwatchlogsServiceCanBeMocked(t *testing.T) {
	var iface cloudwatchlogs_iface.IClient
	iface = &cloudwatchlogs.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudwatchlogs.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateKmsKey", func(t *testing.T) {
        input := &cloudwatchlogs.AssociateKmsKeyInput{}
        output := &cloudwatchlogs.AssociateKmsKeyOutput{}

        mockClient.On("AssociateKmsKey", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateKmsKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelExportTask", func(t *testing.T) {
        input := &cloudwatchlogs.CancelExportTaskInput{}
        output := &cloudwatchlogs.CancelExportTaskOutput{}

        mockClient.On("CancelExportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelExportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDelivery", func(t *testing.T) {
        input := &cloudwatchlogs.CreateDeliveryInput{}
        output := &cloudwatchlogs.CreateDeliveryOutput{}

        mockClient.On("CreateDelivery", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDelivery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExportTask", func(t *testing.T) {
        input := &cloudwatchlogs.CreateExportTaskInput{}
        output := &cloudwatchlogs.CreateExportTaskOutput{}

        mockClient.On("CreateExportTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLogAnomalyDetector", func(t *testing.T) {
        input := &cloudwatchlogs.CreateLogAnomalyDetectorInput{}
        output := &cloudwatchlogs.CreateLogAnomalyDetectorOutput{}

        mockClient.On("CreateLogAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLogAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLogGroup", func(t *testing.T) {
        input := &cloudwatchlogs.CreateLogGroupInput{}
        output := &cloudwatchlogs.CreateLogGroupOutput{}

        mockClient.On("CreateLogGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLogGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLogStream", func(t *testing.T) {
        input := &cloudwatchlogs.CreateLogStreamInput{}
        output := &cloudwatchlogs.CreateLogStreamOutput{}

        mockClient.On("CreateLogStream", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLogStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteAccountPolicyInput{}
        output := &cloudwatchlogs.DeleteAccountPolicyOutput{}

        mockClient.On("DeleteAccountPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataProtectionPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteDataProtectionPolicyInput{}
        output := &cloudwatchlogs.DeleteDataProtectionPolicyOutput{}

        mockClient.On("DeleteDataProtectionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataProtectionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDelivery", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteDeliveryInput{}
        output := &cloudwatchlogs.DeleteDeliveryOutput{}

        mockClient.On("DeleteDelivery", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDelivery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeliveryDestination", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteDeliveryDestinationInput{}
        output := &cloudwatchlogs.DeleteDeliveryDestinationOutput{}

        mockClient.On("DeleteDeliveryDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeliveryDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeliveryDestinationPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteDeliveryDestinationPolicyInput{}
        output := &cloudwatchlogs.DeleteDeliveryDestinationPolicyOutput{}

        mockClient.On("DeleteDeliveryDestinationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeliveryDestinationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeliverySource", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteDeliverySourceInput{}
        output := &cloudwatchlogs.DeleteDeliverySourceOutput{}

        mockClient.On("DeleteDeliverySource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeliverySource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDestination", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteDestinationInput{}
        output := &cloudwatchlogs.DeleteDestinationOutput{}

        mockClient.On("DeleteDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLogAnomalyDetector", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteLogAnomalyDetectorInput{}
        output := &cloudwatchlogs.DeleteLogAnomalyDetectorOutput{}

        mockClient.On("DeleteLogAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLogAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLogGroup", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteLogGroupInput{}
        output := &cloudwatchlogs.DeleteLogGroupOutput{}

        mockClient.On("DeleteLogGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLogGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLogStream", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteLogStreamInput{}
        output := &cloudwatchlogs.DeleteLogStreamOutput{}

        mockClient.On("DeleteLogStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLogStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMetricFilter", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteMetricFilterInput{}
        output := &cloudwatchlogs.DeleteMetricFilterOutput{}

        mockClient.On("DeleteMetricFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMetricFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueryDefinition", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteQueryDefinitionInput{}
        output := &cloudwatchlogs.DeleteQueryDefinitionOutput{}

        mockClient.On("DeleteQueryDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueryDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteResourcePolicyInput{}
        output := &cloudwatchlogs.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRetentionPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteRetentionPolicyInput{}
        output := &cloudwatchlogs.DeleteRetentionPolicyOutput{}

        mockClient.On("DeleteRetentionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRetentionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscriptionFilter", func(t *testing.T) {
        input := &cloudwatchlogs.DeleteSubscriptionFilterInput{}
        output := &cloudwatchlogs.DeleteSubscriptionFilterOutput{}

        mockClient.On("DeleteSubscriptionFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscriptionFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountPolicies", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeAccountPoliciesInput{}
        output := &cloudwatchlogs.DescribeAccountPoliciesOutput{}

        mockClient.On("DescribeAccountPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeliveries", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeDeliveriesInput{}
        output := &cloudwatchlogs.DescribeDeliveriesOutput{}

        mockClient.On("DescribeDeliveries", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeliveries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeliveryDestinations", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeDeliveryDestinationsInput{}
        output := &cloudwatchlogs.DescribeDeliveryDestinationsOutput{}

        mockClient.On("DescribeDeliveryDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeliveryDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDeliverySources", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeDeliverySourcesInput{}
        output := &cloudwatchlogs.DescribeDeliverySourcesOutput{}

        mockClient.On("DescribeDeliverySources", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDeliverySources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDestinations", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeDestinationsInput{}
        output := &cloudwatchlogs.DescribeDestinationsOutput{}

        mockClient.On("DescribeDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExportTasks", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeExportTasksInput{}
        output := &cloudwatchlogs.DescribeExportTasksOutput{}

        mockClient.On("DescribeExportTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExportTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLogGroups", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeLogGroupsInput{}
        output := &cloudwatchlogs.DescribeLogGroupsOutput{}

        mockClient.On("DescribeLogGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLogGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLogStreams", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeLogStreamsInput{}
        output := &cloudwatchlogs.DescribeLogStreamsOutput{}

        mockClient.On("DescribeLogStreams", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLogStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetricFilters", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeMetricFiltersInput{}
        output := &cloudwatchlogs.DescribeMetricFiltersOutput{}

        mockClient.On("DescribeMetricFilters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetricFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeQueries", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeQueriesInput{}
        output := &cloudwatchlogs.DescribeQueriesOutput{}

        mockClient.On("DescribeQueries", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeQueryDefinitions", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeQueryDefinitionsInput{}
        output := &cloudwatchlogs.DescribeQueryDefinitionsOutput{}

        mockClient.On("DescribeQueryDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeQueryDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourcePolicies", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeResourcePoliciesInput{}
        output := &cloudwatchlogs.DescribeResourcePoliciesOutput{}

        mockClient.On("DescribeResourcePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourcePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSubscriptionFilters", func(t *testing.T) {
        input := &cloudwatchlogs.DescribeSubscriptionFiltersInput{}
        output := &cloudwatchlogs.DescribeSubscriptionFiltersOutput{}

        mockClient.On("DescribeSubscriptionFilters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSubscriptionFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateKmsKey", func(t *testing.T) {
        input := &cloudwatchlogs.DisassociateKmsKeyInput{}
        output := &cloudwatchlogs.DisassociateKmsKeyOutput{}

        mockClient.On("DisassociateKmsKey", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateKmsKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFilterLogEvents", func(t *testing.T) {
        input := &cloudwatchlogs.FilterLogEventsInput{}
        output := &cloudwatchlogs.FilterLogEventsOutput{}

        mockClient.On("FilterLogEvents", ctx, input).Return(output, nil)

        result, err := mockClient.FilterLogEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataProtectionPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.GetDataProtectionPolicyInput{}
        output := &cloudwatchlogs.GetDataProtectionPolicyOutput{}

        mockClient.On("GetDataProtectionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataProtectionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDelivery", func(t *testing.T) {
        input := &cloudwatchlogs.GetDeliveryInput{}
        output := &cloudwatchlogs.GetDeliveryOutput{}

        mockClient.On("GetDelivery", ctx, input).Return(output, nil)

        result, err := mockClient.GetDelivery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeliveryDestination", func(t *testing.T) {
        input := &cloudwatchlogs.GetDeliveryDestinationInput{}
        output := &cloudwatchlogs.GetDeliveryDestinationOutput{}

        mockClient.On("GetDeliveryDestination", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeliveryDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeliveryDestinationPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.GetDeliveryDestinationPolicyInput{}
        output := &cloudwatchlogs.GetDeliveryDestinationPolicyOutput{}

        mockClient.On("GetDeliveryDestinationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeliveryDestinationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeliverySource", func(t *testing.T) {
        input := &cloudwatchlogs.GetDeliverySourceInput{}
        output := &cloudwatchlogs.GetDeliverySourceOutput{}

        mockClient.On("GetDeliverySource", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeliverySource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLogAnomalyDetector", func(t *testing.T) {
        input := &cloudwatchlogs.GetLogAnomalyDetectorInput{}
        output := &cloudwatchlogs.GetLogAnomalyDetectorOutput{}

        mockClient.On("GetLogAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.GetLogAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLogEvents", func(t *testing.T) {
        input := &cloudwatchlogs.GetLogEventsInput{}
        output := &cloudwatchlogs.GetLogEventsOutput{}

        mockClient.On("GetLogEvents", ctx, input).Return(output, nil)

        result, err := mockClient.GetLogEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLogGroupFields", func(t *testing.T) {
        input := &cloudwatchlogs.GetLogGroupFieldsInput{}
        output := &cloudwatchlogs.GetLogGroupFieldsOutput{}

        mockClient.On("GetLogGroupFields", ctx, input).Return(output, nil)

        result, err := mockClient.GetLogGroupFields(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLogRecord", func(t *testing.T) {
        input := &cloudwatchlogs.GetLogRecordInput{}
        output := &cloudwatchlogs.GetLogRecordOutput{}

        mockClient.On("GetLogRecord", ctx, input).Return(output, nil)

        result, err := mockClient.GetLogRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryResults", func(t *testing.T) {
        input := &cloudwatchlogs.GetQueryResultsInput{}
        output := &cloudwatchlogs.GetQueryResultsOutput{}

        mockClient.On("GetQueryResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnomalies", func(t *testing.T) {
        input := &cloudwatchlogs.ListAnomaliesInput{}
        output := &cloudwatchlogs.ListAnomaliesOutput{}

        mockClient.On("ListAnomalies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnomalies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLogAnomalyDetectors", func(t *testing.T) {
        input := &cloudwatchlogs.ListLogAnomalyDetectorsInput{}
        output := &cloudwatchlogs.ListLogAnomalyDetectorsOutput{}

        mockClient.On("ListLogAnomalyDetectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListLogAnomalyDetectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cloudwatchlogs.ListTagsForResourceInput{}
        output := &cloudwatchlogs.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsLogGroup", func(t *testing.T) {
        input := &cloudwatchlogs.ListTagsLogGroupInput{}
        output := &cloudwatchlogs.ListTagsLogGroupOutput{}

        mockClient.On("ListTagsLogGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsLogGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.PutAccountPolicyInput{}
        output := &cloudwatchlogs.PutAccountPolicyOutput{}

        mockClient.On("PutAccountPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDataProtectionPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.PutDataProtectionPolicyInput{}
        output := &cloudwatchlogs.PutDataProtectionPolicyOutput{}

        mockClient.On("PutDataProtectionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutDataProtectionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDeliveryDestination", func(t *testing.T) {
        input := &cloudwatchlogs.PutDeliveryDestinationInput{}
        output := &cloudwatchlogs.PutDeliveryDestinationOutput{}

        mockClient.On("PutDeliveryDestination", ctx, input).Return(output, nil)

        result, err := mockClient.PutDeliveryDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDeliveryDestinationPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.PutDeliveryDestinationPolicyInput{}
        output := &cloudwatchlogs.PutDeliveryDestinationPolicyOutput{}

        mockClient.On("PutDeliveryDestinationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutDeliveryDestinationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDeliverySource", func(t *testing.T) {
        input := &cloudwatchlogs.PutDeliverySourceInput{}
        output := &cloudwatchlogs.PutDeliverySourceOutput{}

        mockClient.On("PutDeliverySource", ctx, input).Return(output, nil)

        result, err := mockClient.PutDeliverySource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDestination", func(t *testing.T) {
        input := &cloudwatchlogs.PutDestinationInput{}
        output := &cloudwatchlogs.PutDestinationOutput{}

        mockClient.On("PutDestination", ctx, input).Return(output, nil)

        result, err := mockClient.PutDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDestinationPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.PutDestinationPolicyInput{}
        output := &cloudwatchlogs.PutDestinationPolicyOutput{}

        mockClient.On("PutDestinationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutDestinationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLogEvents", func(t *testing.T) {
        input := &cloudwatchlogs.PutLogEventsInput{}
        output := &cloudwatchlogs.PutLogEventsOutput{}

        mockClient.On("PutLogEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutLogEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMetricFilter", func(t *testing.T) {
        input := &cloudwatchlogs.PutMetricFilterInput{}
        output := &cloudwatchlogs.PutMetricFilterOutput{}

        mockClient.On("PutMetricFilter", ctx, input).Return(output, nil)

        result, err := mockClient.PutMetricFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutQueryDefinition", func(t *testing.T) {
        input := &cloudwatchlogs.PutQueryDefinitionInput{}
        output := &cloudwatchlogs.PutQueryDefinitionOutput{}

        mockClient.On("PutQueryDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.PutQueryDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &cloudwatchlogs.PutResourcePolicyInput{}
        output := &cloudwatchlogs.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRetentionPolicy", func(t *testing.T) {
        input := &cloudwatchlogs.PutRetentionPolicyInput{}
        output := &cloudwatchlogs.PutRetentionPolicyOutput{}

        mockClient.On("PutRetentionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutRetentionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSubscriptionFilter", func(t *testing.T) {
        input := &cloudwatchlogs.PutSubscriptionFilterInput{}
        output := &cloudwatchlogs.PutSubscriptionFilterOutput{}

        mockClient.On("PutSubscriptionFilter", ctx, input).Return(output, nil)

        result, err := mockClient.PutSubscriptionFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartLiveTail", func(t *testing.T) {
        input := &cloudwatchlogs.StartLiveTailInput{}
        output := &cloudwatchlogs.StartLiveTailOutput{}

        mockClient.On("StartLiveTail", ctx, input).Return(output, nil)

        result, err := mockClient.StartLiveTail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartQuery", func(t *testing.T) {
        input := &cloudwatchlogs.StartQueryInput{}
        output := &cloudwatchlogs.StartQueryOutput{}

        mockClient.On("StartQuery", ctx, input).Return(output, nil)

        result, err := mockClient.StartQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopQuery", func(t *testing.T) {
        input := &cloudwatchlogs.StopQueryInput{}
        output := &cloudwatchlogs.StopQueryOutput{}

        mockClient.On("StopQuery", ctx, input).Return(output, nil)

        result, err := mockClient.StopQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagLogGroup", func(t *testing.T) {
        input := &cloudwatchlogs.TagLogGroupInput{}
        output := &cloudwatchlogs.TagLogGroupOutput{}

        mockClient.On("TagLogGroup", ctx, input).Return(output, nil)

        result, err := mockClient.TagLogGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cloudwatchlogs.TagResourceInput{}
        output := &cloudwatchlogs.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestMetricFilter", func(t *testing.T) {
        input := &cloudwatchlogs.TestMetricFilterInput{}
        output := &cloudwatchlogs.TestMetricFilterOutput{}

        mockClient.On("TestMetricFilter", ctx, input).Return(output, nil)

        result, err := mockClient.TestMetricFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagLogGroup", func(t *testing.T) {
        input := &cloudwatchlogs.UntagLogGroupInput{}
        output := &cloudwatchlogs.UntagLogGroupOutput{}

        mockClient.On("UntagLogGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UntagLogGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cloudwatchlogs.UntagResourceInput{}
        output := &cloudwatchlogs.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnomaly", func(t *testing.T) {
        input := &cloudwatchlogs.UpdateAnomalyInput{}
        output := &cloudwatchlogs.UpdateAnomalyOutput{}

        mockClient.On("UpdateAnomaly", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnomaly(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLogAnomalyDetector", func(t *testing.T) {
        input := &cloudwatchlogs.UpdateLogAnomalyDetectorInput{}
        output := &cloudwatchlogs.UpdateLogAnomalyDetectorOutput{}

        mockClient.On("UpdateLogAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLogAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

package ssmincidents_test

// tests for the ssmincidents service interface for this ../../../service/ssmincidents/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssmincidents"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssmincidents/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/ssmincidents/ssmincidents_iface"
	"github.com/stretchr/testify/assert"
)

func TestSsmincidentsServiceCanBeMocked(t *testing.T) {
	var iface ssmincidents_iface.IClient
	iface = &ssmincidents.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := ssmincidents.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetIncidentFindings", func(t *testing.T) {
        input := &ssmincidents.BatchGetIncidentFindingsInput{}
        output := &ssmincidents.BatchGetIncidentFindingsOutput{}

        mockClient.On("BatchGetIncidentFindings", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetIncidentFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicationSet", func(t *testing.T) {
        input := &ssmincidents.CreateReplicationSetInput{}
        output := &ssmincidents.CreateReplicationSetOutput{}

        mockClient.On("CreateReplicationSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResponsePlan", func(t *testing.T) {
        input := &ssmincidents.CreateResponsePlanInput{}
        output := &ssmincidents.CreateResponsePlanOutput{}

        mockClient.On("CreateResponsePlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResponsePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTimelineEvent", func(t *testing.T) {
        input := &ssmincidents.CreateTimelineEventInput{}
        output := &ssmincidents.CreateTimelineEventOutput{}

        mockClient.On("CreateTimelineEvent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTimelineEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIncidentRecord", func(t *testing.T) {
        input := &ssmincidents.DeleteIncidentRecordInput{}
        output := &ssmincidents.DeleteIncidentRecordOutput{}

        mockClient.On("DeleteIncidentRecord", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIncidentRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicationSet", func(t *testing.T) {
        input := &ssmincidents.DeleteReplicationSetInput{}
        output := &ssmincidents.DeleteReplicationSetOutput{}

        mockClient.On("DeleteReplicationSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &ssmincidents.DeleteResourcePolicyInput{}
        output := &ssmincidents.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResponsePlan", func(t *testing.T) {
        input := &ssmincidents.DeleteResponsePlanInput{}
        output := &ssmincidents.DeleteResponsePlanOutput{}

        mockClient.On("DeleteResponsePlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResponsePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTimelineEvent", func(t *testing.T) {
        input := &ssmincidents.DeleteTimelineEventInput{}
        output := &ssmincidents.DeleteTimelineEventOutput{}

        mockClient.On("DeleteTimelineEvent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTimelineEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIncidentRecord", func(t *testing.T) {
        input := &ssmincidents.GetIncidentRecordInput{}
        output := &ssmincidents.GetIncidentRecordOutput{}

        mockClient.On("GetIncidentRecord", ctx, input).Return(output, nil)

        result, err := mockClient.GetIncidentRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReplicationSet", func(t *testing.T) {
        input := &ssmincidents.GetReplicationSetInput{}
        output := &ssmincidents.GetReplicationSetOutput{}

        mockClient.On("GetReplicationSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetReplicationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicies", func(t *testing.T) {
        input := &ssmincidents.GetResourcePoliciesInput{}
        output := &ssmincidents.GetResourcePoliciesOutput{}

        mockClient.On("GetResourcePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResponsePlan", func(t *testing.T) {
        input := &ssmincidents.GetResponsePlanInput{}
        output := &ssmincidents.GetResponsePlanOutput{}

        mockClient.On("GetResponsePlan", ctx, input).Return(output, nil)

        result, err := mockClient.GetResponsePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTimelineEvent", func(t *testing.T) {
        input := &ssmincidents.GetTimelineEventInput{}
        output := &ssmincidents.GetTimelineEventOutput{}

        mockClient.On("GetTimelineEvent", ctx, input).Return(output, nil)

        result, err := mockClient.GetTimelineEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIncidentFindings", func(t *testing.T) {
        input := &ssmincidents.ListIncidentFindingsInput{}
        output := &ssmincidents.ListIncidentFindingsOutput{}

        mockClient.On("ListIncidentFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListIncidentFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIncidentRecords", func(t *testing.T) {
        input := &ssmincidents.ListIncidentRecordsInput{}
        output := &ssmincidents.ListIncidentRecordsOutput{}

        mockClient.On("ListIncidentRecords", ctx, input).Return(output, nil)

        result, err := mockClient.ListIncidentRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRelatedItems", func(t *testing.T) {
        input := &ssmincidents.ListRelatedItemsInput{}
        output := &ssmincidents.ListRelatedItemsOutput{}

        mockClient.On("ListRelatedItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListRelatedItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReplicationSets", func(t *testing.T) {
        input := &ssmincidents.ListReplicationSetsInput{}
        output := &ssmincidents.ListReplicationSetsOutput{}

        mockClient.On("ListReplicationSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListReplicationSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResponsePlans", func(t *testing.T) {
        input := &ssmincidents.ListResponsePlansInput{}
        output := &ssmincidents.ListResponsePlansOutput{}

        mockClient.On("ListResponsePlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListResponsePlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &ssmincidents.ListTagsForResourceInput{}
        output := &ssmincidents.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTimelineEvents", func(t *testing.T) {
        input := &ssmincidents.ListTimelineEventsInput{}
        output := &ssmincidents.ListTimelineEventsOutput{}

        mockClient.On("ListTimelineEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListTimelineEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &ssmincidents.PutResourcePolicyInput{}
        output := &ssmincidents.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartIncident", func(t *testing.T) {
        input := &ssmincidents.StartIncidentInput{}
        output := &ssmincidents.StartIncidentOutput{}

        mockClient.On("StartIncident", ctx, input).Return(output, nil)

        result, err := mockClient.StartIncident(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &ssmincidents.TagResourceInput{}
        output := &ssmincidents.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &ssmincidents.UntagResourceInput{}
        output := &ssmincidents.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeletionProtection", func(t *testing.T) {
        input := &ssmincidents.UpdateDeletionProtectionInput{}
        output := &ssmincidents.UpdateDeletionProtectionOutput{}

        mockClient.On("UpdateDeletionProtection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeletionProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIncidentRecord", func(t *testing.T) {
        input := &ssmincidents.UpdateIncidentRecordInput{}
        output := &ssmincidents.UpdateIncidentRecordOutput{}

        mockClient.On("UpdateIncidentRecord", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIncidentRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRelatedItems", func(t *testing.T) {
        input := &ssmincidents.UpdateRelatedItemsInput{}
        output := &ssmincidents.UpdateRelatedItemsOutput{}

        mockClient.On("UpdateRelatedItems", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRelatedItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReplicationSet", func(t *testing.T) {
        input := &ssmincidents.UpdateReplicationSetInput{}
        output := &ssmincidents.UpdateReplicationSetOutput{}

        mockClient.On("UpdateReplicationSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReplicationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResponsePlan", func(t *testing.T) {
        input := &ssmincidents.UpdateResponsePlanInput{}
        output := &ssmincidents.UpdateResponsePlanOutput{}

        mockClient.On("UpdateResponsePlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResponsePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTimelineEvent", func(t *testing.T) {
        input := &ssmincidents.UpdateTimelineEventInput{}
        output := &ssmincidents.UpdateTimelineEventOutput{}

        mockClient.On("UpdateTimelineEvent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTimelineEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

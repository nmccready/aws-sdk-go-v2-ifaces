package migrationhub_test

// tests for the migrationhub service interface for this ../../../service/migrationhub/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/migrationhub"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhub/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhub/migrationhub_iface"
	"github.com/stretchr/testify/assert"
)

func TestMigrationhubServiceCanBeMocked(t *testing.T) {
	var iface migrationhub_iface.IClient
	iface = &migrationhub.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := migrationhub.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateCreatedArtifact", func(t *testing.T) {
        input := &migrationhub.AssociateCreatedArtifactInput{}
        output := &migrationhub.AssociateCreatedArtifactOutput{}

        mockClient.On("AssociateCreatedArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateCreatedArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDiscoveredResource", func(t *testing.T) {
        input := &migrationhub.AssociateDiscoveredResourceInput{}
        output := &migrationhub.AssociateDiscoveredResourceOutput{}

        mockClient.On("AssociateDiscoveredResource", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDiscoveredResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProgressUpdateStream", func(t *testing.T) {
        input := &migrationhub.CreateProgressUpdateStreamInput{}
        output := &migrationhub.CreateProgressUpdateStreamOutput{}

        mockClient.On("CreateProgressUpdateStream", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProgressUpdateStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProgressUpdateStream", func(t *testing.T) {
        input := &migrationhub.DeleteProgressUpdateStreamInput{}
        output := &migrationhub.DeleteProgressUpdateStreamOutput{}

        mockClient.On("DeleteProgressUpdateStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProgressUpdateStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationState", func(t *testing.T) {
        input := &migrationhub.DescribeApplicationStateInput{}
        output := &migrationhub.DescribeApplicationStateOutput{}

        mockClient.On("DescribeApplicationState", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMigrationTask", func(t *testing.T) {
        input := &migrationhub.DescribeMigrationTaskInput{}
        output := &migrationhub.DescribeMigrationTaskOutput{}

        mockClient.On("DescribeMigrationTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMigrationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateCreatedArtifact", func(t *testing.T) {
        input := &migrationhub.DisassociateCreatedArtifactInput{}
        output := &migrationhub.DisassociateCreatedArtifactOutput{}

        mockClient.On("DisassociateCreatedArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateCreatedArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDiscoveredResource", func(t *testing.T) {
        input := &migrationhub.DisassociateDiscoveredResourceInput{}
        output := &migrationhub.DisassociateDiscoveredResourceOutput{}

        mockClient.On("DisassociateDiscoveredResource", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDiscoveredResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportMigrationTask", func(t *testing.T) {
        input := &migrationhub.ImportMigrationTaskInput{}
        output := &migrationhub.ImportMigrationTaskOutput{}

        mockClient.On("ImportMigrationTask", ctx, input).Return(output, nil)

        result, err := mockClient.ImportMigrationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationStates", func(t *testing.T) {
        input := &migrationhub.ListApplicationStatesInput{}
        output := &migrationhub.ListApplicationStatesOutput{}

        mockClient.On("ListApplicationStates", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationStates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCreatedArtifacts", func(t *testing.T) {
        input := &migrationhub.ListCreatedArtifactsInput{}
        output := &migrationhub.ListCreatedArtifactsOutput{}

        mockClient.On("ListCreatedArtifacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListCreatedArtifacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDiscoveredResources", func(t *testing.T) {
        input := &migrationhub.ListDiscoveredResourcesInput{}
        output := &migrationhub.ListDiscoveredResourcesOutput{}

        mockClient.On("ListDiscoveredResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDiscoveredResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMigrationTasks", func(t *testing.T) {
        input := &migrationhub.ListMigrationTasksInput{}
        output := &migrationhub.ListMigrationTasksOutput{}

        mockClient.On("ListMigrationTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListMigrationTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProgressUpdateStreams", func(t *testing.T) {
        input := &migrationhub.ListProgressUpdateStreamsInput{}
        output := &migrationhub.ListProgressUpdateStreamsOutput{}

        mockClient.On("ListProgressUpdateStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListProgressUpdateStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyApplicationState", func(t *testing.T) {
        input := &migrationhub.NotifyApplicationStateInput{}
        output := &migrationhub.NotifyApplicationStateOutput{}

        mockClient.On("NotifyApplicationState", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyApplicationState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyMigrationTaskState", func(t *testing.T) {
        input := &migrationhub.NotifyMigrationTaskStateInput{}
        output := &migrationhub.NotifyMigrationTaskStateOutput{}

        mockClient.On("NotifyMigrationTaskState", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyMigrationTaskState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourceAttributes", func(t *testing.T) {
        input := &migrationhub.PutResourceAttributesInput{}
        output := &migrationhub.PutResourceAttributesOutput{}

        mockClient.On("PutResourceAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourceAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

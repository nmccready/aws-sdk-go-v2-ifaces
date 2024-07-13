package pcaconnectorscep_test

// tests for the pcaconnectorscep service interface for this ../../../service/pcaconnectorscep/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pcaconnectorscep/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pcaconnectorscep/pcaconnectorscep_iface"
	"github.com/stretchr/testify/assert"
)

func TestPcaconnectorscepServiceCanBeMocked(t *testing.T) {
	var iface pcaconnectorscep_iface.IClient
	iface = &pcaconnectorscep.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pcaconnectorscep.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChallenge", func(t *testing.T) {
        input := &pcaconnectorscep.CreateChallengeInput{}
        output := &pcaconnectorscep.CreateChallengeOutput{}

        mockClient.On("CreateChallenge", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChallenge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnector", func(t *testing.T) {
        input := &pcaconnectorscep.CreateConnectorInput{}
        output := &pcaconnectorscep.CreateConnectorOutput{}

        mockClient.On("CreateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChallenge", func(t *testing.T) {
        input := &pcaconnectorscep.DeleteChallengeInput{}
        output := &pcaconnectorscep.DeleteChallengeOutput{}

        mockClient.On("DeleteChallenge", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChallenge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnector", func(t *testing.T) {
        input := &pcaconnectorscep.DeleteConnectorInput{}
        output := &pcaconnectorscep.DeleteConnectorOutput{}

        mockClient.On("DeleteConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChallengeMetadata", func(t *testing.T) {
        input := &pcaconnectorscep.GetChallengeMetadataInput{}
        output := &pcaconnectorscep.GetChallengeMetadataOutput{}

        mockClient.On("GetChallengeMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetChallengeMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChallengePassword", func(t *testing.T) {
        input := &pcaconnectorscep.GetChallengePasswordInput{}
        output := &pcaconnectorscep.GetChallengePasswordOutput{}

        mockClient.On("GetChallengePassword", ctx, input).Return(output, nil)

        result, err := mockClient.GetChallengePassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnector", func(t *testing.T) {
        input := &pcaconnectorscep.GetConnectorInput{}
        output := &pcaconnectorscep.GetConnectorOutput{}

        mockClient.On("GetConnector", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChallengeMetadata", func(t *testing.T) {
        input := &pcaconnectorscep.ListChallengeMetadataInput{}
        output := &pcaconnectorscep.ListChallengeMetadataOutput{}

        mockClient.On("ListChallengeMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.ListChallengeMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectors", func(t *testing.T) {
        input := &pcaconnectorscep.ListConnectorsInput{}
        output := &pcaconnectorscep.ListConnectorsOutput{}

        mockClient.On("ListConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &pcaconnectorscep.ListTagsForResourceInput{}
        output := &pcaconnectorscep.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &pcaconnectorscep.TagResourceInput{}
        output := &pcaconnectorscep.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &pcaconnectorscep.UntagResourceInput{}
        output := &pcaconnectorscep.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

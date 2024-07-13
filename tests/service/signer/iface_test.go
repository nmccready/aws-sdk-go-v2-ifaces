package signer_test

// tests for the signer service interface for this ../../../service/signer/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/signer/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/signer/signer_iface"
	"github.com/stretchr/testify/assert"
)

func TestSignerServiceCanBeMocked(t *testing.T) {
	var iface signer_iface.IClient
	iface = &signer.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := signer.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddProfilePermission", func(t *testing.T) {
        input := &signer.AddProfilePermissionInput{}
        output := &signer.AddProfilePermissionOutput{}

        mockClient.On("AddProfilePermission", ctx, input).Return(output, nil)

        result, err := mockClient.AddProfilePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSigningProfile", func(t *testing.T) {
        input := &signer.CancelSigningProfileInput{}
        output := &signer.CancelSigningProfileOutput{}

        mockClient.On("CancelSigningProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSigningProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSigningJob", func(t *testing.T) {
        input := &signer.DescribeSigningJobInput{}
        output := &signer.DescribeSigningJobOutput{}

        mockClient.On("DescribeSigningJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSigningJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRevocationStatus", func(t *testing.T) {
        input := &signer.GetRevocationStatusInput{}
        output := &signer.GetRevocationStatusOutput{}

        mockClient.On("GetRevocationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetRevocationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSigningPlatform", func(t *testing.T) {
        input := &signer.GetSigningPlatformInput{}
        output := &signer.GetSigningPlatformOutput{}

        mockClient.On("GetSigningPlatform", ctx, input).Return(output, nil)

        result, err := mockClient.GetSigningPlatform(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSigningProfile", func(t *testing.T) {
        input := &signer.GetSigningProfileInput{}
        output := &signer.GetSigningProfileOutput{}

        mockClient.On("GetSigningProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetSigningProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfilePermissions", func(t *testing.T) {
        input := &signer.ListProfilePermissionsInput{}
        output := &signer.ListProfilePermissionsOutput{}

        mockClient.On("ListProfilePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfilePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSigningJobs", func(t *testing.T) {
        input := &signer.ListSigningJobsInput{}
        output := &signer.ListSigningJobsOutput{}

        mockClient.On("ListSigningJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSigningJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSigningPlatforms", func(t *testing.T) {
        input := &signer.ListSigningPlatformsInput{}
        output := &signer.ListSigningPlatformsOutput{}

        mockClient.On("ListSigningPlatforms", ctx, input).Return(output, nil)

        result, err := mockClient.ListSigningPlatforms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSigningProfiles", func(t *testing.T) {
        input := &signer.ListSigningProfilesInput{}
        output := &signer.ListSigningProfilesOutput{}

        mockClient.On("ListSigningProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListSigningProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &signer.ListTagsForResourceInput{}
        output := &signer.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSigningProfile", func(t *testing.T) {
        input := &signer.PutSigningProfileInput{}
        output := &signer.PutSigningProfileOutput{}

        mockClient.On("PutSigningProfile", ctx, input).Return(output, nil)

        result, err := mockClient.PutSigningProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveProfilePermission", func(t *testing.T) {
        input := &signer.RemoveProfilePermissionInput{}
        output := &signer.RemoveProfilePermissionOutput{}

        mockClient.On("RemoveProfilePermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveProfilePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeSignature", func(t *testing.T) {
        input := &signer.RevokeSignatureInput{}
        output := &signer.RevokeSignatureOutput{}

        mockClient.On("RevokeSignature", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeSignature(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeSigningProfile", func(t *testing.T) {
        input := &signer.RevokeSigningProfileInput{}
        output := &signer.RevokeSigningProfileOutput{}

        mockClient.On("RevokeSigningProfile", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeSigningProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSignPayload", func(t *testing.T) {
        input := &signer.SignPayloadInput{}
        output := &signer.SignPayloadOutput{}

        mockClient.On("SignPayload", ctx, input).Return(output, nil)

        result, err := mockClient.SignPayload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSigningJob", func(t *testing.T) {
        input := &signer.StartSigningJobInput{}
        output := &signer.StartSigningJobOutput{}

        mockClient.On("StartSigningJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartSigningJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &signer.TagResourceInput{}
        output := &signer.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &signer.UntagResourceInput{}
        output := &signer.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

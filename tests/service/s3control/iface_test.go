package s3control_test

// tests for the s3control service interface for this ../../../service/s3control/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3control/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3control/s3control_iface"
	"github.com/stretchr/testify/assert"
)

func TestS3controlServiceCanBeMocked(t *testing.T) {
	var iface s3control_iface.IClient
	iface = &s3control.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := s3control.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAccessGrantsIdentityCenter", func(t *testing.T) {
        input := &s3control.AssociateAccessGrantsIdentityCenterInput{}
        output := &s3control.AssociateAccessGrantsIdentityCenterOutput{}

        mockClient.On("AssociateAccessGrantsIdentityCenter", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAccessGrantsIdentityCenter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessGrant", func(t *testing.T) {
        input := &s3control.CreateAccessGrantInput{}
        output := &s3control.CreateAccessGrantOutput{}

        mockClient.On("CreateAccessGrant", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessGrantsInstance", func(t *testing.T) {
        input := &s3control.CreateAccessGrantsInstanceInput{}
        output := &s3control.CreateAccessGrantsInstanceOutput{}

        mockClient.On("CreateAccessGrantsInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessGrantsInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessGrantsLocation", func(t *testing.T) {
        input := &s3control.CreateAccessGrantsLocationInput{}
        output := &s3control.CreateAccessGrantsLocationOutput{}

        mockClient.On("CreateAccessGrantsLocation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessGrantsLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessPoint", func(t *testing.T) {
        input := &s3control.CreateAccessPointInput{}
        output := &s3control.CreateAccessPointOutput{}

        mockClient.On("CreateAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessPointForObjectLambda", func(t *testing.T) {
        input := &s3control.CreateAccessPointForObjectLambdaInput{}
        output := &s3control.CreateAccessPointForObjectLambdaOutput{}

        mockClient.On("CreateAccessPointForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessPointForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBucket", func(t *testing.T) {
        input := &s3control.CreateBucketInput{}
        output := &s3control.CreateBucketOutput{}

        mockClient.On("CreateBucket", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &s3control.CreateJobInput{}
        output := &s3control.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMultiRegionAccessPoint", func(t *testing.T) {
        input := &s3control.CreateMultiRegionAccessPointInput{}
        output := &s3control.CreateMultiRegionAccessPointOutput{}

        mockClient.On("CreateMultiRegionAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMultiRegionAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStorageLensGroup", func(t *testing.T) {
        input := &s3control.CreateStorageLensGroupInput{}
        output := &s3control.CreateStorageLensGroupOutput{}

        mockClient.On("CreateStorageLensGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStorageLensGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessGrant", func(t *testing.T) {
        input := &s3control.DeleteAccessGrantInput{}
        output := &s3control.DeleteAccessGrantOutput{}

        mockClient.On("DeleteAccessGrant", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessGrantsInstance", func(t *testing.T) {
        input := &s3control.DeleteAccessGrantsInstanceInput{}
        output := &s3control.DeleteAccessGrantsInstanceOutput{}

        mockClient.On("DeleteAccessGrantsInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessGrantsInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessGrantsInstanceResourcePolicy", func(t *testing.T) {
        input := &s3control.DeleteAccessGrantsInstanceResourcePolicyInput{}
        output := &s3control.DeleteAccessGrantsInstanceResourcePolicyOutput{}

        mockClient.On("DeleteAccessGrantsInstanceResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessGrantsInstanceResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessGrantsLocation", func(t *testing.T) {
        input := &s3control.DeleteAccessGrantsLocationInput{}
        output := &s3control.DeleteAccessGrantsLocationOutput{}

        mockClient.On("DeleteAccessGrantsLocation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessGrantsLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessPoint", func(t *testing.T) {
        input := &s3control.DeleteAccessPointInput{}
        output := &s3control.DeleteAccessPointOutput{}

        mockClient.On("DeleteAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessPointForObjectLambda", func(t *testing.T) {
        input := &s3control.DeleteAccessPointForObjectLambdaInput{}
        output := &s3control.DeleteAccessPointForObjectLambdaOutput{}

        mockClient.On("DeleteAccessPointForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessPointForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessPointPolicy", func(t *testing.T) {
        input := &s3control.DeleteAccessPointPolicyInput{}
        output := &s3control.DeleteAccessPointPolicyOutput{}

        mockClient.On("DeleteAccessPointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessPointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessPointPolicyForObjectLambda", func(t *testing.T) {
        input := &s3control.DeleteAccessPointPolicyForObjectLambdaInput{}
        output := &s3control.DeleteAccessPointPolicyForObjectLambdaOutput{}

        mockClient.On("DeleteAccessPointPolicyForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessPointPolicyForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucket", func(t *testing.T) {
        input := &s3control.DeleteBucketInput{}
        output := &s3control.DeleteBucketOutput{}

        mockClient.On("DeleteBucket", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketLifecycleConfiguration", func(t *testing.T) {
        input := &s3control.DeleteBucketLifecycleConfigurationInput{}
        output := &s3control.DeleteBucketLifecycleConfigurationOutput{}

        mockClient.On("DeleteBucketLifecycleConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketLifecycleConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketPolicy", func(t *testing.T) {
        input := &s3control.DeleteBucketPolicyInput{}
        output := &s3control.DeleteBucketPolicyOutput{}

        mockClient.On("DeleteBucketPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketReplication", func(t *testing.T) {
        input := &s3control.DeleteBucketReplicationInput{}
        output := &s3control.DeleteBucketReplicationOutput{}

        mockClient.On("DeleteBucketReplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketTagging", func(t *testing.T) {
        input := &s3control.DeleteBucketTaggingInput{}
        output := &s3control.DeleteBucketTaggingOutput{}

        mockClient.On("DeleteBucketTagging", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJobTagging", func(t *testing.T) {
        input := &s3control.DeleteJobTaggingInput{}
        output := &s3control.DeleteJobTaggingOutput{}

        mockClient.On("DeleteJobTagging", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJobTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMultiRegionAccessPoint", func(t *testing.T) {
        input := &s3control.DeleteMultiRegionAccessPointInput{}
        output := &s3control.DeleteMultiRegionAccessPointOutput{}

        mockClient.On("DeleteMultiRegionAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMultiRegionAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePublicAccessBlock", func(t *testing.T) {
        input := &s3control.DeletePublicAccessBlockInput{}
        output := &s3control.DeletePublicAccessBlockOutput{}

        mockClient.On("DeletePublicAccessBlock", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePublicAccessBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStorageLensConfiguration", func(t *testing.T) {
        input := &s3control.DeleteStorageLensConfigurationInput{}
        output := &s3control.DeleteStorageLensConfigurationOutput{}

        mockClient.On("DeleteStorageLensConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStorageLensConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStorageLensConfigurationTagging", func(t *testing.T) {
        input := &s3control.DeleteStorageLensConfigurationTaggingInput{}
        output := &s3control.DeleteStorageLensConfigurationTaggingOutput{}

        mockClient.On("DeleteStorageLensConfigurationTagging", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStorageLensConfigurationTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStorageLensGroup", func(t *testing.T) {
        input := &s3control.DeleteStorageLensGroupInput{}
        output := &s3control.DeleteStorageLensGroupOutput{}

        mockClient.On("DeleteStorageLensGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStorageLensGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJob", func(t *testing.T) {
        input := &s3control.DescribeJobInput{}
        output := &s3control.DescribeJobOutput{}

        mockClient.On("DescribeJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMultiRegionAccessPointOperation", func(t *testing.T) {
        input := &s3control.DescribeMultiRegionAccessPointOperationInput{}
        output := &s3control.DescribeMultiRegionAccessPointOperationOutput{}

        mockClient.On("DescribeMultiRegionAccessPointOperation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMultiRegionAccessPointOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDissociateAccessGrantsIdentityCenter", func(t *testing.T) {
        input := &s3control.DissociateAccessGrantsIdentityCenterInput{}
        output := &s3control.DissociateAccessGrantsIdentityCenterOutput{}

        mockClient.On("DissociateAccessGrantsIdentityCenter", ctx, input).Return(output, nil)

        result, err := mockClient.DissociateAccessGrantsIdentityCenter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessGrant", func(t *testing.T) {
        input := &s3control.GetAccessGrantInput{}
        output := &s3control.GetAccessGrantOutput{}

        mockClient.On("GetAccessGrant", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessGrantsInstance", func(t *testing.T) {
        input := &s3control.GetAccessGrantsInstanceInput{}
        output := &s3control.GetAccessGrantsInstanceOutput{}

        mockClient.On("GetAccessGrantsInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessGrantsInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessGrantsInstanceForPrefix", func(t *testing.T) {
        input := &s3control.GetAccessGrantsInstanceForPrefixInput{}
        output := &s3control.GetAccessGrantsInstanceForPrefixOutput{}

        mockClient.On("GetAccessGrantsInstanceForPrefix", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessGrantsInstanceForPrefix(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessGrantsInstanceResourcePolicy", func(t *testing.T) {
        input := &s3control.GetAccessGrantsInstanceResourcePolicyInput{}
        output := &s3control.GetAccessGrantsInstanceResourcePolicyOutput{}

        mockClient.On("GetAccessGrantsInstanceResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessGrantsInstanceResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessGrantsLocation", func(t *testing.T) {
        input := &s3control.GetAccessGrantsLocationInput{}
        output := &s3control.GetAccessGrantsLocationOutput{}

        mockClient.On("GetAccessGrantsLocation", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessGrantsLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPoint", func(t *testing.T) {
        input := &s3control.GetAccessPointInput{}
        output := &s3control.GetAccessPointOutput{}

        mockClient.On("GetAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPointConfigurationForObjectLambda", func(t *testing.T) {
        input := &s3control.GetAccessPointConfigurationForObjectLambdaInput{}
        output := &s3control.GetAccessPointConfigurationForObjectLambdaOutput{}

        mockClient.On("GetAccessPointConfigurationForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPointConfigurationForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPointForObjectLambda", func(t *testing.T) {
        input := &s3control.GetAccessPointForObjectLambdaInput{}
        output := &s3control.GetAccessPointForObjectLambdaOutput{}

        mockClient.On("GetAccessPointForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPointForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPointPolicy", func(t *testing.T) {
        input := &s3control.GetAccessPointPolicyInput{}
        output := &s3control.GetAccessPointPolicyOutput{}

        mockClient.On("GetAccessPointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPointPolicyForObjectLambda", func(t *testing.T) {
        input := &s3control.GetAccessPointPolicyForObjectLambdaInput{}
        output := &s3control.GetAccessPointPolicyForObjectLambdaOutput{}

        mockClient.On("GetAccessPointPolicyForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPointPolicyForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPointPolicyStatus", func(t *testing.T) {
        input := &s3control.GetAccessPointPolicyStatusInput{}
        output := &s3control.GetAccessPointPolicyStatusOutput{}

        mockClient.On("GetAccessPointPolicyStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPointPolicyStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPointPolicyStatusForObjectLambda", func(t *testing.T) {
        input := &s3control.GetAccessPointPolicyStatusForObjectLambdaInput{}
        output := &s3control.GetAccessPointPolicyStatusForObjectLambdaOutput{}

        mockClient.On("GetAccessPointPolicyStatusForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPointPolicyStatusForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucket", func(t *testing.T) {
        input := &s3control.GetBucketInput{}
        output := &s3control.GetBucketOutput{}

        mockClient.On("GetBucket", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketLifecycleConfiguration", func(t *testing.T) {
        input := &s3control.GetBucketLifecycleConfigurationInput{}
        output := &s3control.GetBucketLifecycleConfigurationOutput{}

        mockClient.On("GetBucketLifecycleConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketLifecycleConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketPolicy", func(t *testing.T) {
        input := &s3control.GetBucketPolicyInput{}
        output := &s3control.GetBucketPolicyOutput{}

        mockClient.On("GetBucketPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketReplication", func(t *testing.T) {
        input := &s3control.GetBucketReplicationInput{}
        output := &s3control.GetBucketReplicationOutput{}

        mockClient.On("GetBucketReplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketTagging", func(t *testing.T) {
        input := &s3control.GetBucketTaggingInput{}
        output := &s3control.GetBucketTaggingOutput{}

        mockClient.On("GetBucketTagging", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketVersioning", func(t *testing.T) {
        input := &s3control.GetBucketVersioningInput{}
        output := &s3control.GetBucketVersioningOutput{}

        mockClient.On("GetBucketVersioning", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketVersioning(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataAccess", func(t *testing.T) {
        input := &s3control.GetDataAccessInput{}
        output := &s3control.GetDataAccessOutput{}

        mockClient.On("GetDataAccess", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobTagging", func(t *testing.T) {
        input := &s3control.GetJobTaggingInput{}
        output := &s3control.GetJobTaggingOutput{}

        mockClient.On("GetJobTagging", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMultiRegionAccessPoint", func(t *testing.T) {
        input := &s3control.GetMultiRegionAccessPointInput{}
        output := &s3control.GetMultiRegionAccessPointOutput{}

        mockClient.On("GetMultiRegionAccessPoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetMultiRegionAccessPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMultiRegionAccessPointPolicy", func(t *testing.T) {
        input := &s3control.GetMultiRegionAccessPointPolicyInput{}
        output := &s3control.GetMultiRegionAccessPointPolicyOutput{}

        mockClient.On("GetMultiRegionAccessPointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetMultiRegionAccessPointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMultiRegionAccessPointPolicyStatus", func(t *testing.T) {
        input := &s3control.GetMultiRegionAccessPointPolicyStatusInput{}
        output := &s3control.GetMultiRegionAccessPointPolicyStatusOutput{}

        mockClient.On("GetMultiRegionAccessPointPolicyStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetMultiRegionAccessPointPolicyStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMultiRegionAccessPointRoutes", func(t *testing.T) {
        input := &s3control.GetMultiRegionAccessPointRoutesInput{}
        output := &s3control.GetMultiRegionAccessPointRoutesOutput{}

        mockClient.On("GetMultiRegionAccessPointRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.GetMultiRegionAccessPointRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPublicAccessBlock", func(t *testing.T) {
        input := &s3control.GetPublicAccessBlockInput{}
        output := &s3control.GetPublicAccessBlockOutput{}

        mockClient.On("GetPublicAccessBlock", ctx, input).Return(output, nil)

        result, err := mockClient.GetPublicAccessBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStorageLensConfiguration", func(t *testing.T) {
        input := &s3control.GetStorageLensConfigurationInput{}
        output := &s3control.GetStorageLensConfigurationOutput{}

        mockClient.On("GetStorageLensConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetStorageLensConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStorageLensConfigurationTagging", func(t *testing.T) {
        input := &s3control.GetStorageLensConfigurationTaggingInput{}
        output := &s3control.GetStorageLensConfigurationTaggingOutput{}

        mockClient.On("GetStorageLensConfigurationTagging", ctx, input).Return(output, nil)

        result, err := mockClient.GetStorageLensConfigurationTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStorageLensGroup", func(t *testing.T) {
        input := &s3control.GetStorageLensGroupInput{}
        output := &s3control.GetStorageLensGroupOutput{}

        mockClient.On("GetStorageLensGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetStorageLensGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessGrants", func(t *testing.T) {
        input := &s3control.ListAccessGrantsInput{}
        output := &s3control.ListAccessGrantsOutput{}

        mockClient.On("ListAccessGrants", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessGrants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessGrantsInstances", func(t *testing.T) {
        input := &s3control.ListAccessGrantsInstancesInput{}
        output := &s3control.ListAccessGrantsInstancesOutput{}

        mockClient.On("ListAccessGrantsInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessGrantsInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessGrantsLocations", func(t *testing.T) {
        input := &s3control.ListAccessGrantsLocationsInput{}
        output := &s3control.ListAccessGrantsLocationsOutput{}

        mockClient.On("ListAccessGrantsLocations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessGrantsLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessPoints", func(t *testing.T) {
        input := &s3control.ListAccessPointsInput{}
        output := &s3control.ListAccessPointsOutput{}

        mockClient.On("ListAccessPoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessPointsForObjectLambda", func(t *testing.T) {
        input := &s3control.ListAccessPointsForObjectLambdaInput{}
        output := &s3control.ListAccessPointsForObjectLambdaOutput{}

        mockClient.On("ListAccessPointsForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessPointsForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &s3control.ListJobsInput{}
        output := &s3control.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMultiRegionAccessPoints", func(t *testing.T) {
        input := &s3control.ListMultiRegionAccessPointsInput{}
        output := &s3control.ListMultiRegionAccessPointsOutput{}

        mockClient.On("ListMultiRegionAccessPoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListMultiRegionAccessPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegionalBuckets", func(t *testing.T) {
        input := &s3control.ListRegionalBucketsInput{}
        output := &s3control.ListRegionalBucketsOutput{}

        mockClient.On("ListRegionalBuckets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegionalBuckets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStorageLensConfigurations", func(t *testing.T) {
        input := &s3control.ListStorageLensConfigurationsInput{}
        output := &s3control.ListStorageLensConfigurationsOutput{}

        mockClient.On("ListStorageLensConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListStorageLensConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStorageLensGroups", func(t *testing.T) {
        input := &s3control.ListStorageLensGroupsInput{}
        output := &s3control.ListStorageLensGroupsOutput{}

        mockClient.On("ListStorageLensGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListStorageLensGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &s3control.ListTagsForResourceInput{}
        output := &s3control.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccessGrantsInstanceResourcePolicy", func(t *testing.T) {
        input := &s3control.PutAccessGrantsInstanceResourcePolicyInput{}
        output := &s3control.PutAccessGrantsInstanceResourcePolicyOutput{}

        mockClient.On("PutAccessGrantsInstanceResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccessGrantsInstanceResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccessPointConfigurationForObjectLambda", func(t *testing.T) {
        input := &s3control.PutAccessPointConfigurationForObjectLambdaInput{}
        output := &s3control.PutAccessPointConfigurationForObjectLambdaOutput{}

        mockClient.On("PutAccessPointConfigurationForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccessPointConfigurationForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccessPointPolicy", func(t *testing.T) {
        input := &s3control.PutAccessPointPolicyInput{}
        output := &s3control.PutAccessPointPolicyOutput{}

        mockClient.On("PutAccessPointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccessPointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccessPointPolicyForObjectLambda", func(t *testing.T) {
        input := &s3control.PutAccessPointPolicyForObjectLambdaInput{}
        output := &s3control.PutAccessPointPolicyForObjectLambdaOutput{}

        mockClient.On("PutAccessPointPolicyForObjectLambda", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccessPointPolicyForObjectLambda(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketLifecycleConfiguration", func(t *testing.T) {
        input := &s3control.PutBucketLifecycleConfigurationInput{}
        output := &s3control.PutBucketLifecycleConfigurationOutput{}

        mockClient.On("PutBucketLifecycleConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketLifecycleConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketPolicy", func(t *testing.T) {
        input := &s3control.PutBucketPolicyInput{}
        output := &s3control.PutBucketPolicyOutput{}

        mockClient.On("PutBucketPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketReplication", func(t *testing.T) {
        input := &s3control.PutBucketReplicationInput{}
        output := &s3control.PutBucketReplicationOutput{}

        mockClient.On("PutBucketReplication", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketTagging", func(t *testing.T) {
        input := &s3control.PutBucketTaggingInput{}
        output := &s3control.PutBucketTaggingOutput{}

        mockClient.On("PutBucketTagging", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBucketVersioning", func(t *testing.T) {
        input := &s3control.PutBucketVersioningInput{}
        output := &s3control.PutBucketVersioningOutput{}

        mockClient.On("PutBucketVersioning", ctx, input).Return(output, nil)

        result, err := mockClient.PutBucketVersioning(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutJobTagging", func(t *testing.T) {
        input := &s3control.PutJobTaggingInput{}
        output := &s3control.PutJobTaggingOutput{}

        mockClient.On("PutJobTagging", ctx, input).Return(output, nil)

        result, err := mockClient.PutJobTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMultiRegionAccessPointPolicy", func(t *testing.T) {
        input := &s3control.PutMultiRegionAccessPointPolicyInput{}
        output := &s3control.PutMultiRegionAccessPointPolicyOutput{}

        mockClient.On("PutMultiRegionAccessPointPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutMultiRegionAccessPointPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPublicAccessBlock", func(t *testing.T) {
        input := &s3control.PutPublicAccessBlockInput{}
        output := &s3control.PutPublicAccessBlockOutput{}

        mockClient.On("PutPublicAccessBlock", ctx, input).Return(output, nil)

        result, err := mockClient.PutPublicAccessBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutStorageLensConfiguration", func(t *testing.T) {
        input := &s3control.PutStorageLensConfigurationInput{}
        output := &s3control.PutStorageLensConfigurationOutput{}

        mockClient.On("PutStorageLensConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutStorageLensConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutStorageLensConfigurationTagging", func(t *testing.T) {
        input := &s3control.PutStorageLensConfigurationTaggingInput{}
        output := &s3control.PutStorageLensConfigurationTaggingOutput{}

        mockClient.On("PutStorageLensConfigurationTagging", ctx, input).Return(output, nil)

        result, err := mockClient.PutStorageLensConfigurationTagging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitMultiRegionAccessPointRoutes", func(t *testing.T) {
        input := &s3control.SubmitMultiRegionAccessPointRoutesInput{}
        output := &s3control.SubmitMultiRegionAccessPointRoutesOutput{}

        mockClient.On("SubmitMultiRegionAccessPointRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitMultiRegionAccessPointRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &s3control.TagResourceInput{}
        output := &s3control.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &s3control.UntagResourceInput{}
        output := &s3control.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccessGrantsLocation", func(t *testing.T) {
        input := &s3control.UpdateAccessGrantsLocationInput{}
        output := &s3control.UpdateAccessGrantsLocationOutput{}

        mockClient.On("UpdateAccessGrantsLocation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccessGrantsLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJobPriority", func(t *testing.T) {
        input := &s3control.UpdateJobPriorityInput{}
        output := &s3control.UpdateJobPriorityOutput{}

        mockClient.On("UpdateJobPriority", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJobPriority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJobStatus", func(t *testing.T) {
        input := &s3control.UpdateJobStatusInput{}
        output := &s3control.UpdateJobStatusOutput{}

        mockClient.On("UpdateJobStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJobStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStorageLensGroup", func(t *testing.T) {
        input := &s3control.UpdateStorageLensGroupInput{}
        output := &s3control.UpdateStorageLensGroupOutput{}

        mockClient.On("UpdateStorageLensGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStorageLensGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

package cognitoidentity_test

// tests for the cognitoidentity service interface for this ../../../service/cognitoidentity/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cognitoidentity/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cognitoidentity/cognitoidentity_iface"
	"github.com/stretchr/testify/assert"
)

func TestCognitoidentityServiceCanBeMocked(t *testing.T) {
	var iface cognitoidentity_iface.IClient
	iface = &cognitoidentity.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cognitoidentity.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIdentityPool", func(t *testing.T) {
        input := &cognitoidentity.CreateIdentityPoolInput{}
        output := &cognitoidentity.CreateIdentityPoolOutput{}

        mockClient.On("CreateIdentityPool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIdentityPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdentities", func(t *testing.T) {
        input := &cognitoidentity.DeleteIdentitiesInput{}
        output := &cognitoidentity.DeleteIdentitiesOutput{}

        mockClient.On("DeleteIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdentityPool", func(t *testing.T) {
        input := &cognitoidentity.DeleteIdentityPoolInput{}
        output := &cognitoidentity.DeleteIdentityPoolOutput{}

        mockClient.On("DeleteIdentityPool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdentityPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdentity", func(t *testing.T) {
        input := &cognitoidentity.DescribeIdentityInput{}
        output := &cognitoidentity.DescribeIdentityOutput{}

        mockClient.On("DescribeIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIdentityPool", func(t *testing.T) {
        input := &cognitoidentity.DescribeIdentityPoolInput{}
        output := &cognitoidentity.DescribeIdentityPoolOutput{}

        mockClient.On("DescribeIdentityPool", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIdentityPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCredentialsForIdentity", func(t *testing.T) {
        input := &cognitoidentity.GetCredentialsForIdentityInput{}
        output := &cognitoidentity.GetCredentialsForIdentityOutput{}

        mockClient.On("GetCredentialsForIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.GetCredentialsForIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetId", func(t *testing.T) {
        input := &cognitoidentity.GetIdInput{}
        output := &cognitoidentity.GetIdOutput{}

        mockClient.On("GetId", ctx, input).Return(output, nil)

        result, err := mockClient.GetId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityPoolRoles", func(t *testing.T) {
        input := &cognitoidentity.GetIdentityPoolRolesInput{}
        output := &cognitoidentity.GetIdentityPoolRolesOutput{}

        mockClient.On("GetIdentityPoolRoles", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityPoolRoles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOpenIdToken", func(t *testing.T) {
        input := &cognitoidentity.GetOpenIdTokenInput{}
        output := &cognitoidentity.GetOpenIdTokenOutput{}

        mockClient.On("GetOpenIdToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetOpenIdToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOpenIdTokenForDeveloperIdentity", func(t *testing.T) {
        input := &cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput{}
        output := &cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput{}

        mockClient.On("GetOpenIdTokenForDeveloperIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.GetOpenIdTokenForDeveloperIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPrincipalTagAttributeMap", func(t *testing.T) {
        input := &cognitoidentity.GetPrincipalTagAttributeMapInput{}
        output := &cognitoidentity.GetPrincipalTagAttributeMapOutput{}

        mockClient.On("GetPrincipalTagAttributeMap", ctx, input).Return(output, nil)

        result, err := mockClient.GetPrincipalTagAttributeMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentities", func(t *testing.T) {
        input := &cognitoidentity.ListIdentitiesInput{}
        output := &cognitoidentity.ListIdentitiesOutput{}

        mockClient.On("ListIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityPools", func(t *testing.T) {
        input := &cognitoidentity.ListIdentityPoolsInput{}
        output := &cognitoidentity.ListIdentityPoolsOutput{}

        mockClient.On("ListIdentityPools", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityPools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cognitoidentity.ListTagsForResourceInput{}
        output := &cognitoidentity.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLookupDeveloperIdentity", func(t *testing.T) {
        input := &cognitoidentity.LookupDeveloperIdentityInput{}
        output := &cognitoidentity.LookupDeveloperIdentityOutput{}

        mockClient.On("LookupDeveloperIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.LookupDeveloperIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergeDeveloperIdentities", func(t *testing.T) {
        input := &cognitoidentity.MergeDeveloperIdentitiesInput{}
        output := &cognitoidentity.MergeDeveloperIdentitiesOutput{}

        mockClient.On("MergeDeveloperIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.MergeDeveloperIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIdentityPoolRoles", func(t *testing.T) {
        input := &cognitoidentity.SetIdentityPoolRolesInput{}
        output := &cognitoidentity.SetIdentityPoolRolesOutput{}

        mockClient.On("SetIdentityPoolRoles", ctx, input).Return(output, nil)

        result, err := mockClient.SetIdentityPoolRoles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetPrincipalTagAttributeMap", func(t *testing.T) {
        input := &cognitoidentity.SetPrincipalTagAttributeMapInput{}
        output := &cognitoidentity.SetPrincipalTagAttributeMapOutput{}

        mockClient.On("SetPrincipalTagAttributeMap", ctx, input).Return(output, nil)

        result, err := mockClient.SetPrincipalTagAttributeMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cognitoidentity.TagResourceInput{}
        output := &cognitoidentity.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnlinkDeveloperIdentity", func(t *testing.T) {
        input := &cognitoidentity.UnlinkDeveloperIdentityInput{}
        output := &cognitoidentity.UnlinkDeveloperIdentityOutput{}

        mockClient.On("UnlinkDeveloperIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.UnlinkDeveloperIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnlinkIdentity", func(t *testing.T) {
        input := &cognitoidentity.UnlinkIdentityInput{}
        output := &cognitoidentity.UnlinkIdentityOutput{}

        mockClient.On("UnlinkIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.UnlinkIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cognitoidentity.UntagResourceInput{}
        output := &cognitoidentity.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdentityPool", func(t *testing.T) {
        input := &cognitoidentity.UpdateIdentityPoolInput{}
        output := &cognitoidentity.UpdateIdentityPoolOutput{}

        mockClient.On("UpdateIdentityPool", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdentityPool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

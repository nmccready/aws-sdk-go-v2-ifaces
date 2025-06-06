// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package verifiedpermissions_test

// tests for the verifiedpermissions service interface for 
// this ../../../service/verifiedpermissions/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/verifiedpermissions/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/verifiedpermissions/verifiedpermissions_iface"
	"github.com/stretchr/testify/assert"
)

func TestVerifiedpermissionsServiceCanBeMocked(t *testing.T) {
	var iface verifiedpermissions_iface.IClient
	iface = &verifiedpermissions.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := verifiedpermissions.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetPolicy", func(t *testing.T) {
        input := &verifiedpermissions.BatchGetPolicyInput{}
        output := &verifiedpermissions.BatchGetPolicyOutput{}

        mockClient.On("BatchGetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchIsAuthorized", func(t *testing.T) {
        input := &verifiedpermissions.BatchIsAuthorizedInput{}
        output := &verifiedpermissions.BatchIsAuthorizedOutput{}

        mockClient.On("BatchIsAuthorized", ctx, input).Return(output, nil)

        result, err := mockClient.BatchIsAuthorized(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchIsAuthorizedWithToken", func(t *testing.T) {
        input := &verifiedpermissions.BatchIsAuthorizedWithTokenInput{}
        output := &verifiedpermissions.BatchIsAuthorizedWithTokenOutput{}

        mockClient.On("BatchIsAuthorizedWithToken", ctx, input).Return(output, nil)

        result, err := mockClient.BatchIsAuthorizedWithToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIdentitySource", func(t *testing.T) {
        input := &verifiedpermissions.CreateIdentitySourceInput{}
        output := &verifiedpermissions.CreateIdentitySourceOutput{}

        mockClient.On("CreateIdentitySource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIdentitySource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePolicy", func(t *testing.T) {
        input := &verifiedpermissions.CreatePolicyInput{}
        output := &verifiedpermissions.CreatePolicyOutput{}

        mockClient.On("CreatePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePolicyStore", func(t *testing.T) {
        input := &verifiedpermissions.CreatePolicyStoreInput{}
        output := &verifiedpermissions.CreatePolicyStoreOutput{}

        mockClient.On("CreatePolicyStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePolicyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePolicyTemplate", func(t *testing.T) {
        input := &verifiedpermissions.CreatePolicyTemplateInput{}
        output := &verifiedpermissions.CreatePolicyTemplateOutput{}

        mockClient.On("CreatePolicyTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePolicyTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdentitySource", func(t *testing.T) {
        input := &verifiedpermissions.DeleteIdentitySourceInput{}
        output := &verifiedpermissions.DeleteIdentitySourceOutput{}

        mockClient.On("DeleteIdentitySource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdentitySource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicy", func(t *testing.T) {
        input := &verifiedpermissions.DeletePolicyInput{}
        output := &verifiedpermissions.DeletePolicyOutput{}

        mockClient.On("DeletePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicyStore", func(t *testing.T) {
        input := &verifiedpermissions.DeletePolicyStoreInput{}
        output := &verifiedpermissions.DeletePolicyStoreOutput{}

        mockClient.On("DeletePolicyStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicyTemplate", func(t *testing.T) {
        input := &verifiedpermissions.DeletePolicyTemplateInput{}
        output := &verifiedpermissions.DeletePolicyTemplateOutput{}

        mockClient.On("DeletePolicyTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicyTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentitySource", func(t *testing.T) {
        input := &verifiedpermissions.GetIdentitySourceInput{}
        output := &verifiedpermissions.GetIdentitySourceOutput{}

        mockClient.On("GetIdentitySource", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentitySource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &verifiedpermissions.GetPolicyInput{}
        output := &verifiedpermissions.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicyStore", func(t *testing.T) {
        input := &verifiedpermissions.GetPolicyStoreInput{}
        output := &verifiedpermissions.GetPolicyStoreOutput{}

        mockClient.On("GetPolicyStore", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicyTemplate", func(t *testing.T) {
        input := &verifiedpermissions.GetPolicyTemplateInput{}
        output := &verifiedpermissions.GetPolicyTemplateOutput{}

        mockClient.On("GetPolicyTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicyTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchema", func(t *testing.T) {
        input := &verifiedpermissions.GetSchemaInput{}
        output := &verifiedpermissions.GetSchemaOutput{}

        mockClient.On("GetSchema", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIsAuthorized", func(t *testing.T) {
        input := &verifiedpermissions.IsAuthorizedInput{}
        output := &verifiedpermissions.IsAuthorizedOutput{}

        mockClient.On("IsAuthorized", ctx, input).Return(output, nil)

        result, err := mockClient.IsAuthorized(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIsAuthorizedWithToken", func(t *testing.T) {
        input := &verifiedpermissions.IsAuthorizedWithTokenInput{}
        output := &verifiedpermissions.IsAuthorizedWithTokenOutput{}

        mockClient.On("IsAuthorizedWithToken", ctx, input).Return(output, nil)

        result, err := mockClient.IsAuthorizedWithToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentitySources", func(t *testing.T) {
        input := &verifiedpermissions.ListIdentitySourcesInput{}
        output := &verifiedpermissions.ListIdentitySourcesOutput{}

        mockClient.On("ListIdentitySources", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentitySources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicies", func(t *testing.T) {
        input := &verifiedpermissions.ListPoliciesInput{}
        output := &verifiedpermissions.ListPoliciesOutput{}

        mockClient.On("ListPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicyStores", func(t *testing.T) {
        input := &verifiedpermissions.ListPolicyStoresInput{}
        output := &verifiedpermissions.ListPolicyStoresOutput{}

        mockClient.On("ListPolicyStores", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicyStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicyTemplates", func(t *testing.T) {
        input := &verifiedpermissions.ListPolicyTemplatesInput{}
        output := &verifiedpermissions.ListPolicyTemplatesOutput{}

        mockClient.On("ListPolicyTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicyTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &verifiedpermissions.ListTagsForResourceInput{}
        output := &verifiedpermissions.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSchema", func(t *testing.T) {
        input := &verifiedpermissions.PutSchemaInput{}
        output := &verifiedpermissions.PutSchemaOutput{}

        mockClient.On("PutSchema", ctx, input).Return(output, nil)

        result, err := mockClient.PutSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &verifiedpermissions.TagResourceInput{}
        output := &verifiedpermissions.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &verifiedpermissions.UntagResourceInput{}
        output := &verifiedpermissions.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdentitySource", func(t *testing.T) {
        input := &verifiedpermissions.UpdateIdentitySourceInput{}
        output := &verifiedpermissions.UpdateIdentitySourceOutput{}

        mockClient.On("UpdateIdentitySource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdentitySource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePolicy", func(t *testing.T) {
        input := &verifiedpermissions.UpdatePolicyInput{}
        output := &verifiedpermissions.UpdatePolicyOutput{}

        mockClient.On("UpdatePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePolicyStore", func(t *testing.T) {
        input := &verifiedpermissions.UpdatePolicyStoreInput{}
        output := &verifiedpermissions.UpdatePolicyStoreOutput{}

        mockClient.On("UpdatePolicyStore", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePolicyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePolicyTemplate", func(t *testing.T) {
        input := &verifiedpermissions.UpdatePolicyTemplateInput{}
        output := &verifiedpermissions.UpdatePolicyTemplateOutput{}

        mockClient.On("UpdatePolicyTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePolicyTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

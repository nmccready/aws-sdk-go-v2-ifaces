package connectcases_test

// tests for the connectcases service interface for this ../../../service/connectcases/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/connectcases"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectcases/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectcases/connectcases_iface"
	"github.com/stretchr/testify/assert"
)

func TestConnectcasesServiceCanBeMocked(t *testing.T) {
	var iface connectcases_iface.IClient
	iface = &connectcases.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := connectcases.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetField", func(t *testing.T) {
        input := &connectcases.BatchGetFieldInput{}
        output := &connectcases.BatchGetFieldOutput{}

        mockClient.On("BatchGetField", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetField(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutFieldOptions", func(t *testing.T) {
        input := &connectcases.BatchPutFieldOptionsInput{}
        output := &connectcases.BatchPutFieldOptionsOutput{}

        mockClient.On("BatchPutFieldOptions", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutFieldOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCase", func(t *testing.T) {
        input := &connectcases.CreateCaseInput{}
        output := &connectcases.CreateCaseOutput{}

        mockClient.On("CreateCase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &connectcases.CreateDomainInput{}
        output := &connectcases.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateField", func(t *testing.T) {
        input := &connectcases.CreateFieldInput{}
        output := &connectcases.CreateFieldOutput{}

        mockClient.On("CreateField", ctx, input).Return(output, nil)

        result, err := mockClient.CreateField(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLayout", func(t *testing.T) {
        input := &connectcases.CreateLayoutInput{}
        output := &connectcases.CreateLayoutOutput{}

        mockClient.On("CreateLayout", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLayout(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRelatedItem", func(t *testing.T) {
        input := &connectcases.CreateRelatedItemInput{}
        output := &connectcases.CreateRelatedItemOutput{}

        mockClient.On("CreateRelatedItem", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRelatedItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplate", func(t *testing.T) {
        input := &connectcases.CreateTemplateInput{}
        output := &connectcases.CreateTemplateOutput{}

        mockClient.On("CreateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &connectcases.DeleteDomainInput{}
        output := &connectcases.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteField", func(t *testing.T) {
        input := &connectcases.DeleteFieldInput{}
        output := &connectcases.DeleteFieldOutput{}

        mockClient.On("DeleteField", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteField(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLayout", func(t *testing.T) {
        input := &connectcases.DeleteLayoutInput{}
        output := &connectcases.DeleteLayoutOutput{}

        mockClient.On("DeleteLayout", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLayout(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplate", func(t *testing.T) {
        input := &connectcases.DeleteTemplateInput{}
        output := &connectcases.DeleteTemplateOutput{}

        mockClient.On("DeleteTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCase", func(t *testing.T) {
        input := &connectcases.GetCaseInput{}
        output := &connectcases.GetCaseOutput{}

        mockClient.On("GetCase", ctx, input).Return(output, nil)

        result, err := mockClient.GetCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCaseAuditEvents", func(t *testing.T) {
        input := &connectcases.GetCaseAuditEventsInput{}
        output := &connectcases.GetCaseAuditEventsOutput{}

        mockClient.On("GetCaseAuditEvents", ctx, input).Return(output, nil)

        result, err := mockClient.GetCaseAuditEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCaseEventConfiguration", func(t *testing.T) {
        input := &connectcases.GetCaseEventConfigurationInput{}
        output := &connectcases.GetCaseEventConfigurationOutput{}

        mockClient.On("GetCaseEventConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetCaseEventConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomain", func(t *testing.T) {
        input := &connectcases.GetDomainInput{}
        output := &connectcases.GetDomainOutput{}

        mockClient.On("GetDomain", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLayout", func(t *testing.T) {
        input := &connectcases.GetLayoutInput{}
        output := &connectcases.GetLayoutOutput{}

        mockClient.On("GetLayout", ctx, input).Return(output, nil)

        result, err := mockClient.GetLayout(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTemplate", func(t *testing.T) {
        input := &connectcases.GetTemplateInput{}
        output := &connectcases.GetTemplateOutput{}

        mockClient.On("GetTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCasesForContact", func(t *testing.T) {
        input := &connectcases.ListCasesForContactInput{}
        output := &connectcases.ListCasesForContactOutput{}

        mockClient.On("ListCasesForContact", ctx, input).Return(output, nil)

        result, err := mockClient.ListCasesForContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &connectcases.ListDomainsInput{}
        output := &connectcases.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFieldOptions", func(t *testing.T) {
        input := &connectcases.ListFieldOptionsInput{}
        output := &connectcases.ListFieldOptionsOutput{}

        mockClient.On("ListFieldOptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListFieldOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFields", func(t *testing.T) {
        input := &connectcases.ListFieldsInput{}
        output := &connectcases.ListFieldsOutput{}

        mockClient.On("ListFields", ctx, input).Return(output, nil)

        result, err := mockClient.ListFields(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLayouts", func(t *testing.T) {
        input := &connectcases.ListLayoutsInput{}
        output := &connectcases.ListLayoutsOutput{}

        mockClient.On("ListLayouts", ctx, input).Return(output, nil)

        result, err := mockClient.ListLayouts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &connectcases.ListTagsForResourceInput{}
        output := &connectcases.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplates", func(t *testing.T) {
        input := &connectcases.ListTemplatesInput{}
        output := &connectcases.ListTemplatesOutput{}

        mockClient.On("ListTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutCaseEventConfiguration", func(t *testing.T) {
        input := &connectcases.PutCaseEventConfigurationInput{}
        output := &connectcases.PutCaseEventConfigurationOutput{}

        mockClient.On("PutCaseEventConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutCaseEventConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchCases", func(t *testing.T) {
        input := &connectcases.SearchCasesInput{}
        output := &connectcases.SearchCasesOutput{}

        mockClient.On("SearchCases", ctx, input).Return(output, nil)

        result, err := mockClient.SearchCases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchRelatedItems", func(t *testing.T) {
        input := &connectcases.SearchRelatedItemsInput{}
        output := &connectcases.SearchRelatedItemsOutput{}

        mockClient.On("SearchRelatedItems", ctx, input).Return(output, nil)

        result, err := mockClient.SearchRelatedItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &connectcases.TagResourceInput{}
        output := &connectcases.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &connectcases.UntagResourceInput{}
        output := &connectcases.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCase", func(t *testing.T) {
        input := &connectcases.UpdateCaseInput{}
        output := &connectcases.UpdateCaseOutput{}

        mockClient.On("UpdateCase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateField", func(t *testing.T) {
        input := &connectcases.UpdateFieldInput{}
        output := &connectcases.UpdateFieldOutput{}

        mockClient.On("UpdateField", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateField(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLayout", func(t *testing.T) {
        input := &connectcases.UpdateLayoutInput{}
        output := &connectcases.UpdateLayoutOutput{}

        mockClient.On("UpdateLayout", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLayout(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplate", func(t *testing.T) {
        input := &connectcases.UpdateTemplateInput{}
        output := &connectcases.UpdateTemplateOutput{}

        mockClient.On("UpdateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

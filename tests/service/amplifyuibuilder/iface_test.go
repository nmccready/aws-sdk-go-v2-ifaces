package amplifyuibuilder_test

// tests for the amplifyuibuilder service interface for this ../../../service/amplifyuibuilder/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/amplifyuibuilder/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/amplifyuibuilder/amplifyuibuilder_iface"
	"github.com/stretchr/testify/assert"
)

func TestAmplifyuibuilderServiceCanBeMocked(t *testing.T) {
	var iface amplifyuibuilder_iface.IClient
	iface = &amplifyuibuilder.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := amplifyuibuilder.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComponent", func(t *testing.T) {
        input := &amplifyuibuilder.CreateComponentInput{}
        output := &amplifyuibuilder.CreateComponentOutput{}

        mockClient.On("CreateComponent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateForm", func(t *testing.T) {
        input := &amplifyuibuilder.CreateFormInput{}
        output := &amplifyuibuilder.CreateFormOutput{}

        mockClient.On("CreateForm", ctx, input).Return(output, nil)

        result, err := mockClient.CreateForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTheme", func(t *testing.T) {
        input := &amplifyuibuilder.CreateThemeInput{}
        output := &amplifyuibuilder.CreateThemeOutput{}

        mockClient.On("CreateTheme", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteComponent", func(t *testing.T) {
        input := &amplifyuibuilder.DeleteComponentInput{}
        output := &amplifyuibuilder.DeleteComponentOutput{}

        mockClient.On("DeleteComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteForm", func(t *testing.T) {
        input := &amplifyuibuilder.DeleteFormInput{}
        output := &amplifyuibuilder.DeleteFormOutput{}

        mockClient.On("DeleteForm", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTheme", func(t *testing.T) {
        input := &amplifyuibuilder.DeleteThemeInput{}
        output := &amplifyuibuilder.DeleteThemeOutput{}

        mockClient.On("DeleteTheme", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExchangeCodeForToken", func(t *testing.T) {
        input := &amplifyuibuilder.ExchangeCodeForTokenInput{}
        output := &amplifyuibuilder.ExchangeCodeForTokenOutput{}

        mockClient.On("ExchangeCodeForToken", ctx, input).Return(output, nil)

        result, err := mockClient.ExchangeCodeForToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportComponents", func(t *testing.T) {
        input := &amplifyuibuilder.ExportComponentsInput{}
        output := &amplifyuibuilder.ExportComponentsOutput{}

        mockClient.On("ExportComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ExportComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportForms", func(t *testing.T) {
        input := &amplifyuibuilder.ExportFormsInput{}
        output := &amplifyuibuilder.ExportFormsOutput{}

        mockClient.On("ExportForms", ctx, input).Return(output, nil)

        result, err := mockClient.ExportForms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportThemes", func(t *testing.T) {
        input := &amplifyuibuilder.ExportThemesInput{}
        output := &amplifyuibuilder.ExportThemesOutput{}

        mockClient.On("ExportThemes", ctx, input).Return(output, nil)

        result, err := mockClient.ExportThemes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCodegenJob", func(t *testing.T) {
        input := &amplifyuibuilder.GetCodegenJobInput{}
        output := &amplifyuibuilder.GetCodegenJobOutput{}

        mockClient.On("GetCodegenJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetCodegenJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComponent", func(t *testing.T) {
        input := &amplifyuibuilder.GetComponentInput{}
        output := &amplifyuibuilder.GetComponentOutput{}

        mockClient.On("GetComponent", ctx, input).Return(output, nil)

        result, err := mockClient.GetComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetForm", func(t *testing.T) {
        input := &amplifyuibuilder.GetFormInput{}
        output := &amplifyuibuilder.GetFormOutput{}

        mockClient.On("GetForm", ctx, input).Return(output, nil)

        result, err := mockClient.GetForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetadata", func(t *testing.T) {
        input := &amplifyuibuilder.GetMetadataInput{}
        output := &amplifyuibuilder.GetMetadataOutput{}

        mockClient.On("GetMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTheme", func(t *testing.T) {
        input := &amplifyuibuilder.GetThemeInput{}
        output := &amplifyuibuilder.GetThemeOutput{}

        mockClient.On("GetTheme", ctx, input).Return(output, nil)

        result, err := mockClient.GetTheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCodegenJobs", func(t *testing.T) {
        input := &amplifyuibuilder.ListCodegenJobsInput{}
        output := &amplifyuibuilder.ListCodegenJobsOutput{}

        mockClient.On("ListCodegenJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListCodegenJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponents", func(t *testing.T) {
        input := &amplifyuibuilder.ListComponentsInput{}
        output := &amplifyuibuilder.ListComponentsOutput{}

        mockClient.On("ListComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListForms", func(t *testing.T) {
        input := &amplifyuibuilder.ListFormsInput{}
        output := &amplifyuibuilder.ListFormsOutput{}

        mockClient.On("ListForms", ctx, input).Return(output, nil)

        result, err := mockClient.ListForms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &amplifyuibuilder.ListTagsForResourceInput{}
        output := &amplifyuibuilder.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThemes", func(t *testing.T) {
        input := &amplifyuibuilder.ListThemesInput{}
        output := &amplifyuibuilder.ListThemesOutput{}

        mockClient.On("ListThemes", ctx, input).Return(output, nil)

        result, err := mockClient.ListThemes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMetadataFlag", func(t *testing.T) {
        input := &amplifyuibuilder.PutMetadataFlagInput{}
        output := &amplifyuibuilder.PutMetadataFlagOutput{}

        mockClient.On("PutMetadataFlag", ctx, input).Return(output, nil)

        result, err := mockClient.PutMetadataFlag(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRefreshToken", func(t *testing.T) {
        input := &amplifyuibuilder.RefreshTokenInput{}
        output := &amplifyuibuilder.RefreshTokenOutput{}

        mockClient.On("RefreshToken", ctx, input).Return(output, nil)

        result, err := mockClient.RefreshToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCodegenJob", func(t *testing.T) {
        input := &amplifyuibuilder.StartCodegenJobInput{}
        output := &amplifyuibuilder.StartCodegenJobOutput{}

        mockClient.On("StartCodegenJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartCodegenJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &amplifyuibuilder.TagResourceInput{}
        output := &amplifyuibuilder.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &amplifyuibuilder.UntagResourceInput{}
        output := &amplifyuibuilder.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateComponent", func(t *testing.T) {
        input := &amplifyuibuilder.UpdateComponentInput{}
        output := &amplifyuibuilder.UpdateComponentOutput{}

        mockClient.On("UpdateComponent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateForm", func(t *testing.T) {
        input := &amplifyuibuilder.UpdateFormInput{}
        output := &amplifyuibuilder.UpdateFormOutput{}

        mockClient.On("UpdateForm", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateForm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTheme", func(t *testing.T) {
        input := &amplifyuibuilder.UpdateThemeInput{}
        output := &amplifyuibuilder.UpdateThemeOutput{}

        mockClient.On("UpdateTheme", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

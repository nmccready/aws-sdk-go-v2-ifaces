package voiceid_test

// tests for the voiceid service interface for this ../../../service/voiceid/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/voiceid"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/voiceid/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/voiceid/voiceid_iface"
	"github.com/stretchr/testify/assert"
)

func TestVoiceidServiceCanBeMocked(t *testing.T) {
	var iface voiceid_iface.IClient
	iface = &voiceid.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := voiceid.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateFraudster", func(t *testing.T) {
        input := &voiceid.AssociateFraudsterInput{}
        output := &voiceid.AssociateFraudsterOutput{}

        mockClient.On("AssociateFraudster", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateFraudster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &voiceid.CreateDomainInput{}
        output := &voiceid.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWatchlist", func(t *testing.T) {
        input := &voiceid.CreateWatchlistInput{}
        output := &voiceid.CreateWatchlistOutput{}

        mockClient.On("CreateWatchlist", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWatchlist(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &voiceid.DeleteDomainInput{}
        output := &voiceid.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFraudster", func(t *testing.T) {
        input := &voiceid.DeleteFraudsterInput{}
        output := &voiceid.DeleteFraudsterOutput{}

        mockClient.On("DeleteFraudster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFraudster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSpeaker", func(t *testing.T) {
        input := &voiceid.DeleteSpeakerInput{}
        output := &voiceid.DeleteSpeakerOutput{}

        mockClient.On("DeleteSpeaker", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSpeaker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWatchlist", func(t *testing.T) {
        input := &voiceid.DeleteWatchlistInput{}
        output := &voiceid.DeleteWatchlistOutput{}

        mockClient.On("DeleteWatchlist", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWatchlist(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomain", func(t *testing.T) {
        input := &voiceid.DescribeDomainInput{}
        output := &voiceid.DescribeDomainOutput{}

        mockClient.On("DescribeDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFraudster", func(t *testing.T) {
        input := &voiceid.DescribeFraudsterInput{}
        output := &voiceid.DescribeFraudsterOutput{}

        mockClient.On("DescribeFraudster", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFraudster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFraudsterRegistrationJob", func(t *testing.T) {
        input := &voiceid.DescribeFraudsterRegistrationJobInput{}
        output := &voiceid.DescribeFraudsterRegistrationJobOutput{}

        mockClient.On("DescribeFraudsterRegistrationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFraudsterRegistrationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpeaker", func(t *testing.T) {
        input := &voiceid.DescribeSpeakerInput{}
        output := &voiceid.DescribeSpeakerOutput{}

        mockClient.On("DescribeSpeaker", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpeaker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSpeakerEnrollmentJob", func(t *testing.T) {
        input := &voiceid.DescribeSpeakerEnrollmentJobInput{}
        output := &voiceid.DescribeSpeakerEnrollmentJobOutput{}

        mockClient.On("DescribeSpeakerEnrollmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSpeakerEnrollmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWatchlist", func(t *testing.T) {
        input := &voiceid.DescribeWatchlistInput{}
        output := &voiceid.DescribeWatchlistOutput{}

        mockClient.On("DescribeWatchlist", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWatchlist(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFraudster", func(t *testing.T) {
        input := &voiceid.DisassociateFraudsterInput{}
        output := &voiceid.DisassociateFraudsterOutput{}

        mockClient.On("DisassociateFraudster", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFraudster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEvaluateSession", func(t *testing.T) {
        input := &voiceid.EvaluateSessionInput{}
        output := &voiceid.EvaluateSessionOutput{}

        mockClient.On("EvaluateSession", ctx, input).Return(output, nil)

        result, err := mockClient.EvaluateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &voiceid.ListDomainsInput{}
        output := &voiceid.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFraudsterRegistrationJobs", func(t *testing.T) {
        input := &voiceid.ListFraudsterRegistrationJobsInput{}
        output := &voiceid.ListFraudsterRegistrationJobsOutput{}

        mockClient.On("ListFraudsterRegistrationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListFraudsterRegistrationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFraudsters", func(t *testing.T) {
        input := &voiceid.ListFraudstersInput{}
        output := &voiceid.ListFraudstersOutput{}

        mockClient.On("ListFraudsters", ctx, input).Return(output, nil)

        result, err := mockClient.ListFraudsters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSpeakerEnrollmentJobs", func(t *testing.T) {
        input := &voiceid.ListSpeakerEnrollmentJobsInput{}
        output := &voiceid.ListSpeakerEnrollmentJobsOutput{}

        mockClient.On("ListSpeakerEnrollmentJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSpeakerEnrollmentJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSpeakers", func(t *testing.T) {
        input := &voiceid.ListSpeakersInput{}
        output := &voiceid.ListSpeakersOutput{}

        mockClient.On("ListSpeakers", ctx, input).Return(output, nil)

        result, err := mockClient.ListSpeakers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &voiceid.ListTagsForResourceInput{}
        output := &voiceid.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWatchlists", func(t *testing.T) {
        input := &voiceid.ListWatchlistsInput{}
        output := &voiceid.ListWatchlistsOutput{}

        mockClient.On("ListWatchlists", ctx, input).Return(output, nil)

        result, err := mockClient.ListWatchlists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestOptOutSpeaker", func(t *testing.T) {
        input := &voiceid.OptOutSpeakerInput{}
        output := &voiceid.OptOutSpeakerOutput{}

        mockClient.On("OptOutSpeaker", ctx, input).Return(output, nil)

        result, err := mockClient.OptOutSpeaker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFraudsterRegistrationJob", func(t *testing.T) {
        input := &voiceid.StartFraudsterRegistrationJobInput{}
        output := &voiceid.StartFraudsterRegistrationJobOutput{}

        mockClient.On("StartFraudsterRegistrationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartFraudsterRegistrationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSpeakerEnrollmentJob", func(t *testing.T) {
        input := &voiceid.StartSpeakerEnrollmentJobInput{}
        output := &voiceid.StartSpeakerEnrollmentJobOutput{}

        mockClient.On("StartSpeakerEnrollmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartSpeakerEnrollmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &voiceid.TagResourceInput{}
        output := &voiceid.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &voiceid.UntagResourceInput{}
        output := &voiceid.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomain", func(t *testing.T) {
        input := &voiceid.UpdateDomainInput{}
        output := &voiceid.UpdateDomainOutput{}

        mockClient.On("UpdateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWatchlist", func(t *testing.T) {
        input := &voiceid.UpdateWatchlistInput{}
        output := &voiceid.UpdateWatchlistOutput{}

        mockClient.On("UpdateWatchlist", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWatchlist(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

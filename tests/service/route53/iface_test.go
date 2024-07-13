package route53_test

// tests for the route53 service interface for this ../../../service/route53/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53/route53_iface"
	"github.com/stretchr/testify/assert"
)

func TestRoute53ServiceCanBeMocked(t *testing.T) {
	var iface route53_iface.IClient
	iface = &route53.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := route53.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateKeySigningKey", func(t *testing.T) {
        input := &route53.ActivateKeySigningKeyInput{}
        output := &route53.ActivateKeySigningKeyOutput{}

        mockClient.On("ActivateKeySigningKey", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateKeySigningKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateVPCWithHostedZone", func(t *testing.T) {
        input := &route53.AssociateVPCWithHostedZoneInput{}
        output := &route53.AssociateVPCWithHostedZoneOutput{}

        mockClient.On("AssociateVPCWithHostedZone", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateVPCWithHostedZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChangeCidrCollection", func(t *testing.T) {
        input := &route53.ChangeCidrCollectionInput{}
        output := &route53.ChangeCidrCollectionOutput{}

        mockClient.On("ChangeCidrCollection", ctx, input).Return(output, nil)

        result, err := mockClient.ChangeCidrCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChangeResourceRecordSets", func(t *testing.T) {
        input := &route53.ChangeResourceRecordSetsInput{}
        output := &route53.ChangeResourceRecordSetsOutput{}

        mockClient.On("ChangeResourceRecordSets", ctx, input).Return(output, nil)

        result, err := mockClient.ChangeResourceRecordSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestChangeTagsForResource", func(t *testing.T) {
        input := &route53.ChangeTagsForResourceInput{}
        output := &route53.ChangeTagsForResourceOutput{}

        mockClient.On("ChangeTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ChangeTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCidrCollection", func(t *testing.T) {
        input := &route53.CreateCidrCollectionInput{}
        output := &route53.CreateCidrCollectionOutput{}

        mockClient.On("CreateCidrCollection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCidrCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHealthCheck", func(t *testing.T) {
        input := &route53.CreateHealthCheckInput{}
        output := &route53.CreateHealthCheckOutput{}

        mockClient.On("CreateHealthCheck", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHealthCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHostedZone", func(t *testing.T) {
        input := &route53.CreateHostedZoneInput{}
        output := &route53.CreateHostedZoneOutput{}

        mockClient.On("CreateHostedZone", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHostedZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKeySigningKey", func(t *testing.T) {
        input := &route53.CreateKeySigningKeyInput{}
        output := &route53.CreateKeySigningKeyOutput{}

        mockClient.On("CreateKeySigningKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKeySigningKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQueryLoggingConfig", func(t *testing.T) {
        input := &route53.CreateQueryLoggingConfigInput{}
        output := &route53.CreateQueryLoggingConfigOutput{}

        mockClient.On("CreateQueryLoggingConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQueryLoggingConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReusableDelegationSet", func(t *testing.T) {
        input := &route53.CreateReusableDelegationSetInput{}
        output := &route53.CreateReusableDelegationSetOutput{}

        mockClient.On("CreateReusableDelegationSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReusableDelegationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficPolicy", func(t *testing.T) {
        input := &route53.CreateTrafficPolicyInput{}
        output := &route53.CreateTrafficPolicyOutput{}

        mockClient.On("CreateTrafficPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficPolicyInstance", func(t *testing.T) {
        input := &route53.CreateTrafficPolicyInstanceInput{}
        output := &route53.CreateTrafficPolicyInstanceOutput{}

        mockClient.On("CreateTrafficPolicyInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficPolicyInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrafficPolicyVersion", func(t *testing.T) {
        input := &route53.CreateTrafficPolicyVersionInput{}
        output := &route53.CreateTrafficPolicyVersionOutput{}

        mockClient.On("CreateTrafficPolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrafficPolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVPCAssociationAuthorization", func(t *testing.T) {
        input := &route53.CreateVPCAssociationAuthorizationInput{}
        output := &route53.CreateVPCAssociationAuthorizationOutput{}

        mockClient.On("CreateVPCAssociationAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVPCAssociationAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateKeySigningKey", func(t *testing.T) {
        input := &route53.DeactivateKeySigningKeyInput{}
        output := &route53.DeactivateKeySigningKeyOutput{}

        mockClient.On("DeactivateKeySigningKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateKeySigningKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCidrCollection", func(t *testing.T) {
        input := &route53.DeleteCidrCollectionInput{}
        output := &route53.DeleteCidrCollectionOutput{}

        mockClient.On("DeleteCidrCollection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCidrCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHealthCheck", func(t *testing.T) {
        input := &route53.DeleteHealthCheckInput{}
        output := &route53.DeleteHealthCheckOutput{}

        mockClient.On("DeleteHealthCheck", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHealthCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHostedZone", func(t *testing.T) {
        input := &route53.DeleteHostedZoneInput{}
        output := &route53.DeleteHostedZoneOutput{}

        mockClient.On("DeleteHostedZone", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHostedZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKeySigningKey", func(t *testing.T) {
        input := &route53.DeleteKeySigningKeyInput{}
        output := &route53.DeleteKeySigningKeyOutput{}

        mockClient.On("DeleteKeySigningKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKeySigningKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueryLoggingConfig", func(t *testing.T) {
        input := &route53.DeleteQueryLoggingConfigInput{}
        output := &route53.DeleteQueryLoggingConfigOutput{}

        mockClient.On("DeleteQueryLoggingConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueryLoggingConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReusableDelegationSet", func(t *testing.T) {
        input := &route53.DeleteReusableDelegationSetInput{}
        output := &route53.DeleteReusableDelegationSetOutput{}

        mockClient.On("DeleteReusableDelegationSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReusableDelegationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrafficPolicy", func(t *testing.T) {
        input := &route53.DeleteTrafficPolicyInput{}
        output := &route53.DeleteTrafficPolicyOutput{}

        mockClient.On("DeleteTrafficPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrafficPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrafficPolicyInstance", func(t *testing.T) {
        input := &route53.DeleteTrafficPolicyInstanceInput{}
        output := &route53.DeleteTrafficPolicyInstanceOutput{}

        mockClient.On("DeleteTrafficPolicyInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrafficPolicyInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVPCAssociationAuthorization", func(t *testing.T) {
        input := &route53.DeleteVPCAssociationAuthorizationInput{}
        output := &route53.DeleteVPCAssociationAuthorizationOutput{}

        mockClient.On("DeleteVPCAssociationAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVPCAssociationAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableHostedZoneDNSSEC", func(t *testing.T) {
        input := &route53.DisableHostedZoneDNSSECInput{}
        output := &route53.DisableHostedZoneDNSSECOutput{}

        mockClient.On("DisableHostedZoneDNSSEC", ctx, input).Return(output, nil)

        result, err := mockClient.DisableHostedZoneDNSSEC(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateVPCFromHostedZone", func(t *testing.T) {
        input := &route53.DisassociateVPCFromHostedZoneInput{}
        output := &route53.DisassociateVPCFromHostedZoneOutput{}

        mockClient.On("DisassociateVPCFromHostedZone", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateVPCFromHostedZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableHostedZoneDNSSEC", func(t *testing.T) {
        input := &route53.EnableHostedZoneDNSSECInput{}
        output := &route53.EnableHostedZoneDNSSECOutput{}

        mockClient.On("EnableHostedZoneDNSSEC", ctx, input).Return(output, nil)

        result, err := mockClient.EnableHostedZoneDNSSEC(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountLimit", func(t *testing.T) {
        input := &route53.GetAccountLimitInput{}
        output := &route53.GetAccountLimitOutput{}

        mockClient.On("GetAccountLimit", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChange", func(t *testing.T) {
        input := &route53.GetChangeInput{}
        output := &route53.GetChangeOutput{}

        mockClient.On("GetChange", ctx, input).Return(output, nil)

        result, err := mockClient.GetChange(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCheckerIpRanges", func(t *testing.T) {
        input := &route53.GetCheckerIpRangesInput{}
        output := &route53.GetCheckerIpRangesOutput{}

        mockClient.On("GetCheckerIpRanges", ctx, input).Return(output, nil)

        result, err := mockClient.GetCheckerIpRanges(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDNSSEC", func(t *testing.T) {
        input := &route53.GetDNSSECInput{}
        output := &route53.GetDNSSECOutput{}

        mockClient.On("GetDNSSEC", ctx, input).Return(output, nil)

        result, err := mockClient.GetDNSSEC(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGeoLocation", func(t *testing.T) {
        input := &route53.GetGeoLocationInput{}
        output := &route53.GetGeoLocationOutput{}

        mockClient.On("GetGeoLocation", ctx, input).Return(output, nil)

        result, err := mockClient.GetGeoLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHealthCheck", func(t *testing.T) {
        input := &route53.GetHealthCheckInput{}
        output := &route53.GetHealthCheckOutput{}

        mockClient.On("GetHealthCheck", ctx, input).Return(output, nil)

        result, err := mockClient.GetHealthCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHealthCheckCount", func(t *testing.T) {
        input := &route53.GetHealthCheckCountInput{}
        output := &route53.GetHealthCheckCountOutput{}

        mockClient.On("GetHealthCheckCount", ctx, input).Return(output, nil)

        result, err := mockClient.GetHealthCheckCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHealthCheckLastFailureReason", func(t *testing.T) {
        input := &route53.GetHealthCheckLastFailureReasonInput{}
        output := &route53.GetHealthCheckLastFailureReasonOutput{}

        mockClient.On("GetHealthCheckLastFailureReason", ctx, input).Return(output, nil)

        result, err := mockClient.GetHealthCheckLastFailureReason(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHealthCheckStatus", func(t *testing.T) {
        input := &route53.GetHealthCheckStatusInput{}
        output := &route53.GetHealthCheckStatusOutput{}

        mockClient.On("GetHealthCheckStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetHealthCheckStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHostedZone", func(t *testing.T) {
        input := &route53.GetHostedZoneInput{}
        output := &route53.GetHostedZoneOutput{}

        mockClient.On("GetHostedZone", ctx, input).Return(output, nil)

        result, err := mockClient.GetHostedZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHostedZoneCount", func(t *testing.T) {
        input := &route53.GetHostedZoneCountInput{}
        output := &route53.GetHostedZoneCountOutput{}

        mockClient.On("GetHostedZoneCount", ctx, input).Return(output, nil)

        result, err := mockClient.GetHostedZoneCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHostedZoneLimit", func(t *testing.T) {
        input := &route53.GetHostedZoneLimitInput{}
        output := &route53.GetHostedZoneLimitOutput{}

        mockClient.On("GetHostedZoneLimit", ctx, input).Return(output, nil)

        result, err := mockClient.GetHostedZoneLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueryLoggingConfig", func(t *testing.T) {
        input := &route53.GetQueryLoggingConfigInput{}
        output := &route53.GetQueryLoggingConfigOutput{}

        mockClient.On("GetQueryLoggingConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueryLoggingConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReusableDelegationSet", func(t *testing.T) {
        input := &route53.GetReusableDelegationSetInput{}
        output := &route53.GetReusableDelegationSetOutput{}

        mockClient.On("GetReusableDelegationSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetReusableDelegationSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReusableDelegationSetLimit", func(t *testing.T) {
        input := &route53.GetReusableDelegationSetLimitInput{}
        output := &route53.GetReusableDelegationSetLimitOutput{}

        mockClient.On("GetReusableDelegationSetLimit", ctx, input).Return(output, nil)

        result, err := mockClient.GetReusableDelegationSetLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrafficPolicy", func(t *testing.T) {
        input := &route53.GetTrafficPolicyInput{}
        output := &route53.GetTrafficPolicyOutput{}

        mockClient.On("GetTrafficPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrafficPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrafficPolicyInstance", func(t *testing.T) {
        input := &route53.GetTrafficPolicyInstanceInput{}
        output := &route53.GetTrafficPolicyInstanceOutput{}

        mockClient.On("GetTrafficPolicyInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrafficPolicyInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrafficPolicyInstanceCount", func(t *testing.T) {
        input := &route53.GetTrafficPolicyInstanceCountInput{}
        output := &route53.GetTrafficPolicyInstanceCountOutput{}

        mockClient.On("GetTrafficPolicyInstanceCount", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrafficPolicyInstanceCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCidrBlocks", func(t *testing.T) {
        input := &route53.ListCidrBlocksInput{}
        output := &route53.ListCidrBlocksOutput{}

        mockClient.On("ListCidrBlocks", ctx, input).Return(output, nil)

        result, err := mockClient.ListCidrBlocks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCidrCollections", func(t *testing.T) {
        input := &route53.ListCidrCollectionsInput{}
        output := &route53.ListCidrCollectionsOutput{}

        mockClient.On("ListCidrCollections", ctx, input).Return(output, nil)

        result, err := mockClient.ListCidrCollections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCidrLocations", func(t *testing.T) {
        input := &route53.ListCidrLocationsInput{}
        output := &route53.ListCidrLocationsOutput{}

        mockClient.On("ListCidrLocations", ctx, input).Return(output, nil)

        result, err := mockClient.ListCidrLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGeoLocations", func(t *testing.T) {
        input := &route53.ListGeoLocationsInput{}
        output := &route53.ListGeoLocationsOutput{}

        mockClient.On("ListGeoLocations", ctx, input).Return(output, nil)

        result, err := mockClient.ListGeoLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHealthChecks", func(t *testing.T) {
        input := &route53.ListHealthChecksInput{}
        output := &route53.ListHealthChecksOutput{}

        mockClient.On("ListHealthChecks", ctx, input).Return(output, nil)

        result, err := mockClient.ListHealthChecks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHostedZones", func(t *testing.T) {
        input := &route53.ListHostedZonesInput{}
        output := &route53.ListHostedZonesOutput{}

        mockClient.On("ListHostedZones", ctx, input).Return(output, nil)

        result, err := mockClient.ListHostedZones(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHostedZonesByName", func(t *testing.T) {
        input := &route53.ListHostedZonesByNameInput{}
        output := &route53.ListHostedZonesByNameOutput{}

        mockClient.On("ListHostedZonesByName", ctx, input).Return(output, nil)

        result, err := mockClient.ListHostedZonesByName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHostedZonesByVPC", func(t *testing.T) {
        input := &route53.ListHostedZonesByVPCInput{}
        output := &route53.ListHostedZonesByVPCOutput{}

        mockClient.On("ListHostedZonesByVPC", ctx, input).Return(output, nil)

        result, err := mockClient.ListHostedZonesByVPC(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueryLoggingConfigs", func(t *testing.T) {
        input := &route53.ListQueryLoggingConfigsInput{}
        output := &route53.ListQueryLoggingConfigsOutput{}

        mockClient.On("ListQueryLoggingConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueryLoggingConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceRecordSets", func(t *testing.T) {
        input := &route53.ListResourceRecordSetsInput{}
        output := &route53.ListResourceRecordSetsOutput{}

        mockClient.On("ListResourceRecordSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceRecordSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReusableDelegationSets", func(t *testing.T) {
        input := &route53.ListReusableDelegationSetsInput{}
        output := &route53.ListReusableDelegationSetsOutput{}

        mockClient.On("ListReusableDelegationSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListReusableDelegationSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &route53.ListTagsForResourceInput{}
        output := &route53.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResources", func(t *testing.T) {
        input := &route53.ListTagsForResourcesInput{}
        output := &route53.ListTagsForResourcesOutput{}

        mockClient.On("ListTagsForResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrafficPolicies", func(t *testing.T) {
        input := &route53.ListTrafficPoliciesInput{}
        output := &route53.ListTrafficPoliciesOutput{}

        mockClient.On("ListTrafficPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrafficPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrafficPolicyInstances", func(t *testing.T) {
        input := &route53.ListTrafficPolicyInstancesInput{}
        output := &route53.ListTrafficPolicyInstancesOutput{}

        mockClient.On("ListTrafficPolicyInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrafficPolicyInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrafficPolicyInstancesByHostedZone", func(t *testing.T) {
        input := &route53.ListTrafficPolicyInstancesByHostedZoneInput{}
        output := &route53.ListTrafficPolicyInstancesByHostedZoneOutput{}

        mockClient.On("ListTrafficPolicyInstancesByHostedZone", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrafficPolicyInstancesByHostedZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrafficPolicyInstancesByPolicy", func(t *testing.T) {
        input := &route53.ListTrafficPolicyInstancesByPolicyInput{}
        output := &route53.ListTrafficPolicyInstancesByPolicyOutput{}

        mockClient.On("ListTrafficPolicyInstancesByPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrafficPolicyInstancesByPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrafficPolicyVersions", func(t *testing.T) {
        input := &route53.ListTrafficPolicyVersionsInput{}
        output := &route53.ListTrafficPolicyVersionsOutput{}

        mockClient.On("ListTrafficPolicyVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrafficPolicyVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVPCAssociationAuthorizations", func(t *testing.T) {
        input := &route53.ListVPCAssociationAuthorizationsInput{}
        output := &route53.ListVPCAssociationAuthorizationsOutput{}

        mockClient.On("ListVPCAssociationAuthorizations", ctx, input).Return(output, nil)

        result, err := mockClient.ListVPCAssociationAuthorizations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestDNSAnswer", func(t *testing.T) {
        input := &route53.TestDNSAnswerInput{}
        output := &route53.TestDNSAnswerOutput{}

        mockClient.On("TestDNSAnswer", ctx, input).Return(output, nil)

        result, err := mockClient.TestDNSAnswer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHealthCheck", func(t *testing.T) {
        input := &route53.UpdateHealthCheckInput{}
        output := &route53.UpdateHealthCheckOutput{}

        mockClient.On("UpdateHealthCheck", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHealthCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHostedZoneComment", func(t *testing.T) {
        input := &route53.UpdateHostedZoneCommentInput{}
        output := &route53.UpdateHostedZoneCommentOutput{}

        mockClient.On("UpdateHostedZoneComment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHostedZoneComment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrafficPolicyComment", func(t *testing.T) {
        input := &route53.UpdateTrafficPolicyCommentInput{}
        output := &route53.UpdateTrafficPolicyCommentOutput{}

        mockClient.On("UpdateTrafficPolicyComment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrafficPolicyComment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrafficPolicyInstance", func(t *testing.T) {
        input := &route53.UpdateTrafficPolicyInstanceInput{}
        output := &route53.UpdateTrafficPolicyInstanceOutput{}

        mockClient.On("UpdateTrafficPolicyInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrafficPolicyInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}

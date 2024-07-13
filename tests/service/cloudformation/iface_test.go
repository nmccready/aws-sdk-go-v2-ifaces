package cloudformation_test

// tests for the s3 service interface for this ../../../service/cloudformation/iface.go

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudformation/cloudformation_iface"
	"github.com/stretchr/testify/assert"
)

func TestS3ServiceCanBeMocked(t *testing.T) {
	var iface cloudformation_iface.IClient
	iface = &cloudformation.Client{}
	assert.NotNil(t, iface, "not nil")
}

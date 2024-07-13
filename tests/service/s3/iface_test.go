package s3_test

// tests for the s3 service interface for this ../../../service/s3/iface.go

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3/s3_iface"
	"github.com/stretchr/testify/assert"
)

func TestS3ServiceCanBeMocked(t *testing.T) {
	var iface s3_iface.IClient
	iface = &s3.Client{}
	assert.NotNil(t, iface, "not nil")
}

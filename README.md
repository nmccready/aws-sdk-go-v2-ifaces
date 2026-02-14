# aws-sdk-go-v2-ifaces

[![Go Reference](https://pkg.go.dev/badge/github.com/nmccready/aws-sdk-go-v2-ifaces.svg)](https://pkg.go.dev/github.com/nmccready/aws-sdk-go-v2-ifaces)
[![CI](https://github.com/nmccready/aws-sdk-go-v2-ifaces/actions/workflows/tests.yml/badge.svg)](https://github.com/nmccready/aws-sdk-go-v2-ifaces/actions/workflows/tests.yml)
[![Weekly Sync](https://github.com/nmccready/aws-sdk-go-v2-ifaces/actions/workflows/weekly-sync.yml/badge.svg)](https://github.com/nmccready/aws-sdk-go-v2-ifaces/actions/workflows/weekly-sync.yml)
![Go Version](https://img.shields.io/github/go-mod/go-version/nmccready/aws-sdk-go-v2-ifaces)

**Auto-generated Go interfaces for every AWS SDK v2 service client. 413+ services. Updated weekly.**

## The Problem

The AWS SDK for Go v2 ships concrete struct clients with no interfaces, making it painful to mock AWS services in unit tests. This is a [long-standing community request](https://github.com/aws/aws-sdk-go-v2/issues/791) with 100+ upvotes that AWS has not addressed.

### Without interfaces (the painful way)

```go
// Your production code is tightly coupled to the concrete client
type S3Manager struct {
    client *s3.Client // ← concrete type, can't substitute in tests
}

// Testing requires either:
// 1. Hitting real AWS (slow, expensive, flaky)
// 2. Wrapping every method yourself (tedious boilerplate)
// 3. Using something like localstack (heavy setup)
```

### With this package ✅

```go
import (
    s3iface "github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3/s3_iface"
)

// Your production code accepts the interface
type S3Manager struct {
    client s3iface.IClient // ← interface, easily mockable
}

// The real s3.Client already satisfies this interface — zero changes needed
func NewS3Manager(client s3iface.IClient) *S3Manager {
    return &S3Manager{client: client}
}
```

## Quick Start

### Install

```bash
# Import only the services you need
go get github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3
go get github.com/nmccready/aws-sdk-go-v2-ifaces/service/dynamodb
go get github.com/nmccready/aws-sdk-go-v2-ifaces/service/sqs
```

### Use in tests

Each service provides a pre-generated [testify/mock](https://github.com/stretchr/testify) mock:

```go
package mypackage_test

import (
    "context"
    "testing"

    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3/mocks"
    "github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    // Set up expectations
    input := &s3.PutObjectInput{Bucket: aws.String("my-bucket"), Key: aws.String("test.txt")}
    output := &s3.PutObjectOutput{}
    mockClient.On("PutObject", ctx, input).Return(output, nil)

    // Use the mock in your code
    manager := NewS3Manager(mockClient)
    err := manager.Upload(ctx, "my-bucket", "test.txt", data)

    assert.NoError(t, err)
    mockClient.AssertExpectations(t)
}
```

## What's Included

For each of the **413+ AWS services**, this package generates:

| Path | Description |
|------|-------------|
| `service/<svc>/<svc>_iface/iface.go` | `IClient` interface matching every method on the SDK client |
| `service/<svc>/mocks/IClient.go` | Ready-to-use [testify/mock](https://github.com/stretchr/testify) implementation |

### Supported Services (highlights)

<details>
<summary>All 413+ services are supported — click to expand the full list</summary>

`accessanalyzer` · `account` · `acm` · `acmpca` · `amplify` · `apigateway` · `apigatewayv2` · `appconfig` · `appsync` · `athena` · `autoscaling` · `backup` · `batch` · `bedrock` · `bedrockruntime` · `cloudformation` · `cloudfront` · `cloudtrail` · `cloudwatch` · `cloudwatchlogs` · `codebuild` · `codedeploy` · `codepipeline` · `cognitoidentityprovider` · `comprehend` · `configservice` · `connect` · `costexplorer` · `databasemigrationservice` · `datazone` · `directconnect` · `dlm` · `docdb` · `dynamodb` · `dynamodbstreams` · `ebs` · `ec2` · `ecr` · `ecrpublic` · `ecs` · `efs` · `eks` · `elasticache` · `elasticbeanstalk` · `elasticloadbalancingv2` · `emr` · `emrserverless` · `eventbridge` · `firehose` · `fsx` · `glacier` · `glue` · `grafana` · `guardduty` · `health` · `iam` · `imagebuilder` · `inspector2` · `iot` · `kafka` · `kinesis` · `kms` · `lakeformation` · `lambda` · `lightsail` · `location` · `mediaconvert` · `memorydb` · `mq` · `neptune` · `networkfirewall` · `opensearch` · `organizations` · `personalize` · `pi` · `pinpoint` · `polly` · `quicksight` · `ram` · `rds` · `redshift` · `rekognition` · `resiliencehub` · `resourcegroups` · `route53` · `route53domains` · `route53resolver` · `s3` · `s3control` · `sagemaker` · `secretsmanager` · `securityhub` · `ses` · `sesv2` · `sfn` · `shield` · `sns` · `sqs` · `ssm` · `sso` · `ssoadmin` · `storagegateway` · `sts` · `support` · `textract` · `timestreaminfluxdb` · `transcribe` · `transfer` · `translate` · `wafv2` · `wellarchitected` · `workspaces` · `xray` ... and 300+ more

</details>

## Always Up to Date

A [weekly GitHub Actions workflow](https://github.com/nmccready/aws-sdk-go-v2-ifaces/actions/workflows/weekly-sync.yml) automatically:

1. Clones the latest `aws-sdk-go-v2` source
2. Regenerates all interfaces and mocks
3. Opens a PR if anything changed

This means interfaces stay in sync with the latest SDK releases — no manual maintenance needed.

## FAQ

### Why not just write your own interface?

You absolutely can. But with 413 services and methods like S3 having 90+ operations, writing and maintaining interfaces by hand is a lot of boilerplate. This package auto-generates all of it.

### Does the real SDK client satisfy these interfaces?

Yes. The `IClient` interface is generated directly from the SDK client's method signatures. `*s3.Client` (and every other service client) satisfies its corresponding `IClient` interface with zero adapter code.

### Why doesn't the official SDK provide interfaces?

It's been [requested since 2021](https://github.com/aws/aws-sdk-go-v2/issues/791) and has 100+ upvotes, but the AWS SDK team has not implemented it. This package fills that gap.

### How are the mocks generated?

Interfaces are extracted from the SDK source via shell scripts, then [mockery](https://github.com/vektra/mockery) generates the testify/mock implementations.

## Contributing

Contributions are welcome! This project uses:

- **Node.js** for orchestration scripts
- **Go 1.24+** for the generated code
- **[mockery](https://github.com/vektra/mockery)** for mock generation

### Local Development

```bash
# Install dependencies
npm install

# Regenerate all interfaces and mocks from latest SDK
npm run sync

# Run tests
npm test
```

The generation pipeline:
1. `npm run sync` clones `aws-sdk-go-v2` to `./tmp/`
2. Scripts extract client method signatures for each service
3. Interfaces are written to `service/<svc>/<svc>_iface/iface.go`
4. Mockery generates mocks to `service/<svc>/mocks/IClient.go`

## License

See [LICENSE](LICENSE) for details.

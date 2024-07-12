module github.com/nmccready/aws-sdk-go-v2-ifaces

go 1.20

require (
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.32.3
	github.com/aws/aws-sdk-go-v2/service/account v1.19.3
	github.com/aws/aws-sdk-go-v2/service/acm v1.28.4
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.33.3
	github.com/aws/aws-sdk-go-v2/service/amp v1.27.3
	github.com/aws/aws-sdk-go-v2/service/amplify v1.23.3
	github.com/aws/aws-sdk-go-v2/service/amplifybackend v1.25.3
	github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder v1.21.3
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.25.3
	github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi v1.21.3
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.22.3
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.31.3
	github.com/aws/aws-sdk-go-v2/service/appconfigdata v1.16.3
	github.com/aws/aws-sdk-go-v2/service/appfabric v1.9.3
	github.com/aws/aws-sdk-go-v2/service/appflow v1.43.3
	github.com/aws/aws-sdk-go-v2/service/appintegrations v1.27.3
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.30.4
	github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler v1.19.3
	github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice v1.26.3
	github.com/aws/aws-sdk-go-v2/service/applicationinsights v1.26.3
	github.com/aws/aws-sdk-go-v2/service/applicationsignals v1.2.3
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.27.3
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.30.3
	github.com/aws/aws-sdk-go-v2/service/appstream v1.36.3
	github.com/aws/aws-sdk-go-v2/service/appsync v1.34.3
	github.com/aws/aws-sdk-go-v2/service/apptest v1.2.3
	github.com/aws/aws-sdk-go-v2/service/arczonalshift v1.11.3
	github.com/aws/aws-sdk-go-v2/service/artifact v1.4.3
	github.com/aws/aws-sdk-go-v2/service/athena v1.44.3
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.35.3
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.43.3
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.22.3
	github.com/aws/aws-sdk-go-v2/service/b2bi v1.0.0-preview.35
	github.com/aws/aws-sdk-go-v2/service/backup v1.36.3
	github.com/aws/aws-sdk-go-v2/service/backupgateway v1.18.3
	github.com/aws/aws-sdk-go-v2/service/batch v1.43.0
	github.com/aws/aws-sdk-go-v2/service/bcmdataexports v1.5.3
	github.com/aws/aws-sdk-go-v2/service/bedrock v1.12.0
	github.com/aws/aws-sdk-go-v2/service/bedrockagent v1.16.0
	github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime v1.15.0
	github.com/aws/aws-sdk-go-v2/service/bedrockruntime v1.14.0
	github.com/aws/aws-sdk-go-v2/service/billingconductor v1.18.3
	github.com/aws/aws-sdk-go-v2/service/braket v1.29.3
	github.com/aws/aws-sdk-go-v2/service/budgets v1.25.3
	github.com/aws/aws-sdk-go-v2/service/chatbot v1.4.3
	github.com/aws/aws-sdk-go-v2/service/chime v1.32.3
	github.com/aws/aws-sdk-go-v2/service/chimesdkidentity v1.20.3
	github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines v1.18.3
	github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings v1.25.4
	github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging v1.24.3
	github.com/aws/aws-sdk-go-v2/service/chimesdkvoice v1.17.3
	github.com/aws/aws-sdk-go-v2/service/cleanrooms v1.14.3
	github.com/aws/aws-sdk-go-v2/service/cleanroomsml v1.6.3
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.26.3
	github.com/aws/aws-sdk-go-v2/service/cloudcontrol v1.20.3
	github.com/aws/aws-sdk-go-v2/service/clouddirectory v1.22.3
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.53.3
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.38.4
	github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore v1.6.3
	github.com/aws/aws-sdk-go-v2/service/cloudhsm v1.22.3
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.25.2
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v1.24.3
	github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain v1.21.3
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.42.3
	github.com/aws/aws-sdk-go-v2/service/cloudtraildata v1.9.3
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.40.3
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.25.3
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.37.3
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.30.3
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.40.3
	github.com/aws/aws-sdk-go-v2/service/codecatalyst v1.15.3
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.24.3
	github.com/aws/aws-sdk-go-v2/service/codeconnections v1.2.3
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.27.3
	github.com/aws/aws-sdk-go-v2/service/codeguruprofiler v1.22.3
	github.com/aws/aws-sdk-go-v2/service/codegurureviewer v1.27.3
	github.com/aws/aws-sdk-go-v2/service/codegurusecurity v1.10.3
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.30.3
	github.com/aws/aws-sdk-go-v2/service/codestar v1.23.3
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.27.3
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.24.3
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.25.5
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.41.4
	github.com/aws/aws-sdk-go-v2/service/cognitosync v1.21.3
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.33.3
	github.com/aws/aws-sdk-go-v2/service/comprehendmedical v1.24.3
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.37.3
	github.com/aws/aws-sdk-go-v2/service/configservice v1.48.3
	github.com/aws/aws-sdk-go-v2/service/connect v1.104.2
	github.com/aws/aws-sdk-go-v2/service/connectcampaigns v1.13.3
	github.com/aws/aws-sdk-go-v2/service/connectcases v1.19.3
	github.com/aws/aws-sdk-go-v2/service/connectcontactlens v1.23.3
	github.com/aws/aws-sdk-go-v2/service/connectparticipant v1.25.3
	github.com/aws/aws-sdk-go-v2/service/controlcatalog v1.2.3
	github.com/aws/aws-sdk-go-v2/service/controltower v1.16.3
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.26.3
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.40.3
	github.com/aws/aws-sdk-go-v2/service/costoptimizationhub v1.7.3
	github.com/aws/aws-sdk-go-v2/service/customerprofiles v1.39.3
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.40.3
	github.com/aws/aws-sdk-go-v2/service/databrew v1.31.3
	github.com/aws/aws-sdk-go-v2/service/dataexchange v1.30.3
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.23.3
	github.com/aws/aws-sdk-go-v2/service/datasync v1.40.3
	github.com/aws/aws-sdk-go-v2/service/datazone v1.13.2
	github.com/aws/aws-sdk-go-v2/service/dax v1.21.3
	github.com/aws/aws-sdk-go-v2/service/deadline v1.2.3
	github.com/aws/aws-sdk-go-v2/service/detective v1.29.3
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.25.2
	github.com/aws/aws-sdk-go-v2/service/devopsguru v1.32.3
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.27.4
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.27.3
	github.com/aws/aws-sdk-go-v2/service/dlm v1.26.3
	github.com/aws/aws-sdk-go-v2/service/docdb v1.36.3
	github.com/aws/aws-sdk-go-v2/service/docdbelastic v1.11.3
	github.com/aws/aws-sdk-go-v2/service/drs v1.28.3
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.34.3
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.22.3
	github.com/aws/aws-sdk-go-v2/service/ebs v1.25.3
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.170.0
	github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect v1.25.3
	github.com/aws/aws-sdk-go-v2/service/ecr v1.30.3
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.25.3
	github.com/aws/aws-sdk-go-v2/service/ecs v1.44.3
	github.com/aws/aws-sdk-go-v2/service/efs v1.31.3
	github.com/aws/aws-sdk-go-v2/service/eks v1.46.2
	github.com/aws/aws-sdk-go-v2/service/eksauth v1.5.3
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.40.3
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.26.2
	github.com/aws/aws-sdk-go-v2/service/elasticinference v1.21.3
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.26.3
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.33.3
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.30.3
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.25.3
	github.com/aws/aws-sdk-go-v2/service/emr v1.42.2
	github.com/aws/aws-sdk-go-v2/service/emrcontainers v1.30.3
	github.com/aws/aws-sdk-go-v2/service/emrserverless v1.23.3
	github.com/aws/aws-sdk-go-v2/service/entityresolution v1.12.3
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.33.3
	github.com/aws/aws-sdk-go-v2/service/evidently v1.21.3
	github.com/aws/aws-sdk-go-v2/service/finspace v1.26.3
	github.com/aws/aws-sdk-go-v2/service/finspacedata v1.26.3
	github.com/aws/aws-sdk-go-v2/service/firehose v1.31.3
	github.com/aws/aws-sdk-go-v2/service/fis v1.26.3
	github.com/aws/aws-sdk-go-v2/service/fms v1.35.3
	github.com/aws/aws-sdk-go-v2/service/forecast v1.34.3
	github.com/aws/aws-sdk-go-v2/service/forecastquery v1.22.3
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.33.3
	github.com/aws/aws-sdk-go-v2/service/freetier v1.5.3
	github.com/aws/aws-sdk-go-v2/service/fsx v1.47.2
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.33.3
	github.com/aws/aws-sdk-go-v2/service/glacier v1.24.3
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.27.0
	github.com/aws/aws-sdk-go-v2/service/glue v1.91.0
	github.com/aws/aws-sdk-go-v2/service/grafana v1.24.3
	github.com/aws/aws-sdk-go-v2/service/greengrass v1.25.3
	github.com/aws/aws-sdk-go-v2/service/greengrassv2 v1.33.3
	github.com/aws/aws-sdk-go-v2/service/groundstation v1.29.3
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.45.3
	github.com/aws/aws-sdk-go-v2/service/health v1.26.3
	github.com/aws/aws-sdk-go-v2/service/healthlake v1.26.3
	github.com/aws/aws-sdk-go-v2/service/iam v1.34.3
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.25.3
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.35.3
	github.com/aws/aws-sdk-go-v2/service/inspector v1.23.3
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.28.3
	github.com/aws/aws-sdk-go-v2/service/inspectorscan v1.5.3
	github.com/aws/aws-sdk-go-v2/service/internetmonitor v1.16.3
	github.com/aws/aws-sdk-go-v2/service/iot v1.55.3
	github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice v1.21.3
	github.com/aws/aws-sdk-go-v2/service/iot1clickprojects v1.21.3
	github.com/aws/aws-sdk-go-v2/service/iotanalytics v1.24.3
	github.com/aws/aws-sdk-go-v2/service/iotdataplane v1.24.3
	github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor v1.28.3
	github.com/aws/aws-sdk-go-v2/service/iotevents v1.25.3
	github.com/aws/aws-sdk-go-v2/service/ioteventsdata v1.22.3
	github.com/aws/aws-sdk-go-v2/service/iotfleethub v1.22.3
	github.com/aws/aws-sdk-go-v2/service/iotfleetwise v1.17.3
	github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane v1.21.3
	github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling v1.25.3
	github.com/aws/aws-sdk-go-v2/service/iotsitewise v1.39.3
	github.com/aws/aws-sdk-go-v2/service/iotthingsgraph v1.23.3
	github.com/aws/aws-sdk-go-v2/service/iottwinmaker v1.22.3
	github.com/aws/aws-sdk-go-v2/service/iotwireless v1.42.3
	github.com/aws/aws-sdk-go-v2/service/ivs v1.37.3
	github.com/aws/aws-sdk-go-v2/service/ivschat v1.14.3
	github.com/aws/aws-sdk-go-v2/service/ivsrealtime v1.16.3
	github.com/aws/aws-sdk-go-v2/service/kafka v1.35.3
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.19.3
	github.com/aws/aws-sdk-go-v2/service/kendra v1.52.3
	github.com/aws/aws-sdk-go-v2/service/kendraranking v1.9.3
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.12.3
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.29.3
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.23.3
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.28.2
	github.com/aws/aws-sdk-go-v2/service/kinesisvideo v1.25.3
	github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia v1.25.3
	github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia v1.20.3
	github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling v1.21.3
	github.com/aws/aws-sdk-go-v2/service/kinesisvideowebrtcstorage v1.11.3
	github.com/aws/aws-sdk-go-v2/service/kms v1.35.3
	github.com/aws/aws-sdk-go-v2/service/lakeformation v1.35.3
	github.com/aws/aws-sdk-go-v2/service/lambda v1.56.3
	github.com/aws/aws-sdk-go-v2/service/launchwizard v1.6.3
	github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice v1.26.3
	github.com/aws/aws-sdk-go-v2/service/lexmodelsv2 v1.45.3
	github.com/aws/aws-sdk-go-v2/service/lexruntimeservice v1.22.3
	github.com/aws/aws-sdk-go-v2/service/lexruntimev2 v1.27.3
	github.com/aws/aws-sdk-go-v2/service/licensemanager v1.27.3
	github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions v1.12.0
	github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions v1.11.3
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.40.3
	github.com/aws/aws-sdk-go-v2/service/location v1.40.3
	github.com/aws/aws-sdk-go-v2/service/lookoutequipment v1.28.3
	github.com/aws/aws-sdk-go-v2/service/lookoutmetrics v1.29.3
	github.com/aws/aws-sdk-go-v2/service/lookoutvision v1.25.3
	github.com/aws/aws-sdk-go-v2/service/m2 v1.15.3
	github.com/aws/aws-sdk-go-v2/service/machinelearning v1.26.3
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.41.3
	github.com/aws/aws-sdk-go-v2/service/mailmanager v1.2.3
	github.com/aws/aws-sdk-go-v2/service/managedblockchain v1.24.3
	github.com/aws/aws-sdk-go-v2/service/managedblockchainquery v1.14.3
	github.com/aws/aws-sdk-go-v2/service/marketplaceagreement v1.4.3
	github.com/aws/aws-sdk-go-v2/service/marketplacecatalog v1.28.3
	github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics v1.22.3
	github.com/aws/aws-sdk-go-v2/service/marketplacedeployment v1.4.3
	github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice v1.23.3
	github.com/aws/aws-sdk-go-v2/service/marketplacemetering v1.23.3
	github.com/aws/aws-sdk-go-v2/service/mediaconnect v1.32.0
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.57.3
	github.com/aws/aws-sdk-go-v2/service/medialive v1.54.3
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.32.3
	github.com/aws/aws-sdk-go-v2/service/mediapackagev2 v1.14.3
	github.com/aws/aws-sdk-go-v2/service/mediapackagevod v1.32.3
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.22.3
	github.com/aws/aws-sdk-go-v2/service/mediastoredata v1.22.3
	github.com/aws/aws-sdk-go-v2/service/mediatailor v1.40.3
	github.com/aws/aws-sdk-go-v2/service/medicalimaging v1.11.3
	github.com/aws/aws-sdk-go-v2/service/memorydb v1.21.3
	github.com/aws/aws-sdk-go-v2/service/mgn v1.30.3
	github.com/aws/aws-sdk-go-v2/service/migrationhub v1.22.3
	github.com/aws/aws-sdk-go-v2/service/migrationhubconfig v1.23.3
	github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator v1.11.3
	github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces v1.18.3
	github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy v1.19.3
	github.com/aws/aws-sdk-go-v2/service/mobile v1.21.3
	github.com/aws/aws-sdk-go-v2/service/mq v1.25.3
	github.com/aws/aws-sdk-go-v2/service/mturk v1.23.3
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.29.3
	github.com/aws/aws-sdk-go-v2/service/neptune v1.33.3
	github.com/aws/aws-sdk-go-v2/service/neptunedata v1.7.3
	github.com/aws/aws-sdk-go-v2/service/neptunegraph v1.10.3
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.40.3
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.29.3
	github.com/aws/aws-sdk-go-v2/service/networkmonitor v1.5.3
	github.com/aws/aws-sdk-go-v2/service/nimble v1.26.3
	github.com/aws/aws-sdk-go-v2/service/oam v1.13.3
	github.com/aws/aws-sdk-go-v2/service/omics v1.23.3
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.39.2
	github.com/aws/aws-sdk-go-v2/service/opensearchserverless v1.13.3
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.24.3
	github.com/aws/aws-sdk-go-v2/service/opsworkscm v1.25.3
	github.com/aws/aws-sdk-go-v2/service/organizations v1.30.2
	github.com/aws/aws-sdk-go-v2/service/osis v1.12.3
	github.com/aws/aws-sdk-go-v2/service/outposts v1.41.3
	github.com/aws/aws-sdk-go-v2/service/panorama v1.20.3
	github.com/aws/aws-sdk-go-v2/service/paymentcryptography v1.12.3
	github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata v1.12.2
	github.com/aws/aws-sdk-go-v2/service/pcaconnectorad v1.7.3
	github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep v1.2.3
	github.com/aws/aws-sdk-go-v2/service/personalize v1.36.3
	github.com/aws/aws-sdk-go-v2/service/personalizeevents v1.23.3
	github.com/aws/aws-sdk-go-v2/service/personalizeruntime v1.25.3
	github.com/aws/aws-sdk-go-v2/service/pi v1.27.4
	github.com/aws/aws-sdk-go-v2/service/pinpoint v1.32.3
	github.com/aws/aws-sdk-go-v2/service/pinpointemail v1.21.3
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice v1.21.3
	github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2 v1.12.4
	github.com/aws/aws-sdk-go-v2/service/pipes v1.14.3
	github.com/aws/aws-sdk-go-v2/service/polly v1.42.3
	github.com/aws/aws-sdk-go-v2/service/pricing v1.30.3
	github.com/aws/aws-sdk-go-v2/service/privatenetworks v1.11.3
	github.com/aws/aws-sdk-go-v2/service/proton v1.31.3
	github.com/aws/aws-sdk-go-v2/service/qapps v1.0.2
	github.com/aws/aws-sdk-go-v2/service/qbusiness v1.10.2
	github.com/aws/aws-sdk-go-v2/service/qconnect v1.9.3
	github.com/aws/aws-sdk-go-v2/service/qldb v1.23.3
	github.com/aws/aws-sdk-go-v2/service/qldbsession v1.23.3
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.67.3
	github.com/aws/aws-sdk-go-v2/service/ram v1.27.3
	github.com/aws/aws-sdk-go-v2/service/rbin v1.18.3
	github.com/aws/aws-sdk-go-v2/service/rds v1.81.4
	github.com/aws/aws-sdk-go-v2/service/rdsdata v1.23.3
	github.com/aws/aws-sdk-go-v2/service/redshift v1.46.4
	github.com/aws/aws-sdk-go-v2/service/redshiftdata v1.27.3
	github.com/aws/aws-sdk-go-v2/service/redshiftserverless v1.20.3
	github.com/aws/aws-sdk-go-v2/service/rekognition v1.43.2
	github.com/aws/aws-sdk-go-v2/service/repostspace v1.5.3
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.23.3
	github.com/aws/aws-sdk-go-v2/service/resourceexplorer2 v1.12.3
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.24.3
	github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi v1.23.3
	github.com/aws/aws-sdk-go-v2/service/robomaker v1.28.3
	github.com/aws/aws-sdk-go-v2/service/rolesanywhere v1.13.3
	github.com/aws/aws-sdk-go-v2/service/route53 v1.42.3
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.25.3
	github.com/aws/aws-sdk-go-v2/service/route53profiles v1.2.3
	github.com/aws/aws-sdk-go-v2/service/route53recoverycluster v1.21.3
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.23.3
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.19.3
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.30.3
	github.com/aws/aws-sdk-go-v2/service/rum v1.19.3
	github.com/aws/aws-sdk-go-v2/service/s3 v1.58.2
	github.com/aws/aws-sdk-go-v2/service/s3control v1.46.3
	github.com/aws/aws-sdk-go-v2/service/s3outposts v1.26.3
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.150.2
	github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime v1.25.3
	github.com/aws/aws-sdk-go-v2/service/sagemakeredge v1.23.3
	github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime v1.27.3
	github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial v1.12.3
	github.com/aws/aws-sdk-go-v2/service/sagemakermetrics v1.10.3
	github.com/aws/aws-sdk-go-v2/service/sagemakerruntime v1.29.3
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.21.3
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.10.3
	github.com/aws/aws-sdk-go-v2/service/schemas v1.26.3
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.32.3
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.51.3
	github.com/aws/aws-sdk-go-v2/service/securitylake v1.16.3
	github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository v1.22.3
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.30.3
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.28.3
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.31.3
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.23.3
	github.com/aws/aws-sdk-go-v2/service/ses v1.25.2
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.32.3
	github.com/aws/aws-sdk-go-v2/service/sfn v1.29.3
	github.com/aws/aws-sdk-go-v2/service/shield v1.27.3
	github.com/aws/aws-sdk-go-v2/service/signer v1.24.3
	github.com/aws/aws-sdk-go-v2/service/simspaceweaver v1.12.3
	github.com/aws/aws-sdk-go-v2/service/sms v1.22.3
	github.com/aws/aws-sdk-go-v2/service/snowball v1.28.3
	github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement v1.18.3
	github.com/aws/aws-sdk-go-v2/service/sns v1.31.3
	github.com/aws/aws-sdk-go-v2/service/sqs v1.34.3
	github.com/aws/aws-sdk-go-v2/service/ssm v1.52.3
	github.com/aws/aws-sdk-go-v2/service/ssmcontacts v1.24.3
	github.com/aws/aws-sdk-go-v2/service/ssmincidents v1.32.3
	github.com/aws/aws-sdk-go-v2/service/ssmsap v1.15.3
	github.com/aws/aws-sdk-go-v2/service/sso v1.22.3
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.27.3
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.26.4
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.31.3
	github.com/aws/aws-sdk-go-v2/service/sts v1.30.3
	github.com/aws/aws-sdk-go-v2/service/supplychain v1.5.3
	github.com/aws/aws-sdk-go-v2/service/support v1.24.3
	github.com/aws/aws-sdk-go-v2/service/supportapp v1.11.3
	github.com/aws/aws-sdk-go-v2/service/swf v1.25.3
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.26.3
	github.com/aws/aws-sdk-go-v2/service/taxsettings v1.2.3
	github.com/aws/aws-sdk-go-v2/service/textract v1.32.3
	github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb v1.2.3
	github.com/aws/aws-sdk-go-v2/service/timestreamquery v1.25.3
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.27.3
	github.com/aws/aws-sdk-go-v2/service/tnb v1.10.3
	github.com/aws/aws-sdk-go-v2/service/transcribe v1.39.3
	github.com/aws/aws-sdk-go-v2/service/transcribestreaming v1.19.3
	github.com/aws/aws-sdk-go-v2/service/transfer v1.50.3
	github.com/aws/aws-sdk-go-v2/service/translate v1.26.4
	github.com/aws/aws-sdk-go-v2/service/trustedadvisor v1.6.3
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.17.3
	github.com/aws/aws-sdk-go-v2/service/voiceid v1.22.3
	github.com/aws/aws-sdk-go-v2/service/vpclattice v1.10.3
	github.com/aws/aws-sdk-go-v2/service/waf v1.23.3
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.23.3
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.51.4
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.32.3
	github.com/aws/aws-sdk-go-v2/service/wisdom v1.25.3
	github.com/aws/aws-sdk-go-v2/service/workdocs v1.23.3
	github.com/aws/aws-sdk-go-v2/service/worklink v1.22.3
	github.com/aws/aws-sdk-go-v2/service/workmail v1.27.3
	github.com/aws/aws-sdk-go-v2/service/workmailmessageflow v1.21.3
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.44.2
	github.com/aws/aws-sdk-go-v2/service/workspacesthinclient v1.8.3
	github.com/aws/aws-sdk-go-v2/service/workspacesweb v1.21.3
	github.com/aws/aws-sdk-go-v2/service/xray v1.27.3
	github.com/nmccready/go-debug v0.5.0
	github.com/stretchr/testify v1.6.1
)

require (
	github.com/aws/aws-sdk-go-v2 v1.30.3 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.9.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.15 // indirect
	github.com/aws/smithy-go v1.20.3 // indirect
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/nmccready/colorjson v0.1.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)


package ioteventsdata

import (
    "github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
)

// IIoteventsdata defines the interface for ioteventsdata
type IIoteventsdata interface {
 Options() Options 
 BatchAcknowledgeAlarm(ctx context.Context, params *BatchAcknowledgeAlarmInput, optFns ...func(*Options)) (*BatchAcknowledgeAlarmOutput, error) 
 BatchDeleteDetector(ctx context.Context, params *BatchDeleteDetectorInput, optFns ...func(*Options)) (*BatchDeleteDetectorOutput, error) 
 BatchDisableAlarm(ctx context.Context, params *BatchDisableAlarmInput, optFns ...func(*Options)) (*BatchDisableAlarmOutput, error) 
 BatchEnableAlarm(ctx context.Context, params *BatchEnableAlarmInput, optFns ...func(*Options)) (*BatchEnableAlarmOutput, error) 
 BatchPutMessage(ctx context.Context, params *BatchPutMessageInput, optFns ...func(*Options)) (*BatchPutMessageOutput, error) 
 BatchResetAlarm(ctx context.Context, params *BatchResetAlarmInput, optFns ...func(*Options)) (*BatchResetAlarmOutput, error) 
 BatchSnoozeAlarm(ctx context.Context, params *BatchSnoozeAlarmInput, optFns ...func(*Options)) (*BatchSnoozeAlarmOutput, error) 
 BatchUpdateDetector(ctx context.Context, params *BatchUpdateDetectorInput, optFns ...func(*Options)) (*BatchUpdateDetectorOutput, error) 
 DescribeAlarm(ctx context.Context, params *DescribeAlarmInput, optFns ...func(*Options)) (*DescribeAlarmOutput, error) 
 DescribeDetector(ctx context.Context, params *DescribeDetectorInput, optFns ...func(*Options)) (*DescribeDetectorOutput, error) 
 ListAlarms(ctx context.Context, params *ListAlarmsInput, optFns ...func(*Options)) (*ListAlarmsOutput, error) 
 ListDetectors(ctx context.Context, params *ListDetectorsInput, optFns ...func(*Options)) (*ListDetectorsOutput, error) 
}

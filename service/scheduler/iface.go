
package scheduler

import (
    "github.com/aws/aws-sdk-go-v2/service/scheduler"
)

// IClient defines the interface for scheduler
type IClient interface {
 Options() Options 
 CreateSchedule(ctx context.Context, params *CreateScheduleInput, optFns ...func(*Options)) (*CreateScheduleOutput, error) 
 CreateScheduleGroup(ctx context.Context, params *CreateScheduleGroupInput, optFns ...func(*Options)) (*CreateScheduleGroupOutput, error) 
 DeleteSchedule(ctx context.Context, params *DeleteScheduleInput, optFns ...func(*Options)) (*DeleteScheduleOutput, error) 
 DeleteScheduleGroup(ctx context.Context, params *DeleteScheduleGroupInput, optFns ...func(*Options)) (*DeleteScheduleGroupOutput, error) 
 GetSchedule(ctx context.Context, params *GetScheduleInput, optFns ...func(*Options)) (*GetScheduleOutput, error) 
 GetScheduleGroup(ctx context.Context, params *GetScheduleGroupInput, optFns ...func(*Options)) (*GetScheduleGroupOutput, error) 
 ListScheduleGroups(ctx context.Context, params *ListScheduleGroupsInput, optFns ...func(*Options)) (*ListScheduleGroupsOutput, error) 
 ListSchedules(ctx context.Context, params *ListSchedulesInput, optFns ...func(*Options)) (*ListSchedulesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateSchedule(ctx context.Context, params *UpdateScheduleInput, optFns ...func(*Options)) (*UpdateScheduleOutput, error) 
}

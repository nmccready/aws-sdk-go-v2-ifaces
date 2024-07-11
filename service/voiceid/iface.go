
package voiceid

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/voiceid"
)

// IClient defines the interface for voiceid
type IClient interface {
 Options() Options 
 AssociateFraudster(ctx context.Context, params *AssociateFraudsterInput, optFns ...func(*Options)) (*AssociateFraudsterOutput, error) 
 CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) 
 CreateWatchlist(ctx context.Context, params *CreateWatchlistInput, optFns ...func(*Options)) (*CreateWatchlistOutput, error) 
 DeleteDomain(ctx context.Context, params *DeleteDomainInput, optFns ...func(*Options)) (*DeleteDomainOutput, error) 
 DeleteFraudster(ctx context.Context, params *DeleteFraudsterInput, optFns ...func(*Options)) (*DeleteFraudsterOutput, error) 
 DeleteSpeaker(ctx context.Context, params *DeleteSpeakerInput, optFns ...func(*Options)) (*DeleteSpeakerOutput, error) 
 DeleteWatchlist(ctx context.Context, params *DeleteWatchlistInput, optFns ...func(*Options)) (*DeleteWatchlistOutput, error) 
 DescribeDomain(ctx context.Context, params *DescribeDomainInput, optFns ...func(*Options)) (*DescribeDomainOutput, error) 
 DescribeFraudster(ctx context.Context, params *DescribeFraudsterInput, optFns ...func(*Options)) (*DescribeFraudsterOutput, error) 
 DescribeFraudsterRegistrationJob(ctx context.Context, params *DescribeFraudsterRegistrationJobInput, optFns ...func(*Options)) (*DescribeFraudsterRegistrationJobOutput, error) 
 DescribeSpeaker(ctx context.Context, params *DescribeSpeakerInput, optFns ...func(*Options)) (*DescribeSpeakerOutput, error) 
 DescribeSpeakerEnrollmentJob(ctx context.Context, params *DescribeSpeakerEnrollmentJobInput, optFns ...func(*Options)) (*DescribeSpeakerEnrollmentJobOutput, error) 
 DescribeWatchlist(ctx context.Context, params *DescribeWatchlistInput, optFns ...func(*Options)) (*DescribeWatchlistOutput, error) 
 DisassociateFraudster(ctx context.Context, params *DisassociateFraudsterInput, optFns ...func(*Options)) (*DisassociateFraudsterOutput, error) 
 EvaluateSession(ctx context.Context, params *EvaluateSessionInput, optFns ...func(*Options)) (*EvaluateSessionOutput, error) 
 ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) 
 ListFraudsterRegistrationJobs(ctx context.Context, params *ListFraudsterRegistrationJobsInput, optFns ...func(*Options)) (*ListFraudsterRegistrationJobsOutput, error) 
 ListFraudsters(ctx context.Context, params *ListFraudstersInput, optFns ...func(*Options)) (*ListFraudstersOutput, error) 
 ListSpeakerEnrollmentJobs(ctx context.Context, params *ListSpeakerEnrollmentJobsInput, optFns ...func(*Options)) (*ListSpeakerEnrollmentJobsOutput, error) 
 ListSpeakers(ctx context.Context, params *ListSpeakersInput, optFns ...func(*Options)) (*ListSpeakersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWatchlists(ctx context.Context, params *ListWatchlistsInput, optFns ...func(*Options)) (*ListWatchlistsOutput, error) 
 OptOutSpeaker(ctx context.Context, params *OptOutSpeakerInput, optFns ...func(*Options)) (*OptOutSpeakerOutput, error) 
 StartFraudsterRegistrationJob(ctx context.Context, params *StartFraudsterRegistrationJobInput, optFns ...func(*Options)) (*StartFraudsterRegistrationJobOutput, error) 
 StartSpeakerEnrollmentJob(ctx context.Context, params *StartSpeakerEnrollmentJobInput, optFns ...func(*Options)) (*StartSpeakerEnrollmentJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDomain(ctx context.Context, params *UpdateDomainInput, optFns ...func(*Options)) (*UpdateDomainOutput, error) 
 UpdateWatchlist(ctx context.Context, params *UpdateWatchlistInput, optFns ...func(*Options)) (*UpdateWatchlistOutput, error) 
}

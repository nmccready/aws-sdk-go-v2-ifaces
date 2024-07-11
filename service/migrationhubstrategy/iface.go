
package migrationhubstrategy

import (
    "github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy"
)

// IMigrationhubstrategy defines the interface for migrationhubstrategy
type IMigrationhubstrategy interface {
 Options() Options 
 GetApplicationComponentDetails(ctx context.Context, params *GetApplicationComponentDetailsInput, optFns ...func(*Options)) (*GetApplicationComponentDetailsOutput, error) 
 GetApplicationComponentStrategies(ctx context.Context, params *GetApplicationComponentStrategiesInput, optFns ...func(*Options)) (*GetApplicationComponentStrategiesOutput, error) 
 GetAssessment(ctx context.Context, params *GetAssessmentInput, optFns ...func(*Options)) (*GetAssessmentOutput, error) 
 GetImportFileTask(ctx context.Context, params *GetImportFileTaskInput, optFns ...func(*Options)) (*GetImportFileTaskOutput, error) 
 GetLatestAssessmentId(ctx context.Context, params *GetLatestAssessmentIdInput, optFns ...func(*Options)) (*GetLatestAssessmentIdOutput, error) 
 GetPortfolioPreferences(ctx context.Context, params *GetPortfolioPreferencesInput, optFns ...func(*Options)) (*GetPortfolioPreferencesOutput, error) 
 GetPortfolioSummary(ctx context.Context, params *GetPortfolioSummaryInput, optFns ...func(*Options)) (*GetPortfolioSummaryOutput, error) 
 GetRecommendationReportDetails(ctx context.Context, params *GetRecommendationReportDetailsInput, optFns ...func(*Options)) (*GetRecommendationReportDetailsOutput, error) 
 GetServerDetails(ctx context.Context, params *GetServerDetailsInput, optFns ...func(*Options)) (*GetServerDetailsOutput, error) 
 GetServerStrategies(ctx context.Context, params *GetServerStrategiesInput, optFns ...func(*Options)) (*GetServerStrategiesOutput, error) 
 ListAnalyzableServers(ctx context.Context, params *ListAnalyzableServersInput, optFns ...func(*Options)) (*ListAnalyzableServersOutput, error) 
 ListApplicationComponents(ctx context.Context, params *ListApplicationComponentsInput, optFns ...func(*Options)) (*ListApplicationComponentsOutput, error) 
 ListCollectors(ctx context.Context, params *ListCollectorsInput, optFns ...func(*Options)) (*ListCollectorsOutput, error) 
 ListImportFileTask(ctx context.Context, params *ListImportFileTaskInput, optFns ...func(*Options)) (*ListImportFileTaskOutput, error) 
 ListServers(ctx context.Context, params *ListServersInput, optFns ...func(*Options)) (*ListServersOutput, error) 
 PutPortfolioPreferences(ctx context.Context, params *PutPortfolioPreferencesInput, optFns ...func(*Options)) (*PutPortfolioPreferencesOutput, error) 
 StartAssessment(ctx context.Context, params *StartAssessmentInput, optFns ...func(*Options)) (*StartAssessmentOutput, error) 
 StartImportFileTask(ctx context.Context, params *StartImportFileTaskInput, optFns ...func(*Options)) (*StartImportFileTaskOutput, error) 
 StartRecommendationReportGeneration(ctx context.Context, params *StartRecommendationReportGenerationInput, optFns ...func(*Options)) (*StartRecommendationReportGenerationOutput, error) 
 StopAssessment(ctx context.Context, params *StopAssessmentInput, optFns ...func(*Options)) (*StopAssessmentOutput, error) 
 UpdateApplicationComponentConfig(ctx context.Context, params *UpdateApplicationComponentConfigInput, optFns ...func(*Options)) (*UpdateApplicationComponentConfigOutput, error) 
 UpdateServerConfig(ctx context.Context, params *UpdateServerConfigInput, optFns ...func(*Options)) (*UpdateServerConfigOutput, error) 
}

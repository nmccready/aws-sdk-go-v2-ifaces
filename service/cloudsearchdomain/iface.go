
package cloudsearchdomain

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
)

// ICloudsearchdomain defines the interface for cloudsearchdomain
type ICloudsearchdomain interface {
 Options() Options 
 Search(ctx context.Context, params *SearchInput, optFns ...func(*Options)) (*SearchOutput, error) 
 Suggest(ctx context.Context, params *SuggestInput, optFns ...func(*Options)) (*SuggestOutput, error) 
 UploadDocuments(ctx context.Context, params *UploadDocumentsInput, optFns ...func(*Options)) (*UploadDocumentsOutput, error) 
}

package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SimpleCloudfrontDistributionProps struct {
	S3Bucket awss3.Bucket
}

func NewSimpleCloudfrontDistribution(scope constructs.Construct, id *string, props SimpleCloudfrontDistributionProps) awscloudfront.Distribution {
	return awscloudfront.NewDistribution(scope, id, &awscloudfront.DistributionProps{
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			AllowedMethods:       awscloudfront.AllowedMethods_ALLOW_GET_HEAD(),
			CachedMethods:        awscloudfront.CachedMethods_CACHE_GET_HEAD(),
			CachePolicy:          awscloudfront.CachePolicy_CACHING_OPTIMIZED(),
			Compress:             jsii.Bool(true),
			ViewerProtocolPolicy: awscloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
			Origin:               awscloudfrontorigins.NewS3Origin(props.S3Bucket, nil),
		},
		Certificate: nil,
		DomainNames: nil,
		EnableIpv6:  jsii.Bool(true),
	})
}

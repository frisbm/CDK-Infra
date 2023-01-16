package constructs

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdklabs/cdk-nag-go/cdknag/v2"
)

type SimpleCloudfrontDistributionProps struct {
	S3Bucket  awss3.Bucket
	LogBucket awss3.Bucket
}

func NewSimpleCloudfrontDistribution(scope constructs.Construct, id *string, props SimpleCloudfrontDistributionProps) awscloudfront.Distribution {
	originAccessIdentityId := jsii.String(fmt.Sprintf("%v-access", *id))
	originAccessIdentity := awscloudfront.NewOriginAccessIdentity(scope, originAccessIdentityId, &awscloudfront.OriginAccessIdentityProps{
		Comment: originAccessIdentityId,
	})

	props.S3Bucket.GrantRead(originAccessIdentity, nil)
	cloudfrontDistribution := awscloudfront.NewDistribution(scope, id, &awscloudfront.DistributionProps{
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			AllowedMethods:       awscloudfront.AllowedMethods_ALLOW_GET_HEAD(),
			CachedMethods:        awscloudfront.CachedMethods_CACHE_GET_HEAD(),
			CachePolicy:          awscloudfront.CachePolicy_CACHING_OPTIMIZED(),
			Compress:             jsii.Bool(true),
			ViewerProtocolPolicy: awscloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
			Origin: awscloudfrontorigins.NewS3Origin(props.S3Bucket, &awscloudfrontorigins.S3OriginProps{
				OriginAccessIdentity: nil,
			}),
		},
		Certificate:   nil,
		DomainNames:   nil,
		HttpVersion:   awscloudfront.HttpVersion_HTTP2_AND_3,
		EnableIpv6:    jsii.Bool(true),
		EnableLogging: jsii.Bool(true),
		LogFilePrefix: jsii.String("distribution-access-logs"),
		LogBucket:     props.LogBucket,
	})

	cdknag.NagSuppressions_AddResourceSuppressions(cloudfrontDistribution, &[]*cdknag.NagPackSuppression{
		{
			Id:     jsii.String("AwsSolutions-CFR4"),
			Reason: jsii.String("No certificate at the moment"),
		},
	}, nil)

	return cloudfrontDistribution
}

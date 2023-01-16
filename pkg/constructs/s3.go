package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SimpleS3BucketProps struct {
	BucketName *string
	LogBucket  awss3.Bucket
}

func NewSimpleS3Bucket(scope constructs.Construct, id *string, props SimpleS3BucketProps) awss3.Bucket {
	return awss3.NewBucket(scope, id, &awss3.BucketProps{
		BucketName:             props.BucketName,
		Versioned:              jsii.Bool(true),
		AccessControl:          awss3.BucketAccessControl_LOG_DELIVERY_WRITE,
		BlockPublicAccess:      awss3.BlockPublicAccess_BLOCK_ALL(),
		ServerAccessLogsBucket: props.LogBucket,
		ServerAccessLogsPrefix: jsii.String("bucket-access-logs"),
		AutoDeleteObjects:      jsii.Bool(true),
		RemovalPolicy:          awscdk.RemovalPolicy_DESTROY,
		EnforceSSL:             jsii.Bool(true),
		Encryption:             awss3.BucketEncryption_S3_MANAGED,
	})
}

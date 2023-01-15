package constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SimpleS3BucketProps struct {
	BucketName *string
}

func NewSimpleS3Bucket(scope constructs.Construct, id *string, props SimpleS3BucketProps) awss3.Bucket {
	return awss3.NewBucket(scope, id, &awss3.BucketProps{
		AccessControl:     awss3.BucketAccessControl_PUBLIC_READ,
		AutoDeleteObjects: jsii.Bool(true),
		BucketName:        props.BucketName,
		Encryption:        awss3.BucketEncryption_S3_MANAGED,
		PublicReadAccess:  jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	})
}

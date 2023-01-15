package stacks

import (
	"github.com/MatthewFrisby/aws-cdk-infra/pkg/constructs"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	awsconstructs "github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FileDistributionStackProps struct {
	awscdk.StackProps
	BucketName *string
}

func NewFileDistributionStack(scope awsconstructs.Construct, id string, props *FileDistributionStackProps) awscdk.Stack {
	var stackProps awscdk.StackProps
	if props != nil {
		stackProps = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &stackProps)

	s3Bucket := constructs.NewSimpleS3Bucket(stack, jsii.String("Bucket"), constructs.SimpleS3BucketProps{
		BucketName: props.BucketName,
	})

	_ = constructs.NewSimpleCloudfrontDistribution(stack, jsii.String("Distribution"), constructs.SimpleCloudfrontDistributionProps{
		S3Bucket: s3Bucket,
	})

	return stack
}

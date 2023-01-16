package stacks

import (
	"fmt"
	"github.com/MatthewFrisby/aws-cdk-infra/pkg/constructs"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	awsconstructs "github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"strings"
)

type CDNStackProps struct {
	awscdk.StackProps
	BucketName *string
}

func NewCDNStack(scope awsconstructs.Construct, id string, props *CDNStackProps) awscdk.Stack {
	var stackProps awscdk.StackProps
	if props != nil {
		stackProps = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &stackProps)

	logBucket := constructs.NewSimpleS3Bucket(stack, jsii.String("LogBucket"), constructs.SimpleS3BucketProps{
		BucketName: jsii.String(fmt.Sprintf("%v-logbucket", strings.ToLower(id))),
	})

	s3Bucket := constructs.NewSimpleS3Bucket(stack, jsii.String("Bucket"), constructs.SimpleS3BucketProps{
		BucketName: props.BucketName,
		LogBucket:  logBucket,
	})

	_ = constructs.NewSimpleCloudfrontDistribution(stack, jsii.String("Distribution"), constructs.SimpleCloudfrontDistributionProps{
		S3Bucket:  s3Bucket,
		LogBucket: logBucket,
	})

	return stack
}

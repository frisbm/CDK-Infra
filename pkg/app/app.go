package app

import (
	"github.com/MatthewFrisby/aws-cdk-infra/pkg/config"
	"github.com/MatthewFrisby/aws-cdk-infra/pkg/stacks"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

type App struct {
	props  *awscdk.AppProps
	env    *awscdk.Environment
	config *config.Config
}

func NewApp(configPath string, props *awscdk.AppProps) App {
	newConfig := config.NewConfig(configPath)
	return App{
		props:  props,
		config: newConfig,
		env: &awscdk.Environment{
			Account: newConfig.Get("cdk.account"),
			Region:  newConfig.Get("cdk.region"),
		},
	}
}

func (a *App) SynthAllStacks() {
	defer jsii.Close()

	app := awscdk.NewApp(a.props)

	stacks.NewFileDistributionStack(app, "FileDistributionStack", &stacks.FileDistributionStackProps{
		StackProps: awscdk.StackProps{
			Env: a.env,
		},
		BucketName: a.config.Get("file-distribution-stack.bucket-name"),
	})

	app.Synth(nil)
}

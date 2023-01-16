package app

import (
	"github.com/MatthewFrisby/aws-cdk-infra/pkg/config"
	"github.com/MatthewFrisby/aws-cdk-infra/pkg/stacks"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdklabs/cdk-nag-go/cdknag/v2"
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

var nagRuleGroups = []awscdk.IAspect{
	cdknag.NewAwsSolutionsChecks(nil),
}

func addNagRuleGroups(app awscdk.App) {
	aspects := awscdk.Aspects_Of(app)
	for _, nagRuleGroup := range nagRuleGroups {
		aspects.Add(nagRuleGroup)
	}
}

func (a *App) SynthAllStacks() {
	defer jsii.Close()

	app := awscdk.NewApp(a.props)

	stacks.NewCDNStack(app, "CDNStack", &stacks.CDNStackProps{
		StackProps: awscdk.StackProps{
			Env: a.env,
		},
		BucketName: a.config.Get("cdn-stack.bucket-name"),
	})

	addNagRuleGroups(app)
	app.Synth(nil)
}

package main

import "github.com/MatthewFrisby/aws-cdk-infra/pkg/app"

func main() {
	cdkApp := app.NewApp("config.json", nil)
	cdkApp.SynthAllStacks()
}

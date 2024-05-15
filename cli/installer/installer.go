package installer

import (
	cliConfig "github.com/intelops/qualitytrace/cli/config"
	cliUI "github.com/intelops/qualitytrace/cli/ui"
)

var (
	Force             = false
	RunEnvironment    = NoneRunEnvironmentType
	InstallationMode  = NotChosenInstallationModeType
	KubernetesContext = ""
)

const createIssueMsg = "If you need help, please create an issue: https://github.com/intelops/qualitytrace/issues/new/choose"

func Start() {
	ui := cliUI.DefaultUI

	ui.Banner(cliConfig.Version)

	ui.Println(`
Hi! Welcome to the Qualitytrace server installer. I'll help you set up your Qualitytrace server by asking you a few questions
and configuring your system with all the requirements, so you can start Qualitytraceing right away!

To get more info about Qualitytrace, you can check our docs at https://intelops.github.io/qualitytrace/

If you have any issues, please let us know by creating an issue (https://github.com/intelops/qualitytrace/issues/new/choose)
or reach us on Slack https://dub.sh/qualitytrace-community

`)

	if RunEnvironment == DockerRunEnvironmentType { // check if docker was previously chosen as a CLI arg
		ui.Println("How do you want to run Qualitytrace?")
		ui.Println("  > Using Docker Compose")
		dockerCompose.Install(ui)
		return
	}

	if RunEnvironment == KubernetesRunEnvironmentType { // check if kubernetes was previously chosen as a CLI arg
		ui.Println("How do you want to run Qualitytrace?")
		ui.Println("  > Using Kubernetes")
		kubernetes.Install(ui)
		return
	}

	option := ui.Select("How do you want to run Qualitytrace?", []cliUI.Option{
		{Text: "Using Docker Compose", Fn: dockerCompose.Install},
		{Text: "Using Kubernetes", Fn: kubernetes.Install},
	}, 0)

	option.Fn(ui)
}

type installer struct {
	name      string
	preChecks []preChecker
	configs   []configurator
	installFn func(config configuration, ui cliUI.UI)
}

func (i installer) PreCheck(ui cliUI.UI) {
	ui.Title("Let's check if your system has everything we need")
	for _, pc := range i.preChecks {
		pc(ui)
	}

	ui.Title("Your system is ready! Now, let's configure Qualitytrace")
}

func (i installer) Configure(ui cliUI.UI) configuration {
	config := newConfiguration(ui)
	config.set("installer", i.name)

	setInstallationType(ui, config)

	for _, confFn := range i.configs {
		config = confFn(config, ui)
	}

	return config
}

func (i installer) Install(ui cliUI.UI) {
	i.PreCheck(ui)

	conf := i.Configure(ui)

	ui.Title("Thanks! We are ready to install Qualitytrace now")

	i.installFn(conf, ui)
}

type preChecker func(ui cliUI.UI)

func setInstallationType(ui cliUI.UI, config configuration) {
	if InstallationMode == WithoutDemoInstallationModeType { // check if it was previously chosen
		ui.Println("Do you have OpenTelemetry based tracing already set up, or would you like us to install a demo tracing environment and app?")
		ui.Println("  > I have a tracing environment already. Just install Qualitytrace")
		config.set("installer.only_qualitytrace", true)
		return
	}

	if InstallationMode == WithDemoInstallationModeType { // check if it was previously chosen
		ui.Println("Do you have OpenTelemetry based tracing already set up, or would you like us to install a demo tracing environment and app?")
		ui.Println("  > Just learning tracing! Install Qualitytrace, OpenTelemetry Collector and the sample app.")
		config.set("installer.only_qualitytrace", false)
		return
	}

	option := ui.Select("Do you have OpenTelemetry based tracing already set up, or would you like us to install a demo tracing environment and app?", []cliUI.Option{
		{Text: "I have a tracing environment already. Just install Qualitytrace", Fn: func(ui cliUI.UI) {
			config.set("installer.only_qualitytrace", true)
		}},
		{Text: "Just learning tracing! Install Qualitytrace, OpenTelemetry Collector and the sample app.", Fn: func(ui cliUI.UI) {
			config.set("installer.only_qualitytrace", false)
		}},
	}, 0)

	option.Fn(ui)
}

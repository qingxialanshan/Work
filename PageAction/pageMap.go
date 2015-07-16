package main

import (
	"pageaction"
)

type Actions interface {
	PostAction(arg ...string)
}

var comps = map[string]Actions{
	"setup":              &pageaction.Setup{},
	"suggestion":         &pageaction.Suggestion{},
	"install_dependency": &pageaction.InstallDependency{},
	"install_config":     &pageaction.InstallConfig{},
	"installation":       &pageaction.Installation{},
	"device_setup":       &pageaction.DeviceSetup{},
	"deploy_samples":     &pageaction.DeploySamples{},
	"finished":           &pageaction.Finished{},
}

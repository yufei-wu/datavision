package commands

import (
	"github.com/grafana/grafana/pkg/cmd/datavision-cli/logger"
	s "github.com/grafana/grafana/pkg/cmd/datavision-cli/services"
)

func listremoteCommand(c CommandLine) error {
	plugin, err := s.ListAllPlugins(c.RepoDirectory())

	if err != nil {
		return err
	}

	for _, i := range plugin.Plugins {
		pluginVersion := ""
		if len(i.Versions) > 0 {
			pluginVersion = i.Versions[0].Version
		}

		logger.Infof("id: %v version: %s\n", i.Id, pluginVersion)
	}

	return nil
}

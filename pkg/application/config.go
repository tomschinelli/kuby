package application

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tomschinelli/kummy/pkg/cluster"
)

type Config struct {
	err          error
	outputFormat string
	cluster      *cluster.Cluster
	ctx          context.Context
}

func ParseConfigForCommand(cmd *cobra.Command) *Config {
	return &Config{
		outputFormat: "table",
	}
}

func (c *Config) UseCluster(cluster *cluster.Cluster) *Config {
	c.cluster = cluster
	return c
}

func (c *Config) UseContext(ctx context.Context) *Config {
	c.ctx = ctx
	return c
}

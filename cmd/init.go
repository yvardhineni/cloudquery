package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/google/uuid"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const initHelpMsg = "Generate initial config.hcl for fetch command"

var (
	initCmd = &cobra.Command{
		Use:   "init [choose one or more providers (aws gcp azure okta ...)]",
		Short: initHelpMsg,
		Long:  initHelpMsg,
		Example: `
  # Downloads aws provider and generates config.hcl for aws provider
  cloudquery init aws

  # Downloads aws,gcp providers and generates one config.hcl with both providers
  cloudquery init aws gcp`,
		Args: cobra.MinimumNArgs(1),
		Run: handleCommand(func(ctx context.Context, _ *console.Client, cmd *cobra.Command, args []string) error {
			return Initialize(ctx, args)
		}),
	}
)

func Initialize(ctx context.Context, providers []string) error {
	fs := afero.NewOsFs()

	configPath := viper.GetString("configPath")

	if info, _ := fs.Stat(configPath); info != nil {
		ui.ColorizedOutput(ui.ColorError, "Error: Config file %s already exists\n", configPath)
		return diag.FromError(fmt.Errorf("config file %q already exists", configPath), diag.USER)
	}
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()
	requiredProviders := make([]*config.RequiredProvider, len(providers))
	for i, p := range providers {
		organization, providerName, provVersion, err := registry.ParseProviderNameWithVersion(p)
		if err != nil {
			return fmt.Errorf("could not parse requested provider: %w", err)
		}
		rp := config.RequiredProvider{
			Name:    providerName,
			Version: provVersion,
		}
		if organization != registry.DefaultOrganization {
			source := fmt.Sprintf("%s/%s", organization, providerName)
			rp.Source = &source
		}
		requiredProviders[i] = &rp
		providers[i] = providerName // overwrite "provider@version" with just "provider"
	}
	// TODO: build this manually with block and add comments as well
	cqBlock := gohcl.EncodeAsBlock(&config.CloudQuery{
		PluginDirectory: "./cq/providers",
		PolicyDirectory: "./cq/policies",
		Providers:       requiredProviders,
		Connection: &config.Connection{
			Username: "postgres",
			Password: "pass",
			Host:     "localhost",
			Port:     5432,
			Database: "postgres",
			SSLMode:  "disable",
		},
	}, "cloudquery")

	// Update connection block to remove unwanted keys
	if b := cqBlock.Body().FirstMatchingBlock("connection", nil); b != nil {
		bd := b.Body()
		bd.RemoveAttribute("dsn")
		bd.RemoveAttribute("type")
		bd.RemoveAttribute("extras")
	}

	rootBody.AppendBlock(cqBlock)
	cfg, diags := config.NewParser(
		config.WithEnvironmentVariables(config.EnvVarPrefix, os.Environ()),
	).LoadConfigFromSource("init.hcl", f.Bytes())
	if diags.HasErrors() {
		return diags
	}

	cfg.CloudQuery.Connection.DSN = "" // Don't connect
	c, err := console.CreateClientFromConfig(ctx, cfg, uuid.Nil)
	if err != nil {
		return err
	}
	defer c.Close()
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	rootBody.AppendNewline()
	rootBody.AppendUnstructuredTokens(hclwrite.Tokens{
		{
			Type:  hclsyntax.TokenComment,
			Bytes: []byte("// All Provider Configurations"),
		},
	})
	rootBody.AppendNewline()
	var buffer bytes.Buffer
	buffer.WriteString("// Configuration AutoGenerated by CloudQuery CLI\n")
	if n, err := buffer.Write(f.Bytes()); n != len(f.Bytes()) || err != nil {
		return err
	}
	for _, p := range providers {
		pCfg, diags := core.GetProviderConfiguration(ctx, c.PluginManager, &core.GetProviderConfigOptions{
			Provider: c.ConvertRequiredToRegistry(p),
		})

		if diags.HasErrors() {
			return diags
		}
		buffer.Write(pCfg.Config)
		buffer.WriteString("\n")
	}
	mm := module.NewManager(nil, nil)
	mm.Register(drift.New())
	if mex := mm.ExampleConfigs(providers); len(mex) > 0 {
		buffer.WriteString("\n// Module Configurations\nmodules {\n")
		for _, c := range mex {
			buffer.WriteString(c)
			buffer.WriteString("\n")
		}
		buffer.WriteString("}\n")
	}

	formattedData := hclwrite.Format(buffer.Bytes())
	_ = afero.WriteFile(fs, configPath, formattedData, 0644)
	ui.ColorizedOutput(ui.ColorSuccess, "configuration generated successfully to %s\n", configPath)
	return nil
}

func init() {
	initCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(initCmd)
}

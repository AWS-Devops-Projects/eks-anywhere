package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/config"
	"github.com/aws/eks-anywhere/pkg/curatedpackages"
	"github.com/aws/eks-anywhere/pkg/kubeconfig"
	"github.com/aws/eks-anywhere/pkg/version"
)

type generatePackageOptions struct {
	source      curatedpackages.BundleSource
	kubeVersion string
	registry    string
}

var gpOptions = &generatePackageOptions{}

func init() {
	generateCmd.AddCommand(generatePackageCommand)
	generatePackageCommand.Flags().Var(&gpOptions.source, "source", "Location to find curated packages: (cluster, registry)")
	if err := generatePackageCommand.MarkFlagRequired("source"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
	generatePackageCommand.Flags().StringVar(&gpOptions.kubeVersion, "kubeversion", "", "Kubernetes Version of the cluster to be used. Format <major>.<minor>")
	generatePackageCommand.Flags().StringVar(&gpOptions.registry, "registry", "", "Used to specify an alternative registry for package generation")
}

var generatePackageCommand = &cobra.Command{
	Use:          "packages [flags]",
	Aliases:      []string{"package", "packages"},
	Short:        "Generate package(s) configuration",
	Long:         "Generates Kubernetes configuration files for curated packages",
	PreRunE:      preRunPackages,
	SilenceUsage: true,
	RunE:         runGeneratePackages,
}

func runGeneratePackages(cmd *cobra.Command, args []string) error {
	if err := curatedpackages.ValidateKubeVersion(gpOptions.kubeVersion, gpOptions.source); err != nil {
		return err
	}
	return generatePackages(cmd.Context(), args)
}

func generatePackages(ctx context.Context, args []string) error {
	kubeConfig := kubeconfig.FromEnvironment()
	deps, err := newDependenciesForPackages(ctx, kubeConfig)
	if err != nil {
		return fmt.Errorf("unable to initialize executables: %v", err)
	}
	bm := curatedpackages.CreateBundleManager(gpOptions.kubeVersion)
	username, password, err := config.ReadCredentials()
	if err != nil && gpOptions.registry != "" {
		return err
	}
	registry, err := curatedpackages.NewRegistry(deps, gpOptions.registry, gpOptions.kubeVersion, username, password)
	if err != nil {
		return err
	}

	b := curatedpackages.NewBundleReader(
		kubeConfig,
		gpOptions.kubeVersion,
		gpOptions.source,
		deps.Kubectl,
		bm,
		version.Get(),
		registry,
	)

	bundle, err := b.GetLatestBundle(ctx)
	if err != nil {
		return err
	}

	packageClient := curatedpackages.NewPackageClient(
		bundle,
		deps.Kubectl,
		args...,
	)
	packages, err := packageClient.GeneratePackages()
	if err != nil {
		return err
	}
	if err = packageClient.WritePackagesToStdOut(packages); err != nil {
		return err
	}
	return nil
}

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sundaram2021/dockit/pkg/builder"
	"github.com/sundaram2021/dockit/pkg/git"
)

var (
	repoURL     string
	branch      string
	imageTag    string
	outputDir   string
	verbose     bool
)

var rootCmd = &cobra.Command{
	Use:   "dockit",
	Short: "Dockit - A CLI tool to build Docker images from GitHub repositories",
	Long: `Dockit simplifies the process of building Docker images directly from GitHub repositories.
	
Usage:
  dockit [github-repo-url] [flags]

Example:
  dockit https://github.com/example/repo --branch main --tag my-image`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setupLogging()

		if len(args) > 0 {
			repoURL = args[0]
		}

		if repoURL == "" {
			logrus.Fatal("GitHub repository URL is required")
		}

		if outputDir == "" {
			tempDir, err := os.MkdirTemp("", "dockit-*")
			if err != nil {
				logrus.Fatalf("Failed to create temporary directory: %v", err)
			}
			outputDir = tempDir
			defer os.RemoveAll(tempDir)
		}

		repoPath, err := git.CloneRepository(repoURL, branch, outputDir)
		if err != nil {
			logrus.Fatalf("Repository cloning failed: %v", err)
		}

		imageName := imageTag
		if imageName == "" {
			imageName = filepath.Base(repoURL)
		}

		err = builder.BuildDockerImage(repoPath, imageName)
		if err != nil {
			logrus.Fatalf("Docker image build failed: %v", err)
		}

		logrus.Infof("Successfully built Docker image: %s", imageName)
	},
}

func setupLogging() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&repoURL, "url", "", "GitHub repository URL")
	rootCmd.Flags().StringVar(&branch, "branch", "main", "Branch to clone")
	rootCmd.Flags().StringVar(&imageTag, "tag", "", "Custom Docker image tag")
	rootCmd.Flags().StringVar(&outputDir, "output", "", "Output directory for cloned repository")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose logging")
}
package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	dockerfileURL string
	sourceCodeURL string
	rootCmd       = &cobra.Command{
		Use:   "dockit",
		Short: "dockit is a Simple Docker Tool",
		Long: `dockit is a CLI tool that wraps Docker commands and helps to run Dockerfiles present without downloading source code.
	dockit -d <dockerfile-link> -s <source-code-folder-link>`,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&dockerfileURL, "dockerfile", "d", "", "URL of the Dockerfile")
	rootCmd.PersistentFlags().StringVarP(&sourceCodeURL, "source", "s", "", "URL of the source code folder")
	rootCmd.MarkPersistentFlagRequired("dockerfile")
	rootCmd.MarkPersistentFlagRequired("source")

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		ExecuteDockerfile()
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

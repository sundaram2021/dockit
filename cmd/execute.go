package cmd

import (
	"os"
	"os/exec"
)

func ExecuteDockerfile() {

}

func runDockerBuild(dockerfilePath, contextDir string) error {
	cmd := exec.Command("docker", "build", "-f", dockerfilePath, contextDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

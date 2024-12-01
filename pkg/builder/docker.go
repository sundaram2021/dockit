package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type ImageInfo struct {
	ID           string   `json:"Id"`
	RepoTags     []string `json:"RepoTags"`
	Created      string   `json:"Created"`
	Size         int64    `json:"Size"`
	VirtualSize  int64    `json:"VirtualSize"`
}

func BuildDockerImage(contextPath, imageName string) error {
	dockerfilePath, err := findDockerfile(contextPath)
	if err != nil {
		return fmt.Errorf("no Dockerfile found in the repository: %v", err)
	}

	dockerfileDir := filepath.Dir(dockerfilePath)

	// Validate Docker is installed
	_, err = exec.LookPath("docker")
	if err != nil {
		return fmt.Errorf("docker is not installed: %v", err)
	}

	buildCmd := exec.Command(
		"docker", 
		"build", 
		"-t", 
		imageName, 
		"-f", 
		dockerfilePath, 
		dockerfileDir,
	)

	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	logrus.Debugf("Executing command: %v", buildCmd.String())

	err = buildCmd.Run()
	if err != nil {
		return fmt.Errorf("docker build failed: %v", err)
	}

	logrus.Infof("Successfully built Docker image: %s", imageName)

	if err := logImageDetails(imageName); err != nil {
		logrus.Warnf("Failed to retrieve image details: %v", err)
	}

	err = os.RemoveAll(contextPath)
	if err != nil {
		logrus.Warnf("Failed to remove temporary directory %s: %v", contextPath, err)
	} else {
		logrus.Debugf("Removed temporary directory: %s", contextPath)
	}

	return nil
}

func logImageDetails(imageName string) error {
	cmd := exec.Command("docker", "image", "inspect", imageName)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to inspect image: %v", err)
	}

	var imageInfoList []ImageInfo
	err = json.Unmarshal(out.Bytes(), &imageInfoList)
	if err != nil {
		return fmt.Errorf("failed to parse image info: %v", err)
	}

	if len(imageInfoList) == 0 {
		return fmt.Errorf("no image information found")
	}

	imageInfo := imageInfoList[0]

	sizeInMB := float64(imageInfo.Size) / (1024 * 1024)

	logrus.Info("🖼️ Docker Image Details:")
	logrus.Infof("  Image Name   : %s", strings.Join(imageInfo.RepoTags, ", "))
	logrus.Infof("  Image ID     : %s", imageInfo.ID)
	logrus.Infof("  Created      : %s", imageInfo.Created)
	logrus.Infof("  Size         : %.2f MB", sizeInMB)

	return nil
}

func findDockerfile(rootPath string) (string, error) {
	var dockerfilePath string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == "Dockerfile" {
			dockerfilePath = path
			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if dockerfilePath == "" {
		return "", fmt.Errorf("no Dockerfile found")
	}

	return dockerfilePath, nil
}
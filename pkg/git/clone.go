package git

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/sirupsen/logrus"
)

func CloneRepository(repoURL, branch, outputDir string) (string, error) {
	logrus.Debugf("Cloning repository: %s (branch: %s)", repoURL, branch)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", err
	}

	repoPath := filepath.Join(outputDir, filepath.Base(repoURL))

	cloneOptions := &git.CloneOptions{
		URL:           repoURL,
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
	}

	_, err := git.PlainClone(repoPath, false, cloneOptions)
	if err != nil {
		logrus.Errorf("Repository cloning failed: %v", err)
		return "", err
	}

	logrus.Infof("Repository cloned successfully to: %s", repoPath)
	return repoPath, nil
}
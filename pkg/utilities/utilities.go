package utilities

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/autonomy/conform/pkg/constants"
	"github.com/docker/docker/client"
)

// ImageName formats the image name based on the status of a git repository.
func ImageName(repository, sha string, isClean bool) (image string, err error) {
	if !isClean {
		image = formatImageNameDirty(repository)
	} else {
		image = formatImageNameSHA(repository, sha)
	}

	return
}

func formatImageNameDirty(repository string) string {
	return fmt.Sprintf("%s:dirty", repository)
}

func formatImageNameSHA(repository, sha string) string {
	return fmt.Sprintf("%s:%s", repository, sha)
}

// func formatImageNameTag(repository, tag string) string {
// 	return fmt.Sprintf("%s:%s", repository, tag)
// }
//
// func formatImageNameLatest(repository string) string {
// 	return fmt.Sprintf("%s:latest", repository)
// }

// CheckDockerVersion checks the Docker server version and returns an error if
// it is an incompatible version.
func CheckDockerVersion() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	serverVersion, err := cli.ServerVersion(context.Background())
	if err != nil {
		return err
	}
	minVersion, err := semver.NewVersion(constants.MinDockerVersion)
	if err != nil {
		return err
	}
	serverSemVer := semver.MustParse(serverVersion.Version)
	i := serverSemVer.Compare(minVersion)
	if i < 0 {
		err = fmt.Errorf("At least Docker version %s is required", constants.MinDockerVersion)

		return err
	}

	return nil
}

package utilities

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/autonomy/conform/pkg/constants"
	"github.com/docker/docker/client"
)

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

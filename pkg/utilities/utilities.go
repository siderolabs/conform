package utilities

import "fmt"

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

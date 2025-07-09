package ffmpeg

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Gets a single frame from a stream and saves to a file, it checks if theres already a thumbnail generated
func GenerateThumbnail(uri, path string, cacheDuration time.Duration) error {
	splittedPath := strings.Split(path, "/")
	dir := strings.Join(splittedPath[0:len(splittedPath)-1], "/")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0755); err != nil {
			return err
		}
	}

	// checks if there's already a thumbnail for this uri
	// if so, use cache
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	args := []string{"-y", "-i", uri, "-vframes", "1", path}

	cmd := exec.Command("ffmpeg", args...)
	if err := cmd.Run(); err != nil {
		return err
	}

	go func() {
		<-time.After(cacheDuration)
		logrus.Debugf("removing cache for %s", uri)
		if err := os.Remove(path); err != nil {
			logrus.Errorf("error removing thumb cache for %s.Error: %s", uri, err)
		}
	}()

	return nil
}

package redpill

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	cgroupsFile    = "/proc/1/cgroup"
	dockerNsPrefix = "/docker/"
)

// InContainer returns true if the process is running in a container
func InContainer() (bool, error) {
	ns, err := GetContainerID()
	return ns != "", err
}

// GetContainerID returns the container id if running in docker, or "" otherwise
func GetContainerID() (string, error) {
	f, err := os.Open(cgroupsFile)
	if err != nil {
		return "", fmt.Errorf("redpill: error detecting container id: %v", err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	// going to get lines like "11:pids:/docker/ed807a7d59accf3b9e70c05d52cc494fc125377b82e4"
	// or "11:pids:/" if not in container

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		if len(line) != 3 {
			// don't really know what's going on here so ignore
			continue
		}

		ns := strings.TrimSpace(line[2])

		if strings.HasPrefix(ns, dockerNsPrefix) {
			return strings.TrimPrefix(ns, dockerNsPrefix), nil
		}
	}

	return "", nil
}

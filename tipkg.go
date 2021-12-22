package tipkg

import (
	"fmt"
	"os"
	"path/filepath"
)

var tiup_home string

func init() {
	tiup_home = os.Getenv("TIUP_HOME")
	if tiup_home == "" {
		home, _ := os.UserHomeDir()
		tiup_home = filepath.Join(home, ".tiup")
	}
}

// Download package to TIUP_HOME
func Download(name, version string) error {
	if tiup_home == "" {
		return fmt.Errorf("cannot found env $TIUP_HOME nor $HOME")
	}
	return nil
}

func Remove(name, version string) error {
	if tiup_home == "" {
		return fmt.Errorf("cannot found env $TIUP_HOME nor $HOME")
	}
	return nil
}

func Install(name, version string) error {
	if tiup_home == "" {
		return fmt.Errorf("cannot found env $TIUP_HOME nor $HOME")
	}
	return nil
}

func Uninstall(name, version string) error {
	if tiup_home == "" {
		return fmt.Errorf("cannot found env $TIUP_HOME nor $HOME")
	}
	return nil
}

func GetExactlyVersion(name string) (string, error) {
	switch name {
	case "", "stable":
		return "TBD", nil
	case "nightly":
		return "TBD", nil
	}
	return "", nil
}

func UpdateRootManifest() error {
	return nil
}

func UpdateTimestampManifest() error {
	return nil
}

func UpdateSnapshootManifest() error {
	return nil
}

func UpdateIndexManifest() error {
	return nil
}
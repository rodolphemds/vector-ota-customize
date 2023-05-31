package patcher

import (
	"errors"
	"fmt"

	"github.com/kercre123/wire-os/pkg/vars"
)

func RunPatches(version vars.Version, target int) error {
	if target < 0 || target > 3 {
		return errors.New("target must be between 0 and 3")
	}
	if !vars.IsMounted() {
		return errors.New("image is not mounted")
	}
	for _, patch := range WireOSPatches {
		fmt.Println("Running: " + patch.Name + " (" + patch.Description + ")")
		err := patch.Patch(version, target)
		if err != nil {
			return err
		}
		vars.PatchLogger("Completed without error.")
	}
	return nil
}

package dirSync

import (
	"os"

	cp "github.com/otiai10/copy"
)

func dirSync(syncFrom string, syncTo string, copy bool) error {

	syncFrom = os.ExpandEnv(syncFrom)
	syncTo = os.ExpandEnv(syncTo)

	if _, err := os.Stat(syncFrom); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(syncTo); os.IsNotExist(err) {
		os.Create(syncTo)
	}

	if copy == true {
		err := cp.Copy(syncFrom, syncTo)
		if err != nil {
			return err
		}
	} else {
		err := os.Symlink(syncFrom, syncTo)
		if err != nil {
			return err
		}
	}

	return nil
}

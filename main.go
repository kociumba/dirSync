package dirSync

import "os"

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
		err := os.Rename(syncFrom, syncTo)
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

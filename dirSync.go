// A very small package that syncs the contents of a directory to another with the option to create symlinks instead of copying
package dirSync

import (
	"os"

	cp "github.com/otiai10/copy"
)

// DirSync synchronizes the contents of the syncFrom directory to the syncTo directory.
//
// It takes two parameters: syncFrom, which is the source directory to sync from, and syncTo, which is the target directory to sync to.
// Both syncFrom and syncTo can contain env variables, wich are going to be expanded during runtime.
//
// The function has an optional parameter named copy, which indicates whether to perform a copy or a symlink operation.
// If copy is set to true, the function will copy the files from syncFrom to syncTo.
// If copy is set to false, the function will create symbolic links from syncFrom to syncTo.
//
// The function returns an error if there is any error during the synchronization process aswell as if the syncFrom directory does not exist.
// If the synchronization is successful, it returns nil.
func DirSync(syncFrom string, syncTo string, copy bool) error {

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

// i hate git tags

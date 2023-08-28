package dirSync

import (
	"testing"
)

func TestDirSync(t *testing.T) {
	// Test case: syncFrom directory does not exist
	err := DirSync("/path/to/nonexistent", "$HOMEPATH/dirSync/syncTo", true)
	if err == nil {
		t.Log(err)
	} else {
		t.Log(err)
		t.Errorf("Expected error when syncFrom directory does not exist")
	}

	// Test case: syncTo directory does not exist and is created
	err = DirSync("$HOMEPATH/dirSync/syncFrom", "/path/to/nonexistent", true)
	if err != nil {
		t.Log(err)
		t.Errorf("Expected nil error when creating syncTo directory")
	}

	// Test case: copy is true, files are copied successfully
	err = DirSync("$HOMEPATH/dirSync/syncFrom", "$HOMEPATH/dirSync/syncTo", true)
	if err != nil {
		t.Log(err)
		t.Errorf("Expected nil error when copying files")
	}

	// Test case: copy is false, symbolic links are created successfully
	err = DirSync("$HOMEPATH/dirSync/syncFrom", "syncToSymLink", false)
	if err != nil {
		t.Log(err)
		t.Errorf("Expected nil error when creating symbolic links")
	}
}

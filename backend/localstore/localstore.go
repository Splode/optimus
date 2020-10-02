package localstore

import (
	"github.com/vrischmann/userdir"
	"io/ioutil"
	"os"
	"path"
)

// LocalStore provides reading and writing application data to the user's
// configuration directory.
type LocalStore struct {
	ConfDir string
}

// NewLocalStore returns a localStore instance.
func NewLocalStore() *LocalStore {
	return &LocalStore{ConfDir: path.Join(userdir.GetConfigHome(), "Optimus")}
}

// Load reads the given file in the user's configuration directory and returns
// its contents.
func (l *LocalStore) Load(filename string) ([]byte, error) {
	p := path.Join(l.ConfDir, filename)
	d, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return d, err
}

// Store writes data to the user's configuration directory at the given
// filename.
func (l *LocalStore) Store(data []byte, filename string) error {
	p := path.Join(l.ConfDir, filename)
	if err := ensureDirExists(l.ConfDir); err != nil {
		return err
	}
	if err := ioutil.WriteFile(p, data, 0777); err != nil {
		return err
	}
	return nil
}

// ensureDirExists checks for the existence of the directory at the given path,
// which is created if it does not exist.
func ensureDirExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err = os.Mkdir(path, 0777); err != nil {
			return err
		}
	}
	return nil
}

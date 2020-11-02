package mangostine

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

//check the config and data file exist and remove
func fileexists(configpath string, datapath string) error {
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		return err
	}

	err := os.RemoveAll(configpath)
	if err != nil {
		return err
	}

	if _, err := os.Stat(datapath); os.IsNotExist(err) {
		return err
	}

	err = os.RemoveAll(datapath)
	if err != nil {
		return err
	}

	return nil
}

func folderexists(configpath string, datapath string) error {
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		return err
	}

	err := os.Remove(configpath)
	if err != nil {
		return err
	}
	if _, err := os.Stat(datapath); os.IsNotExist(err) {
		return err
	}
	err = os.Remove(datapath)
	if err != nil {
		return err
	}

	return nil
}

func copyKeyNode(keyPath string, nodepath []string) error {
	for _, path := range nodepath {
		n1 := filepath.Join(path, "keys")
		err := CopyDir(keyPath, n1)
		if err != nil {
			return fmt.Errorf("Unable to create new account: %v", err)
		}
	}
	return nil
}

// CopyFile copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func CopyFile(src, dst string) (err error) {

	if fileExists(dst) {
		return nil
	} else {

		in, err := os.Open(src)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer func() {
			if e := out.Close(); e != nil {
				err = e
			}
		}()

		_, err = io.Copy(out, in)
		if err != nil {
			return err
		}

		err = out.Sync()
		if err != nil {
			return err
		}

		si, err := os.Stat(src)
		if err != nil {
			return err
		}
		err = os.Chmod(dst, si.Mode())
		if err != nil {
			return err
		}
	}
	return err
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func EnsuredeleteDir(dir string, mode os.FileMode) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return fmt.Errorf("Could not create directory %v. %v", dir, err)
	}
	return nil
}
func EnsureDir(dir string, mode os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, mode)
		if err != nil {
			return fmt.Errorf("Could not create directory %v. %v", dir, err)
		}
	}
	return nil
}

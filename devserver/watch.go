package devserver

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// track changes from the following directories:
// - source directory (excluding components dir, if it is within the source directory)
// OR if the components directory is within the source directory then just ONLY track the source directory anyways
// - components directory by itself
// from all tracked directories, output dir must be excluded along with common excluded directories

// TODO: dir parameter removed in favour of source and components dir and excludes list (when above changes are implemented)
// also refer to commands/serve.go for more information
func Watch(dir string) (<-chan bool, <-chan error) {
	events := make(chan bool)
	errs := make(chan error)

	go func() {
		defer close(events)
		defer close(errs)

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			errs <- err
			return
		}
		defer watcher.Close()

		// recursively watch ALL subdirectories
		err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return watcher.Add(path)
			}

			return nil
		})
		if err != nil {
			errs <- err
		}

		for {
			select {
			case e, ok := <-watcher.Events:
				if !ok {
					return
				}

				// handle new directories dynamically
				if e.Op&fsnotify.Create != 0 {
					if info, err := os.Stat(e.Name); err == nil && info.IsDir() {
						_ = watcher.Add(e.Name)
					}
				}

				// we only want writes, creates and deletes
				if e.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove|fsnotify.Rename) != 0 {
					events <- true
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				errs <- err
			}
		}
	}()

	return events, errs
}

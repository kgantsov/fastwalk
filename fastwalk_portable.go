// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine !linux,!darwin,!freebsd,!openbsd,!netbsd

package fastwalk

import (
	"io/ioutil"
	"os"
)

// readDir calls fn for each directory entry in dirName.
// It does not descend into directories or follow symlinks.
// If fn returns a non-nil error, readDir returns with that error
// immediately.
func readDir(dirName string) (Dirs, error) {
	dirs := make(Dirs, 0)

	fis, err := ioutil.ReadDir(dirName)
	if err != nil {
		return dirs, err
	}
	skipFiles := false
	for _, fi := range fis {
		if fi.Mode().IsRegular() && skipFiles {
			continue
		}
		dirs = append(dirs, Dir{dirName, fi.Name(), fi.Mode() & os.ModeType})
	}
	return dirs, nil
}

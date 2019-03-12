// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fastwalk provides a faster version of filepath.Walk for file system
// scanning tools.
package fastwalk

import (
	"errors"
	"os"
	"path/filepath"
)

// TraverseLink is used as a return value from WalkFuncs to indicate that the
// symlink named in the call may be traversed.
var TraverseLink = errors.New("fastwalk: traverse symlink, assuming target is a directory")

// SkipFiles is a used as a return value from WalkFuncs to indicate that the
// callback should not be called for any other files in the current directory.
// Child directories will still be traversed.
var SkipFiles = errors.New("fastwalk: skip remaining files in directory")

type WalkFunc func(path string, fileType os.FileMode) error

func walk(path string, fileType os.FileMode, walkFn WalkFunc) error {
	if fileType != os.ModeDir {
		return walkFn(path, fileType)
	}

	files, err := ReadDir(path)
	err1 := walkFn(path, fileType)

	if err != nil || err1 != nil {
		return err1
	}
	// err1 := walkFn(path, info, err)
	// If err != nil, walk can't walk into this directory.
	if err != nil {
		// The caller's behavior is controlled by the return value, which is decided
		// by walkFn. walkFn may ignore err and return nil.
		// If walkFn returns SkipDir, it will be handled by the caller.
		// So walk should return whatever walkFn returns.
		return err
	}

	for _, file := range files {
		filename := filepath.Join(path, file.Name)

		if file.Type == os.ModeDir {
			err = walk(filename, file.Type, walkFn)
			if err != nil {
				if file.Type != os.ModeDir || err != filepath.SkipDir {
					return err
				}
			}
		} else {
			if err := walkFn(filename, file.Type); err != nil && err != filepath.SkipDir {
				return err
			}
		}
	}
	return nil
}

// Walk walks the file tree rooted at root, calling walkFn for each file or
// directory in the tree, including root. All errors that arise visiting files
// and directories are filtered by walkFn. The files are walked in lexical
// order, which makes the output deterministic but means that for very
// large directories Walk can be inefficient.
// Walk does not follow symbolic links.
func Walk(root string, walkFn WalkFunc) error {
	info, err := os.Lstat(root)
	fileType := info.Mode() & os.ModeType
	if err != nil {
		err = walkFn(root, fileType)
	} else {
		err = walk(root, fileType, walkFn)
	}
	if err == filepath.SkipDir {
		return nil
	}
	return err
}

// ReadDir calls fn for each directory entry in dirName.
// It does not descend into directories or follow symlinks.
// If fn returns a non-nil error, ReadDir returns with that error
// immediately.
func ReadDir(dirName string) (Dirs, error) {
	return readDir(dirName)
}

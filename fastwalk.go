// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fastwalk provides a faster version of filepath.Walk for file system
// scanning tools.
package fastwalk

import (
	"os"
	"path/filepath"
)

type WalkFunc func(path string, fileType os.FileMode) error

func walk(path string, fileType os.FileMode, walkFn WalkFunc) error {
	if fileType != os.ModeDir {
		return walkFn(path, fileType)
	}

	files, err := ReadDir(path)
	if err != nil {
		return err
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

		err = walk(filename, file.Type, walkFn)
		if err != nil {
			if file.Type != os.ModeDir {
				return err
			}
		}

		if err := walkFn(filename, file.Type); err != nil {
			return err
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
		err = walkFn(root, fileType)
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

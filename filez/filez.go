// Package filez provides various utilities for working with paths, files, and directories.
package filez

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/ibrt/golang-utils/errorz"
)

// MustAbs is like [filepath.Abs], but panics on error.
func MustAbs(path string) string {
	path, err := filepath.Abs(path)
	errorz.MaybeMustWrap(err)
	return path
}

// MustRel is like [filepath.Rel] but panics on error.
func MustRel(src, dst string) string {
	path, err := filepath.Rel(MustAbs(src), MustAbs(dst))
	errorz.MaybeMustWrap(err)
	return path
}

// MustGetwd is like [os.Getwd], but panics on error.
func MustGetwd() string {
	wd, err := os.Getwd()
	errorz.MaybeMustWrap(err)
	return wd
}

// MustChdir is like [os.Chdir], but panics on error.
func MustChdir(wd string) string {
	errorz.MaybeMustWrap(os.Chdir(wd))
	return wd
}

// MustUserHomeDir is like [os.UserHomeDir], but panics on error.
func MustUserHomeDir() string {
	dirPath, err := os.UserHomeDir()
	errorz.MaybeMustWrap(err)
	return dirPath
}

// MustRemoveAll is like [os.RemoveAll], but panics on error.
func MustRemoveAll(path string) {
	errorz.MaybeMustWrap(os.RemoveAll(path))
}

// MustReadFile reads a file, panics on error.
func MustReadFile(filePath string) []byte {
	buf, err := os.ReadFile(filePath)
	errorz.MaybeMustWrap(err)
	return buf
}

// MustReadFileString reads a file, panics on error.
func MustReadFileString(filePath string) string {
	return string(MustReadFile(filePath))
}

// MustWriteFile creates a file with the given mode and contents, also ensuring the containing folder exists.
func MustWriteFile(filePath string, dirMode, fileMode os.FileMode, contents []byte) string {
	errorz.MaybeMustWrap(os.MkdirAll(filepath.Dir(filePath), dirMode))
	errorz.MaybeMustWrap(os.WriteFile(filePath, contents, fileMode))
	return filePath
}

// MustWriteFileString creates a file with the given mode and contents, also ensuring the containing folder exists.
func MustWriteFileString(filePath string, dirMode, fileMode os.FileMode, contents string) string {
	return MustWriteFile(filePath, dirMode, fileMode, []byte(contents))
}

// MustCreateTempFile creates a temporary file with the given contents.
func MustCreateTempFile(contents []byte) string {
	fd, err := os.CreateTemp("", "golang-utils-")
	errorz.MaybeMustWrap(err)
	defer errorz.MustClose(fd)

	_, err = io.Copy(fd, bytes.NewReader(contents))
	errorz.MaybeMustWrap(err)
	return fd.Name()
}

// WithMustCreateTempFile calls MustCreateTempFile, then the given closure, then deletes the temp file.
func WithMustCreateTempFile(contents []byte, f func(filePath string)) {
	filePath := MustCreateTempFile(contents)
	defer func() { errorz.MaybeMustWrap(os.RemoveAll(filePath)) }()

	f(filePath)
}

// MustCreateTempFileString creates a temporary file with the given contents.
func MustCreateTempFileString(contents string) string {
	return MustCreateTempFile([]byte(contents))
}

// WithMustCreateTempFileString calls MustCreateTempFileString, then the given closure, then deletes the temp file.
func WithMustCreateTempFileString(contents string, f func(filePath string)) {
	filePath := MustCreateTempFileString(contents)
	defer func() { MustRemoveAll(filePath) }()

	f(filePath)
}

// MustCreateTempDir is like [os.MkdirTemp], but panics on error.
func MustCreateTempDir() string {
	dirPath, err := os.MkdirTemp("", "golang-utils-")
	errorz.MaybeMustWrap(err)
	return dirPath
}

// WithMustCreateTempDir calls WithMustCreateTempDir, then the given closure, then deletes the temp dir.
func WithMustCreateTempDir(f func(dirPath string)) {
	dirPath := MustCreateTempDir()
	defer func() { MustRemoveAll(dirPath) }()

	f(dirPath)
}

// MustPrepareDir deletes the given directory and its contents (if present) and recreates it.
func MustPrepareDir(dirPath string, dirMode os.FileMode) {
	errorz.MaybeMustWrap(os.RemoveAll(dirPath))
	errorz.MaybeMustWrap(os.MkdirAll(dirPath, dirMode))
}

// MustCheckPathExists checks if the given path exists, panics on errors other than [os.ErrNotExist].
func MustCheckPathExists(fileOrDirPath string) bool {
	if _, err := os.Stat(fileOrDirPath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		errorz.MustWrap(err)
	}
	return true
}

// MustCheckFileExists checks if the given path exists and is a regular file, panics on errors other than
// [os.ErrNotExist].
func MustCheckFileExists(fileOrDirPath string) bool {
	stat, err := os.Stat(fileOrDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		errorz.MustWrap(err)
	}

	return stat.Mode().IsRegular()
}

// MustIsChild returns true if "childPath" is lexically determined to be a child of "parentPath". Panics on error.
func MustIsChild(parentPath, childPath string) bool {
	rel := MustRel(parentPath, childPath)
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

// MustRelForDisplay converts "path" to relative if (1) it is an absolute path, and (2) it is a child of the current
// working directory. It returns "path" cleaned otherwise. Panics on error.
func MustRelForDisplay(path string) string {
	if wd := MustGetwd(); filepath.IsAbs(path) && MustIsChild(wd, path) {
		return MustRel(wd, path)
	}

	return filepath.Clean(path)
}

// MustExport exports an embedded FS to disk.
func MustExport(eFS embed.FS, rootDirPath, outDirPath string) {
	errorz.MaybeMustWrap(fs.WalkDir(eFS, rootDirPath, func(path string, d fs.DirEntry, err error) error {
		errorz.MaybeMustWrap(err)

		if d.Type().IsRegular() {
			buf, err := eFS.ReadFile(path)
			errorz.MaybeMustWrap(err)
			MustWriteFile(
				MustAbs(filepath.Join(outDirPath, strings.TrimPrefix(path, rootDirPath+"/"))),
				0777, 0666, buf)
		}

		return nil
	}))
}

// MustListRegularFilePaths returns a list of regular file paths found in the given root directory, relative to it.
func MustListRegularFilePaths(rootDirPath string) []string {
	filePaths := make([]string, 0)

	errorz.MaybeMustWrap(filepath.WalkDir(MustAbs(rootDirPath), func(path string, d fs.DirEntry, err error) error {
		errorz.MaybeMustWrap(err)

		if d.Type().IsRegular() {
			filePaths = append(filePaths, MustRel(rootDirPath, path))
		}

		return nil
	}))

	return filePaths
}

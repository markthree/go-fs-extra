package test

import (
	"os"
	"path/filepath"
	"testing"

	fe "github.com/markthree/go-fs-extra"
	"github.com/stretchr/testify/assert"
)

func TestCopyFile(t *testing.T) {
	t.Parallel()

	is := assert.New(t)

	src, dest := "fixture/copySrc.js", "fixture/copyDest.js"

	err := fe.CopyFile(src, dest)
	if err != nil {
		t.Error(err)
		return
	}

	srcData, err := os.ReadFile(src)
	if err != nil {
		t.Error(err)
		return
	}

	destData, err := os.ReadFile(dest)
	if err != nil {
		t.Error(err)
		return
	}

	is.Equal(string(srcData), string(destData))
}

func TestIsDir(t *testing.T) {
	t.Parallel()

	is := assert.New(t)

	d, f := "fixture", "fixture/src.js"

	is.Equal(fe.IsDir(d), true)
	is.Equal(fe.IsDir(f), false)
}

func TestWriteFile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	path, data := "fixture/writeFile.js", "writeFile.js"

	err := fe.WriteFile(path, data)
	if err != nil {
		t.Error(err)
		return
	}

	fileData, err := os.ReadFile(path)

	if err != nil {
		t.Error(err)
		return
	}

	is.Equal(string(fileData), data)
}

func TestReadFile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	path, data := "fixture/readFile.js", "readFile.js\r\n"

	fileData, err := fe.ReadFile[[]byte](path)
	if err != nil {
		t.Error(err)
		return
	}

	fileStringData, err := fe.ReadFile[string](path)
	if err != nil {
		t.Error(err)
		return
	}

	is.Equal(fileData, []byte(data))
	is.Equal(fileStringData, data)
}

func TestPathExists(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(fe.PathExists("fixture/pathExists.js"), true)
	is.Equal(fe.PathExists("fixture/pathExists2.js"), false)
}

func TestEnsureDir(t *testing.T) {
	t.Parallel()

	is := assert.New(t)

	f := "fixture/ensureDir/1/foo.js"
	err := fe.EnsureDir(f)
	if err != nil {
		t.Error(err)
		return
	}

	is.Equal(fe.PathExists(filepath.Dir(f)), true)
}

func TestEnsureFile(t *testing.T) {
	t.Parallel()

	is := assert.New(t)

	src := "fixture/ensureFile/1/foo.js"
	f, err := fe.EnsureFile(src)
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	is.Equal(fe.PathExists(src), true)

	str := "foo.js"
	f.Write([]byte(str))
	fileStr, err := fe.ReadFile[string](src)
	if err != nil {
		t.Error(err)
		return
	}
	is.Equal(str, fileStr)
}

package test

import (
	"os"
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

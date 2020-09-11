package u

import (
	"io"
	"io/ioutil"
	"os"
)

// TempfileWithContent creates a tempfile with specified content written in it, it also seeks the file pointer so you can read it directly.
// The second returned parameter is a cleanup function that closes and removes the temp file.
func TempfileWithContent(content []byte) (*os.File, func(), error) {
	// create temp file
	tmpfile, err := ioutil.TempFile("", "u")
	if err != nil {
		return nil, nil, err
	}

	// write content
	_, err = tmpfile.Write(content)
	if err != nil {
		return nil, nil, err
	}

	// seek at the beginning of file
	_, err = tmpfile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = tmpfile.Close()
		_ = os.Remove(tmpfile.Name())
	}

	return tmpfile, cleanup, nil
}

// MustTempfileWithContent wraps TempfileWithContent and panics if initialization fails.
func MustTempfileWithContent(content []byte) (*os.File, func()) {
	f, cleanup, err := TempfileWithContent(content)
	if err != nil {
		panic(err)
	}
	return f, cleanup
}
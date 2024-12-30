package fileIO

import (
	"io"
	"log"
	"os"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

// ~~~~~~~~~~~~~~~~ File

type File struct {
	filename string
	suffix   string
	relPath  string

	contentBuffer []byte
	lines         int
	linesRead     int
	hasBeenRead   bool

	bytesRead int
}

func NewFileWithSuffix(fn string, suff string, path string) *File {
	f := &File{filename: fn, suffix: suff, relPath: path}
	f.setDefaults()
	return f
}

func NewFile(path, fn string) *File {
	fn, suff := splitFileName(fn)
	return NewFileWithSuffix(fn, suff, path)
}

func (f *File) setDefaults() *File {
	f.hasBeenRead = false
	f.linesRead = 0
	return f
}

func OpenFile(path string, d os.DirEntry) (f *File) { // TODO make the File struct use byte slice rather than string slice
	name := path + d.Name()

	osF, err := os.Open(name)
	handle(err)

	fInf, _ := osF.Stat()
	byteArr := make([]byte, fInf.Size())
	osF.Read(byteArr)

	strArr := helpers.BytesToString(byteArr)

	f = NewFile(path, d.Name())
	f.Append(strArr)

	return
}

func EmptyFile() *File {
	return &File{}
}

func (f *File) IsEmpty() bool {
	return len(f.contentBuffer) == 0
}

func (f *File) Name() string {
	return helpers.Format("%s.%s", f.filename, f.suffix)
}

func (f *File) Contents() []byte {
	return f.contentBuffer
}

func (f *File) ClearFile() {
	if err := os.WriteFile(f.Name(), make([]byte, 0), 0666); err != nil {
		log.Fatal(err)
	}
}

// copied from io.go
// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
func (f *File) Write(p []byte) (n int, err error) {
	if err := os.WriteFile(f.Name(), p, 0664); err != nil {
		log.Fatal(err)
	}

	return
}

// copied from io.go
// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// If len(p) == 0, Read should always return n == 0. It may return a
// non-nil error if some error condition is known, such as EOF.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
func (f File) Read(p []byte) (n int, err error) {
	l := len(f.contentBuffer)
	i := 0

	for ; i < l; i++ {
		v := f.contentBuffer[i]
		if i < len(p) {
			p[i] = v
			n++
		} else if i == len(p) {
			// err
			err = io.EOF
			break
		}
	}

	if i == l {
		err = io.EOF
	}

	f.bytesRead += n
	// helpers.Print(l, len(p), f.Name(), "bytes read", f.bytesRead)
	// time.Sleep(time.Millisecond * 400)
	return
}

func (f *File) String() string {
	return f.Name()
}

func (f *File) Append(s string) {
	f.AppendLine(s)
}

func (f *File) AppendLine(s string) {
	// adds string s to the end of the buffer, newline determines if it is the end of a line,
	// however that really should be determined when writing to the file so lets ignore that
	// especially since we can use the index we are appending to to determine if it's a new line or now
	f.contentBuffer = append(f.contentBuffer, []byte(s)...)
}

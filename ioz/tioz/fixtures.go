package tioz

import (
	"io"
)

// TestByteReader is a mock blueprint.
type TestByteReader interface {
	io.ByteReader
}

// TestByteScanner is a mock blueprint.
type TestByteScanner interface {
	io.ByteScanner
}

// TestByteWriter is a mock blueprint.
type TestByteWriter interface {
	io.ByteWriter
}

// TestCloser is a mock blueprint.
type TestCloser interface {
	io.Closer
}

// TestReadCloser is a mock blueprint.
type TestReadCloser interface {
	io.ReadCloser
}

// TestReadSeekCloser is a mock blueprint.
type TestReadSeekCloser interface {
	io.ReadSeekCloser
}

// TestReadSeeker is a mock blueprint.
type TestReadSeeker interface {
	io.ReadSeeker
}

// TestReadWriteCloser is a mock blueprint.
type TestReadWriteCloser interface {
	io.ReadWriteCloser
}

// TestReadWriteSeeker is a mock blueprint.
type TestReadWriteSeeker interface {
	io.ReadWriteSeeker
}

// TestReadWriter is a mock blueprint.
type TestReadWriter interface {
	io.ReadWriter
}

// TestReader is a mock blueprint.
type TestReader interface {
	io.Reader
}

// TestReaderAt is a mock blueprint.
type TestReaderAt interface {
	io.ReaderAt
}

// TestReaderFrom is a mock blueprint.
type TestReaderFrom interface {
	io.ReaderFrom
}

// TestRuneReader is a mock blueprint.
type TestRuneReader interface {
	io.RuneReader
}

// TestRuneScanner is a mock blueprint.
type TestRuneScanner interface {
	io.RuneScanner
}

// TestSeeker is a mock blueprint.
type TestSeeker interface {
	io.Seeker
}

// TestStringWriter is a mock blueprint.
type TestStringWriter interface {
	io.StringWriter
}

// TestWriteCloser is a mock blueprint.
type TestWriteCloser interface {
	io.WriteCloser
}

// TestWriteSeeker is a mock blueprint.
type TestWriteSeeker interface {
	io.WriteSeeker
}

// TestWriter is a mock blueprint.
type TestWriter interface {
	io.Writer
}

// TestWriterAt is a mock blueprint.
type TestWriterAt interface {
	io.WriterAt
}

// TestWriterTo is a mock blueprint.
type TestWriterTo interface {
	io.WriterTo
}

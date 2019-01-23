package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_io_reader(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	temp := make([]byte, 1024, 1024)
	for i := range temp {
		temp[i] = byte(rand.Intn(255))
	}

	src := bytes.NewBuffer(temp)
	result := bytes.NewBuffer(nil)

	// err := Filter(io.LimitReader(src, 10), result)
	// err := Filter(src, os.Stdout)
	err := Filter(src, result)

	require.NoError(t, err)

	fmt.Printf("resulted in %d bytes\n", result.Len())
}

func Filter(in io.Reader, out io.Writer) error {
	buf := make([]byte, 100)
	for {
		_, err := in.Read(buf)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errors.WithMessage(err, "read failed")
		}
		// the filter
		var outBuf []byte
		for _, b := range buf {
			if b >= 32 && b <= 128 {
				outBuf = append(outBuf, b)
			}
		}

		// forward on
		_, err = out.Write(outBuf)
		if err != nil {
			return errors.WithMessage(err, "write failed")
		}
	}
}

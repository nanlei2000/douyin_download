package utils

import "io"

func SafeClose(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

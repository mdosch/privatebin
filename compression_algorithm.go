// Copyright (c) 2020-2024 Bryan Frimin <bryan@frimin.fr>.
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
// OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

package privatebin

import (
	"encoding/json"
	"fmt"
)

const (
	CompressionAlgorithmUnknow CompressionAlgorithm = iota
	CompressionAlgorithmNone
	CompressionAlgorithmGZip
)

type CompressionAlgorithm uint8

func (ca CompressionAlgorithm) MarshalJSON() ([]byte, error) {
	var s string

	switch ca {
	case CompressionAlgorithmNone:
		s = "none"
	case CompressionAlgorithmGZip:
		s = "zlib"
	default:
		return nil, fmt.Errorf("invalid CompressionAlgorithm value: %v", ca)
	}

	return json.Marshal(s)
}

func (ca *CompressionAlgorithm) UnmarshalJSON(data []byte) error {
	var (
		v CompressionAlgorithm
		s string
	)

	err := json.Unmarshal(data, &s)
	if err != nil {
		return fmt.Errorf("cannot decode compression algorithm: %w", err)
	}

	switch s {
	case "none":
		v = CompressionAlgorithmNone
	case "zlib":
		v = CompressionAlgorithmGZip
	default:
		return fmt.Errorf("invalid compression algorithm value: %v", s)
	}

	*ca = v

	return nil
}

func (ca CompressionAlgorithm) String() string {
	switch ca {
	case CompressionAlgorithmNone:
		return "none"
	case CompressionAlgorithmGZip:
		return "zlib"
	default:
		return "unknow"
	}
}
// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factom

import (
	"encoding/binary"
	"bytes"
	"crypto/sha256"
	"time"

	"golang.org/x/crypto/sha3"
)

var (
	server = "localhost:8088"
	serverFct = "localhost:8089"
)

func SetServer(s string) {
	server = s
}

// milliTime returns a 6 byte slice representing the unix time in milliseconds
func milliTime() (r []byte) {
	buf := new(bytes.Buffer)
	t := time.Now().UnixNano()
	m := t / 1e6
	binary.Write(buf, binary.BigEndian, m)
	return buf.Bytes()[2:]
}

// shad Double Sha256 Hash; sha256(sha256(data))
func shad(data []byte) []byte {
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}

// sha23 combination sha256 and sha3 Hash; sha256(data + sha3(data))
func sha23(data []byte) []byte {
	h1 := sha3.Sum256(data)
	h2 := sha256.Sum256(append(data, h1[:]...))
	return h2[:]
}

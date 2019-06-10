// Copyright 2019 Collin Kreklow
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package adsb

import (
	"encoding/hex"
	"testing"
)

// TestShort tests an incorrect length message
func TestShort(t *testing.T) {
	raw, err := hex.DecodeString("ff00")
	if err != nil {
		t.Error("received unexpected error", err)
	}
	msg := new(Message)
	err = msg.Decode(raw)
	if err.Error() != "invalid message length" {
		t.Error("received unexpected error", err)
	}
}

// TestUnknown tests an unknown message format
func TestUnknown(t *testing.T) {
	raw, err := hex.DecodeString("ff0000000000ff")
	if err != nil {
		t.Error("received unexpected error", err)
	}
	msg := new(Message)
	err = msg.Decode(raw)
	if err.Error() != "unsupported format: 31" {
		t.Error("received unexpected error", err)
	}
}

// Copyright 2019 Andrei Tudor Călin
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package mem_test

import (
	"testing"
)

func TestOffsetof(t *testing.T) {
	tests := []struct {
		Field  string
		Offset int
	}{
		{Field: "version", Offset: 0},
		{Field: "ihl", Offset: 4},
		{Field: "tos", Offset: 8},
		{Field: "frag_off", Offset: 48},
		{Field: "flags", Offset: 48},
		{Field: "offset", Offset: 51},
		{Field: "ttl", Offset: 64},
		{Field: "protocol", Offset: 72},
	}
	if err := ipv4Header.Init(); err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		got, ok := ipv4Header.Offsetof[tt.Field]
		if !ok {
			t.Fatalf("no field %q in layout", tt.Field)
		}
		if got != tt.Offset {
			t.Errorf("Offsetof[%q] = %d, want %d", tt.Field, got, tt.Offset)
		}
	}
}

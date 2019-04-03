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

package mem

import (
	"golang.org/x/exp/errors/fmt"
)

// Layout is a memory layout, composed of multiple fields. All field names
// and names in nested layouts must be unique for a given Layout.
//
// Callers should set Fields to a collection of fields, then call
// Init, which, if successful, causes the Offsetof map to be populated.
type Layout struct {
	Fields   []Field
	Offsetof map[string]int
}

// Field is a field in a memory layout.
type Field struct {
	Name   string // name of the field; required
	Desc   string // description; optional
	Bits   int    // width of the field, in bits; required
	Layout Layout // layout of the field itself; optional
}

// Reserved is a special field name to use for reserved or unused fields.
const Reserved = "__reserved"

// Init validates the Layout and populates the Offsetof map. Once Init has
// been called successfully on a Layout, subsequent calls have no effect.
func (l *Layout) Init() error {
	if l.Offsetof != nil {
		return nil
	}

	seen := map[string]struct{}{}
	offsetof := map[string]int{}

	var walk func(f Field, offset *int) error
	walk = func(f Field, offset *int) error {
		if _, ok := seen[f.Name]; ok {
			return duplicateFieldError(f.Name)
		}
		if f.Name != Reserved {
			seen[f.Name] = struct{}{}
		}

		offsetof[f.Name] = *offset
		defer func() { *offset += f.Bits }()

		if len(f.Layout.Fields) == 0 {
			return nil
		}

		childbits := 0
		for _, child := range f.Layout.Fields {
			childbits += child.Bits
		}

		if childbits != f.Bits {
			return childSizeError{childbits, f.Bits, f.Name}
		}

		tmpoffset := new(int)
		*tmpoffset = *offset
		for _, child := range f.Layout.Fields {
			if err := walk(child, tmpoffset); err != nil {
				return err
			}
		}
		return nil
	}

	offset := 0

	for _, f := range l.Fields {
		if err := walk(f, &offset); err != nil {
			return err
		}
	}
	l.Offsetof = offsetof
	return nil
}

type duplicateFieldError string

func (dfe duplicateFieldError) Error() string {
	return fmt.Sprintf("mem: duplicate field name %q", string(dfe))
}

type childSizeError struct {
	Got   int
	Want  int
	Field string
}

func (clse childSizeError) Error() string {
	return fmt.Sprintf("mem: field %q declares size %d, but child view has size %d", clse.Field, clse.Want, clse.Got)
}

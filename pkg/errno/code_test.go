/**
@author: yeebing
@date: 2022/9/24
**/

package errno

import (
	"errors"
	"reflect"
	"testing"
)

// TestMustRegister test Register and MustRegister
func TestMustRegister(t *testing.T) {
	t.Run("must register success", func(t *testing.T) {
		coder := New(100, 200, "success")
		MustRegister(coder)
	})

	t.Run("register repeated success", func(t *testing.T) {
		coder := New(100, 200, "override")
		Register(coder)
	})

}

// TestParseCoder test ParseCoder and WithCode, WithCodef
func TestParseCoder(t *testing.T) {
	t.Run("parse nil error", func(t *testing.T) {
		var err error = nil
		got := ParseCoder(err)
		want := Coder{}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %#+v want %#+v, given %#+v", got, want, err)
		}
	})

	t.Run("parse registered code", func(t *testing.T) {
		code := 200

		err := WithCode(code, errors.New("registered error"))
		want := New(code, 500, "registered error")
		Register(want)
		got := ParseCoder(err)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("got %#+v want %#+v, given %#+v", got, want, err)
		}
	})

	t.Run("parse unregistered code", func(t *testing.T) {
		err := WithCode(400, errors.New("unregistered error"))
		want := unknownCoder
		got := ParseCoder(err)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("got %#+v want %#+v, given %#+v", got, want, err)
		}
	})

	t.Run("parse normal error without code", func(t *testing.T) {
		err := errors.New("normal code")
		want := unknownCoder
		got := ParseCoder(err)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("got %#+v want %#+v, given %#+v", got, want, err)
		}
	})

}

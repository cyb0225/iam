/**
@author: yeebing
@date: 2022/9/24
**/

package errno

import (
	"errors"
	pkgerr "github.com/pkg/errors"

	"testing"
)

// TestStackTrace test print error's stack trace.
func TestStackError(t *testing.T) {
	err := StackA()
	t.Logf("%+v", StackError(err))
}

func StackA() error {
	err := StackB()
	return Wrap(err, "got err in StackA")
}

func StackB() error {
	err := WithCodef(200, "init err in StackB")
	return err
}

func TestWrap(t *testing.T) {
	err := pkgerr.New("row error")
	got := Wrap(err, "with wrap")
	gotC := Wrap(err, "with wrap")

	t.Logf("%+v", got.Error())
	t.Logf("%+v", gotC.Error())
}

func TestUnwrap(t *testing.T) {
	err := pkgerr.New("row error")
	msg := "with wrap"

	t.Run("test unwrap normal error", func(t *testing.T) {
		got := Wrap(err, msg)
		want := pkgerr.Wrap(err, msg)
		if errors.Is(got, want) {
			t.Fatalf("got %v want %v, given %v", got, want, err)
		}
	})

	t.Run("test unwrap code error", func(t *testing.T) {
		got := WithCode(100, err)
		got = Wrap(got, msg)
		want := pkgerr.Wrap(err, msg)

		if errors.Is(got, want) {
			t.Fatalf("got %v want %v, given %v", got, want, err)
		}
	})

}

// TestIs TODO:
func TestIs(t *testing.T) {
	t.Run("normal error compare with normal error", func(t *testing.T) {
		original := errors.New("normal error")
		processed := Wrap(original, "processed")
		got := Is(original, processed)
		t.Logf("got %v", got)
		want := errors.Is(original, processed)
		if got != want {
			t.Fatalf("got %v want %v, given %v, %v", got, want, original, processed)
		}
	})

	t.Run("code error compare with normal error", func(t *testing.T) {
		original := errors.New("normal error")
		processed := Wrap(WithCode(200, original), "processed")
		if !Is(original, processed) {
			t.Fatalf("got %v want %v, given %v, %v", false, true, original, processed)
		}
	})

	t.Run("code error compare with code error", func(t *testing.T) {
		original := errors.New("normal error")
		processed1 := Wrap(WithCode(200, original), "processed 1")
		processed2 := Wrap(WithCode(200, original), "processed 2")
		if !Is(processed1, processed2) {
			t.Fatalf("got %v want %v, given %v, %v", false, true, processed1, processed2)
		}
	})
}

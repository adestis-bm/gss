package gss

import (
	"reflect"
	"testing"
)

func TestGSS(t *testing.T) {
	txt := "Hello, World!"
	bytes := []byte(txt)

	ss, key, err := New(txt)
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	defer ss.Destroy()

	if reflect.DeepEqual(bytes, ss.sealed) {
		t.Fatalf("string and sealed string do match")
	}

	res, err := ss.String(key)
	if err != nil {
		t.Fatalf("String() failed with: %s", err)
	}

	if res != txt {
		t.Fatalf("String() is %q, expected %q", res, txt)
	}
}

func genStrings() []string {
	return []string{
		"Hello, World!",
		"Another Test",
		"I do not care",
		"733t",
	}
}

func TestGSSMulti(t *testing.T) {
	for _, txt := range genStrings() {
		bytes := []byte(txt)

		ss, key, err := New(txt)
		if err != nil {
			t.Fatalf("New() failed with: %s", err)
		}
		defer ss.Destroy()

		if reflect.DeepEqual(bytes, ss.sealed) {
			t.Fatalf("string and sealed string do match")
		}

		res, err := ss.String(key)
		if err != nil {
			t.Fatalf("String() failed with: %s", err)
		}

		if res != txt {
			t.Fatalf("String() is %q, expected %q", res, txt)
		}
	}
}

func TestGSSReuse(t *testing.T) {
	ss0, key, err := New("")
	if err != nil {
		t.Fatalf("New() failed with: %s", err)
	}
	defer ss0.Destroy()

	for _, txt := range genStrings() {
		bytes := []byte(txt)

		ss, err := key.New(txt)
		if err != nil {
			t.Fatalf("New() failed with: %s", err)
		}
		defer ss.Destroy()

		if reflect.DeepEqual(bytes, ss.sealed) {
			t.Fatalf("string and sealed string do match")
		}

		res, err := ss.String(key)
		if err != nil {
			t.Fatalf("String() failed with: %s", err)
		}

		if res != txt {
			t.Fatalf("String() is %q, expected %q", res, txt)
		}
	}
}

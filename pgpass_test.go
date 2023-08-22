package pgpassfile

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParsePassFile(t *testing.T) {
	buf := bytes.NewBufferString(`# A comment
	test1:5432:larrydb:larry:whatstheidea
	test1:5432:moedb:moe:imbecile
	test1:5432:curlydb:curly:nyuknyuknyuk
	test2:5432:*:shemp:heymoe
	test2:5432:*:*:test\\ing\:
	localhost:*:*:*:sesam
		`)

	passfile, err := ParsePassfile(buf)
	requireNil(t, err)

	assertEqual(t, len(passfile.Entries), 6)

	assertEqual(t, "whatstheidea", passfile.FindPassword("test1", "5432", "larrydb", "larry"))
	assertEqual(t, "imbecile", passfile.FindPassword("test1", "5432", "moedb", "moe"))
	assertEqual(t, `test\ing:`, passfile.FindPassword("test2", "5432", "something", "else"))
	assertEqual(t, "sesam", passfile.FindPassword("localhost", "9999", "foo", "bare"))

	assertEqual(t, "", passfile.FindPassword("wrong", "5432", "larrydb", "larry"))
	assertEqual(t, "", passfile.FindPassword("test1", "wrong", "larrydb", "larry"))
	assertEqual(t, "", passfile.FindPassword("test1", "5432", "wrong", "larry"))
	assertEqual(t, "", passfile.FindPassword("test1", "5432", "larrydb", "wrong"))
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v but got: %#v", expected, actual)
	}
}

func requireNil(t *testing.T, object interface{}) {
	if object != nil {
		t.Fatalf("Expected nil, but got: %#v", object)
	}
}

package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test value"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test value"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAddKey(t *testing.T) {

	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		definition := "test value"

		err := dictionary.AddKey(key, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, definition)
	})

	t.Run("existing key", func(t *testing.T) {
		key := "test"
		definition := "test value"
		dictionary := Dictionary{key: definition}
		err := dictionary.AddKey(key, "new test")

		assertError(t, err, ErrKeyExists)
		assertDefinition(t, dictionary, key, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{key: definition}
		newDefinition := "new definition"

		err := dictionary.Update(key, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newDefinition)
	})

	t.Run("new key", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, definition)

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "test definition"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	assertError(t, err, ErrNotFound)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added key:", err)
	}
	assertStrings(t, got, definition)
}

package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("Word not found", func(t *testing.T) {
		_, err := dictionary.Search("notFound")

		assertStrings(t, err.Error(), errNotFound.Error())
		if err == nil {
			t.Errorf("Expected Error")
		}
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "added")

	got, err := dictionary.Search("test")
	want := "added"

	if err != nil {
		t.Errorf("Expect Error not to have occured")
	}

	assertStrings(t, got, want)

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

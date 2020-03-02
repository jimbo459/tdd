package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name     string
	location Location
}

type Location struct {
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string",
			struct {
				Name string
			}{"James"},
			[]string{"James"},
		},
		{
			"Struct with two strings",
			struct {
				Name    string
				Surname string
			}{"James", "Norman"},
			[]string{"James", "Norman"},
		},
		{
			"Struct with differing types",
			struct {
				Name string
				Age  int
			}{"James", 30},
			[]string{"James"},
		},
		{
			"Struct with nested fields",
			Person{
				"James",
				Location{"Amersham"},
			},
			[]string{"James", "Amersham"},
		},
		{
			"Struct passed in as pointer",
			&Person{
				"James",
				Location{"Amersham"},
			},
			[]string{"James", "Amersham"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(test.ExpectedCalls, got) {
				t.Errorf("Walk return value incorrect. Expected %v, got %v", test.ExpectedCalls, got)
			}
		})
	}
}

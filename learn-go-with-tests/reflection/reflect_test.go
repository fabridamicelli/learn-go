package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},

		{
			"struct with 2 string field",
			struct {
				Name string
				City string
			}{"Chris", "HH"},
			[]string{"Chris", "HH"},
		},

		{
			"struct with string and int field",
			struct {
				Name string
				Age  int
			}{"Chris", 21},
			[]string{"Chris"},
		},
		{
			"struct with nested struct",
			Person{
				"Chris",
				Profile{33, "HH"},
			},
			[]string{"Chris", "HH"},
		},

		{
			"pointers",
			&Person{
				"Chris",
				Profile{33, "HH"},
			},
			[]string{"Chris", "HH"},
		},
		{
			"slices",
			[]Profile{
				{32, "London"},
				{33, "HH"},
			},
			[]string{"London", "HH"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "HH"},
			},
			[]string{"London", "HH"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"cow":   "moo",
			"sheep": "meeh",
		}
		var got []string
		walk(aMap, func(inp string) {
			got = append(got, inp)
		})

		assertContains(t, got, "moo")
		assertContains(t, got, "meeh")

	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

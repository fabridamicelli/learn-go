package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"blogposts"
)

func assertPost(t *testing.T, got blogposts.Post, exp blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("got %+v, want %+v", got, exp)
	}

}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Blogpost content
a1
a2
a3`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Blogpost content
b1
b2
b3`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Blogpost content
a1
a2
a3`,
	})

	assertPost(t, posts[1], blogposts.Post{
		Title:       "Post 2",
		Description: "Description 2",
		Tags:        []string{"rust", "borrow-checker"},
		Body: `Blogpost content
b1
b2
b3`,
	})

}

func TestNewBlogPostsWithWrongBody(t *testing.T) {
	const (
		wrongBodySep = `Title: foo
Description: bar
Tags: foo, bar
--
Blogpost content
`
	)

	fs := fstest.MapFS{
		"wrong.md": {Data: []byte(wrongBodySep)},
	}

	_, err := blogposts.NewPostsFromFS(fs)
	if err == nil {
		t.Fatal("want error for invalid body separator")
	}
}

package blogposts

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil

}

func getPost(fileSystem fs.FS, fname string) (Post, error) {
	postFile, err := fileSystem.Open(fname)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

const (
	titleSep = "Title: "
	descSep  = "Description: "
	tagsSep  = "Tags: "
)

func readBody(scanner *bufio.Scanner) (string, error) {
	scanner.Scan()
	bodySep := scanner.Text()
	if bodySep != "---" {
		return "", errors.New(fmt.Sprintf("Wrong body separator, expected ---, got %s", bodySep))
	}
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	body := strings.Join(lines, "\n")
	return body, nil
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)

	}
	title := readMetaLine(titleSep)
	desc := readMetaLine(descSep)
	tags := strings.Split(readMetaLine(tagsSep), ", ")
	body, err := readBody(scanner)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: title, Description: desc, Tags: tags, Body: body}
	return post, nil
}

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc) // get all <a> tag

	var links []Link
	for _, n := range nodes {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, Link{a.Val, n.FirstChild.Data})
			}
		}
	}
	return links, err
}

func linkNodes(n *html.Node) []*html.Node {
	// donot care nested tag
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...) // variadic functions
	}
	return nodes
}

func main() {
	htmls := []string{
		"ex1.html",
		"ex2.html",
		"ex3.html",
		"ex4.html",
	}
	for _, html := range htmls {
		fmt.Println("=====", html, "=====")
		r := readFile(html)
		links, err := Parse(r)
		if err != nil {
			panic(err)
		}
		for _, l := range links {
			fmt.Println(l)
		}
	}
}

func readFile(filename string) io.Reader {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	r := strings.NewReader(string(file))
	return r
}

package main

import "fmt"

type Markdown string
type HTML string
type Path string
type Title string

type Node interface {
	Value() string
	Children() []Node
	RenderTitle(int) string // is there no other way to do this?
}

type Folder struct {
	Name     string
	Contents []Node
}

func (f Folder) Value() string {
	return f.Name
}

func (f Folder) Children() []Node {
	return f.Contents
}

func (f Folder) RenderTitle(depth int) string {
	return fmt.Sprintf("<li><h%d>%s</h%d></li>\n", depth, f.Name, depth)
}

type Page struct {
	Title string
	Body  string
}

func (p Page) Value() string {
	return p.Title
}

func (p Page) Children() []Node {
	return []Node{}
}

func (p Page) RenderTitle(_ int) string {
	return fmt.Sprintf("<li>%s</li>\n", p.Title)
}

func RenderIndex(start Node) string {

	var renderIndex func(int, Node) string
	renderIndex = func(depth int, n Node) string {

		renderedChildren := ""
		for _, c := range n.Children() {
			renderedChildren += renderIndex(depth+1, c)
		}

		return n.RenderTitle(depth) + renderedChildren
	}
	return renderIndex(1, start)
}

func main() {
	fmt.Println("hello?")
}

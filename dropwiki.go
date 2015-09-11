package main

import "fmt"

type Node interface {
	Value() string
	Children() []Node
	renderIndex(int) string
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

func (p Page) renderIndex(d int) string {
	return fmt.Sprintf("<li>%s</li>\n", p.Title)
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

func (f Folder) renderIndex(startingDepth int) string {
	renderedChildren := ""
	for _, c := range f.Children() {
		renderedChildren += c.renderIndex(startingDepth + 1)
	}

	renderedTitle := fmt.Sprintf("<li><h%d>%s</h%d></li>\n", startingDepth, f.Name, startingDepth)
	return renderedTitle + renderedChildren
}

func RenderIndex(start Node) string {
	return start.renderIndex(1)
}

func main() {
	fmt.Println("hello?")
}

package htmlparser

import (
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Url  string
	Text string
}

func Run() []Link {

	file, err := os.OpenFile("html_parser/ex1.html", os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(err.Error())
	}

	node, _ := html.Parse(file)

	links := Dfs(node)

	return links
}

func Dfs(node *html.Node) []Link {

	var links []Link

	link := Link{}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, items := range node.Attr {
			link.Url = items.Val
			break
		}
		link.Text = getText(node)
		links = append(links, link)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, Dfs(c)...)
	}
	return links

}

func getText(node *html.Node) string {

	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}

	var text string

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}
	return text
}

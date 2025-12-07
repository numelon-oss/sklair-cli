package htmlUtilities

import "golang.org/x/net/html"

func Clone(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}

	// hah
	clown := &html.Node{
		Type:     n.Type,
		DataAtom: n.DataAtom,
		Data:     n.Data,
		Attr:     append([]html.Attribute{}, n.Attr...),
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		clown.AppendChild(Clone(child))
	}

	return clown
}

func FindTag(n *html.Node, tag string) *html.Node {
	if n.Type == html.ElementNode && n.Data == tag {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := FindTag(c, tag); found != nil {
			return found
		}
	}

	return nil
}

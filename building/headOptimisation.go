package building

import (
	"sklair/htmlUtilities"

	"golang.org/x/net/html"
)

func OptimiseHead(head *html.Node) {
	DeduplicateHeadPass(head)
}

func DeduplicateHeadPass(head *html.Node) {
	seenHead := make(map[uint64]struct{})

	var toRemove []*html.Node

	for c := head.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}

		key := htmlUtilities.WeakHashNode(c)
		if key == 0 {
			continue
		}

		if _, seen := seenHead[key]; seen {
			//head.RemoveChild(c) // WE'RE GOING OVER A LINKED LIST - CANNOT MUTATE IT IMMEDIATELY!!!
			toRemove = append(toRemove, c)
			continue
		}
		seenHead[key] = struct{}{}
	}

	for _, node := range toRemove {
		head.RemoveChild(node)
	}
}

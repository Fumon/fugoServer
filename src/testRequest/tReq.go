package main

import (
	"fmt"
	"html"
	"http"
	"regexp"
)

var nullLine = regexp.MustCompile("\n")

/*
type logEntry struct {

}

func storeLog(h *html.Node) {

}
*/

func saveLogT(h *html.Node) {
	//Print the button contents.
	fmt.Println(h.Child[0].Child[0].Data)
	var fn func(*html.Node)
	fn = func(hn *html.Node) {
		if hn.Type == html.TextNode && !nullLine.MatchString(hn.Data){
			//PRINT!
			fmt.Printf("%v\n", hn.Data)
		}
		for _, v := range hn.Child {
			fn(v)
		}
	}

	fn(h.Child[0].Child[2])
}

func main() {
	res, _, err := http.Get("http://www.mspaintadventures.com/?s=6&p=005366")
	if err != nil {
		fmt.Printf("Error detected as %s", err)
		return
	}

	z, err := html.Parse(res.Body)
	if err != nil {
		fmt.Printf("Error while parsing: %v", err)
		return
	}

	foundSpoiler := false

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" && len(n.Attr) != 0 {
			//Check to see if it's the spoiler div
			at0 := n.Attr[0]
			if at0.Key == "class" && at0.Val == "spoiler" {
				//It's the spoiler div, store the stuff.
				fmt.Println("Found a spoiler div!")
				//storeLog(n)
				saveLogT(n)
				foundSpoiler = true
			}
		} else if n.Type == html.ElementNode && n.Data == "table" && len(n.Attr) != 0 {
			if n.Attr[0].Key == "width" && n.Attr[1].Val == "90%" {
				fmt.Println("I found what looks like the table")
			}
		}
		for _, c := range n.Child {
			f(c)
		}
	}
	f(z)
}

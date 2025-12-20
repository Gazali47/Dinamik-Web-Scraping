package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

func extractLinks(htmlContent string) []string {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil
	}
	linksMap := make(map[string]bool)
	var links []string
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.HasPrefix(attr.Val, "http") {
					if !linksMap[attr.Val] {
						linksMap[attr.Val] = true
						links = append(links, attr.Val)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)
	return links
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <URL>")
		return
	}

	rawURL := os.Args[1]
	fileNameBase := strings.ReplaceAll(strings.ReplaceAll(rawURL, "https://", ""), "http://", "")
	fileNameBase = strings.ReplaceAll(strings.ReplaceAll(fileNameBase, "/", "_"), ".", "_")

	fmt.Printf("Hedef URL işleniyor: %s\n", rawURL)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	var htmlContent string
	var buf []byte
	var statusCode int64

	err := chromedp.Run(ctx,
		chromedp.Navigate(rawURL),
		chromedp.Evaluate(`performance.getEntriesByType("navigation")[0].responseStatus`, &statusCode),
		chromedp.OuterHTML("html", &htmlContent),
		chromedp.FullScreenshot(&buf, 90),
	)

	if err != nil {
		log.Fatalf("Hata oluştu: %v", err)
	}

	htmlFileName := fileNameBase + ".html"
	err = os.WriteFile(htmlFileName, []byte(htmlContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	imgFileName := fileNameBase + ".png"
	err = os.WriteFile(imgFileName, buf, 0644)
	if err != nil {
		log.Fatal(err)
	}

	links := extractLinks(htmlContent)
	linksFileName := fileNameBase + "_links.txt"
	err = os.WriteFile(linksFileName, []byte(strings.Join(links, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Durum Kodu: %d\n", statusCode)
	fmt.Printf("Kaydedilen Dosyalar:\n - %s\n - %s\n - %s\n", htmlFileName, imgFileName, linksFileName)
}

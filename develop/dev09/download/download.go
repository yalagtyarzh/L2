package download

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func Store(dest string, url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	doc.Find("img").Each(findElement("src"))
	doc.Find("script").Each(findElement("src"))
	doc.Find("link").Each(findElement("href"))
	return nil
}

func findElement(attr string) func(i int, selection *goquery.Selection) {
	return func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr(attr)
		// БОЛЬШОЕ СПАСИБО GOQUERY ЗА ВАЛИДАЦИЮ ПО ПУСТОЙ СТРОКЕ А НЕ ПО OK, XD
		if src == "" {
			return
		}

		r := []rune(src)
		sub := strings.Split(src, ".")

		if r[0] == '/' && r[0] != r[1] && sub[len(sub)-1] != "html" {
			fmt.Println(src)
		}
	}
}

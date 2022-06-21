package download

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Download(dest string, url string) error {
	fname, err := download(dest, url)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fname, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return err
	}

	resources := make([]string, 0)

	doc.Find("img").Each(findElement("src", &resources))
	doc.Find("script").Each(findElement("src", &resources))
	doc.Find("link").Each(findElement("rel", &resources))

	for _, resource := range resources {
		_, err = download(dest, url+resource)
		if err != nil {
			return err
		}
	}

	return nil
}

func download(dest string, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//dest = addSlash(dest)
	path := []string{dest}
	divided := strings.Split(resp.Request.URL.String(), "/")
	path = append(path, divided[2:]...)

	err = os.MkdirAll(strings.Join(path[:len(path)-1], "/"), 0777)

	var file *os.File
	if resp.Header.Get("content-type") == "text/html" {
		file, err = os.Create(strings.Join(path, "/") + "index.html")
	} else {
		file, err = os.Create(strings.Join(path, "/"))
	}

	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	return file.Name(), nil
}

func findElement(attr string, output *[]string) func(i int, selection *goquery.Selection) {
	return func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr(attr)
		if src == "stylesheet" {
			src, _ = selection.Attr("href")
		}

		// БОЛЬШОЕ СПАСИБО GOQUERY ЗА ВАЛИДАЦИЮ ПО ПУСТОЙ СТРОКЕ А НЕ ПО OK, XD
		if src == "" {
			return
		}

		r := []rune(src)

		if r[0] == '/' && r[0] != r[1] {
			*output = append(*output, src)
			selection.SetAttr(attr, "."+src)
		}
	}
}

//func addSlash(dest string) string {
//	runes := []rune(dest)
//	if runes[len(runes)-1] != '/' {
//		runes = append(runes, '/')
//	}
//
//	dest = string(runes)
//	return dest
//}

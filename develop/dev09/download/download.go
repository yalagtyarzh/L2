package download

import (
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var wg sync.WaitGroup

func Download(dest string, url string) error {
	wg.Add(1)
	filename, err := download(dest, url)
	if err != nil {
		return err
	}
	wg.Wait()

	f, err := os.OpenFile(filename, os.O_RDWR, os.FileMode(0644))
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
	doc.Find("source").Each(findElement("srcset", &resources))

	s, err := doc.Html()
	if err != nil {
		return err
	}

	err = f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = f.WriteString(s)
	if err != nil {
		return err
	}

	for _, resource := range resources {
		wg.Add(1)
		r := resource
		go func() {
			_, err := download(dest, url+r)
			if err != nil {
				return
			}
		}()
	}
	wg.Wait()

	return nil
}

func download(dest string, url string) (string, error) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	path := []string{dest}
	divided := strings.Split(resp.Request.URL.String(), "/")
	path = append(path, divided[2:]...)

	err = os.MkdirAll(strings.Join(path[:len(path)-1], "/"), 0777)

	var file *os.File
	if strings.Contains(resp.Header.Get("content-type"), "text/html") {
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
		a := attr
		src, _ := selection.Attr(attr)
		if a == "rel" {
			if src == "stylesheet" {
				a = "href"
				src, _ = selection.Attr(a)
			}
		}

		if src == "" {
			return
		}

		r := []rune(src)

		if r[0] == '/' && r[0] != r[1] {
			*output = append(*output, src)
			selection.SetAttr(a, "."+src)
		}
	}
}

package download

import (
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// wg для синхронизации горутин
var wg sync.WaitGroup

// Download скачивает сайт с интернета по определенному url и помещает его в заданном dest
// После в файле сайта преобразуют все абсолютные ссылки на ресурсы страницы в относительные, после происходит скачка
// найденных ресурсов
func Download(dest string, url string) error {
	wg.Add(1)
	// Скачиваем сайт и получаем его путь с наименованием
	filename, err := download(dest, url)
	if err != nil {
		return err
	}
	wg.Wait()

	// Открываем файл для чтения-записи
	f, err := os.OpenFile(filename, os.O_RDWR, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer f.Close()

	// Инициализируем парсер HTML в заданном файле
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return err
	}

	// resources хранит в себе ссылки ресурсов
	resources := make([]string, 0)

	// Ищем ресурсы в HTML заполняем массив resources
	doc.Find("img").Each(findElement("src", &resources))
	doc.Find("script").Each(findElement("src", &resources))
	doc.Find("link").Each(findElement("rel", &resources))
	doc.Find("source").Each(findElement("srcset", &resources))

	// Получаем преобразованный html документ
	s, err := doc.Html()
	if err != nil {
		return err
	}

	err = rewrite(f, s)
	if err != nil {
		return err
	}

	// Скачиваем ресурсы с помощью нескольких горутин
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
	// Ждем, когда все скачается
	wg.Wait()

	return nil
}

// download скачивает файл по url и помещает их по dest, возвращая путь скачанного файла
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
	// Если в ответе запроса имеется хедер content-type со значением text/html, то именуем полученный файл
	if strings.Contains(resp.Header.Get("content-type"), "text/html") {
		file, err = os.Create(strings.Join(path, "/") + "/index.html")
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

// findElement ищет все элементы с переданным аттрибутом attr, получает оттуда ссылки
// и данными ссылкам срез output. Так же преобразовывает абсолютные ссылки найденных элементов в относительные
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

		// Преобразовываем абсолютную ссылку в относительную
		if r[0] == '/' && r[0] != r[1] {
			*output = append(*output, src)
			selection.SetAttr(a, "."+src)
		}
	}
}

// rewrite перезаписывает файл file контентом, передающийся в input
func rewrite(file *os.File, input string) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.WriteString(input)
	if err != nil {
		return err
	}

	return nil
}

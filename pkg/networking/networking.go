package networking

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func (s *NetworkingService) CreateClient(jar *http.CookieJar) (*http.Client, error) {
	if jar == nil {
		cookie, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
		return &http.Client{Jar: cookie}, nil
	}

	return &http.Client{Jar: *jar}, nil
}

func (s *NetworkingService) GetPage(client *http.Client, url string) (*http.Response, *goquery.Document, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, doc, nil
}

func (s *NetworkingService) GetJsonToStruct(client *http.Client, url string, structure interface{}) (*http.Response, interface{}, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&structure); err != nil {
		return nil, nil, err
	}

	return res, structure, nil
}

func (s *NetworkingService) PostForm(client *http.Client, url string, values *url.Values) (*http.Response, *goquery.Document, error) {
	res, err := client.PostForm(url, *values)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, doc, nil
}

func (s *NetworkingService) CorrectLocation(res http.Response, targetUrl string) error {
	location := res.Request.URL.String()
	if location != targetUrl {
		return errors.New("incorrect location")
	}

	return nil
}

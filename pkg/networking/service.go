package networking

//go:generate mockgen -source=service.go -destination=service_mock.go -package=networking

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
)

type IFNetworkingService interface {
	// CreateClient : Create HTTP Client
	CreateClient(jar *http.CookieJar) (*http.Client, error)
	// GetPage : GET リクエストを送信し、レスポンスを返す
	GetPage(client *http.Client, url string) (*http.Response, *goquery.Document, error)
	// GetJsonToStruct : GET リクエストを送信し、レスポンスを構造体に変換する
	GetJsonToStruct(client *http.Client, url string, structure interface{}) (*http.Response, interface{}, error)
	// PostForm : POST リクエストを送信し、レスポンスを返す
	PostForm(client *http.Client, url string, values *url.Values) (*http.Response, *goquery.Document, error)
	// CorrectLocation : 正しい URL にいるかどうかを確認する
	CorrectLocation(res http.Response, targetUrl string) error
}

type ifPrivateNetworkingService interface {
}

type NetworkingService struct {
	privateService ifPrivateNetworkingService
}

func NewNetworkingService() *NetworkingService {
	srv := &NetworkingService{}
	srv.privateService = srv

	return srv
}

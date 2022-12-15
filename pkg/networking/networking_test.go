package networking

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
	"testing"
)

func TestNetworkingService_CreateClient(t *testing.T) {
	type Returns struct {
		client *http.Client
		err    error
	}
	type testContext struct {
		returns Returns
	}
	tests := []struct {
		name        string
		testContext testContext
	}{
		{
			name: "正常",
			testContext: testContext{
				returns: Returns{
					client: &http.Client{
						Jar: &cookiejar.Jar{},
					},
					err: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetworkingService{}
			gotClient, err := s.CreateClient(nil)
			if (err != nil) != (tt.testContext.returns.err != nil) {
				t.Errorf("NetworkingService.CreateClient() error = %v, wantErr %v", err, tt.testContext.returns.err)
				return
			}
			if gotClient == nil {
				t.Errorf("NetworkingService.CreateClient() gotClient = %v, want %v", gotClient, tt.testContext.returns.client)
			}
		})
	}
}

func TestNetworkingService_GetPage(t *testing.T) {
	type Args struct {
		client *http.Client
		url    string
	}
	type Returns struct {
		res *http.Response
		doc *goquery.Document
		err error
	}
	type testContext struct {
		args    Args
		returns Returns
	}
	tests := []struct {
		name        string
		testContext testContext
	}{
		{
			name: "正常",
			testContext: testContext{
				args: Args{
					client: &http.Client{},
					url:    "https://example.com/",
				},
				returns: Returns{
					res: &http.Response{},
					doc: &goquery.Document{},
					err: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetworkingService{}
			gotRes, gotDoc, err := s.GetPage(tt.testContext.args.client, tt.testContext.args.url)
			if (err != nil) != (tt.testContext.returns.err != nil) {
				t.Errorf("NetworkingService.GetPage() error = %v, wantErr %v", err, tt.testContext.returns.err)
				return
			}
			if gotRes == nil {
				t.Errorf("NetworkingService.GetPage() gotRes = %v, want %v", gotRes, tt.testContext.returns.res)
			}
			if gotDoc == nil {
				t.Errorf("NetworkingService.GetPage() gotDoc = %v, want %v", gotDoc, tt.testContext.returns.res)
			}
		})
	}
}

func TestNetworkingService_GetJsonToStruct(t *testing.T) {
	type Args struct {
		client    *http.Client
		url       string
		structure interface{}
	}
	type Returns struct {
		res       *http.Response
		structure interface{}
		err       error
	}
	type testContext struct {
		args    Args
		returns Returns
	}
	type testGetJsonToStructStructure struct {
		EntityPrefix         string        `json:"entityPrefix"`
		AssignmentCollection []interface{} `json:"assignment_collection"`
	}
	tests := []struct {
		name        string
		testContext testContext
	}{
		{
			name: "正常",
			testContext: testContext{
				args: Args{
					client:    &http.Client{},
					url:       "https://hoppii.hosei.ac.jp/direct/assignment/my.json",
					structure: &testGetJsonToStructStructure{},
				},
				returns: Returns{
					res: &http.Response{
						StatusCode: 200,
					},
					structure: &testGetJsonToStructStructure{
						EntityPrefix:         "assignment",
						AssignmentCollection: []interface{}{},
					},
					err: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetworkingService{}
			gotRes, gotStructure, err := s.GetJsonToStruct(tt.testContext.args.client, tt.testContext.args.url, tt.testContext.args.structure)
			if (err != nil) != (tt.testContext.returns.err != nil) {
				t.Errorf("NetworkingService.GetJsonToStruct() error = %v, wantErr %v", err, tt.testContext.returns.err)
				return
			}
			if gotRes == nil {
				t.Errorf("NetworkingService.GetJsonToStruct() gotRes = %v, want %v", gotRes, tt.testContext.returns.res)
			}
			// Body は Close されると無くなるので、非破壊の Status Code で比較する
			if gotRes.StatusCode != tt.testContext.returns.res.StatusCode {
				t.Errorf("NetworkingService.GetJsonToStruct() gotRes = %v, want %v", gotRes, tt.testContext.returns.res)
			}
			if gotStructure == nil {
				t.Errorf("NetworkingService.GetJsonToStruct() gotStructure = %v, want %v", gotStructure, tt.testContext.returns.structure)
			}
			if !reflect.DeepEqual(gotStructure, tt.testContext.returns.structure) {
				t.Errorf("NetworkingService.GetJsonToStruct() gotStructure = %v, want %v", gotStructure, tt.testContext.returns.structure)
			}
		})
	}
}

func TestNetworkingService_PostForm(t *testing.T) {
	type Args struct {
		client *http.Client
		url    string
		data   *url.Values
	}
	type Returns struct {
		res *http.Response
		doc *goquery.Document
		err error
	}
	type testContext struct {
		args    Args
		returns Returns
	}
	tests := []struct {
		name        string
		testContext testContext
	}{
		{
			name: "正常",
			testContext: testContext{
				args: Args{
					client: &http.Client{},
					url:    "https://example.com/",
					data:   &url.Values{},
				},
				returns: Returns{
					res: &http.Response{
						StatusCode: 200,
					},
					doc: &goquery.Document{},
					err: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetworkingService{}
			gotRes, gotDoc, err := s.PostForm(tt.testContext.args.client, tt.testContext.args.url, tt.testContext.args.data)
			if (err != nil) != (tt.testContext.returns.err != nil) {
				t.Errorf("NetworkingService.PostForm() error = %v, wantErr %v", err, tt.testContext.returns.err)
				return
			}
			if gotRes == nil {
				t.Errorf("NetworkingService.PostForm() gotRes = %v, want %v", gotRes, tt.testContext.returns.res)
			}
			// Body は Close されると無くなるので、非破壊の Status Code で比較する
			if gotRes.StatusCode != tt.testContext.returns.res.StatusCode {
				t.Errorf("NetworkingService.PostForm() gotRes = %v, want %v", gotRes, tt.testContext.returns.res)
			}
			if gotDoc == nil {
				t.Errorf("NetworkingService.PostForm() gotDoc = %v, want %v", gotDoc, tt.testContext.returns.res)
			}
		})
	}
}

func TestNetworkingService_CorrectLocation(t *testing.T) {
	type Args struct {
		res       http.Response
		targetUrl string
	}
	type Returns struct {
		err error
	}
	type testContext struct {
		args    Args
		returns Returns
	}
	tests := []struct {
		name        string
		testContext testContext
	}{
		{
			name: "正常",
			testContext: testContext{
				args: Args{
					res: http.Response{
						Request: &http.Request{
							URL: &url.URL{
								Scheme: "https",
								Host:   "example.com",
								Path:   "/",
							},
						},
					},
					targetUrl: "https://example.com/",
				},
				returns: Returns{
					err: nil,
				},
			},
		},
		{
			name: "異常: URL が異なる",
			testContext: testContext{
				args: Args{
					res: http.Response{
						Request: &http.Request{
							URL: &url.URL{
								Scheme: "https",
								Host:   "example.com",
								Path:   "/",
							},
						},
					},
					targetUrl: "https://example.com/other",
				},
				returns: Returns{
					err: errors.New("incorrect location"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NetworkingService{}
			if err := s.CorrectLocation(tt.testContext.args.res, tt.testContext.args.targetUrl); (err != nil) != (tt.testContext.returns.err != nil) {
				t.Errorf("NetworkingService.CorrectLocation() error = %v, wantErr %v", err, tt.testContext.returns.err)
			}
		})
	}
}

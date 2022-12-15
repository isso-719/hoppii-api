package user

import (
	"errors"
	"net/url"
)

// TODO: Cookie がファイルとして渡された時の対処
func (s *UserService) CreateUser(input *UserInput) (*UserOutput, error) {
	cookie, err := s.privateService.login(input.StudentId, input.Password)
	if err != nil {
		return nil, err
	}

	return &UserOutput{
		StudentId: input.StudentId,
		Password:  input.Password,
		Cookie:    cookie,
	}, nil
}

// TODO : Cookie　がファイルとして渡された時の対処
func (s *UserService) login(username, password string) (*UserCookie, error) {
	client, err := s.networkingService.CreateClient(nil)
	if err != nil {
		return nil, err
	}

	_, doc, err := s.networkingService.GetPage(client, "https://hoppii.hosei.ac.jp/portal/login")
	if err != nil {
		return nil, err
	}

	action, ok := doc.Find("form").Attr("action")
	if !ok {
		return nil, errors.New("login failed (action attribute not found): check your credentials or network")
	}
	values := &url.Values{
		"j_username":       {username},
		"j_password":       {password},
		"_eventId_proceed": {""},
	}
	actionUrl := "https://idp.hosei.ac.jp" + action
	_, doc, err = s.networkingService.PostForm(client, actionUrl, values)
	if err != nil {
		return nil, err
	}

	action, ok = doc.Find("form").Attr("action")
	if !ok {
		return nil, errors.New("login failed (action attribute not found): check your credentials or network")
	}
	relayState, ok := doc.Find("input[name=RelayState]").Attr("value")
	if !ok {
		return nil, errors.New("login failed (RelayState not found): check your credentials or network")
	}
	samlResponse, ok := doc.Find("input[name=SAMLResponse]").Attr("value")
	if !ok {
		return nil, errors.New("login failed (SAMLResponse not found): check your credentials or network")
	}
	values = &url.Values{
		"RelayState":   {relayState},
		"SAMLResponse": {samlResponse},
	}
	res, _, err := s.networkingService.PostForm(client, action, values)
	if err != nil {
		return nil, err
	}

	if err := s.networkingService.CorrectLocation(*res, "https://hoppii.hosei.ac.jp:443/portal"); err != nil {
		return nil, errors.New("login failed (incorrect location): check your credentials or network")
	}

	return &UserCookie{
		Cookie: client.Jar,
	}, nil
}

// TODO: サービスに追加
func (s *UserService) SaveUserCookieFile() {

}

package announcement

// TODO: オプションの考慮
func (s *AnnouncementService) User(input *AnnouncementUserInput) (*AnnouncementUserOutput, error) {
	client, err := s.networkingService.CreateClient(&input.UserCookie.Cookie)
	if err != nil {
		return nil, err
	}

	_, aur, err := s.networkingService.GetJsonToStruct(client, "https://hoppii.hosei.ac.jp/direct/announcement/user.json", &announcementUserResult{})
	if err != nil {
		return nil, err
	}

	return &AnnouncementUserOutput{
		AnnouncementUserResult: aur.(*announcementUserResult),
	}, nil
}

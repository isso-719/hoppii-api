package assignment

func (s *AssignmentService) My(input *AssignmentMyInput) (*AssignmentMyOutput, error) {
	client, err := s.networkingService.CreateClient(&input.UserCookie.Cookie)
	if err != nil {
		return nil, err
	}

	_, amr, err := s.networkingService.GetJsonToStruct(client, "https://hoppii.hosei.ac.jp/direct/assignment/my.json", &assignmentMyResult{})
	if err != nil {
		return nil, err
	}

	return &AssignmentMyOutput{
		AssignmentMyResult: amr.(*assignmentMyResult),
	}, nil
}

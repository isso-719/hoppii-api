package assignment

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/isso-719/hoppii-api/pkg/networking"
	"github.com/isso-719/hoppii-api/pkg/user"
	"reflect"
	"testing"
	"time"
)

func TestAssignmentService_My(t *testing.T) {
	type Fields struct {
		networkingService networking.IFNetworkingService
	}
	type Args struct {
		input *AssignmentMyInput
	}
	type Returns struct {
		output *AssignmentMyOutput
	}
	type testContext struct {
		fields  Fields
		args    Args
		returns Returns
	}

	tests := []struct {
		name        string
		testContext func(ctrl *gomock.Controller) *testContext
	}{
		{
			name: "正常",
			testContext: func(ctrl *gomock.Controller) *testContext {
				input := &AssignmentMyInput{
					UserCookie: &user.UserCookie{},
				}

				networkingService := networking.NewMockIFNetworkingService(ctrl)
				networkingService.EXPECT().CreateClient(gomock.Any()).Return(nil, nil)
				networkingService.EXPECT().GetJsonToStruct(
					gomock.Any(),
					"https://hoppii.hosei.ac.jp/direct/assignment/my.json",
					&assignmentMyResult{},
				).Return(
					nil,
					&assignmentMyResult{
						EntityPrefix: "assignment",
						AssignmentCollection: []struct {
							Access             string        `json:"access"`
							AllPurposeItemText interface{}   `json:"allPurposeItemText"`
							Attachments        []interface{} `json:"attachments"`
							Author             string        `json:"author"`
							AuthorLastModified interface{}   `json:"authorLastModified"`
							CloseTime          struct {
								EpochSecond int `json:"epochSecond"`
								Nano        int `json:"nano"`
							} `json:"closeTime"`
							CloseTimeString time.Time   `json:"closeTimeString"`
							Content         interface{} `json:"content"`
							Context         string      `json:"context"`
							Creator         interface{} `json:"creator"`
							DropDeadTime    struct {
								EpochSecond int `json:"epochSecond"`
								Nano        int `json:"nano"`
							} `json:"dropDeadTime"`
							DropDeadTimeString time.Time `json:"dropDeadTimeString"`
							DueTime            struct {
								EpochSecond int `json:"epochSecond"`
								Nano        int `json:"nano"`
							} `json:"dueTime"`
							DueTimeString       time.Time     `json:"dueTimeString"`
							GradeScale          string        `json:"gradeScale"`
							GradeScaleMaxPoints string        `json:"gradeScaleMaxPoints"`
							GradebookItemID     interface{}   `json:"gradebookItemId"`
							GradebookItemName   interface{}   `json:"gradebookItemName"`
							Groups              []interface{} `json:"groups"`
							ID                  string        `json:"id"`
							Instructions        string        `json:"instructions"`
							ModelAnswerText     interface{}   `json:"modelAnswerText"`
							OpenTime            struct {
								EpochSecond int `json:"epochSecond"`
								Nano        int `json:"nano"`
							} `json:"openTime"`
							OpenTimeString  time.Time   `json:"openTimeString"`
							Position        int         `json:"position"`
							PrivateNoteText interface{} `json:"privateNoteText"`
							Section         string      `json:"section"`
							Status          string      `json:"status"`
							SubmissionType  string      `json:"submissionType"`
							TimeCreated     struct {
								EpochSecond int `json:"epochSecond"`
								Nano        int `json:"nano"`
							} `json:"timeCreated"`
							TimeLastModified struct {
								EpochSecond int `json:"epochSecond"`
								Nano        int `json:"nano"`
							} `json:"timeLastModified"`
							Title             string `json:"title"`
							AllowResubmission bool   `json:"allowResubmission"`
							Draft             bool   `json:"draft"`
							EntityReference   string `json:"entityReference"`
							EntityURL         string `json:"entityURL"`
							EntityID          string `json:"entityId"`
							EntityTitle       string `json:"entityTitle"`
						}{
							{
								Access:             "SITE",
								AllPurposeItemText: nil,
								Attachments:        nil,
								Author:             "12345678-9abc-defg-hijk-lmnopqrstuvw",
								AuthorLastModified: nil,
								CloseTime: struct {
									EpochSecond int `json:"epochSecond"`
									Nano        int `json:"nano"`
								}{
									EpochSecond: 0,
									Nano:        0,
								},
								CloseTimeString: time.Time{},
								Content:         nil,
								Context:         "12345",
								Creator:         nil,
								DropDeadTime: struct {
									EpochSecond int `json:"epochSecond"`
									Nano        int `json:"nano"`
								}{},
								DropDeadTimeString: time.Time{},
								DueTime: struct {
									EpochSecond int `json:"epochSecond"`
									Nano        int `json:"nano"`
								}{
									EpochSecond: 0,
									Nano:        0,
								},
								DueTimeString:       time.Time{},
								GradeScale:          "SCORE_GRADE_TYPE",
								GradeScaleMaxPoints: "10.00",
								GradebookItemID:     nil,
								GradebookItemName:   nil,
								Groups:              nil,
								ID:                  "12345678-9abc-defg-hijk-lmnopqrstuvw",
								Instructions:        "This is Test Assignment",
								ModelAnswerText:     nil,
								OpenTime: struct {
									EpochSecond int `json:"epochSecond"`
									Nano        int `json:"nano"`
								}{
									EpochSecond: 0,
									Nano:        0,
								},
								OpenTimeString:  time.Time{},
								Position:        0,
								PrivateNoteText: nil,
								Section:         "",
								Status:          "OPEN",
								SubmissionType:  "TEXT_AND_ATTACHMENT_ASSIGNMENT_SUBMISSION",
								TimeCreated: struct {
									EpochSecond int `json:"epochSecond"`
									Nano        int `json:"nano"`
								}{
									EpochSecond: 0,
									Nano:        0,
								},
								TimeLastModified: struct {
									EpochSecond int `json:"epochSecond"`
									Nano        int `json:"nano"`
								}{
									EpochSecond: 0,
									Nano:        0,
								},
								Title:             "Test Assignment",
								AllowResubmission: true,
								Draft:             false,
								EntityReference:   "\\/assignment\\/12345678-9abc-defg-hijk-lmnopqrstuvw",
								EntityURL:         "https:\\/\\/hoppii.hosei.ac.jp:443\\/direct\\/assignment\\/12345678-9abc-defg-hijk-lmnopqrstuvw",
								EntityID:          "12345678-9abc-defg-hijk-lmnopqrstuvw",
								EntityTitle:       "Test Assingment",
							},
						},
					},
					nil,
				)
				return &testContext{
					fields: Fields{
						networkingService: networkingService,
					},
					args: Args{
						input: input,
					},
					returns: Returns{
						output: &AssignmentMyOutput{
							AssignmentMyResult: &assignmentMyResult{
								EntityPrefix: "assignment",
								AssignmentCollection: []struct {
									Access             string        `json:"access"`
									AllPurposeItemText interface{}   `json:"allPurposeItemText"`
									Attachments        []interface{} `json:"attachments"`
									Author             string        `json:"author"`
									AuthorLastModified interface{}   `json:"authorLastModified"`
									CloseTime          struct {
										EpochSecond int `json:"epochSecond"`
										Nano        int `json:"nano"`
									} `json:"closeTime"`
									CloseTimeString time.Time   `json:"closeTimeString"`
									Content         interface{} `json:"content"`
									Context         string      `json:"context"`
									Creator         interface{} `json:"creator"`
									DropDeadTime    struct {
										EpochSecond int `json:"epochSecond"`
										Nano        int `json:"nano"`
									} `json:"dropDeadTime"`
									DropDeadTimeString time.Time `json:"dropDeadTimeString"`
									DueTime            struct {
										EpochSecond int `json:"epochSecond"`
										Nano        int `json:"nano"`
									} `json:"dueTime"`
									DueTimeString       time.Time     `json:"dueTimeString"`
									GradeScale          string        `json:"gradeScale"`
									GradeScaleMaxPoints string        `json:"gradeScaleMaxPoints"`
									GradebookItemID     interface{}   `json:"gradebookItemId"`
									GradebookItemName   interface{}   `json:"gradebookItemName"`
									Groups              []interface{} `json:"groups"`
									ID                  string        `json:"id"`
									Instructions        string        `json:"instructions"`
									ModelAnswerText     interface{}   `json:"modelAnswerText"`
									OpenTime            struct {
										EpochSecond int `json:"epochSecond"`
										Nano        int `json:"nano"`
									} `json:"openTime"`
									OpenTimeString  time.Time   `json:"openTimeString"`
									Position        int         `json:"position"`
									PrivateNoteText interface{} `json:"privateNoteText"`
									Section         string      `json:"section"`
									Status          string      `json:"status"`
									SubmissionType  string      `json:"submissionType"`
									TimeCreated     struct {
										EpochSecond int `json:"epochSecond"`
										Nano        int `json:"nano"`
									} `json:"timeCreated"`
									TimeLastModified struct {
										EpochSecond int `json:"epochSecond"`
										Nano        int `json:"nano"`
									} `json:"timeLastModified"`
									Title             string `json:"title"`
									AllowResubmission bool   `json:"allowResubmission"`
									Draft             bool   `json:"draft"`
									EntityReference   string `json:"entityReference"`
									EntityURL         string `json:"entityURL"`
									EntityID          string `json:"entityId"`
									EntityTitle       string `json:"entityTitle"`
								}{
									{
										Access:             "SITE",
										AllPurposeItemText: nil,
										Attachments:        nil,
										Author:             "12345678-9abc-defg-hijk-lmnopqrstuvw",
										AuthorLastModified: nil,
										CloseTime: struct {
											EpochSecond int `json:"epochSecond"`
											Nano        int `json:"nano"`
										}{
											EpochSecond: 0,
											Nano:        0,
										},
										CloseTimeString: time.Time{},
										Content:         nil,
										Context:         "12345",
										Creator:         nil,
										DropDeadTime: struct {
											EpochSecond int `json:"epochSecond"`
											Nano        int `json:"nano"`
										}{},
										DropDeadTimeString: time.Time{},
										DueTime: struct {
											EpochSecond int `json:"epochSecond"`
											Nano        int `json:"nano"`
										}{
											EpochSecond: 0,
											Nano:        0,
										},
										DueTimeString:       time.Time{},
										GradeScale:          "SCORE_GRADE_TYPE",
										GradeScaleMaxPoints: "10.00",
										GradebookItemID:     nil,
										GradebookItemName:   nil,
										Groups:              nil,
										ID:                  "12345678-9abc-defg-hijk-lmnopqrstuvw",
										Instructions:        "This is Test Assignment",
										ModelAnswerText:     nil,
										OpenTime: struct {
											EpochSecond int `json:"epochSecond"`
											Nano        int `json:"nano"`
										}{
											EpochSecond: 0,
											Nano:        0,
										},
										OpenTimeString:  time.Time{},
										Position:        0,
										PrivateNoteText: nil,
										Section:         "",
										Status:          "OPEN",
										SubmissionType:  "TEXT_AND_ATTACHMENT_ASSIGNMENT_SUBMISSION",
										TimeCreated: struct {
											EpochSecond int `json:"epochSecond"`
											Nano        int `json:"nano"`
										}{
											EpochSecond: 0,
											Nano:        0,
										},
										TimeLastModified: struct {
											EpochSecond int `json:"epochSecond"`
											Nano        int `json:"nano"`
										}{
											EpochSecond: 0,
											Nano:        0,
										},
										Title:             "Test Assignment",
										AllowResubmission: true,
										Draft:             false,
										EntityReference:   "\\/assignment\\/12345678-9abc-defg-hijk-lmnopqrstuvw",
										EntityURL:         "https:\\/\\/hoppii.hosei.ac.jp:443\\/direct\\/assignment\\/12345678-9abc-defg-hijk-lmnopqrstuvw",
										EntityID:          "12345678-9abc-defg-hijk-lmnopqrstuvw",
										EntityTitle:       "Test Assingment",
									},
								},
							},
						},
					},
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			ctrl, ctx := gomock.WithContext(ctx, t)
			defer ctrl.Finish()

			go func() {
				tc := tt.testContext(ctrl)
				s := &AssignmentService{
					networkingService: tc.fields.networkingService,
				}
				got, err := s.My(tc.args.input)
				if err != nil {
					t.Error("AssignmentService.My() error = ", err)
				}
				if !reflect.DeepEqual(got, tc.returns.output) {
					t.Errorf("AssignmentService.My() = %v, want %v", got, tc.returns.output)
				}
				cancel()
			}()
			<-ctx.Done()
		})
	}
}

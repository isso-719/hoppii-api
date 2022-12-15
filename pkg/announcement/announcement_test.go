package announcement

import (
	"context"
	"github.com/golang/mock/gomock"
	"hoppii-api/pkg/networking"
	"hoppii-api/pkg/user"
	"reflect"
	"testing"
)

func TestAnnouncementService_User(t *testing.T) {
	type Fields struct {
		networkingService networking.IFNetworkingService
	}
	type Args struct {
		input *AnnouncementUserInput
	}
	type Returns struct {
		output *AnnouncementUserOutput
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
				input := &AnnouncementUserInput{
					UserCookie: &user.UserCookie{},
				}

				networkingService := networking.NewMockIFNetworkingService(ctrl)
				networkingService.EXPECT().CreateClient(gomock.Any()).Return(nil, nil)
				networkingService.EXPECT().GetJsonToStruct(
					gomock.Any(),
					"https://hoppii.hosei.ac.jp/direct/announcement/user.json",
					&announcementUserResult{},
				).Return(
					nil,
					&announcementUserResult{
						EntityPrefix: "announcement",
						AnnouncementCollection: []struct {
							AnnouncementID       string        `json:"announcementId"`
							Attachments          []interface{} `json:"attachments"`
							Body                 string        `json:"body"`
							Channel              string        `json:"channel"`
							CreatedByDisplayName string        `json:"createdByDisplayName"`
							CreatedOn            int64         `json:"createdOn"`
							ID                   string        `json:"id"`
							SiteID               string        `json:"siteId"`
							SiteTitle            string        `json:"siteTitle"`
							Title                string        `json:"title"`
							EntityReference      string        `json:"entityReference"`
							EntityURL            string        `json:"entityURL"`
							EntityID             string        `json:"entityId"`
							EntityTitle          string        `json:"entityTitle"`
						}{
							{
								AnnouncementID:       "12345678-9abc-defg-hijk-lmnopqrstuvw",
								Attachments:          nil,
								Body:                 "This is Test",
								Channel:              "main",
								CreatedByDisplayName: "Test User",
								CreatedOn:            0,
								ID:                   "12345:main:12345678-9abc-defg-hijk-lmnopqrstuvw",
								SiteID:               "12345",
								SiteTitle:            "Home",
								Title:                "Test",
								EntityReference:      "/announcement\\/12345:main:12345678-9abc-defg-hijk-lmnopqrstuvw",
								EntityURL:            "ttps:\\/\\/hoppii.hosei.ac.jp:443\\/direct\\/announcement\\/12345:main:12345678-9abc-defg-hijk-lmnopqrstuvw",
								EntityID:             "12345:main:12345678-9abc-defg-hijk-lmnopqrstuvw",
								EntityTitle:          "Test",
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
						output: &AnnouncementUserOutput{
							AnnouncementUserResult: &announcementUserResult{
								EntityPrefix: "announcement",
								AnnouncementCollection: []struct {
									AnnouncementID       string        `json:"announcementId"`
									Attachments          []interface{} `json:"attachments"`
									Body                 string        `json:"body"`
									Channel              string        `json:"channel"`
									CreatedByDisplayName string        `json:"createdByDisplayName"`
									CreatedOn            int64         `json:"createdOn"`
									ID                   string        `json:"id"`
									SiteID               string        `json:"siteId"`
									SiteTitle            string        `json:"siteTitle"`
									Title                string        `json:"title"`
									EntityReference      string        `json:"entityReference"`
									EntityURL            string        `json:"entityURL"`
									EntityID             string        `json:"entityId"`
									EntityTitle          string        `json:"entityTitle"`
								}{
									{
										AnnouncementID:       "12345678-9abc-defg-hijk-lmnopqrstuvw",
										Attachments:          nil,
										Body:                 "This is Test",
										Channel:              "main",
										CreatedByDisplayName: "Test User",
										CreatedOn:            0,
										ID:                   "12345:main:12345678-9abc-defg-hijk-lmnopqrstuvw",
										SiteID:               "12345",
										SiteTitle:            "Home",
										Title:                "Test",
										EntityReference:      "/announcement\\/12345:main:12345678-9abc-defg-hijk-lmnopqrstuvw",
										EntityURL:            "ttps:\\/\\/hoppii.hosei.ac.jp:443\\/direct\\/announcement\\/12345:main:12345678-9abc-defg-hijk-lmnopqrstuvw",
										EntityID:             "12345:main:12345678-9abc-defg-hijk-lmnopqrstuvw",
										EntityTitle:          "Test",
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
				s := &AnnouncementService{
					networkingService: tc.fields.networkingService,
				}
				got, err := s.User(tc.args.input)
				if err != nil {
					t.Error("AnnouncementService.User() error = ", err)
				}
				if !reflect.DeepEqual(got, tc.returns.output) {
					t.Errorf("AnnouncementService.User() = %v, want %v", got, tc.returns.output)
				}
				cancel()
			}()
			<-ctx.Done()
		})
	}
}

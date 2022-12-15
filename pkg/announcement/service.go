package announcement

//go:generate mockgen -source=service.go -destination=service_mock.go -package=announcement

import (
	"hoppii-api/pkg/networking"
	"hoppii-api/pkg/user"
)

type IFAnnouncementService interface {
	User(input *AnnouncementUserInput) (*AnnouncementUserOutput, error)
}

type ifPrivateAnnouncementService interface {
}

type AnnouncementService struct {
	networkingService networking.IFNetworkingService

	privateService ifPrivateAnnouncementService
}

// TODO: オプションの考慮
type AnnouncementUserInput struct {
	UserCookie *user.UserCookie
}

type announcementUserResult struct {
	EntityPrefix           string `json:"entityPrefix"`
	AnnouncementCollection []struct {
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
	} `json:"announcement_collection"`
}

type AnnouncementUserOutput struct {
	AnnouncementUserResult *announcementUserResult
}

func NewAnnouncementService() *AnnouncementService {
	srv := &AnnouncementService{
		networkingService: networking.NewNetworkingService(),
	}
	srv.privateService = srv

	return srv
}

package hoppii_api

import (
	"github.com/isso-719/hoppii-api/pkg/announcement"
	"github.com/isso-719/hoppii-api/pkg/assignment"
	"github.com/isso-719/hoppii-api/pkg/networking"
	"github.com/isso-719/hoppii-api/pkg/user"
)

type HoppiiApi struct {
	Announcement announcement.IFAnnouncementService
	Assignment   assignment.IFAssignmentService
	Networking   networking.IFNetworkingService
	User         user.IFUserService
}

func NewHoppiiApi() *HoppiiApi {
	return &HoppiiApi{
		Announcement: announcement.NewAnnouncementService(),
		Assignment:   assignment.NewAssignmentService(),
		Networking:   networking.NewNetworkingService(),
		User:         user.NewUserService(),
	}
}

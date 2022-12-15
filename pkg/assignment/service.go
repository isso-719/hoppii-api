package assignment

//go:generate mockgen -source=service.go -destination=service_mock.go -package=assignment

import (
	"hoppii-api/pkg/networking"
	"hoppii-api/pkg/user"
	"time"
)

type IFAssignmentService interface {
	My(input *AssignmentMyInput) (*AssignmentMyOutput, error)
}

type ifPrivateAssignmentService interface {
}

type AssignmentService struct {
	networkingService networking.IFNetworkingService

	privateService ifPrivateAssignmentService
}

type AssignmentMyInput struct {
	UserCookie *user.UserCookie
}

type assignmentMyResult struct {
	EntityPrefix         string `json:"entityPrefix"`
	AssignmentCollection []struct {
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
	} `json:"assignment_collection"`
}

type AssignmentMyOutput struct {
	AssignmentMyResult *assignmentMyResult
}

func NewAssignmentService() *AssignmentService {
	srv := &AssignmentService{
		networkingService: networking.NewNetworkingService(),
	}
	srv.privateService = srv

	return srv
}

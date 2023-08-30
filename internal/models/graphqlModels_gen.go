// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AbsoluteMediaHTTP struct {
	ID          string `json:"id"`
	URI         string `json:"uri"`
	URIAbsolute string `json:"uri_absolute"`
}

type ChatHTTP struct {
	ID    string    `json:"id"`
	User1 *UserHTTP `json:"user1"`
	User2 *UserHTTP `json:"user2"`
}

type ChatMutationResult struct {
	ID    string `json:"id"`
	User1 string `json:"user1"`
	User2 string `json:"user2"`
}

type CourseAPIMediaCollectionHTTP struct {
	ID          string             `json:"id"`
	BannerImage *AbsoluteMediaHTTP `json:"banner_image,omitempty"`
	CourseImage *MediaHTTP         `json:"course_image,omitempty"`
	CourseVideo *MediaHTTP         `json:"course_video,omitempty"`
	Image       *ImageHTTP         `json:"image,omitempty"`
}

type CourseHTTP struct {
	ID               string                        `json:"id"`
	BlocksURL        string                        `json:"blocks_url"`
	Effort           string                        `json:"effort"`
	EnrollmentStart  string                        `json:"enrollment_start"`
	EnrollmentEnd    string                        `json:"enrollment_end"`
	End              string                        `json:"end"`
	Name             string                        `json:"name"`
	Number           string                        `json:"number"`
	Org              string                        `json:"org"`
	ShortDescription string                        `json:"short_description"`
	Start            string                        `json:"start"`
	StartDisplay     string                        `json:"start_display"`
	StartType        string                        `json:"start_type"`
	Pacing           string                        `json:"pacing"`
	MobileAvailable  bool                          `json:"mobile_available"`
	Hidden           bool                          `json:"hidden"`
	InvitationOnly   bool                          `json:"invitation_only"`
	Overview         *string                       `json:"overview,omitempty"`
	CourseID         string                        `json:"course_id"`
	Media            *CourseAPIMediaCollectionHTTP `json:"media"`
}

type CoursesListHTTP struct {
	Courses   []*CourseHTTP `json:"courses"`
	CountRows int           `json:"countRows"`
}

type ImageHTTP struct {
	ID    string `json:"id"`
	Raw   string `json:"raw"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type MediaHTTP struct {
	ID  string `json:"id"`
	URI string `json:"uri"`
}

type MessageEdge struct {
	Node   *MessageHTTP `json:"node,omitempty"`
	Cursor string       `json:"cursor"`
}

type MessageHTTP struct {
	ID       string     `json:"id"`
	Payload  string     `json:"payload"`
	Sender   *UserHTTP  `json:"sender"`
	Receiver *UserHTTP  `json:"receiver"`
	ChatID   string     `json:"chatID"`
	Time     *time.Time `json:"time,omitempty"`
}

type MessagesFromUserInput struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
}

type NewChat struct {
	User1 string `json:"user1"`
	User2 string `json:"user2"`
}

type NewMessage struct {
	Payload  string `json:"payload"`
	Receiver string `json:"receiver"`
}

type NewUser struct {
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Role       Role    `json:"role"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Middlename *string `json:"middlename,omitempty"`
	Nickname   string  `json:"nickname"`
}

type NewUserResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Role       int    `json:"role"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
}

type PageInfo struct {
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type ProjectPageHTTP struct {
	ID               string `json:"id"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
	AuthorID         string `json:"authorId"`
	ProjectID        string `json:"projectId"`
	ProjectUpdatedAt string `json:"projectUpdatedAt"`
	Title            string `json:"title"`
	Instruction      string `json:"instruction"`
	Notes            string `json:"notes"`
	LinkToScratch    string `json:"linkToScratch"`
	IsShared         bool   `json:"isShared"`
	IsBanned         bool   `json:"isBanned"`
}

type ProjectPageHTTPList struct {
	ProjectPages []*ProjectPageHTTP `json:"projectPages"`
	CountRows    int                `json:"countRows"`
}

type Response struct {
	Ok bool `json:"ok"`
}

type Settings struct {
	ActivationByLink bool `json:"activationByLink"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignUp struct {
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Nickname   string  `json:"nickname"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Middlename *string `json:"middlename,omitempty"`
}

type UpdateProjectPage struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Instruction string `json:"instruction"`
	Notes       string `json:"notes"`
	IsShared    bool   `json:"isShared"`
}

type UpdateUser struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Middlename string `json:"middlename"`
	Nickname   string `json:"nickname"`
}

type UserHTTP struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           Role   `json:"role"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Middlename     string `json:"middlename"`
	Nickname       string `json:"nickname"`
	IsActive       bool   `json:"isActive"`
	ActivationLink string `json:"activationLink"`
}

type UsersList struct {
	Users     []*UserHTTP `json:"users"`
	CountRows int         `json:"countRows"`
}

type Role string

const (
	RoleAnonymous  Role = "Anonymous"
	RoleStudent    Role = "Student"
	RoleParent     Role = "Parent"
	RoleTeacher    Role = "Teacher"
	RoleUnitAdmin  Role = "UnitAdmin"
	RoleSuperAdmin Role = "SuperAdmin"
)

var AllRole = []Role{
	RoleAnonymous,
	RoleStudent,
	RoleParent,
	RoleTeacher,
	RoleUnitAdmin,
	RoleSuperAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAnonymous, RoleStudent, RoleParent, RoleTeacher, RoleUnitAdmin, RoleSuperAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

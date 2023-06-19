package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"fmt"

	"github.com/skinnykaen/rpa_clone/internal/models"
)

// GetCourseContent is the resolver for the GetCourseContent field.
func (r *queryResolver) GetCourseContent(ctx context.Context, courseID string) (*models.CourseHTTP, error) {
	panic(fmt.Errorf("not implemented: GetCourseContent - GetCourseContent"))
}

// GetCoursesByUser is the resolver for the GetCoursesByUser field.
func (r *queryResolver) GetCoursesByUser(ctx context.Context) (*models.CoursesListHTTP, error) {
	panic(fmt.Errorf("not implemented: GetCoursesByUser - GetCoursesByUser"))
}

// GetAllPublicCourses is the resolver for the GetAllPublicCourses field.
func (r *queryResolver) GetAllPublicCourses(ctx context.Context, pageNumber string) (*models.CoursesListHTTP, error) {
	panic(fmt.Errorf("not implemented: GetAllPublicCourses - GetAllPublicCourses"))
}

// GetEnrollments is the resolver for the GetEnrollments field.
func (r *queryResolver) GetEnrollments(ctx context.Context, username string) (*models.EnrollmentsListHTTP, error) {
	panic(fmt.Errorf("not implemented: GetEnrollments - GetEnrollments"))
}
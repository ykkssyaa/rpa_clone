package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"strconv"

	"github.com/skinnykaen/rpa_clone/internal/consts"
	"github.com/skinnykaen/rpa_clone/internal/models"
)

// CreateProjectPage is the resolver for the CreateProjectPage field.
func (r *mutationResolver) CreateProjectPage(ctx context.Context) (*models.ProjectPageHTTP, error) {
	newProjectPage, err := r.projectPageService.CreateProjectPage(ctx.Value(consts.KeyId).(uint))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	projectPageHttp := models.ProjectPageHTTP{}
	projectPageHttp.FromCore(newProjectPage)
	return &projectPageHttp, nil
}

// UpdateProjectPage is the resolver for the UpdateProjectPage field.
func (r *mutationResolver) UpdateProjectPage(ctx context.Context, input models.UpdateProjectPage) (*models.ProjectPageHTTP, error) {
	atoi, err := strconv.Atoi(input.ID)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	projectPage := models.ProjectPageCore{
		ID:          uint(atoi),
		Title:       input.Title,
		Instruction: input.Instruction,
		Notes:       input.Notes,
		IsShared:    input.IsShared,
	}
	updatedProjectPage, err := r.projectPageService.UpdateProjectPage(projectPage, ctx.Value(consts.KeyId).(uint))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	projectPageHttp := models.ProjectPageHTTP{}
	projectPageHttp.FromCore(updatedProjectPage)
	return &projectPageHttp, nil
}

// DeleteProjectPage is the resolver for the DeleteProjectPage field.
func (r *mutationResolver) DeleteProjectPage(ctx context.Context, id string) (*models.Response, error) {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	if r.projectPageService.DeleteProjectPage(uint(atoi), ctx.Value(consts.KeyId).(uint)); err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	return &models.Response{Ok: true}, nil
}

// GetProjectPageByID is the resolver for the GetProjectPageById field.
func (r *queryResolver) GetProjectPageByID(ctx context.Context, id string) (*models.ProjectPageHTTP, error) {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	project, err := r.projectPageService.GetProjectPageById(uint(atoi), ctx.Value(consts.KeyId).(uint), ctx.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	projectPageHttp := models.ProjectPageHTTP{}
	projectPageHttp.FromCore(project)
	return &projectPageHttp, nil
}

// GetAllProjectPagesByAuthorID is the resolver for the GetAllProjectPagesByAuthorId field.
func (r *queryResolver) GetAllProjectPagesByAuthorID(ctx context.Context, id string, page *int, pageSize *int) (*models.ProjectPageHTTPList, error) {
	atoi, err := strconv.Atoi(id)
	projects, countRows, err := r.projectPageService.GetProjectsPageByAuthorId(uint(atoi), page, pageSize)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	return &models.ProjectPageHTTPList{
		ProjectPages: models.FromProjectPagesCore(projects),
		CountRows:    int(countRows),
	}, nil
}

// GetAllProjectPagesByAccessToken is the resolver for the GetAllProjectPagesByAccessToken field.
func (r *queryResolver) GetAllProjectPagesByAccessToken(ctx context.Context, page *int, pageSize *int) (*models.ProjectPageHTTPList, error) {
	projects, countRows, err := r.projectPageService.GetAllProjectPages(page, pageSize, ctx.Value(consts.KeyId).(uint), ctx.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, err
	}
	return &models.ProjectPageHTTPList{
		ProjectPages: models.FromProjectPagesCore(projects),
		CountRows:    int(countRows),
	}, nil
}

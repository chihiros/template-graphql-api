package usecase

import (
	"context"
)

type UserUseCase interface {
	Get(context.Context) (Response, error)
	GetByID(context.Context, int) (Response, error)
	Post(context.Context, Request) (Response, error)
	Put(context.Context, Request) (Response, error)
	Delete(context.Context, int) error
}

type UserRepository interface {
	Get(context.Context) (Response, error)
	GetByID(context.Context, int) (Response, error)
	Post(context.Context, Request) (Response, error)
	Put(context.Context, Request) (Response, error)
	Delete(context.Context, int) error
}

type UserUsecase struct {
	Repository UserRepository
}

func (u *UserUsecase) Get(ctx context.Context) (Response, error) {
	return u.Repository.Get(ctx)
}

func (u *UserUsecase) GetByID(ctx context.Context, id int) (Response, error) {
	return u.Repository.GetByID(ctx, id)
}

func (u *UserUsecase) Post(ctx context.Context, req Request) (Response, error) {
	return u.Repository.Post(ctx, req)
}

func (u *UserUsecase) Put(ctx context.Context, req Request) (Response, error) {
	return u.Repository.Put(ctx, req)
}

func (u *UserUsecase) Delete(ctx context.Context, id int) error {
	return u.Repository.Delete(ctx, id)
}

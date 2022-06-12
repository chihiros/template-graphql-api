package repository

import (
	"context"
	"fmt"
	"tamaribacms/ent"
	"tamaribacms/ent/user"
	"tamaribacms/usecase"
	"time"

	"entgo.io/ent/dialect/sql"
)

type UserRepository struct {
	DBConn *ent.Client
}

func NewUserRepository(conn *ent.Client) *UserRepository {
	return &UserRepository{
		DBConn: conn,
	}
}

func (r *UserRepository) Get(ctx context.Context) (usecase.Response, error) {
	users, err := r.DBConn.User.Query().All(ctx)
	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: users}
	return res, err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (usecase.Response, error) {
	user, err := r.DBConn.User.Query().
		Where(user.IDEQ(id)).
		All(ctx)

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: user}
	return res, err
}

func (r *UserRepository) Post(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	user, err := r.DBConn.User.Create().
		SetUsername(req.Username).
		SetAge(req.Age).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			// ent側の制約エラー
			return usecase.Response{}, fmt.Errorf("duplicate")
		}
	}

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: user}
	return res, err
}

func (r *UserRepository) Put(ctx context.Context, req usecase.Request) (usecase.Response, error) {
	id, err := r.DBConn.User.Create().
		SetUsername(req.Username).
		SetAge(req.Age).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		OnConflict(
			sql.ConflictColumns(user.FieldUsername),
		).
		Update(func(u *ent.UserUpsert) {
			u.SetUsername(req.Username)
			u.SetAge(req.Age)
			u.UpdateUpdatedAt()
		}).
		ID(ctx)

	if err != nil {
		panic(err)
	}

	// 更新されたユーザー情報を取得する
	user, err := r.DBConn.User.Query().
		Where(user.IDEQ(id)).
		All(ctx)

	if err != nil {
		panic(err)
	}

	res := usecase.Response{Data: user}
	return res, err
}
func (r *UserRepository) Delete(ctx context.Context, id int) error {
	_, err := r.DBConn.User.Delete().
		Where(user.IDEQ(id)).
		Exec(ctx)

	if err != nil {
		panic(err)
	}

	return err
}

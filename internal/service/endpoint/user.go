package endpoint

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"github.com/txvier/user-srv/internal/dbconn"
	"github.com/txvier/user-srv/internal/entity"

	user "github.com/txvier/user-srv/api/proto/user"
)

type UserHandler struct {
}

func (h *UserHandler) Hello(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Log("Received Foo.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (h *UserHandler) CreateUser(ctx context.Context, in *user.CreateUserRequest, out *user.CreateUserResponse) error {
	u := &entity.User{
		Name:  in.Name,
		Pwd:   in.Pwd,
		Tel:   in.Tel,
		Email: in.Email,
	}
	affected, err := dbconn.Get().Insert(u)
	if err != nil {
		return err
	}
	out.Affected = affected
	return nil
}

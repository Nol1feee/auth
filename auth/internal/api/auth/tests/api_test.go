package tests

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/api/auth"
	"github.com/Nol1feee/CLI-chat/auth/internal/model"
	"github.com/Nol1feee/CLI-chat/auth/internal/service"
	desc "github.com/Nol1feee/CLI-chat/auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	//t.Parallel()
	type authServiceMock func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		email = gofakeit.Email()
		name  = gofakeit.Name()
		role  = desc.Role(0)

		req = &desc.CreateRequest{
			UserInfo: &desc.UserInfo{
				Name:  name,
				Email: email,
				Role:  role,
			}, Password: "", PasswordConfirm: "",
		}

		info = &model.UserInfo{
			Name:  name,
			Email: email,
			Role:  model.RoleUser,
		}

		resp = &desc.CreateResponse{Id: id}
	)

	tests := []struct {
		name            string
		args            args
		expectedRes     *desc.CreateResponse
		expectedErr     error
		authServiceMock authServiceMock
	}{
		{name: "first test",
			args:        args{ctx: ctx, req: req},
			expectedRes: resp,
			expectedErr: nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
	}
	defer t.Cleanup(mc.Finish)
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := auth.NewImplementation(authServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.expectedErr, err)
			require.Equal(t, tt.expectedRes, newID)
		})
	}
}

//
//func TestTest(t *testing.T) {
//	//t.Parallel()
//	type authServiceMock func(mc *minimock.Controller) service.AuthService
//
//	type args struct {
//		ctx context.Context
//		req *desc.CreateRequest
//	}
//
//	var (
//		ctx = context.Background()
//		mc  = minimock.NewController(t)
//
//		id    = gofakeit.Int64()
//		email = gofakeit.Email()
//		name  = gofakeit.Name()
//		role  = desc.Role(0)
//
//		req = &desc.CreateRequest{
//			UserInfo: &desc.UserInfo{
//				Name:  name,
//				Email: email,
//				Role:  role,
//			}, Password: "", PasswordConfirm: "",
//		}
//
//		info = &model.UserInfo{
//			Name:  name,
//			Email: email,
//			Role:  model.RoleUser,
//		}
//
//		resp = &desc.CreateResponse{Id: id}
//	)
//
//	tests := []struct {
//		name            string
//		args            args
//		expectedRes     *desc.CreateResponse
//		expectedErr     error
//		authServiceMock authServiceMock
//	}{
//		{name: "first test",
//			args:        args{ctx: ctx, req: req},
//			expectedRes: resp,
//			expectedErr: nil,
//			authServiceMock: func(mc *minimock.Controller) service.AuthService {
//				mock := NewAuthServiceMock(mc)
//				mock.CreateMock.Expect(ctx, info).Return(id, nil)
//				return mock
//			},
//		},
//	}
//	defer t.Cleanup(mc.Finish)
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//
//			authServiceMock := tt.authServiceMock(mc)
//			api := auth.NewImplementation(authServiceMock)
//
//			newID, err := api.Create(tt.args.ctx, tt.args.req)
//			require.Equal(t, tt.expectedErr, err)
//			require.Equal(t, tt.expectedRes, newID)
//		})
//	}
//}

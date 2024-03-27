package auth

import (
	"context"
	"errors"
	ssov1 "github.com/17HIERARCH70/messageService/api-contracts/gen/go/sso"
	"github.com/17HIERARCH70/messageService/sso/internal/services/auth"
	"github.com/17HIERARCH70/messageService/sso/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)

	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)

	IsAdmin(
		ctx context.Context,
		userID int64,
	) (bool, error)

	GetIdByEmail(
		ctx context.Context,
		email string,
	) (int64, error)

	RestorePassword(
		ctx context.Context,
		email string,
	) error
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})
}

func ValidateLoginRequest(req *ssov1.LoginRequest) error {
	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	if req.AppId == 0 {
		return status.Error(codes.InvalidArgument, "app id is required")
	}

	return nil
}

func ValidateRegisterRequest(req *ssov1.RegisterRequest) error {
	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {

	if err := ValidateLoginRequest(in); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, in.GetEmail(), in.GetPassword(), int(in.GetAppId()))

	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}

		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &ssov1.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {

	if err := ValidateRegisterRequest(in); err != nil {
		return nil, err
	}

	uid, err := s.auth.RegisterNewUser(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &ssov1.RegisterResponse{UserId: uid}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	in *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	isAdmin, err := s.auth.IsAdmin(ctx, in.GetUserId())
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		return nil, status.Error(codes.Internal, "failed to check admin status")
	}

	return &ssov1.IsAdminResponse{IsAdmin: isAdmin}, nil
}

func (s *serverAPI) GetIdByEmail(
	ctx context.Context,
	in *ssov1.GetIdByEmailRequest,
) (*ssov1.GetIdByEmailResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	uid, err := s.auth.GetIdByEmail(ctx, in.GetEmail())
	if err != nil {

	}
	return &ssov1.GetIdByEmailResponse{UserId: uid}, nil
}

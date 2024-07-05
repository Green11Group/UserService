package serviceuser

import (
	pb "GreenThumb/User-Service/genproto"
	"GreenThumb/User-Service/storage/postgres"
	"context"
	"database/sql"
)

type UserServer struct {
	pb.UnimplementedUserServer
	db   *sql.DB
	user *postgres.UserRepo
}

func NewUserServer(db *sql.DB, userRepo *postgres.UserRepo) *UserServer {
	return &UserServer{db: db, user: userRepo}
}

func (s *UserServer) GetUSerByID(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.user.GetUserByID(req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
	}, nil
}

func (s *UserServer) UpdateUserByID(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	user, err := s.user.UpdateUserByID(req)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
	}, nil
}

func (s *UserServer) DeleteUserByID(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	user, err := s.user.DeleteUserByID(req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{
		Message: user.Message,
	}, nil
}

func (s *UserServer) GetUserProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	profile, err := s.user.GetUserProfileByID(req)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileResponse{
		FullName:  profile.FullName,
		Bio:       profile.Bio,
		Expertise: profile.Expertise,
		Location:  profile.Location,
		AvatarUrl: profile.AvatarUrl,
	}, nil
}

func (s *UserServer) UpdateUserProfile(ctx context.Context, req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error) {
	profile, err := s.user.UpdateUserProfileByID(req)
	if err != nil {
		return nil, err
	}
	return &pb.ProfileUpdateResponse{
		FullName:  profile.FullName,
		Bio:       profile.Bio,
		Expertise: profile.Expertise,
		Location:  profile.Location,
		AvatarUrl: profile.AvatarUrl,
	}, nil
}

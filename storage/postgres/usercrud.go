package postgres

import (
	"database/sql"
	"time"

	pb "GreenThumb/User-Service/genproto"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetUserByID(req *pb.UserRequest) (*pb.UserResponse, error) {
	query := `select username, email, password_hash from users where id = $1`

	row := u.db.QueryRow(query, req.UserID)
	var user pb.UserResponse
	err := row.Scan(&user.UserName, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) UpdateUserByID(req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	query := `update users set username = $2, email = $3 where id = $1`

	var user pb.UpdateResponse

	row := u.db.QueryRow(query, req.UserID, user.UserName, user.Email)
	err := row.Scan(&user.UserName, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) DeleteUserByID(req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	query := `update users set deleted_at = $2 from users where id = $1`

	now := time.Now()
	DeletedAt := now.Unix()

	_, err := u.db.Exec(query, req.UserID, DeletedAt)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Message: "Successfully deleted user by id"}, nil
}

func (u *UserRepo) GetUserProfileByID(req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	query := `select fullname, bio, expertise, location, avatar_url from users_profile where id = $1`
	row := u.db.QueryRow(query, req.UserID)
	var profile pb.ProfileResponse
	err := row.Scan(&profile.FullName, &profile.Bio, &profile.Expertise, &profile.Location, &profile.AvatarUrl)
	if err != nil {
		return nil, row.Err()
	}
	return &profile, nil
}

func (u *UserRepo) UpdateUserProfileByID(req *pb.ProfileUpdateRequest) (*pb.ProfileUpdateResponse, error) {
	query := `update users_profile set fullname = $2, bio =$3 , expertise =$4, location = $5, avatar_url = $6 where id = $1`

	var profile pb.ProfileUpdateResponse

	row := u.db.QueryRow(query, req.UserID, profile.FullName, profile.Bio, profile.Expertise, profile.Location, profile.AvatarUrl)

	err := row.Scan(&profile.FullName, &profile.Bio, &profile.Expertise, &profile.Location, &profile.AvatarUrl)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

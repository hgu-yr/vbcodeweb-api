package service

import (
	"errors"

	"vbCodeWeb/repo"
	"vbCodeWeb/request"
	"vbCodeWeb/utils"
)

// AuthService 登录鉴权业务逻辑。
type AuthService struct {
	userRepo repo.UserRepo
}

func NewAuthService(userRepo repo.UserRepo) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// UserInfo 登录返回的用户信息。
type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
}

// LoginResult 登录成功结果。
type LoginResult struct {
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}

// Login 校验账号密码并签发 token。
func (s *AuthService) Login(req *request.LoginReq) (*LoginResult, error) {
	u, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	if u.Status != 1 {
		return nil, errors.New("账号已被禁用")
	}
	if !utils.CheckPassword(u.Password, req.Password) {
		return nil, errors.New("用户名或密码错误")
	}
	token, err := utils.GenerateToken(u.ID, u.Username, u.Role)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		Token: token,
		User: UserInfo{
			ID:       u.ID,
			Username: u.Username,
			Nickname: u.Nickname,
			Role:     u.Role,
		},
	}, nil
}

package service

import (
	"auth/config"
	"auth/model"
	"auth/repository"
	"auth/schema"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User interface {
	GetByID(Id uint) (record model.User, err error)
	Create(data schema.RegisterRequest) (model.User, error)
	GetForLogin(request schema.LoginRequest) (string, error)
	ParseToken(token string) (*jwt.Token, error)
	ExtractUserIDFromJWT(tokenString string) (int, error)
}

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetByID(Id uint) (record model.User, err error) {
	return model.User{}, err
}
func (s *UserService) Create(data schema.RegisterRequest) (model.User, error) {
	var user model.User
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hash)
	user.Username = data.Username
	return s.userRepo.Create(user)
}
func (s *UserService) GetForLogin(request schema.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByUsername(request.Username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	cfg, _ := config.GetConfig()
	tokenString, err := token.SignedString([]byte(cfg.Security.SecretKey))
	fmt.Println("token " + tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *UserService) ParseToken(tokenString string) (*jwt.Token, error) {
	cfg, _ := config.GetConfig()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка алгоритма подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Неподдерживаемый метод подписи: %v", token.Header["alg"])
		}
		return []byte(cfg.Security.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Проверка валидности токена
	if !token.Valid {
		return nil, fmt.Errorf("Недействительный токен")
	}

	return token, nil
}

func (s *UserService) ExtractUserIDFromJWT(tokenString string) (int, error) {
	token, err := s.ParseToken(tokenString)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("Ошибка извлечения утверждений из токена")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, fmt.Errorf("Ошибка извлечения ID пользователя из токена")
	}

	userID := int(sub)
	return userID, nil
}

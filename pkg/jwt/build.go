package jwt

import (
	"time"

	jwt_pack "github.com/dgrijalva/jwt-go"
	"github.com/yueqing2617/XFLY/service/conf"
)

type Jwt struct {
	Config *conf.JwtConfig
}

func NewJwt() *Jwt {
	return &Jwt{Config: conf.Config.JwtConfig}
}

type CustomClaims struct {
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
	jwt_pack.StandardClaims
}

// MakeToken 生成token
func (j *Jwt) MakeToken(userId, roleId int64) (string, error) {
	token := jwt_pack.NewWithClaims(jwt_pack.SigningMethodHS256, &CustomClaims{
		UserId: userId,
		RoleId: roleId,
		StandardClaims: jwt_pack.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.Config.ExpiresAt)).Unix(),
			Audience:  j.Config.Audience,
			Issuer:    j.Config.Issuer,
			Subject:   j.Config.Subject,
		},
	})
	return token.SignedString([]byte(j.Config.PrivateKey))
}

// ParseToken 解析token
func (j *Jwt) ParseToken(token string) (*CustomClaims, error) {
	t, err := jwt_pack.ParseWithClaims(token, &CustomClaims{}, func(token *jwt_pack.Token) (interface{}, error) {
		return []byte(j.Config.PrivateKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}

package web

import (
	"tiktok/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(uid int64, username string) {
	var (
		exTime = 3 * 24 * time.Hour //三天会过期
		sign   = []byte("tiktok")
	)

	claim := models.Jwt{
		Uid:      uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exTime).Unix(), //过期时间
			Issuer:    "tiktok",                      //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim) //获取token
	token.SignedString(sign)                                  //sign为签名，此操作可获得完整的签名令牌(sign token)
}

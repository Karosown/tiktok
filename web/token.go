package web

import (
	"errors"
	"tiktok/models"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	exTime = time.Minute //3 * 24 * time.Hour //三天会过期
	sign   = []byte("tiktok")
)

/*传入uid和username
 *1、成功返回token字符串
 *2、失败返回错误信息
 **/
func CreateToken(uid int64, username string) (string, error) { //不同客户端登录识别问题
	//无则创建，有则返回
	claim := models.MyClaims{
		Uid:      uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exTime).Unix(), //过期时间
			Issuer:    "tiktok",                      //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) //获取token        //密钥类型问题
	return token.SignedString(sign)                           //sign为签名，此操作可获得完整的签名令牌(sign token)
}

/*传入token字符串
 *1、如果未过期返回models.MyClaims指针，内部包含
 *type MyClaims struct {
 *	Uid                int64
 *	Username           string
 *	jwt.StandardClaims
 * }
 *2、如果过期返回空指针和错误信息err
 **/

func ParseToken(tokenstr string) (*models.MyClaims, error) {
	claim := models.MyClaims{}
	token, err := jwt.ParseWithClaims(tokenstr, claim, func(token *jwt.Token) (i interface{}, err error) {
		return sign, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return &claim, nil
	}
	return nil, errors.New("invalid token")
}

package dms

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type GoogleOAuthService struct{}

func (s *GoogleOAuthService) HandleGoogleOAuth(code string) (token string, err error) {
	// Use the code to get the id and access tokens
	tokenRes, err := utils.GetGoogleOauthToken(code)

	if err != nil {
		return "", err
	}

	user, err := utils.GetGoogleUser(tokenRes.AccessToken, tokenRes.IdToken)
	if err != nil {
		return "", err
	}

	var dbUser system.SysUser
	err = global.GVA_DB.Model(&system.SysUser{}).Where("username = ?", user.Email).First(&dbUser).Error
	if err != nil {
		return "", err
	}

	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        dbUser.UUID,
		ID:          dbUser.ID,
		NickName:    dbUser.NickName,
		Username:    dbUser.Username,
		AuthorityId: dbUser.AuthorityId,
	})

	token, err = j.CreateToken(claims)

	if err != nil {
		return "", err
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		return token, nil
	}

	return "", errors.New("cannot generate new token")
}

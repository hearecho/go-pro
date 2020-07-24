package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-pro/go-web/pkg/resp"
	"github.com/hearecho/go-pro/go-web/pkg/utils"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		r := resp.R{}.Ok().SetPath(context.Request.URL.Path)

		token := context.Query("token")
		if token == "" {
			r = r.SetCode(resp.INVALID_PARAMS).SetMsg(resp.MsgFlags[resp.INVALID_PARAMS])
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				r = r.SetCode(resp.ERROR_AUTH_CHECK_TOKEN_FAIL).SetMsg(resp.MsgFlags[resp.ERROR_AUTH_CHECK_TOKEN_FAIL])
			} else if time.Now().Unix() > claims.ExpiresAt {
				r = r.SetCode(resp.ERROR_AUTH_CHECK_TOKEN_TIMEOUT).SetMsg(resp.MsgFlags[resp.ERROR_AUTH_CHECK_TOKEN_TIMEOUT])
			}
		}
		if r.Code != resp.SUCCESS {
			context.JSON(http.StatusUnauthorized,r)
			context.Abort()
			return
		}
		context.Next()
	}
}
package middleware


import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/appleboy/gin-jwt.v2"

	"github.com/xintangli/monitor/models"
	"github.com/xintangli/monitor/utils"
)

var AuthMiddleware *jwt.GinJWTMiddleware
var GlobalMiddleware gin.HandlerFunc
var AdminHandlers = []string{"main.userCreate", "main.userDelete"}

func InitMiddleware() {
	//globalMiddleware = GlobalMiddleware()
	AuthMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "TalkingData",
		Key:        []byte("testKey"),
		Timeout:    time.Hour * 24 * 7,
		MaxRefresh: time.Hour * 24 * 7,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			user := &models.User{
				Username:userId,
			}
			//查找用户
			if err := user.ReadByUserName(); err != nil {
				c.Set("code", http.StatusNotFound)
				c.Set("message", "user not found")
				return userId, false
			}

			if user.Status != 1 {
				c.Set("code", http.StatusUnauthorized)
				c.Set("message", "user is disable")
				return userId, false
			}
			//密码校验
			if userId == user.Username && utils.MD5(strings.TrimSpace(password)) == user.Password {
				return userId, true
			}
			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			user := &models.User{
				Username:userId,
			}

			if err := user.ReadByUserName(); err != nil {
				c.Set("code", http.StatusNotFound)
				c.Set("message", "user not found")
				return false
			}

			if user.Role == models.USER {
				for _, handler := range AdminHandlers {
					if c.HandlerName() == handler {
						return false
					}
				}
			}
			return true
		},
		PayloadFunc: func(userId string) map[string]interface{} {
			user_info := make(map[string]interface{})
			user_info["userId"] = userId
			return user_info
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			if new_code, ok := c.Get("code"); ok {
				code = new_code.(int)
				new_message, _ := c.Get("message")
				message = new_message.(string)
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
				return
			}

			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	}
}

/*
func GlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := &models.User{
			Username: c.Get("userId"),
		}
		err := user.ReadByUserName()

		//for user not found
		if err != nil || user.ID <= 0 {
			c.Abort()
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "user not found",
			})
			return
		}

		c.Set("user", user)

		//for user status
		if user.Status != 1 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "user is disable",
			})
			return
		}

		//for opertion journal
		operations := models.Operations{}
		operations.OperationResult = true
		operations.OperationTime = time.Now()
		operations.IP = strings.Split(c.Request.RemoteAddr, ":")[0]
		rip := c.Request.Header.Get("X-Real-Ip")
		if len(rip) > 0 {
			operations.IP = rip
		}
		operations.Operator = user.Username
		operations.Content = fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.String())

		c.Next()

		if strings.Contains(c.Request.URL.String(), "/operations") != true {
			if c.Writer.Status() <= 599 && c.Writer.Status() >= 400 {
				operations.OperationResult = false
			}
			//mydb.Create(&operations)
		}
	}
}*/

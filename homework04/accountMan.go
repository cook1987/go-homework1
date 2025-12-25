package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// 自定义 Claims
type MyClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const JWT_SECRET string = "mygintest"

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db := GetDBFromContext(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}
	existUser := &User{}
	if err := db.Where("username = ?", user.Username).First(&existUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username has been used! please transfer another one!"})
		return
	}
	if err := db.Where("email = ?", user.Email).First(&existUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email has been used! please transfer another one!"})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	logger.WithField("username", user.Username).Info("创建用户")
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var userl UserLogin
	if err := c.ShouldBindJSON(&userl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := GetDBFromContext(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}
	var storedUser User
	if err := db.Where("username = ?", userl.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userl.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token, err := GenerateToken(storedUser.ID, storedUser.Username, []byte(JWT_SECRET))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	logger.WithField("username", storedUser.Username).Info("用户登录")
	// 返回 token 给前端
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GenerateToken(userId uint, username string, secretKey []byte) (string, error) {
	claims := MyClaims{
		ID:       userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
			Issuer:    "cook1987",                                         // 签发者
			Subject:   "ginTest",                                          // 主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从 header 获取 token
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "缺少认证令牌，请先登录"})
			ctx.Abort() // 终止后续处理
			return
		}
		// 验证 token
		claim, err := ParseToken(token, []byte(JWT_SECRET))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort() // 终止后续处理
			return
		}
		// 将 claims 存入 context，供后续 handler 使用
		ctx.Set("claim", claim)
		// token 验证通过，继续处理
		ctx.Next()
	}
}

func ParseToken(tokenString string, secretKey []byte) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (any, error) {
		// 验证签名方法
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

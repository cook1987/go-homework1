package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:870101@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	return db
}

// DBMiddleware 将 DB 实例注入到请求上下文中
func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 将 DB 实例存储到 Context 中
		ctx.Set("db", db)
		ctx.Next()
	}
}

// 从 Context 中获取 DB 实例
func GetDBFromContext(c *gin.Context) *gorm.DB {
	if db, exists := c.Get("db"); exists {
		if gormDB, ok := db.(*gorm.DB); ok {
			return gormDB
		}
	}
	return nil
}

func main() {
	// 初始化日志
	initLogger()
	db := InitDB()
	// 自动迁移模型
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	router := gin.Default()
	// 使用日志和恢复中间件
	router.Use(GinLogrusLogger(), GinLogrusRecovery())
	router.Use(DBMiddleware(db))

	blogGrp := router.Group("/blog")
	{
		blogGrp.POST("/register", Register)                      // 用户注册接口
		blogGrp.POST("/login", Login)                            // 登录接口
		blogGrp.GET("/getAllPosts", GetAllPosts)                 // 获取所有文章列表接口
		blogGrp.GET("/getPostDetail/:id", GetPostDetail)         // 获取单个文章的详细信息接口
		blogGrp.GET("/getCommentsOfPost/:id", GetCommentsOfPost) // 评论读取接口，支持获取某篇文章的所有评论列表。
		// 需要验证 token 的接口分组
		bizGrp := blogGrp.Group("/biz")
		bizGrp.Use(AuthMiddleware())

		bizGrp.POST("/createPost", CreatePost)       // 文章创建接口
		bizGrp.PATCH("/updatePost", UpdatePost)      // 文章更新接口，只有文章的作者才能更新自己的文章。
		bizGrp.DELETE("/deletePost/:id", DeletePost) // 文章删除接口，只有文章的作者才能删除自己的文章。
		bizGrp.POST("/createComment", CreateComment) // 评论创建接口，已认证的用户可以对文章发表评论。
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
	// 启动服务器
	logger.Info("服务器启动在 :8080")
}

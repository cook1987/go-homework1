package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	cl, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	claim := cl.(*MyClaims)
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		logger.WithError(err).Warn("用户数据绑定失败")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := GetDBFromContext(c)
	if db == nil {
		logger.Error("Database connection not avaiable")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}
	postN := Post{Title: post.Title, Content: post.Content, UserID: claim.ID}
	if err := db.Create(&postN).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Post"})
		return
	}
	logger.WithField("title", post.Title).WithField("Content", post.Content).WithField("UserID", claim.ID).Info("创建文章")
	c.JSON(http.StatusOK, gin.H{"msg": "CreatePost ok"})
}

func GetAllPosts(c *gin.Context) {
	db := GetDBFromContext(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}
	var posts []Post
	db.Select("*").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func GetPostDetail(c *gin.Context) {
	db := GetDBFromContext(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please assign a post id!"})
		return
	}

	pid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please assign a correct post id!"})
		return
	}
	var post Post
	db.Where("id = ?", pid).Find(&post)
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data not exists!"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	cl, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	claim := cl.(*MyClaims)
	var postu PostUpdate
	if err := c.ShouldBindJSON(&postu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := GetDBFromContext(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}
	var post Post
	db.Where("id = ?", postu.ID).Find(&post)
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data not exists!"})
		return
	}
	if post.UserID != claim.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can't update post of others!"})
		return
	}
	db.Model(&post).Updates(map[string]interface{}{"title": postu.Title, "content": postu.Content})
	logger.WithField("title", postu.Title).WithField("Content", postu.Content).WithField("id", post.ID).WithField("UserID", claim.ID).Info("更新文章")
	c.JSON(http.StatusOK, gin.H{"msg": "UpatePost ok"})
}

func GetCommentsOfPost(c *gin.Context) {
	db := GetDBFromContext(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please assign a post id!"})
		return
	}

	pid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please assign a correct post id!"})
		return
	}
	var comments []Comment
	db.Where("post_id = ?", pid).Find(&comments)
	c.JSON(http.StatusOK, &comments)
}

func DeletePost(c *gin.Context) {
	db := GetDBFromContext(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please assign a post id!"})
		return
	}

	pid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please assign a correct post id!"})
		return
	}
	var post Post
	db.Where("id = ?", pid).Find(&post)
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data not exists!"})
		return
	}
	cl, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	claim := cl.(*MyClaims)
	if post.UserID != claim.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can't delete post of others!"})
		return
	}
	db.Delete(&post)
	logger.WithField("id", post.ID).WithField("title", post.Title).WithField("Content", post.Content).WithField("UserID", claim.ID).Info("删除文章")
	c.JSON(http.StatusOK, post)
}

func CreateComment(c *gin.Context) {
	cl, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	claim := cl.(*MyClaims)
	var cc CommentCreate
	if err := c.ShouldBindJSON(&cc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := GetDBFromContext(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not avaiable"})
		return
	}

	var post Post
	db.Where("id = ?", cc.PostID).Find(&post)
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not exists!"})
		return
	}
	var comment = Comment{Content: cc.Content, PostID: cc.PostID, UserID: claim.ID}
	db.Save(&comment)
	logger.WithField("Content", cc.Content).WithField("PostID", cc.PostID).WithField("UserID", claim.ID).Info("添加评论")
	c.JSON(http.StatusOK, gin.H{"msg": "Save ok"})
}

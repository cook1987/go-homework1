package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID      uint
	Name    string
	Posts   []Post
	PostNum int
}

type Post struct {
	ID           uint
	Title        string
	Content      string
	UserId       uint
	Comments     []Comment
	CommentNum   int
	CommentState string
}

type Comment struct {
	ID      uint
	Content string
	PostID  uint
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:870101@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	return db
}

// 题目 1
func initTable() {
	db := InitDB()
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	user1 := User{
		Name: "zhangsan",
		Posts: []Post{
			{Title: "one", Content: "zhangsan one", Comments: []Comment{
				{Content: "lisi one yes"},
				{Content: "wangwu one yes"},
				{Content: "zhaoliu one yes"},
			}},
			{Title: "two", Content: "zhangsan two", Comments: []Comment{
				{Content: "lisi two yes"},
				{Content: "zhaoliu two no"},
			}},
		},
	}
	db.Create(&user1)
	user2 := User{
		Name: "lisi",
		Posts: []Post{
			{Title: "three", Content: "lisi three", Comments: []Comment{
				{Content: "lisi three yes"},
			}},
			{Title: "four", Content: "lisi four", Comments: []Comment{
				{Content: "zhangsan four yes"},
				{Content: "zhaoliu four no"},
			}},
		},
	}
	db.Create(&user2)
	user3 := User{
		Name: "zhaoliu",
		Posts: []Post{
			{Title: "five", Content: "zhaoliu five", Comments: []Comment{
				{Content: "lisi five yes"},
			}},
			{Title: "six", Content: "zhaoliu six", Comments: []Comment{
				{Content: "zhangsan six yes"},
			}},
		},
	}
	db.Create(&user3)

}

// 题目 2
func referenceSearch() {
	db := InitDB()
	// 使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	db.Preload("Posts.Comments").Preload(clause.Associations).Find(&user, "name=?", "zhangsan")
	fmt.Println(user)

	// 使用Gorm查询评论数量最多的文章信息。
	var post Post
	db.Raw("select * from posts p join (select post_id from comments group by post_id order by count(*) desc limit 1)c on c.post_id=p.id").Scan(&post)
	fmt.Println(post)
}

// 题目 3
//
//	为Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{ID: p.UserId}).Update("post_num", gorm.Expr("post_num + 1"))
	return
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	tx.Model(&Post{ID: c.PostID}).Update("comment_num", gorm.Expr("comment_num - 1"))
	var post Post
	post.ID = c.PostID
	tx.Find(&post)
	if post.CommentNum <= 0 {
		post.CommentNum = 0
		post.CommentState = "无评论"
	}
	tx.Save(&post)
	return
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&Post{ID: c.PostID}).Update("comment_num", gorm.Expr("comment_num + 1"))
	return
}

// Title        string
// 	Content      string
// 	UserId       uint

func testHook() {
	db := InitDB()
	// db.Create(&Post{Title: "test_add_post", Content: "my creator is zhangsan", UserId: 1})

	var comment Comment
	comment.ID = 2
	db.Find(&comment)
	db.Delete(&comment)
}

func main() {
	// initTable()
	// referenceSearch()
	testHook()
}

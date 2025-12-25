## 一、项目运行环境：<br/>
go版本：go1.25.4<br/>
mysql版本: 8.0<br/>
## 二、依赖安装步骤
gorm: go get -u gorm.io/gorm<br/>
mysql: go get -u gorm.io/driver/mysql<br/>
gin: go get -u github.com/gin-gonic/gin<br/>
jwt: go get github.com/golang-jwt/jwt/v5<br/>
日志：go get github.com/sirupsen/logrus<br/>
## 三、启动方式
homework04> go run .
## 四、接口列表
### 1. 用户注册接口：POST http://localhost:8080/blog/register
BODY:<br/>

{
    "username":"cook3",
    "password":"123456",
    "email":"cc2@qq.com"
}
<br/>
响应：<br/>
{
	"message": "User registered successfully"
}
### 2. 登录接口：POST http://localhost:8080/blog/login
BODY:<br/>
{
    "username":"cook3",
    "password":"123456"
}
<br/>
响应：<br/>
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJjb29rMyIsImlzcyI6ImNvb2sxOTg3Iiwic3ViIjoiZ2luVGVzdCIsImV4cCI6MTc2NjczMDcyMiwibmJmIjoxNzY2NjQ0MzIyLCJpYXQiOjE3NjY2NDQzMjJ9.VkzGML7lRuBMT9oHoHFxqkJuslgSheRuYklozlCpBlU"
}
### 3. 文章创建接口：POST http://localhost:8080/blog/biz/createPost
Header:
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJjb29rMyIsImlzcyI6ImNvb2sxOTg3Iiwic3ViIjoiZ2luVGVzdCIsImV4cCI6MTc2NjczMDcyMiwibmJmIjoxNzY2NjQ0MzIyLCJpYXQiOjE3NjY2NDQzMjJ9.VkzGML7lRuBMT9oHoHFxqkJuslgSheRuYklozlCpBlU

BODY:
{
    "title": "99999广东人，把雷克萨斯买成了年度最硬气日系车2",
    "content": "99999虽然2025年还没结束，但雷克萨斯实现年销量同比正增长，应该是没有太大的问题。放眼这个进口豪华车市场，这大概率是独一家。数据显示，雷克萨斯今年1—11月的终端销量约16.7万辆，而2024年全年相同口径下的销量则约17.8万辆。年底冲刺背景下，这个日系豪华品牌预计今年的销量将介于18万—19万辆之间。"
}

响应：
{
	"msg": "CreatePost ok"
}
### 4. 获取所有文章列表接口：GET http://localhost:8080/blog/getAllPosts

响应：
[
	{
		"ID": 2,
		"CreatedAt": "2025-12-25T10:17:03.363+08:00",
		"UpdatedAt": "2025-12-25T10:17:03.363+08:00",
		"DeletedAt": null,
		"Title": "强生被判赔偿一女子约110亿元，创爽身粉致癌案赔付新高",
		"Content": "据央视财经报道，当地时间12月22日，美国马里兰州陪审团作出一项裁定：强生公司需向一名因使用其婴儿爽身粉罹患癌症的女性支付15.6亿美元（约合人民币110亿元）赔偿，这一金额创下强生滑石粉致癌诉讼15年来单一原告获赔的最高纪录。"
	},
	{
		"ID": 4,
		"CreatedAt": "2025-12-25T11:00:31.502+08:00",
		"UpdatedAt": "2025-12-25T11:00:31.502+08:00",
		"DeletedAt": null,
		"Title": "广东人，把雷克萨斯买成了年度最硬气日系车",
		"Content": "虽然2025年还没结束，但雷克萨斯实现年销量同比正增长，应该是没有太大的问题。放眼这个进口豪华车市场，这大概率是独一家。数据显示，雷克萨斯今年1—11月的终端销量约16.7万辆，而2024年全年相同口径下的销量则约17.8万辆。年底冲刺背景下，这个日系豪华品牌预计今年的销量将介于18万—19万辆之间。"
	}
]
### 5. 获取单个文章的详细信息接口：GET http://localhost:8080/blog/getPostDetail/2

响应：
{
	"ID": 2,
	"CreatedAt": "2025-12-25T10:17:03.363+08:00",
	"UpdatedAt": "2025-12-25T10:17:03.363+08:00",
	"DeletedAt": null,
	"Title": "强生被判赔偿一女子约110亿元，创爽身粉致癌案赔付新高",
	"Content": "据央视财经报道，当地时间12月22日，美国马里兰州陪审团作出一项裁定：强生公司需向一名因使用其婴儿爽身粉罹患癌症的女性支付15.6亿美元（约合人民币110亿元）赔偿，这一金额创下强生滑石粉致癌诉讼15年来单一原告获赔的最高纪录。"
}

### 6. 文章更新接口：PATCH http://localhost:8080/blog/biz/updatePost

Header:
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJjb29rMyIsImlzcyI6ImNvb2sxOTg3Iiwic3ViIjoiZ2luVGVzdCIsImV4cCI6MTc2NjczMDcyMiwibmJmIjoxNzY2NjQ0MzIyLCJpYXQiOjE3NjY2NDQzMjJ9.VkzGML7lRuBMT9oHoHFxqkJuslgSheRuYklozlCpBlU

BODY:
{
    "id": 2,
    "title": "索尼独供时代终结! upate 333333",
    "content": "快科技12月25日消息，据媒体报道，三星将替代索尼成为iPhone 18"
}

响应：
{
	"msg": "UpatePost ok"
}

### 7. 文章删除接口：DELETE http://localhost:8080/blog/biz/deletePost/6

Header:
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJjb29rMyIsImlzcyI6ImNvb2sxOTg3Iiwic3ViIjoiZ2luVGVzdCIsImV4cCI6MTc2NjczMDcyMiwibmJmIjoxNzY2NjQ0MzIyLCJpYXQiOjE3NjY2NDQzMjJ9.VkzGML7lRuBMT9oHoHFxqkJuslgSheRuYklozlCpBlU

响应：
{
	"ID": 6,
	"CreatedAt": "2025-12-25T14:32:36.479+08:00",
	"UpdatedAt": "2025-12-25T14:32:36.479+08:00",
	"DeletedAt": "2025-12-25T14:33:07.434+08:00",
	"Title": "99999广东人，把雷克萨斯买成了年度最硬气日系车2",
	"Content": "99999虽然2025年还没结束，但雷克萨斯实现年销量同比正增长，应该是没有太大的问题。放眼这个进口豪华车市场，这大概率是独一家。数据显示，雷克萨斯今年1—11月的终端销量约16.7万辆，而2024年全年相同口径下的销量则约17.8万辆。年底冲刺背景下，这个日系豪华品牌预计今年的销量将介于18万—19万辆之间。"
}

### 8. 评论创建接口：POST http://localhost:8080/blog/biz/createComment

Header:
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJjb29rMyIsImlzcyI6ImNvb2sxOTg3Iiwic3ViIjoiZ2luVGVzdCIsImV4cCI6MTc2NjczMDcyMiwibmJmIjoxNzY2NjQ0MzIyLCJpYXQiOjE3NjY2NDQzMjJ9.VkzGML7lRuBMT9oHoHFxqkJuslgSheRuYklozlCpBlU

BODY:
{
    "content": "广东人，把雷克萨斯买成了年度最硬气日系车",
    "postID": 2
}

响应：
{
	"msg": "Save ok"
}

### 9. 评论读取接口：GET http://localhost:8080/blog/getCommentsOfPost/2

响应：
[
	{
		"ID": 1,
		"CreatedAt": "2025-12-25T13:28:48.306+08:00",
		"UpdatedAt": "2025-12-25T13:28:48.306+08:00",
		"DeletedAt": null,
		"Content": "广东人，把雷克萨斯买成了年度最硬气日系车",
		"UserID": 1,
		"PostID": 2
	}
]

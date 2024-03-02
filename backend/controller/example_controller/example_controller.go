package example_controller

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	// "fmt"
	"github.com/gin-gonic/gin"
	"os"
	// "template/constants"
	// daoImpl "template/dao/impl"
	// "template/db/mysql"
	// "template/model"
	// "template/model/params"
	// serviceImpl "template/service/example_service"
)

type ExampleController struct {
}

// 注册请求参数结构体
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 注册响应数据结构体
type SignUpResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		UserID string `json:"userId"`
	} `json:"data"`
}

// SignUpHandler 处理注册请求
func (o ExampleController) SignUpHandler(c *gin.Context) {
	// 1. 解析请求参数
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": err.Error(), "data": gin.H{}})
		return
	}

	// 2. 校验参数
	if req.Username == "" || req.Password == "" {
		c.JSON(200, gin.H{"code": 400, "msg": "用户名和密码不能为空", "data": gin.H{}})
		return
	}

	// 3. 处理业务逻辑，模拟注册成功
	resp := SignUpResponse{
		Code: 1001,
		Msg:  "注册成功",
	}
	resp.Data.UserID = "123456"

	// 4. 返回结果
	c.JSON(200, gin.H{"code": resp.Code, "msg": resp.Msg, "data": resp.Data})
}

// 登录请求参数结构体
type LoginRequest struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

// 登录响应数据结构体
type LoginResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data struct{} `json:"data"`
}

// loginHandler 处理登录请求
func (o ExampleController) LoginHandler(c *gin.Context) {
	// 1. 解析请求参数
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, LoginResponse{Code: 1004, Msg: "登录失败", Data: struct{}{}})
		return
	}

	// 2. 校验参数
	if req.UserID == "" || req.Password == "" {
		c.JSON(200, LoginResponse{Code: 1004, Msg: "登录失败", Data: struct{}{}})
		return
	}

	// 3. 处理业务逻辑，模拟登录成功
	resp := LoginResponse{
		Code: 1003,
		Msg:  "登录成功",
		Data: struct{}{},
	}

	// 4. 返回结果
	c.JSON(200, resp)
}

type TitleContentListResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (o ExampleController) GetTitleListHandler(c *gin.Context) {
	//读取./articles文件下的所有文件名
	files, err := os.ReadDir("./articles")
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "读取文件失败", "data": gin.H{}})
		return
	}
	//读取每个文件,文件内容的第一行为title，其他的为content
	var titleContentList []TitleContentListResponse
	for _, file := range files {
		f, err := os.Open("./articles/" + file.Name())
		if err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "打开文件失败", "data": gin.H{}})
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		// 读取第一行作为title
		scanner.Scan()
		title := scanner.Text()

		fmt.Println("title:", title)
		// 读取剩余的所有内容作为content
		content := ""
		for scanner.Scan() {
			content += scanner.Text() + "\n"
		}

		// 检查扫描过程中是否有错误发生
		if err := scanner.Err(); err != nil {
			c.JSON(200, gin.H{"code": 400, "msg": "读取文件内容失败", "data": gin.H{}})
			return
		}

		titleContentList = append(titleContentList, TitleContentListResponse{
			Title:   title,
			Content: strings.TrimSpace(content), // 移除尾部多余的换行符
		})
	}
	c.JSON(200, gin.H{"code": 200, "msg": "获取文章列表成功", "data": titleContentList})
}

const (
	baoyan = iota //保研
	jiuye         //就业
	chuguo        //出国
	kaoyan        //考研
)

type Step struct {
	Time        string `xml:"time"`
	Description string `xml:"description"`
}

type Route struct {
	Name  string `xml:"name,attr"`
	Steps []Step `xml:"step"`
}

type Routes struct {
	XMLName xml.Name `xml:"routes"`
	Routes  []Route  `xml:"route"`
}

func (o ExampleController) GetRoadsHandler(c *gin.Context) {
	// 示例XML数据，实际应用中您可能需要从文件中读取
	//从./roads/route/xml文件中读取
	fileName := fmt.Sprintf("./roads/route%d.xml", 0)
	xmlFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	// 解析XML
	var routes Routes
	if err := xml.Unmarshal([]byte(xmlData), &routes); err != nil {
		fmt.Println("Error parsing XML:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// 取出路线id
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// 使用示例map
	routeTypes := map[int]string{
		baoyan: "保研路线",
		jiuye:  "就业路线",
		chuguo: "出国路线",
		kaoyan: "考研路线",
	}

	target := routeTypes[id] //找到目标路线

	//随机出1或者2，但是要确保逻辑与ID或路线名称匹配
	randomInt := rand.Intn(2) + 1 //生成1或2
	key := fmt.Sprintf("%s%d", target, randomInt)
	fmt.Println("key:", key)

	// 找到对应的路线
	for _, route := range routes.Routes {
		if route.Name == key {
			// 如果找到，返回路线信息
			c.JSON(http.StatusOK, route)
			return
		}
	}

	// 如果没有找到匹配的路线
	c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
}

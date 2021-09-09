package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"sync"
	"time"
)

var (
	wg = sync.WaitGroup{}
)

const (
	timeFormat = "2006-01-02"
)

type MyTime time.Time

func (myTime *MyTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	t := time.Time(*myTime)
	b = t.AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (*MyTime) getTime() {
	fmt.Println("my time...")
}

func (myTime *MyTime) UnmarshalJSON(data []byte) error {
	fmt.Println("日期格式:", string(data))
	if location, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local); err != nil {
		return err
	} else {
		*myTime = MyTime(location)
		return nil
	}
}

type User struct {
	Username string                 `json:"username,omitempty" form:"username,default=defaultName"`
	Password string                 `json:"password,omitempty" form:"password"`
	Score    float32                `json:"score,omitempty" form:"score"`
	Birthday time.Time              `json:"birthday,omitempty" form:"birthday" time_format:"2006-01-02" binding:"required"`
	StuArr   []int                  `json:"stuArr,omitempty" form:"stuArr"`
	StuMap   map[string]interface{} `json:"stuMap,omitempty" form:"stuMap"`
	Addr     Addr                   `json:"addr" form:"addr"`
}

type Addr struct {
	Addr string `json:"addr,omitempty"`
	Desc string `json:"desc,omitempty"`
}

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		fmt.Println("处理请求前...")
		if match, err := path.Match("/gin/test/**", c.FullPath()); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		} else {
			if match {
				c.Set("key1", "value1")
				fmt.Println("路径匹配")
			} else {
				fmt.Println("路径不匹配")
			}
		}
		c.Next()
		fmt.Println("处理请求后...")
	})
	router.GET("/gin/test/hello", func(c *gin.Context) {
		log.Println("处理请求...")
		v, exists := c.Get("key1")
		if exists {
			fmt.Println("key1", v)
		} else {
			fmt.Println("key1不存在")
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
		})
	})
	router.POST("gin/test/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"message": "文件上传失败",
				"cause":   err.Error(),
			})
			return
		}
		c.SaveUploadedFile(file, "temp/"+file.Filename)
	})

	router.POST("/gin/test/form", func(c *gin.Context) {
		user := &User{}
		if err := c.ShouldBind(user); err != nil {
			fmt.Println(err)
		}
		fmt.Println(user)
		c.String(http.StatusOK, "ok")
	})
	router.GET("/gin/test/user/:id", func(c *gin.Context) {
		//username := c.Query("username")
		//password, b := c.GetQuery("password")
		id := c.Param("id")
		fmt.Println(id)
	})
	router.GET("/gin/user/:id/*action", func(c *gin.Context) {
		param := c.Param("action")
		fmt.Println(param)
		c.JSON(http.StatusOK, gin.H{
			"param": param,
		})
	})
	router.POST("/gin/test/json", func(c *gin.Context) {
		stu := &User{}
		err := c.BindJSON(stu)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "参数格式异常",
				"cause":   err.Error(),
			})
			return
		}
		fmt.Println(stu)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
			"data":    stu,
		})
	})
	router.GET("/gin/test/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	router.POST("/gin/test/time", func(c *gin.Context) {
		user := &User{}
		if err := c.ShouldBindJSON(user); err != nil {
			log.Println(err)
		}
		//log.Println(time.Time(*user.Birthday))
		c.String(http.StatusOK, "ok")
	})
	router.Run(":8080")
}

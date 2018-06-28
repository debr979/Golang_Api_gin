package ucontroller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	m "mysqlProc"
	"strings"
)

type usrMsg struct {
	UID      string `form:"uid" json:"uid"`
	PWD      string `form:"pwd" json:"pwd"`
	UNAME    string `json:"uname"`
	AGE      int    `json:"age"`
	EMAIL    string `json:"email"`
	TOCHANGE string `json:"tochange"`
}

var msg usrMsg

func cJson(c *gin.Context, code int, r string) {
	//簡化c.JSON寫法
	c.JSON(code, gin.H{"Status": r})

}
func JsonTurn(c *gin.Context, i interface{}) {
	//解析JSON格式，傳入C與介面格式
	buf := make([]byte, 1024)
	j, _ := c.Request.Body.Read(buf)
	err := json.Unmarshal(buf[0:j], i)
	if err != nil {
		c.JSON(200, gin.H{"Status": "Fail"})
		return
	}
}

func Register(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	conType := c.GetHeader("content-type")

	if strings.Contains(conType, "x-www") {
		msg.UID = c.PostForm("UID")
		msg.PWD = c.PostForm("PWD")
	} else if strings.Contains(conType, "json") {
		JsonTurn(c, &msg)
	}
	if msg.UID != "" && msg.PWD != "" {
		pwd := map[string]string{"PWD": msg.PWD}
		check := m.UserAction(msg.UID, pwd, 0)

		if strings.Contains(check, "PRIMARY") {
			cJson(c, 200, "帳號重複")
		} else {
			cJson(c, 200, check)
		}
	} else {
		cJson(c, 200, "帳號或密碼不可空白")
	}
}

func Login(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	conType := c.GetHeader("content-type")
	if strings.Contains(conType, "x-www") {
		msg.UID = c.PostForm("UID")
		msg.PWD = c.PostForm("PWD")
	} else if strings.Contains(conType, "json") {
		JsonTurn(c, &msg)
	}
	if msg.UID != "" && msg.PWD != "" {
		pwd := map[string]string{"PWD": msg.PWD}
		check := m.UserAction(msg.UID, pwd, 1)
		if check != "" {
			cJson(c, 200, "登入成功")
		} else {
			cJson(c, 200, "帳號或密碼錯誤")
		}
	} else {
		cJson(c, 200, "帳號或密碼不可空白")

	}
}

func Delete(c *gin.Context) {
	p1 := c.Param("action")
	p2 := map[string]string{"UID": c.Param("param")}
	delResult := m.UserAction(p1, p2, 2)
	cJson(c, 200, delResult)
}

func Update(c *gin.Context) {
	c.Header("content-type", "application/json")
	contype := c.GetHeader("content-type")
	if strings.Contains(contype, "x-www") {
		msg.UID = c.PostForm("UID")
		msg.PWD = c.PostForm("PWD")
		msg.TOCHANGE = c.PostForm("TOCHANGE")
	} else if strings.Contains(contype, "json") {
		JsonTurn(c, &msg)
	}

	passwd := map[string]string{
		"PWD":      msg.PWD,
		"TOCHANGE": msg.TOCHANGE,
	}
	udResult := m.UserAction(msg.UID, passwd, 3)
	cJson(c, 200, udResult)
}
func Get(c *gin.Context) {
	c.Header("content-type", "application/json")
	data := map[string]interface{}{
		"account": c.Param("account"),
		"actID":   0,
	}
	res := m.DbControl(data)
	cJson(c, 200, res)

}

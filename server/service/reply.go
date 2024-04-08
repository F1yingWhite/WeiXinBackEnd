package service

import (
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
	"unicode"
	"weixin_backend/models"

	"github.com/gin-gonic/gin"
)

type ReceiveMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MsgId        string   `xml:"MsgId"`
	TextContent  string   `xml:"Content"`
	PicUrl       string   `xml:"PicUrl"`
	MediaId      string   `xml:"MediaId"`
}

type ReplyTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

func (r *ReplyTextMsg) send() string {
	XmlForm := `
				<xml>
					<ToUserName><![CDATA[%s]]></ToUserName>
					<FromUserName><![CDATA[%s]]></FromUserName>
					<CreateTime>%d</CreateTime>
					<MsgType><![CDATA[%s]]></MsgType>
					<Content><![CDATA[%s]]></Content>
				</xml>
				`
	var jsonRaw interface{}
	err := json.Unmarshal([]byte(r.Content), &jsonRaw)
	if err != nil {
		return fmt.Sprintf(XmlForm, r.ToUserName, r.FromUserName, r.CreateTime, r.MsgType, r.Content)
	}

	jsonFormatted, err := json.MarshalIndent(jsonRaw, "", "    ")
	if err != nil {
		return fmt.Sprintf(XmlForm, r.ToUserName, r.FromUserName, r.CreateTime, r.MsgType, r.Content)
	}

	indentedJSON := string(jsonFormatted)
	return fmt.Sprintf(XmlForm, r.ToUserName, r.FromUserName, r.CreateTime, r.MsgType, indentedJSON)
}

type ReplyImageMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MediaId      string
}

func (r *ReplyImageMsg) send() string {
	XmlForm := `
            <xml>
                <ToUserName><![CDATA[%s]]></ToUserName>
                <FromUserName><![CDATA[%s]]></FromUserName>
                <CreateTime>%d</CreateTime>
                <MsgType><![CDATA[%s]]></MsgType>
                <Image>
                    <MediaId><![CDATA[%s]]></MediaId>
                </Image>
            </xml>
            `
	return fmt.Sprintf(XmlForm, r.ToUserName, r.FromUserName, r.CreateTime, r.MsgType, r.MediaId)
}

func parseXml(data []byte) (*ReceiveMsg, error) {
	var msg ReceiveMsg
	err := xml.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// 对内的接口，前端传来的消息需要满足 {"method":"get","url":"/search?company=华为&city=北京","json_data":"..."}
type InnerApi struct {
	Method   string                 `json:"method"`
	Url      string                 `json:"url"`
	JsonData map[string]interface{} `json:"json_data"`
}

func cleanInput(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}
		return r
	}, input)
}

func matchInnerApi(jsonStr string) (InnerApi, error) {
	jsonStr = cleanInput(jsonStr)
	var api InnerApi
	err := json.Unmarshal([]byte(jsonStr), &api)
	if err != nil {
		return InnerApi{}, err
	}
	return api, nil
}

type ReplyQuery struct {
	Signature    string `form:"signature" binding:"required"`
	Timestamp    string `form:"timestamp" binding:"required"`
	Nonce        string `form:"nonce" binding:"required"`
	Openid       string `form:"openid" binding:"required"`
	EncryptType  string `form:"encrypt_type" binding:"required"`
	MsgSignature string `form:"msg_signature" binding:"required"`
}

func checkSignature(signature string, timestamp string, nonce string) bool {
	token := "xyh"
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)
	tmpStr := ""
	for _, v := range tmpArr {
		tmpStr += v
	}
	hash := sha1.New()
	hash.Write([]byte(tmpStr))
	tmpStr = fmt.Sprintf("%x", hash.Sum(nil))

	if tmpStr == signature {
		return true
	} else {
		return false
	}
}

func Reply(c *gin.Context) {
	replyQuery := ReplyQuery{}
	c.ShouldBindQuery(&replyQuery)
	if !checkSignature(replyQuery.Signature, replyQuery.Timestamp, replyQuery.Nonce) {
		c.String(403, "Invalid signature")
		return
	}
	webData, _ := io.ReadAll(c.Request.Body)
	recMsg, err := parseXml(webData)
	if err != nil {
		c.String(500, "Failed to parse XML")
		return
	}
	log.Printf("Receive: %s\n", recMsg)
	//检查是否是新用户
	_, err = models.GetUserById(recMsg.FromUserName)
	if err != nil {
		models.CreateUser(recMsg.FromUserName)
	}
	switch recMsg.MsgType {
	case "text": //如果接受的是文本
		replyMsg := ReplyTextMsg{
			ToUserName:   recMsg.FromUserName,
			FromUserName: recMsg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      "text",
			Content:      "你好",
		}
		// log.Printf("Reply: %s\n", replyMsg.send())
		innerapi, err := matchInnerApi(recMsg.TextContent)
		if err != nil {
			log.Printf("Failed to match inner api: %s\n", err.Error())
			replyMsg.Content = `你好鸭，这是由超帅的许一涵开发的后端，如果希望获取更多信息，请输入正确的格式，例如：
{"method":"GET","url":"/user/get"}
{
    "method":"POST",
    "url":"/user/update",
    "json_data":{
        "username":"许一涵",
        "signature":"我很帅"
    }
}


{"method":"GET","url":"/salary?page_size=5&page=1&company=狐&city=北京"}
{"method":"GET","url":"/salary/getById?id=1"}
{"method":"GET","url":"/salary/getByUserId?page_size=5&page=1&user_id=obGiG6n3SPlTapeLcCVx2VAg1la8"}
{
    "method":"POST",
    "url":"/salary/create",
    "json_data":{
        "Company":"华为",
        "City":"北京",
        "Salary":"20w",
        "Major":"软件工程",
        "CategoryFirst":"技术/开发",
        "CategorySecond":"前端开发"
    }
}

{
    "method":"DELETE",
    "url":"/salary?id=?"
}

{
    "method":"POST",
    "url":"/salary/creates",
    "json_data":{
        "salaries":[
            {
                "Company":"字节跳动",
                "City":"北京",
                "Salary":"230w",
                "Major":"软件工程",
                "CategoryFirst":"技术/开发",
                "CategorySecond":"前端开发"
            },
            {
                "Company":"字节跳",
                "City":"北京",
                "Salary":"300w",
                "Major":"软件工程",
                "CategoryFirst":"技术/开发",
                "CategorySecond":"前端开发"
            }
        ]
        
    }
}`
			c.String(200, replyMsg.send())
			return
		} else {
			client := &http.Client{}
			jsonDataBytes, err := json.Marshal(innerapi.JsonData)
			if err != nil {
				replyMsg.Content = "转换JSON数据失败: " + err.Error()
				c.String(500, replyMsg.send())
				return
			}
			jsonDataStr := string(jsonDataBytes)
			req, err := http.NewRequest(innerapi.Method, "http://0.0.0.0:8888/api"+innerapi.Url, strings.NewReader(jsonDataStr))
			if err != nil {
				replyMsg.Content = "创建请求失败: " + err.Error()
				c.String(500, replyMsg.send())
				return
			}
			req.Header.Add("Authorization", recMsg.FromUserName)
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				replyMsg.Content = "请求失败: " + err.Error()
				c.String(500, replyMsg.send())
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				replyMsg.Content = "读取响应失败: " + err.Error()
				c.String(500, replyMsg.send())
				return
			}

			replyMsg.Content = string(body)
			c.String(200, replyMsg.send())
		}
	case "image": //如果接受的是图片
		replyMsg := ReplyImageMsg{
			ToUserName:   recMsg.FromUserName,
			FromUserName: recMsg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      "image",
			MediaId:      recMsg.MediaId,
		}
		c.String(200, replyMsg.send())
	default:
		c.String(200, "success")
	}
}

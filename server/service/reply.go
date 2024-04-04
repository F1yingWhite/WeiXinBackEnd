package service

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"time"

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
	return fmt.Sprintf(XmlForm, r.ToUserName, r.FromUserName, r.CreateTime, r.MsgType, r.Content)
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

func Reply(c *gin.Context) {
	webData, _ := io.ReadAll(c.Request.Body)
	recMsg, err := parseXml(webData)
	if err != nil {
		c.String(500, "Failed to parse XML")
		return
	}

	switch recMsg.MsgType {
	case "text": //如果接受的是文本
		replyMsg := ReplyTextMsg{
			ToUserName:   recMsg.FromUserName,
			FromUserName: recMsg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      "text",
			Content:      "test",
		}
		log.Printf("Reply: %s\n", replyMsg.send())
		c.String(200, replyMsg.send())
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

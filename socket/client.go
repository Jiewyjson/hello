package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	clientIp   string
	clientPort int
	clientName string
	conn       net.Conn
	flag       int
	C          chan string
}

func NewClient(clientIp string, clientPort int) *Client {
	client := &Client{
		clientIp:   clientIp,
		clientPort: clientPort,
		flag:       999,
		C:          make(chan string),
	}
	//连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", client.clientIp, client.clientPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}
	client.conn = conn
	return client
}

func init() {
	//接收命令行参数
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址,默认为127.0.1")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口,默认为8888")
}

// DealResponse 处理server回应的消息， 直接显示到标准输出即可
func (client *Client) DealResponse() {
	//一旦client.conn有数据，就直接copy到stdout标准输出上, 永久阻塞监听
	io.Copy(os.Stdout, client.conn)
}

func (client *Client) menu() bool {
	var flag int
	
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	
	fmt.Scanln(&flag)
	
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>请输入合法范围内的数字<<<<")
		return false
	}
	
}

// SelectUsers 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}

// PrivateChat 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	
	client.SelectUsers()
	fmt.Println(">>>>请输入聊天对象[用户名], exit退出:")
	fmt.Scanln(&remoteName)
	
	for remoteName != "exit" {
		fmt.Println(">>>>请输入消息内容, exit退出:")
		fmt.Scanln(&chatMsg)
		
		for chatMsg != "exit" {
			//消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}
			
			chatMsg = ""
			fmt.Println(">>>>请输入消息内容, exit退出:")
			fmt.Scanln(&chatMsg)
		}
		
		client.SelectUsers()
		fmt.Println(">>>>请输入聊天对象[用户名], exit退出:")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) PublicChat() {
	//提示用户输入消息
	var chatMsg string
	
	fmt.Println(">>>>请输入聊天内容，exit退出.")
	fmt.Scanln(&chatMsg)
	
	for chatMsg != "exit" {
		//发给服务器
		
		//消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}
		
		chatMsg = ""
		fmt.Println(">>>>请输入聊天内容，exit退出.")
		fmt.Scanln(&chatMsg)
	}
	
}

func (client *Client) UpdateName() bool {
	
	fmt.Println(">>>>请输入用户名:")
	fmt.Scanln(&client.clientName)
	
	sendMsg := "rename|" + client.clientName + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		
		//根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			//公聊模式
			client.PublicChat()
			break
		case 2:
			//私聊模式
			client.PrivateChat()
			break
		case 3:
			//更新用户名
			client.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

func main() {
	//命令行参数解析
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>> connect server failed, exit!")
		return
	}
	go client.DealResponse()
	fmt.Println(">>>>> connect server success!")
	//go client.menu()
	//启动客户端业务
	go client.Run()
	//阻塞
	select {}
}

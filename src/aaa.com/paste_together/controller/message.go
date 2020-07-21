package controller

import (
	"aaa.com/paste_together/common"
	"encoding/json"
	"errors"
	"html"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// 消息控制器
type MessageController struct {
	locker          sync.RWMutex
	TemplateDirPath string
	StorageDirPath  string
	storageFilePath string
	backupFilePath  string
}

// 消息分隔符
var msgSplitSymbol = "\n-------" + string(0x01) + "\n"

// 初始化
func (c *MessageController) Init() {
	c.storageFilePath = filepath.Join(c.StorageDirPath, "message.data")
	c.backupFilePath = filepath.Join(c.StorageDirPath, "message.data.bak")
}

// 获取所有消息
func (c *MessageController) GetContents(w http.ResponseWriter, r *http.Request) {
	// 读取文件内容
	content, err := c.readContents()
	if err != nil {
		common.ResponseJsonError(w, err)
		return
	}

	// 内容处理
	if len(content) == 0 {
		content = "No Data"
	} else {
		content = html.EscapeString(content)
		re, _ := regexp.Compile(msgSplitSymbol)
		content = re.ReplaceAllString(content, "<hr>")
		re, _ = regexp.Compile("\n")
		content = re.ReplaceAllString(content, "<br>")
	}

	// 渲染页面
	t, _ := template.ParseFiles(filepath.Join(c.TemplateDirPath, "index.html"))
	data := map[string]interface{}{
		"content": template.HTML(content),
	}

	t.Execute(w, data)
}

// 添加消息
func (c *MessageController) Create(w http.ResponseWriter, r *http.Request) {
	// 获取表单数据
	err := r.ParseForm()
	if err != nil {
		common.ResponseJsonError(w, err)
		return
	}
	formData := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		common.ResponseJsonError(w, err)
		return
	}

	// 获取消息
	msg := formData["message"].(string)
	msg = strings.Trim(msg, " ")
	if len(msg) == 0 {
		common.ResponseJsonError(w, errors.New("Message can not be empty"))
		return
	}

	// 添加消息到文件
	err = c.appendMsg(msg)
	if err != nil {
		common.ResponseJsonError(w, err)
		return
	}

	// 响应
	data := map[string]interface{}{
		"message": msg,
	}
	common.ResponseJson(w, http.StatusCreated, data)
}

// 删除所有消息
func (c *MessageController) DeleteAll(w http.ResponseWriter, r *http.Request) {
	// 备份当前消息文件内容，备份后清空掉
	c.backupFileAndClear()

	// 响应
	data := map[string]interface{}{}
	common.ResponseJson(w, http.StatusCreated, data)
}

// 读取文件内容
func (c *MessageController) readContents() (content string, err error) {
	c.locker.RLock()
	defer c.locker.RUnlock()

	// 读取文件内容
	common.CreateFileIfNotExist(c.storageFilePath)
	b, err := ioutil.ReadFile(c.storageFilePath)
	if err != nil {
		return
	}
	content = string(b)
	return
}

// 添加消息到文件
func (c *MessageController) appendMsg(msg string) (err error) {
	c.locker.Lock()
	defer c.locker.Unlock()

	// 消息处理
	msg += msgSplitSymbol

	// 添加内容到文件开头
	common.CreateFileIfNotExist(c.storageFilePath)
	err = common.AppendContentToFileStart(c.storageFilePath, msg)
	return
}

// 备份当前消息文件内容，备份后清空掉
func (c *MessageController) backupFileAndClear() {
	c.locker.Lock()
	defer c.locker.Unlock()

	// 创建备份文件
	common.CreateFileIfNotExist(c.backupFilePath)

	// 读取消息文件内容
	buf, err := ioutil.ReadFile(c.storageFilePath)
	if err != nil {
		return
	}

	// 读取备份文件内容
	buf2, err := ioutil.ReadFile(c.backupFilePath)
	if err != nil {
		return
	}

	// 写入备份文件
	err = common.WriteMultiBufToFile(c.backupFilePath, buf, buf2)
	if err != nil {
		return
	}

	// 将消息文件清空
	err = common.ClearFileContent(c.storageFilePath)
	if err != nil {
		return
	}
}

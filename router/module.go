package router

import (
	"net/http"
	"encoding/json"
	"html/template"
)

//内部模块组

//获取template路径
//param name string 路径末尾文件名称
//return string 路径
func getTemplateSrc(name string) string{
	return "." + pathSep + "template" + pathSep + name
}

//输出一个简单文本
//param w http.ResponseWriter 写入http句柄
//param r *http.Request 读取http句柄
//param contentByte []byte 输出文本
func postText(w http.ResponseWriter, r *http.Request, contentByte []byte) {
	_, err := w.Write(contentByte)
	if err != nil {
		sendLog("router/module.go",r.RemoteAddr,"postText","write-content",err.Error())
		return
	}
}

//跳转到URL
//param w http.ResponseWriter 写入http句柄
//param r *http.Request 读取http句柄
//param gotoURL string 跳转到URL
func goURL(w http.ResponseWriter, r *http.Request, gotoURL string) {
	http.Redirect(w, r, gotoURL, http.StatusFound)
}

//输出JSON
//param w http.ResponseWriter 写入http句柄
//param r *http.Request 读取http句柄
//param data interface{} 输出数据
//param b bool 获取数据是否成功
func postJSONData(w http.ResponseWriter, r *http.Request,data interface{},b bool) {
	res := make(map[string]interface{})
	res["status"] = b
	res["data"] = data
	res["login"] = false
	resJson,err := json.Marshal(res)
	if err != nil{
		sendLog("router/module.go",r.RemoteAddr,"postJSONData","get-json",err.Error())
		return
	}
	postText(w, r, resJson)
}

//检查post完整性
//param r *http.Request 读取http句柄
//return bool 是否存在post数据
func checkPost(r *http.Request) bool{
	var err error
	err = r.ParseForm()
	if err != nil {
		sendLog("router/module.go",r.RemoteAddr,"checkPost","check-form",err.Error())
	}
	return err == nil
}

//输出模版内容
//param w http.ResponseWriter 写入http句柄
//param r *http.Request 读取http句柄
//param templateFileName string 模版文件名称
//param data interface{} 输出参数
func showTemplate(w http.ResponseWriter, r *http.Request, templateFileName string, data interface{}){
	t, err := template.ParseFiles(getTemplateSrc(templateFileName),getTemplateSrc("page-header.html"),getTemplateSrc("page-menu.html"),getTemplateSrc("page-footer.html"),getTemplateSrc("page-menu-nologin.html"))
	if err != nil {
		sendLog("router/module.go",r.RemoteAddr,"showTemplate","show-template",err.Error())
		return
	}
	t.Execute(w, data)
}

//启动会话
func startSession(){
	if glob.SessionOperate.Status == false{
		glob.SessionOperate.Create(glob.AppName)
	}
}

//获取IP地址
//param r *http.Request 读取http句柄
//return string IP地址
func getIPAddrs(r *http.Request) string{
	return r.RemoteAddr
}
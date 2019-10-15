package handler

import (
	"drink_check/lib/common"
	"drink_check/lib/http_tools"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type DoUpload struct {
	DrinkCheckAPI string
	TokenAPI      string
}
type TokenResponse struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
}

func (d *DoUpload) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	log.Println("start do upload...")
	file_name := "upload_image"
	//接收文件
	file, header, err := req.FormFile(file_name)
	if err != nil {
		log.Printf("%v", err)
	}
	defer file.Close()
	log.Printf("upload file header %v", header.Header)
	log.Printf("the name of the file uploaded:%s,file size:%d", header.Filename, header.Size)

	//上传的文件不能超过3M
	if header.Size > 3000000 {
		log.Printf("the file is too big:%s", header.Filename)
		res.Write([]byte("{\"success\":\"1\",\"message\":\"上传的文件不能超过3M\"}"))
		return
	}
	//保存上传的图片
	f, err := os.OpenFile("./upload/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("%v", err)
	}

	_, err = io.Copy(f, file)
	if err != nil {
		log.Printf("create file failed:%v", err)
	}
	f.Close()

	accessToken, _ := d.getAccessToken()
	//去取文件内容并base64编码
	imageRaw, err := ioutil.ReadFile("./upload/" + header.Filename)
	if err != nil {
		log.Println(err)
	}
	raw := base64.StdEncoding.EncodeToString(imageRaw)

	requestData := fmt.Sprintf("{\"image\":\"%s\"}", raw)
	response := http_tools.HttpRequest(d.DrinkCheckAPI, "POST", "access_token="+accessToken, requestData, map[string]string{"Content-Type": "application/json"})

	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(response))
	log.Println("end do upload...")
}

func (d *DoUpload) getAccessToken() (string, error) {
	//优先从缓存文件中读取access token
	cacheFile := "./access_token"
	isExist, _ := common.FileExists(cacheFile)
	log.Printf("%v", isExist)
	if isExist {
		fileContent, err := ioutil.ReadFile(cacheFile)
		if err != nil {
			return "", err
		}
		//判断accesstoken是否有效
		tokenSections := strings.Split(string(fileContent), "|")
		createTime, _ := common.TimeToTimestamp("2006-01-02 15:04:06", tokenSections[1])
		//如果access token的创建时间大于两天 则重新获取
		if time.Now().Unix()-createTime < 3600*48 {
			return tokenSections[0], nil
		}
	}

	tokenParam := "grant_type=client_credentials&client_id=EdEuYT9AMRMFKVQMk4N6CVeR&client_secret=77kzmR2lXTWSFGTlLUGAKSU8KUgmytv0"
	res := http_tools.HttpRequest(d.TokenAPI, "GET", tokenParam, "", map[string]string{"Content-Type": "application/x-www-form-urlencoded"})

	tokenObj := TokenResponse{}
	err := json.Unmarshal([]byte(res), &tokenObj)
	if err != nil {
		log.Printf("json decode error:%v", err)
		return "", err
	}
	//将token写入文件缓存
	var f *os.File
	if isExist {
		//os.O_TRUNC 覆盖写入，不加则追加写入
		f, _ = os.OpenFile(cacheFile, os.O_TRUNC, 0666)
	} else {
		f, _ = os.Create(cacheFile)
	}
	defer f.Close()
	_, err = io.WriteString(f, tokenObj.AccessToken+"|"+time.Now().Format("2006-01-02 15:04:06"))
	if err != nil {
		log.Printf("cache token fail:%v", err)
	}

	return tokenObj.AccessToken, nil
}

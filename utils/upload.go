package utils

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"log"
	"meetingBooking/config"
	"mime/multipart"
	"os"
	"strconv"
)

// dirExistOrNot  判断文件是否存在
func dirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// createDir 创建文件夹
func createDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// UploadAvatarToLocalStatic 本地上传头像
func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + config.AvatarPath + "user" + bId + "/"
	if !dirExistOrNot(basePath) {
		createDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

// UploadImageToQiQiu 封装上传图片到七牛云然后返回状态和图片的url，单张
func UploadImageToQiQiu(file multipart.File, fileSize int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: config.Bucket,
	}
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}

	url := config.QiNiuServer + ret.Key
	return url, nil
}

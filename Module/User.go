package Module

import (
	"CeylonPaperOnline-BackEnd/Middleware"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func aes256Encrypt(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

func getMd5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func EncodingPassword(password, username string) string {
	// 加密算法为AES-256，密钥为用户名的MD5(全小写)
	// 最终结果为加密算法结果的MD5加上用户名的MD5
	h := md5.New()
	h.Write([]byte(username))
	key := h.Sum(nil)
	subStr := hex.EncodeToString(h.Sum(nil))

	result, err := aes256Encrypt(password, key)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		h.Write([]byte(result))
		return hex.EncodeToString(h.Sum(nil)) + subStr
	}
}

type User struct {
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	UserMajor string `json:"user_major"`
	UserGrade int `json:"user_grade"`
	UserClass int `json:"user_class"`
	UserStatus int `json:"user_status"`
	UserGroup string `json:"user_group"`
	UserSex int `json:"user_sex"`
	UserPass string `json:"user_pass"`
	UserEmail string `json:"user_email"`
	UserPhone string `json:"user_phone"`
}

func (u* User) SetPassword(password string) {
	u.UserPass = EncodingPassword(password, u.UserName)
}

func (u* User) Deserialization(jsonString string) {
	err := json.Unmarshal([]byte(jsonString), u)
	if err != nil {
		Middleware.Console.Log(err, "")
		return
	}
}
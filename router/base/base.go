package base

import (
	"github.com/speps/go-hashids"
	"golang.org/x/crypto/bcrypt"
)

type (

	ViewData struct {
		Data interface{}
		RouteName string
		Title string
		Description string
		VoiceNoteUrl string
	}

	JsonData struct {
		Err int `json:"err"`
		Data interface{} `json:"data"`
	}

	TableData struct {
		Data interface{} `json:"data"`
		Draw int `json:"draw"`
		RecordsTotal int `json:"recordsTotal"`
		RecordsFiltered int `json:"recordsFiltered"`
	}
)

const (
	HASHIDS_ENCRYPT_FILE_SALT = "voice note viettel file"
	JWT_HASH_KEY = "@#voice#$voice$%abc@#"
)

func EncodeHashId(id int) string {
	hd := hashids.NewData()
	hd.Salt = HASHIDS_ENCRYPT_FILE_SALT

 	h,_ := hashids.NewWithData(hd)
 	rs,_ := h.Encode([]int{id})

 	return rs
}

func EncodeBcrypt(password string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return ""
	} else {
		return string(h)
	}
}

package util

import (
	"encoding/base64"
	"goflow/model"
	"log"
	"sync"

	"github.com/sony/sonyflake"
	"gorm.io/gorm"
)

type Config struct {
	Port string
	Dsn  string
}

var Conf = Config{}
var Db *gorm.DB
var processMap sync.Map
var flake=sonyflake.NewSonyflake(sonyflake.Settings{})

func NextId() string {
	id, err := flake.NextID()
	if err!=nil{
		log.Println(err.Error())
	}
	//convert id to byte
	idByte := make([]byte, 8)
	for i := 0; i < 8; i++ {
		idByte[i] = byte(id >> uint(i*8))
	}
	//base64 encode
	return base64.URLEncoding.EncodeToString([]byte(idByte))
}

func GetProcess(id string) *model.Process {
	if v, ok := processMap.Load(id); ok {
		return v.(*model.Process)
	}

	return nil
}

func PutProcess(id string, process *model.Process) {
	processMap.Store(id, process)
}

func IsStringIn(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

func GetIndex(num int64) (res []int) {
	if num <= 1 {
		res = append(res, 0)
		return
	}
	if num%2 != 0 {
		res = append(res, 0)
		num--
	}

	// now num is even
	for num != 0 {
		var minus int64
		var i int64 = 2
		exp := 0
		for ; i <= num; i *= 2 {
			minus = i
			exp++
		}
		res = append(res, exp)
		num -= minus
	}
	return
}

func GetIntValue(num int) (res []int) {
	if num <= 1 {
		res = append(res, 1)
		return
	}
	if num%2 != 0 {
		res = append(res, 1)
		num--
	}

	// now num is even
	for num != 0 {
		minus := 0
		for i := 2; i <= num; i *= 2 {
			minus = i
		}
		res = append(res, minus)
		num -= minus
	}
	return
}
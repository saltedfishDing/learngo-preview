package helper

import (
	"math/rand"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err){
		return false, err
	}

	return true, nil
}

func RandomString(length int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	res := make([]rune, length)
	for i := 0;i<length;i++ {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(res)
}

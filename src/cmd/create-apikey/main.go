package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func generatePassword(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func main() {
	password, err := generatePassword(25)
	if err != nil {
		fmt.Println("パスワードの生成に失敗しました！")
		return
	}

	// パスワードをSHA-256でハッシュ化
	hash := sha256.Sum256([]byte(password))
	hashString := hex.EncodeToString(hash[:])

	// ログ出力
	fmt.Println("APIキーとそのパスワードを生成しました！")
	fmt.Println("APIキー:", hashString)
	fmt.Println("パスワード:", password)
}

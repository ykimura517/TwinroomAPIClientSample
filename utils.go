package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/big"
	"regexp"
)

type Target string

const (
	TargetMsg   Target = "msg"
	TargetVoice Target = "voice"
)

func getTargetPartFromResString(resString string, targetPart Target) *string {
	//lineから'msg=*'で囲まれた部分を正規表現で抽出
	var pattern string

	//Be more careful when processing in production
	if targetPart == TargetMsg {
		pattern = `msg=(.*?)\s`
	} else if targetPart == TargetVoice {
		pattern = `voice=(.*?)\s`
	} else {
		panic("Invalid targetPart")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		panic(err)
	}

	match := re.FindStringSubmatch(resString)
	if len(match) > 1 {
		if match[1] == "None" {
			return nil
		}
		if len(match[1]) < 2 {
			return nil
		}
		// 最初と最後の文字を除去
		trimmed := match[1][1 : len(match[1])-1]
		return &trimmed
	} else {
		fmt.Println("No match found")
	}
	return nil
}

// ランダムなファイル名を生成する関数
func generateRandomFileName() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.mp3", n.String()), nil
}

// Base64エンコードされた文字列をデコードしてファイルに保存する関数
func saveBase64EncodedAudio(base64Audio string) error {
	// Base64デコード
	audioData, err := base64.StdEncoding.DecodeString(base64Audio)
	if err != nil {
		return err
	}

	// ランダムなファイル名の生成
	fileName, err := generateRandomFileName()
	if err != nil {
		return err
	}
	folder := "voice/"
	// ファイルへの書き込み
	err = ioutil.WriteFile(folder+fileName, audioData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Audio file saved as: %s\n", fileName)
	return nil
}

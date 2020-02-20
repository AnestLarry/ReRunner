package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"libs"
	"os"
	"sync"
	"time"
)

var (
	RRSTemplate = map[string]interface{}{"WatchFiles": map[string]string{"filename1": "filehash", "filename2": "filehash"},
		"RunCommands": [][]string{[]string{"cmd.exe", "/c", "cls"}, []string{"cmd.exe", "/c", "dir"}},
		"ByLine":      false}

	RRSettingsFileNameString = "./RRSetting.json"
	TemplateBytes, _         = json.Marshal(RRSTemplate)
	Version                  = "Feb 15,2020."
	RRSettings               = map[string]interface{}{}
	IsRun                    = false
	wg                       sync.WaitGroup
)

func Init() {
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "h", "-h", "help":
				fmt.Println(
					"Run it if you need a setting template\n" +
						"Args:\n" +
						" h  - show this help\n" +
						" v  - get version\n")
				os.Exit(0)
			case "v", "-v", "version":
				fmt.Println(Version)
				os.Exit(0)
			default:
				fmt.Println("Do you mean \"-h\" ?")
				os.Exit(0)
			}
		}
	}
	if libs.LibsXExists(RRSettingsFileNameString) {
		if !libs.LibsXIsFile(RRSettingsFileNameString) {
			fmt.Println("RRSetting.json is not a file.")
			os.Exit(0)
		}
		data, err := ioutil.ReadFile(RRSettingsFileNameString)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(data, &RRSettings)
	} else {
		ioutil.WriteFile(RRSettingsFileNameString, TemplateBytes, 0644)
		os.Exit(0)
	}

}

func CheckFileHash() {
	for k, v := range RRSettings["WatchFiles"].(map[string]interface{}) {
		if libs.LibsXExists(k) {

			data := libs.LibsXSha1FileString(k)
			if v != data {
				RRSettings["WatchFiles"].(map[string]interface{})[k] = data
				IsRun = true
			}
		}
	}
	time.Sleep(500 * time.Millisecond)

}

func RunCommands() {
	if IsRun {
		libs.LibsXClear()
		for _, v := range RRSettings["RunCommands"].([]interface{}) {
			var cArgs []string
			for _, value := range v.([]interface{}) {
				cArgs = append(cArgs, value.(string))
			}
			if RRSettings["ByLine"].(bool) {
				libs.LibsXexecCommand(cArgs[0], cArgs[1:])
			} else {
				libs.LibsXExecShell(cArgs[0], cArgs[1:])
			}
		}
		SaveToRRSettings(RRSettings)
		IsRun = false
	}
}

func SaveToRRSettings(v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(RRSettingsFileNameString, data, 0644)
}

func Process() {
	for {
		CheckFileHash()
		RunCommands()
	}
}

func main() {
	Init()
	//RRSettings["RunCommands"].(map[string]interface{})["cmd.exe"] = []string{"command"}
	//RRSettings["WatchFiles"].(map[string]interface{})["filename1"] = "filehash"
	wg.Add(1)
	go Process()
	wg.Wait()
}

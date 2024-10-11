package main

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
)

type FileObj struct {
    Id      int64  `json:"id"`
    Name    string `json:"name"`
    Size    int64  `json:"size"`
    ModTime string `json:"modTime"`
}

func main() {
    // 指定要遍历的文件夹路径
    // 获取命令行参数
    folderPath := os.Args[1]
    outputFile := os.Args[2]

    files := make([]FileObj, 0)

    // 创建一个新文件，用于保存文件名
    file, err := os.Create(outputFile)
    if err != nil {
        fmt.Println("无法创建文件:", err)
        return
    }
    defer file.Close()

    var idCnt int64 = 1

    // 遍历指定文件夹下的所有文件
    err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println("无法遍历文件夹:", err)
            return nil
        }

        // 如果文件不是目录，则将其名称写入文件
        if !info.IsDir() {
            files = append(files, FileObj{
                Id:      idCnt,
                Name:    path,
                Size:    info.Size(),
                ModTime: info.ModTime().String(),
            })
            idCnt += 1
        }

        return nil
    })

    // fmt.Println(files)
    jsonData, err := json.Marshal(files)

    _, err = file.Write(jsonData)
    if err != nil {
        fmt.Println("无法写入文件:", err)
    } else {
        fmt.Println("文件名已保存到", outputFile)
        // fmt.Println(string(jsonData))
    }
}

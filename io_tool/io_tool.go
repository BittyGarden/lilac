package io_tool

import (
    "os"
)

//file exists return true
func FileExists(filePath string) bool {
    if _, err := os.Stat(filePath); err == nil || os.IsExist(err) {
        return true
    }
    return false
}

//file not exists return true
func FileNotExists(filePath string) bool {
    return !FileExists(filePath)
}

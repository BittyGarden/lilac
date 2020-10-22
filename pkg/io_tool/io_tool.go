package io_tool

import (
    "os"
)

func FileExists(filePath string) bool {
    if _, err := os.Stat(filePath); err == nil || os.IsExist(err) {
        return true
    }
    return false
}

func FileNotExists(filePath string) bool {
    return !FileExists(filePath)
}

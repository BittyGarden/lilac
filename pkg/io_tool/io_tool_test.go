package io_tool

import (
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
)

func TestFileExists(t *testing.T) {
    filePath := "awukknxs.txt"
    assert.False(t, FileExists(filePath), "file should not exists")
    file, err := os.Create(filePath)
    file.Close()
    if err == nil {
        assert.True(t, FileExists(filePath), "file should exists")
        os.Remove(filePath)
    }
    assert.False(t, FileExists(filePath), "file should not exists")
}

func TestFileNotExists(t *testing.T) {
    filePath := "awukknxs.txt"
    assert.True(t, FileNotExists(filePath), "file should not exists")
    file, err := os.Create(filePath)
    file.Close()
    if err == nil {
        assert.False(t, FileNotExists(filePath), "file should exists")
        os.Remove(filePath)
    }
    assert.True(t, FileNotExists(filePath), "file should not exists")
}

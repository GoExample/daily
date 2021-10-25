package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "fmt"
)

func main() {
    nonce := "37b8e8a308c354048d245f6d"
    key := "AES256Key-32Characters1234567890"
    plainText := "hello"
    cipherText := encryptString(plainText, key, nonce)
    newPlain := decryptString(cipherText, key, nonce)

    fmt.Println("plain:", plainText)
    fmt.Println("cipher:", cipherText)
    fmt.Println("new plain:", newPlain)
}

func encryptString(src, k, n string) string {
    // The key argument should be the AES key, either 16 or 32 bytes
    // to select AES-128 or AES-256.
    key := []byte(k)
    plaintext := []byte(src)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err.Error())
    }

    nonce, _ := hex.DecodeString(n)
    fmt.Println(nonce)

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }

    ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
    return fmt.Sprintf("%x", ciphertext)
}

func decryptString(src, k, n string) string {
    // The key argument should be the AES key, either 16 or 32 bytes
    // to select AES-128 or AES-256.
    key := []byte(k)
    ciphertext, _ := hex.DecodeString(src)

    nonce, _ := hex.DecodeString(n)

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err.Error())
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }

    plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        panic(err.Error())
    }

    return string(plaintext)
}

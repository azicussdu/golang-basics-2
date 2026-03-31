package main

//import (
//	"crypto/sha256"
//	"encoding/hex"
//	"fmt"
//)
//
//type User struct {
//	Name     string
//	Password string
//	Hash     string
//}
//
//func NewUser(name, password string) User {
//	hash := sha256.Sum256([]byte(password))
//	hashedPassword := hex.EncodeToString(hash[:])
//
//	user := User{Name: name, Password: password, Hash: hashedPassword}
//	return user
//}
//
//func main() {
//	us := NewUser("Arman", "Qwerty")
//	fmt.Println(us)
//}

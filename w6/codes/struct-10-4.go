package main

//import "fmt"
//
//type DBConfig struct {
//	IpAddress string
//	Port      int
//	Username  string
//	Password  string
//}
//
//func NewDBConfig(ipAdd string, port int, uname, pass string) (*DBConfig, error) {
//	if port < 1 || port > 65535 {
//		return nil, fmt.Errorf("invalid port %d", port)
//	}
//	return &DBConfig{
//		IpAddress: ipAdd,
//		Port:      port,
//		Username:  uname,
//		Password:  pass,
//	}, nil
//}
//
//func main() {
//	dbc, err := NewDBConfig("203.21.43.54", 4323, "dbuser", "dbpassword")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(dbc)
//	}
//}

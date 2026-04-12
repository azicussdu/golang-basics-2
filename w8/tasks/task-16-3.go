package main

type Notifier interface {
	Notify() string
}

type User struct {
	Name       string
	MyNotifier Notifier // MyNotifier = EmailNotifier{} | SmsNotifier{}
}

type EmailNotifier struct {
	Email string
}

func (en EmailNotifier) Notify() string {
	return "Sending email to " + en.Email
}

type SmsNotifier struct {
	PhoneNumber string
}

func (sn SmsNotifier) Notify() string {
	return "Sending SMS to number " + sn.PhoneNumber
}

//func main() {
//	userNurb := User{
//		Name:       "Nurbakhyt",
//		MyNotifier: SmsNotifier{PhoneNumber: "8707232312"},
//	}
//
//	userRakhm := User{
//		Name:       "Rakhmanberdi",
//		MyNotifier: EmailNotifier{Email: "rakhmanberdi@gmail.com"},
//	}
//
//	fmt.Println(userNurb.MyNotifier.Notify())
//	fmt.Println(userRakhm.MyNotifier.Notify())
//}

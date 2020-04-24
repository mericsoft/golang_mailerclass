# golang_mailerclass
a simple smtp mail library 
this library uses smtp auth
if you want to attach any file, you must insert full path

example 
////////////////////////////////////////////////////////////////////////////////////////
func main() {args := os.Args
 
	maliim := Mailer{
		Conn:        MailCon{"smtp.gmail.com", "587", "abc@gmail.com","xxxxxx"},
		Name:        "John Doe",
		From:        info@xyz.com,
		To:          []string{"abc@gmail.com"},
		Subject:     "test message",
		ContentType: "text/html",
		Body:        "this is an <b>example</b> message",
    File :       "/home/aaa/Pictures/abc.png",  
	}
	
	fmt.Println(maliim.SendMail())
}
///////////////////////////////////////////////////////////////////////////////////////////////

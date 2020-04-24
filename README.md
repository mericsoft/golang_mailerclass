# golang_mailerclass
a simple smtp mail library 


example 

////////////////////////////////////////////////////////////////////////////////////////

 
	maliim := Mailer{
		Conn:        MailCon{"smtp.gmail.com", "587", "abc@gmail.com","xxxxxx"},
		Name:        "John Doe",
		From:        info@xyz.com,
		To:          []string{"abc@gmail.com"},
		Subject:     "test message",
		ContentType: "text/html",
		Body:        "this is an <b>example</b> message",
                File :       "/home/aaa/Pictures/abc.png",  
	        ReturnMessage:"your mail has been sent",
	}
	
	fmt.Println(maliim.SendMail())
}

///////////////////////////////////////////////////////////////////////////////////////////////

for your questions : mericsoftware@gmail.com

package main

import (
	"bytes"
	"fmt"
	"github.com/mxk/go-imap/imap"
	"log"
	"net/mail"
	"net/smtp"
	"time"
)

func send() {
	auth := smtp.PlainAuth("", "recipient@example.com", "passwd", "mailhost.example.com")

	to := []string{"recipient@example.com"}
	msg := []byte("To: recipient@example.com\r\n" +
		"Subject: This is a test message\r\n" +
		"\r\n" +
		"This is the test body of the message.\r\n")
	err := smtp.SendMail("relay.example.com:25", auth, "noreply@example.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func check_imap() {
	var (
		c   *imap.Client
		cmd *imap.Command
		rsp *imap.Response
	)

	c, _ = imap.Dial("mailhost.example.com")
	defer c.Logout(30 * time.Second)
	fmt.Println("Server says hello:", c.Data[0].Info)
	c.Data = nil

	if c.Caps["STARTTLS"] {
		c.StartTLS(nil)
	}

	if c.State() == imap.Login {
		c.Login("recipient@example.com", "passwd")
	}

	cmd, _ = imap.Wait(c.List("", "%"))
	fmt.Println("\nTop-Level mailboxes:")

	for _, rsp = range cmd.Data {
		fmt.Println("Server data: ", rsp)
	}

	c.Data = nil

	c.Select("INBOX", true)
	fmt.Println("\nMailto status:", c.Mailbox)

	set, _ := imap.NewSeqSet("")
	if c.Mailbox.Messages >= 10 {
		set.AddRange(c.Mailbox.Messages-9, c.Mailbox.Messages)
	} else {
		set.Add("1:*")
	}

	cmd, _ = c.Fetch(set, "RFC822.HEADER")

	fmt.Println("\nMost recent Messages:")
	for cmd.InProgress() {
		c.Recv(-1)

		for _, rsp = range cmd.Data {
			header := imap.AsBytes(rsp.MessageInfo().Attrs["RFC822.HEADER"])
			if msg, _ := mail.ReadMessage(bytes.NewReader(header)); msg != nil {
				fmt.Println("|--", msg.Header.Get("Subject"))
			}
		}
		cmd.Data = nil
	}

	if rsp, err := cmd.Result(imap.OK); err != nil {
		if err == imap.ErrAborted {
			fmt.Println("Fetch command aborted")
		} else {
			fmt.Println("Fetch error:", rsp.Info)
		}
	}
}

func main() {
	send()
	check_imap()
}

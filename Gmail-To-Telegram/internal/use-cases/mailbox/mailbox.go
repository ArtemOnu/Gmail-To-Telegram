package mailbox

import (
	log "gomod/internal/entities"
	"io"
	"strings"

	"github.com/emersion/go-imap"
	Imap "github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"golang.org/x/net/html"
)

type Mail struct {
	client  *Imap.Client
	mailbox *imap.MailboxStatus
	envelop *imap.Envelope
	uid     uint32
	text    []byte
}

// Read the contents
func (m *Mail) reader(r imap.Literal) {
	mread, _ := mail.CreateReader(r)
	for {
		p, err := mread.NextPart()
		//We come to the end and stop the cycle
		if err == io.EOF {
			break
		}
		m.text, _ = io.ReadAll(p.Body)

	}
}

// Processing the letter
func (m *Mail) Fetch() {
	seq := new(imap.SeqSet)
	seq.AddNum(m.mailbox.Messages)
	var section imap.BodySectionName
	//We choose, what exactly we will take from the letter
	items := []imap.FetchItem{
		imap.FetchEnvelope,
		imap.FetchBodyStructure,
		section.FetchItem(),
	}
	//Create channels
	done := make(chan error, 1)
	messages := make(chan *imap.Message, 1)

	go func() {
		done <- m.client.Fetch(seq, items, messages)
	}()

	msg := <-messages
	m.uid = msg.Uid
	err := <-done
	if err != nil {
		panic(err)
	}

	m.envelop = msg.Envelope
	m.reader(msg.GetBody(&section))
}

func (m *Mail) Disconnect() {
	m.client.Logout()
	log.Log("Bot discconect")
}

// Selecting a mailbox from which to receive emails
func (m *Mail) seelect() error {
	var err error
	m.mailbox, err = m.client.Select("INBOX", false)
	return err
}

func (m *Mail) Connect(mail string, token string) error {
	var err error
	m.client, err = Imap.DialTLS("imap.mail.ru:993", nil)
	if err == nil {
		err = m.client.Login(mail, token)
		if err == nil {
			m.seelect()
			return err
		}
		return err
	}
	return err
}

// AI insert
// We extract the text of the letter, which is presented in HTML format.
func (m *Mail) ClearText() (string, error) {
	doc, err := html.Parse(strings.NewReader(string(m.text)))
	if err == nil {
		var textbuffer strings.Builder
		var walking func(*html.Node)
		walking = func(n *html.Node) {
			if n.Type == html.TextNode {
				t := strings.TrimSpace(n.Data)
				if t != "" {
					textbuffer.WriteString(t)
					textbuffer.WriteString("\n")
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				walking(c)
			}
		}
		walking(doc)
		return textbuffer.String(), nil
	}
	return "", nil
}

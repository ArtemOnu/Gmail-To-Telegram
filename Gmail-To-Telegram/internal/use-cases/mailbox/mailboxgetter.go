package mailbox

import "time"

//Receive html letters
func (m *Mail) GetBody() []byte {
	return m.text
}

//Convert HTML to Text (Get Information Only)
func (m *Mail) GetFormaBody() (string, error) {
	buf, err := m.ClearText()
	if err == nil {
		return buf, nil
	}
	return "", err
}

//Date formatting
func (m *Mail) GetDate() string {
	location, _ := time.LoadLocation("Asia/Vladivostok")
	return m.envelop.Date.In(location).Format("02 Jan 2006 15:04")
}

func (m *Mail) GetBox() string {
	str := m.envelop.To[0].MailboxName + "@" + m.envelop.To[0].HostName
	return str
}

//Get the author of the letter
func (m *Mail) GetAuthor() string {
	return m.envelop.From[0].PersonalName
}

//We check whether the mailbox is empty or not
func (m *Mail) CheckMessage() bool {
	if m.mailbox.Messages > 0 {
		return true
	}
	return false
}

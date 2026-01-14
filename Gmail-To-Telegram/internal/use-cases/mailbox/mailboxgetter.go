package mailbox

//Receive html letters
func (m *Mail) GetBody() string {
	return string(m.text)
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
	return m.envelop.Date.Format("02-Jan-2006 15:04")
}

//Get the author of the letter
func (m *Mail) GetAuthor() string {
	return m.GetAuthor()
}

//We check whether the mailbox is empty or not
func (m *Mail) CheckMessage() bool {
	if m.mailbox.Messages > 0 {
		return true
	}
	return false
}

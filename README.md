
# Mail to telegram

This is my second project, developed for educational purposes. The GMAIL email service is currently down. The project was originally written for mail.ru.



## Run Locally

Clone the project

```bash
  git clone https://github.com/ArtemOnu/Gmail-To-Telegram
```

Go to the project directory

```bash
  cd Gmail-To-Telegram
```

Install dependencies

```bash
  go get "github.com/emersion/go-imap"
  go get "github.com/emersion/go-imap/client"
  go get "github.com/emersion/go-message/mail"
  go get "golang.org/x/net/html"
  go get "github.com/go-telegram-bot-api/telegram-bot-api/v5"
```
Customize your json

```
  {
    "Token": "Token telegram bot",
    "Chat-ID": "Chat ID",
    "Mail-addres": "your mail",
    "Password": "your password",
    "Host": "your host"
}
```
Build project

```
  cd cmd
  go build main.go
```
Start (If you decide not to compile the project, otherwise just run the .exe)

```bash
  cd cmd
  go run main.go
```

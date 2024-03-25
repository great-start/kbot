<p align="center">
  <h3 align="center">Telegram bot</p>
</p>

<p align="center">Telegram bot that returns particular pictures after pressing a button</p>

#### [Telegram link](https://t.me/ev_bill_kbot)

### Description

Telegram bot that returns particular pictures after pressing a button

### Technologies

- Golang (v1.22.1)
- Telebot (gopkg.in/telebot.v3)

## Getting started

---

### Clone repository

```bash
$ git clone https://github.com/great-start/kbot.git
```

### Install dependencies

```bash
$ go get
```

### Setting Environment

```bash
$ export TELEGRAM_BOT_TOKEN=<bot_token>
```

### Start bot

```bash
$ go build -ldflags "-X=github.com/great-start/kbot/cmd.appVersion=v1.0.3"
```

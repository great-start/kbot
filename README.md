<h2 align="center">Telegram bot</h2>

<h3 align="center">Bot returns particular pictures after pressing a buttons</h3>

---

#### [TelegramBot link](https://t.me/ev_bill_kbot)

### Technologies

- Golang (v1.22.1)
- Telebot (gopkg.in/telebot.v3)

---

## Getting started

### Clone repository

```bash
git clone https://github.com/great-start/kbot.git
```

### Setting up environment

```bash
$ export TELEGRAM_BOT_TOKEN=<bot_token>
```

### Install dependecies & Build

```bash
go build -ldflags "-X=github.com/great-start/kbot/cmd.appVersion=v1.0.3"
```

### Start bot

```bash
./kbot start
```

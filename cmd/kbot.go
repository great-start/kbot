/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

var (
	// Telegram bot token
	TeleBotToken = os.Getenv("TELE_TOKEN")
	menu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	apple        = menu.Text("apple")
	car          = menu.Text("car")
	flower       = menu.Text("flower")
	bird         = menu.Text("bird")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("bot %s started", appVersion)
		sett := tele.Settings{
			URL:    "",
			Token:  os.Getenv("TELE_TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}

		bot, err := tele.NewBot(sett)
		if err != nil {
			log.Fatalf("Check TELE_TOKEN env. %s", err)
			return
		}

		menu.Reply(
			menu.Row(apple, car),
			menu.Row(flower, bird),
		)

		bot.Handle(tele.OnMyChatMember, func(ctx tele.Context) error {
			// ctx.Send("Hi! 👋")
			    if err := ctx.Send("Hi! 👋"); err != nil {
        return err
    }
			return ctx.Send("Push a button to show picture", menu)
		})

		dir, err := os.Getwd()

		apple_pic := &tele.Photo{File: tele.FromDisk(fmt.Sprintf("%s/.img/apple.jpeg", dir))}
		car_pic := &tele.Photo{File: tele.FromDisk(fmt.Sprintf("%s/.img/car.jpeg", dir))}
		flower_pic := &tele.Photo{File: tele.FromDisk(fmt.Sprintf("%s/.img/flower.jpeg", dir))}
		bird_pic := &tele.Photo{File: tele.FromDisk(fmt.Sprintf("%s/.img/bird.jpeg", dir))}

		bot.Handle(tele.OnText, func(ctx tele.Context) error {
			switch ctx.Text() {
			case "apple":
				_, err := bot.SendAlbum(ctx.Sender(), tele.Album{apple_pic})
				if err != nil {
					return err
				}
				return nil
			case "car":
				_, err := bot.SendAlbum(ctx.Sender(), tele.Album{car_pic})
				if err != nil {
					return err
				}
				return nil
			case "flower":
				_, err := bot.SendAlbum(ctx.Sender(), tele.Album{flower_pic})
				if err != nil {
					return err
				}
				return nil
			case "bird":
				_, err := bot.SendAlbum(ctx.Sender(), tele.Album{bird_pic})
				if err != nil {
					return err
				}
				return nil

			}

			return err
		})

		bot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

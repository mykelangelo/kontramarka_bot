/*
Package telego provides one-to-one Telegram Bot API method & types.

Telego features all methods and types described in official Telegram documentation (https://core.telegram.org/bots/api).
It achieves this by generating methods and types from docs (generation is in internal/generator package).

The main goal was and is to create one-to-one library, so that if you know how Telegram bots work, you will immediately
know how to implement that in Go using Telego.

One-to-one implementation

All types named and contains same information as documented by Telegram, for methods it's exactly the same, however,
some minor differences may be present (like use of interfaces or combined types).
Also, all generated code has same description as in Telegram docs, so there is actually no need to go to docs (but
still, be careful as it is not a full copy of docs due to text only limitation).

Telego was also created to simplify work with Telegram API, so some additional methods for more convenient usage
located in helpers.go and telegoutil package.

When you are working with things like chat ID which can be an integer or string Telego provides combined types:

	type ChatID struct {
		ID       int64
		Username string
	}

or input files that can be URL, file ID or actual file data:

	type InputFile struct {
		File telegoapi.NamedReader
		FileID string
		URL string
	}

you will specify only one of fields and Telego will figure out what to do with that.

For more flexibility file data for InputFile is provided via simple interface:

	type NamedReader interface {
		io.Reader
		Name() string
	}

this interface already implemented by os.File, so you can use it directly.

Example

Most of the examples can be seen in examples folder.

Simple echo bot:

	package main

	import (
		"fmt"
		"os"

		"github.com/mymmrac/telego"
		tu "github.com/mymmrac/telego/telegoutil"
	)

	func main() {
		botToken := os.Getenv("TOKEN")

		// Create Bot with debug on
		bot, err := telego.NewBot(botToken, telego.WithDefaultLogger(true, true))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Get updates channel
		updates, _ := bot.UpdatesViaLongPulling(nil)

		// Stop reviving updates from updates channel
		defer bot.StopLongPulling()

		// Loop through all updates when they came
		for update := range updates {
			// Check if update contains message
			if update.Message != nil {
				// Get chat ID from message
				chatID := tu.ID(update.Message.Chat.ID)

				// Copy sent message back to user
				_, _ = bot.CopyMessage(&telego.CopyMessageParams{
					ChatID:     chatID,
					FromChatID: chatID,
					MessageID:  update.Message.MessageID,
				})
			}
		}
	}

This bot will send same messages as you sent to him.
*/
package telego

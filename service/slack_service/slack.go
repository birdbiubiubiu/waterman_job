package slack_service

import (
	"fmt"
	"github.com/slack-go/slack"
	"time"
	"waterman_job/models"
	"waterman_job/pkg/logging"
	"waterman_job/pkg/setting"
)

var SwapWhaleCh = make(chan *models.Whales, 100)

func LiquidityAlert() {
	for msg := range SwapWhaleCh {
		msgText := ""
		tm := time.Unix(int64(msg.Timestamp), 0)

		if msg.Action == "mint" {
			msgText = fmt.Sprintf("Added to Liquidity! \n %s \n \n %s - %f + \n %s - %f \n Worth: $ %f USD\n https://etherscan.io/tx/%s",
				tm.Format("2006-01-02 15:04:05"), msg.Token0, msg.Amount0, msg.Token1, msg.Amount1, msg.AmountUsd, msg.TransactionId)
		} else {
			msgText = fmt.Sprintf("Removed to Liquidity! \n %s \n \n %s - %f + \n %s - %f \n Worth: $ %f USD\n https://etherscan.io/tx/%s",
				tm.Format("2006-01-02 15:04:05"), msg.Token0, msg.Amount0, msg.Token1, msg.Amount1, msg.AmountUsd, msg.TransactionId)
		}
		title := fmt.Sprintf("%s 🐋 Whale Alert 🐋", msg.Platform)
		send(title, msgText, setting.SlackSetting.Token, setting.SlackSetting.BridgeChannelId)
		// slack 消息频率限制
		time.Sleep(time.Second)
	}
}

func send(title, msgText, token, channelId string) {
	api := slack.New(token)
	attachment := slack.Attachment{
		Text: msgText,
	}

	_, _, err := api.PostMessage(
		channelId,
		slack.MsgOptionText(title, false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)

	if err != nil {
		logging.Error(err)
		return
	}
}

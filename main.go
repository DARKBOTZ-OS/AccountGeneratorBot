//   Account Generator Bot
//   Copyright (C) 2021 AnonyIndian (@xnony)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.


package main

import (
	"fmt"
	"math/rand"
	"strings"
	"strconv"
	"time"
	"os"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

var MSG string

func main() {
	// Put Your Bot Token via ENV Vars
	b, err := gotgbot.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(b, nil)
	dispatcher := updater.Dispatcher

	// Handlers for runnning commands.
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("gen", gen))

	err = updater.StartPolling(b, &ext.PollingOpts{Clean: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\nMade with ❤️ by @DGABHI .\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		// Put your accounts here
		// "email:pass",
		// Following are some demo accounts
		"radhab144@gmail.com:Chakri18",
                "shush2128@gmail.com:chinmun9",
                "sukhap@gmail.com:visu1985#
                 "rupesh.tech@gmail.com:rupesh80", "siriparuchurisiri@gmail.com:Aruna@31",
                 "senthazai.s@gmail.com:flower.143",
                 "Raghav.online1989@gmail.com:Raghav.online1989", "kesinenieswar11@gmail.com:Honey@1211",
                 "prithwishgiri2010@gmail.com:Dulu@123",
                 "abhishek11.tripathi@gmail.com:gabbudada",
                 "sri18.satya44@gmail.com:Bhavikaraj18b12",
                 "sandeep.br@gmail.com:Shravya2009",
                 "pandey.mantosh@gmail.com:mantosh1212",
                 "chinnagopal@gmail.com:bngrm@143",
                 "muji1990@gmail.com:muji@1990", 
                 "roshan.anand222@gmail.com:roshan.anand222",
                 "shoumya.taps@gmail.com:ghontu.123",
                 "vinayakahk@yahoo.com:pulsar100", 
                 "sweta992001@yahoo.com:mantu975",
                 "priyal.goyal@hotmail.com:Prishi@123",
                 "sandip.some@gmail.com:sandip.123", 
                 "nannapanenibhargavnagesh@gmail.com:raghu2358",
                 "pgrover005@gmail.com:p@rth005",
                 "rxtrived@gmail.com:kriti131", 
                "shantanuresistivity@gmail.com:Shantanu_7", 
                "sathiyaprakash13@gmail.com:Aisu@1996",
                "mrunalpalande059@gmail.com:mrunal1212",
                "vasan869@gmail.com:sakshi869",
                "santhosh.sivaraman@outlook.com:Naresh@1128",
                "sanjivboy2000@gmail.com:optimist",
                "shwetamishra02vit@gmail.com:shweta#12345",
                "richa.arora25@gmail.com:anju12vinod",
                "supriya.chalagam@gmail.com:vamsi921",
                "saurabh.subu@gmail.com:drsharma85",
                "sshakeel@hotmail.com:Ash$2011",
                "vasan869@gmail.com:sakshi869",
                "uramu123@gmail.com:udatha123",
                "simzcool15@gmail.com:jasvirsingh04",
                "bchinna807@gmail.com:Lakshmikanth.1",
                "jagadeesh41@yahoo.co.in:Rukmani123!",
                "anikakuk@gmail.com:anika05011988",
                "sujithredrocks@gmail.com:Sujith@12",
                "ursrihari.08@gmail.com:@sri1984",
                "shroffmv@hotmail.com:iindiann25",
                "manvibhardwaj1993@gmail.com:manvi0201",
                "prakashashok999@gmail.com:9963200281",
                "neethugeetha10@gmail.com:neethu@123#",
                "tkd.mca@gmail.com:tapaN@123",
                "sahil1986scorpio@gmail.com:ktmrc390",
                "rhymes2rashmi@gmail.com:Bindiya-12",
                "rakesh2meera@gmail.com:rakesh2meera",
                "swathi.pandurangi@gmail.com:swalokam1234",
                "sumanganguly1985@gmail.com:Pa11word",
               "durgaprasad.kota@gmail.com:prasad9111",
               "sarojiniramesh15@gmail.com:saro1711",
               "sai.ravipudi@gmail.com:ramsai999",
               "prbatta@gmail.com:prbatt1A",
               "shamik98@gmail.com:Shamik98*",
               "bishwaroop.dey@gmail.com:mnbv1234",
               "rp.preet39@gmail.com:raman4447",
               "naren_tataisp@yahoo.com:naren8867",
               "vijayasai.talluri@gmail.com:Sunnysony$",
              "shalluchhillar@gmail.com:Shalluchhillar2004",
              "sknadeem7869@gmail.com:Mohammed14@",
              "shayandas88@gmail.com:dodgevipeR_1",
              "n.sreenivasulu2710@gmail.com:vijaya2710",
              "sharmisthabhattacharyya1820@gmail.com:Misti@18",
              "grk1124@gmail.com:Bangaram@1124",
              "pal.abhijit1@gmail.com:pal.abhijit1",
              "Praveen6938@gmail.com:Praveen@63",
              "mohit.bhtngr@gmail.com:mani2515",
              "sampath.trainer@gmail.com:126011207",
              "responc@gmail.com:farakka14",
             "priyanka.v1010@gmail.com:Deeptrisha@2008",
             "shadab.sm@gmail.com:shaddy161",
	}

	return formats[rand.Intn(len(formats))]
}

func start(ctx *ext.Context) error {
	// To ensure bot does not reply outside of private chats
	if ctx.EffectiveChat.Type != "private" {
		return nil
	}
	// Following string is replied to cmd user on /start 
	MSG = "*Hi %v,\n" +
		"I am an Account Generator Bot\n" +
		"-------------------------------------------------\n" +
		"I can provide premium accounts of different services\n" +
		"--------------------------------------------------\n" +
		"Do /gen to generate an account\n" +
		"--------------------------------------------------\n" +
		"❤️Brought to You By @DGABHI❤️\n*"

	user := ctx.EffectiveUser
	channel_id, cerror := strconv.Atoi(os.Getenv("CHANNEL_ID"))
	if cerror != nil {fmt.Println("Please Provide me a valid Channel ID")}
	member, eror := ctx.Bot.GetChatMember(int64(channel_id), user.Id)
	if eror != nil {
		ctx.Bot.SendMessage(ctx.EffectiveChat.Id, "*Bot not admin in JoinCheck Channel.*", nil)
		return nil
	}
	// For Checking either user joined channel or not
	if member.Status == "member" || member.Status == "administrator" || member.Status == "creator" {
		_, err := ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf(MSG, user.FirstName), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})
		if err != nil {
			fmt.Println("failed to send: " + err.Error())
		}
	} else {
		// An Error message replied to command user if he's not member of the JoinCheck Channel
		ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf("*You must join %v To use me.*", os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{ParseMode: "Markdown"})
	}
	if strings.ToLower(os.Getenv("LOGS")) != "false"{
		logs, log_err := strconv.Atoi(os.Getenv("LOGS"))
		if log_err != nil{fmt.Println(log_err.Error())}
		// Following message is sent in Logs Group (if set)
		ctx.Bot.SendMessage(int64(logs), fmt.Sprintf("#Start\n\nBot Started By %v(%v)", user.FirstName, user.Id), nil)
	}
	return nil
}

func gen(ctx *ext.Context) error {
	// To ensure bot does not reply outside of private chats
	if ctx.EffectiveChat.Type != "private" {
		return nil
	}
	Combo := strings.Split(randomFormat(), ":")
	// Following string is replied to cmd user on /gen 
	MSG = "𝙃𝙚𝙧𝙚 𝙄𝙨 𝙔𝙤𝙪𝙧 %v 𝘼𝙘𝙘𝙤𝙪𝙣𝙩" +
		"\n\n𝙀𝙢𝙖𝙞𝙡: `%v`" +
		"\n𝙋𝙖𝙨𝙨: `%v`" +
		"\n𝙂𝙚𝙣𝙚𝙧𝙖𝙩𝙚𝙙 𝘽𝙮: *%v*" +
		"\n\n𝙏𝙝𝙖𝙣𝙠 𝙮𝙤𝙪 𝙛𝙤𝙧 𝙪𝙨𝙞𝙣𝙜 𝙢𝙚!\n❤️𝙎𝙝𝙖𝙧𝙚 & 𝙎𝙪𝙥𝙥𝙤𝙧𝙩 *%v*❤️"

	user := ctx.EffectiveUser
	channel_id, cerror := strconv.Atoi(os.Getenv("CHANNEL_ID"))
	if cerror != nil {fmt.Println("Please Provide me a valid Channel ID")}
	member, eror := ctx.Bot.GetChatMember(int64(channel_id), user.Id)
	if eror != nil {
		ctx.Bot.SendMessage(ctx.EffectiveChat.Id, "Bot not admin in JoinCheck Channel", nil)
		return nil
	}
	// For Checking either user joined channel or not
	if member.Status == "member" || member.Status == "administrator" || member.Status == "creator" {
		_, err := ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf(MSG, os.Getenv("ACC_NAME"), Combo[0], Combo[1], user.FirstName, os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})
		if err != nil {
			fmt.Println("failed to send: " + err.Error())
		}
	} else {
		// An Error message replied to command user if he's not member of the JoinCheck Channel
		ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf("*You must join %v to use me.*", os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{ParseMode: "Markdown"})
	}
	if strings.ToLower(os.Getenv("LOGS")) != "false"{
		logs, log_err := strconv.Atoi(os.Getenv("LOGS"))
		if log_err != nil{fmt.Println(log_err.Error())}
		// Following message is sent in Logs Group (if set)
		ctx.Bot.SendMessage(int64(logs), fmt.Sprintf("#AccClaimed\n\n%v generated a new account.", user.FirstName), nil)
	}
	return nil
}

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
		"modith.raj@gmail.com:modith8568",
"syedabdulquader@gmail.com:London12",
"sk.bommidi@gmail.com:1Welcome",
"vinodup@gmail.com:biotec@21",
"deekshishetty697@gmail.com:shetty12",
"anandkamalgoel@gmail.com:Tathagat9",
"shru.nanjunda@gmail.com:Feb@1991",
"sumanganguly1985@gmail.com:Pa11word",
"sumanganguly1985@gmail.com:Pa11word",
"prem.galaxy191@gmail.com:Kalka@2137",
"poonamgr@gmail.com:poonam5757",
"syedabdulquader@gmail.com:London12",
"chandana1985hitha@gmail.com:chandu85",
"prem.galaxy191@gmail.com:Kalka@2137",
"pavithratresa@gmail.com:pavithratresa14",
"pardhu.pathella@gmail.com:welcomepps",
"pspriyanka.singh2000@gmail.com:dphhayden@123",
"paridhig13@gmail.com:omsai401",
"prabinyovan@gmail.com:Laana@123",
"poojashinde0182@gmail.com:Pooja1512",
"jainaj.123@gmail.com:akshay11",
"psrivastava4u@gmail.com:Jims@2017",
"pagarrohan12345@gmail.com:rdx_0807",
"nsd.prasad@gmail.com:King@208",
"rajesh.cobra201@gmail.com:anshika@201",
"puneet.chaturvedi@gmail.com:aditi1942",
"raineha743@gmail.com:raineha743",
"rajojunc65@gmail.com:rksai123",
"prachiranade2@gmail.com:Prachi1987#",
"rahulgargatte@gmail.com:rashi2117",
"msdakshu94@gmail.com:mahiakshu",
"monali_jeanie@hotmail.com:isspalko",
"narinder.sh59@gmail.com:shashi834",
"priyakrishnan1905@gmail.com:Krishna123#$",
"pawar.519@gmail.com:arnav!20",
"prbatta@gmail.com:prbatt1A",
"padmabhaskaram71@gmail.com:udbhav@123",
"nitin260659@gmail.com:bharti2493",
"pratima.kanumuri88@gmail.com:mangaraju88",
"pankaj@khemkas.in:308khemka",
"pridwik@gmail.com:lakshmiP1",
"priti.nahar@ymail.com:S@unshine8",
"raghava1684@gmail.com:raghava168143",
"patilabhi25296@gmail.com:abhi9139",
"prattaydutta@gmail.com:prattay3",
"pskuttappa@gmail.com:Series123@",
"poojajainfd@gmail.com:mishtii07",
"nishika.guttu@gmail.com:Nishika16",
"nandukiran11@gmail.com:Kiran.11",
"prabir69a@gmail.com:kelokoreche9",
"rajeevbhatt3999@gmail.com:6uec60ha",
"sachinjgawade@yahoo.com:Sach@1234",
"pal.othersites@gmail.com:palothersites",
"raighanshyam97@gmail.com:Ghannu@54321",
"radiraj@gmail.com:arkumar@1",
"himesh.thakur13@gmail.com:ihatemyself",
"pragati.v.bora@gmail.com:Kanchan321#",
"priscilla1585@gmail.com:babli@15",
"parixit87@gmail.com:0501206087",
"onkar.pvg@gmail.com:Swati19#",
"pramoj@gmail.com:tanu1680",
"jainaj.123@gmail.com:akshay11",
"nirajkrm@yahoo.com:Nishu00!",
"raghavendra.alwala@gmail.com:Kumar731",
"phanindranemala@gmail.com:annapurna@22",
"prernachaudhari64@gmail.com:prerna1902",
"padmalrm@gmail.com:9448039772",
"praveensh93@gmail.com:Praveen@306",
"monti2006@gmail.com:Qwerty@1323",
"priyanka.v1010@gmail.com:Deeptrisha@2008",
"sk.bommidi@gmail.com:1Welcome",
"prins.patel441@gmail.com:dhanu1416",
"rajesh.prabha10@gmail.com:Jimcarry@2",
"priyadarshi.mohapatra@gmail.com:Mohapatra71",
"deekshishetty697@gmail.com:shetty12",
"psa20011@rediffmail.com:@surya2016",
"rajeshwar.sable@gmail.com:Sindra@77",
"rajnish781@yahoo.co.in:madhur.01",
"priyanka.kantakari@gmail.com:Pinky@1293",
"pratiksha00@gmail.com:248803123",
"p8790600439@gmail.com:8790600439",
"archana.silpa@gmail.com:Silpa123",
"pratichi.mishra197@gmail.com:Adishree2357",
"praneeth.180694@gmail.com:Praneeth1$",
"priyankapgm.21@gmail.com:pgm@0210P",
"chandana1985hitha@gmail.com:chandu85",
"KONETICHANDRAREDDY@GMAIL.COM:ammananna@123",
"pavanbobbi@gmail.com:Pavan@1502",
"vibhu12345@gmail.com:123456yu",
"raghavim90@gmail.com:rag@1990",
"rahulsabharwal1@hotmail.com:27Decembe",
"nnayan2727@gmail.com:Nayan@27",
"pta1013@gmail.com:Angel@784",
"nabikrish@gmail.com:krish2807",
"raghukachalappa@gmail.com:Raghu@#87",
"pani.dasyam@gmail.com:chakri9999",
"pinkey0905@gmail.com:I'mhere4u",
"parimalseema@gmail.com:parimalseema1990",
"vinodup@gmail.com:biotec@21",
"naveen789sh@gmail.com:Naveen@405",
"mrunmayeek94@gmail.com:monu021119",
"rahulkamble1385@gmail.com:Dhamma358",
"psingh99@gmail.com:Star@99news",
"prajithprakashpillai@gmail.com:pillai.123",
"pathansalman886@gmail.com:Samzishu89",
"online.grv@gmail.com:snghl28grv",
"nitinvashistha17@gmail.com:Rocking#123",
"nikitamsheth@gmail.com:nikita#1711",
"priyank.bajani@gmail.com:Dentist@no1",
"prateekror@gmail.com:10316235",
"prasshanthshettyy@gmail.com:Nisha@1110",
"oviyaashokan12@gmail.com:oviyaashokan",
"sachinjgawade@yahoo.com:Sach@1234",

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

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
		"radhab144@gmail.com:Chakri18
shush2128@gmail.com:chinmun9
sukhap@gmail.com:visu1985#
supathra@rediffmail.com:Supa1hra
phanindranemala@gmail.com:annapurna@22
responc@gmail.com:farakka14
rajraj1901@yahoo.com:rajraj1901
vishu.9896@gmail.com:2Jangurpilov
shru.nanjunda@gmail.com:Feb@1991
Raghav.online1989@gmail.com:Raghav.online1989
raokishore@yahoo.com:Neelima30
shru.nanjunda@gmail.com:Feb@1991
shwetamishra02vit@gmail.com:shweta#12345
ssrini78@gmail.com:lakshya-12
s.paul@rediffmail.com:P!@ssw0rd
vadav.06@gmail.com:vaibhavyadav18
gauravrawatt@gmail.com:pakhi101010
thyaga81@gmail.com:thyaga81
nuthan.kotian@gmail.com:janu0422
saugata.1980@gmail.com:Mosaboni5
rayboutique@gmail.com:9323242599
s.paul@rediffmail.com:P!@ssw0rd
sajidns143@gmail.com:Sajid143#
abhinavadeepak123@gmail.com:abhi@994
tejagokina@gmail.com:tejagokina
praveenkmr78@gmail.com:pkumar@77
sudheer.gss@gmail.com:$suresh1
nitin260659@gmail.com:bharti2493
pragati.v.bora@gmail.com:Kanchan321#
mukatiraj@gmail.com:G@ecis999a
nidhishrdyk@gmail.com:Nidhish17
mukatiraj@gmail.com:G@ecis999a
parihar.ashvin@gmail.com:Cr@ft@123
carohit.soni@gmail.com:Abcd*1234
atul.a.patil78@gmail.com:Advait@2009
sircaranjan@gmail.com:ranu3434
sahejapo@gmail.com:03june1999
sandeep007z@yahoo.in:S@ndeep007z
naveen789sh@gmail.com:Naveen@405
swetavnaik0923@gmail.com:shravani@22
archana.silpa@gmail.com:Silpa123
ranavir.roy@gmail.com:ytia1234
dasturseema@gmail.com:Gulabjamun25
psrivastava4u@gmail.com:Jims@2017
pratiksha00@gmail.com:248803123
vg249089@icloud.com:Babita0812
asajrawat1993@gmail.com:211993asaj
shraddhajituri@gmail.com:Chetana@24
arinneeds@gmail.com:More#1234
shailesh.rbims@gmail.com:shailesh1305
sandyjaan1986@gmail.com:Sandy@3131
pani.dasyam@gmail.com:chakri9999
eds.sanjay@gmail.com:Chuttus#2013S
swarnava1979@gmail.com:ron@2011
KONETICHANDRAREDDY@GMAIL.COM:ammananna@123
dusane.suhas@gmail.com:Swara@1409
rocksays1911@gmail.com:nagamani@16
deepa369g@gmail.com:deepa369g
sreedharala.aishwarya@gmail.com:Aishu@92
nishika.guttu@gmail.com:Nishika16
salil.info@gmail.com:P@ssion02
ranavir.roy@gmail.com:ytia1234
sntshjgtp@gmail.com:Winsave@24362
sujatha3664@gmail.com:sanjay13189
neelcute@gmail.com:knsrao@123
syedabdulquader@gmail.com:London12
pandey.mantosh@gmail.com:mantosh1212
saurabh.subu@gmail.com:drsharma85
p8790600439@gmail.com:8790600439
srikanth8mantri@gmail.com:srikanth*88
samagata88@gmail.com:Samagata@88
nowaj.sharif@gmail.com:@raveena
padhulaknar@gmail.com:kavya0911
sai.ravipudi@gmail.com:ramsai999
pinkey0905@gmail.com:I'mhere4u
swetha.tall@gmail.com:Sarayu@31
sakshithakkar22@gmail.com:pices2002
shubra1977@gmail.com:shanaya2206
roshnireddyg@gmail.com:Roshnireddy9!
lohithravilla94@gmail.com:RSLRlohith94#
setuakhoury@gmail.com:oku@3122
setuakhoury@gmail.com:oku@3122
vijendrasingh2111@gmail.com:vijendr@1
sherestha@gmail.com:Bluejay14*
rxtrived@gmail.com:kriti131
sankar.rgb@gmail.com:sarovar@1
nidhi.kh@gmail.com:direction@123
sweetlilgal_in@yahoo.co.in:Dhruvi@21
pradeepm1406@gmail.com:latha1406
raokishore@yahoo.com:Neelima30
siddhu8531@gmail.com:150885pintoo
neethugeetha10@gmail.com:neethu@123#
sachinshiromany@gmail.com:care0312
sai.ravipudi@gmail.com:ramsai999
venkat.avula221@gmail.com:venky@1989
prins.patel441@gmail.com:dhanu1416
shavish18@gmail.com:shavish1870
shubra1977@gmail.com:shanaya2206
ashviniaesd@gmail.com:Explorer!2
shubra1977@gmail.com:shanaya2206
textmadi@gmail.com:Hanuman@@123
shaikpasha0108@gmail.com:borabanda123
parixit87@gmail.com:0501206087
masterpratyush08@gmail.com:Pratyush100
sunilk9@gmail.com:au703304
sujithredrocks@gmail.com:Sujith@12
parimalseema@gmail.com:parimalseema1990
pal.othersites@gmail.com:palothersites
pagarrohan12345@gmail.com:rdx_0807
surajshetty303@gmail.com:Suraj@303
vartichauhan13@gmail.com:Vartika13@
supriya.chalagam@gmail.com:vamsi921
tamboli1912@gmail.com:tamboli@1912
phanindranemala@gmail.com:annapurna@22
KONETICHANDRAREDDY@GMAIL.COM:ammananna@123
pratichi.mishra197@gmail.com:Adishree2357
alagappan2110@gmail.com:visu13071953
santoshgade@rediffmail.com:raghava@7
ramji_52@yahoo.co.in:ramjiansc115
sameen.neha@gmail.com:Neha@1108
pgrover005@gmail.com:p@rth005
richiee21@live.com:$Godisgreat1
mrunalpalande059@gmail.com:mrunal1212
sahejapo@gmail.com:03june1999
stynemcloy@gmail.com:rayalanagar
santhunaanu@gmail.com:BASAVANNA1
subh2601@gmail.com:subhash26*
neiwete@gmail.com:lomi2017
pridwik@gmail.com:lakshmiP1
poonamgr@gmail.com:poonam5757
sukrita.paulkumar@gmail.com:nairobi49
pinal83shah@gmail.com:pinal83shah
mrunalpalande059@gmail.com:mrunal1212
shazia.farhana@gmail.com:Farhanas3
shilpisinghal29@gmail.com:tanushilpi
rashitalwarbhatia@gmail.com:Aditya2506
amitmukherjii@gmail.com:timasonu
sanketshah11@gmail.com:Life@008
pratiksha00@gmail.com:248803123
thyaga81@gmail.com:thyaga81
sudhamshi.reddy@gmail.com:marchipoku7
roshan.anand222@gmail.com:roshan.anand222
sanjeevgahlot@yahoo.com:sanj7517
tiwarishivam@rediffmail.com:M@idenIr0n
chandana1985hitha@gmail.com:chandu85
rahulkamble1385@gmail.com:Dhamma358
alagappan2110@gmail.com:visu13071953
abhishek11.tripathi@gmail.com:gabbudada
badal.badala@gmail.com:udaipur321
ve.ramesh.babu@gmail.com:jagram1064
stellarcables@gmail.com:anand@79
poojajainfd@gmail.com:mishtii07
prithwishgiri2010@gmail.com:Dulu@123
rakesh.h1984@gmail.com:Ricky@1984
pragati.v.bora@gmail.com:Kanchan321#
shishubalak6@gmail.com:Anushi@123
sameen.neha@gmail.com:Neha@1108
seshaiahashok@gmail.com:ashok@1809
bhawanachampawat044@gmail.com:anand6014
raighanshyam97@gmail.com:Ghannu@54321
sankar.rgb@gmail.com:sarovar@1
pragyesh07cs31@gmail.com:pkm12345@
stalinpaul14.sp@gmail.com:Stal@1411
aayushishah756@gmail.com:KAJALmanish
sunischalreddy@yahoo.com:shikaa99
sekhar301@gmail.com:Pallavi@301
rajnish781@yahoo.co.in:madhur.01
sbhargavi.ec@gmail.com:sweety89#
raveenaacharya@gmail.com:avani2609
sandeepterala.59@gmail.com:sandeepkumar
prachiranade2@gmail.com:Prachi1987#
nitin260659@gmail.com:bharti2493
gupta.ashima8@gmail.com:9086110108
tuhinabanerjee195@gmail.com:Anu12jan
varma874@gmail.com:ajvarma7
gbyhospital.mohan@gmail.com:ka02eb52
KONETICHANDRAREDDY@GMAIL.COM:ammananna@123
praveenkmr78@gmail.com:pkumar@77
reachkumar16@gmail.com:Men@w0rk
narayanptripathy@gmail.com:bablu#1991
saipowerstar25@gmail.com:Srisai@25
sapnarenu11@gmail.com:drbhagwan
siddhu8531@gmail.com:150885pintoo
rani.p.july01@gmail.com:143ullas
mohit.bhtngr@gmail.com:mani2515
tashusrivastav1993@gmail.com:rebornpart12
sivalingakumar140@gmail.com:siva@sai143
supriya.chalagam@gmail.com:vamsi921
tanwirarshed@gmail.com:03532699173
spdhanushya@gmail.com:Dhanushya99
shopkeny@gmail.com:fun@shop
nuthan.kotian@gmail.com:janu0422
archsa2801@gmail.com:ghar1993
senthazai.s@gmail.com:flower.143
sambmish@gmail.com:#Bishop@123
surajshetty303@gmail.com:Suraj@303
shanky.verma02@gmail.com:shanky@1989
syedaurooj2305@gmail.com:Urooj2305
sahusumita@gmail.com:sumi1241
tejagokina@gmail.com:tejagokina
vasuvinothkumar@gmail.com:jkem12fr
poojita99@gmail.com:poojita99
swetha.tall@gmail.com:Sarayu@31
ravishpandey1989@gmail.com:ravish1989
sethisrishti11@gmail.com:srishtisethi11
savchelsea@gmail.com:Allscripts@16
tambe.ganesh@gmail.com:pranav*01
poojaa9.92@gmail.com:Viji@1990
pal.othersites@gmail.com:palothersites
subhashanker@gmail.com:subhashini9*
jayasinha020908@gmail.com:NewPass@1
jphosp@gmail.com:tatroo80
tamil293@gmail.com:deepsel
samitha095@gmail.com:S@mitha095
paridhig13@gmail.com:omsai401
stalinpaul14.sp@gmail.com:Stal@1411
shavish18@gmail.com:shavish1870
ranjithkp625@gmail.com:Ranjith@45
psrireddy9@gmail.com:Mukyaprana@9
sanjaybaranwal@gmail.com:Aligarh1@
nowaj.sharif@gmail.com:@raveena
salil.info@gmail.com:P@ssion02
prasshanthshettyy@gmail.com:Nisha@1110
suneeldr@gmail.com:chagsun
vinayakahk@yahoo.com:pulsar100
mustak_saiyad26@yahoo.com:TINU@9786
vartichauhan13@gmail.com:Vartika13@
dusane.suhas@gmail.com:Swara@1409
nikhildchaudhari@gmail.com:Nikhil@2807
ritayan.hello@gmail.com:sKv2CdHT$
neiwete@gmail.com:lomi2017
sehkar1980@gmail.com:Sekar_1980
narendra.devrani@gmail.com:Narendra*12
carekhasoni@gmail.com:canro341304
priti.nahar@ymail.com:S@unshine8
sandyjaan1986@gmail.com:Sandy@3131
rajib_biswas80@yahoo.co.in:Butter1$#
suhas.barmase@gmail.com:arvind@1312
srjtmitra5@gmail.com:srijit87
sivalingakumar140@gmail.com:siva@sai143
harneet@pm.me:shobhade7
vamsikrishnakola@gmail.com:japanise2
neethugeetha10@gmail.com:neethu@123#
rajkumarampolu@gmail.com:rajkiran37
supathra@rediffmail.com:Supa1hra
padmabhaskaram71@gmail.com:udbhav@123
yadav.sunil1351@gmail.com:Anilkumar@1995
rampelli.pradeep@gmail.com:Rampelli@123
nitin260659@gmail.com:bharti2493
nsd.prasad@gmail.com:King@208
skams1994@gmail.com:kamakshi1994
c4@2die4.com:dOors1024
vg249089@icloud.com:Babita0812
dhuruvsharma@gmail.com:Bittoo912$
puneet.chaturvedi@gmail.com:aditi1942
priyanka.v1010@gmail.com:Deeptrisha@2008
shrutig.naik@gmail.com:Rishi0210#
nimmikr@gmail.com:Manya-123
sabaresh.p@gmail.com:Logica12345
sanvb@rediffmail.com:sanvaish9
priyatyagidocs@gmail.com:Priya@2212
abhishek.tiwari2388@gmail.com:Abhi2326@
pani.dasyam@gmail.com:chakri9999
rocksays1911@gmail.com:nagamani@16
tejaswani.mamidi7@gmail.com:LondoN07
paridhig13@gmail.com:omsai401
naveen789sh@gmail.com:Naveen@405
roopdeb@gmail.com:Honda1042
vishu.9896@gmail.com:2Jangurpilov
pathansalman886@gmail.com:Samzishu89
shru.nanjunda@gmail.com:Feb@1991
vadav.06@gmail.com:vaibhavyadav18
rakeshbabu@gmail.com:Bodla123
tabassumkhan02@gmail.com:reyman@786
sukrita.paulkumar@gmail.com:nairobi49
poojita99@gmail.com:poojita99
sandeshkd9@gmail.com:Sandesh@1989
saugata.1980@gmail.com:Mosaboni5",
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	pl("Hello! Welcome to my little Go game.\nIt's still unfinished, so if you find any bugs, please report them to cmyui#5585 on discord.\nThank you <3\n\nThe game will start in 5 seconds..")
	sleeps(5)
	CallClear()
	pl("You're sitting on an old park bench in Gage Park in Brampton, Ontario.. Wearing a heavy jacket, some jeans and boots, you\ngot bored of walking around and enjoying the scenery. You take out your phone and begin watching some YouTube videos..\nA man in a hoodie and slick black sweatpants approaches you..")
	td()
	sleeps(5)
	pl("??: Hey, do you happen to know what time it is?")
	sleeps(2)
	pl("Me: Oh yea, its uhh.. 12:35")
	sleeps(2)
	pl("??: Awesome, thanks")
	sleeps(2)
	pl("He sits down beside you.. He's quite sweaty and looks like he just finished his run..")
	td()
	sleeps(2)
	pl("??: Oh, sorry for not introducing myself. I'm Josh..")
	pl("He extends his hand, and it takes me a second to react")
	p("Me: Oh, right, sorry, I'm: ")
	var name string
	var nameLower string
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	name = string(bytes)

	nameLower = strings.ToLower(name)
	if nameLower == "josh" {
		td()
		CallClear()
		sleeps(1)
		pl("F͍̖̣̿̃ͅͅͅÙ̖C̜̯̤̗̓ͤ̉͆ͅͅK̝͓ͪͮ ͖̓ͧ̋O͓̺̟̜̦̽ͣ͗́̀̚F̥̘͖ͯͩ̀̑͑̓̚F̝͑͑ͪ̇̊͌ͥ ̰̇A̫͍̠͙͚̟N̳͓̟͖̎̊ͅD͇̼̠̱̱̩̙̈̓ͣ̏ ̪̗ͫ̅Nͣ̌ͬͬ̔̈E͕̱͚̗̠̠ͨ͑̀̉̌V̫͖̲̟̹͖E̬͕͕͈R͎͉ ̤̹̙̲R̪̳ͯ͛̃ͦ̎ͫ̌E̖̝͈̱͈̫͓͑ͪ̓ͬͨ̊T̬͓͖͓̠̗͎̔̃U̖ͧ̄̔R̤̰̼͎͆̇̋ͧ̒̿ͯNͮ̉")
		sleepm(500)
		end()
	}
	pl(name + ": Nice to meet you too!")
	sleeps(1)
	td()
	pl("Josh: Mm.. Sorry if this seems out of context, but have you ever heard of a game called osu?")
	sleepm(600)
	var playOsuString string
	p(name + ": ")
	pos, _, _ := reader.ReadLine()
	playOsuString = string(pos)
	playOsuString = strings.ToLower(playOsuString)

	if playOsuString[0] == 'y' {
		sleeps(1)
		pl("Josh: Wait.. Seriously? ")
		sleeps(2)
		pl("Josh: Do you actually play the game?")
		sleepm(600)
		var playOsu string
		p(name + ": ")
		po, _, _ := reader.ReadLine()
		playOsu = string(po)
		playOsu = strings.ToLower(playOsu)

		if playOsu[0] == 'y' {
			sleeps(2)
			pl("Josh: Oh! Holy shit! You're the first person I've met that actually plays the game!")
			sleeps(2)
			pl(name + ": Haha, yea, not a lot of people know about it, and even less play it.")
		} else {
			sleeps(2)
			pl("Josh: Aw.. Damn.. I really wante to meet someone else that played it..")
			sleeps(3)
			pl("Josh: Anyways.. I've got to head home. See you around!")
			sleepm(1500)
			pl(name + ": Yea dude, have a good one!")
			sleeps(1)
			softEnd()
		}
	} else {
		sleeps(2)
		pl("Josh: Aw.. Damn.. I really wanted to meet someone else that played it..")
		sleeps(2)
		pl("Josh: Anyways.. I've got to head home. See you around!")
		sleepm(1500)
		pl(name + ": Yea dude, have a good one!")
		sleeps(1)
		softEnd()
	}
	sleepm(1500)
	pl("Josh: So, since we both play osu!.. Who's your favourite player?")
	sleepm(600)
	p(name + ": ")
	var favourite string
	var kill bool
	fa, _, _ := reader.ReadLine()
	favourite = string(fa)
	favourite = strings.ToLower(favourite)
	sleeps(1)
	switch favourite {
	case "kablaze":
		kill = false
		sleeps(1)
		td()
		pl("Josh: Wait.. what?")
		sleeps(2)
		pl("Josh: Really..?")
		sleeps(2)
		pl("Josh: Huh.. He's actually probably my favourite aswell..")
		sleeps(2)
		pl("Josh: I think he's a really unrated and undervalued player overall..")
		sleeps(3)
		pl("Josh: Well, thats sortof his fault.. He doesn't really put himself 'out there' very much..")
		sleeps(2)
	case "cookiezi":
		kill = false
		pl("Josh: Haha. Yea, I definitely can't blame you for that one. I definitely think hes the BEST player,")
		sleeps(1)
		pl("Josh: Although, he's not my favourite..")
	case "howl":
		kill = false
		pl("Josh: :^)")
	case "solis":
		kill = false
		pl("Josh: oh yeah dude solis is my favourite relax cheater too :^)")
	case "cmyui":
	case "chase":
		kill = false
		pl("Josh: <3")
	case "rafis", "azerite", "zirba", "bubbleman",
		"rohulk", "angelsim", "emilia", "woey", "osu player84", "mathi",
		"totoki", "_ryuk", "ryuk", "mcy3", "spare", "danyl", "piggey",
		"bro_gamer72", "gayz", "gayzmcgee", "dustice", "idke", "happystick",
		"_index", "index", "karthy", "yaong", "digitalhypno", "toy", "aireu",
		"wubwoofwolf", "talala", "resia", "adamqs", "wilchq", "mouseeasy",
		"flyingtuna", "thepoon", "recia", "dumii", "mlaw22", "mlaw", "mshake",
		"xilver", "ceptin", "plz enjoy game", "apraxia", "doomsday", "vaxei",
		"filsdelama", "dustice":
		kill = false
		pl("Josh: Yea, I can't blame you for that. He's extremely talented. Although..")
		sleeps(1)
		pl("Josh: He's not my favourite, but he's definitely up there.")
	case "rrtyui":
		kill = false
		pl("Josh: Meh.. Irrelevant trash player, but i'll let it slide..")
	case "hvick", "hvick225":
		kill = false
		pl("Josh: Yea, I can't blame you for that. She's extremely talented. Although..")
		sleeps(1)
		pl("Josh: She's not my favourite, but she's definitely up there.")
	default:
		kill = true
		sleeps(1)
		pl("Josh: ..")
		sleeps(2)
		pl("Josh: Who..?")
		sleeps(1)
		pl("Josh: Oh well.. whatever.. It's not like I cared anyways")
		sleeps(4)
	}
	if kill == true {
		td()
		pl("Josh: Actually what the fuck? How do you not appreciate any good players?")
		sleeps(3)
		pl("I literally have no idea who you're even talking about. What the hell is wrong with you??")
		sleeps(3)
		pl("Josh: Are you mentally insane? I don't understand your thinking")
		sleeps(3)
		pl("Johs: Yea.. You're probably insane. You should go get that checked out..")
		sleeps(3)
		pl("josh: Dude go to the hospital already what the hell are you doing")
		sleeps(3)
		pl("Josh: You're clearly stupid and delusional if you don't think kablaze, cmyui, or cookiezi is the best player.")
		sleeps(2)
		pl("J̢̙̮̙̮̬͍̰̺̗͚̥̠͉̹̙͕̹ͅͅoŞ̧̡̨͍̘̖̭̭̗͕̫̭͉̰̥͎͈̻͈h͙̖͇̘̱̦̩̱̞̝̞̬̯̣̹̞̰͉̮̯:͇̱̬͎̣ ̢̰͕̟̪̹̯͍̥͍͍̤̺̫̱̟̟͚̰͜W̡̡͕̺̬̠Ḩ͍̬̠͖̲̝̼̗͙̤̞̤̗̹̰̩̣͜͜Y̧̧̢̧̡̧̹̲̖̺̩̲̞͙̹̬̭̦̣ͅ ̡̢̫̺̥̦̘̻͔̦͔̙̣̝͔͓̫̭̤ͅA̢̡̧̡̼̰̬̗͚͔̠̜̺̲̝̯͕̫͓̭�̧̨̮̲͈̭͚̗̗̗̠̺̱̳͙̣͇̼͖̳�̡̧̧̨̙͉͇̲̗̣̹͙͇̤̜̣̖̟̩̠�̢̠͉̯͇̰͓͔̞̯̺̥̝͉̤͍̳̞͔̮R̺̭͇̼͖Ȩ̞̮̭̩͇̟̥̥̠̖̟̲̮̖̰͔͜ͅO̢̢̫͍̲͇̘̦̮̘͙̙̼̳͙̤͙̞͕̹U̡̥͇͔̞̝̝̩͇͍͔̰̭̜̯̫̯͇͜͜�̢̧̢̧͎̻̘͚̣̻͎̮͇̟̖̺̯͖̭̟  ̨̻̗̭̭͇̘̼̹͍͙͖̬̤̞̼͕ͅE̢̠͔̬̳̙͖̰̭̺̪̗̠̲̘͚̞̗̭͜Ã͕̦̯̜̺͉͎͇̜̟͔͎̬̞̘̹̞̘̥ͅ³̢̯̘̼̯̤̭̬̳̟̫̲͓ķ̫̫̞͈̳̜̦̳̤̺͕̪̥̲̟̰̹͜ͅǫ̡̙̣̯̤̘͔͙̙͎̹̪̞̻̟̬̣͙͜r͚̥̙̘̪ͅȨ̮̥͇̤͉͍̦̞̹͉̞̹̳̤̣̫̜̻̬?̢̝̬̹̣̝͓͍̹̱̺̩̪̺̤̭̗͉̩̞?̨̡̙̮͖̞̟͙̺͖̩̭̻̯͓̖͉̺̦͜")
		sleepm(800)
		pl("jalapeÃ±o: L E A V E")
		sleepm(800)
		count := 1
		for count < 20 {
			pl("\n� /? / ? LFKS�?fsajj �")
			count++
		}
		end()
	}

	sleeps(2)
	td()
	pl("Josh: To be honest, i'm amazed you've made it this far..")
	td()
	sleeps(2)
	pl(name + ": What?")
	sleeps(2)
	pl("Josh: Nothing.. It's almost like..")
	sleeps(2)
	pl("Josh: You're trying to read my mind..")
	pl(name + ": Uh.. Sorry I'm not sure I know what you're talking about..")
	sleeps(2)
	pl("Josh: Heh... Of course not..")
	sleeps(4)
	td()
	pl("Josh: Um.. Have you ever heard of Akatsuki then?")
	sleepm(300)
	p(name + ": ")
	var akatsuki string
	a, _, _ := reader.ReadLine()
	akatsuki = string(a)
	akatsuki = strings.ToLower(akatsuki)

	if akatsuki[0] == 'y' {
		sleeps(2)
		pl("Josh: Yea.. It's been getting quite a lot of attention recently..")
		sleeps(2)
		pl("I'm actually the creator of it..")
		sleeps(2)
		pl(name + ": Oh dude, thats kind of cool.. although.. isnt it litrally just ripples code..?")
		dotspam()
		sleeps(2)
		pl("Josh: �.")
		sleeps(3)
		pl("Jos�̡̢̡̝̥͇̠̗͖̙̹̥̩̣̻̞͎̜̹͇̻͚̖̭̖̟̼ͅh̨̢̨̨̭̞̖̲̝̭̠̜̳͔̱̫̞̼̠̖͎͎̠̰̦̦̬̭:̡̺̰̲̣̲̼̲̖͔̗̞̲͇͈̖̠̯̜̲̮̦̼̹̻̠̱ͅ ̡̧̢̢̟̩͈̞̳̙͔͔̜͓̼̥̭̝̘͔͓̝̻̪̯̯̹ͅI̧̻̯͍̜͓͈̹͖̰̬̪̹͇̱̝̣̥̠͙̣͉̲̠̺̠̲ͅ'̨̧̢̢̡̨̡̙̟̙̳̥͓͎̼̩̰̠̟̞̗͍̺͙̪̬̝ͅm̢̡̧̨̩̜̠̝̱͕̲̦͇̰̬͎͙̠̱̼̹̯̥̘̰͈̥̣ ̡̡̹͓̥̙̣̙͈͈̘̠̣̰̞͈͎̲̩̣̼͎̝̳̣͍͜ͅs̢̢̫̜̘͖̞̫̩̩̟̻͔̻̙̼̜͓̭̜͔̬͕̣͜͜ͅͅo̡̢̨̲̼͉͇̟͖̯̰͉̦͍̻͓̜̻̫͇͍͖̮̰͍̯͚͜ŗ̥͉̜̩̰͉̩̼͖̮̱̻͍̩͎̮͔͕̲͈̹͇̦̪ͅͅ    rRrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr-")
		end()
	} else {
		sleeps(2)
		pl("Josh: Oh.. Yea it's an osu! private server I made using Ripple's source as a base..")
		sleeps(2)
		pl("Josh: It sortof gets a lot of hate due to the fact that most people just see it as a ripple clone..")
		sleeps(2)
		pl("Which I actually understand, and even agree with.. but I wish that people wouldn't spread so much hate about it..")
		sleeps(2)
		pl("Oh well.. This isn't the place to vent about that kind of stuff anyways..")
		sleeps(2)
	}

	pl("Have you done anything interesting recently?")
	sleepm(600)
	p(name + ": ")
	var interesting string
	in, _, _ := reader.ReadLine()
	interesting = string(in)
	interesting = strings.ToLower(interesting)

	if interesting[0] == 'y' {
		sleepm(900)
		pl("Oh dude, tell me about it!")
		sleepm(600)
		print(name + ": ")
		var blabla string
		bb, _, _ := reader.ReadLine()
		blabla = string(bb)
		blabla = strings.ToLower(blabla)

		if blabla[0] == 'n' {
			sleepm(900)
			pl("ẅ̮̙̪͚̭̪̳͓̩́̐̈̎͊̌̃̊̀͗ͅĕ̮͉̖̫̣̱̦̰̜̽̑̿́͆̿̈͂͠ͅl̪͓̘̲̙͓̠̱̟̎͑̎̌̿͑̎͑̇̕ͅl̡̨͎̭͔̖̹̺̻̗̅̑̉̌̂́̄̒̎̕ ̢̗̞̻̟͓̘͖͎̲͆̒͌́͆̎͛̂͆͝f̹͙̗̲̝̮͎̮̜͎̆̎̉͌̊̈̃̍̎̚u̡̜̩̗̺̺̱̺͖̗̔̍̈̓̈́͒̎͛͝͝ç̬̪̟̟̤͙̹̟͔̽̋̀̈́͛̾͋̿̈͝k̦̥̩̬̞̗̟̳͎͇͒̓͗̊̾̿̒̈́̆̊ ̢̜͚̳̙̖͉͇͇͂̑̂̇̄͆̌̉̈́͘͜y̹̰͍͍̙̲̰̞̳͉͒̐̓̆̉̓̆̆̚͝o̞̫̻̗̖͎̬̝̝̽̐̈̆̔͑̾̎͘͝ͅu̠̞͕̮͎̫̦̻͈͔̐̆͊͋͛͛̎̔̏͝ ̨͙̳̝̭͚̰̲̰̘̋̾͐̌̅̈́̈̃͆͘ẗ̛̜̝̠̖̪̬̹̗͍̺͂͆͛́͂̒͊͘ơ̧̹͚͇̯͎͕̳̘͍͌̓̑̉̐͐̈́̊̋ǫ̧̡̟̟͇̥̯̠̱̈́͗̌͗̈́̒̍̉͝͠")
			end()
		}
		sleepm(1500)
		pl("Wow.. That's super cool! I hope I can do something like that someday..")
		sleeps(1)
	} else {
		sleepm(600)
		pl("Josh: Oh.. Well that kinda sucks.. I haven't done much either I guess..")
		sleeps(2)
		pl("Josh: Well.. At least other th�n the osu! server. But that wasn't even me really. I really owe it to Howl and Nyo.")
		sleeps(2)
		pl(name + ": Well.. Why not do something for them then..?")
		sleeps(2)
		pl("mmm..")
		sleeps(2)
		pl("Josh: I'm really now sure what I could even do.. One of their developers doesnt like me at all\nand I'm not sure what the other one thinks..")
		sleeps(4)
		pl(name + ": Oh well Sucks lole")
		end()
	}

	pl(name + ": I'm sure you'll do alot more in your life. haha")
	sleeps(2)
	pl("Josh: Yup.. Most likely. I'm just waiting for the day where I come out of my shell.")
	sleeps(2)
	pl("Josh: So.. Out of curiousity, are you any good at osu..? What's your rank?")
	sleepm(600)
	p(name + ": ")
	var rank int
	var digit int
	fmt.Scanln(&rank)
	if rank >= 100000 {
		digit = 6
		sleepm(400)
		pl("Josh: Oh, so you're still pretty new to the game then!")
		sleepm(1200)
		pl("Josh: I remember when I was that rank.. I first got my tablet right around 250k, haha")
	} else if rank > 10000 {
		digit = 5
		sleepm(400)
		pl("Josh: Oh. So you've been playing for a bit, I guess. Still pretty new though, haha.")
		sleepm(1200)
		pl("Josh: You'll probably start to see a slowdown in progress around the rank you're at now though. That's when it starts to actually get difficult..")
	} else if rank >= 3000 {
		digit = 4
		sleepm(400)
		pl("Josh: Oh damn.. You're catching up to me.. I got banned at about 2.9k, but that was after 2 cheated plays.. In reality i'm probably around there now.")
		sleepm(1200)
		pl("Josh: Best of lucking kicking my ass, haha.")
	} else if rank >= 1000 {
		digit = 4
		sleepm(400)
		pl("Josh: Oh damn, you're a bit further than I am then..")
		sleepm(1200)
		pl("Josh: We might actually be able to compete, I'll be sure to try to snipe some of your scores, haha")
	} else if rank >= 100 {
		digit = 3
		sleepm(400)
		pl("Josh: Oh wow.. You're really good damn. I hope I'll reach that sort of level one day..")
		sleepm(1200)
		pl("Josh: Not sure if I can, but I think it's a good goal.")
	} else if rank >= 35 {
		sleepm(400)
		pl("Jo�h: ?  ????????? ? ?       ????????????  ?  ? ? ???\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? �?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? �?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? ?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? ?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? ?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? ?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?")
		end()
	} else if rank >= 20 {
		digit = 2
		sleepm(400)
		pl("Josh: Bro im finna smash ur ass like grass at ANIME NORTH CON 2018 :^)")
	} else {
		sleepm(400)
		var count int
		count = 0
		for count < 20 {
			pl("\nJo̧̡̧̨̨̧̧̢͓͓̼̘̝͇͈͍̼͓̦̟̝͉͍͍̬͈͔̜̥̜̱͕̲̲̗̥̖̻̦̠̠̹̣̣͖̞̺̻͖̤̪̭̳͖͈͚̦̳͔̥̰̜̪̩̱̗͚̞̯̘̮͕͔͙͍̭͙̳̺̱̥̣̗͖͙͎̥̣̭̞͇̦̠̘̥͜ͅͅͅͅ࿊hs�h: loO  Le h�e SDlo i am   coo   kiez i xDD℻DS�႟DDSD��᥀᥀℆℆℆D")
			count++
		}
		end()
	}
	sleeps(1)
	pl("Josh: Anyways, I've got to go home.. Do you use discord by any chance?")
	sleepm(600)
	p(name + ": ")
	var haveDiscord string
	hd, _, _ := reader.ReadLine()
	haveDiscord = string(hd)
	haveDiscord = strings.ToLower(haveDiscord)

	if haveDiscord[0] == 'y' {
		sleepm(600)
		pl("Josh: Oh, sweet! What's your ID? Mine is cmyui#5585")
		sleepm(300)
	} else {
		sleepm(600)
		pl("Josh: Oh damn.. Well.. Maybe we'll meet again one day, haha.")
		sleeps(1)
		pl(name + ": Maybe.. Well.. See ya!")
		sleeps(1)
		pl("Josh: Cya!")
		end()
	}
	var discord string
	p(name + "Yea sure, mines: ")
	d, _, _ := reader.ReadLine()
	discord = string(d)
	discord = strings.ToLower(discord)
	switch discord {
	case "cmyui#5585":
		pl("J̢̢̧̨̨͓̱̹̝̣̳̥͎̯̫̣̭͇̗̮͚̼̬͈͕̫̲̩̮̠̖̗̺̘͉̝͍͙͇̼̥͖͙͇̖̬̤̘̭̠͜͜͜ͅS̢̧͚̰͔̰̮̞͕͇͉̪̝̼̟̜̠̟̻̬͎̠̲̘̭̹̗͓̺̱̹̤͎̤̲̣̳̯̺̖̪͜ͅH̨̠̫̤̹̞͙͉͈̻͓̬̩̲̖̜̰̦̥̬̙͚̥͙̥̹̠̣̝̖͖̰̜͎̮͜ͅͅ℆̨̨̢̨̢̨̨̨͙͙̘͓͖̜̹̯̥̪͍̫̗̝̬͖̬̞̰͖̭̝̮̫̤̫̦̩͙̝̞̟͇͉̟̹̤̙̞̺̖̤̳̲̣̟̹͍̦͙̫͚̣̜͕̜͕͔̙̤̥̻̝͓̝͓̙̞̹̺͇̠̟̬̬͎̰͖͙̝̫̫͎̟͓̩͍͕͜͜ͅͅS̢̡̡̨̧̢̡͈̜̩͓͍̰͚̪̰͓̲̦̺̦̬̜̝͍̙̜͕̥̟̼͈̣̟̥̼̤̤̤̮͎̱̻̻͈̹͎̘̺͉̞̯͎͈̘̹͔̼̭̳̖͈̞͇͉͓̹̣͔̭̪͔̠̞̫̯̥̼͉͍͉͍͜͜͜ͅͅ ̨̡̧̡̣͙̳͓͖͙͈̭̤͖͍̘͈̠͕̤͈͎̙̦̩̯̜͔͙̜̦͖̻̪͙̙̲̟̱̝̝͜ͅJ̡̢̡̧̨̨̨͉̩̥͙̞̫̥̥̜̼̩̤̮̣̳̥̰̖͍̲̮̗̖̭͈̗̻̮͍̼̯̰̣̪͔̱̤̗̟̝̝͔̖̻̣̼̹͚̰̬̻̪̣̳̥̜̗͔̭̠̲͙̠̥̟͖̟͈̗̘͍̰̰͎̦͚̤̼͓̹͈͜͜ͅH̨̧̡̨̡̧̨̡̢̧̢̟͉͓̤̞̭̥̘̦̜͓͉̘̣͖͕̥̜̣͉̟̣̣̞͔͉̳̠͓͍̦̘̳̬͙̟̣͍̩͎̯͉̹͚̗͖̦͇͖͖̻̩͍͕̹̪̲̳͍̭̟̬̖͍͕̠͍̲̹͕̪̳̺̭̟̱͕͖̙̯̙͔̲̫͜ͅͅͅͅ ̧̧̡̢̢̧̧̡̡̨̨̡̧͉͕͚̘̞̭̻̫̺͔̠̜̝̭̟̬͓̝̘̯̙̬̰̞̯̫͔̼̼͈̰͖͍̗̯̳̗̮̗̣̹̗͉̻͕̟̼̦͔͇̩̝̥̹̖̰̥͎͕̮̗̜̭̬̻̹̙̖͇͔̱͙̙̗̮̜̜̗̰̟̥͕͜͜͜ͅͅą̨̨̧̨̢̡̧̨̡̧͈̼̹̟̭̣͈̟̰̣͔̘̲̘͚̫̥͈̲͓͈̝̺̳̪͕̗̻͈̱̠̩͚͙͍̜̦͓͍̜̝̹̙̺͕̺̻̝͔̺͕͚̤̟͖̪͉ͅͅͅy̡̢̡̢̢̧̢̢̢̢̨̻̣̜̫͈̥̗̭̱̼̗̪̯̹͓̝͇̙͇̞͈͓̘̣͚͔͍̺̲̳̦̤̖͔̼͓̝̣̩͖̬̜̹͕̺̜̬͉̺̙̰̘̟̫̦͓̩̦̦̻̮̖̗͔͈̠̺͓̹̥̦̫͙͇̯̟͔̥̥̺̞͚̜̘̲̝̖̘͕̣̯̻̳̻͇͎̭͎̣̳͚͎͜͜͜͜͜͜͜͜͜ͅͅͅͅͅb̨̡̢̨̡̢̢̡̨̡̢̡̥̦̫̫̤͉̤͓͕̲̩̭̠͖͍̹̝͙̹̥̱̟̹͍̙̻̥̻̳̠̟̙͕̮̮̪͈͔̰̰̮͍̪̻͕̗̱̠͈̤̪͇̳͔̲̭̗̜͖̬͕̝̤̳͎̝͙̩͚͜ͅͅF̢̡̧̡̧̧̣̪̜͙͖̱̺̞͙͙̼̫̲̞̭͔̳̫̦̲̰̹̥̟̹͓͈̜̝͉͚͎̫͕̖̟̗͖͓̦̙̜͓̝̥͖̺͉̱͕̘̭͓̬̙̠̠̯͜͜ͅy̧̢̺̲͔̠̮̬̬͇̦̼͖̤͙̗̻͇̬̞̲͙̮̳̼̙̠̭̫̮̱̳͓͉̜̜̮̣̣̞̘͓͎ͅͅ\n̢̢̡̡̡̟̟̰̪̩͇̝̻͔̪̖̲͔̮̰̞̫͕̙͚̮̹͕͈̩̼͕͚̗̗͚͖͚̼̜͉͔͓̺̣̬̜̱͇̤̙̲̖̰̳̼̮͔̱̹̬̱̥̰̞̥̰̻̰̣̟͖͚̰̺̜̜̟͎̼̯̹̙͜ǫ̨̨̨̢̡̡̨̧̢̧̟̙̰̘͎͈̮̟͕̜̱̥̞͕̙͖͉̬̞͎̟̣̫̱̼̘̭̝̙̲̗̭̠͍̹͙̬͔̭̰̖̯̞̩̩͇̲̜͈͔̪̮͎̮̬̱͖̱͓͕̪̜̱̪̖͈͙̣̜͍̣̰͈̱̼͓̹̩̙͜͜͜͜ͅͅͅͅͅͅ.")
		end()
		break
	case "kahri#9828", "Emily | Sunpy#5213", "Hanyeol 한열#8257", "Kip#1120", "Frosti#0602", "ninjin#3198", "Howl#0940":
		pl("Josh: Huh.. That sounds oddly familliar..")
		sleeps(2)
		pl("Josh: well, whatever.")
		break
	default:
		pl("Josh: Alright cool!")
		break
	}
	sleeps(1)
	pl("Josh: I'll make sure to add you when I get on!")
	sleeps(2)
	pl("Josh heads off towards his home, and I do the same..")
	sleeps(1)
	pl("..")
	sleeps(1)
	pl("..")
	sleeps(1)
	pl("..")
	sleeps(1)
	pl("..")
	sleeps(1)
	pl("..")
	sleeps(1)
	pl("I open my door, take a seat and turn on my PC.")
	sleeps(1)
	pl("Mmm.. It's been a long day..")
	sleeps(2)
	pl("Discord opens up, and I head to the Friends section, and see that cmyui#5585 has already added me as a friend.")
	sleeps(2)
	pl("I accept the request, and send him a PM saying 'Hey, long time no see, haha'")
	sleeps(3)
	pl("5 minutes passes, and still no reply.. I assume he must be busy doing something, so I open up osu!")
	sleeps(2)
	switch digit {
	case 6:
		pl("I've really been improving recently, even though im only a 6 digit, I think that I might be able to climb again now..")
	case 5:
		pl("I've really been improving recently, even though im only a 5 digit, I think that I might be able to climb again now..")
	case 4:
		pl("I've really been improving recently, even though im a 4 digit, I think that I might be able to climb again now..")
	case 3:
		pl("I've really been improving recently, even though im already decent.. I think that I might be able to climb again now..")
	case 2:
		pl("I've really been improving recently, even though im already so high in the rankings.. I think that I might be able to climb again now..")
	}
	sleeps(2)
	pl("osu! boots up, and I realize that I never actually got to know how good Josh was..")
	sleeps(2)
	pl("I hit F9 and search up his name, and sure enough, there he is. A 4 digit, rank #3482.")
	sleeps(2)
	switch digit {
	case 6:
		pl("Wow, what the hell? He's really good!")
	case 5:
		pl("Wow.. That's really good..")
	case 4:
		pl("Oh, so we're around the same skill. Neat.")
	case 3:
		pl("Oh, he's a bit worse than I am, haha.")
	case 2:
		pl("Oh, he's decent I guess..")
	}
	sleeps(2)
	//Oh maybe I'll make this something identity theft something something idk im winging this while writing it.
	pl("Hes playing a song, I guess that's why he didnt reply on discord.. Although that song is only a minute long.. weird..")
	sleeps(2)
	pl("I send him a PM ingame")
	sleeps(2)
	pl(name + ": Hey!")
	sleeps(2)
	pl("cmyui: uh.. do i know you?")
	sleeps(2)
	pl(name + ": Yea! I'm the guy you met today, haha.")
	sleeps(2)
	pl("cmyui: huh?")
	sleeps(2)
	pl("uh.. what do i say.. the guy from irl?")
	sleeps(2)
	pl(name + ": The guy you met irl like 20 minutes ago or so. I decided i'd just message you on here since you didn't reply on discord")
	sleeps(2)
	pl("cmyui: uh.. i think you might have the wrong person, i haven't left my room today lmao")
	sleeps(4)
	pl("...")
	sleeps(1)
	pl("Huh? how is this not him?")
	sleepm(1500)
	pl("cmyui: Also.. the person on discord was me aswell, I think you might have the wrong username or something?")
	sleepm(500)
	pl("cmyui: actually wait no, because discord has a # ID, how'd you manage to find my osu name AND discord id wtf")
	sleeps(3)
	pl("... What? Did the guy give me someone elses ID or something..?")
	sleeps(1)
	pl(name + ": Uhh.. Maybe the guy gave me someone elses ID..? Do you know a guy named Josh from Brampton?")
	sleeps(2)
	pl("cmyui: yea.. me..")
	sleeps(1)
	pl(name + ": Huh?")
	sleeps(2)
	pl("cmyui: I'm a guy named Josh from brampton..")
	sleeps(5)
	pl(name + ": wtf.. I think someone might be acting like they're you then..")
	sleeps(3)
	pl(name + ": wait, then why did you add me on discord first? I came home and you had already requested me as a friend..")
	sleeps(5)
	pl("He doesn't reply..")
	sleeps(2)
	pl("I'm beyond confused at this point..")
	sleeps(2)
	realEnd()
}

var x int

func td() {
	sleepm(750)
	p(".")
	sleepm(750)
	p(".")
	sleepm(750)
	p(".\n")
}
func dotspam() {
	p(".")
	sleepm(750)
	p(".")
	sleepm(750)
	p(".")
	sleepm(750)
	p(".")
	var count int
	count = 0
	for count < 100 {
		p(".")
		sleepm(5)
		count++
	}
}
func realEnd() {
	sleeps(3)
	CallClear()
	pl("Congratulations! You've reached the end of the game.. for now.. Feel free to contribute, or just tell cmyui you've reached this point,\nas theres only a few that have <3")
	sleeps(10)
	os.Exit(3)
}
func softEnd() {
	sleeps(3)
	CallClear()
	pl("Game Over")
	sleeps(10)
	os.Exit(3)
}
func end() {
	sleeps(3)
	CallClear()
	pl("An Error has Occurred")
	pl("Please check err.log for more information")
	file, err := os.Create("err.log")
	if err != nil {
		//If we can't write the file for any reason, just display it in console I guess
		log.Fatal("L̨̤͉̖̹̟̲͖͇̤̪͕̤̤̠̥ͨ̔ͮͤͥͭ͑ͧͭ̑̊̈́ͨ̿͐̚É̵̢͔̟͕̝̹͎͚͎͈̳̔̇̃ͩ͝A̵̯̜̬̖͙ͥ́ͪͦ͊̍ͨ͆V̡ͩ̔̾͗̚͟҉̭̝̱̣̥̥̗̱̪͓͕͈̯͙͈͖Ĕ̛̦̗̫̟̖̦̞̠̙̝͉̖̦̤̻̏ͬͫͯ́̂͋̃̑̀͞ ̵̮̠̠̂̈ͨ͗̑̅͞Ä͇̙͕̲̪̭͕̮͚͈̭͍̫̖̺͙͚́͆̑ͭ́̒̎͘͜͡ͅN̎͒ͭͨ̊ͤ͌ͤ͑ͮ̐͆͡҉̹̲̯͘D̷ͭ͋̌̒͋ͣ͒ͩ̾̃̌̽̀ͫͦ͞҉̯̳̫̺̳̞̹̖̜̖ ̷̢̡̱͉̲͔̻̙̳̜̳͑̿̍͗͌ͦͬͨ̂͋̉̌͠N̵̛̲̟̦̠̟̻̲̯̲̦̝͍̳̣̳̺̤ͪ͐͛͆ͣ͒ͤ̊̍͊̊̓ͥ̈ͬ͒̈́ͬ̀͘̕ͅȄ̶̺̠̞̹̱̯̘̘͖̥̗̔̏̑͐̃̅͗͊͗̌̄̈́̿ͪ̀̕͘V́͆ͮ̔͂̔̉̌̋͐̀͟͝͏̺̬̘͕͙̫̺̫̜̗̘̝̼̫̺̫̹̞͘E̟̻͉͓̣̰͓̮̫̎̐͊̂̆̉ͩ͛̌̐́̅͊̐ͨ̊̀ͫ̀ͅR̸̴̨̈́͂ͣͨͤ̐̇̑ͯͫ̓ͭ́ͦ͐͡҉̩͙̼ͅ ̸̢͓̱͎̥̱̺͚̎͐̈́́̊ͮ͛̀̈̅̎́ͫ̌ͭ̒͆́̚͞͞R̸̵͚̯͇̹̹̟̖͖̞͓̯͚͔̆̈́̆̆̍͡Ḗ̞͖̮̣̒̓̿̓ͪ̀̈́͌̂͆ͦͤ̐͌ͥ́͟͡T̷̵͉̪̼̗̼̭̙̲̖̰̤̲͍͕̥̱̅ͩ̇ͪ͌ͭ͌ͩ̆͒̀ͅƯ̯̞̜͎̣͈̻̠͈̦̠͚͋ͥͮ͐́ͬ͌ͥ̐ͨ͗̓̔͑͝R̷̨̧̎̃̓̏̋̂ͩ̾̓ͥ̍ͫ͒̊̈̓̆̑҉̹̲̝̜͎̠̥̣̥̭̪͎̙̖͎̟̳N̷̙̖̮̩̭̪̜͉̝͙̤̫̲̮̘͇̙̺̝̑ͯͥͮ̑̓͆̒̚", err)
	}
	defer file.Close()
	fmt.Fprintf(file, "L̨̤͉̖̹̟̲͖͇̤̪͕̤̤̠̥ͨ̔ͮͤͥͭ͑ͧͭ̑̊̈́ͨ̿͐̚É̵̢͔̟͕̝̹͎͚͎͈̳̔̇̃ͩ͝A̵̯̜̬̖͙ͥ́ͪͦ͊̍ͨ͆V̡ͩ̔̾͗̚͟҉̭̝̱̣̥̥̗̱̪͓͕͈̯͙͈͖Ĕ̛̦̗̫̟̖̦̞̠̙̝͉̖̦̤̻̏ͬͫͯ́̂͋̃̑̀͞ ̵̮̠̠̂̈ͨ͗̑̅͞Ä͇̙͕̲̪̭͕̮͚͈̭͍̫̖̺͙͚́͆̑ͭ́̒̎͘͜͡ͅN̎͒ͭͨ̊ͤ͌ͤ͑ͮ̐͆͡҉̹̲̯͘D̷ͭ͋̌̒͋ͣ͒ͩ̾̃̌̽̀ͫͦ͞҉̯̳̫̺̳̞̹̖̜̖ ̷̢̡̱͉̲͔̻̙̳̜̳͑̿̍͗͌ͦͬͨ̂͋̉̌͠N̵̛̲̟̦̠̟̻̲̯̲̦̝͍̳̣̳̺̤ͪ͐͛͆ͣ͒ͤ̊̍͊̊̓ͥ̈ͬ͒̈́ͬ̀͘̕ͅȄ̶̺̠̞̹̱̯̘̘͖̥̗̔̏̑͐̃̅͗͊͗̌̄̈́̿ͪ̀̕͘V́͆ͮ̔͂̔̉̌̋͐̀͟͝͏̺̬̘͕͙̫̺̫̜̗̘̝̼̫̺̫̹̞͘E̟̻͉͓̣̰͓̮̫̎̐͊̂̆̉ͩ͛̌̐́̅͊̐ͨ̊̀ͫ̀ͅR̸̴̨̈́͂ͣͨͤ̐̇̑ͯͫ̓ͭ́ͦ͐͡҉̩͙̼ͅ ̸̢͓̱͎̥̱̺͚̎͐̈́́̊ͮ͛̀̈̅̎́ͫ̌ͭ̒͆́̚͞͞R̸̵͚̯͇̹̹̟̖͖̞͓̯͚͔̆̈́̆̆̍͡Ḗ̞͖̮̣̒̓̿̓ͪ̀̈́͌̂͆ͦͤ̐͌ͥ́͟͡T̷̵͉̪̼̗̼̭̙̲̖̰̤̲͍͕̥̱̅ͩ̇ͪ͌ͭ͌ͩ̆͒̀ͅƯ̯̞̜͎̣͈̻̠͈̦̠͚͋ͥͮ͐́ͬ͌ͥ̐ͨ͗̓̔͑͝R̷̨̧̎̃̓̏̋̂ͩ̾̓ͥ̍ͫ͒̊̈̓̆̑҉̹̲̝̜͎̠̥̣̥̭̪͎̙̖͎̟̳N̷̙̖̮̩̭̪̜͉̝͙̤̫̲̮̘͇̙̺̝̑ͯͥͮ̑̓͆̒̚")
	sleeps(10)
	os.Exit(3)
}
func sleeps(x int) {
	time.Sleep(time.Duration(x) * time.Second)
}
func sleepm(x int) {
	time.Sleep(time.Duration(x) * time.Millisecond)
}
func pl(x string) {
	fmt.Println(x)
}
func p(x string) {
	fmt.Print(x)
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("The screen was supposed to be cleared here, but it seems your OS is unsupported. Sorry about that.")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello! Welcome to my crappy little game. It's still unfinished, so if you find any bugs, please report them to cmyui#5585 on discord.\nThank you <3\n\nThe game will start in 10 seconds")
	sleeps(10)
	CallClear()
	fmt.Println("??: Hey dude! Nice to meet you! I'm Josh, what's your name?")
	sleepm(600)
	var name string
	fmt.Scanln(&name)
	fmt.Println(name + ": Hey.. I'm " + name + ". Nice to meet you too!")

	if name == "josh" && name == "Josh" {
		sleepm(1400)
		fmt.Println("F͍̖̣̿̃ͅͅͅÙ̖C̜̯̤̗̓ͤ̉͆ͅͅK̝͓ͪͮ ͖̓ͧ̋O͓̺̟̜̦̽ͣ͗́̀̚F̥̘͖ͯͩ̀̑͑̓̚F̝͑͑ͪ̇̊͌ͥ ̰̇A̫͍̠͙͚̟N̳͓̟͖̎̊ͅD͇̼̠̱̱̩̙̈̓ͣ̏ ̪̗ͫ̅Nͣ̌ͬͬ̔̈E͕̱͚̗̠̠ͨ͑̀̉̌V̫͖̲̟̹͖E̬͕͕͈R͎͉ ̤̹̙̲R̪̳ͯ͛̃ͦ̎ͫ̌E̖̝͈̱͈̫͓͑ͪ̓ͬͨ̊T̬͓͖͓̠̗͎̔̃U̖ͧ̄̔R̤̰̼͎͆̇̋ͧ̒̿ͯNͮ̉")
		end()
	}

	sleeps(1)
	fmt.Println("Josh: Mm.. Have you ever heard of a game called osu?")
	sleepm(600)
	var playOsuString string
	fmt.Print(name + ": ")
	fmt.Scanln(&playOsuString)
	playOsuString = strings.ToLower(playOsuString)

	if playOsuString[0] == 'y' {
		sleeps(1)
		fmt.Println("Josh: Oh wow, really?")
		sleeps(2)
		fmt.Println("Josh: Do you actually play the game?")
		sleepm(600)
		var playOsu string
		fmt.Print(name + ": ")
		fmt.Scanln(&playOsu)
		playOsu = strings.ToLower(playOsu)

		if playOsu[0] == 'y' {
			sleepm(1500)
			fmt.Println("Josh: Oh! Holy shit! You're the first person I've met that actually plays the game!")
			sleepm(1500)
			fmt.Println(name + ": Haha, yea, not a lot of people know about it, and even less play it.")
		} else {
			sleeps(2)
			fmt.Println("Josh: Aw.. Damn.. I really wante to meet someone else that played it..")
			sleeps(2)
			fmt.Println("Josh: Anyways.. I've got to head home. See you around!")
			end()
		}
	} else {
		sleeps(2)
		fmt.Println("Josh: Aw.. Damn.. I really wanted to meet someone else that played it..")
		sleeps(2)
		fmt.Println("Josh: Anyways.. I've got to head home. See you around!")
		end()
	}
	sleepm(1500)
	fmt.Println("Josh: So, since we both play osu!.. Whos your favourite player?")
	sleepm(600)
	var favourite string
	var kill bool
	fmt.Print(name + ": ")
	fmt.Scanln(&favourite)
	favourite = strings.ToLower(favourite)
	sleeps(1)
	switch favourite {
	case "kablaze":
		kill = false
		sleeps(1)
		fmt.Println("Josh: Wait.. what?")
		sleeps(2)
		fmt.Println("Josh: Really..?")
		sleeps(2)
		fmt.Println("Josh: Huh.. He's actually probably my favourite aswell..")
		sleeps(2)
		fmt.Println("Josh: I think he's a really unrated and undervalued player overall..")
		sleeps(3)
		fmt.Println("Josh: Well, thats sortof his fault.. He doesn't really put himself 'out there' very much..")
		sleeps(2)
	case "cookiezi":
		kill = false
		fmt.Println("Josh: Haha. Yea, I definitely can't blame you for that one. I definitely think hes the BEST player,")
		sleeps(1)
		fmt.Println("Josh: Although, he's not my favourite..")
	case "howl":
		kill = false
		fmt.Println("Josh: :^)")
	case "solis":
		kill = false
		fmt.Println("Josh: oh yeah dude solis is my favourite relax cheater too :)")
	case "cmyui":
	case "chase":
		kill = false
		fmt.Println("Josh: <3")
	case "rafis", "azerite", "zirba", "hvick", "hvick225", "bubbleman",
		"rohulk", "angelsim", "emilia", "woey", "osu player84", "mathi",
		"totoki", "_ryuk", "ryuk", "mcy3", "spare", "danyl", "piggey",
		"bro_gamer72", "gayz", "gayzmcgee", "dustice", "idke", "happystick",
		"_index", "index", "karthy", "yaong", "digitalhypno", "toy", "aireu",
		"wubwoofwolf", "talala", "resia", "adamqs", "wilchq", "mouseeasy",
		"flyingtuna", "thepoon", "recia", "dumii", "mlaw22", "mlaw", "mshake",
		"xilver", "ceptin", "plz enjoy game", "apraxia", "doomsday":
		kill = false
		fmt.Println("Josh: Yea, I can't blame you for that. He's extremely talented. Although..")
		sleeps(1)
		fmt.Println("Josh: He's not my favourite, but he's definitely up there.")
	case "rrtyui":
		kill = false
		fmt.Println("Josh: Meh.. Irrelevant trash player, but i'll let it slide..")
	default:
		kill = true
		sleeps(1)
		fmt.Println("Josh: ..")
		sleeps(2)
		fmt.Println("Josh: Who..?")
		sleeps(1)
		fmt.Println("Josh: Oh well.. whatever.. It's not like I cared anyways")
		sleeps(4)
	}
	if kill == true {
		fmt.Println("Josh: Actually what the fuck? How do you not appreciate any good players?")
		sleeps(3)
		fmt.Println("I literally have no idea who you're even talking about. What the hell is wrong with you??")
		sleeps(3)
		fmt.Println("Josh: Are you mentally insane? I don't understand your thinking")
		sleeps(3)
		fmt.Println("Johs: Yea.. You're probably insane. You should go get that checked out..")
		sleeps(3)
		fmt.Println("josh: Dude go to the hospital already what the hell are you doing")
		sleeps(3)
		fmt.Println("Josh: You're clearly stupid and delusional if you don't think kablaze, cmyui, or cookiezi is the best player.")
		sleeps(2)
		fmt.Println("J̢̙̮̙̮̬͍̰̺̗͚̥̠͉̹̙͕̹ͅͅoŞ̧̡̨͍̘̖̭̭̗͕̫̭͉̰̥͎͈̻͈h͙̖͇̘̱̦̩̱̞̝̞̬̯̣̹̞̰͉̮̯:͇̱̬͎̣ ̢̰͕̟̪̹̯͍̥͍͍̤̺̫̱̟̟͚̰͜W̡̡͕̺̬̠Ḩ͍̬̠͖̲̝̼̗͙̤̞̤̗̹̰̩̣͜͜Y̧̧̢̧̡̧̹̲̖̺̩̲̞͙̹̬̭̦̣ͅ ̡̢̫̺̥̦̘̻͔̦͔̙̣̝͔͓̫̭̤ͅA̢̡̧̡̼̰̬̗͚͔̠̜̺̲̝̯͕̫͓̭�̧̨̮̲͈̭͚̗̗̗̠̺̱̳͙̣͇̼͖̳�̡̧̧̨̙͉͇̲̗̣̹͙͇̤̜̣̖̟̩̠�̢̠͉̯͇̰͓͔̞̯̺̥̝͉̤͍̳̞͔̮R̺̭͇̼͖Ȩ̞̮̭̩͇̟̥̥̠̖̟̲̮̖̰͔͜ͅO̢̢̫͍̲͇̘̦̮̘͙̙̼̳͙̤͙̞͕̹U̡̥͇͔̞̝̝̩͇͍͔̰̭̜̯̫̯͇͜͜�̢̧̢̧͎̻̘͚̣̻͎̮͇̟̖̺̯͖̭̟  ̨̻̗̭̭͇̘̼̹͍͙͖̬̤̞̼͕ͅE̢̠͔̬̳̙͖̰̭̺̪̗̠̲̘͚̞̗̭͜Ã͕̦̯̜̺͉͎͇̜̟͔͎̬̞̘̹̞̘̥ͅ³̢̯̘̼̯̤̭̬̳̟̫̲͓ķ̫̫̞͈̳̜̦̳̤̺͕̪̥̲̟̰̹͜ͅǫ̡̙̣̯̤̘͔͙̙͎̹̪̞̻̟̬̣͙͜r͚̥̙̘̪ͅȨ̮̥͇̤͉͍̦̞̹͉̞̹̳̤̣̫̜̻̬?̢̝̬̹̣̝͓͍̹̱̺̩̪̺̤̭̗͉̩̞?̨̡̙̮͖̞̟͙̺͖̩̭̻̯͓̖͉̺̦͜")
		sleepm(800)
		fmt.Println("jalapeÃ±o: L E A V E")
		sleepm(800)
		count := 1
		for count < 20 {
			fmt.Println("\n� /? / ? LFKS�?fsajj �")
			count++
		}
		end()
	}

	sleeps(2)
	fmt.Println("Josh: To be honest, i'm amazed you've made it this far..")
	sleeps(2)
	fmt.Println(name + ": What?")
	sleeps(2)
	fmt.Println("Josh: mm.. Nothing.. It's almost like you're trying to read my mind is all..")
	sleeps(2)
	fmt.Println(name + ": Uh.. Sorry I'm not sure I know what you're talking about..")
	sleeps(2)
	fmt.Println("Josh: heh... Of course not..")
	sleeps(4)
	fmt.Println("Josh: Have you ever heard of Akatsuki?")
	sleepm(300)
	fmt.Print(name + ": ")
	var akatsuki string
	fmt.Scanln(&akatsuki)
	akatsuki = strings.ToLower(akatsuki)

	if akatsuki[0] == 'y' {
		sleeps(2)
		fmt.Println("Josh: Yea.. It's been getting quite a lot of attention recently..")
		sleeps(2)
		fmt.Println("I'm actually the creator of it.. haha")
		sleeps(2)
		fmt.Println(name + ": Oh dude, thats kind of cool.. although.. isnt it litrally just ripples code..?")
		sleeps(2)
		fmt.Println("Josh: �.")
		sleeps(4)
		fmt.Println("Jos�̡̢̡̝̥͇̠̗͖̙̹̥̩̣̻̞͎̜̹͇̻͚̖̭̖̟̼ͅh̨̢̨̨̭̞̖̲̝̭̠̜̳͔̱̫̞̼̠̖͎͎̠̰̦̦̬̭:̡̺̰̲̣̲̼̲̖͔̗̞̲͇͈̖̠̯̜̲̮̦̼̹̻̠̱ͅ ̡̧̢̢̟̩͈̞̳̙͔͔̜͓̼̥̭̝̘͔͓̝̻̪̯̯̹ͅI̧̻̯͍̜͓͈̹͖̰̬̪̹͇̱̝̣̥̠͙̣͉̲̠̺̠̲ͅ'̨̧̢̢̡̨̡̙̟̙̳̥͓͎̼̩̰̠̟̞̗͍̺͙̪̬̝ͅm̢̡̧̨̩̜̠̝̱͕̲̦͇̰̬͎͙̠̱̼̹̯̥̘̰͈̥̣ ̡̡̹͓̥̙̣̙͈͈̘̠̣̰̞͈͎̲̩̣̼͎̝̳̣͍͜ͅs̢̢̫̜̘͖̞̫̩̩̟̻͔̻̙̼̜͓̭̜͔̬͕̣͜͜ͅͅo̡̢̨̲̼͉͇̟͖̯̰͉̦͍̻͓̜̻̫͇͍͖̮̰͍̯͚͜ŗ̥͉̜̩̰͉̩̼͖̮̱̻͍̩͎̮͔͕̲͈̹͇̦̪ͅͅ    rRrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr-")
		end()
	} else {
		sleeps(2)
		fmt.Println("Josh: Oh.. Yea it's an osu! private server I made using Ripple's source as a base..")
		sleeps(2)
		fmt.Println("Josh: It sortof gets a lot of hate due to the fact that most people just see it as a ripple clone..")
		sleeps(2)
		fmt.Println("Which I actually understand, and even agree with.. but I wish that people wouldn't spread so much hate about it..")
		sleeps(2)
		fmt.Println("Oh well.. This isn't the place to vent about that kind of stuff anyways..")
		sleeps(2)
	}

	fmt.Println("Have you done anything interesting recently?")
	sleepm(600)
	fmt.Print(name + ": ")
	var interesting string
	fmt.Scanln(&interesting)
	interesting = strings.ToLower(interesting)

	if interesting[0] == 'y' {
		sleepm(900)
		fmt.Println("Oh dude, tell me about it!")
		sleepm(600)
		print(name + ": ")
		var blabla string
		blabla = strings.ToLower(blabla)

		if blabla[0] == 'n' {
			sleepm(900)
			fmt.Println("ẅ̮̙̪͚̭̪̳͓̩́̐̈̎͊̌̃̊̀͗ͅĕ̮͉̖̫̣̱̦̰̜̽̑̿́͆̿̈͂͠ͅl̪͓̘̲̙͓̠̱̟̎͑̎̌̿͑̎͑̇̕ͅl̡̨͎̭͔̖̹̺̻̗̅̑̉̌̂́̄̒̎̕ ̢̗̞̻̟͓̘͖͎̲͆̒͌́͆̎͛̂͆͝f̹͙̗̲̝̮͎̮̜͎̆̎̉͌̊̈̃̍̎̚u̡̜̩̗̺̺̱̺͖̗̔̍̈̓̈́͒̎͛͝͝ç̬̪̟̟̤͙̹̟͔̽̋̀̈́͛̾͋̿̈͝k̦̥̩̬̞̗̟̳͎͇͒̓͗̊̾̿̒̈́̆̊ ̢̜͚̳̙̖͉͇͇͂̑̂̇̄͆̌̉̈́͘͜y̹̰͍͍̙̲̰̞̳͉͒̐̓̆̉̓̆̆̚͝o̞̫̻̗̖͎̬̝̝̽̐̈̆̔͑̾̎͘͝ͅu̠̞͕̮͎̫̦̻͈͔̐̆͊͋͛͛̎̔̏͝ ̨͙̳̝̭͚̰̲̰̘̋̾͐̌̅̈́̈̃͆͘ẗ̛̜̝̠̖̪̬̹̗͍̺͂͆͛́͂̒͊͘ơ̧̹͚͇̯͎͕̳̘͍͌̓̑̉̐͐̈́̊̋ǫ̧̡̟̟͇̥̯̠̱̈́͗̌͗̈́̒̍̉͝͠")
			end()
		}
		sleepm(1500)
		fmt.Println("Wow.. That's super cool! I hope I can do something like that someday..")
		sleeps(1)
	} else {
		sleepm(600)
		fmt.Println("Josh: Oh.. Well that kinda sucks.. I haven't done much either I guess..")
		sleeps(2)
		fmt.Println("Josh: Well.. At least other th�n the osu! server. But that wasn't even me really. I really owe it to Howl and Nyo.")
		sleeps(2)
		fmt.Println(name + ": Well.. Why not do something for them then..?")
		sleeps(2)
		fmt.Println("mmm..")
		sleeps(2)
		fmt.Println("Josh: I'm really now sure what I could even do.. One of their developers doesnt like me at all\nand I'm not sure what the other one thinks..")
		sleeps(4)
		fmt.Println(name + ": Oh well Sucks lole")
		end()
	}

	fmt.Println(name + ": I'm sure you'll do alot more in your life. haha")
	sleeps(2)
	fmt.Println("Josh: Yup.. Most likely. I'm just waiting for the day where I come out of my shell.")
	sleeps(2)
	fmt.Println("Josh: So.. Out of curiousity, are you any good at osu..? What's your rank?")
	sleepm(600)
	fmt.Print(name + ": ")
	var rank int
	fmt.Scanln(&rank)
	if rank >= 100000 {
		sleepm(400)
		fmt.Println("Josh: Oh, so you're still pretty new to the game then!")
		sleepm(1200)
		fmt.Println("Josh: I remember when I was that rank.. I first got my tablet right around 250k, haha")
	} else if rank > 10000 {
		sleepm(400)
		fmt.Println("Josh: Oh. So you've been playing for a bit, I guess. Still pretty new though, haha.")
		sleepm(1200)
		fmt.Println("Josh: You'll probably start to see a slowdown in progress around the rank you're at now though. That's when it starts to actually get difficult..")
	} else if rank >= 3000 {
		sleepm(400)
		fmt.Println("Josh: Oh damn.. You're catching up to me.. I got banned at about 2.9k, but that was after 2 cheated plays.. In reality i'm probably around there now.")
		sleepm(1200)
		fmt.Println("Josh: Best of lucking kicking my ass, haha.")
	} else if rank >= 1000 {
		sleepm(400)
		fmt.Println("Josh: Oh damn, you're a bit further than I am then..")
		sleepm(1200)
		fmt.Println("Josh: We might actually be able to compete, I'll be sure to try to snipe some of your scores, haha")
	} else if rank >= 100 {
		sleepm(400)
		fmt.Println("Josh: Oh wow.. You're really good damn. I hope I'll reach that sort of level one day..")
		sleepm(1200)
		fmt.Println("Josh: Not sure if I can, but I think it's a good goal.")
	} else if rank >= 35 {
		sleepm(400)
		fmt.Println("Jo�h: ?  ????????? ? ?       ????????????  ?  ? ? ???\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? �?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? �?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? ?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? ?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? ?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? ?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?")
		end()
	} else if rank >= 20 {
		sleepm(400)
		fmt.Println("Josh: Bro im finna smash ur ass like grass at ANIME NORTH CON 2018 :^)")
	} else {
		sleepm(400)
		var count int
		count = 0
		for count < 20 {
			fmt.Println("\nJo̧̡̧̨̨̧̧̢͓͓̼̘̝͇͈͍̼͓̦̟̝͉͍͍̬͈͔̜̥̜̱͕̲̲̗̥̖̻̦̠̠̹̣̣͖̞̺̻͖̤̪̭̳͖͈͚̦̳͔̥̰̜̪̩̱̗͚̞̯̘̮͕͔͙͍̭͙̳̺̱̥̣̗͖͙͎̥̣̭̞͇̦̠̘̥͜ͅͅͅͅ࿊hs�h: loO  Le h�e SDlo i am   coo   kiez i xDD℻DS�႟DDSD��᥀᥀℆℆℆D")
			count++
		}
		end()
	}
	sleeps(1)
	fmt.Println("Josh: Anyways, I've got to go home.. Do you use discord by any chance?")
	sleepm(600)
	fmt.Print(name + ": ")
	var haveDiscord string
	fmt.Scanln(&haveDiscord)
	haveDiscord = strings.ToLower(haveDiscord)

	if haveDiscord[0] == 'y' {
		sleepm(600)
		fmt.Println("Josh: Oh, sweet! What's your ID? Mine is cmyui#5585")
		sleepm(300)
	} else {
		sleepm(600)
		fmt.Println("Josh: Oh damn.. Well.. Maybe we'll meet again one day, haha.")
		sleeps(1)
		fmt.Println(name + ": Maybe.. Well.. See ya!")
		sleeps(1)
		fmt.Println("Josh: Cya!")
		end()
	}
	var discord string
	fmt.Print(name + "Yea sure, mines: ")
	fmt.Scanln(&discord)
	discord = strings.ToLower(discord)
	switch discord {
	case "cmyui#5585":
		fmt.Println("J̢̢̧̨̨̡̨̧̨͓̱̹̝̣̳̥͎̯̫̣̭͇̗̮͚̼̬͈͕̫̲̩̮̠̖̗̺̘͉̝͍͙͇̼̥͖͙͇̖̬̤̘̭̠̱̺͖̮̖̘̹͓͈̩̲̞̹̪̪̞͕̗̺̬̤̫̟̹͇̦͖͎͓͉̻͍̩͈͓̼̫̰̠͜͜͜͜ͅͅͅo̢̡̨̧̨̧̡̨̡̡̡̢̭̼͉̭͇̞͔̻̪̤̼̯̣̼̥̮͉̣͚̩̯̩͖͖̜͇̖̥͚͙͎̟̪̘̻̝̥͖̖̬̖̩̪̦͎̝̩̜̤̫̳̤̹̙̩̭̘̳̖̥̠͎̫͈͈̙̣̬̺̤̳̩̟̰̳͍̮̝͓͙̰̦̗̟̬͚ͅͅS̢̧̢̨̡̧͚̰͔̰̮̞͕͇͉̪̝̼̟̜̠̟̻̬͎̠̲̘̭̹̗͓̺̱̹̤͎̤̲̣̳̯̺̖̪͙͚̖͈͈̩̯͓̻̯͚̲̮͇̲̳̞̬͙̠̖̘̣̹͍͈̤͍̗̗̥͚͎̞͔͙̟̤͖̭̪͚͈͈̙̥̠̦̫̱̥͜ͅͅͅH̨̧̡̡̧̠̫̤̹̞͙͉͈̻͓̬̩̲̖̜̰̦̥̬̙͚̥͙̥̹̠̣̝̖͖̰̜͎̮̬͓͔̠͉̭͎̯̣̻̰͔̖̩͕̜̰̤̭̩̪̥͚̘̦̖͎͚̝̬͉̠̺̼͍̻̣̟̗͉̳̱̩̰̯͇̬̠̘͙͕̖͖͜͜͜ͅͅͅͅͅ℆̨̨̢̨̢̨̨̨͙͙̘͓͖̜̹̯̥̪͍̫̗̝̬͖̬̞̰͖̭̝̮̫̤̫̦̩͙̝̞̟͇͉̟̹̤̙̞̺̖̤̳̲̣̟̹͍̦͙̫͚̣̜͕̜͕͔̙̤̥̻̝͓̝͓̙̞̹̺͇̠̟̬̬͎̰͖͙̝̫̫͎̟͓̩͍͕͜͜ͅͅS̢̡̡̨̧̢̡̡͈̜̩͓͍̰͚̪̰͓̲̦̺̦̬̜̝͍̙̜͕̥̟̼͈̣̟̥̼̤̤̤̮͎̱̻̻͈̹͎̘̺͉̞̯͎͈̘̹͔̼̭̳̖͈̞͇͉͓̹̣͔̭̪͔̠̞̫̯̥̼͉͍͉͍̠͚̮̞̻̼̖̦̞͎̬̘͜͜͜ͅͅ:̧̢̢̨̢̧̡̨̭̗̯͔̦̙͙̼͇͙̟̻̹͎̠̗̝̦͚̭̣̟̫͉̪͍̦̬̥͚̪̖̱͙̫̭̯̯̲̼̘̞̦̪̝̤̫͈̰̗̗̩͉͙̹̺̙̤͕̲̹̣̦͈͚͔̞̳̞̖̼͎͈̲͙̰̬̯͚̙̫͕̣̻̞̲̙̞͜͜ͅͅ ̡̡̡̡̨̡̧̡̻̫̗̫̟̼̫̣͉̩̯̯̺̭͈̦̺̙͖̗̺͙̣͚͕̥͖͚̩̪͇̘̜̤̘̮̳͇̠̱̯̦̬͎͓̝̼̹͙̣͙̳͓͖͙͈̭̤͖͍̘͈̠͕̤͈͎̙̦̩̯̜͔͙̜̦͖̻̪͙̙̲̟̱̝̝͜͜͜͜ͅͅJ̡̢̡̧͉̩̥͙̞̫̥̥̜̼̩̤̮̣̳̥̰̖͍̲̮̗̖̭͈̗̻̮͍̼̯̰̣̪̮̰͇̤̩̙͇̟͈̺̲̯̳͉̱̘̪̬͉̤͕͉̜̞̰̣͚͎̣͉̮̞̪̪͈̙̟̺̘̝̲̮̝̦̭̰̥͙̙̝̗̱̫̳̩̟̹̬̳͜͜ͅS̡̢̡̨̢̢̨̜̹̗͙̥̬͈̪͈͍͈̺̦͖͇̻̤͎͔̞̘̥̝̯̜̻̝̬̖͈̪̱̮͇̗̝͔̘͉͉̮͍̣͈̻̩͖͓̖̝̖̪͍̟̹̰͍̼͓̪̖̖̦̯̺͚̲̦̺̥̘̯̤͍̺͈̠̰͔͎͙̻̳̠̜̰̳̙͜͜ͅͅͅI̧̧̡̧̨̨̨̮̪̥̠̼͙̹̘̲̱͈͉̗̘͈̘̘͔͎̗̣͙̹͕̙͎̫͇͇̼͚̩͍̮̱̗̮̦̦̥̭̲͓̻͔̱̤̗̟̝̝͔̖̻̣̼̹͚̰̬̻̪̣̳̥̜̗͔̭̠̲͙̠̥̟͖̟͈̗̘͍̰̰͎̦͚̤̼͓̹͈͜ͅH̨̧̡̨̡̧̨̡̢̧̢̟͉͓̤̞̭̥̘̦̜͓͉̘̣͖͕̥̜̣͉̟̣̣̞͔͉̳̠͓͍̦̘̳̬͙̟̣͍̩͎̯͉̹͚̗͖̦͇͖͖̻̩͍͕̹̪̲̳͍̭̟̬̖͍͕̠͍̲̹͕̪̳̺̭̟̱͕͖̙̯̙͔̲̫͜ͅͅͅͅ ̧̧̡̢̢̧̧̡̡̨̨̡̧͉͕͚̘̞̭̻̫̺͔̠̜̝̭̟̬͓̝̘̯̙̬̰̞̯̫͔̼̼͈̰͖͍̗̯̳̗̮̗̣̹̗͉̻͕̟̼̦͔͇̩̝̥̹̖̰̥͎͕̮̗̜̭̬̻̹̙̖͇͔̱͙̙̗̮̜̜̗̰̟̥͕͜͜͜ͅͅą̨̨̧̨̢̡̧̨̡̧̨̡̨͈̼̹̟̭̣͈̟̰̣͔̘̲̘͚̫̥͈̲͓͈̝̺̳̪͕̗̻͈̱̠̩͚͙͍̜̦͓͍̜̝̹̙̺͕̺̻̝͔̺͕͚̤̟͖̪͉̤̘̳̟̬̥̞̠̱̝͎͉̫̠̱̮̹̺̻ͅͅͅͅy̡̢̡̢̢̻̣̜̫͈̥̗̭̱̼̗̪̯̹͓̝͇̙͇̞͈͓̘̣͚͔͍̺̲̳̦̤̖͔̼͓̝̣̩͖̬̜̹͕̺̜̬̱̰̼̰̤̞͜͜͜͜ͅͅş̢̢̢̢̨̯͉̺̙̰̘̟̫̦͓̩̦̦̻̮̖̗͔͈̠̺͓̹̥̦̫͙͇̯̟͔̥̥̺̞͚̜̘̲̝̖̘͕̣̯̻̳̻͇͎̭͎̣̳͚͎̼̙͔͖͚̖͙̜̮̫̤̲͍͚̮̮͇͜͜͜͜͜ͅͅͅş̨̡̢̨̨̡̨̨̧̙̗͖̪̖̪͔̹̥̱̜̻̞͙̩̪͓̺͔̻̩̪̙͔̺̭͈̻̝̫̘͓̦̗͚̹̜͔̳̳̺̬̳̪̟̠͍̥̪͍̜̖̰̹̭̥̲͈͈͔̹͖͕̭̙̯̤̳̝̠̫̭̗͔͍͕̳̬̣̠̗̣̱͜͜͜͜͜͜ͅy̡̢̧̨̨̧̡̯͎̝̲͓̳̪̭͈̬̹̱̞̰̮̖̪̥̥̮͕͉̺̠͕̞̣̟̞̺̫̟͍̞̥͔͕̹̙͓̺͕̬̟̼̱͙͕͇̪̣̘̞̯͍͚͎̘̜͓̪̻̠̫̻̮̝̖̬͚̫͙͈̘̟̲͖̜̱̬̯͚͍̠̪͚͓̼̲͜͜ͅͅb̨̡̢̨̡̢̢̡̨̡̢̡̥̦̫̫̤͉̤͓͕̲̩̭̠͖͍̹̝͙̹̥̱̟̹͍̙̻̥̻̳̠̟̙͕̮̮̪͈͔̰̰̮͍̪̻͕̗̱̠͈̤̪͇̳͔̲̭̗̜͖̬͕̝̤̳͎̝͙̩͚̳͕͔̙̹̫̤̻͚̱̜͇͖̯̰͈͜ͅͅF̢̨̢̨̢̧̧̣̪̜͙͖̱̺̞͙͙̼̫̲̞̭͔̳̫̦̲̰̹̥̟̹͓͈̜̝͉̫͇͇͉̭͎͕͍̦̘̳̮͕̙͚̫̰̰͓͇̤͕̫̰͉̬̦̮͕̪̖͕͚͕͍̖̯̮̗̮͓̩͖̖̘͚̣̟͍͓̦̘̪̞̦͖̪̤̹̲͜͜ͅB̢̢̡̢̡̡̢̧̢͍̙̠̦̦̳̙̭̯̹̹̗̗̻͕̻̫͓͎̪̙͙̪̫͕͕̭̯͓͈̱͍̞̦͖̬̙̻̜͙̳̖̼͖͖͔͙̲̮̩̱̟̺̹̻̤͓̠̹̘̬͕͚̥̝̦̝̯̯̼̹͕̤̺͚͖͎̫̤̖͔͎̦̦̫̯͔͜ͅͅͅS̨̡̡̧̨̨̢̡̭͇̙̲͈̦͉̹̝̦̝̦͓̻̙͍̺͉͇̘͇̣̭͎͔̭͚͔̯̝̦̹̰̮͙̱͉̜̘̫̺̯̻̣̻̺̯̠̪͕̳̖͚̦̬͙̮͙̦͓̼̰͓͕͍̩̪̗̞̩̝̣̥̮͍̗̜̲͇͎̯̳̖͎̞̞̣̭͚͇̞ͅF̨̨̨̧̢̧̢̧̨̢̠̺̯̩̟̭͈̰̬̻̖͍̼͍͍̱̭̭̞͖̮̪̙̖̻̩̫͍͔̟̝̯̮͈̺̜̳̥̮̟̺̤̩̹͙͉̞̖̭̥͚͚͓̦̹͎͈̥̻̰̤̱̻̺̙̗̝̥̣̱̱̮̝̜͍̺͔̣̪͎͕͕͜͜͜͜͜͜͜ͅḄ̨̧̨̢̢̡̢̭̮̟̲̪̘̯̝̫̘̣̲̖̪̙̠̼̦̭̯̪̥͉̹̬̩͍̺̩̲͖͚̲̬̙͈̦̠̞̖͔̙̗̭͙͈̲̘̠͓̖̖͜͜b̢̢̧̧̡̧̡̧̧̢̨̧̖̳̳̰̬̩̭̱̬̜̳̺͍͚͎̫͕̖̟̗͖͓̦̙̜͓̝̥͖̺͉̱͕̘̭͓̬̙̠̠̯̖͔̗̖̞͕͉̝̗͙͚̻̙̥͕̟̠̠͈̜̜̳̪̼̜̮̙̟͔̟̹͈̫̭͍̳̘͎̤̖͚̠̼͚͜ͅͅy̧̢̨̧̡̺̲͔̠̮̬̬͇̦̼͖̤͙̗̻͇̬̞̲͙̮̳̼̙̠̭̫̮̱̳͓͉̜̜̮̣̣̞̘͓͎͕̺̻͓̬͓̹̹̲̩͕̥͔͈̖̻̖̤͕͚͎̘̠̟̺͉͕̞̟͚̤͓͖̥̜̯̹͉̯̝̪͍͚̦̦̪̖̪͓ͅͅͅͅͅ\nN̢̧̡̢̡̡̡̟̟̰̪̩͇̝̻͔̪̖̲͔̮̰̞̫͕̙͚̮̹͕͈̩̼͕͚̗̗͚͖̖̻͈̟͖̠̥̝̳̦̝͙̩͚̼̜͉͔͓̺̣̬̜̱͇̤̙̲̖̰̳̼̮͔̱̹̬̱̥̰̞̥̰̻̰̣̟͖͚̰̺̜̜̟͎̼̯̹̙͜͜ǫ̨̨̨̢̡̡̨̧̢̧̟̙̰̘͎͈̮̟͕̜̱̥̞͕̙͖͉̬̞͎̟̣̫̱̼̘̭̝̙̲̗̭̠͍̹͙̬͔̭̰̖̯̞̩̩͇̲̜͈͔̪̮͎̮̬̱͖̱͓͕̪̜̱̪̖͈͙̣̜͍̣̰͈̱̼͓̹̩̙͜͜͜͜ͅͅͅͅͅͅ.")
		end()
		break
	case "kahri#9828":
	case "Emily | Sunpy#5213":
	case "Hanyeol 한열#8257":
	case "Kip#1120":
	case "Frosti#0602":
	case "ninjin#3198":
	case "Howl#0940":
		fmt.Println("Josh: Huh.. That sounds oddly familliar..")
		sleeps(2)
		fmt.Println("Josh: well, whatever.")
		break
	default:
		fmt.Println("Josh: Alright cool!")
		break
	}
	sleeps(1)
	fmt.Println("Josh: I'll make sure to add you when I get on!")
	sleeps(2)
	fmt.Println("Josh heads off towards his home, and I do the same..")
	sleeps(1)
	fmt.Println("..")
	sleeps(1)
	fmt.Println("..")
	sleeps(1)
	fmt.Println("..")
	sleeps(1)
	fmt.Println("..")
	sleeps(1)
	fmt.Println("..")
	sleeps(1)
	fmt.Println("I open my door, take a seat and turn on my PC.")
	sleeps(1)
	fmt.Println("Mmm.. It's been a long day..")
	sleeps(2)
	realEnd()
}

var x int

func realEnd() {
	sleeps(3)
	CallClear()
	fmt.Println("Congratz, you've reached the end of the game.. for now.. Feel free to contribute, or just tell cmyui you've reached this point,\nas theres only a few that have <3")
	sleeps(10)
	os.Exit(3)
}
func end() {
	sleeps(3)
	CallClear()
	fmt.Println("Game Over")
	sleeps(10)
	os.Exit(3)
}
func sleeps(x int) {
	time.Sleep(time.Duration(x) * time.Second)
}
func sleepm(x int) {
	time.Sleep(time.Duration(x) * time.Millisecond)
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Woops.. The screen was supposed to be cleared here, but it seems your OS is unsupported. Sorry about that >.>")
	}
}

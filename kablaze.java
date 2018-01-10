import java.util.Scanner;

/*
* Author: cmyui#5585 // https://akatsuki.pw
* Hey! Welcome to my first ever Java game. It's probably going to suck, and it's literally just text.
* But I'm proud of it so far lmfao. It's supposed to be some sort of little game to see how far you
* can get before the game closes on you. Try not to look at the source code too much, and try to pretend
* you're me so you can predict the answers :^)
*/

public class kablaze {
	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		String playOsuString, akatsuki, interesting;
		int careAmt;
		boolean playOsu;

		System.out.println("??: Hey dude! Nice to meet you! I'm Josh, what's your name?");
		String name = scan.nextLine();
		System.out.println(name + ": Hey.. I'm " + name + ". Nice to meet you too!");

		if ("Josh".equalsIgnoreCase(name)) {
			Sleep.main(1400);
			System.out.println("Josh: Oh fuck off..");
			Sleep.main(1600);
			System.exit(0);
		}
		Sleep.main(1000);
		System.out.println("Josh: Do you play osu?");
		playOsuString = scan.nextLine();

		if ("y".equalsIgnoreCase(playOsuString.charAt(0))) {
			System.out.println(name + ": " + playOsuString + ".");
			Sleep.main(200);
			System.out.println("Josh: Oh dude, sweet! Finally another person that plays osu!");
			playOsu = true;
		} else {
			System.out.println(name + ": " + playOsuString + ".");
			Sleep.main(200);
			System.out.println("Josh: Aw.. Damn.. I really want to meet another person that plays the game..");
			Sleep.main(1000);
			System.out.println("Josh: Anwyays.. I've got to head home. See you around!");
			Sleep.main(1600);
			System.exit(0);
		}
		System.out.println("Josh: Ahh dude I forgot.. Could you tell me what kablaze's current osu! rank is?");
		int number = scan.nextInt();
		System.out.println(name + " " + number + ".");
		Sleep.main(400);
		System.out.println("Josh: Hmm...");
		Sleep.main(1600);

		if (number <= 25) {
			System.out.println("Josh: Well.. Fuck.. That means the worls will be ending tomorrow..");
			Sleep.main(400);
			System.out.println(name + ": Huh?... Oh.. Ok.. I guess..");
		} else {
			System.out.println("Josh: Ahhhhh... Balance has been restored to the universe...");
			Sleep.main(1000);
			System.out.println("Josh: Goodbye..");
			Sleep.main(1600);
			System.exit(0);
		}
		Sleep.main(3000);
		System.out.println("Josh: Anyways.. Since you're doomed anyways, might as well have a discussion.");
		Sleep.main(500);
		System.out.println("Josh: I'm guessing we both play osu!.. So whos your favourite player?");
		String favourite = scan.nextLine();

		if ("kablaze".equalsIgnoreCase(favourite)) {
			System.out.println(name + ": kablaze.");
			Sleep.main(400);
			System.out.println("Josh: Holy shit no way!!! ME TOOOO!!!!!!!");
			careAmt = 5;
			Sleep.main(4000);
		} else if ("cookiezi".equalsIgnoreCase(favourite)) {
			System.out.println(name + ": Cookiezi.");
			Sleep.main(400);
			System.out.println("Josh: Yea dude he's pretty good aswell");
			careAmt = 2;
			Sleep.main(4000);
		} else if ("howl".equalsIgnoreCase(favourite)) {
			System.out.println(name + ": Howl.");
			Sleep.main(400);
			System.out.println(":^)");
			careAmt = 2;
			Sleep.main(4000);
		} else if ("cmyui".equalsIgnoreCase(favourite)) {
			System.out.println(name + ": cmyui.");
			Sleep.main(400);
			System.out.println("Josh: Haha.. Thanks man. <3");
			careAmt = 3;
			Sleep.main(4000);
		} else {
			System.out.println(name + ": " + favourite + ".");
			Sleep.main(400);
			System.out.println("Josh: ..");
			Sleep.main(1800);
			System.out.println("Josh: Who..?");
			Sleep.main(3600);
			System.out.println("Josh: Oh well.. whatever.. It's not like I cared anyways");
			careAmt = 1;
			Sleep.main(4000);
		}

		if (careAmt == 1) {
			System.out.println("Josh: Actually what the fuck? How do you not appreciate any good players?\nI literally have no idea who you're even talking about. What the hell is wrong with you??");
			Sleep.main(300);
			System.out.println("Josh: Are you mentally insane? I don't understand your thinking");
			Sleep.main(300);
			System.out.println("Johs: Yea.. You're probably insane. You should go get that checked out..");
			Sleep.main(200);
			System.out.println("josh: Dude go to the hospital already what the hell are you doing");
			Sleep.main(150);
			System.out.println("Josh: You're clearly stupid and delusional if you don't think kablaze, cmyui, or cookiezi is the best player.");
			Sleep.main(120);
			System.out.println("JoSh: WHY A���RE YOU� STILL HEÃ³korRE??");
			Sleep.main(80);
			System.out.println("jalapeÃ±o: L E A V E");
			Sleep.main(20);
			int count = 0; 
			while (count < 20) {
				System.out.println("�J ௸�    O �୰�Shs: G ௵  �Ã³k߶E T    O� U҂T.   ���   .");
				count++;
			}
			Sleep.main(60);
			System.exit(0);
		} else if (careAmt == 5) {
			System.out.println("Josh: I'm actually really suprised you think kablaze is the best..");
			Sleep.main(700);
			System.out.println("Josh: He's usually a super underrated player so it's really nice when I can find someone\nThat can appreciate what he brings to the table..");
			Sleep.main(950);
			System.out.println("..");
		}
		Sleep.main(1000);
		System.out.println("Josh: To be honest, i'm amazed you've made it this far..");
		Sleep.main(300);
		System.out.println(name + ": What?");
		Sleep.main(500);
		System.out.println("Josh: mm.. Nothing.. It's almost like you're trying to read my mind is all..");
		Sleep.main(600);
		System.out.println(name + ": Uh.. Sorry I'm not sure I know what you're talking about..");
		Sleep.main(600);
		System.out.println("Josh: heh... Of course not..");
		Sleep.main(3600);
		System.out.println("Josh: Have you ever heard of Akatsuki?");
		Sleep.main(600);
		akatsuki = scan.nextLine();

		if ("y".equalsIgnoreCase(akatsuki.charAt(0))) {
			System.out.println(name + ": " + akatsuki + ".");
			Sleep.main(200);
			System.out.println("Josh: Yea.. It's been getting quite a lot of attention recently..");
			Sleep.main(400);
			System.out.println("I'm actually the creator of it.. haha");
			Sleep.main(400);
			System.out.println(name + ": Oh dude, thats kind of cool.. although.. isnt it literally just ripples code..?");
			Sleep.main(1400);
			System.out.println("Josh: ..");
			Sleep.main(6000);
			System.out.println("Josh: I'm sorry");
			Sleep.main(100);
			System.exit(0);
		} else {
			System.out.println(name + ": " + akatsuki + ".");
			Sleep.main(200);
			System.out.println("Josh: Oh.. Yea it's an osu! private server I made using Ripple's source as a base..");
			Sleep.main(500);
			System.out.println("Josh: It sortof gets a lot of hate due to the fact that most people just see it as a ripple clone..");
			Sleep.main(1600);
			System.out.println("Which I actually understand, and even agree with.. but I wish that people wouldn't spread so much hate about it..");
			Sleep.main(500);
			System.out.println("Oh well.. This isn't the place to vent about that kind of stuff anyways..");
			Sleep.main(600);
		}
		System.out.println("Have you done anything interesting recently?");
		interesting = scan.nextLine();

		if ("y".equalsIgnoreCase(interesting.charAt(0))) {
			System.out.println(name + ": " + interesting + ".");
			Sleep.main(400);
			System.out.println("Oh dude, tell me about it!");
			String blabla = scan.nextLine();
			Sleep.main(500);
			System.out.println("Wow.. That's super cool! I hope I can do something like that someday..");
			Sleep.main(500);
		} else {
			System.out.println(name + ": " + interesting + ".");
			Sleep.main(400);
			System.out.println("Josh: Oh.. Well that kinda sucks.. I haven't done much either I guess..");
			Sleep.main(400);
			System.out.println("Josh: Well.. At least other than the osu! server. But that wasn't even me really. I really owe it to Howl and Nyo.");
			Sleep.main(1000);
			System.out.println(name + ": Well.. Why not do something for them then..?");
			Sleep.main(2000);
			System.out.println("mmm..");
			Sleep.main(1000);
			System.out.println("Josh: I'm really now sure what I could even do.. One of their developers doesnt like me at all\nand I'm not sure what the other one thinks..");
			Sleep.main(4000);
			System.out.println(name + ": Oh well Sucks lole");
			Sleep.main(100);
			System.out.println("The End");
			Sleep.main(1000);
			System.exit(0);
		}
		System.out.println(name + ": I'm sure you'll do alot more in your life. haha");
		Sleep.main(500);
		System.out.println("Josh: Yup.. Most likely. I'm just waiting for the day where I come out of my shell.");
		Sleep.main(1500);
		System.out.println("Josh: So.. Out of curiousity, are you any good at osu..? What's your rank?");
		int rank = scan.nextInt();

		if (rank => 100000) {
			System.out.println(name + ": " + rank + ".");
			Sleep.main(400);
			System.out.println("Josh: Oh, so you're still pretty new to the game then!");
			Sleep.main(1200);
			System.out.println("Josh: I remember when I was that rank.. I first got my tablet right around 250k, haha");
		} else if (rank > 10000) {
			System.out.println(name + ": " + rank + ".");
			Sleep.main(400);
			System.out.println("Josh: Oh. So you've been playing for a bit, I guess. Still pretty new though, haha.");
			Sleep.main(1200);
			System.out.println("Josh: You'll probably start to see a slowdown in progress around the rank you're at now though. That's when it starts to actually get difficult..");
		} else if (rank > 3000) {
			System.out.println(name + ": " + rank + ".");
			Sleep.main(400);
			System.out.println("Josh: Oh damn.. You're catching up to me.. I got banned at about 2.9k, but that was after 2 cheated plays.. In reality i'm probably around there now.");
			Sleep.main(1200);
			System.out.println("Josh: Best of lucking kicking my ass, haha.");
		} else if (rank > 1000) {
			System.out.println(name + ": " + rank + ".");
			Sleep.main(400);
			System.out.println("Josh: Oh damn, you're a bit further than I am then..");
			Sleep.main(1200);
			System.out.println("Josh: We might actually be able to compete, I'll be sure to try to snipe some of your scores, haha");
		} else if (rank > 100)  {
			System.out.println(name + ": " + rank + ".");
			Sleep.main(400);
			System.out.println("Josh: Oh wow.. You're really good damn. I hope I'll reach that sort of level one day..");
			Sleep.main(1200);
			System.out.println("Josh: Not sure if I can, but I think it's a good goal.");
		} else if (rank > 35) {
			System.out.println(name + ": " + rank + ".");
			Sleep.main(400);
			System.out.println("Jo�h: ?  ????????? ? ?       ????????????  ?  ? ? ???\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? �?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? �?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? ?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? ?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? ?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? ?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?");
			Sleep.main(200);
			System.exit(0);
		} else if (rank > 20) {
			System.out.println(name + ": " + rank + ".");
			Sleep.main(400);
			System.out.println("Josh: Bro im finna smash ur ass like grass at ANIME NORTH CON 2018 :^)");
		} else {
			System.out.println(name + ": " + rank + ".");
			Sleep.main(400);
			while (count < 100) {
				System.out.println("Jo࿊hs�h: loO  Le h�e SDlo i am   coo   kiez i xDD℻DS�႟DDSD��᥀᥀℆℆℆D");
				count++;
			}
			Sleep.main(200);
			System.exit(0);
		}
		Sleep.main(600);
		System.out.println("Josh: Anyways, I've got to go home.. Do you use discord by any chance?");
		String haveDiscord = scan.nextLine();
		if ("y".equalsIgnoreCase(haveDiscord.charAt(0))) {
			System.out.println(name + ": " + haveDiscord + "!");
			Sleep.main(600);
			System.out.println("Josh: Oh sweet, whats your ID? Mine is cmyui#5585.");
			Sleep.main(850);
			System.out.print(name + ": Mine is: ");
			String discord = scan.nextLine();
		}
		System.out.print(name + ": Yea, ofcourse! My ID is: ");
		String discord = scan.nextLine();
		Sleep.main(600);
		System.out.println("Josh: Awesome. I'll make sure to add you when I get on!");
		//Unfinished
		}
	}
}
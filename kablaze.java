import java.util.Scanner;

/*
* Author: cmyui#5585 // https://akatsuki.pw
* Hey! Honestly I have no idea what im doing here but it's just something im doing to learn java. It's pretty cringy and thats basically the goal of it.
* Enjoy getting as far as you can before the game ends on you :^)
*/

public class kablaze {
	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		String playOsuString, akatsuki, interesting;
		int careAmt;

		System.out.println("??: Hey dude! Nice to meet you! I'm Josh, what's your name?");
		String name = scan.nextLine();
		System.out.println(name + ": Hey.. I'm " + name + ". Nice to meet you too!");

		if ("Josh".equalsIgnoreCase(name) && "cmyui".equalsIgnoreCase(name)) {
			paws(1400);
			System.out.println("J̸̢̨̨̦͚̠̺̱̜̙͙̣̗̱͔̝͍̮͕̠̱̥͎̮̟͇̞̾͂͌̈́͆̈̍̄͒̓̈̈́͐̓͛̐̈́͗͆̇̑̓̽̽̑̊̊̚ͅo̴̧̨̧̼͕̫̭̭̫̞̙͎̯̱̩̩͚̦͉̱̞̜̲͍̮̮͙̥͆͑͛̓͌́̉̏̎͐͐́͂̎̐̍̋̀̀́̇̈́̀̓̚̚̕͝ş̶̢̛̛̯̯͎̭̤̦̲̩̤̜͔̳͚̫̜̪̯̜͇͍͍̤̯̙̺̈́͂͗̒̈́̎̿̎͊̈́͑̈́̀̆́̽̈́̃͋̔̀̉͛͒̃̂ͅh̵̨̗̝͙̞̭͖̜̼͕̞͎̲͔̟̝͓̲̭͔̺̻͕̲̟̹̓̈́̉̊̄̈́̏̇̄̊̓̅̊̀̋̍̽̒̄͘̚̕̚̚͜͠͝͠͝ͅ:̴̢̺̘̩̩̩̠͕̬̳̼̘̩̬̺̱̯͓͔͓̣̗̜̲̱͉̣̂̃͂̍͒̏̋̇̅̓́͛͋͊̅̎̽̂̒͊̀̉̄̅̐͘͝͝ͅ ̸̛̛͓̤̩͕̗̠̬̗̮͉͈̫̖̻̠̲̘͇̠̪͓̙̭͉͖̙̳͎̄̍̓̐̾͗̾̎͆̌́̈̀̽́̀̆͑̌̓̎̆̽̑͘͘F̵̧̨̨̞̦̥͎̦̙͈̠̜̦̠̲͓̹̺̳̗̗͎̩̫̪̦̘͐̌̽̐͒̔͒͌̄̆̍͌̈́̅͒̆̍̈́̉̋̍̆̔̑̕͠͝͝ͅƯ̸̡̢̧͚͖͓̦̯̜̩̘͔̜̗̠̣̩̰̣̤̻̲͕͉͖͛̋̈́̄̂̈͗̄̂͂̔̇͊̇̐́̀̓́̇̆̐̆̕͘̚͜͝ͅͅÇ̶̨̻͖͉͓̻̣̹̺͔̥̠̜͚̤͓̤͈͕͔̲̞̤͖̼͂̈̇̓͐̂̄̉̉̎̊͗̄͋̀̈́̇̃̆́̎́́̍̆̚͜͠͠ͅĶ̵̧̨̳̺̭̤̘̯͕̳̝͙͈̩̝͖̼̦̫̥̜̭̫̞̰́̀̒̀̿̆͊͑̍̐͐̄̽͊̿̀́̈́͗̐̋̓̒͑̆͘͜͝͝ͅ ̷̡̲̼̜̪͓̞͈͙͎̦͇͍̙̘͚̦͈̝̬̺̤͚̘̖̠͔͙̾̓̾͑̏͆̈́̉̀̑̈̏̅͐͊̆͑̌͌̌̈͆͂̔͌̚̚͝ǭ̵̧̡̝̬̰͕̙͎̖̯͇̟͈̤̩̲̣̥̣̠̫͈̭̘͙͒͊̀̏͐́̎̈́̄̀̒͆̒́̋͆̈́̄̈́̑́͘̚̕͝͝͝ͅͅf̴̛̛̲̪̦̟̰̟̺̟̝̯̳̪̹̥̮͇͉̳͇̲͓̺̮̳͉̠͋̆̽̑̋̈̅͂̓͋̒̎͊̀̐͌̏͊͊̿̓́̚͘̚͜͠ͅ ̶̢̡̢̡̝̭̘̠̼̬̪̮͙̣̣̫̗̼̪̫̗̲̲̖͍̲̦͐̆̂̍͆̓̓̓͛̿̒̏̽̔̉͐͂͂̏͑́̔̀̑̏̚͝͝ͅf̶̡̛̲͈̮͕̩̫̰̻͇̯̱͖̪̮̤̪̤͕̱̠̝̮̰̤͇͇͆͐͛̾̇̑̀͑̒̌̅̀̇͌̓̓̍̓̑̈̒̾̃͂͘̚͝ͅf̵̡̢̡̢̨̧̛̲͖̘̝̼͔̖͕͕̝̥̝̞̣̮̱̼͕̩̻̜̈́̈́̽̅̈́̅̉̄̅̇̈́̃̄̍̈́́̂̂̌͑̽͒̐͗̕͘͘ͅ ̶̧̧̛̣̣̭̩̘͎̗̠͉͍͈̭͓̜͈̼̬̠͎͈͎̯͉̉̀̌̽̈́̑͛̈́͌́̌̀̈́̊̽͂̿̀̒͐̑͐̕͘͜͝͝͠ͅͅḒ̸̨̧̛̰̱͕̭͕̘͕̩̼̝͕̩̙̱̠̪̝͔̖̫̰͈̾̅̔͊̊̈́̒̔̈́̉̾́̈́̆̄̅̄̃̀̎̈́̄̽͂͘̚͠ͅͅͅi̶̡̨̢̛̩͕̯͔̫͍̗̯͕̥͖̟̟̤̫̟̞͙͙̬͕͍͚͛̂͌̃̈͗̈́́͗̐̀̈̑́̄̈́̓͊̾̓̑̓͐͘̕̕͝ͅͅë̶̢̧̡̛͉̘̳̩̬̩̗̳͖̮̤̫̝̝͖͙͙̞̱̞̹̥̱̯̬́̔̀̏̓̿̑̃̑̿͑̈́̇̐̍͊̎͗̿́̽̿̄̈́̈́͘͘ ̴̢̧̢̘̹̰͈̥̫͙̞͔͕̺͕̥͕͎̪̙̝̝̗̖̱̮̲̰̏̇͋̅͆̃̃́͐͌̓̄̎͗͛̂̈́̎͒̇̆̊̀̓͠͝͝͝á̴̡̨̧̡̛̞͖͕̪̫̭̥͎̱̱͙͕̲̞͍͎͇̜̞͔̫͍̤͙̐́̃̌̊͑̋͑͋̃̊̋̀̊̈́̆̊̈̈́̆͌͋̕̚̚͠ņ̸̢̛̥̰̙̖̝̤͔͎͖͔̭͙̝̲͇̗̖̘̖̝͇͈͙͓͌̾̐̐̃͐̀̽̂͆̈́͊͒́͛̽̔͗͗̔̈́͘̕͘͝͝͠ͅͅd̴̢̛͓̦̠̝͎̜̳̖̼̲͚͎͎̤͙̞̝̘̘̳͉̥̞͔̳͍̓̄̿̈̍̎̉̀́͒̍̋̔͆͐͆̉͑͑͒͗͑͋́͘͠͝ͅ ̷̢̨̢̨̡̧̘̝͍͓͓̮̹̦̠̟͎͔̦̮͓̭̖̹̣͕̰̹́̍͑̉̓̒͆̀͐̃̾͛̐͆̏̒̍̀̑̀̎͐̈́̀̓̕͠͝N̶̘̥̖͕̳̥͈̖͎̝̼̺͓̯̜̹̺͉͙̻͉͔̩͓̫̲̯͒́̋̈̄̀́́̇̋̌͐͑̎̊̐̆͑͂̉͂̂̈́͘̕͝͝͠ͅè̴̢̨̡̨̛̛̯̙̻̥̝͔̗̫̜̯̟̪̦͓͚͎̝̦̖̥̙̈̀̔̈́̋̒̒͂̆̓͊̍̎̐̅͑̈͆͋̒͑́̅̇͜͝ͅͅv̴̡̧̛̛̠̠̳̦͖̞͙͚̠̗̪̝̘̤̼̲̠̣̪̝̣̞̺̯̑̃́́̌̆͛̍̈́͒̈́͐͐̇̓͒̑̇̈́̿̉̉̈̓͘͝ͅͅĘ̸̢̧͓͙̪̮̱̞͔̠̱̰͕̜̱͎͇̫̩̱̯̜̼͓̹̅̂́̎̑̄̑̊́̐̑̿͗͋̅̐̉̉̀͂̂̅͊̑̅́̌̉͜͜R̷̡̛̛̺̮͉̞͎͔̻̭̮͕̞̫̻̪̻̖̱̠͇̹̟̤̜̖̈́̄͗͋̾̿̏̍̍̃̆̈́̂̑̆̂̾̈̐͂͑̓̆̋̔̃͜͜ͅ ̴̨̨̡̡̡̞̫̙̜̰͕̞͔̙̘͚͚̘̤̝̥̻͚̯̻̻͕̺̓͂͂̔͗̏̇̈́̿͐̅̉̿̐̈́̄̓̋̓̿͋̓̑̆̀̆͝͠R̸̢̭̳̙̺͔̫͚͇̪̟͕̥̟̻̣̬̗̻̥͔̯̻̼͉͍̩̊̎́̉̍͛̽̉̇̔̊̾̄̏̔̂̈́̀͌͊̈͘͘̕͘͝͠͠ͅĘ̴̛̠͙̦̱͎͔̣͙̻͙̗̙̟̫͕͔̘̯̦͔̹̥̼͚͓̣̱̀̓̑̈́͆̌̀̇́̈͒̐͛͌͌͌̓̉̅͒̑̈́̔̍̄͝͝T̵̡̧̨̧̨̜̩̺̲͍̰̮͉͈̹̬̗̲̠͉̗̟̙̟̹̰̥̭͑͆͌̅͗́̅́̈́͊̓̀̀͋͒͐̉̑͑̌́̎̃̿͑͛̏̾Ư̴̢̨̡̢̛̯̗̻̣̮͓̥̗̲̞͚̫͚̯̤̣̖̣͈͖̤̮̞̽̀̒̆̈́̂̽̔̇͊̄̐͗̌̍̌̆̌̈́̐̓̊͂̚͜͝͝R̴̠̰͖̱͍͚͉̙͔̮̺̮̹̭͍̰͕̫͚͈̠̺̘͓̦͉̾̏͗̇̾̓́̀͆̐̿̐̍̏̍̓̋̇́̐̈͆͊̚̚͜͜͠͠͝ñ̸̢̢̧̨͙̥̟̩̦͇͉̦̱̜̪͔̗̣̜̯̗̫̹̹̜̐̐̋͛̍̔̍̑̃̑͆́͊̀̎̌̈̇̔̈́̓͋̕͘͜͠͝͝ͅͅṟ̶̡̨̛̤͖̱͕̬̹̣̪̟͓͎͓̦͓̮̘̬̯͕̗̪͎͍̤̓͋̈́̓͛͌̆̑̑͑̃̔͐̈́̉̈́̀̂̐̿̌͘̚̕̚͝͝ͅ");
			paws(1600);
			end(0);
		}

		paws(1000);
		System.out.println("Josh: Mm.. Have you ever heard of a game called osu?");
		System.out.print(name + ": ");
		playOsuString = scan.nextLine();

		if ("y".equalsIgnoreCase(playOsuString.substring(0, 1))) {
			paws(200);
			System.out.println("Josh: Oh wow, really?");
			paws(400);
			System.out.println("Josh: Do you actually play the game?");
			System.out.print(name + ": ");
			String playOsu = scan.nextLine();
			if ("y".equalsIgnoreCase(playOsu.substring(0, 1))) {
				paws(300);
				System.out.println("Josh: Oh holy shit! You're the first person I've met that actually plays the game!");
				paws(600);
				System.out.println(name + ": Haha, yea not a lot of people know about it.");
			}
		} else {
			paws(200);
			System.out.println("Josh: Aw.. Damn.. I really want to meet another person that plays the game..");
			paws(1000);
			System.out.println("Josh: Anwyays.. I've got to head home. See you around!");
			paws(1600);
			end(0);
		}

		paws(500);
		System.out.println("Josh: So, since we both play osu!.. Whos your favourite player?");
		System.out.print(name + ": ");
		String favourite = scan.nextLine();

		if ("kablaze".equalsIgnoreCase(favourite)) {
			paws(400);
			System.out.println("Josh: Holy shit no way!!! ME TOOOO!!!!!!!");
			careAmt = 5;
			paws(4000);
		} else if ("cookiezi".equalsIgnoreCase(favourite)) {
			paws(400);
			System.out.println("Josh: Yea dude he's pretty good aswell");
			careAmt = 2;
			paws(4000);
		} else if ("howl".equalsIgnoreCase(favourite)) {
			paws(400);
			System.out.println(":^)");
			careAmt = 2;
			paws(4000);
		} else if ("cmyui".equalsIgnoreCase(favourite)) {
			paws(400);
			System.out.println("Josh: Haha.. Thanks man. <3");
			careAmt = 3;
			paws(4000);
		} else {
			paws(400);
			System.out.println("Josh: ..");
			paws(1800);
			System.out.println("Josh: Who..?");
			paws(3600);
			System.out.println("Josh: Oh well.. whatever.. It's not like I cared anyways");
			careAmt = 1;
			paws(4000);
		}

		if (careAmt == 1) {
			System.out.println("Josh: Actually what the fuck? How do you not appreciate any good players?\nI literally have no idea who you're even talking about. What the hell is wrong with you??");
			paws(300);
			System.out.println("Josh: Are you mentally insane? I don't understand your thinking");
			paws(300);
			System.out.println("Johs: Yea.. You're probably insane. You should go get that checked out..");
			paws(200);
			System.out.println("josh: Dude go to the hospital already what the hell are you doing");
			paws(150);
			System.out.println("Josh: You're clearly stupid and delusional if you don't think kablaze, cmyui, or cookiezi is the best player.");
			paws(120);
			System.out.println("J̢̙̮̙̮̬͍̰̺̗͚̥̠͉̹̙͕̹ͅͅo̢̢̨̧͓̗̩̫̙̠̪̪̯̯͙̱̬̼͍͜Ş̧̡̨͍̘̖̭̭̗͕̫̭͉̰̥͎͈̻͈h͙̖͇̘̱̦̩̱̞̝̞̬̯̣̹̞̰͉̮̯:̧̡͇̱̬͎̣̠̖̞̰͙̝͍̱̟̝̺̰͙ ̢̰͕̟̪̹̯͍̥͍͍̤̺̫̱̟̟͚̰͜W̡̡̨͕̺̬̠̦̮̼̼̭̭̳̗̯͕͉̥͜Ḩ͍̬̠͖̲̝̼̗͙̤̞̤̗̹̰̩̣͜͜Y̧̧̢̧̡̧̹̲̖̺̩̲̞͙̹̬̭̦̣ͅ ̡̢̫̺̥̦̘̻͔̦͔̙̣̝͔͓̫̭̤ͅA̢̡̧̡̼̰̬̗͚͔̠̜̺̲̝̯͕̫͓̭�̧̨̮̲͈̭͚̗̗̗̠̺̱̳͙̣͇̼͖̳�̡̧̧̨̙͉͇̲̗̣̹͙͇̤̜̣̖̟̩̠�̢̠͉̯͇̰͓͔̞̯̺̥̝͉̤͍̳̞͔̮Ŗ̡̨̺̭͇̼͖͉͎̹̫̮͍̫͕̝̬̫͜Ȩ̞̮̭̩͇̟̥̥̠̖̟̲̮̖̰͔̗͜ͅ ̡̢̪͉̤̫̱̝̦̺͖͍͇̫̱͈̫͕͓͜Y̨̡̟̤̹̯̮͈̤͙͈͓̗̤̮̳̮͚͓͈O̢̢̫͍̲͇̘̦̮̘͙̙̼̳͙̤͙̞͕̹U̡̥͇͔̞̝̝̩͇͍͔̰̭̜̯̫̯͇͜͜�̢̧̢̧͎̻̘͚̣̻͎̮͇̟̖̺̯͖̭̟ ͙͖̤̹̻̳̭̠̳̜̳͇̗̙̼̼̩̞̺̲S̡̳̼̩̱̘͎̠̜͖̣̘̩͖̻̠̱̳̞ͅŢ̡̯͚̰̤̯̮͚̫̹̦͚̩̰̟̻̻̩͜I̡̖̱̜͈͉̪͓̤̲̙͇̙̜̝̮̳̠̮̱L̨̢̨̢̡̞̠̫͉͕͈̩̬̘͉͚̠̻̯͜Ļ͕̪̳̺̲̹̤͓̰͖̥̹̪͔̺̲̣͓ͅ ̨̻̗̭̭͇̘̼̹͍͙͖̬̤̞̼͕̻̪ͅH̢̡̧͙̦̻̤̫͙͎͈͔̗͓̭̙̲͍̺͜E̢̠͔̬̳̙͖̰̭̺̪̗̠̲̘͚̞̗̭͜Ã͕̦̯̜̺͉͎͇̜̟͔͎̬̞̘̹̞̘̥ͅ³̢̯̘̼̯̤̭̬̫̼͙̹͙̬͎̳̟̫̲͓ķ̫̫̞͈̳̜̦̳̤̺͕̪̥̲̟̰̹͜ͅǫ̡̙̣̯̤̘͔͙̙͎̹̪̞̻̟̬̣͙͜ŗ͚̥̙̘̪͓͉͉̦̘͓̻̝̞̺̘͔ͅͅR̢̢̨̢̭͍̙̜̼̲̮͚̬̙̦͙̝̱͓͜Ȩ̮̥͇̤͉͍̦̞̹͉̞̹̳̤̣̫̜̻̬?̢̝̬̹̣̝͓͍̹̱̺̩̪̺̤̭̗͉̩̞?̨̡̙̮͖̞̟͙̺͖̩̭̻̯͓̖͉̺̦͜");
			paws(80);
			System.out.println("jalapeÃ±o: L E A V E");
			paws(20);
			int count = 0;
			while (count < 20) {
				System.out.println("�̡̡̡̢̧̨̧̢̨͎͎͕̙̣̪̞͙̲̰̭̞̮͕͇̪̮̫̮̯̮̜͎̦̫̗̪̩̥͔̪͈̮̞̦̬̭͚̟̖͕͎̥̖̞̱̠̖̺̠̻͖͔̮̬̦̪͙̦͇̲̖̫͈͙̲̹͎̭͖͖̣͍̱̹̣̯͎̬͕̟̱̰̺̜̣͙̫͜ͅͅJ̨̨̢̨̡̢̨̡̧̡̨̢̺̼͉̼̜̠͍͎͖̳̰̮͕̠̬̱͔̮̺̦̝̺̺̭̙͎̺̦̘̺̖͙̻̰̰̟̣̮̪̻͙͇͈̝̤̩̼̟̬̯̠̖̰̝̤̩̞̣͚̺̤͚̻͓̖̲̤̩̝̥̖̼̻̥̪̼̲̼͙͕̗̯͍̻̹͜͜ͅ ̧̨̨̡̡̢̡̧̨̢̡̢̢͙͈̝̫̮̥̼̺͔͚̖͍̲̖͍̜̟͈̞͈͉̻͙̮̱̯̲̣̪͇͎̤̣̞̬̘̖͓͎͉͕̬̰̞̠̼̗͍̣͍͉̫͓̣̻̦͕̬̟̯̦͚̩̣̝̮̝̰̯̭͎͚̜͖̬͖̱͖̻̠͎͜ͅͅͅͅͅ௸̡̧̨͖̩̳̪͎̬̬̱̤͇̼͉̖̬̤̠͎͙͖̼̪̯̟̩͍̦͉̗̤͚̫̗̣̬͚̬̪̜̺̹͍̳̥̜͕͔̟̠̙̪̳͔̙̞̫̟̥͚͔̟̤̥̪͔̦̫̩̘̼̩̬͕̗̺̺̩̭͔͈̠̭̟̪̻͇̯̼͍̰͙̳̗̠͜ͅͅͅ�̨̢̧̢̢̨̢̢̢̢̭̱̫͉̘̮͙̖̭̥̬̝̘͎̪̩̪̤̯̦̝̪̺̦̤͖̘̗̠̻̠̜̻̳̩̝͔̬͓̬̭̱͚̬̘̗̖͓̖̯̪͓͎͚̰̫͈̼̠̫̖͙͎̤̻̼̮͇̘̖̗͔͎̹͙̘̳̟̩̗̩̩͙͜͜͜ͅͅͅͅ ̡̧̢̢̢̡̡̧̧̨͔͖͍͖̩̤͇̩͍̥̜̥͓̬̗̯̖̝͔̪̲͚̦͉̦̘̘̤͚̭̟̯̯̦͍̩͎̺̜̟͙̣̝̯̦͖̮̪̞̤͉͇͉̜̳̠̻̺̻̭̩̤̘̥̖̗͍͍̮̘̣̭͇͖̱̻̪̩̥̗̗̘͜͜͜͜ͅͅͅͅ ̨̧̨̡̡̨̠̗͕̦̖̻̱̮̗̮̯̭͚̞̝̪͇̺͉̭̤͎̣̥͇͚͔̼̺̘̖̣̙͍͇̖̖̻͈̪͉͉̗̥̫͔̗̼̻͈͕̩͇͖̳͍̫̤̦̥̩͙͕͉͔̝͚̱̣͖͍͎̮̺͙̠̗̮̘̮̜̻̩͈͎͜͜͜ͅͅͅͅͅͅ ̡̡̢̨̧̨̡̪̫̗͕̖̠͚͖͈̟̩̞̘̲͚̭̥͇̙̻̜̲̰̜̻̦͇̦̻̠͚̦̩̘̩̟͖̤̯̯͓͖̘̪͚̦̦̘̜͍͓̠̤̲̫̩̰̞͈̫͎̭̣̮̫̹̯̺̯̩̝̗̪̻̗̺̜̟̞̘̱̭̟̺̣̱͔̙͈̯̗͜ͅ ̧̨̨̨̢͙̘͇̪̟̤̗̱̩̹̳̻͉̣̙̪̫̣̪̝̬̥͔̘̥̳̱̜̘̳̤̺̘̟̙̣̪͔̞͔̠̪̭̹̝͍̜̰̫̰̳̭̣͕̗̜̣͙̹̘̼̲̫̗̰̤̗͖̗͚̝̺̜̪̥͍͓̜̼̠̱̻̯̙̺̠͈͙̹͎͍͜͜ͅͅ�̨̨̢̧̧̢̧̡̡̨̡̢̡͔͍̩̟̤̳̟̣̤̬͈̗̘̞̬͍̭̫̮̘̹̣͖̬̮͕͈̻͙̹̭̦̣͎͖͎̹̹̹̩̼͕͔͎̗̪̜̝̼̺̥̥̩̹̯͈̜̻̟͉̦͉̻̙͚̱͚̺͕̫̭̩͓̱̣̖̟͖͇͇̪̝̞͉͔̖ͅ�̢̧̢̢̢̧̢̙̖̹͙͈̳͎͎͈̭̹̞̙̩͇͇̝̬̥̠̺͖͎̣̹͚̞̥͙̭̜̮̞̖̩̦̖̺̺̱͓̞̺̦͙̠͚̰͎͔̙̘͉̦̳̼͎̥̬̯̜̗̩̟͔̦̖̯̤̗̗͚͈͔͎̣̭̳̫̗̪̖̯̤͔͉͈̲̳͓ͅͅͅ�̧̡̢̨̡̨̢̨̢̨̡̢̡̧̧̝̣͈̩̥͙͎̭̺̣͔̯̻̠̝̻̦̦͉̙̻̝͙̗̥̹̹̼̜̥͎̜̤̖̮͇͈̗͈̟̘͎̹̠̮̬̻̦̳̗̩̯̬̲̥̪͚̫̤͕̗̱̩̜͇̺̼̙̜̥̦̼̫̰̞͖̱̥̼̣͜͜ͅͅͅ ̧̡̡̡̡̧̡̡̨̢̼̯͔̮̯͕̯͔̫͍̜̹̪̱̜̘͕̩̯̯͉̺̱͈̬̗̪̖̬̱̥̦͕̭̥̰̦͚͍̱̜̦̦̳̤̩̣͚̲̘̣͇̳̝̙͈̲̻͉̹̼͔̝̲̗̲̫͚̙̠̭̹͕̜̹̳̲̠̹̯͉̘̰͇̺͈̼͜͜ͅ ̨̢̨̦̻͚̮̜̥̪̮̰̙̼̙̼͖̬̦̖̮͇̻̙̦͕͖̤͓̤̲̱̺͈̙͚̳̝̘̰̘̮͇̼͈͖̙͚͍̩̲̬͓̗͚͖̻͉̭̠̯̝̥͙͈͎͔͕̻̱͔̝͇̟̗̗̪̟͈̠̼̲̼̹͇̗̺̘͍̲̩̝̮͜͜ͅͅͅͅͅ ̨̧̢̢̡̢̢̢̧̢̡̢̡̟̫̜̱̮͙̲̰̦͚͎͈͉̹̜̗̟̥͕̳̖̗̳̞̫̻͕̳̺̼͔̬̯̞̘͖̼̠̬͔̯̗̹̙̲̰̬̫͎̣̩͎̝̦̯̣͚̼̤̲͖̥̮̜̻͔̮͔͖͙̱̥̘̯̻̤̟̻͜͜͜͜͜͜͜͜ͅ.̡̢̢̡̧̨̧̡̢̜͖̼̭̗̞͚̯̥̝̩̜̘̙̹̳̗̞̰͖̰̬͔̥̯͇̪̫̫̥͚̗̠̜̹̦̯̠̱͉̫͇̪͎͚̹̣̰͓͚̭̫͉̲̖͙̜̼͖̪̤̘̳̠̞̬̣̺̞̺͎̦̠͈̹̤͙̜̖͚̳̞͖͍̘͔̘͓̤͜ͅ");
				count++;
			}
			paws(60);
			end(0);
		} else if (careAmt == 5) {
			System.out.println("Josh: I'm actually really suprised you think kablaze is the best..");
			paws(700);
			System.out.println("Josh: He's usually a super underrated player so it's really nice when I can find someone\nThat can appreciate what he brings to the table..");
			paws(950);
			System.out.println("..");
		}

		paws(1000);
		System.out.println("Josh: To be honest, i'm amazed you've made it this far..");
		paws(300);
		System.out.println(name + ": What?");
		paws(500);
		System.out.println("Josh: mm.. Nothing.. It's almost like you're trying to read my mind is all..");
		paws(600);
		System.out.println(name + ": Uh.. Sorry I'm not sure I know what you're talking about..");
		paws(600);
		System.out.println("Josh: heh... Of course not..");
		paws(3600);
		System.out.println("Josh: Have you ever heard of Akatsuki?");
		paws(600);
		System.out.print(name + ": ");
		akatsuki = scan.nextLine();

		if ("y".equalsIgnoreCase(akatsuki.substring(0, 1))) {
			paws(200);
			System.out.println("Josh: Yea.. It's been getting quite a lot of attention recently..");
			paws(400);
			System.out.println("I'm actually the creator of it.. haha");
			paws(400);
			System.out.println(name + ": Oh dude, thats kind of cool.. although.. isnt it litrally just ripples code..?");
			paws(1400);
			System.out.println("Josh: �.");
			paws(6000);
			System.out.println("Jos�̡̢̡̝̥͇̠̗͖̙̹̥̩̣̻̞͎̜̹͇̻͚̖̭̖̟̼ͅh̨̢̨̨̭̞̖̲̝̭̠̜̳͔̱̫̞̼̠̖͎͎̠̰̦̦̬̭:̡̺̰̲̣̲̼̲̖͔̗̞̲͇͈̖̠̯̜̲̮̦̼̹̻̠̱ͅ ̡̧̢̢̟̩͈̞̳̙͔͔̜͓̼̥̭̝̘͔͓̝̻̪̯̯̹ͅI̧̻̯͍̜͓͈̹͖̰̬̪̹͇̱̝̣̥̠͙̣͉̲̠̺̠̲ͅ'̨̧̢̢̡̨̡̙̟̙̳̥͓͎̼̩̰̠̟̞̗͍̺͙̪̬̝ͅm̢̡̧̨̩̜̠̝̱͕̲̦͇̰̬͎͙̠̱̼̹̯̥̘̰͈̥̣ ̡̡̹͓̥̙̣̙͈͈̘̠̣̰̞͈͎̲̩̣̼͎̝̳̣͍͜ͅs̢̢̫̜̘͖̞̫̩̩̟̻͔̻̙̼̜͓̭̜͔̬͕̣͜͜ͅͅo̡̢̨̲̼͉͇̟͖̯̰͉̦͍̻͓̜̻̫͇͍͖̮̰͍̯͚͜ŗ̥͉̜̩̰͉̩̼͖̮̱̻͍̩͎̮͔͕̲͈̹͇̦̪ͅͅ    rRrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr-");
			paws(100);
			end(0);
		} else {
			paws(200);
			System.out.println("Josh: Oh.. Yea it's an osu! private server I made using Ripple's source as a base..");
			paws(500);
			System.out.println("Josh: It sortof gets a lot of hate due to the fact that most people just see it as a ripple clone..");
			paws(1600);
			System.out.println("Which I actually understand, and even agree with.. but I wish that people wouldn't spread so much hate about it..");
			paws(500);
			System.out.println("Oh well.. This isn't the place to vent about that kind of stuff anyways..");
			paws(600);
		}

		System.out.println("Have you done anything interesting recently?");
		System.out.print(name + ": ");
		interesting = scan.nextLine();

		if ("y".equalsIgnoreCase(interesting.substring(0, 1))) {
			paws(400);
			System.out.println("Oh dude, tell me about it!");
			System.out.print(name + ": ");
			String blabla = scan.nextLine();
			if ("no".equalsIgnoreCase(blabla)) {
				paws(200);
				System.out.println("ẅ̮̙̪͚̭̪̳͓̩́̐̈̎͊̌̃̊̀͗ͅĕ̮͉̖̫̣̱̦̰̜̽̑̿́͆̿̈͂͠ͅl̪͓̘̲̙͓̠̱̟̎͑̎̌̿͑̎͑̇̕ͅl̡̨͎̭͔̖̹̺̻̗̅̑̉̌̂́̄̒̎̕ ̢̗̞̻̟͓̘͖͎̲͆̒͌́͆̎͛̂͆͝f̹͙̗̲̝̮͎̮̜͎̆̎̉͌̊̈̃̍̎̚u̡̜̩̗̺̺̱̺͖̗̔̍̈̓̈́͒̎͛͝͝ç̬̪̟̟̤͙̹̟͔̽̋̀̈́͛̾͋̿̈͝k̦̥̩̬̞̗̟̳͎͇͒̓͗̊̾̿̒̈́̆̊ ̢̜͚̳̙̖͉͇͇͂̑̂̇̄͆̌̉̈́͘͜y̹̰͍͍̙̲̰̞̳͉͒̐̓̆̉̓̆̆̚͝o̞̫̻̗̖͎̬̝̝̽̐̈̆̔͑̾̎͘͝ͅu̠̞͕̮͎̫̦̻͈͔̐̆͊͋͛͛̎̔̏͝ ̨͙̳̝̭͚̰̲̰̘̋̾͐̌̅̈́̈̃͆͘ẗ̛̜̝̠̖̪̬̹̗͍̺͂͆͛́͂̒͊͘ơ̧̹͚͇̯͎͕̳̘͍͌̓̑̉̐͐̈́̊̋ǫ̧̡̟̟͇̥̯̠̱̈́͗̌͗̈́̒̍̉͝͠")
				end(0);
			}
			paws(500);
			System.out.println("Wow.. That's super cool! I hope I can do something like that someday..");
			paws(500);
		} else {
			paws(400);
			System.out.println("Josh: Oh.. Well that kinda sucks.. I haven't done much either I guess..");
			paws(400);
			System.out.println("Josh: Well.. At least other th�n the osu! server. But that wasn't even me really. I really owe it to Howl and Nyo.");
			paws(1000);
			System.out.println(name + ": Well.. Why not do something for them then..?");
			paws(2000);
			System.out.println("mmm..");
			paws(1000);
			System.out.println("Josh: I'm really now sure what I could even do.. One of their developers doesnt like me at all\nand I'm not sure what the other one thinks..");
			paws(4000);
			System.out.println(name + ": Oh well Sucks lole");
			paws(1000);
			end(0);
		}

		System.out.println(name + ": I'm sure you'll do alot more in your life. haha");
		paws(500);
		System.out.println("Josh: Yup.. Most likely. I'm just waiting for the day where I come out of my shell.");
		paws(1500);
		System.out.println("Josh: So.. Out of curiousity, are you any good at osu..? What's your rank?");
		System.out.print(name + ": ");
		int rank = scan.nextInt();

		if (rank >= 100000) {
			paws(400);
			System.out.println("Josh: Oh, so you're still pretty new to the game then!");
			paws(1200);
			System.out.println("Josh: I remember when I was that rank.. I first got my tablet right around 250k, haha");
		} else if (rank > 10000) {
			paws(400);
			System.out.println("Josh: Oh. So you've been playing for a bit, I guess. Still pretty new though, haha.");
			paws(1200);
			System.out.println("Josh: You'll probably start to see a slowdown in progress around the rank you're at now though. That's when it starts to actually get difficult..");
		} else if (rank > 3000) {
			paws(400);
			System.out.println("Josh: Oh damn.. You're catching up to me.. I got banned at about 2.9k, but that was after 2 cheated plays.. In reality i'm probably around there now.");
			paws(1200);
			System.out.println("Josh: Best of lucking kicking my ass, haha.");
		} else if (rank > 1000) {
			paws(400);
			System.out.println("Josh: Oh damn, you're a bit further than I am then..");
			paws(1200);
			System.out.println("Josh: We might actually be able to compete, I'll be sure to try to snipe some of your scores, haha");
		} else if (rank > 100)  {
			paws(400);
			System.out.println("Josh: Oh wow.. You're really good damn. I hope I'll reach that sort of level one day..");
			paws(1200);
			System.out.println("Josh: Not sure if I can, but I think it's a good goal.");
		} else if (rank > 35) {
			paws(400);
			System.out.println("Jo�h: ?  ????????? ? ?       ????????????  ?  ? ? ???\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? �?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? �?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? ?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? ?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?\n???????????? ? ??  ?? ?    ????? ??  ?\n????? ?? ?? ? ?? ? ?        ?????????  ??? ?\n ?? ??          ??? ?     ?? ??? ? ?\n? ?  ? ? ??????? ? ?? ? ?? ? ?? ? ? \n ? ? ?? ?   ????? ?  ?? ? ? ? ? ? ?? ?");
			paws(200);
			end(0);
		} else if (rank > 20) {
			paws(400);
			System.out.println("Josh: Bro im finna smash ur ass like grass at ANIME NORTH CON 2018 :^)");
		} else {
			paws(400);
			int count = 0;
			while (count < 100) {
				System.out.println("Jo̧̡̧̨̨̧̧̢͓͓̼̘̝͇͈͍̼͓̦̟̝͉͍͍̬͈͔̜̥̜̱͕̲̲̗̥̖̻̦̠̠̹̣̣͖̞̺̻͖̤̪̭̳͖͈͚̦̳͔̥̰̜̪̩̱̗͚̞̯̘̮͕͔͙͍̭͙̳̺̱̥̣̗͖͙͎̥̣̭̞͇̦̠̘̥͜ͅͅͅͅ࿊hs�h: loO  Le h�e SDlo i am   coo   kiez i xDD℻DS�႟DDSD��᥀᥀℆℆℆D");
				count++;
			}
			paws(200);
			end(0);
		}

		paws(600);
		System.out.println("Josh: Anyways, I've got to go home.. Do you use discord by any chance?");
		System.out.print(name + ": ");
		String haveDiscord = scan.nextLine();

		if ("y".equalsIgnoreCase(haveDiscord.substring(0, 1))) {
			paws(600);
			System.out.println("Josh: Oh sweet, whats your ID? Mine is cmyui#5585.");
			paws(850);
		} else {
			paws(600);
			System.out.println("Josh: Oh damn, well.. Maybe we'll meet again one day.");
			paws(1000);
			System.out.println(name + ": Maybe.. Well.. See ya!");
			paws(600);
			System.out.println("Josh: See ya!");
			paws(600);
			end(0);
		}

		System.out.print(name + ": Mine's: ");
		System.out.print(name + ": ");
		String discord = scan.nextLine();
		paws(1000);

		switch (discord.toLowerCase()) {
			case "cmyui#5585":
				System.out.println("J̢̢̧̨̨̡̨̧̨͓̱̹̝̣̳̥͎̯̫̣̭͇̗̮͚̼̬͈͕̫̲̩̮̠̖̗̺̘͉̝͍͙͇̼̥͖͙͇̖̬̤̘̭̠̱̺͖̮̖̘̹͓͈̩̲̞̹̪̪̞͕̗̺̬̤̫̟̹͇̦͖͎͓͉̻͍̩͈͓̼̫̰̠͜͜͜͜ͅͅͅo̢̡̨̧̨̧̡̨̡̡̡̢̭̼͉̭͇̞͔̻̪̤̼̯̣̼̥̮͉̣͚̩̯̩͖͖̜͇̖̥͚͙͎̟̪̘̻̝̥͖̖̬̖̩̪̦͎̝̩̜̤̫̳̤̹̙̩̭̘̳̖̥̠͎̫͈͈̙̣̬̺̤̳̩̟̰̳͍̮̝͓͙̰̦̗̟̬͚ͅͅS̢̧̢̨̡̧͚̰͔̰̮̞͕͇͉̪̝̼̟̜̠̟̻̬͎̠̲̘̭̹̗͓̺̱̹̤͎̤̲̣̳̯̺̖̪͙͚̖͈͈̩̯͓̻̯͚̲̮͇̲̳̞̬͙̠̖̘̣̹͍͈̤͍̗̗̥͚͎̞͔͙̟̤͖̭̪͚͈͈̙̥̠̦̫̱̥͜ͅͅͅH̨̧̡̡̧̠̫̤̹̞͙͉͈̻͓̬̩̲̖̜̰̦̥̬̙͚̥͙̥̹̠̣̝̖͖̰̜͎̮̬͓͔̠͉̭͎̯̣̻̰͔̖̩͕̜̰̤̭̩̪̥͚̘̦̖͎͚̝̬͉̠̺̼͍̻̣̟̗͉̳̱̩̰̯͇̬̠̘͙͕̖͖͜͜͜ͅͅͅͅͅ℆̨̨̢̨̢̨̨̨͙͙̘͓͖̜̹̯̥̪͍̫̗̝̬͖̬̞̰͖̭̝̮̫̤̫̦̩͙̝̞̟͇͉̟̹̤̙̞̺̖̤̳̲̣̟̹͍̦͙̫͚̣̜͕̜͕͔̙̤̥̻̝͓̝͓̙̞̹̺͇̠̟̬̬͎̰͖͙̝̫̫͎̟͓̩͍͕͜͜ͅͅS̢̡̡̨̧̢̡̡͈̜̩͓͍̰͚̪̰͓̲̦̺̦̬̜̝͍̙̜͕̥̟̼͈̣̟̥̼̤̤̤̮͎̱̻̻͈̹͎̘̺͉̞̯͎͈̘̹͔̼̭̳̖͈̞͇͉͓̹̣͔̭̪͔̠̞̫̯̥̼͉͍͉͍̠͚̮̞̻̼̖̦̞͎̬̘͜͜͜ͅͅ:̧̢̢̨̢̧̡̨̭̗̯͔̦̙͙̼͇͙̟̻̹͎̠̗̝̦͚̭̣̟̫͉̪͍̦̬̥͚̪̖̱͙̫̭̯̯̲̼̘̞̦̪̝̤̫͈̰̗̗̩͉͙̹̺̙̤͕̲̹̣̦͈͚͔̞̳̞̖̼͎͈̲͙̰̬̯͚̙̫͕̣̻̞̲̙̞͜͜ͅͅ ̡̡̡̡̨̡̧̡̻̫̗̫̟̼̫̣͉̩̯̯̺̭͈̦̺̙͖̗̺͙̣͚͕̥͖͚̩̪͇̘̜̤̘̮̳͇̠̱̯̦̬͎͓̝̼̹͙̣͙̳͓͖͙͈̭̤͖͍̘͈̠͕̤͈͎̙̦̩̯̜͔͙̜̦͖̻̪͙̙̲̟̱̝̝͜͜͜͜ͅͅJ̡̢̡̧͉̩̥͙̞̫̥̥̜̼̩̤̮̣̳̥̰̖͍̲̮̗̖̭͈̗̻̮͍̼̯̰̣̪̮̰͇̤̩̙͇̟͈̺̲̯̳͉̱̘̪̬͉̤͕͉̜̞̰̣͚͎̣͉̮̞̪̪͈̙̟̺̘̝̲̮̝̦̭̰̥͙̙̝̗̱̫̳̩̟̹̬̳͜͜ͅS̡̢̡̨̢̢̨̜̹̗͙̥̬͈̪͈͍͈̺̦͖͇̻̤͎͔̞̘̥̝̯̜̻̝̬̖͈̪̱̮͇̗̝͔̘͉͉̮͍̣͈̻̩͖͓̖̝̖̪͍̟̹̰͍̼͓̪̖̖̦̯̺͚̲̦̺̥̘̯̤͍̺͈̠̰͔͎͙̻̳̠̜̰̳̙͜͜ͅͅͅI̧̧̡̧̨̨̨̮̪̥̠̼͙̹̘̲̱͈͉̗̘͈̘̘͔͎̗̣͙̹͕̙͎̫͇͇̼͚̩͍̮̱̗̮̦̦̥̭̲͓̻͔̱̤̗̟̝̝͔̖̻̣̼̹͚̰̬̻̪̣̳̥̜̗͔̭̠̲͙̠̥̟͖̟͈̗̘͍̰̰͎̦͚̤̼͓̹͈͜ͅH̨̧̡̨̡̧̨̡̢̧̢̟͉͓̤̞̭̥̘̦̜͓͉̘̣͖͕̥̜̣͉̟̣̣̞͔͉̳̠͓͍̦̘̳̬͙̟̣͍̩͎̯͉̹͚̗͖̦͇͖͖̻̩͍͕̹̪̲̳͍̭̟̬̖͍͕̠͍̲̹͕̪̳̺̭̟̱͕͖̙̯̙͔̲̫͜ͅͅͅͅ ̧̧̡̢̢̧̧̡̡̨̨̡̧͉͕͚̘̞̭̻̫̺͔̠̜̝̭̟̬͓̝̘̯̙̬̰̞̯̫͔̼̼͈̰͖͍̗̯̳̗̮̗̣̹̗͉̻͕̟̼̦͔͇̩̝̥̹̖̰̥͎͕̮̗̜̭̬̻̹̙̖͇͔̱͙̙̗̮̜̜̗̰̟̥͕͜͜͜ͅͅą̨̨̧̨̢̡̧̨̡̧̨̡̨͈̼̹̟̭̣͈̟̰̣͔̘̲̘͚̫̥͈̲͓͈̝̺̳̪͕̗̻͈̱̠̩͚͙͍̜̦͓͍̜̝̹̙̺͕̺̻̝͔̺͕͚̤̟͖̪͉̤̘̳̟̬̥̞̠̱̝͎͉̫̠̱̮̹̺̻ͅͅͅͅy̡̢̡̢̢̢̨̢̨̢̡̧̢̧̧̻̣̜̫͈̥̗̭̱̼̗̪̯̹͓̝͇̙͇̞͈͓̘̣͚͔͍̺̲̳̦̤̖͔̼͓̝̣̩͖̬̗̰̟̜̭͙̤͉̜͓̟̭̗͚̗̫͙̻͖̙͍̝͙͕̟̯̫͖̣̗̳͈̰͖̖̣̬͜͜͜͜͜ͅͅf̢̧̨̨̨̡̧̜̙̤̼̯̰̟̗͕̭͚̬̹̻͈̣̱̱̹̦̪͔̮̰͙͔̩̟̬̫̱̤̜̜͈͍̰̹̩̻̦̖̖̣̬͇͈̩̳͔͖̣̭͚̰͉̙̩͙̜̗̗̯̩͙̹̤̠̘͇̜̣͚̫̺̤͈̯̯̬͓̹̳̻̝̤͚̘͎̞͜ͅͅg̨̢̡̨̧̨̢̡̨̧̡̡͉̝̭̼͔̝̤̼̻͎̣̯̘͚͉̪̦̞̭̻̜̩̩̱͎̭̹̳͙̫͕̦̱̠̙̪̦̻̠̦͖͚̺̳̬͎̬̩̜͍̙̬̖̪̳̹͕͖̱͎̪̻̟̭̹͎̱̗͚̝̣̣̩̜̹͕̺̜̬̱̰̼̰̤̞͜ͅͅş̢̢̢̢̨̯͉̺̙̰̘̟̫̦͓̩̦̦̻̮̖̗͔͈̠̺͓̹̥̦̫͙͇̯̟͔̥̥̺̞͚̜̘̲̝̖̘͕̣̯̻̳̻͇͎̭͎̣̳͚͎̼̙͔͖͚̖͙̜̮̫̤̲͍͚̮̮͇͜͜͜͜͜ͅͅͅş̨̡̢̨̨̡̨̨̧̙̗͖̪̖̪͔̹̥̱̜̻̞͙̩̪͓̺͔̻̩̪̙͔̺̭͈̻̝̫̘͓̦̗͚̹̜͔̳̳̺̬̳̪̟̠͍̥̪͍̜̖̰̹̭̥̲͈͈͔̹͖͕̭̙̯̤̳̝̠̫̭̗͔͍͕̳̬̣̠̗̣̱͜͜͜͜͜͜ͅy̡̢̧̨̨̧̡̯͎̝̲͓̳̪̭͈̬̹̱̞̰̮̖̪̥̥̮͕͉̺̠͕̞̣̟̞̺̫̟͍̞̥͔͕̹̙͓̺͕̬̟̼̱͙͕͇̪̣̘̞̯͍͚͎̘̜͓̪̻̠̫̻̮̝̖̬͚̫͙͈̘̟̲͖̜̱̬̯͚͍̠̪͚͓̼̲͜͜ͅͅb̨̡̢̨̡̢̢̡̨̡̢̡̥̦̫̫̤͉̤͓͕̲̩̭̠͖͍̹̝͙̹̥̱̟̹͍̙̻̥̻̳̠̟̙͕̮̮̪͈͔̰̰̮͍̪̻͕̗̱̠͈̤̪͇̳͔̲̭̗̜͖̬͕̝̤̳͎̝͙̩͚̳͕͔̙̹̫̤̻͚̱̜͇͖̯̰͈͜ͅͅF̢̨̢̨̢̧̧̣̪̜͙͖̱̺̞͙͙̼̫̲̞̭͔̳̫̦̲̰̹̥̟̹͓͈̜̝͉̫͇͇͉̭͎͕͍̦̘̳̮͕̙͚̫̰̰͓͇̤͕̫̰͉̬̦̮͕̪̖͕͚͕͍̖̯̮̗̮͓̩͖̖̘͚̣̟͍͓̦̘̪̞̦͖̪̤̹̲͜͜ͅB̢̢̡̢̡̡̢̧̢͍̙̠̦̦̳̙̭̯̹̹̗̗̻͕̻̫͓͎̪̙͙̪̫͕͕̭̯͓͈̱͍̞̦͖̬̙̻̜͙̳̖̼͖͖͔͙̲̮̩̱̟̺̹̻̤͓̠̹̘̬͕͚̥̝̦̝̯̯̼̹͕̤̺͚͖͎̫̤̖͔͎̦̦̫̯͔͜ͅͅͅS̨̡̡̧̨̨̢̡̭͇̙̲͈̦͉̹̝̦̝̦͓̻̙͍̺͉͇̘͇̣̭͎͔̭͚͔̯̝̦̹̰̮͙̱͉̜̘̫̺̯̻̣̻̺̯̠̪͕̳̖͚̦̬͙̮͙̦͓̼̰͓͕͍̩̪̗̞̩̝̣̥̮͍̗̜̲͇͎̯̳̖͎̞̞̣̭͚͇̞ͅF̨̨̨̧̢̧̢̧̨̢̠̺̯̩̟̭͈̰̬̻̖͍̼͍͍̱̭̭̞͖̮̪̙̖̻̩̫͍͔̟̝̯̮͈̺̜̳̥̮̟̺̤̩̹͙͉̞̖̭̥͚͚͓̦̹͎͈̥̻̰̤̱̻̺̙̗̝̥̣̱̱̮̝̜͍̺͔̣̪͎͕͕͜͜͜͜͜͜͜ͅḄ̨̧̨̢̢̡̢̡̨̢̡̢̧̢̭̮̟̲̪̘̯̝̫̘̣̲̖̪̙̠̼̦̭̯̪̥͉̹̬̩͍̺̩̲͖͚̲̬̙͈̦̠̞̖͔̙̗̭͙͈̲̘̠͓̖̖͍̲̤̹̫͇̫̗̗͈̲̖̪͕̥̤̙̭͍̬̣̼̞̠̯͇͙̤͈̙͜͜͜b̢̢̧̧̡̧̡̧̧̢̨̧̖̳̳̰̬̩̭̱̬̜̳̺͍͚͎̫͕̖̟̗͖͓̦̙̜͓̝̥͖̺͉̱͕̘̭͓̬̙̠̠̯̖͔̗̖̞͕͉̝̗͙͚̻̙̥͕̟̠̠͈̜̜̳̪̼̜̮̙̟͔̟̹͈̫̭͍̳̘͎̤̖͚̠̼͚͜ͅͅy̧̢̨̧̡̺̲͔̠̮̬̬͇̦̼͖̤͙̗̻͇̬̞̲͙̮̳̼̙̠̭̫̮̱̳͓͉̜̜̮̣̣̞̘͓͎͕̺̻͓̬͓̹̹̲̩͕̥͔͈̖̻̖̤͕͚͎̘̠̟̺͉͕̞̟͚̤͓͖̥̜̯̹͉̯̝̪͍͚̦̦̪̖̪͓ͅͅͅͅͅ\nN̢̧̡̢̡̡̡̟̟̰̪̩͇̝̻͔̪̖̲͔̮̰̞̫͕̙͚̮̹͕͈̩̼͕͚̗̗͚͖̖̻͈̟͖̠̥̝̳̦̝͙̩͚̼̜͉͔͓̺̣̬̜̱͇̤̙̲̖̰̳̼̮͔̱̹̬̱̥̰̞̥̰̻̰̣̟͖͚̰̺̜̜̟͎̼̯̹̙͜͜ǫ̨̨̨̢̡̡̨̧̢̧̟̙̰̘͎͈̮̟͕̜̱̥̞͕̙͖͉̬̞͎̟̣̫̱̼̘̭̝̙̲̗̭̠͍̹͙̬͔̭̰̖̯̞̩̩͇̲̜͈͔̪̮͎̮̬̱͖̱͓͕̪̜̱̪̖͈͙̣̜͍̣̰͈̱̼͓̹̩̙͜͜͜͜ͅͅͅͅͅͅ.");
				end(0);
				break;
			case "kahri#9828": case "Emily | Sunpy#5213": case "Hanyeol 한열#8257":
			case "Kip#1120": case "Frosti#0602": case "ninjin#3198": case "Howl#0940":
				System.out.println("Josh: Huh.. That sounds oddly familliar..");
				paws(1600);
				System.out.println("Josh: well, whatever.");
				break;
			default:
				System.out.println("Josh: Alright cool!");
				break;
		}
		paws(600);
		System.out.println("Josh: I'll make sure to add you when I get on!");
		paws(1600);
		System.out.println("Josh heads off towards his home, and I do the same..");
		paws(800);
		System.out.println("..");
		paws(800);
		System.out.println("..");
		paws(800);
		System.out.println("..");
		paws(800);
		System.out.println("..");
		paws(800);
		System.out.println("..");
		paws(800);
		System.out.println("")
		}
	//Red text & Reset
	public static final String ANSI_RESET = "\u001B[0m";
	public static final String ANSI_RED = "\u001B[31m";
	//Pauses inbetween text for realistic timing
	public static void paws(int x) {
		try {
			Thread.sleep(x);
		} catch ( java.lang.InterruptedException ie) {
			System.out.println(ie);
		}
	}
	//Function for whenever the player hits a dead end
	public static void end(int x) {
			System.out.println(ANSI_RED + "Game Over" + ANSI_RESET);
			System.exit(x);
	}

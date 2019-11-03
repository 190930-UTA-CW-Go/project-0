package managers

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
This file contains every encounter in the game. Each encounter is labelled with a number, and each returns a number.
This number corresponds to a different, unique branch of the story, similar to pages in an adventure book.
Therefore, whichever int the function returns determines which encounter will come next.
i.e if Scene1 returns a 2, in main.go, Scene2 will be executed.
*/

//Scene1 is the first encounter in the game.
func Scene1(p *Player) int {
	io.WriteString(os.Stdout,
		`You awaken. The world around you is pitch black and cold, and you find yourself unable to move as you struggle against the void.
Slowly but surely the darkness around you begins to fade, heat flowing over you as you began to hear the sound of bustling streets, 
the lively crowd around you passing by without a second thought of your presence. You knew you were somewhere, but you couldn't see a single thing.
You thought to yourself, and you knew exactly what you had to do.`+"\n",
	)

	a := 0
	attempts := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "open eyes":
			a++
			selection = 2
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			fmt.Println("Your eyes are closed.")
		default:
			if attempts == 0 {
				fmt.Println("Just open your eyes.")
			} else if attempts == 1 {
				fmt.Println("Type 'open eyes' into the console.")
			} else if attempts == 2 {
				fmt.Println("Really?")
			} else if attempts == 3 {
				fmt.Println("You're kidding, right? If you don't type 'open eyes' this time, I'm actually going to explode.")
			} else {
				panic("The narrator's face became a deep red and he promptly exploded, killing you in the process. Maybe you should've just opened your eyes. Game over.")
			}
			attempts++
		}

	}
	return selection
}

//Scene2 is the second encounter in the game, branching off of Scene1.
func Scene2(p *Player) int {
	io.WriteString(os.Stdout,
		`You opened your eyes and were partially blinded by the sheer intensity of the world around you. As your eyes adjusted, 
your mouth hung open with wonder at the sight of the most vibrant city you had ever seen. Tall skyscarpers were covered in colorful billboards, 
illuminating the world around you to such a degree that you could hardly tell it was late into the night. The people around you were dressed 
in an odd fashion, their style reminescent of some grungy, dystopian future you might find in a 1970s fiction. You looked down and noticed 
you were dressed very similar to them, and to your surprise, your left arm was completely cybernetic. You shuddered as you imagined what your hair might look like 
as you began to walk about the street, looking around as you did your best to dodge the crowd. The first thing you noticed was a large skyscraper marked with a '107',
to the left of it you saw what appeared to be a nightclub named 'Paradiso', and to the right of it was a humble info kiosk. Each likely held answers as to where and when you were,
but the skyscraper appeared to be closed for now. You travelled to the...`+"\n",
	)

	a := 0
	attempts := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			fmt.Println("Your options are 'go to nightclub' and 'go to kiosk'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			if attempts < 1 {
				fmt.Println("Someone rudely bumps into you, scoffing at you as they walk away. You notice a small card fall from their bag and take it.")
				GetMoney(500, p)
				attempts++
			} else {
				fmt.Println("The crowd is too busy -- you can't find anything of value.")
			}
		case "go to nightclub":
			selection = 4
			a++
		case "go to kiosk":
			selection = 5
			a++
		case "go to skyscraper":
			selection = 3
			a++
		default:
			fmt.Println("You can't think of a reason why you would want to do that right now.")
		}
	}
	return selection
}

//Scene3 is the skyscraper area.
func Scene3(p *Player) int {
	io.WriteString(os.Stdout,
		`You approach the skyscraper and find that the entrance is in fact locked as they fail to slide open. You find this curious, but that thought is quickly
pushed from your mind as you recoil in horror. To your dismay, the skyscraper has glass paneling, and your hair looks about ten times more awful than you could
have imagined. You sigh as you come to terms with your truly disgusting appearance before someone approaches you from the side. "Yo, you tryin' to get
in here?" You turned and saw a woman who's hair was somehow worse than your own, chewing gum and blowing a large bubble as it broke with a pop. She turned
to look into the building as she continued, "This building is owned by EdenTech, those bigwigs that control all the infrastructure that runs this city.
This is their HQ, and you ain't gettin' in unless you get someone inside to let you." She places her hand on the glass before clenching it into a fist, mimicking
a punch. "Can't break in neither, cause the drones'll getcha and chew ya up like a dog starved for weeks." You shot her a puzzled glance, and she seemed to
understand that you had no idea what she was talking about as she sighed. "Look, just stay away from here, 'less you wanna cause trouble. If you think you might,
go ahead and follow me. I know a place." She began to walk away and waved, and you considered following her. However, you could still visit the nightclub
or the info kiosk from here.`+"\n",
	)

	a := 0
	attempts := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			fmt.Println("Your options are 'follow', 'go to nightclub', and 'go to kiosk'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			if attempts < 1 {
				fmt.Println("You take a look around you and find a flyer for EdenTech on the ground, which you take.")
				InventoryAdd("EdenTech Flyer")
				attempts++
			} else {
				fmt.Println("You quickly scan the area for something interesting, but you didn't find anything.")
			}
		case "follow":
			selection = 6
			a++
		case "go to nightclub":
			selection = 4
			a++
		case "go to kiosk":
			selection = 5
			a++
		default:
			fmt.Println("You can't think of a reason why you would want to do that right now.")
		}
	}
	return selection
}

//Scene4 is the nightclub scene.
func Scene4(p *Player) int {
	io.WriteString(os.Stdout,
		`You decide to pay a visit to the nightclub, crossing the street and pushing your way through the crowd. You begin to hear the sound of a thumping 
bass as you approach, and oddly enough, you notice that there doesn't seem to be a queue -- or a door, for that matter. You head down the stairs and through the 
opening in the wall, and you're greeted by flashing lights and neon signs. The music was incredibly loud and drowned out almost everything else, a loud chatter 
filling the room as you once again pushed through another crowd. Cyborgs were dancing together on the floor, their metallic bodies covered in lights that flashed 
in time with the music as they raved, the sight almost blinding to you as you decided to keep to the edges of the club. As you're looking around, you accidentally bump 
into an incredibly large man, his attention immediately focused on you as he looked down at you, scowling. "Watch where you're goin', punk," he grunted at you 
as he slowly walked off, joined a group of thugs as they all shot you malicious glances. You figured you'd best stay away from them for now. You looked around the club 
for someone to talk to; maybe someone in here could give you some answers. You decided the bartender might know something, but you could also leave and go back to the info kiosk. 
You also supposed you could ignore caution and ask your new, threatening, friends. You decided to...`+"\n",
	)

	a := 0
	attempts := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			fmt.Println("Your options are 'talk to bartender', 'talk to thugs', and 'go to kiosk'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			if attempts < 1 {
				fmt.Println("You find a bottle of pain pills on the ground. Kind of shady and disgusting, but you decide to take one. You gain 25 hp.")
				TakeDamage(-25, p)
				attempts++
			} else {
				fmt.Println("This area is a bit too crowded for thorough searching.")
			}
		case "talk to bartender":
			selection = 7
			a++
		case "talk to thugs":
			selection = 8
			a++
		case "go to kiosk":
			selection = 5
			a++
		default:
			fmt.Println("You thought about doing that, but the music was so loud it caused you to forget what it was.")
		}
	}
	return selection
}

//Scene5 is the kiosk scene.
func Scene5(p *Player) int {
	io.WriteString(os.Stdout,
		`You figured the safest thing to do would be to pay the info kiosk a visit. You took in the city a bit more as you approached the stand, 
still wondering how it was that you ended up in such a place. You arrived at the kiosk and noticed it was covered in stickers, mostly of cats, but some
seemed to be pictures of cartoonized food with cute faces. You tried to look for a pamphlet of some kind, but the only thing you could see was a service 
bell. You decided the only thing to do was ring it, and you gave it a quick smack. Almost immediately someone popped up from under the counter, and you had
never regretted a decision faster in your life. You were greeted by what appeared to be a short, crazy-eyed girl with pigtails, her work uniform a bit tattered. Her hands were 
behind her back for a moment, and you flinched away as she quickly thrust a large knife into the countertop, giggling at you as her wild eyes took you in, 
almost like she was trying to determine how hard you'd be to kill, her smile making you shiver. "I'm so happy to have someone to help! What can I do for you today? And don't you DARE leave 
before I've answered your questions!" You held up your hands and tried to stay back from the counter. You'd better get your questions out before this girl loses it, 
but you could always ignore her and try to run.`+"\n",
	)

	a := 0
	attempts := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			fmt.Println("Your options are 'where am i?', 'who are you?', and 'run away'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			fmt.Println("You'd better not make any sudden movements. You might catch that knife.")
		case "where am i?":
			selection = 9
			a++
		case "who are you?":
			if attempts < 1 {
				io.WriteString(os.Stdout, `"What a stupid, annoying question!" She let out a yell and pulled the knife up from the countertop, 
throwing it at you. You tried to dodge but the knife was faster, and it sank into your leg with a sickening shlick. You screamed in pain and pulled it out, 
noticing that you somehow weren't bleeding, but you were still in incredible pain. You take 50 damage, but keep the knife for yourself.`+"\n")
				InventoryAdd("Knife")
				TakeDamage(50, p)
				if Hp <= 0 {
					GameOver(5)
				}
				attempts++
			} else {
				fmt.Println("You must be crazier than her if you're going to ask that again.")
			}
		case "run away":
			fmt.Println("Bad move. You try to run, but the girl pulls out two more knives from under the counter and throws them your way, stabbing your vitals with deadly precision.")
			fmt.Println("You take 99999 damage.")
			TakeDamage(99999, p)
			if Hp <= 0 {
				GameOver(5)
			}
		default:
			fmt.Println("You decided against that. You wouldn't wanna incite this girl.")
		}
	}
	return selection
}

//Scene6 is the sewer hideout scene.
func Scene6(p *Player) int {
	io.WriteString(os.Stdout,
		`You followed the mysterious girl across the streets as she quickly weaved through the crowd, leaving you struggling to keep up. 
She led you into a large, dark alley which connected into several smaller alleys, almost like a labyrinth of buildings. You picked 
up the pace as you were determined not to lose her, watching her as he lifted a sewer cover and dropped down. You weren't excited about going after her, 
but you pressed on anyway, dropping down the hole with a thud as the girl turned back to look at you, pressing a finger to her lips to shush you. 
You nodded and followed her to a small metal door which she knocked on in a peculiar way, likely as a password. The door slowly creaked open 
and you walked inside. You stayed close to her as you looked around, seeing people practicing hand to hand combat, firing guns at a small range, 
and moving supplies back and forth from an underground river deep within the sewer. You were led to a small office and you couldn't help but 
notice the almost completely cybernetic man sitting at the desk before you, his eyes glowing as red as the end of his cigar as he smoked it, grunting at the girl as 
she brought you in. "Who's this?" "This one's saying they're wanting to stir up some trouble," she replied. You couldn't remember saying any 
such thing as the terminator in front of you looked you over, nodding to you as he saw your arm. "If it ain't obvious to you, we're starting 
up a rebellion here. You might think us terrorists or somethin', but those EdenTech suits are the true enemies of the state. They've been 
screwing over people since before the turn of the 22nd century, and we're gettin' sick of it. Ava here's told me she ain't never seen you around before, 
and you're just unassuming enough that we may have a job for someone like you. You wanna help push this city in the right direction?" he asked you, 
leaning forwards in his chair, his red eyes staring into yours intently. You weren't sure whether to take his assessment of you as a compliment, 
and you were more than a bit confused, somehow finding yourself in the middle of a revolution. You looked over at Ava and she rolled her eyes, 
saying, "Just git on with it, already." This boiled down to a simple yes or no, but you were uneasy about what might happen if you declined. 
You decided your answer was going to be...`+"\n",
	)

	a := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			fmt.Println("Your options are 'yes' and 'no'. Not sure why you needed a hint for this one.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			fmt.Println("That's probably not a good idea.")
		case "yes":
			selection = 10
			a++
		case "no":
			selection = 11
			a++
		default:
			fmt.Println("You had best just give them a clear yes or no.")
		}
	}
	return selection
}

//Scene7 is the bartender conversation.
func Scene7(p *Player) int {
	io.WriteString(os.Stdout,
		`You head over to the bar, deciding that the bartender might be your best bet at understanding your situation. You took a seat and a hovering, happy 
looking robot greeted you. You noticed his nametag read 'Barry :)'. "Hi there! I'm the B4R bartending bot, but you can call me Barry! What can I make for 
you today? Perhaps some coke and vodka, or the house's special oil?" His voice was the most annoying thing you'd ever heard, and you knew you definitely didn't 
want anything to do with that oil. You calmly asked Barry where you were, and he replied, "You're at club Paradiso, friend! Our little heaven on earth!" You 
began to rub your temples and decided that you should have been more specific. You ask Barry what city you were in. "Oh! You're in the ward of New Shibuya, 
buddy! After it was wiped off the map in the late 21st century EdenTech built it back from the ground up to be their new headquarters! Apparently their CEO has 
a huge thing for eastern culture and --" You let the bot drone on as you can't bear to pay attention to him anymore. So you were in Shibuya, or the 'New' one, anyway. 
You were curious as to how you ended up in Japan or why you were there, but even more curious as to why there seemed to be no trace of the Japanese language in sight. 
You struggled to think as Barry kept talking, however, he slowly began to float upwards as he spoke, eventually colliding with the ceiling of his bar with a loud clang. 
"Ouch!" he said cheerfully as he floated back down to talk to you some more. "So what brings you here? I've been here for a while you know. I was made in the year --" 
You quickly push Barry out of your senses before he notices your frustration. "Aw, I'm sorry. It's just that no one's come to talk with me in a while. I'd really love 
it if we could keep chatting." You begin to feel slightly bad for the bot as you had noticed you had been the only one seated there, but you also weren't sure if 
you could take much more of this abuse. You could either stay a while and listen, or leave immediately and pretend you had never met Barry. You decided to...`+"\n",
	)

	a := 0
	attempts := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			fmt.Println("Your options are 'stay' or 'leave'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			if attempts < 1 {
				fmt.Println("You find that someone had left their belongings on the table and you weren't surprised in the slightest. You take their cards for yourself.")
				GetMoney(3000, p)
				attempts++
			} else {
				fmt.Println("There doesn't seem to be anything else here that you can take.")
			}
		case "stay":
			fmt.Println("You decide to stay with Barry. For all your compassion, you slowly succumb as he literally talks you to death. Game over.")
			os.Exit(7)
		case "leave":
			fmt.Println("You make the incredibly sound decision of bailing on Barry, and you decide to head for the skyscraper.")
			selection = 3
			a++
		default:
			fmt.Println("That doesn't seem like the right call right now.")
		}
	}
	return selection
}

//Scene8 is the conversation with the thugs.
func Scene8(p *Player) int {
	io.WriteString(os.Stdout,
		`You approached the thugs as they all turned to look at you, giving you the nastiest stink eyes you had ever seen. The one you had bumped into earlier 
stepped up, crossing his arms in front of you. "Ya got some nerve, punk. I oughta beat you for pushing into me like that," he growled, making you shudder a bit. 
You calmed yourself and explained that it was an accident. You were new here, and you simply wanted to know where you were. The lead thug started laughing 
and his cronies followed suit, their bellows almost as loud as the music as you tilted your head in confusion. "Look boys, we got a freshie! Listen punk, I'm gonna 
let you in on a little secret. Me and my boys are not people to be messed with, and I think I'm gonna teach you that real quick. He reaches down and quickly grabs 
you by your shirt, lifting you up into the air as you struggled, unable to break free from his grip. "You're gonna pay up 300 credits for your little misstep earler, 
and if you don't, I'm gonna beat your lights out. What's it gonna be, freshie?" You kept struggling against his hold as you assessed your options. You really didn't 
want to give this guy any money, but you definitely couldn't take him in a fight. Maybe you could go for a sucker punch -- hit him before he hits you and break free, 
but if that went sideways you knew you'd be in trouble. What do you do?`+"\n",
	)

	a := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			fmt.Println("Your options are 'pay' or 'fight'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			fmt.Println("You're being held in the air by a man who's built like a tank. Searching isn't a viable option for you.")
		case "pay":
			if Money < 300 {
				fmt.Println("You told the thug you didn't have that kind of money. He shrugged with a grunt, then promptly hit you in the face.")
				fmt.Println("You took 25 damage and lost conciousness.")
				TakeDamage(25, p)
				if Hp <= 0 {
					GameOver(8)
				}
				fmt.Println("After a little while, you wake up on the cold concrete outside the club. With an aching head, you make for the skyscraper.")
				selection = 3
				a++
			} else {
				fmt.Println("You decided to give the thug what he wanted.")
				GetMoney(-300, p)
				fmt.Println(`"Good. If more punks paid up we'd have to beat less of 'em. Get out of my sight."`)
				fmt.Println("You were released from his grip and pushed away, landing on your back. You brushed yourself off and headed for the bar.")
				selection = 7
				a++
			}
		case "fight":
			fmt.Println("You decided to fight back, and you gave the thug a swift kick between the legs.")
			fmt.Println("You were horrified to hear a clang of metal as your foot made contact, the thug laughing heartily at you.")
			fmt.Println(`"You think I'd be that stupid, kid? Now you're gonna get what you deserve."`)
			fmt.Println("The thug pummeled you hard, dealing 45 damage and knocking you out cold. You shouldn't pick fights you can't win.")
			TakeDamage(45, p)
			if Hp <= 0 {
				GameOver(8)
			}
			fmt.Println("After a little while, you wake up on the cold concrete outside the club. With an aching head, you make for the skyscraper.")
			selection = 3
			a++

		default:
			fmt.Println("You're not in a position to think of something more creative.")
		}
	}
	return selection
}

//Scene9 is the second kiosk scene.
func Scene9(p *Player) int {
	io.WriteString(os.Stdout,
		`As you asked where you were, the crazy girl twitched and you saw a spark fly from her head, a bit of smoke coming out of her ear. Her demeanor 
completely changed, going from a raving lunatic to a depressed lump. She let out a heavy sigh and leaned forward on the counter, pulling the knife out and 
toying with it a little bit, looking up at you with tired eyes. "If you really couldn't tell by the skycraper in front of you, this is New Shibuya. It's the 
base of operations for the EdenTech corporation. They're the guys that made me and all the other bots, and they probably made that metal arm of yours too." She 
squinted at you, saying, "You must be pretty dumb if you don't know that. Either that or you're some kind of alien. Everyone knows who EdenTech is. I'm not 
gonna ask, though. It would just bore me." You were thankful she didn't pry -- not that you would have known what to say anyway. Though she had nearly killed 
you, you began to feel a bit of sympathy for this android. She had probably just been driven insane by her thankless job as an info desk clerk. She let out 
another sigh, rolling her eyes at you as she lazily waved the knife at you. "Can ya get going? I'd really like to have some peace and quiet, and your hair bothers me." 
Your self esteem destroyed, you slinked away from the kiosk, and you decided to go check out the skyscraper.`+"\n",
	)
	selection := 3
	return selection
}

//Scene10 is the bomb steup scene.
func Scene10(p *Player) int {
	io.WriteString(os.Stdout,
		`The metal man let out a synthetic laugh as you nodded in agreement, pulling the cigar from his mouth and putting it out on his desk. "Good, glad to hear it. 
Yer job's gonna be very simple, so even if you're a complete moron you would have a hard time screwing it up." You tilted your head and wondered exactly how stupid you 
looked as he leaned down to rummage under his desk, pulling up a satchel and tossing it to you. You noticed it was a bit heavy, and you opened it up to see a large brick 
of C4 explosives. You looked at the man in surprise as he gave you a smirk, point at the brick with a wagging finger. "So you're gonna go to that skyscraper you paid a visit 
to earlier, and you're gonna plant it right at one of the support beams at the base. One you've done that, you're gonna run back here and I'm gonna flip the switch. EdenTech's 
tower of greed's gonna come down hard, and we're all gonna have a party when it does. Ava's gonna trail you to make sure you do your job." "Ugh, really? I don--" 
The man slammed his fists on his desk, cracking it down the middle as it made the both of you jump. "You are gonna go with him to make sure he gets it done. If it goes south, 
you need to make sure you get that C4 and bring it back. You got that?" He glared at her almost evilly, and she nodded, looking away. "Got it. Come on." She waved to you and 
you followed her out of the sewer, keeping the satchel close as you hurried back to the skyscraper through the dark alleyways. Ava was keeping close behind as you made your way 
back out of the maze, and you found yourself overwhelmed by your situation. You weren't sure what to do. You knew if you went through with this a lot of innocent people 
might get hurt; you didn't forget how many were roaming in the square in front of the tower. However, if this corporation actually was bad news, you might be doing the 
world a favor by taking out their main headquarters. You glanced back at Ava as you approached the building. She motioned for you to go plant the bomb, and you hesitated. 
You wondered if you could somehow get away, but you figured Ava would probably try to kill you. You tried to subtley wave her over and she rushed over, her face grimacing. 
"What are you doin'? Just plant the bomb!" If you were gonna get away, now was your last chance. Either that, or plant the bomb. You decided to...`+"\n",
	)

	a := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			fmt.Println("Your options are 'plant bomb' or 'escape'.")
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			fmt.Println("You try desperately to search for some other way out, but you find none.")
		case "plant bomb":
			fmt.Println("You embrace the Chaos of revolution.")
			selection = 12
			a++
		case "escape":
			if InventoryCheck("Knife") == true {
				fmt.Println("You embrace the comfort of Order.")
				selection = 13
				a++
			} else {
				fmt.Println("Seeing no other choice, you attempt to run away. Ava rolls her eyes, an indifferent expression on her face as she shoots you.")
				GameOver(10)
			}
		default:
			fmt.Println("In your panic, you forget what it was you were just going to try to do.")
		}
	}
	return selection
}

//Scene11 is the sewer escape scene.
func Scene11(p *Player) int {
	io.WriteString(os.Stdout,
		`You shake your head no, uninterested in anything they might have asked of you. The cybernetic man scoffs and takes a long drag from his 
cigar, slamming his fist on the desk. "Shame, kid. Now we've gotta kill ya." He pointed at Ava, and then at you. "Take care of them, NOW." She nods 
and you quickly turn tail and run as she draws her pistol. She fires at you as you run out the door but misses, and you quickly sprint back out where 
you came from as the others in the hideout looked on in confusion. You hear Ava scream orders to hunt you down as you began to hear a thunder of footsteps 
chasing after you down the waterways. You were letting your adrenaline carry you as you barreled down the sewers with reckless abandon, desperate to get 
away from your pursuers as quickly as possible. You scanned the sewers for any semblance of a hiding place, but it was proving difficult. You knew you 
had to keep moving no matter the cost, but you were beginning to get tired. A hiding place might save you -- you just had to find one first. You...`+"\n",
	)

	a := 0
	attempts := 0
	selection := 0

	for {
		if a != 0 {
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		switch line {
		case "hint":
			if attempts < 1 {
				fmt.Println("Your only option for now is 'keep running'.")
			} else {
				fmt.Println("Your options are 'keep running' and 'hide'.")
			}
		case "check":
			PlayerCheck(p)
		case "backpack":
			Backpack()
		case "search":
			if attempts < 1 {
				fmt.Println("You're able to locate a crack in the wall ahead that you think you could squeeze into.")
				fmt.Println("'hide' option added.")
				attempts++
			} else {
				fmt.Println("You've no time to be picky. If you want to hide, the wall is your only shot.")
			}
		case "keep running":
			fmt.Println("You keep running as fast as you can, but eventually the rebels catch up to you. With no way out, you beg for mercy, of which there was none.")
			GameOver(11)
		case "hide":
			if attempts < 1 {
				fmt.Println("You haven't found a hiding spot yet.")
			} else {
				fmt.Println("You embrace the neutrality of Pacifism.")
				selection = 14
				a++
			}
		default:
			fmt.Println("You've no time to think about something like that right now.")
		}
	}
	return selection
}

//Scene12 is the chaos ending.
func Scene12(p *Player) {
	io.WriteString(os.Stdout,
		`You decide to plant the bomb. After fiddling around with it a bit, you stick the duct-taped C4 to the designated support beam. You and Ava make 
a break for it before you're spotted, and you make your way back to the hideout. Once the boss ensures everyone in the resistance is safely inside, he triggers 
the detonator. The walls around you shake as you can feel the '107' tower crumbling down, the room erupting in a roar as the rebels celebrate a small victory 
over their corporate overlords. You're unsure of the impact of what you just did, but you assure yourself that it was the right choice. Ava smiles as you as the 
world begins to fade from your view, the sounds around you slowly muting as you awaken with a breath. The dream is done, and the game is over.`+"\n",
	)
}

//Scene13 is the order ending.
func Scene13(p *Player) {
	io.WriteString(os.Stdout,
		`You recall the knife you so graciously obtained from the info clerk earlier. With a shaking hand you reach into your pocket and pull it out, 
quickly thrusting it into Ava's chest before she can react. She doesn't fight her fate as she quietly slumps over, falling onto her back. You cover your mouth 
in horror at what you had just done, but compose yourself before anyone notices your crime. You throw the satchel on top of her for the police to find, and you 
duck into an alley just as her body is noticed. You begin to hear a siren as you delve deeper into the maze of buildings, finally settling on somewhere to rest 
as you find a secluded spot, sitting on the ground and hugging your knees to your chest. You try to assure yourself that you had done the right thing, knowing 
that you had saved all of those people in the square by taking one life and prolonging a bloody revolution a little while longer. Suddenly, your vision begins 
to fade, the sounds around you being drowned out by silence as you awaken with a gasp. The dream is done, and the game is over.`+"\n",
	)
}

//Scene14 is the pacifist ending.
func Scene14(p *Player) {
	io.WriteString(os.Stdout,
		`You managed to evade your pursuers enough to quickly hide in the wall's opening. You keep silent and stay your breath as the group of rebels pass you by, 
letting out a low sigh of relief when you no longer hear their footsteps. You decide to remain there for a while, keeping yourself scarce while you were hunted. 
You eventually emerge from your hiding place, finding your way back to the surface through a different exit. You manage to make your way outside the city, knowing
setting foot there again may be your last mistake. As you try to decide on your next move, your vision begins to fade to black, your senses slowly dissipating to 
nothingness. You slowly awaken, your eyes fluttering as they take in the morning sun. The dream is done, and the game is over.`+"\n",
	)
}

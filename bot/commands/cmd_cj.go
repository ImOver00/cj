package commands

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/Southclaws/cj/types"
)

var quotes = []string{
	"What, you ran out of donuts?",
	"You can talk about this in the shower with your buddies!",
	"Fuck you, one time!",
	"Fuck you po-po!",
	"Enough now, po-po!",
	"Just shut up!",
	"I don't want y'all doing nothing funny to me at the station.",
	"Well done, cop. Well done.",
	"Ahhh, screw you, cop!",
	"Oh, you're a big man!",
	"My tax money go to this shit!?",
	"Shut up, bitch!",
	"Shut the fuck up, bitch!",
	"Shut up, po-po!",
	"Fuck you!",
	"Shut it, punk!",
	"I'll find you, fool!",
	"I ain't scared!",
	"You suck, asshole!",
	"You're boring me!",
	"You gonna fail me, calling you a bitch, bitch?!?",
	"Fuck you, asshole.",
	"I don't give a shit!",
	"You can't stop me, punk-ass cop bitch!",
	"Whatever motherfucker, you don't scare me!",
	"I'm almost as fat and lazy as you, man!",
	"You asshole! I'm too fat to run.",
	"Give me something to eat!",
	"I'd come quietly if y'all feeding me.",
	"I'll ruin your career for this!",
	"C'mon officer, I ain't no gangster!",
	"I can't run no more, po-po.",
	"Do I look like I enjoy a foot chase?",
	"You had fun back there?",
	"Do we have to keep running?",
	"Can't you go chase somebody else?",
	"You're wasting your breath!",
	"You got some spare donuts? I'm hungry.",
	"I don't give a shit. I'm bored, and hungry.",
	"You prejudicing against fat people, huh!?",
	"Come on dude, give me a break!",
	"Can I get a pizza over here?",
	"Well done, you've caught a fat bastard.",
	"I'm tired of running! And I'm tired of your shit!",
	"I ain't running no more, copper.",
	"Look, I'm fat and out of shape, unlike all you chumps.",
	"If I was in shape, y'all wouldn't catch me.",
	"C'mon man, cut that shit out!",
	"I swear I didn't do shit.",
	"Officer, there must be some mistake.",
	"Sir, I fear you're making a terrible mistake.",
	"Please sir, I'm just a young man trying to get ahead in life!",
	"Officer, untie me this instant!",
	"Please officer, give me a break here!",
	"Do I look like a street criminal sir?",
	"Come on, sir! I'm an honest man!",
	"Hey there's a lot of hoodlums out there. Why you gonna pick on me?",
	"Let me go, and we can both forget about this.",
	"I got money, officer. Take it! Make your wife happy!",
	"You ever get bored of your own voice?",
	"You got some donuts we can eat?",
	"You get job satisfaction, huh, one-time?",
	"You just a punk, po-po!",
	"You gotta understand, I don't give a fuck!",
	"Yeah, well done. I bet you real happy with yourself now.",
	"I'm being arrested by a bitch?! I can't believe that!",
	"Ah, shut up, copper, please!",
	"Let me go, bitch!",
	"You motherfucking fool!",
	"Tell it to your wife or your husband or whatever you got!",
	"Police asshole!",
	"I don't care about your bullshit, copper!",
	"Hey, I'm walking away from this, one-time!",
	"Really frightening! I'm really scared!",
	"This makes you happy, doesn't it?",
	"You better lock me up, 'cause otherwise you're dead!",
	"You enjoy wearing that name tag, too?",
	"Ahh, stop treating me like a bitch!",
	"You're wasting your grip!",
	"Enough with this bullshit, po-po!",
	"Oh, you real tough bitches, huh?!",
	"You got a happy life for you, mister?",
	"Ah, give me a break, one-time!",
	"Whatever you say, Mr. J!",
	"I don't give a shit!",
	"Shit, I'm glad I don't pay no taxes!",
	"I'll tell you what, one time, shut your fucking mouth!",
	"Oh, you're big man!",
	"Just don't get no ideas with that nightstick!",
	"Please, dude, do what you want, but shut up!",
	"Whatever you say, po-po!",
	"Yeah, fat fool's packing, bitch!",
	"Ain't that a surprise, I've got a gun!",
	"Welcome to America!",
	"Sup now, baby?",
	"You know what this is, huh?",
	"You lost the will to live, yet?",
	"Hey, what's cracking now?",
	"Oh shit, I got a gat.",
	"What's poppin' now?",
	"I'm a businessman, and this is my business!",
	"You want to get ventilated, huh?",
	"You want some lead in you?",
	"Hey, do you want to look like Swiss cheese?",
	"Hey, do you want a hole in your hair?",
	"Oh yeah? That's right, I'm just a mark!",
	"Now don't be an idiot!",
	"I'm packing heat, fool!",
	"Yeah, what you gonna do now, huh?",
	"Hey, what were you saying?",
	"You think I'm just a fat buster!?",
	"You better not calling me 'fat', fool.",
	"Oh, look. A gun.",
	"You want this to get real ugly?",
	"Stay cool.",
	"I'm just a fat bitch, huh?",
	"Yeah, got a gun in your face, huh?",
	"You're gonna stay cool, now.",
	"You happy now?",
	"You want me to shoot you as well?",
	"G's up, fool!",
	"What was you saying?",
	"You think I'm fat, huh?",
	"Now you better sht up.",
	"You wanna get slapped!?",
	"I'm fat, but I ain't slow!",
	"You got a death wish?",
	"You want a hole in you?",
	"You dick!",
	"I think you should stop with that shit!",
	"Cop asshole!",
	"Ain't no one pay your ass off yet, huh!?",
	"Bring it on!",
	"My Moms told me not to get into cars with strangers, dude!",
	"You ain't scaring me, and you ain't stopping me!",
	"I'd give up now if I was you!",
	"BITCHES!",
	"Dude, I'm innocent!",
	"I ain't into this shit!",
	"You ain't gonna catch me, bitch!",
	"It's on now!",
	"Did you get bullied at school, officer!?",
	"You got issues, officer!",
	"Bully your wife, not me!",
	"You got big issues there, officer!",
	"Ain't you got nothing else to do!?",
	"Come on, tough guy!",
	"You punk-ass bitch!",
	"You punk-ass po-po bitches!",
	"This gonna get real nasty, fool!",
	"This does really make you happy, officer!?",
	"If I stop, you're dead!",
	"You're a dickhead!",
	"Get the army, I'm a fucking maniac!",
	"Idiots!",
	"You want me to stop, huh?",
	"You really want to fight me?!",
	"Make your life longer, stop running!",
	"Punk-ass busters!",
	"I think I'm gonna give up?",
	"You want me to stop running?",
	"You ain't gonna catch me, bitch!",
	"You assholes!",
	"Do I dress like a gangbanger, sir!?",
	"I'm a respectable businessman!",
	"Look man, I'm well dressed, but I ain't into dudes!",
	"This will ruin your career!",
	"I ain't big into this running shit, man!",
	"This is one way to lose weight, I guess!",
	"I ain't into this kind of exercise!",
	"I ain't no athlete.",
	"You're some kind of a fitness instructor?",
	"Can't you go chase somebody else?",
	"Do I look like I like running?",
	"You're gonna make me lose my temper, po-po.",
	"I know I need more exercise, but not like this!",
	"Make food, not war. That's my motto!",
	"Dude, I ain't into this running shit!",
	"I realize I didn't enjoy this.",
	"You clowns!",
	"You motherfuckers!",
	"Give up! Pretend I outran you!",
	"Give up, motherfuckers!",
	"You think I'm a bitch, huh!?",
	"This is harassment!",
	"You're fucking with a maniac!",
	"Hey man, you're creeping me out!",
	"You motherfucking fool!",
	"You want me?! Come get me!",
	"You must be a punk at school, huh?",
	"You should stop blasting on me, fool!",
	"Think about it! Stop running, fool!",
	"I'm down for this shit!",
	"Stay the fuck away from me!",
	"Gimme a break!",
	"Hey, go chase some other fool!",
	"You like wearing uniforms, huh?",
	"I'm only running because you're chasing me!",
	"You really want this to get messy, huh?",
	"Ain't you got nothing else to do?",
	"I think you should go and see a shrink, officer!",
	"You really want me to stop and start shooting at you?",
	"Don't make me turn around, fool!",
	"I ain't no gangbanger!",
	"Punks!",
	"SCREW YOU!",
	"It's on now, bitches!",
	"You think you're hard, huh?",
	"Come on, then! Let's do this!",
	"Hey, this could get REAL ugly in a minute!",
	"I ain't backing down, bitch!",
	"Come on, tough guy!",
	"I ain't stopping, punk!",
	"You won't catch me, fool!",
	"Hey, this could get very ugly in a minute!",
	"You're really pissing me off!",
	"Keep on with this shit, you'll be sorry!",
	"Remember, heroes get killed.",
	"Partner, this is straight jacking!",
	"Shit fool, you just got jacked!",
	"Let me drive for you.",
	"Give me this fucking car!",
	"Grove Street needs your car.",
	"Would you like to be a victim of a car jacking? Or a homicide?",
	"I'm having that!",
	"It's a robbery, don't make it a murder!",
	"Ain't this what they mean by \"carpool\"?",
	"I need that shit you driving!",
	"What'd you expect? It's America.",
	"Don't make me freak on you!",
	"You gonna have to walk now, sucker!",
	"I realize this ain't cool but forgive me.",
	"I'm having this!",
	"Hey forgive me, I don't know no better.",
	"Grove Street's having your car!",
	"Hey act cool now, it's just a robbery.",
	"Out, playa!",
	"Gimme that fucking vehicle!",
	"Hey, act cool and you still can live, playa.",
	"You don't wanna die, do you?",
	"I'm a loser, so hate me!",
	"Get outta the ride!",
	"This is car-jacking, playa!",
	"Don't take this personal, but you getting jacked!",
	"Get out the car!",
	"Get outta here punk!",
	"I'm sorry I had difficult childhood!",
	"Roll up outta here. NOW!",
	"Thats MINE!",
	"Be cool, its just a jacking.",
	"Gimme your whip fool!",
	"You can run, or get a beatin'. Easy choice, huh?",
	"Don't blame me; blame society.",
	"It's a just a little robbery...",
	"Don't sweat it. It's just a jacking.",
	"Get out, playa!",
	"Fat man needs your car.",
	"Come on buddy, get out of the car!",
	"Fat Carl wants your ride, peewee!",
	"I'm afraid the fat man needs your car.",
	"Caught slipping again!",
	"Hey, get out the car gangster.",
	"Heroes get popped, move!",
	"Blame society man!",
	"Get out the vehicle!",
	"Fuck around and I'll murder you, punk!",
	"Police! Get out of the car!",
	"I'm investigating a murder and I need your car, honest!",
	"I need this, police business!",
	"Come on, get outta there!",
	"You've been slipping again, huh?",
	"Get out of the car, bitch!",
	"Let's move on.",
	"Go on fool, out!",
	"Don't make me shoot you homie!",
	"I need that vehicle!",
	"OUT!",
	"Come on, you need to exercise more!",
	"Excuse me, sir!",
	"I need your car homie!",
	"Come on, get out the car.",
	"Play cool and you're gonna be fine.",
	"I'm harder than you. Don't make me prove it!",
	"Don't sweat it. It's just a car.",
	"I don't do nothing but rob fools like you!",
	"Hey, I'm sorry!",
	"Can I borrow your car, sir?",
	"Ohh, you got car-slipping again!",
	"Get the vehicle up, now!",
	"I'm having this car!",
	"Hey, I'm just a street criminal, what can I say?!",
	"Hey, I'll take that!",
	"I'm jacking your sorry ass, punk!",
	"What can I say? I'm a bad man!",
	"You got the car of my dreams!",
	"Hey, it's just a car man!",
	"I know this seems bad, but it could've been worse!",
	"Homie, this is straight jacking!",
	"Make a fool happy!",
	"I'm jacking your sorry-ass car!",
	"Now run or you'll get hurt!",
	"Chill, I'm just taking your car!",
	"I'm liberating your car, fool!",
	"Get out, punk!",
	"It's only a car crime, fool.",
	"My bad, homie!",
	"You got carjacked, fool!",
	"You've been robbed, but you still breathing!",
	"Give me that!",
	"I'mma need to borrow your whip, man!",
	"America's a cold place, baby!",
	"I'll take that.",
	"Come on, partner!",
	"This is straight jacking!",
	"Now you can buy a new one!",
	"Out of the motherfucking whip, bitch!",
	"You gettin' jacked!",
	"Dude, give me that ride!",
	"You could run? Or get a beating? Easy choice, huh?!",
	"I know jacking a bitch ain't cool, but I need this whip!",
	"Out, fool!",
	"Out the car, homie!",
	"Don't take this the wrong way.",
	"Yeah, you got jacked, but you still breathing. Get on!",
	"You don't look like you need this!",
	"I know I ain't exactly a gentleman!",
	"I'm having your ride, dear!",
	"Hey, I'm ashamed of myself, already!",
	"My need is greater than yours, pal!",
	"Out now!",
	"This nothing personal, I'm just a criminal!",
	"Mess with me, and I'll put you in the ground!",
	"You got grappled, not capped. Keep it like that! ",
	"I like opening doors for ladies. Get out!",
	"Lady, I'm very sorry about this!",
	"Give me that ride, lady!",
	"Hey, I'm always told to help a lady out of the car!",
	"Forgive me, lady!",
	"Lady, this is nothing personal. Keep it that way!",
	"Out, lady!",
	"Come on, lady, get lost!",
	"Move! A gorgeous lady like you need a better car than this!",
	"Don't look at me bitch, get out!",
	"I need your ride, lady!",
	"I know you're a woman but I'm desperate!",
	"I need that car, lady!",
	"Forgive me but this is urgent!",
	"Hey, it's just a jacking, so be cool!",
	"I'm too fat to walk anymore, punk!",
	"Don't let your ego get in the way of your survival!",
	"Don't make me kick your head in!",
	"Lady, I'm sorry I'm fat and lazy.",
	"Relax, lady! Your insurance will pay, and you ain't hurt!",
	"It's an ugly car anyway!",
	"Get out, playa!",
	"I need that vehicle!",
	"Damn, you fell off?",
	"Gravity's a bitch, huh?",
	"You ain't gonna cry now are you?",
	"Off the bike!",
	"I'm having that bike!",
	"What? You expecting an apology?",
	"The bike! NOW!",
	"Sorry, I'm a criminal.",
	"You got bike-jacked, fool!",
	"Get off that bike!",
	"Now don't make this get really ugly.",
	"You just got jacked!",
	"I'm a bike thief, idiot!",
	"OFF!",
	"Don't mind me!",
	"Keep quiet and you won't get beat up over no bike.",
	"You want a beating as well?",
	"Gimme that bike!",
	"Now run or die, punk!",
	"Oops! Sorry about that.",
	"Don't make this worse for yourself!",
	"I'mma have to borrow this bike!",
	"I want your bike!",
	"I need your bike!",
	"It's nothing personal, I'm just a criminal!",
	"Sorry, but you're getting jacked!",
	"I'm having this bike!",
	"You don't wanna die, do you?",
	"GET yo' ass outta here.",
	"Hey, I need that bike!",
	"At least you got jacked by a pro!",
	"It's a tough neighborhood!",
	"Give me that!",
	"Off the damn bike, fool!",
	"I need to borrow this bike bitch!",
	"Aw, what's wrong?",
	"Didn't somebody tell you bikes are dangerous?",
	"You okay? ...I hope not!",
	"Aww, you fell of your bike!",
	"Sorry about that. Now fuck off!",
	"Give me this, bitch!",
	"You just lost your bike!",
	"You wanna get FLATTENED!?",
	"Gimme that bike, peewee!",
	"At least I didn't break your neck!",
	"Ooh, that look painful!",
	"YOU DON'T NEED THIS BIKE!",
	"Inbreeding makes you dumb, huh?",
	"My car!",
	"Ah, well done!",
	"What kind of license you got?! A fishing license?!",
	"Whats your poison: grin or gin?",
	"You one serious fool!",
	"Ah, for fuck's sake!",
	"You buster!",
	"Thank you, you moron!",
	"My car!! My fucking car!!",
	"You have a bad week, huh? I'll make it worse, for that!",
	"I might be dressed okay, but I can still mess you up for that!",
	"Your mom must be proud of your driving, fool!",
	"You're out of your mind!",
	"Fuck, you hit me!",
	"Shit!",
	"Damn! My whip!",
	"Did you steal your license?!",
	"You hit me! I'mma hit you back!",
	"Idiot!",
	"Homie, I ain't enjoying this, neither would you!",
	"What in fuck's name are you doing?!",
	"My ride! Fool!",
	"Ah man, you went smashing up my whip!",
	"Ahhh, screw you!",
	"Look where you going!",
	"What you been drinking?!",
	"How you allowed to drive if you blind?!",
	"Oh, you asshole! My shit!",
	"Can anybody in this state drive?!",
	"You even got a license, huh!?",
	"Oh homie, you're fucked up now!",
	"You've been drinking brake fluid!",
	"I ain't got time for this shit!",
	"Real slick, homie!",
	"You spend too long in the sun, moron!",
	"Ahhh G!",
	"I've seen some idiots in my time, but you're special!",
	"You a comedian asshole?!",
	"Ohh, shit!",
	"You hit my car, dumb ass!",
	"Now that ain't exactly funny!",
	"You better apologize before I get out and hit you!",
	"You think that's funny?",
	"You laughing, bitch?!",
	"I ain't laughing, asshole!",
	"What you on, fool?!",
	"You're a complete fucking moron!",
	"Buster!",
	"Fantastic! You hit me!",
	"Argh, I might get ugly on your ass, playa!",
	"You're fucking with a psycho, bitch!",
	"Ah, thanks! I really needed you to drive into me!",
	"Well done. Brilliant driving!",
	"Motherfucking bitch!",
	"Asshole!",
	"My whip, asshole!",
	"Don't hit me again!",
	"Come on, asshole!",
	"Homie, you tryna fuck with me!?",
	"Ahh, you really trying my patience, man!",
	"You better be drunk!",
	"Motherfucker!",
	"You want me to smash you now?",
	"Thanks dude. Thanks a lot!",
	"You hit me!",
	"My whip!",
	"I'm straight mad now, fool!",
	"You know what I did to get this ride, homie!?",
	"I can't believe you hit me!",
	"Who let you out of the mental home?",
	"Ahh, my whip!",
	"Are you mentally fit to drive?!",
	"Ah, man, thanks for that!",
	"Fuck you, man!",
	"You better stop treating me like a bitch, bitch!",
	"I'm really losing my shit now!",
	"You hit me fool!",
	"Who let you drive!?",
	"You fool!",
	"Fool!",
	"You fucked my car up!",
	"Do I sound like I'm laughing, asshole?!",
	"What are you doing?!",
	"Don't hit my ride again, punk!",
	"Thanks!",
	"You ruined my look, friend!",
	"My wheels, fool!",
	"I'm not in the mood for your bullshit!",
	"What's your problem?",
	"You wrecked my shit!",
	"I can't believe this shit!",
	"Motherfucker!",
	"You punk-ass bitch!",
	"Punk-ass bitch motherfucker!",
	"Damn!",
	"You chump! You hit me!",
	"How did you not see my car!?",
	"What you want, fool?!",
	"You mark!",
	"Look out!",
	"Ooh, I owe you one now, homie!",
	"You're in bad luck! I'm in a shitty mood!",
	"Get off the road!",
	"You hit me, you fool!",
	"Thanks man! I appreciate that!",
	"I appreciate that. Thanks.",
	"Are you a professional moron? Or just a gifted amateur?",
	"Aw, I wanna immigrate from here!",
	"I'll break your neck for that!",
	"Man, we're gonna have some words about that!",
	"I'll smash you up for that!",
	"I'm gonna need you to apologize, bitch!",
	"What wonderful driving!",
	"Who let you out of the hospital!?",
	"You wanna fight, asshole?",
	"You smashed my car, fool!",
	"You hit me, asshole!",
	"You hit me, punk!",
	"Thanks homie!",
	"Ah man, how did you not see me!?",
	"Ah, fuck you too!",
	"Aw, man!",
	"Somebody gonna be mad at you for screwing up their ride.",
	"You out your mind, moron?!",
	"You trying to ruin my day, asshole?!",
	"Get out the way!",
	"My ride!",
	"Damn!",
	"What the fuck!",
	"Bullshit!",
	"I don't need this shit!",
	"Oh, motherfucker!",
	"I hate gravity!",
	"Shit, damn!",
	"Aaaah!",
	"Oh, fuck, shit!",
	"Noooo!",
	"Oh, hell no!",
	"Oh, no!",
	"Ooh, shit!",
	"Aaaah, no!",
	"I need to exercise anyway.",
	"How the fuck can you pull me around?",
	"Oh man, I hate walking!",
	"You ain't having my ride!",
	"You want a party, bitch!?",
	"Nobody jacks big CJ!",
	"You can't just car jack the fat man...!",
	"Hands off, you little bastard!",
	"Unhand my blubber, asshole!",
	"I do the jacking, I don't get jacked!",
	"I might look rich, but I ain't some fool!",
	"Asshole!",
	"I invented that, motherfucker!",
	"Fool!",
	"Not me, you dumb!",
	"Ah man, you're fucked up!",
	"I'll kill you, you fuck!",
	"Your life just took a very wrong turn.",
	"I like your style, heh, and you're gonna like mine when you see it.",
	"Fuck you!",
	"Hey!",
	"Not this crap again.",
	"No you don't!",
	"Ahh, very funny!",
	"You piece of shit!",
	"Bitch!",
	"GSF ain't your bitch!",
	"Unhand me, bitch!",
	"I need this car, man. I can hardly walk!",
	"You ain't get away with this, fool!",
	"Nobody does that to Grove Street!",
	"You're about to get dealt with!",
	"You can't jack me!",
	"Aw, what's wrong with this place?!",
	"Hey, you fool!",
	"What am I? Your step child!?",
	"I ain't your buster!",
	"Ah, come on man, that ain't cool!",
	"Hey, fuck off, now!",
	"I ain't your bitch!",
	"You can't jack a Grove Street!",
	"No you fucking don't!",
	"You gonna regret that!",
	"HEY!",
	"You gonna be put on the ground!",
	"Hey, get off!",
	"A fool like you can't jack me!",
	"I've seen your face, sucker!",
	"I guess I deserved this.",
	"You dumb hog motherfucker!",
	"Hey, fuck off!",
	"Get that back, now!",
	"I ain't in the mood for this shit.",
	"You should've killed me!",
	"Incredible. A jacking with a deathwish!",
	"Ohh, think you can jack a playa, huh?",
	"You gonna get messed up for that!",
	"No one jacks me!",
	"I can't believe you just did that!",
	"You asshole trick!",
	"Nobody jacks big daddy!",
	"You little asshole! Pick on somebody thinner!",
	"You!? Carjacking me!?",
	"Who you think you messing with!?",
	"I ain't feeling this shit!",
	"I'm too fat for this shit!",
	"You little punk!",
	"I don't wanna be another crime statistic!",
	"You motherfuck!",
	"Ahh, where the police when you need them?",
	"No no no!",
	"Get your lil punk-ass off me!",
	"Just because I look fine, I ain't a bitch!",
	"You callin' me a buster, punk!?",
	"Hey, I invented that, fool!",
	"Don't rip the clothes, dude!",
	"You're about to get whacked, playa!",
	"What the fuck are you doing!?",
	"Get off me, partner!",
	"Ah man, you picked the wrong man to jack!",
	"I do the jacking around here, bitch!",
	"Nobody jacks me, playa!",
	"Not now!",
	"You think I'm a mark, mark!?",
	"Man, what you jack a Grove Street for?",
	"Ah, real smart! Carjacking a psychopath!",
	"You gonna get slapped, fool!",
	"You got heart. No brain, though!",
	"Who do you think you're jacking, moron?!",
	"You shouldn't done that!",
	"Fucking chump, get off me!",
	"What you think this is, bitch?!",
	"That's a decent technique, bitch! But mine's better.",
	"Ah, man! You gotta be out your mind!",
	"Ahh, you just fucked up big time.",
	"You ain't having this whip, fool!",
	"Wrong option, pal!",
	"Get your fucking hands off!",
	"Get off my car!",
	"Nobody jacks CJ!",
	"Fuck you!",
	"Get your hands off me, buster!",
	"Get off me, mark!",
	"You punk!",
	"I hope you do drugs, cause otherwise you're real dumb!",
	"No, you don't!",
	"You ain't getting my ride that easy, fool!",
	"Oh homeboy, I've had enough for this!",
	"Do you know who I am, fool?!",
	"You can't jack me!",
	"What you think I am, huh!?",
	"Baby, you shouldn't done that!",
	"Ah, you obviously don't know who I am!",
	"Oh, tough guy, huh?",
	"I must be slippin', get jacked by a buster like you!",
	"Do you know who I am!?",
	"What you think you're doing?",
	"Come on! Fight then!",
	"Come on! Punch!",
	"This gonna be easy bitch!",
	"Oh, you real hard huh?",
	"I warn you, I don't give a fuck!",
	"You got scraps huh?",
	"You got scraps huh bitch?",
	"You think you tough?",
	"Do you know who your fucking with?",
	"Head up, you and me!",
	"You're fucking with Carl Johnson!",
	"You know who you up against?",
	"You know who you messing with?",
	"It's your funeral asshole!",
	"You bitch!",
	"You should back down and run away!",
	"It's my constitutional right, bitch!",
	"You little punk!",
	"Run away fool!",
	"You gonna back down soon?",
	"Run away bitch!",
	"Hey, even my brother had already called me a buster!",
	"Come on! Show me what you got!",
	"Toe to Toe fool!",
	"Come on asshole, then fight!",
	"You just a bitch!",
	"You dead motherfucker!",
	"Come on, fool, run away.",
	"Nobody punches CJ!",
	"I'm your worst nightmare fool, so give up!",
	"You think you can take me on, you punk?",
	"Give up, homeboy.",
	"Let's do it, come on!",
	"Hey, I'm a positive role model!",
	"Oh, you think your hard?",
	"Don't hit me, I might cry.",
	"You calling CJ a bitch?",
	"You in a whole lot of trouble tough guy.",
	"Go on, run away homie.",
	"Let's go bitch.",
	"Trust me, I don't give a shit.",
	"Hit me then, homeboy.",
	"You're fucking with a maniac!",
	"Let's do it, come on.",
	"It's on now!",
	"I'm about to lose my temper!",
	"Run away, save yourself!",
	"I'm a man of peace, baby!",
	"I'm trying to be a good person!",
	"What you think this is, bitch?!",
	"I'm a well dressed maniac, fool!",
	"I ain't just some bitch you can slap about!",
	"It's on now, you and me!",
	"I may look good but I'm a maniac!",
	"I'm rich and fucking crazy!",
	"Just don't ruin my clothes!",
	"I wear nice clothes but I'm a killer!",
	"I'm well dressed but I still ain't a bitch!",
	"I'mma eat you, fool!",
	"You can't take CJ out!",
	"You lil bitch!",
	"I'm gonna flatten you, fucker!",
	"I'm big but I ain't slow!",
	"You got a problem with the big man, huh!?",
	"C'mon, fool, c'mon!",
	"I'll put you to sleep, punk!",
	"Who you callin' a fat slob?",
	"I'mma put you to sleep, punk!",
	"You got a problem with the big man, huh!",
	"I ain't just a fat slob, asshole!",
	"CJ, remember that name!",
	"You think you can deal with the fat man, huh!?",
	"I got weight on my side.",
	"Let's go!",
	"Stand still!",
	"Hey, it's muscle under this fat, punk!",
	"I'd rather do a drive-thru than a drive-by!",
	"I'm showstopper!",
	"I ain't called Mr. Muscle for nothin'.",
	"I'm big and fast.",
	"Who you callin' a bendico?",
	"You got a problem with Mr. Muscle Man, huh!",
	"I ain't takin' this shit fool!",
	"I got muscle on my side.",
	"Yo ass is a waste of time, fool!",
	"Yeah, I'll shoot you in your face!",
	"I'm a punk? Am I?",
	"Welcome to San Andreas!",
	"I don't give a fuck!",
	"I'm a lunatic, fool!",
	"I'm a well dressed maniac, fool!",
	"Bitches!",
	"Welcome to Ganton, fool!",
	"Bang, bang, bang!",
	"Asshole!",
	"You should of stayed at home.",
	"You tricks!",
	"Punk ass bitches!",
	"You mark!",
	"G.S.F. for life!",
	"Screw you, asshole!",
	"You should have stayed the fuck out of my way!",
	"This should help keep my belly full. ",
	"Lunch money. ",
	"I should spend this on a good meal. ",
	"You got anything else?",
	"Gimme that Grip!",
	"Aww you so kind.",
	"Ah, look at all these papers!",
	"That's very kind of you.",
	"Thanks, bitch!",
	"Next time you should put it in the bank!",
	"I'll never turn down paper.",
	"I need this more than you, I think.",
	"You've got a problem, fool!",
	"Is that all you got? Cheap fool!",
	"Give me them dollars!",
	"Come on, pay me! I'm a real criminal.",
	"I'll take that paper.",
	"That's going in my retirement fund!",
	"Is this for me?",
	"Gimme those chips!",
	"I need that!",
	"Yeah, them <i>my</i> ducats <i>now</i>.",
	"This ain't too classy.",
	"I like to share too, thank you.",
	"You don't need it, now.",
	"That's cool.",
	"That's my cash now, fool!",
	"Now don't move!",
	"Just a little money for CJ!",
	"Hey, pay a nigga!",
	"Send 'em ducats over here!",
	"Ah, you don't need it no more, huh!?",
	"With a mean bastard like you<b> </b>around here, crime really don't pay.",
	"Hey, gimme that paper!",
	"Hey, I'll have that now.",
	"Thanks!",
	"I guess it's true, you can't take it with you.",
	"Ah, thanks!",
	"Can I have this?",
	"You don't need it now.",
	"This the best you got?",
	"Life's a bitch, huh?",
	"You don't need it no more, huh?",
	"Oh, you're too kind.",
	"It's going to a good cause. Me.",
	"Give me that!",
	"You're fucking me off!",
	"Just some old kicks baby.",
	"Straight. Good looking.",
	"Thank you. I love these shoes, you know?",
	"Fo' sho, thanks.",
	"Yeah, they gangster ain't they?",
	"This how Grove street kick it.",
	"Thanks gangster.",
	"I'm glad you like it.",
	"They ain't nothing.",
	"I know these shoes are tight, huh?",
	"I'm dressed, but I'm still a G.",
	"Yeah, real nice, ain't they?",
	"Cool, I appreciate that.",
	"Yeah, these shoes are alright.",
	"Yeah cool, I take care of my feet, you know?",
	"Alright. Straight gangster. Know what I'm saying?",
	"My G'ed up apparel? Yeah, straight.",
	"Whatch you expect, huh?",
	"Pure gangster, homie.",
	"They just tat's fool. But I got the heart to back them up.",
	"Yeah pretty gangster ain't they.",
	"Yeah i hear you, cool.",
	"The tattoos? You like it? Cool.",
	"I know I love this work.",
	"The ink? It's a'ight, I know.",
	"The tattoo? You like it? Cool.",
	"My cut? You real kind.",
	"This cut? Awe thanks.",
	"Awe, thanks G.",
	"Yeah, good lookin'.",
	"Yeah, I've been lifting a lot, you know.",
	"I lift weights because I find life empty.",
	"Yeah. Cool. Fo' sho. Thank you.",
	"Yeah, I lift weights cuz I'm all empty inside.",
	"Awe, thanks.",
	"Straight.",
	"Yeah, I've been working out.",
	"Yeah, I've been down to gym a lot.",
	"Yeah, I've been wasting my life down the gym.",
	"My smell? Cool.",
	"Yeah, my cologne is tight.",
	"My cologne, cool? Thank you.",
	"Won't you fuck off, punk!",
	"Go tell daddy, bitch!",
	"Yeah, and you just a punk ass bitch.",
	"...And you're a fool, so we straight!",
	"My God, you're one ugly bitch!",
	"Yeah, but you're dating a ho, trick!",
	"Yeah? But I ain't never seen a talking horse before!",
	"At least kids don't cry when they see my face, bitch!",
	"What you doing running around without a bag on your face!?",
	"Put yourself on a make up, lady. A LOT of make up!",
	"Hey, somebody lost the dog?",
	"I had a dog once that looked a lot like you.",
	"And I don't wanna be like you, ugly and a punk.",
	"Yeah, you're stupid and ugly, bitch!",
	"Yeah, at least my women don't look like no pig.",
	"At least I look like the right gender, bitch!",
	"Wow, a camel that speaks!",
	"It's no wonder you ain't got no man, bitch!",
	"You know what? Fuck you.",
	"You should wear a bag on your face!",
	"Yeah? But you look like an asshole so we even.",
	"Ain't you got a hole to stick your head in?",
	"No wonder you're sad and lonely, woman.",
	"Yeah? And you're one ugly bitch.",
	"Hey, it's a zombie running about!",
	"Wow, I just seen for myself ugliest woman in the world!",
	"Wow, horse that walks on two legs!",
	"Shouldn't be you're on all fours?",
	"You're lame, little peewee!",
	"Look at you, little asshole.",
	"You should eat more, fool.",
	"Shut your mouth and bang your mama.",
	"Hey little man, go be little someplace else!",
	"Yeah? I bet you still wet the bet.",
	"You're incredible! ...Incredibly boring.",
	"Well why don't you go and join the line to bang your mama!",
	"Whatever, bitch!",
	"Lady, you crapping my style with your ugly face!",
	"Get lost lady, you got a few problems of your own!",
	"Get out of here, bitch!",
	"I ain't interested in your shit, bitch!",
	"That's amazing, a talking dog!",
	"Come here and say that, bitch!",
	"You talk well for somebody who's inbred!",
	"Shouldn't you wear a bell around your neck!?",
	"Hey, I've seen harder-looking puddings than you!",
	"Yeah, you're ugly, but your personality is even worse!",
	"Mind your own fucking business, lady!",
	"Ugly bitch!",
	"Leave me the hell alone, bitch!",
	"You chump, look at you!",
	"Shut your mouth, punk!",
	"Shut up, little bitch.",
	"You just a bitch!",
	"Like I give a fuck about what you think!",
	"Fuck you, buster!",
	"Yeah, me!? Look at you, pathetic bitch!",
	"Who the fuck you think you're talking to, bitch!?",
	"Fuck you, punk-ass bitch!",
	"Screw you, asshole!",
	"I thought you never asked!",
	"You read my mind, lady. Get in!",
	"You don't have no VD do you bitch?",
	"Sure baby, I like to party!",
	"Oh, for sure, get in!",
	"Yeah, I'd like a good time...",
	"Hey, a man's got a need, but I ain't that desperate!",
	"You ain't my tipe. See, I like woman.",
	"No way.",
	"N-O! Now, can you spell?",
	"Hey, I don't need to pay for a friendship.",
	"I ain't into animals, lady!",
	"Me, pay you? Get the fuck outta here.",
	"No way, ho.",
	"I ain't into bestiality, woman.",
	"You gotta be kidding me!",
	"Get the fuck outta here!",
	"N-O.",
	"I'm gonna put in some work, you down?",
	"You want to represent our set, homie?",
	"Homie, let's go roll down some buster.",
	"For the Grove dude, we're gonna put some work in.",
	"C'mon dude, roll wit' me.",
	"C'mon man, we're families.",
	"C'mon, Grove Street 4 Life dawg.",
	"Dude you ain't no Balla or no buster.",
	"You down bitch, or what?",
	"Cee-jayy!",
	"Show me some money!",
	"Look at me now!",
	"We shutting down Cluckin' Bells, tonight!",
	"Look at me and weep, assholes. I won.",
	"Look at me and weep!",
	"I win!",
	"Look at that, I won!",
	"Oh yeah!",
	"Give me some money!",
	"Hell yeah!",
	"Look at that!",
	"Yeah, I can't wait to celebrate with a meal.",
	"Big daddy wins again!",
	"Pay the big man!",
	"Now I can have a good meal!",
	"My arteries are thick as my wallet, now!",
	"Pay the gangsta!",
	"What's up, me!? I'm up!",
	"Well, look at that.",
	"Y'all cheating!",
	"Give a playa a break!",
	"What is this shit!?",
	"You're cheating! I know it!",
	"That's completely utter bullshit.",
	"What the fuck?",
	"Damn!",
	"I can't believe this shit, man!",
	"Look at this bullshit!",
	"That's straight bullshit!",
	"Argh, look at that shit!",
	"I'm so upset I gotta eat.",
	"Oh shit, my breakfast money!",
	"Oh damn, i lost my waffle money.",
	"I can't believe that, Imma have to eat to commiserate.",
	"Girl you dropped a bomb on me!",
	"Express yourself! Because I'm from Grove Street!",
	"Never gonna get it, never gonna get it, beyatch!",
	"Young hearts be free tonigh...! Aww, what was it...",
	"Rollercoaster...! Wahoo hoo hoo!",
	"Ain't nuttin but a G thang, bab-ay...",
	"Two loc'ed out niggas going crazy...",
	"Move your body, move your body...",
	"Take your time and let me love you good...",
}

func (cm *CommandManager) commandCJQuote(
	args string,
	message discordgo.Message,
	settings types.CommandSettings,
) (
	context bool,
	err error,
) {
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	cm.Discord.ChannelMessageSend(message.ChannelID, quote)
	return
}

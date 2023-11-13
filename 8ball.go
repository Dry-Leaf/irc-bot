package main

import (
    "math/rand"
    "regexp"
    
    "github.com/thoj/go-ircevent"
)

var eightBallAnswers = map[int]string{
    0: "It is certain.", 1: "It is decidedly so.",
    2: "Yes definitely.", 3: "You may rely on it.",
    4: "As I see it, yes.", 5: "Most likely.",
    6: "Outlook good.", 7: "Yes.",
    8: "Signs point to yes.", 9: "Reply hazy, try again.",
    10: "Ask again later.", 11: "Better not tell you now.",
    12: "Cannot predict now.", 13: "Concentrate and ask again.",
    14: "Don't count on it.", 15: "My reply is no.",
    16: "My sources say no.", 17: "Outlook not so good.",
    18: "Very doubtful.", 19: "Without a doubt."}

var fortuneAnswers = map[int]string{
    0: "A dubious friend may be an enemy in camouflage.", 1: "A faithful friend is a strong defense.",
    2: "A feather in the hand is better than a bird in the air.",
    3: "A friend asks only for your time not your money.",
    4: "A gambler not only will lose what he has, but also will lose what he doesn’t have.",
    5: "A golden egg of opportunity falls into your lap this month.",
    6: "A fresh start will put you on your way.",
    7: "A good friendship is often more important than a passionate romance.",
    8: "A good time to finish up old tasks.",
    9: "A hunch is creativity trying to tell you something.",
    10: "A light heart carries you through all the hard times.",
    11: "A new outlook brightens your image and brings new friends.",
    12: "A new perspective will come with the new year.",
    13: "A person is never too old to learn.",
    14: "A person of words and not deeds is like a garden full of weeds.",
    15: "A pleasant surprise is waiting for you.",
    16: "A short pencil is usually better than a long memory any day.",
    17: "A small donation is call for. It’s the right thing to do.",
    18: "A smile is your personal welcome mat.",
    19: "A smooth long journey! Great expectations.",
    20: "A soft voice may be awfully persuasive.",
    21: "A truly rich life contains love and art in abundance.",
    22: "Accept something that you cannot change, and you will feel better.",
    23: "Adventure can be real happiness.",
    24: "Advice is like kissing. It costs nothing and is a pleasant thing to do.",
    25: "Advice, when most needed, is least heeded.",
    26: "All the effort you are making will ultimately pay off.",
    27: "All the troubles you have will pass away very quickly.",
    28: "All will go well with your new project.",
    29: "All your hard work will soon pay off.",
    30: "Allow compassion to guide your decisions.",
    31: "An acquaintance of the past will affect you in the near future.",
    32: "An agreeable romance might begin to take on the appearance.",
    33: "An important person will offer you support.",
    34: "An inch of time is an inch of gold.",
    35: "Any day above ground is a good day.",
    36: "Any decision you have to make tomorrow is a good decision.",
    37: "At the touch of love, everyone becomes a poet.",
    38: "Be careful or you could fall for some tricks today.",
    39: "Beauty in its various forms appeals to you.",
    40: "Because you demand more from yourself, others respect you deeply.",
    41: "Believe in yourself and others will too.",
    42: "Believe it can be done.",
    43: "Pantsu.",
    44: "Better ask twice than lose yourself once.",
    45: "Bide your time, for success is near.",
    46: "Carve your name on your heart and not on marble.",
    47: "Chance favors those in motion.",
    48: "Change is happening in your life, so go with the flow!",
    49: "Competence like yours is underrated.",
    50: "Congratulations! You are on your way.",
    51: "Could I get some directions to your heart?",
    52: "Courtesy begins in the home.", 53: "Courtesy is contagious.",
    54: "Curiosity kills boredom. Nothing can kill curiosity.",
    55: "Dedicate yourself with a calm mind to the task at hand.",
    56: "Depart not from the path which fate has you assigned.",
    57: "Determination is what you need now.",
    58: "Diligence and modesty can raise your social status.",
    59: "Disbelief destroys the magic.",
    60: "Distance yourself from the vain.",
    61: "Do not be intimidated by the eloquence of others.",
    62: "Do not demand for someone’s soul if you already got his heart.",
    63: "Do not let ambitions overshadow small success.",
    64: "Do not make extra work for yourself.",
    65: "Do not underestimate yourself. Human beings have unlimited potentials.",
    66: "Do you know that the busiest person has the largest amount of time?",
    67: "Don’t be discouraged, because every wrong attempt discarded is another step forward.",
    68: "Don’t confuse recklessness with confidence.",
    69: "Don’t expect romantic attachments to be strictly logical or rational.",
    70: "Don’t just spend time. Invest it.",
    71: "Don’t just think, act!",
    72: "Don’t let friends impose on you, work calmly and silently.",
    73: "Don’t let the past and useless detail choke your existence.",
    74: "Don’t let your limitations overshadow your talents.",
    75: "Don’t worry; prosperity will knock on your door soon.",
    76: "Each day, compel yourself to do something you would rather not do.",
    77: "Education is the ability to meet life’s situations.",
    78: "Embrace this love relationship you have!",
    79: "Emulate what you admire in your parents.",
    80: "Emulate what you respect in your friends.",
    81: "Every flower blooms in its own sweet time.",
    82: "Every wise man started out by asking many questions.",
    83: "Everyday in your life is a special occasion.",
    84: "Everywhere you choose to go, friendly faces will greet you.",
    85: "Expect much of yourself and little of others.",
    86: "Failure is the chance to do better next time.",
    87: "Failure is the path of lease persistence.",
    88: "Fear and desire – two sides of the same coin.",
    89: "Fearless courage is the foundation of victory.",
    90: "Feeding a cow with roses does not get extra appreciation.",
    91: "First think of what you want to do; then do what you have to do.",
    92: "Follow the middle path. Neither extreme will make you happy.",
    93: "For hate is never conquered by hate. Hate is conquered by love.",
    94: "For the things we have to learn before we can do them, we learn by doing them.",
    95: "Fortune Not Found: Abort, Retry, Ignore?",
    96: "From listening comes wisdom and from speaking repentance.",
    97: "From now on your kindness will lead you to success.",
    98: "Get your mind set – confidence will lead you on.",
    99: "Get your mind set…confidence will lead you on.",
    100: "Go for the gold today! You’ll be the champion of whatever.",
    101: "Go take a rest; you deserve it.",
    102: "Good news will be brought to you by mail.",
    103: "Good news will come to you by mail.",
    104: "Good to begin well, better to end well.",
    105: "Happiness begins with facing life with a smile and a wink.",
    106: "Happiness will bring you good luck.",
    107: "Happy life is just in front of you.",
    108: "Hard words break no bones, fine words butter no parsnips.",
    109: "Have a beautiful day.",
    110: "He who expects no gratitude shall never be disappointed.",
    111: "He who knows he has enough is rich.",
    112: "He who knows others is wise. He who knows himself is enlightened.",
    113: "Help! I’m being held prisoner in a chinese bakery!",
    114: "How many of you believe in psycho-kinesis? Raise my hand.",
    115: "How you look depends on where you go.",
    116: "I learn by going where I have to go.",
    117: "If a true sense of value is to be yours it must come through service.",
    118: "If certainty were truth, we would never be wrong.",
    119: "If you continually give, you will continually have.",
    120: "If you look in the right places, you can find some good offerings.",
    121: "If you think you can do a thing or think you can’t do a thing, you’re right.",
    122: "If you wish to see the best in others, show the best of yourself.",
    123: "If your desires are not extravagant, they will be granted.",
    124: "If your desires are not to extravagant they will be granted.",
    125: "If you’re feeling down, try throwing yourself into your work.",
    126: "Imagination rules the world.",
    127: "In order to take, one must first give.",
    128: "In the end all things will be known.",
    129: "In this world of contradiction, It’s better to be merry than wise.",
    130: "It could be better, but its [sic] good enough.",
    131: "It is better to be an optimist and proven a fool than to be a pessimist and be proven right.",
    132: "It is better to deal with problems before they arise.",
    133: "It is honorable to stand up for what is right, however unpopular it seems.",
    134: "It is worth reviewing some old lessons.",
    135: "It takes courage to admit fault.",
    136: "It’s not the amount of time you devote, but what you devote to the time that counts.",
    137: "It’s time to get moving. Your spirits will lift accordingly.",
    138: "Keep your face to the sunshine and you will never see shadows.",
    139: "Let the world be filled with tranquility and goodwill.",
    140: "Like the river flow into the sea. Something are just meant to be.",
    141: "Listen not to vain words of empty tongue.",
    142: "Listen to everyone. Ideas come from everywhere.",
    143: "Living with a commitment to excellence shall take you far.",
    144: "Long life is in store for you.",
    145: "Love is a warm fire to keep the soul warm.",
    146: "Love is like sweet medicine, good to the last drop.",
    147: "Love lights up the world.",
    148: "Love truth, but pardon error.",
    149: "Man is born to live and not prepared to live.",
    150: "Man’s mind, once stretched by a new idea, never regains it’s original dimensions.",
    151: "Many will travel to hear you speak.",
    152: "Meditation with an old enemy is advised.",
    153: "Miles are covered one step at a time.",
    154: "Nature, time and patience are the three great physicians.",
    155: "Never fear! The end of something marks the start of something new.",
    156: "New ideas could be profitable.",
    157: "New people will bring you new realizations, especially about big issues.",
    158: "No one can walk backwards into the future.",
    159: "Nothing is more difficult, and therefore more precious, than to be able to decide.",
    160: "Now is a good time to buy stock.",
    161: "Now is the time to go ahead and pursue that love interest!",
    162: "Now is the time to try something new",
    163: "Now is the time to try something new.",
    164: "Observe all men, but most of all yourself.",
    165: "One can never fill another’s shoes, rather he must outgrow the old shoes.",
    166: "One of the first things you should look for in a problem is its positive side.",
    167: "Others can help you now.",
    168: "Pennies from heaven find their way to your doorstep this year!",
    169: "People are attracted by your Delicate features.",
    170: "People find it difficult to resist your persuasive manner.",
    171: "Perhaps you’ve been focusing too much on saving.",
    172: "Physical activity will dramatically improve your outlook today.",
    173: "Pick battles big enough to matter, small enough to win.",
    174: "Place special emphasis on old friendship.",
    175: "Please visit us at www.wontonfood.com",
    176: "Po Says: Pandas like eating bamboo, but I prefer mine dipped in chocolate.",
    177: "Practice makes perfect.",
    178: "Protective measures will prevent costly disasters.",
    179: "Put your mind into planning today. Look into the future.",
    180: "Remember the birthday but never the age.",
    181: "Remember to share good fortune as well as bad with your friends.",
    182: "Rest has a peaceful effect on your physical and emotional health.",
    183: "Resting well is as important as working hard.",
    184: "Romance moves you in a new direction.",
    185: "Savor your freedom – it is precious.",
    186: "Say hello to others. You will have a happier day.",
    187: "Self-knowledge is a life long process.",
    188: "Share your joys and sorrows with your family.",
    189: "Sift through your past to get a better idea of the present.",
    190: "Sloth makes all things difficult; industry all easy.",
    191: "Small confidences mark the onset of a friendship.",
    192: "Smile when you are ready.",
    193: "Society prepares the crime; the criminal commits it.",
    194: "Someone you care about seeks reconciliation.",
    195: "Soon life will become more interesting.",
    196: "Stand tall. Don’t look down upon yourself.",
    197: "Staying close to home is going to be best for your morale today",
    198: "Stop searching forever, happiness is just next to you.",
    199: "Strong reasons make strong actions.",
    200: "Success is a journey, not a destination.",
    201: "Success is failure turned inside out.",
    202: "Success is going from failure to failure without loss of enthusiasm.",
    203: "Swimming is easy. Stay floating is hard.",
    204: "Take care and sensitivity you show towards others will return to you.",
    205: "Take the high road.",
    206: "Technology is the art of arranging the world so we do not notice it.",
    207: "The austerity you see around you covers the richness of life like a veil.",
    208: "The best prediction of future is the past.",
    209: "The change you started already have far-reaching effects. Be ready.",
    210: "The change you started already have far-reaching effects. Be ready.",
    211: "The first man gets the oyster, the second man gets the shell.",
    212: "The greatest achievement in life is to stand up again after falling.",
    213: "The harder you work, the luckier you get.",
    214: "The night life is for you.",
    215: "The one that recognizes the illusion does not act as if it is real.",
    216: "The only people who never fail are those who never try.",
    217: "The person who will not stand for something will fall for anything.",
    218: "The philosophy of one century is the common sense of the next.",
    219: "The saints are the sinners who keep on trying.",
    220: "The secret to good friends is no secret to you.",
    221: "The small courtesies sweeten life, the greater ennoble it.",
    222: "The smart thing to do is to begin trusting your intuitions.",
    223: "The strong person understands how to withstand substantial loss.",
    224: "The sure way to predict the future is to invent it.",
    225: "The truly generous share, even with the undeserving.",
    226: "The value lies not within any particular thing, but in the desire placed on that thing.",
    227: "The weather is wonderful.",
    228: "There is a time for caution, but not for fear.",
    229: "There is no mistake so great as that of being always right.",
    230: "There is no wisdom greater than kindness.",
    231: "There is not greater pleasure than seeing your lived  ones prosper.",
    232: "There’s no such thing as an ordinary cat.",
    233: "Things don’t just happen; they happen just.",
    234: "Those who care will make the effort.",
    235: "Time and patience are called for many surprises await you!.",
    236: "Time is precious, but truth is more precious than time",
    237: "To know oneself, one should assert oneself.",
    238: "To the world you may be one person, but to one person you may be the world.",
    239: "Today is the conserve yourself, as things just won’t budge.",
    240: "Today, your mouth might be moving but no one is listening.",
    241: "Tonight you will be blinded by passion.",
    242: "Use your eloquence where it will do the most good.",
    243: "We first make our habits, and then our habits make us.",
    244: "Welcome change.",
    245: "“Welcome” is a powerful word.",
    246: "Well done is better than well said.",
    247: "What’s hidden in an empty box?",
    248: "What’s that in your eye? Oh…it’s a sparkle.",
    249: "What’s yours in mine, and what’s mine is mine.",
    250: "When more become too much. It’s same as being not enough.",
    251: "When your heart is pure, your mind is clear.",
    252: "Wish you happiness.",
    253: "With age comes wisdom.",
    254: "You  adventure could lead to happiness.",
    255: "You always bring others happiness.",
    256: "You are a person of another time.",
    257: "You are a talented storyteller.",
    258: "You are admired by everyone for your talent and ability.",
    259: "You are almost there.",
    260: "You are busy, but you are happy.",
    261: "You are generous to an extreme and always think of the other fellow.",
    262: "You are going to have some new clothes.",
    263: "You are in good hands this evening.",
    264: "You are interested in higher education, whether material or spiritual.",
    265: "You are modest and courteous.",
    266: "You are never selfish with your advice or your help.",
    267: "You are next in line for promotion in your firm.",
    268: "You are offered the dream of a lifetime. Say yes!",
    269: "You are open-minded and quick to make new friends.",
    270: "You are solid and dependable.",
    271: "You are soon going to change your present line of work.",
    272: "You are talented in many ways.",
    273: "You are the master of every situation.",
    274: "You are very expressive and positive in words, act and feeling.",
    275: "You are working hard.",
    276: "You begin to appreciate how important it is to share your personal beliefs.",
    277: "You can keep a secret.",
    278: "You can see a lot just by looking.",
    279: "You can’t steal second base and keep your foot on first.",
    280: "You desire recognition and you will find it.",
    281: "You have a deep appreciation of the arts and music.",
    282: "You have a deep interest in all that is artistic.",
    283: "You have a friendly heart and are well admired.",
    284: "You have a shrewd knack for spotting insincerity.",
    285: "You have a yearning for perfection.",
    286: "You have an active mind and a keen imagination.",
    287: "You have an ambitious nature and may make a name for yourself.",
    288: "You have an unusual equipment for success, use it properly.",
    289: "You have exceeded what was expected.",
    290: "You have had a good start. Work harder!",
    291: "You have the power to write your own fortune.",
    292: "You have yearning for perfection.",
    293: "You know where you are going and how to get there.",
    294: "You look pretty.",
    295: "You love challenge.",
    296: "You love chinese food.",
    297: "You make people realize that there exist other beauties in the world.",
    298: "You never hesitate to tackle the most difficult problems.",
    299: "You never know who you touch.",
    300: "You only treasure what you lost.",
    301: "You seek to shield those you love and like the role of provider.",
    302: "You should be able to make money and hold on to it.",
    303: "You should be able to undertake and complete anything.",
    304: "You should pay for this check. Be generous.",
    305: "You understand how to have fun with others and to enjoy your solitude.",
    306: "You will always be surrounded by true friends.",
    307: "You will always get what you want through your charm and personality.",
    308: "You will always have good luck in your personal affairs.",
    309: "You will be a great success both in the business world and society.",
    310: "You will be blessed with longevity.",
    311: "You will be pleasantly surprised tonight.",
    312: "You will be sharing great news with all the people you love.",
    313: "You will be successful in your work.",
    314: "You will be traveling and coming into a fortune.",
    315: "You will be unusually successful in business.",
    316: "You will become a great philanthropist in your later years.",
    317: "You will become more and more wealthy.",
    318: "You will enjoy good health.",
    319: "You will enjoy good health; you will be surrounded by luxury.",
    320: "You will find great contentment in the daily, routine activities.",
    321: "You will have a fine capacity for the enjoyment of life.",
    322: "You will have gold pieces by the bushel.",
    323: "You will inherit a large sum of money.",
    324: "You will make change for the better.",
    325: "You will soon be surrounded by good friends and laughter.",
    326: "You will take a chance in something in near future.",
    327: "You will travel far and wide, both pleasure and business.",
    328: "You will travel far and wide,both pleasure and business.",
    329: "Your abilities are unparalleled.",
    330: "Your ability is appreciated.",
    331: "Your ability to juggle many tasks will take you far.",
    332: "Your biggest virtue is your modesty.",
    333: "Your character can be described as natural and unrestrained.",
    334: "Your difficulties will strengthen you.",
    335: "Your dreams are never silly; depend on them to guide you.",
    336: "Your dreams are worth your best efforts to achieve them.",
    337: "Your energy returns and you get things done.",
    338: "Your family is young, gifted and attractive.",
    339: "Your first love has never forgotten you.",
    340: "Your goal will be reached very soon.",
    341: "Your happiness is before you, not behind you! Cherish it.",
    342: "Your hard work will payoff today.",
    343: "Your heart will always make itself known through your words.",
    344: "Your home is the center of great love.",
    345: "Your ideals are well within your reach.",
    346: "Your infinite capacity for patience will be rewarded sooner or later.",
    347: "Your leadership qualities will be tested and proven.",
    348: "Your life will be happy and peaceful.",
    349: "Your life will get more and more exciting.",
    350: "Your love life will be happy and harmonious.",
    351: "Your love of music will be an important part of your life.",
    352: "Your loyalty is a virtue, but not when it’s wedded with blind stubbornness.",
    353: "Your mentality is alert, practical, and analytical.",
    354: "Your mind is creative, original and alert.",
    355: "Your mind is your greatest asset.",
    356: "Your moods signal a period of change.",
    357: "Your quick wits will get you out of a tough situation.",
    358: "Your reputation is your wealth.",
    359: "Your success will astonish everyone.",
    360: "Your talents will be recognized and suitably rewarded.",
    361: "Your work interests can capture the highest status or prestige.",
    362: "Error 404: Fortune not found.",
    363: "Two days from now, tomorrow will be yesterday.",
    364: "You will soon have an out-of-money experience.",
    365: "When working towards the solution of a problem, it always helps if you know the answer.",
    367: "You are cleverly disguised as a responsible adult.",
    368: "Be cautious while walking in the darkness alone.",
    369: "I see money in your future… it is not yours though.",
    370: "Hard work pays off in the future. Laziness pays off now.",
}

var eightBallReg = regexp.MustCompile(`(?i)\A\.8ball(?:\s+|\z)`)
var fortuneReg = regexp.MustCompile(`(?i)\A\.fortune(?:\s+|\z)`)

func EightBall(stored, ch string, conn *irc.Connection) {
    if eightBallReg.MatchString(stored) {
        reply := eightBallAnswers[rand.Intn(20)]
        conn.Privmsg(ch, reply)
    }

    if fortuneReg.MatchString(stored) {
        reply := fortuneAnswers[rand.Intn(371)]
        conn.Privmsg(ch, reply)
    }
}

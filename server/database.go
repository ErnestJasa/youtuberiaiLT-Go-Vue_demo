package main

type Channel struct {
	Id              string
	Title           string
	CustomUrl       string
	Description     string
	SubscriberCount uint64
	VideoCount      uint64
	Categories      []string
}

var Channels = []Channel{
	{
		Id:              "UCj5sdOdmY70LG1tIQxgk6nA",
		Title:           "Omantas",
		CustomUrl:       "/@OmantasJokubaitis",
		Description:     `"NEJUOKINGA" Pokalbinė yra laida, kuri suburia du žmones atviram bei nuoširdžiam pokalbiui - kaip namuose.`,
		SubscriberCount: 54500,
		VideoCount:      199,
		Categories: []string{
			"PODCAST",
		},
	},
	{
		Id:              "UCef7VJO5aZ80ZBsTy5cyrBA",
		Title:           "Overcrow",
		CustomUrl:       "/@Overcrow",
		Description:     `overcrowding with degeneracy`,
		SubscriberCount: 1910,
		VideoCount:      31,
		Categories: []string{
			"IVAIRUS",
			"ZAIDIMAI",
			"ANIMACIJA",
		},
	},
	{
		Id:              "UCV_59afytAhjWv9WWeixTFg",
		Title:           "Kitokie pasikalbėjimai",
		CustomUrl:       "/@pasikalbejimai",
		Description:     `„Kitokie pasikalbėjimai“ – vėlyvojo vakaro pokalbių šou`,
		SubscriberCount: 160000,
		VideoCount:      1000,
		Categories: []string{
			"PODCAST",
		},
	},
	{
		Id:              "UC1D_O7QCx4pvx-hHv8yuYZg",
		Title:           "Pergalė",
		CustomUrl:       "/@pergale7421",
		Description:     `Viskas tik dėl šokolado! Kiekviena plytelė, saldainis, desertas, batonėlis, kurį pagaminame, yra unikalus ir mums pats svarbiausias. Prisijunk ir susipažink su tikrosiomis „Pergalės“ šokolado žvaigždėmis`,
		SubscriberCount: 1100,
		VideoCount:      47,
		Categories: []string{
			"IVAIRUS",
			"PREKINIS ZENKLAS",
		},
	},
	{
		Id:        "UC6rlTjdr8cqk4IJe1GoWukQ",
		Title:     "PildykLT",
		CustomUrl: "/@PildykLT",
		Description: `Čia tavęs laukia daugybė PILDYK YouTuber’ių nuotykių: Talžūnas, Stimo, Aqva ir Demino nuolat dalyvauja įvairiuose iššūkiuose, filmuoja vlogus, geimina, žaidžia kitus žaidimus, atsakinėja į aštrius klausimus, daro LIVE transliacijas ir kitus smagius video.

`,
		SubscriberCount: 162000,
		VideoCount:      2200,
		Categories: []string{
			"VLOG",
			"HUMORAS",
			"PREKINIS ZENKLAS",
		},
	},
	{
		Id:              "UCB84fyfY4-FF-eDEfJ69kMg",
		Title:           "PixelSimkE",
		CustomUrl:       "/@pixelsimke7210",
		Description:     ``,
		SubscriberCount: 48,
		VideoCount:      7,
		Categories: []string{
			"RECENZIJOS",
		},
	},
	{
		Id:              "UCJ5WlR_2Wddnrk-lr3kkg2w",
		Title:           "PonasAkiniuotis",
		CustomUrl:       "/@PonasAkiniuotis",
		Description:     `ponasakiniuotis.business@gmail.com`,
		SubscriberCount: 48200,
		VideoCount:      374,
		Categories: []string{
			"ZAIDIMAI",
			"MINECRAFT",
		},
	},
	{
		Id:        "UCKeHmmYwZZJAg8zS6mZWLIQ",
		Title:     "Povilo Kanalas",
		CustomUrl: "/@PoviloKanalas",
		Description: `Hello Would You Like to Know More About Me? Read Below.
My Name Is Paul
I am 17 years old
I like to play games and film them and upload them to YouTube.
------------------------------------------------- ---------------------------------------------
What Programs Do I Use?
For filming - Bandicam and Shadowplay
For streaming - Obs
Video Editing - Pixlr.com
Video Editing - Sony Vegas Pro 17
For Audio Editing - Audacity
------------------------------------------------- ---------------------------------------------
My Contact Email: povilastires132@gmail.com
------------------------------------------------- --------------------
Computer Specifications
CPU - Amd Ryzen 5 2600 3.4 GHz
Ram - 8 Gb DDR4 T-Force
GPU - GTX 1660 TI OC 6 Gb
Motherboard - Asus Prime B450M-K`,
		SubscriberCount: 10400,
		VideoCount:      235,
		Categories: []string{
			"IVAIRUS",
			"ZAIDIMAI",
			"VLOG",
		},
	},
	{
		Id:              "UCJyL6oBXV1XNpPVF8_gkHHg",
		Title:           "pragaras",
		CustomUrl:       "/@pragaras",
		Description:     `pragaro karaliaus perliukai, clipukai ir visokios kitokios nesamones`,
		SubscriberCount: 826,
		VideoCount:      59,
		Categories: []string{
			"IVAIRUS",
		},
	},
	{
		Id:        "UCra6CoauJ0zgggbwoN8ww7A",
		Title:     "prefis",
		CustomUrl: "/@prefislt",
		Description: `Filmuoju žaidimukus
`,
		SubscriberCount: 33300,
		VideoCount:      795,
		Categories: []string{
			"ZAIDIMAI",
		},
	},
	{
		Id:              "UCmEKDSs3dB3rRaIaVEEGmyA",
		Title:           "Valentinas",
		CustomUrl:       "/@PressButtonTVLT",
		Description:     `Sveiki. Aš esu Valentinas ir domiuosi video žaidimais bei kinu. Mano kanale pamatysite vaizdo įrašus, kuriuose dalinuose savo nuomone apie šias meno formas.`,
		SubscriberCount: 7280,
		VideoCount:      167,
		Categories: []string{
			"RECENZIJOS",
			"EDUKACIJA",
			"ZAIDIMAI",
		},
	},
	{
		Id:              "UCOGHTLkrwzzckcebvaF8njQ",
		Title:           "Proflame",
		CustomUrl:       "/@proflame2122",
		Description:     ``,
		SubscriberCount: 23200,
		VideoCount:      48,
		Categories: []string{
			"MUZIKA",
		},
	},
	{
		Id:              "UCe22CG3JAXvpP19aYO71bGg",
		Title:           "Prologai",
		CustomUrl:       "/@prologai3443",
		Description:     `Prologai – kasmėnesinė tinklalaidė apie kritinės minties teoriją ir kairiąją politiką su Jokūbu Sąlyga. Jos tikslas pristatyti kairiųjų autorių darbus, aptarti šiandien aktualias analizės prizmes ir teorinius debatus, bei prisidėti prie progresyvių politinių praktikų. Naujus epizodus galite klausyti kiekvieną paskutinįjį mėnesio penktadienį. `,
		SubscriberCount: 136,
		VideoCount:      6,
		Categories: []string{
			"PODCAST",
		},
	},
	{
		Id:              "UC681PVzbyqGOxS4_p23Oosw",
		Title:           "Pro Patria",
		CustomUrl:       "/@ProPatria1",
		Description:     ``,
		SubscriberCount: 8000,
		VideoCount:      209,
		Categories: []string{
			"EDUKACIJA",
			"PODCAST",
		},
	},
	{
		Id:        "UC-rrTSS2Z0Nsa3Oc7dSWRQA",
		Title:     "PUKELIO",
		CustomUrl: "/@PUKELIO",
		Description: `Sveiki atvyke i mano kanala!
Kvieciu paziureti visus ka tik domina automobiliu turinys! Turini stengiuosi kelti kuo dazniau.

💎SOCIAL MEDIA💎
INSTAGRAM: https://www.instagram.com/pukeliooo/

`,
		SubscriberCount: 2960,
		VideoCount:      59,
		Categories: []string{
			"VLOG",
			"MASINOS",
		},
	},
	{
		Id:        "UC0Zw-isGKZFKuljDWWjjrXg",
		Title:     "punktò",
		CustomUrl: "/@punkto9973",
		Description: `- Los Ingredientes -

A stupid smile with a taste of blood in the mouth
-
A crunchy cocktail of social clichés
-
The fragility of fluffy reality
-
etc. / make up your own
-

`,
		SubscriberCount: 1590,
		VideoCount:      9,
		Categories: []string{
			"MUZIKA",
		},
	},
	{
		Id:        "UCv40SKmi_cc41mY_yOsfvwA",
		Title:     "Qbuilds",
		CustomUrl: "/@Qbuildslt",
		Description: `Labas, mes - QBuilds.

Kiekvienas mūsų kompiuteris surinktas su dideliu dėmesiu detalei ir kokybei. Ieškome geriausių sprendimų sukurti balansui tarp kokybės, funkcionalumo ir vizualų.

Mūsų kanale rasi geriausių rinkoje esančių kompiuterių apžvalgas, naujienas tech pasaulyje ir clean setup’us.

`,
		SubscriberCount: 3270,
		VideoCount:      208,
		Categories: []string{
			"RECENZIJOS",
			"IVAIRUS",
		},
	},
	{
		Id:              "UCU8AueUTR9QAmyqyvgcTBYg",
		Title:           "Radijo Ranča",
		CustomUrl:       "/@RadijoRanca",
		Description:     `Trys eterio kaubojai - Mindaugas Stasiulis, Gintas Vaičikauskas ir Saulius Žvirgždas susitinka Radijo Rančoje, kurioje ganosi laisvė, gera nuotaika, bajeris, žvengas ir palaidi liežuviai!`,
		SubscriberCount: 4790,
		VideoCount:      24,
		Categories: []string{
			"PODCAST",
		},
	},
	{
		Id:        "UCsPHjbmGQQXLBGfsxcRWS5g",
		Title:     "thomasas",
		CustomUrl: "/@thomasas1190",
		Description: `Thomso antras kanalas kuriame pilna šūdo.
Original kanalas: https://www.youtube.com/channel/UCDpdakUrU8QgIqFIShbe_7g`,
		SubscriberCount: 283,
		VideoCount:      17,
		Categories: []string{
			"IVAIRUS",
			"HUMORAS",
		},
	},
	{
		Id:              "UC-AGx7rM_mBqesECSVp-x8w",
		Title:           "Skirmantas Malinauskas",
		CustomUrl:       "/@tiesayrasvarbu",
		Description:     `Padėkite man kurti šį kanalą. Ačiū!`,
		SubscriberCount: 187000,
		VideoCount:      206,
		Categories: []string{
			"EDUKACIJA",
			"DOKUMENTIKA",
		},
	},
	{
		Id:        "UC1uXcCFBq4c1L1je3-SuWnQ",
		Title:     "Titukis",
		CustomUrl: "/@Titukis",
		Description: `➤ Support-a-creator kodas: Titukis

➤ Discord serveris: https://discord.gg/a7kNAfxQZY
➤ Instagram: https://www.instagram.com/titukisyt/

➦ Susisiekti galite: titastamu@gmail.com

Logo sukūrė Eftyen
`,
		SubscriberCount: 2290,
		VideoCount:      97,
		Categories: []string{
			"ZAIDIMAI",
			"MINECRAFT",
			"FORTNITE",
		},
	},
	{
		Id:        "UCvNTp3DQkb46WYLM58o009g",
		Title:     "Tomuxas",
		CustomUrl: "/@Tomuxas",
		Description: `⬇️Instagram : tomuxas__
`,
		SubscriberCount: 36500,
		VideoCount:      375,
		Categories: []string{
			"ZAIDIMAI",
			"IVAIRUS",
			"GTA",
		},
	},
	{
		Id:        "UCJBpu_3csr7ZsfTWbY11gAA",
		Title:     "Tortasis",
		CustomUrl: "/@Tortasis",
		Description: `pabuciuosiu jei pasubinsi

Pasiekimai:
1 prenumeratorius: ✅
100 prenumeratorių: ✅
500 prenumeratorių: ✅
1000 prenumeratorių: ✅
5000 prenumeratorių: ✅
10k prenumeratorių: ✅
15k prenumeratorių: ❌
20k prenumeratorių: ❌
Milijardas Begalybių prenumeratorių: ❌
`,
		SubscriberCount: 11200,
		VideoCount:      450,
		Categories: []string{
			"ZAIDIMAI",
			"IVAIRUS",
		},
	},
	{
		Id:        "UCdKdE8ddDpV9VoiICXolxSg",
		Title:     "TreciaLentyna",
		CustomUrl: "/@TreciaLentyna",
		Description: `Sveiki, aš vardu Arnas. O čia - mano kanalas, kuris tapo mano hobiu, o vėliau ir vienu pagrindinių užsiėmimų.

Čia rasite visko: nuo šiurpių, bauginančių, įdomių faktų ir dokumentikos iki šiek tiek linksmesnių video. Kiekvienam smalsuoliui! Kanalas jau peržengė 160 tūkstančių prenumeratorių ribą, dėl ko labai džiaugiuosi ir nekantrauju išvysti tai, kur mane tai nuves. Tikiuosi Jums tai suteiks tiek pat džiaugsmo, kiek suteikia man!

Ačiū, kad užsukote užmesti akį. Susitiksime naujame mano epizode!
`,
		SubscriberCount: 167000,
		VideoCount:      306,
		Categories: []string{
			"DOKUMENTIKA",
			"IVAIRUS",
		},
	},
	{
		Id:        "UCZnxnV3Lypuf8N6WOOyJJjQ",
		Title:     "Twisteris",
		CustomUrl: "/@Twisteris",
		Description: `● SUSISIEKITE SU MANIMI/CONTACT ME ●
twisteris.business@gmail.com

● INSTAGRAM ●
instagram.com/jonas.hofsteteris

● PC SPECIFIKACIJA ●
CPU: AMD Ryzen 7 3800X
GPU: MSI RTX 2080 Ti Gaming X Trio
RAM: G.SKILL TridentZ 32GB 3200MHz
Motherborad: MSI X570 MEG ACE
SSD/HDD: SAMSUNG 970 EVO Plus 1TB + WD Black 6TB

● ĮRANGA ●
Mouse: Logitech G502 PROTEUS SPECTRUM
Keyboard: Corsair K95 RGB PLATINUM (MX Speed)
Headphones: Sennheiser HD598 SR
Monitors: AORUS AD27QD + 2xAOC 27G2U
Microphone: RODE NT1
Camera: Sony A7 III

`,
		SubscriberCount: 22500,
		VideoCount:      293,
		Categories: []string{
			"ZAIDIMAI",
			"GTA",
			"COUNTER-STRIKE",
			"MINECRAFT",
		},
	},
	{
		Id:        "UCyNx5wHqvCXDK21Xe6eFaiQ",
		Title:     "UNBOX LT",
		CustomUrl: "/@UNBOXLT",
		Description: `Sveiki visi.
Mėgstu siuntinius, ypač nemokamus. Siunčiuosi, išpakuoju ir žiūriu, ką gavau ir dalinuosi video su Jumis.
Kam patiko, prenumeruokite, spauskite LIKE. Kam nepatiko irgi spauskite LIKE.
Nesu aš profesionalus apžvalgininkas, tad prašom neįsižeisti, jei ko neparodau kanale, ar kažko nepaaiškinu.
Nesu tarpininkas ir prekių neužsakinėju, komercija neužsiimu:)
Visos prekės, kurias rodau kanale, tai mano arba draugų užsakytos.

Unboxingltu@gmail.com`,
		SubscriberCount: 4710,
		VideoCount:      621,
		Categories: []string{
			"RECENZIJOS",
		},
	},
	{
		Id:              "UCBLvjJODX0mi5IfKAiq_JBA",
		Title:           "Undergroundinis Knygynas",
		CustomUrl:       "/@undergroundinisknygynas",
		Description:     ``,
		SubscriberCount: 4100,
		VideoCount:      552,
		Categories: []string{
			"DOKUMENTIKA",
			"HUMORAS",
		},
	},
	{
		Id:        "UC3Ccgql7ENMo9ZDup8V6yrQ",
		Title:     "Urbietis",
		CustomUrl: "/@urbietis1473",
		Description: `Urbietis - nepriklausomo turinio kūrėjas. Pagrindinė jo skleidžiama mintis yra apie pasaulio suvokimo dualumą. Nėra tik vienos "teisingos" pusės, o jei pastebima tik viena, vadinasi mąstyme dominuoja iškreiptos, susikurtos iliuzijos.

"Urbietis podcast" - greitai populiarėjanti laida Lietuvoje, jau pasiekusi daugiau nei 1 milijono lietuvių ausis. Laidoje paliečiamos temos gali prisidėti prie mąstymo pokyčių ir padaryti didelę įtaką žmogaus ateities veiksmams.

Taip pat, Urbietis plėtoja šviečiamąją veiklą socialiniuose tinkluose, o visus, norinčius įgyti praktinių žinių (mąstymo pokyčiams, elgesio ar įpročių transformavimui), kviečia prisijungti į Urbiečio akademiją (www.urbiecioakademija.lt).

"Eidamas per skaudžias patirtis, atradau savo gyvenimo kelią. Tos patirtys lydėjo ne vienus metus, bet tik pasitraukęs iš aukos pozicijos ir prisiėmęs atsakomybę už savo gyvenimą, aš pradėjau jį keisti. Šios laidos skirtos tau, mielas žiūrove, jeigu nori iš gyvenimo daugiau. Gero žiūrėjimo!"
`,
		SubscriberCount: 5740,
		VideoCount:      147,
		Categories: []string{
			"PODCAST",
		},
	},
	{
		Id:        "UC6Nv3jCkfsiZykM5YqKKQ-Q",
		Title:     "Vilnius Archdiocese Caritas",
		CustomUrl: "/@va-caritas",
		Description: `Vilnius Archdiocese Caritas is part of one of the world’s largest aid networks. Caritas operates in more than 200 countries and territories worldwide.

Since 1929, we have been providing assistance to people experiencing personal crises in Lithuania. We are the largest social assistance network in the Vilnius region, with more than 100 employees and 1,000 volunteers providing support through 14 centers and 30 parishes.
Our primary mission is to provide needs-based assistance to those experiencing personal crises and to bring together people of goodwill.

Caritas does not divide people, ensuring respect for the dignity of those facing difficulties or suffering. We are close to people in crisis and choose to be with them.

We provide assistance to children, youth, families, the elderly, migrants and refugees, prisoners, those experiencing psychological and emotional crises, homelessness, and those suffering from poverty and social exclusion.`,
		SubscriberCount: 143,
		VideoCount:      90,
		Categories: []string{
			"INSTITUCIJA",
			"IVAIRUS",
			"PODCAST",
		},
	},
	{
		Id:        "UCDAEVVGCCr4Tgn0GJZ_1LvA",
		Title:     "Vaciuks",
		CustomUrl: "/@Vaciuks",
		Description: `
SUB 0 (2016-12-24)
SUB 100 (2017-05-07)
SUB 1000 (2018-01-08)
SUB 3000(2018-11-25)
SUB 5000(2019-08-18)
Sugrįžimas (2023-09-19)

Vaclovas.poc@gmail.com`,
		SubscriberCount: 6810,
		VideoCount:      90,
		Categories: []string{
			"IVAIRUS",
			"VLOG",
		},
	},
	{
		Id:              "UCQXYaIOo5Ml4bk9C5MKD_sg",
		Title:           "valdelio",
		CustomUrl:       "/@valdelio420",
		Description:     `Kuriu video ir tiek!`,
		SubscriberCount: 11500,
		VideoCount:      153,
		Categories: []string{
			"ZAIDIMAI",
			"IVAIRUS",
			"VLOG",
			"GTA",
			"COUNTER-STRIKE",
		},
	},
	{
		Id:              "UCFbArsn6pmZGxv2A2skKa5w",
		Title:           "vanilla killa b",
		CustomUrl:       "/@vanillakillab",
		Description:     ``,
		SubscriberCount: 23300,
		VideoCount:      44,
		Categories: []string{
			"IVAIRUS",
			"HUMORAS",
		},
	},
	{
		Id:              "UCMs400z7tu_WeRkH7DfUFkA",
		Title:           "VAROM! Party",
		CustomUrl:       "/@varomparty2819",
		Description:     ``,
		SubscriberCount: 136,
		VideoCount:      99,
		Categories: []string{
			"IVAIRUS",
			"VLOG",
			"MUZIKA",
		},
	},
	{
		Id:        "UCPzzSNeZoxjJy9frS2FK3Sg",
		Title:     "Vestro",
		CustomUrl: "/@VestroShow",
		Description: `Jei man įdomu - bus video. Jei tau įdomu - bus daugiau.

👇 Turite temą, kurią reikėtų aptarti?
👇 Dėl reklamos?
👉vestroshow@gmail.com
`,
		SubscriberCount: 25400,
		VideoCount:      10,
		Categories: []string{
			"DOKUMENTIKA",
			"ESE",
		},
	},
	{
		Id:        "UCvOAJ_8-W5mzE6KK25TBWoA",
		Title:     "Vestuviu filmavimas - VIZA studio",
		CustomUrl: "/@vestuviufilmavimas-vizastu4034",
		Description: `VIZA studio - oficialus Lietuvos vestuvių apdovanojimų filmuotojas.
Nuotakos, sako, kad mūsų filmai tarsi iš Holivudo.
Žiūrint juos - šiurpuliukai eina per kūną.
www.vizastudio.lt
Mes kuriame video istorijas, kurias norisi prisiminti.`,
		SubscriberCount: 3010,
		VideoCount:      119,
		Categories: []string{
			"EDUKACIJA",
			"VLOG",
			"PREKINIS ZENKLAS",
		},
	},
	{
		Id:        "UCFahEHPubHircXz4OOlleqA",
		Title:     "ViktisLT",
		CustomUrl: "/@Viktislt",
		Description: `Esu įvairių žaidimų video kūrėjas - Youtube veteranas. Mano kanale išvysit įvairiausio žanro turinį. Nuo Minecraft iki CS:GO ar League of Legends video gidų. Žaidimai išrenkami fanų.

***Koks mano vardas?***
- Viktoras
***Kiek man metų?***
- 29
***Kuo užsiimu gyvenime?***
- Dievo keliai nežinomi, šiuo metu įstojau toliau mokytis :D
***Iš kur kilęs?***
- Klaipėda.
***Su kuo filmuoju?***
- Didžioji dauguma žaidimų filmuoti su Dxtory/OBS.
***Su kuo editinu video?***
 - Sony Vegas Pro 10/ Adobe After Effects CS5/CS6
***Kada bus naujausia serija?***
 - Sek mano FB/YT pasistengsiu pranešti kai pats apie tai žinosiu.
***Ar norėčiau kartu sužaisti Lets Play?***
 - Turėjau tokių užklausų begales, sukūrus serverius su fanais - visada prašom.
***Ar galiu kreiptis pagalbos?***
- Mano prigimtis padėti VISADA. Jeigu esi naujokas, ko nors nesupranti - paprašyk, kad sukurčiau video pamoką :)`,
		SubscriberCount: 11700,
		VideoCount:      292,
		Categories: []string{
			"ZAIDIMAI",
			"MINECRAFT",
			"LEAGUE OF LEGENDS",
		},
	},
	{
		Id:              "UCW4UCKB80NlgnD7SBJZMAlQ",
		Title:           "VILNIUS TECH",
		CustomUrl:       "/@VILNIUSTECH",
		Description:     `Oficialus Vilniaus Gedimino technikos universiteto - VILNIUS TECH bendruomenės video kanalas.`,
		SubscriberCount: 1570,
		VideoCount:      477,
		Categories: []string{
			"IVAIRUS",
			"ZAIDIMAI",
			"INSTITUCIJA",
			"PODCAST",
		},
	},
	{
		Id:              "UC33m79Zn4lFfIgk4gm-ZzVg",
		Title:           "Zaidziam",
		CustomUrl:       "/@Zaidziam",
		Description:     `Raganėjam žaidimuose.`,
		SubscriberCount: 6250,
		VideoCount:      225,
		Categories: []string{
			"ZAIDIMAI",
			"COUNTER-STRIKE",
			"LEAGUE OF LEGENDS",
		},
	},
	{
		Id:        "UCZikG4ExMa6X4B3sSjajIGw",
		Title:     "Žaismo DNR",
		CustomUrl: "/@ZaismoDNR",
		Description: `Internetinė TV apie žaidimus ir žaidimų kultūrą "Žaismo DNR".

Dėl laimėtų prizų kreipkitės žinute mūsų facebook paskyroje. Jei apie laimėjimą nepranešite  FB žinute per 3 savaites nuo paskelbimo, pasiliekame teisę prizo neatiduoti.`,
		SubscriberCount: 14700,
		VideoCount:      1100,
		Categories: []string{
			"ZAIDIMAI",
			"RECENZIJOS",
			"ZINIOS",
		},
	},
	{
		Id:              "UCrj1lKDSh_bxfqv2LuszXqw",
		Title:           "Žaliasis Ekranas",
		CustomUrl:       "/@Zaliasis_Ekranas",
		Description:     ``,
		SubscriberCount: 22,
		VideoCount:      25,
		Categories: []string{
			"ZAIDIMAI",
			"LEAGUE OF LEGENDS",
		},
	},
	{
		Id:        "UC3VYJScZFP-abuT_Dc1vYng",
		Title:     "Žaidimų Balsas",
		CustomUrl: "/@zbalsas",
		Description: `Žaidimų naujienos iš viso pasaulio. Tikrų "Geimerių" kanalas.
Turite kokių pasiūlymų Žaidimų Balsui - rašykite mums adresu admin@zaidimubalsas.lt.
Norite padėkoti ir prisidėti prie tolimesnio Žaidimų Balso augimo? Paremkite projektą - https://www.paypal.me/zaidimubalsas
AČIŪ, kad esate su Žaidimų Balsu ;)

`,
		SubscriberCount: 76800,
		VideoCount:      1900,
		Categories: []string{
			"RECENZIJOS",
			"PODCAST",
			"ZAIDIMAI",
		},
	},
	{
		Id:        "UC2c986H3lldYah8mV_tlEGg",
		Title:     "Žentelis",
		CustomUrl: "/@zenteliokanalas",
		Description: `https://www.youtube.com/@zenteliokanalas?sub_confirmation=1
OCPC.lt
Ocpclounge.lt
Kosminestovykla.lt
`,
		SubscriberCount: 9480,
		VideoCount:      53,
		Categories: []string{
			"RECENZIJOS",
		},
	},
	{
		Id:        "UCoAb2aR-ggdV82hRAxtBhww",
		Title:     "Žolininkė",
		CustomUrl: "/@zolininke",
		Description: `Kalbame apie anime ir manga ♥
O kartais žaidžiame žaidimus
`,
		SubscriberCount: 1580,
		VideoCount:      113,
		Categories: []string{
			"IVAIRUS",
			"ANIMACIJA",
		},
	},
	{
		Id:        "UCtEYjtoCAo03xiRyy0VsVIw",
		Title:     "Zygiz",
		CustomUrl: "/@Zygiz",
		Description: `😜 Žaidžiu, pasakoju ir reaguoju!

👋 Labukas! Šis kanalas pagrinde yra orientuotas į žaidimų turinį, bet pasitaiko įvairaus žanro vaizdo įrašų - istorijų pasakojimų, tierlistų, reakcijų ir dar visokių keistenybių!
💛 Stengiuosi kelti vaizdo įrašus bent kartą per savaitę!
💚 Prenumeruok kanalą ir prisijunk prie mano YouTubo kelionės!

🔴 Instagram・https://instagram.com/zygizyt
🟡 Discord serveris・https://discord.gg/bWjShuqsDK
⚪ Antrasis mano kanalas・https://youtube.com/@Zygiz2
⚫ TikTok・Keliu savo video ištraukas・https://tiktok.com/@zygizyt
🟤 TikTok・Keliu lietuviškus prikolus・https://tiktok.com/@zygizz

✉️ Business email・zygizyt@gmail.com

🔴 Prenumeratorių pasiekimus gali rasti šioje nuorodoje・https://pastebin.com/NYiNEqjh
`,
		SubscriberCount: 2029,
		VideoCount:      342,
		Categories: []string{
			"IVAIRUS",
			"ZAIDIMAI",
			"MINECRAFT",
		},
	},
	{
		Id:              "UCQXpN_Iqle7iBkmpHGuHvKA",
		Title:           "ZyyllMusic",
		CustomUrl:       "/@zyyllmusic9464",
		Description:     `Hello! We are ZyyllMusic, bringing you the best, new musics that is there. This channel is ownered by Serial Entrepreneur SEO Manager Zygimantas Mazeika.`,
		SubscriberCount: 17,
		VideoCount:      6,
		Categories: []string{
			"MUZIKA",
		},
	},
	{
		Id:              "UC-VCxhn6a-nnSxjgAQrLbcA",
		Title:           "Labandėliai",
		CustomUrl:       "/c/Laband%C4%97liai",
		Description:     `Pavalgom?`,
		SubscriberCount: 824,
		VideoCount:      40,
		Categories: []string{
			"RECENZIJOS",
			"KULINARIJA",
		},
	},
	{
		Id:              "UC1cAUvq1UlNN_VR83rm9M7w",
		Title:           "Lintrovert Lietuviškai",
		CustomUrl:       "/@LintrovertLietuviskai",
		Description:     `Sveiki! Aš esu dr. Linas Jonušauskas ir sveikinu atvykus į savo kanalą! Čia pasiūlysiu Jums savo politinius komentarus apie Lietuvos ir pasaulio aktualijas. Visas šio kanalo turinys bus išimtinai lietuvių kalba.`,
		SubscriberCount: 3360,
		VideoCount:      284,
		Categories: []string{
			"EDUKACIJA",
			"DOKUMENTIKA",
		},
	},
	{
		Id:        "UC962MKsp3HK-ker1plMA9NQ",
		Title:     "Pralaužk Vieną Šaltą",
		CustomUrl: "/c/Pralau%C5%BEkVien%C4%85%C5%A0alt%C4%85",
		Description: `Pralaužk Vieną Šaltą - jau ketvirtus metus skaičiuojantis ir sustoti nesiruošiantis podkastas. Kalbėti vis dar neišmokom, bet jau ir nebežadam, kad išmoksim. Bet prižadėti, kad ir toliau išvysite daugybę įdomių, mažiau ar daugiau žinomų pašnekovų, mes galime. Neįpareigojantys pokalbiai, atskleistos paslaptys, daug nejuokingų ir juokingų bajerių visada lauks jūsų mūsų laidose!
`,
		SubscriberCount: 65700,
		VideoCount:      377,
		Categories: []string{
			"PODCAST",
		},
	},
	{
		Id:              "UCAolQGvDZj5HWiAS4q8O1Qw",
		Title:           "VvVėjai",
		CustomUrl:       "/c/VyraujantysVakar%C5%B3V%C4%97jai",
		Description:     `Klarko tikras vardas yra "Balius"`,
		SubscriberCount: 7260,
		VideoCount:      72,
		Categories: []string{
			"IVAIRUS",
			"ZAIDIMAI",
		},
	},
}
var Categories = []string{
	"PODCAST",
	"IVAIRUS",
	"ZAIDIMAI",
	"ANIMACIJA",
	"PREKINIS ZENKLAS",
	"VLOG",
	"HUMORAS",
	"RECENZIJOS",
	"MINECRAFT",
	"EDUKACIJA",
	"MUZIKA",
	"MASINOS",
	"DOKUMENTIKA",
	"FORTNITE",
	"GTA",
	"COUNTER-STRIKE",
	"INSTITUCIJA",
	"ESE",
	"LEAGUE OF LEGENDS",
	"ZINIOS",
	"KULINARIJA",
}

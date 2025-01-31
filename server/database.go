package main

type Channel struct {
	Id              string
	Title           string
	CustomUrl       string
	Thumbnail       string
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
		Thumbnail:       "https://yt3.googleusercontent.com/C4tXv5SuSvKRvq9FO19zfmD317xy9TFh71lNRicHLfkbxVznajx7cgcZL5Kmmb7aEb-KhErDOw=s900-c-k-c0x00ffffff-no-rj",
		Description:     `"NEJUOKINGA" Pokalbinė yra laida, kuri suburia du žmones atviram bei nuoširdžiam pokalbiui - kaip namuose.`,
		SubscriberCount: 54500,
		VideoCount:      181,
		Categories: []string{
			"PODCAST",
		},
	},
	{
		Id:              "UCef7VJO5aZ80ZBsTy5cyrBA",
		Title:           "Overcrow",
		CustomUrl:       "/@Overcrow",
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_mYmLymFurZ0K8ugau2lBeWNJOUmAuNDNLumd1jFkegqQ4=s900-c-k-c0x00ffffff-no-rj",
		Description:     `overcrowding with degeneracy`,
		SubscriberCount: 1930,
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
		Thumbnail:       "https://yt3.googleusercontent.com/-8f0WPNDaZA7wGcrbYGvKQQ5qSbSShbfFZ_6VVWqZ8gIlF96Fuh77CHvbLkLSE5WCdly_QA=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/hDc0a8xJrRBXWjk-lkn8XJFPWKQn9ZrSVBf_qW6-xUJnc40ernxfsYcIYLvz1CGtJpX1jxTh=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/ytc/AIdro_mAxb5pG8TqHVkFpfvxd7H1L2iiw18yCuOCLirB5gG6Yig=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_mCVxuE378GsjadkONmmIUcmqHqQtFO22tSG9BaAixhmA=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/yhObpWLQ1HRgR_AzB3kClMcU_8DQG8v5P81CK14Fnq8J46FVcbOojF16u0MHde8gAgERNsov=s900-c-k-c0x00ffffff-no-rj",
		Description:     `ponasakiniuotis.business@gmail.com`,
		SubscriberCount: 48300,
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
		Thumbnail: "https://yt3.googleusercontent.com/aSHh6H2mzcpyM6lFA9X-M1dCIcyBcC7mRwIAgrabqQx9fckkL9lJGcEeLTaLNdUyoGxrMidK=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/6lYzEC0OwRzBgrCPV7dck00IiJljqhTrbSvzJe7hJMcnbZPLaOyveQBvJ-s6cvCsoKehjy_eBeY=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/2RfclUOCKenqNTsOU7SIl54dcbLJllefgl8VGHO5h60OPtjGD86yn5U2UEGxbu_XAs0HWbNj=s900-c-k-c0x00ffffff-no-rj",
		Description: `Filmuoju žaidimukus
`,
		SubscriberCount: 33200,
		VideoCount:      795,
		Categories: []string{
			"ZAIDIMAI",
		},
	},
	{
		Id:              "UCmEKDSs3dB3rRaIaVEEGmyA",
		Title:           "Valentinas",
		CustomUrl:       "/@PressButtonTVLT",
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_nH53Fnb4oVaoD0dlNsYfnEFw3_sY37qqS_OQq7Ir9RpCY=s900-c-k-c0x00ffffff-no-rj",
		Description:     `Sveiki. Aš esu Valentinas ir domiuosi video žaidimais bei kinu. Mano kanale pamatysite vaizdo įrašus, kuriuose dalinuose savo nuomone apie šias meno formas.`,
		SubscriberCount: 7270,
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
		Thumbnail:       "https://yt3.googleusercontent.com/6My7e-IBOhtgbSTnAViioHU8XnOm-dsjUMA_8RGsPyNpykFlg5_iBTp3UORaXejwbyCCr9iESg=s900-c-k-c0x00ffffff-no-rj",
		Description:     ``,
		SubscriberCount: 23300,
		VideoCount:      48,
		Categories: []string{
			"MUZIKA",
		},
	},
	{
		Id:              "UCe22CG3JAXvpP19aYO71bGg",
		Title:           "Prologai",
		CustomUrl:       "/@prologai3443",
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_mO6ym3ia6qxXVaST2PgbqDSrFQP459sGeXO2LrweW7Rw=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_nfF-YTClMdu7u2hAn3tCJ6i3G_vfbbFC1tspPiDXHttQ=s900-c-k-c0x00ffffff-no-rj",
		Description:     ``,
		SubscriberCount: 8010,
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
		Thumbnail: "https://yt3.googleusercontent.com/ytc/AIdro_mEPdNJStenF4fwQT1qCJtyPePLltpw2Nn977LmmVJFHRI=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/kk2vjdOL_ilU065jL1aBi1py9cGvoFX2GNiYlxibvQNjOuJfDmmGdwiHSUd_6dv10PnxfEHC=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/kS457YIIzSG3i-Y9TPOhl2Z4MjX66Wf8OXdXxAjZ8PaWqV5xIdoRhOS4rN27wVpGSKBY43NyVO0=s900-c-k-c0x00ffffff-no-rj",
		Description: `Labas, mes - QBuilds.

Kiekvienas mūsų kompiuteris surinktas su dideliu dėmesiu detalei ir kokybei. Ieškome geriausių sprendimų sukurti balansui tarp kokybės, funkcionalumo ir vizualų.

Mūsų kanale rasi geriausių rinkoje esančių kompiuterių apžvalgas, naujienas tech pasaulyje ir clean setup’us. 

`,
		SubscriberCount: 3320,
		VideoCount:      210,
		Categories: []string{
			"RECENZIJOS",
			"IVAIRUS",
		},
	},
	{
		Id:              "UCU8AueUTR9QAmyqyvgcTBYg",
		Title:           "Radijo Ranča",
		CustomUrl:       "/@RadijoRanca",
		Thumbnail:       "https://yt3.googleusercontent.com/8RhtUm5F3CS_dLAakHW7IZxH4mjKjtJt9feks_XJ-NYbX7slu33Q5LEE8fknWYQuS5DKrCKwbA=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/ytc/AIdro_leHqqCPXx-0SjRUjxGxcfEHbOAv0aEWnEJHizmXOhGL1U=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_lLVIg6saqreDs0omkJKd1WpADO0OlYz3W1jyJMRSN7YpE=s900-c-k-c0x00ffffff-no-rj",
		Description:     `Padėkite man kurti šį kanalą. Ačiū!`,
		SubscriberCount: 188000,
		VideoCount:      207,
		Categories: []string{
			"EDUKACIJA",
			"DOKUMENTIKA",
		},
	},
	{
		Id:        "UC1uXcCFBq4c1L1je3-SuWnQ",
		Title:     "Titukis",
		CustomUrl: "/@Titukis",
		Thumbnail: "https://yt3.googleusercontent.com/07qEFASWnlPn7KxuZY_xqnb3AdErwukHpSog8w9FmELLcHKvmCf-47cRSTuiLegCTrKWRQdgcGA=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/1GZ5utZG7oiNS2BrgNqJKKoUgsUbY8CRtgj0k3E5jydw5nntNQt3O2HDxDejrxAs1Zq2CL607w=s900-c-k-c0x00ffffff-no-rj",
		Description: `⬇️Instagram : tomuxas__
`,
		SubscriberCount: 36700,
		VideoCount:      377,
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
		Thumbnail: "https://yt3.googleusercontent.com/t_oqxgRsKqF3YdyBPfGADgkezEtGcQ7zh6Ae_UTvDl4tvQMb1tIez2mLLHSyBWKT2BDyky0xCw=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/P4TL3uuPS6rSZB7mBx6_TSEVfuJMc_PjWDpGZ7xQMkZQbypKCby7-K0kaLocIL9lvaxvyDono3s=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/ytc/AIdro_nItOVGIlNapPJMQKnh9N4itRMAX3RrDq2-kbIEDdX_Oy8=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/Q2d3GH2PJ60GZuJjBKd51koqvaj090FoKTkPtgSZcZQShLzcdztZ1HMryOQ4zS8VmxSD-spu=s900-c-k-c0x00ffffff-no-rj",
		Description: `Sveiki visi. 
Mėgstu siuntinius, ypač nemokamus. Siunčiuosi, išpakuoju ir žiūriu, ką gavau ir dalinuosi video su Jumis.
Kam patiko, prenumeruokite, spauskite LIKE. Kam nepatiko irgi spauskite LIKE.
Nesu aš profesionalus apžvalgininkas, tad prašom neįsižeisti, jei ko neparodau kanale, ar kažko nepaaiškinu.
Nesu tarpininkas ir prekių neužsakinėju, komercija neužsiimu:)
Visos prekės, kurias rodau kanale, tai mano arba draugų užsakytos.

Unboxingltu@gmail.com`,
		SubscriberCount: 4840,
		VideoCount:      621,
		Categories: []string{
			"RECENZIJOS",
		},
	},
	{
		Id:              "UCBLvjJODX0mi5IfKAiq_JBA",
		Title:           "Undergroundinis Knygynas",
		CustomUrl:       "/@undergroundinisknygynas",
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_nV7z3v6NZDpIQQyjpdMQnwQIHl0MXbdlGgg3V5VYGanQ=s900-c-k-c0x00ffffff-no-rj",
		Description:     ``,
		SubscriberCount: 4110,
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
		Thumbnail: "https://yt3.googleusercontent.com/ytc/AIdro_lCIEhtQchCKJvQma1go6OT0GgUkK3t0OxHKHMG5PcuW2A=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/QTvq1pl2cu8RhYGhCT8qeyDVppO9AoXi8bOSG3fStPFFX9pDVg1M8d332Ufs7emwnIYWCNVjByQ=s900-c-k-c0x00ffffff-no-rj",
		Description: `Vilnius Archdiocese Caritas is part of one of the world’s largest aid networks. Caritas operates in more than 200 countries and territories worldwide.

Since 1929, we have been providing assistance to people experiencing personal crises in Lithuania. We are the largest social assistance network in the Vilnius region, with more than 100 employees and 1,000 volunteers providing support through 14 centers and 30 parishes.
Our primary mission is to provide needs-based assistance to those experiencing personal crises and to bring together people of goodwill.

Caritas does not divide people, ensuring respect for the dignity of those facing difficulties or suffering. We are close to people in crisis and choose to be with them.

We provide assistance to children, youth, families, the elderly, migrants and refugees, prisoners, those experiencing psychological and emotional crises, homelessness, and those suffering from poverty and social exclusion.`,
		SubscriberCount: 144,
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
		Thumbnail: "https://yt3.googleusercontent.com/YmbboEYkCcvWNvHRjNNIrv8nVIzQyPT2wQjrSBQk4s0JYEzyOcGTsUAOc_iOiiM1Ms9plk1vCg=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_n0ePFxVXvKeMHduPtuGtdbcGnUA3djXbXYghxDruyUgw=s900-c-k-c0x00ffffff-no-rj",
		Description:     `Kuriu video ir tiek!`,
		SubscriberCount: 11500,
		VideoCount:      156,
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_lmleROtlMSaeHAo9eiaiWUeVEwXquV_TZkXlPfbuOVpFA=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_ngrCgWPYyefuJYdwBRzAB0LLyMdRrbWiao-lOXtXSzXUY=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/2uFu4smIUWWgtqVXI5rOMVksoGh_2ZA1-7kn3ljnphwFfFQQpzSu8icrx_pENg8dpx5aQYEUzn0=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/ytc/AIdro_k3hQdgeqQwqAVBj3Wnm09B7vzdc5EE0qSS4ZRgduG5BA=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/ytc/AIdro_loNtIXxOLxVbNxsigc_EbgqriVHk5fvoBAme3qHCuNCA=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_lhz-xsTo_FkX-UF1Mc8dRV9kpsP3hkxVu5x1DfyI_fzcQ=s900-c-k-c0x00ffffff-no-rj",
		Description:     `Oficialus Vilniaus Gedimino technikos universiteto - VILNIUS TECH bendruomenės video kanalas.`,
		SubscriberCount: 1580,
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_mzuFVrK3Vk1rGL_L9X4S-J5Rqgdk3Iae0OqLmjWE0joA=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/ytc/AIdro_l6y4GPmwp8klVMscTpHPUvZ4DdMQYOav5Fg-VvWoV5eQ=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/7GFA4VTeiS1EsYWDMZH7P5TEJYq_uWL4wW_Tp0dxIqsSTZKaLOPoC0_SuRxHhATenRoW8TKNlw=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/v4Rk9TznJurufc0IKgEpiJkcu6EVJU62qE1GzRBddN7OE96O9fhjD04NjpRMqdjO874VMUge_5w=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail: "https://yt3.googleusercontent.com/GeWBFAHFhh3MRGjmKh_stdEJwM345CvALQ7DYwqEVBApkMp2GlG-jwI44hIIV1o7VqLL_ZrR=s900-c-k-c0x00ffffff-no-rj",
		Description: `https://www.youtube.com/@zenteliokanalas?sub_confirmation=1
OCPC.lt
Ocpclounge.lt
Kosminestovykla.lt
`,
		SubscriberCount: 9630,
		VideoCount:      54,
		Categories: []string{
			"RECENZIJOS",
		},
	},
	{
		Id:        "UCoAb2aR-ggdV82hRAxtBhww",
		Title:     "Žolininkė",
		CustomUrl: "/@zolininke",
		Thumbnail: "https://yt3.googleusercontent.com/DjYzZYfPW7zzgmSCvp7GUzoqLTy2AyAXMfPQ0JCY_luGvHvcjoftSdl9TZJQxRblWZXoWVdIDWg=s900-c-k-c0x00ffffff-no-rj",
		Description: `Kalbame apie anime ir manga ♥
O kartais žaidžiame žaidimus
`,
		SubscriberCount: 1590,
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
		Thumbnail: "https://yt3.googleusercontent.com/HR7U3s_BrVhJGjsAaduulKpFohfWXbSvQKxUQDIqQaVfI1zu_vbv9zENPOFEOm5zPkqnGt46OgQ=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/m2q1hB6fsAlhSautzU3ttE_9N4Jha1CP1nQVEEyCy8kkv0n0uex0huREhJkuLamgGTvKZnqIUYw=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_nGzI0BLjAUfPjv8f2TxuTkgDXgJcsP4o17yRT8s562sg=s900-c-k-c0x00ffffff-no-rj",
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
		Thumbnail:       "https://yt3.googleusercontent.com/huG1AlqOg-p0VtNCr7qWloH1AwfyJr5vjjTPJ7sndI71qM4EZgM3V3qfmjWtuaI_54mfymGUAg=s900-c-k-c0x00ffffff-no-rj",
		Description:     `Sveiki! Aš esu dr. Linas Jonušauskas ir sveikinu atvykus į savo kanalą! Čia pasiūlysiu Jums savo politinius komentarus apie Lietuvos ir pasaulio aktualijas. Visas šio kanalo turinys bus išimtinai lietuvių kalba.`,
		SubscriberCount: 3410,
		VideoCount:      287,
		Categories: []string{
			"EDUKACIJA",
			"DOKUMENTIKA",
		},
	},
	{
		Id:        "UC962MKsp3HK-ker1plMA9NQ",
		Title:     "Pralaužk Vieną Šaltą",
		CustomUrl: "/c/Pralau%C5%BEkVien%C4%85%C5%A0alt%C4%85",
		Thumbnail: "https://yt3.googleusercontent.com/HOmKtkPSxCosgCeOg5HVNoEedh1E_WQfhuhaHgUm6YNQCSJh9-4Vltx-2qXphbhpl5tONYTg=s900-c-k-c0x00ffffff-no-rj",
		Description: `Pralaužk Vieną Šaltą - jau ketvirtus metus skaičiuojantis ir sustoti nesiruošiantis podkastas. Kalbėti vis dar neišmokom, bet jau ir nebežadam, kad išmoksim. Bet prižadėti, kad ir toliau išvysite daugybę įdomių, mažiau ar daugiau žinomų pašnekovų, mes galime. Neįpareigojantys pokalbiai, atskleistos paslaptys, daug nejuokingų ir juokingų bajerių visada lauks jūsų mūsų laidose! 
`,
		SubscriberCount: 65900,
		VideoCount:      377,
		Categories: []string{
			"PODCAST",
		},
	},
	{
		Id:              "UCAolQGvDZj5HWiAS4q8O1Qw",
		Title:           "VvVėjai",
		CustomUrl:       "/c/VyraujantysVakar%C5%B3V%C4%97jai",
		Thumbnail:       "https://yt3.googleusercontent.com/ytc/AIdro_lU9em4LnXGey7XG7uEpV35sbXL9mZIT_r7ORa7mhkmc6s=s900-c-k-c0x00ffffff-no-rj",
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

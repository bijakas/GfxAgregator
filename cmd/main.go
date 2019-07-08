package main

import (
    "fmt"
    "github.com/bijakas/GenflixAgregator/pkg/funct"
)

func main() {

    jsonRes := `{
        "status": 1,
        "total_data": 29,
        "data": [
            {
                "video_type": "series",
                "content_id": "368",
                "title": "Shizuku-chan Eps 01",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Shizuku chan",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___368___/Cover_Episode_1.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___368___/Cover_Episode_1.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/368"
            },
            {
                "video_type": "series",
                "content_id": "558",
                "title": "Fantastic Children Episode 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Fantastic Children",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___558___/cu_FantasticChildren.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___558___/cu_FantasticChildren.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/558"
            },
            {
                "video_type": "series",
                "content_id": "559",
                "title": "Before  Green Gables Episode 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Before Green Gables",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___559___/cu_BeforeGreenGables.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___559___/cu_BeforeGreenGables.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/559"
            },
            {
                "video_type": "series",
                "content_id": "602",
                "title": "Make Up Cosplay",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Sushiroll Make Up",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___602___/Sushiroll_Jeanice.Makeup_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___602___/Sushiroll_Jeanice.Makeup_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/602"
            },
            {
                "video_type": "series",
                "content_id": "573",
                "title": "Love Pledge  - Carlo",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Love Pledge",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___573___/c_LovePledge.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___573___/c_LovePledge.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/573"
            },
            {
                "video_type": "series",
                "content_id": "569",
                "title": "Top 10 Quotes Anime",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Top Popular Quotes",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/569"
            },
            {
                "video_type": "series",
                "content_id": "568",
                "title": "Top 10 Karakter Anime  ",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Top Popular Karakter Anime",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/568"
            },
            {
                "video_type": "series",
                "content_id": "615",
                "title": "Hane Bado! Episode 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Hanebado!",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___615___/cu_Hanebado.1_.300_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___615___/cu_Hanebado.1_.300_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/615"
            },
            {
                "video_type": "series",
                "content_id": "617",
                "title": "Fight League: Gear Gadget Generators Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Fight League",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___617___/cu_FightLeague.Eps_.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___617___/cu_FightLeague.Eps_.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/617"
            },
            {
                "video_type": "series",
                "content_id": "621",
                "title": "Panda Fanfare Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Panda Fanfare",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___621___/cu_PandaFanfare.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___621___/cu_PandaFanfare.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/621"
            },
            {
                "video_type": "series",
                "content_id": "623",
                "title": "You Dont Know Gunma Yet! Episode 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "You dont know gunma yet",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___623___/cu_YouDontKnowGunmaYet.300_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___623___/cu_YouDontKnowGunmaYet.300_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/623"
            },
            {
                "video_type": "series",
                "content_id": "627",
                "title": "Bananya Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Bananya",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___627___/cu_sr_Bananya_eps_1.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___627___/cu_sr_Bananya_eps_1.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/627"
            },
            {
                "video_type": "series",
                "content_id": "658",
                "title": "Canaan Ep 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Canaan",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___658___/cu_Canaan.Eps_.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___658___/cu_Canaan.Eps_.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/658"
            },
            {
                "video_type": "series",
                "content_id": "664",
                "title": "Girls Und Panzer Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Girls Und Panzer",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___664___/cu_GirlsUndPanzer.Eps_.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___664___/cu_GirlsUndPanzer.Eps_.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/664"
            },
            {
                "video_type": "series",
                "content_id": "660",
                "title": "Merc Storia Ep 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Merc Storia",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___660___/merc_storia_ep_1.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___660___/merc_storia_ep_1.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/660"
            },
            {
                "video_type": "series",
                "content_id": "662",
                "title": "Ms. vampire who lives in my neighborhood Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Ms. vampire who lives in my neighborhood.",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___662___/vampire_eps_1.jpg.png",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___662___/vampire_eps_1.jpg.png"
                },
                "player": "http://sushiroll.co.id/smartfren/play/662"
            },
            {
                "video_type": "series",
                "content_id": "682",
                "title": "Tari Tari Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Tari Tari",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___682___/WhatsApp_Image_2019-05-31_at_13.55_.37_.jpeg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___682___/WhatsApp_Image_2019-05-31_at_13.55_.37_.jpeg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/682"
            },
            {
                "video_type": "series",
                "content_id": "694",
                "title": "Meiji Tokyo Renka Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Meiji Tokyo Renka",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___694___/cu_MeijiTokyoRenka.Eps_.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___694___/cu_MeijiTokyoRenka.Eps_.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/694"
            },
            {
                "video_type": "series",
                "content_id": "748",
                "title": "Mao sama. Retry Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": "Mao sama. Retry!",
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___748___/cu_MousamaRetry.Eps_.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___748___/cu_MousamaRetry.Eps_.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/748"
            },
            {
                "video_type": "movies",
                "content_id": "72",
                "title": "Million Doll Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___72___/1-Million-doll-300x400.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___72___/1-Million-doll-300x400.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/72"
            },
            {
                "video_type": "movies",
                "content_id": "83",
                "title": "Orenchi no Furo Jijou Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___83___/ORENCHI-NO-FURO-JIJO-1-300x400.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___83___/ORENCHI-NO-FURO-JIJO-1-300x400.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/83"
            },
            {
                "video_type": "movies",
                "content_id": "96",
                "title": "Hakone-chan Eps 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___96___/1-hakone-chan-300x400.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___96___/1-hakone-chan-300x400.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/96"
            },
            {
                "video_type": "movies",
                "content_id": "368",
                "title": "Shizuku-chan Eps 01",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___368___/Cover_Episode_1.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___368___/Cover_Episode_1.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/368"
            },
            {
                "video_type": "movies",
                "content_id": "558",
                "title": "Fantastic Children Episode 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___558___/cu_FantasticChildren.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___558___/cu_FantasticChildren.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/558"
            },
            {
                "video_type": "movies",
                "content_id": "559",
                "title": "Before  Green Gables Episode 1",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___559___/cu_BeforeGreenGables.1_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___559___/cu_BeforeGreenGables.1_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/559"
            },
            {
                "video_type": "movies",
                "content_id": "568",
                "title": "Top 10 Karakter Anime  ",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/568"
            },
            {
                "video_type": "movies",
                "content_id": "569",
                "title": "Top 10 Quotes Anime",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/569"
            },
            {
                "video_type": "movies",
                "content_id": "578",
                "title": "Silver Noah-Love Pledge",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___578___/c_LovePledge.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___578___/c_LovePledge.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/578"
            },
            {
                "video_type": "movies",
                "content_id": "602",
                "title": "Make Up Cosplay",
                "synopsis": "",
                "artist": "",
                "genre": [
                    "Anime"
                ],
                "package_name": null,
                "poster": {
                    "s": "http://sushiroll.co.id/assets/_preview/vid/___602___/Sushiroll_Jeanice.Makeup_.jpg",
                    "m": "http://sushiroll.co.id/assets/_preview/vid/___602___/Sushiroll_Jeanice.Makeup_.jpg"
                },
                "player": "http://sushiroll.co.id/smartfren/play/602"
            }
        ]
    }`
        fmt.Println("test..")
       // funct.MakeVideoBulkMutationString (jsonRes)
        funct.MakeGraphQlRequest(funct.QueryToRequest(( funct.MakeVideoBulkMutationString (jsonRes)  )))
    
        
jsonResSushiRoll := `{
    "status": 1,
    "total_data": 29,
    "data": [
        {
            "video_type": "series",
            "content_id": "368",
            "title": "Shizuku-chan Eps 01",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Shizuku chan",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___368___/Cover_Episode_1.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___368___/Cover_Episode_1.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/368"
        },
        {
            "video_type": "series",
            "content_id": "558",
            "title": "Fantastic Children Episode 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Fantastic Children",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___558___/cu_FantasticChildren.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___558___/cu_FantasticChildren.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/558"
        },
        {
            "video_type": "series",
            "content_id": "559",
            "title": "Before  Green Gables Episode 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Before Green Gables",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___559___/cu_BeforeGreenGables.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___559___/cu_BeforeGreenGables.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/559"
        },
        {
            "video_type": "series",
            "content_id": "602",
            "title": "Make Up Cosplay",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Sushiroll Make Up",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___602___/Sushiroll_Jeanice.Makeup_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___602___/Sushiroll_Jeanice.Makeup_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/602"
        },
        {
            "video_type": "series",
            "content_id": "573",
            "title": "Love Pledge  - Carlo",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Love Pledge",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___573___/c_LovePledge.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___573___/c_LovePledge.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/573"
        },
        {
            "video_type": "series",
            "content_id": "569",
            "title": "Top 10 Quotes Anime",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Top Popular Quotes",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/569"
        },
        {
            "video_type": "series",
            "content_id": "568",
            "title": "Top 10 Karakter Anime  ",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Top Popular Karakter Anime",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/568"
        },
        {
            "video_type": "series",
            "content_id": "615",
            "title": "Hane Bado! Episode 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Hanebado!",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___615___/cu_Hanebado.1_.300_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___615___/cu_Hanebado.1_.300_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/615"
        },
        {
            "video_type": "series",
            "content_id": "617",
            "title": "Fight League: Gear Gadget Generators Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Fight League",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___617___/cu_FightLeague.Eps_.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___617___/cu_FightLeague.Eps_.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/617"
        },
        {
            "video_type": "series",
            "content_id": "621",
            "title": "Panda Fanfare Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Panda Fanfare",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___621___/cu_PandaFanfare.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___621___/cu_PandaFanfare.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/621"
        },
        {
            "video_type": "series",
            "content_id": "623",
            "title": "You Dont Know Gunma Yet! Episode 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "You dont know gunma yet",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___623___/cu_YouDontKnowGunmaYet.300_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___623___/cu_YouDontKnowGunmaYet.300_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/623"
        },
        {
            "video_type": "series",
            "content_id": "627",
            "title": "Bananya Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Bananya",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___627___/cu_sr_Bananya_eps_1.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___627___/cu_sr_Bananya_eps_1.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/627"
        },
        {
            "video_type": "series",
            "content_id": "658",
            "title": "Canaan Ep 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Canaan",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___658___/cu_Canaan.Eps_.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___658___/cu_Canaan.Eps_.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/658"
        },
        {
            "video_type": "series",
            "content_id": "664",
            "title": "Girls Und Panzer Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Girls Und Panzer",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___664___/cu_GirlsUndPanzer.Eps_.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___664___/cu_GirlsUndPanzer.Eps_.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/664"
        },
        {
            "video_type": "series",
            "content_id": "660",
            "title": "Merc Storia Ep 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Merc Storia",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___660___/merc_storia_ep_1.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___660___/merc_storia_ep_1.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/660"
        },
        {
            "video_type": "series",
            "content_id": "662",
            "title": "Ms. vampire who lives in my neighborhood Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Ms. vampire who lives in my neighborhood.",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___662___/vampire_eps_1.jpg.png",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___662___/vampire_eps_1.jpg.png"
            },
            "player": "http://sushiroll.co.id/smartfren/play/662"
        },
        {
            "video_type": "series",
            "content_id": "682",
            "title": "Tari Tari Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Tari Tari",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___682___/WhatsApp_Image_2019-05-31_at_13.55_.37_.jpeg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___682___/WhatsApp_Image_2019-05-31_at_13.55_.37_.jpeg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/682"
        },
        {
            "video_type": "series",
            "content_id": "694",
            "title": "Meiji Tokyo Renka Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Meiji Tokyo Renka",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___694___/cu_MeijiTokyoRenka.Eps_.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___694___/cu_MeijiTokyoRenka.Eps_.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/694"
        },
        {
            "video_type": "series",
            "content_id": "748",
            "title": "Mao sama. Retry Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": "Mao sama. Retry!",
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___748___/cu_MousamaRetry.Eps_.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___748___/cu_MousamaRetry.Eps_.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/748"
        },
        {
            "video_type": "movies",
            "content_id": "72",
            "title": "Million Doll Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___72___/1-Million-doll-300x400.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___72___/1-Million-doll-300x400.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/72"
        },
        {
            "video_type": "movies",
            "content_id": "83",
            "title": "Orenchi no Furo Jijou Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___83___/ORENCHI-NO-FURO-JIJO-1-300x400.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___83___/ORENCHI-NO-FURO-JIJO-1-300x400.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/83"
        },
        {
            "video_type": "movies",
            "content_id": "96",
            "title": "Hakone-chan Eps 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___96___/1-hakone-chan-300x400.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___96___/1-hakone-chan-300x400.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/96"
        },
        {
            "video_type": "movies",
            "content_id": "368",
            "title": "Shizuku-chan Eps 01",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___368___/Cover_Episode_1.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___368___/Cover_Episode_1.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/368"
        },
        {
            "video_type": "movies",
            "content_id": "558",
            "title": "Fantastic Children Episode 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___558___/cu_FantasticChildren.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___558___/cu_FantasticChildren.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/558"
        },
        {
            "video_type": "movies",
            "content_id": "559",
            "title": "Before  Green Gables Episode 1",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___559___/cu_BeforeGreenGables.1_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___559___/cu_BeforeGreenGables.1_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/559"
        },
        {
            "video_type": "movies",
            "content_id": "568",
            "title": "Top 10 Karakter Anime  ",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/568"
        },
        {
            "video_type": "movies",
            "content_id": "569",
            "title": "Top 10 Quotes Anime",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/569"
        },
        {
            "video_type": "movies",
            "content_id": "578",
            "title": "Silver Noah-Love Pledge",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___578___/c_LovePledge.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___578___/c_LovePledge.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/578"
        },
        {
            "video_type": "movies",
            "content_id": "602",
            "title": "Make Up Cosplay",
            "synopsis": "",
            "artist": "",
            "genre": [
                "Anime"
            ],
            "package_name": null,
            "poster": {
                "s": "http://sushiroll.co.id/assets/_preview/vid/___602___/Sushiroll_Jeanice.Makeup_.jpg",
                "m": "http://sushiroll.co.id/assets/_preview/vid/___602___/Sushiroll_Jeanice.Makeup_.jpg"
            },
            "player": "http://sushiroll.co.id/smartfren/play/602"
        }
    ]
}`

        funct.MakeGraphQlRequest(funct.QueryToRequest(( funct.MakeVideoBulkMutationString (jsonResSushiRoll)  )))
    }

package main

import (
    "strings"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "crypto/tls"
    "strconv"
   // "github.com/tidwall/gjson"

)

const (
    VERSION = 1
    PROVIDER_URL = "http://api-dev.genflix.co.id/basic/smartfren/content"
    GRAPHQL_URL = "http://10.16.5.159:8001/query"
)

type GenflixResp struct {
	Status        int    `json:"status"`
	Message       string `json:"message"`
	Next          bool   `json:"next"`
	RequestOffset int    `json:"request_offset"`
	RequestLimit  int    `json:"request_limit"`
	TotalData     int    `json:"total_data"`
	Data          []struct {
		VideoType string   `json:"video_type"`
		ContentID string   `json:"content_id"`
		Title     string   `json:"title"`
		Synopsis  string   `json:"synopsis"`
		Duration  string   `json:"duration"`
		Artist    string   `json:"artist"`
		Genre     []string `json:"genre"`
		Poster    struct {
			S string `json:"s"`
			M string `json:"m"`
			L string `json:"l"`
		} `json:"poster"`
		Player string `json:"player"`
	} `json:"data"`
}

type VideoBank struct {
    Provider string
    ProviderShort string
    Tags string
    ProviderLabel string
    Title string 
    TitlePackage string
    VideoType string
    Genre string
    Year string
    Duration string 
    ContentId string 
    Synopsis string 
    Cast string 
    PlayerUrl string 
    S string 
    M string 
    L string 
    Director string 
    ContentType string 
    IsActive string
}

func jsonEscape(i string) string {
    b,err :=json.Marshal(i)
    if err != nil {
        panic(err)
    }
    s:= string(b)
    return s[1:len(s)-1]
} 

func buildGraphQlMutationString(vb VideoBank) string{
    // videoType genre year duration contentType
    string := `
           {
            provider: "`+vb.Provider+`",
            providerShort: "`+vb.ProviderShort+`",
            tags: "`+vb.Tags+`",
            providerLabel: "`+vb.ProviderLabel+`",   
            title: "`+jsonEscape(vb.Title)+`",
            titlePackage: "`+jsonEscape(vb.TitlePackage)+`",
            videoType: "`+"MOVIES"+`",
            genre: ["HORROR", "FANTASY"], 
            year: 2019,
            duration: 0, 
            contentId: `+vb.ContentId+`, 
            synopsis: "`+jsonEscape(vb.Synopsis)+`",
            cast: "`+jsonEscape(vb.Cast)+`",
            playerUrl: "`+jsonEscape(vb.PlayerUrl)+`",
            poster: {
              s: "`+vb.S+`",
              m: "`+vb.M+`",
              l: "`+vb.L+`",
            },
            director: "`+vb.Director+`",
            contentType: "`+"FREE"+`",
            isActive: `+vb.IsActive+`,   
          
          },
      `
    return string
}


func main() {

    jsonRes := `{
        "status": 1,
        "message": "SUCCESS",
        "next": true,
        "request_offset": 1,
        "request_limit": 100,
        "total_data": 179,
        "data": [
            
            {
                "video_type": "movies",
                "content_id": "6",
                "title": "Hollow Man",
                "synopsis": "Hollow man mengisahkan tentang sebuah kelompok peneliti yang tengah melakukan penelitan sebuah serum yang mampu membuat objek yang disuntikannya bisa menghilang. Hingga kemudian ketika anti serum  formula tersebut tidak cocok dengan manusia. Sehingga membuat peneliti tidak bisa kembali lagi terlihat. Sekaligus menjadi awal dari segala masalah",
                "duration": "119",
                "artist": "Kevin Bacon , Elisabeth Shue",
                "genre": [
                    "Fiksi",
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/6/cu_HollowMan300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/6/cu_HollowMan400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/6/cu_HollowMan1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/hollow-man"
            },
            {
                "video_type": "movies",
                "content_id": "7",
                "title": "Bridget Jones Diary",
                "synopsis": "Selama ini Bridget kesulitan menemukan cintanya, sampai ia memutuskan menulis buku harian dan mendokumentasikan hidupnya. Tak diduga, datanglah dua pria yang tertarik untuk berkencan dengannya.",
                "duration": "97",
                "artist": "Renée Zellweger , Colin Firth",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/7/Bridget-jones-diary-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/7/bridget-jones-diary-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/7/Bridget-jones-1280x720-diary.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/bridget-jones-diary"
            },
            {
                "video_type": "movies",
                "content_id": "8",
                "title": "Unbreakable",
                "synopsis": "Seorang pria mempelajari sesuatu yang luar biasa dalam dirinya setelah mengalami kecelakan kereta api yang dahsyat.",
                "duration": "107",
                "artist": "Bruce Willis , Samuel L Jackson",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/8/Unbreakable-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/8/Unbreakable-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/8/Unbreakable-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/unbreakable"
            },
            {
                "video_type": "movies",
                "content_id": "9",
                "title": "Behind Enemy Lines",
                "synopsis": "Pesawat F-18 yang ditumpangi Chris Bunett dan pilotnya ditembak oleh pasukan separatis Bosnia yang kejam. Nyawa Bunet pun terancam karena bala tentara menganggapnya seorang musuh.",
                "duration": "111",
                "artist": "Gene Hackman , Owen Wilson",
                "genre": [
                    "Drama",
                    "Horor"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/9/behinde-enemy-lines-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/9/behind-enemy-lines-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/9/behind-enemy-lines-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/behind-enemy-lines"
            },
            {
                "video_type": "movies",
                "content_id": "10",
                "title": "Enemy at the Gates",
                "synopsis": "Peperangan 2 penembak jitu (sniper-warfare) pada Pertempuran Stalingrad tahun 1942, Vasily Zaitsev dari kubu Uni Soviet melawan rivalnya Mayor Heinz Thorvald dari Jerman.",
                "duration": "111",
                "artist": "Gene Hackman , Owen Wilson",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/10/enemy-at-the-gates300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/10/enemy-at-the-gates400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/10/enemy-at-the-gates-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/enemy-at-the-gates"
            },
            {
                "video_type": "movies",
                "content_id": "11",
                "title": "Far From Heaven",
                "synopsis": "Catchy dan keluarga hidup dengan bahagia bahkan nyaris sempurna. Namun semua berubah ketika dia mendapati suaminya memiliki hubungan gelap dengan wanita lain. Dalam kondisi tersebut, Cathy menjadi dekat dengan tukang kebun rumahnya hingga perasaan terlarang pun tumbuh di hati Cathy.",
                "duration": "81",
                "artist": "Julianne Moore , Dennis Quaid",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/11/Far-From-Heaven-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/11/Far-From-Heaven-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/11/Far-From-Heaven-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/far-from-heaven"
            },
            {
                "video_type": "movies",
                "content_id": "12",
                "title": "Legally Blonde",
                "synopsis": "Setelah putus dari kekasihnya karena dianggap berotak dangkal, Elle gadis berambut pirang memutuskan untuk masuk ke sekolah hukum Harvard. Disana dia berubah menjadi wanita yang cerdas.",
                "duration": "111",
                "artist": "Reese Witherspoon , Luke Wilson",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/12/legally-blonde-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/12/legally-blonde-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/12/legally-blonde-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/legally-blonde"
            },
            {
                "video_type": "movies",
                "content_id": "13",
                "title": "Men of Honor",
                "synopsis": "Sekelompok Angkatan Laut AS berkulit hitam direkrut untuk mengatasi rasisme dan menjadi penyelam kulit hitam pertama pasukan itu, bahkan sampai rela kehilangan kakinya.",
                "duration": "129",
                "artist": "Cuba Gooding Jr. , Robert De Niro",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/13/men-of-honor-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/13/men-of-honor-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/13/men-of-honor-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/men-of-honor"
            },
            {
                "video_type": "movies",
                "content_id": "14",
                "title": "The Pitch Black",
                "synopsis": "Pesawat yang menampung banyak penumpang terdampar di sebuah planet misterius yang ditinggali oleh makhluk mengerikan yang siap memburu mangsanya pada malam hari.",
                "duration": "107",
                "artist": "Jim Wheat , Radha Mitchell",
                "genre": [
                    "Horor"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/14/pitch-black-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/14/pitch-black-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/14/pitch-black-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-pitch-black"
            },
            {
                "video_type": "movies",
                "content_id": "15",
                "title": "Replicant",
                "synopsis": "Untuk menangkap tersangka pembunuhan berantai, para ilmuwan membuat tiruan genetika dari pembunuh tersebut. Bersama dengan kepolisian, mampukah mereka menangkap sang pembunuh?",
                "duration": "111",
                "artist": "Van Damme , Michael Rooker",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/15/replicant-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/15/replicant-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/15/replicant-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/replicant"
            },
            {
                "video_type": "movies",
                "content_id": "16",
                "title": "Snatch",
                "synopsis": "Promotor tinju, bos mafia judi, pencuri profesional, petinju, ex-KGB, pembunuh bayaran, dan penjual permata terjebak peristiwa tidak terduga yang melibatkan pencurian permata dan pertandingan tinju ilegal.",
                "duration": "81",
                "artist": "Jason Statham , Brad Pitt",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/16/snatch-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/16/snatch-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/16/snatch-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/snatch"
            },
            {
                "video_type": "movies",
                "content_id": "17",
                "title": "The Beach",
                "synopsis": "Richard seorang backpacker asal Amerika melakukan perjalanan ke Thailand, disana dia menerima sebuah peta yang konon berisi perjalanan menuju tempat yang disebut surga.",
                "duration": "111",
                "artist": "Leonardo DCaprio , Daniel York",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/17/the-beach-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/17/the-beach-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/17/the-beach-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-beach"
            },
            {
                "video_type": "series",
                "content_id": "18",
                "title": "The Jones",
                "synopsis": "cerita tentang 2 cowok yang selalu apes dalam mengejar pacar impian mereka",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/18/jones-cov-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/18/jones-cov-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/18/jones-cov-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/the-jones"
            },
            {
                "video_type": "movies",
                "content_id": "19",
                "title": "The Others",
                "synopsis": "Grace tinggal dirumah megah namun usang bersama kedua anaknya yang mengidap penyakit photosensitive. Dia mempekerjakan tiga pembantu rumah tangga yang membawa mereka pada kejadian yang tidak diinginkan.",
                "duration": "94",
                "artist": "Nicole Kidman , C Eccleston",
                "genre": [
                    "Laga",
                    "Horor"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/19/the-others2-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/19/the-others2-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/19/the-others2-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-others"
            },
            {
                "video_type": "movies",
                "content_id": "20",
                "title": "Erau Kota Raja",
                "synopsis": "Jurnalis wanita yang menjalin asmara dengan Reza, penduduk setempat yang menjadi pendampingnya saat bertugas meliput festival budaya di Tenggarong, Kutai Kartanegara, Kalimantan Timur.",
                "duration": "91",
                "artist": "Nadine C , Denny Sumargo",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/20/Erau-Kota-Raja-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/20/Erau-Kota-Raja-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/20/Erau-Kota-Raja-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/erau-kota-raja"
            },
            {
                "video_type": "movies",
                "content_id": "21",
                "title": "Liar",
                "synopsis": "Perjuangan Indra dan Bayu dua anak muda yang bercita - cita menjadi pembalap motor. Berbagai rintangan dan cara mereka tempuh demi meraih impian mereka.",
                "duration": "108",
                "artist": "Raffi Ahmad , Asmirandah",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/21/liar-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/21/liar-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/21/liar-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/liar"
            },
            {
                "video_type": "movies",
                "content_id": "22",
                "title": "Sword Fish",
                "synopsis": "Gabriel Shear mantan agen Mossad, Israel yang memiliki dendam tersendiri kepada pemerintah. Dia memanfaatkan Stanley seorang hacker handal untuk menjebol keamanan bank.",
                "duration": "111",
                "artist": "John Travolta , Hugh Jackman",
                "genre": [
                    "Horor"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/22/sword-fish-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/22/sword-fish-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/22/sword-fish-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/sword-fish"
            },
            {
                "video_type": "movies",
                "content_id": "23",
                "title": "Badai Pasti Berlalu",
                "synopsis": "Siska yang baru dikhianati oleh kekasihnya, kembali mengalami patah hati akibat dipermainkan oleh Leo. Hingga kemudian dia jatuh kedalam pelukan Helmi manajer kafe yang berkarakter licik.",
                "duration": "102",
                "artist": "Vino Bastian , Raihaanun",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/23/Badai-Pasti-Berlalu-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/23/Badai-Pasti-Berlalu-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/23/Badai-Pasti-Berlalu-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/badai-pasti-berlalu"
            },
            {
                "video_type": "movies",
                "content_id": "24",
                "title": "Paranormal Activity",
                "synopsis": "Pasangan muda, Katie dan Micah dihantui oleh kekuatan supranatural di rumah mereka. Mereka memasang kamera untuk mengetahui apa yang menghantui selama ini.",
                "duration": "84",
                "artist": "Katie F , Micah Sloat",
                "genre": [
                    "Horor"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/24/Thumbnail_Paranormal-Activity_300x400-.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/24/Thumbnail_Paranormal-Activity_400x300-.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/24/Thumbnail_Paranormal-Activity_1280x720-.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/paranormal-activity"
            },
            {
                "video_type": "movies",
                "content_id": "25",
                "title": "The Score",
                "synopsis": "nick wells, seorang kriminal yang memiliki banyak catatan dalam tindakan kejahatannya. Sudah banyak sekali tindakan kriminal yang dilakukan olehnya terutama dalam bidang pencurian. Nick wells terkenal sebagai pencuri yang handal dan memiliki beragam cara dan rencana sehingga membuat segala tindakan pencuriannya selalu berjalan mulus. Namun setiap orang selalu menemukan titik balik dalam hidupnya. Tidak terkecuali bagi nick wells.",
                "duration": "124",
                "artist": "Robert De Niro ,  Edward Norton",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/25/Score-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/25/Score-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/25/Score-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-score"
            },
            {
                "video_type": "movies",
                "content_id": "27",
                "title": "Ada Surga di Rumahmu",
                "synopsis": "Ramadhan, anak nakal yang sering menyusahkan orang tuanya. Setelah dewasa banyak peristiwa - peristiwa yang membenturnya hingga membuatnya sadar akan kesalahannya.",
                "duration": "101",
                "artist": "Elma Theana , Husein Alatas",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/27/adasurgadirumahmu300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/27/adasurgadirumahmu400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/27/adasurgadirumahmu1280720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/ada-surga-di-rumahmu"
            },
            {
                "video_type": "movies",
                "content_id": "28",
                "title": "3 Hati, Dua Dunia, Satu Cinta",
                "synopsis": "Film Tiga Hati Dua Dunia Satu Cinta adalah sebuah kisah cinta Seorang pemuda muslim. Seorang gadis katolik. Will they live happily ever after?",
                "duration": "103",
                "artist": "Reza Rahadian , Laura Basuki, Arumi Bachsin",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/28/3hati2dunia1cinta300400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/28/3hati2dunia1cinta400300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/28/3hati2dunia1cinta1280720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/3-hati-dua-dunia-satu-cinta"
            },
            {
                "video_type": "movies",
                "content_id": "30",
                "title": "catatan akhir kuliah",
                "synopsis": "3 orang pria yang sudah brsahabat, mereka bertekad pasti mereka akan wisuda bersama yang bernama Sam Maulana, Sobari, Ajeb. Namun 1 diantara mereka yaitu Sam belum mengerjakan tugas akhir sekolah yaitu Skripsi.",
                "duration": "111",
                "artist": "Muhadkly Acho , Ajun Perwira",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/30/catatan-akhir-kuliah-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/30/catatan-akhir-kuliah-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/30/Catatan-Akhir-Kuliah-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/catatan-akhir-kuliah"
            },
            {
                "video_type": "movies",
                "content_id": "32",
                "title": "Garuda 19",
                "synopsis": "Perjuangan spektakuler Timnas U-19 dalam menorehkan prestasi ditengah kondisi ‘kritis’.",
                "duration": "98",
                "artist": "Mathias Muchus , Ibnu Jamil",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/32/garuda-19-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/32/garuda-19-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/32/garuda191280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/garuda-19"
            },
            {
                "video_type": "movies",
                "content_id": "35",
                "title": "The Bang Bang Club",
                "synopsis": "Kisah nyata dari empat fotografer muda pemberani yang masuk ke dalam pertempuran ras yang terjadi di Afrika Selatan pada era 90an.",
                "duration": "90",
                "artist": "Ryan Phillippe , Malin Akerman",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/35/TheBangBangClub300400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/35/TheBangBangClub400300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/35/thebangbangclub1280720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-bang-bang-club"
            },
            {
                "video_type": "movies",
                "content_id": "40",
                "title": "Crazy Little Thing Called Love",
                "synopsis": "Khun Nam yang duduk dikelas M.1 (1 SMP) sudah jatuh hati pada seniornya, Khun Shone yang duduk dikelas M.4 (1 SMA). Nam yang berkulit gelap, berambut pendek, berkacamata dan culun merasa tidak cocok dengan Shone. Terlebih lagi, Shone sangat populer disekolahnya. Tapi karena cintanya, Nam bersemangat untuk menjadi yang terbaik dan bisa berbicara juga dekat dengan Shone.",
                "duration": "118",
                "artist": "Pimchanok Luevisadpaibul , Mario Maurer",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/40/cuCrazyLittleThingCalledLove300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/40/cuCrazyLittleThingCalledLove400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/40/cuCrazyLittleThingCalledLove1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/crazy-little-thing-called-love"
            },
            {
                "video_type": "movies",
                "content_id": "41",
                "title": "SITI",
                "synopsis": "Satu hari dalam kehidupan Siti, 24 tahun. Siti, ibu muda, mengurusi ibu mertuanya, Darmi, anaknya, Bagas, dan suaminya, Bagus. Bagus mengalami kecelakaan saat melaut setahun sebelumnya yang mengakibatkan kelumpuhan. Kapal Bagus yang baru dibeli dengan uang pinjaman hilang di laut.",
                "duration": "88",
                "artist": "Sekar Sari , Bintang Timur",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/41/siti-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/41/siti-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/41/siti-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/siti"
            },
            {
                "video_type": "movies",
                "content_id": "42",
                "title": "Ghost World",
                "synopsis": "Kehidupan dua sahabat Enid dan Rebecca yang selalu berbuat iseng dan membuat orang lain jengkel akan tingkahnya.",
                "duration": "110",
                "artist": "Thora Birch , Scarlett J",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/42/ghost-world-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/42/ghost-world-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/42/ghost-world-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/ghost-world"
            },
            {
                "video_type": "movies",
                "content_id": "45",
                "title": "Erin Brockovich",
                "synopsis": "Seorang ibu beranak tiga yang bercerai dua kali, berusaha mencari keadilan dan selalu dihadapkan dengan orang-orang jahat namun dia tetap menang.",
                "duration": "128",
                "artist": "Julia Roberts , David Brisbin",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/45/brockovich-300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/45/erin-brockovich-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/45/erin-brockovich-1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/erin-brockovich"
            },
            {
                "video_type": "series",
                "content_id": "46",
                "title": "Halal Bites",
                "synopsis": "",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Religi"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/46/HalalBites300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/46/HalalBites400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/46/Halalbites1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/halal-bites"
            },
            {
                "video_type": "movies",
                "content_id": "48",
                "title": "Rock Star",
                "synopsis": "Chris adalah tukang reparasi dan juga vokalis band tribute dimana ia bisa menirukan secara luar biasa suara penyanyi aslinya. Dengan perjalanan nasib, ia menjadi pemimpin band yang sebenarnya.",
                "duration": "101",
                "artist": "Mark Wahlberg , Jennifer Aniston",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/48/rockstar300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/48/rockstar400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/48/rockstar1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/rock-star"
            },
            {
                "video_type": "movies",
                "content_id": "50",
                "title": "The Gift",
                "synopsis": "Seorang wanita yang memiliki indra keenam diminta untuk membantu menemukan seorang wanita muda yang hilang.",
                "duration": "81",
                "artist": "Cate Blanchett , Katie Holmes",
                "genre": [
                    "Horor"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/50/thegift300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/50/thegift400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/50/thegift1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-gift"
            },
            {
                "video_type": "movies",
                "content_id": "51",
                "title": "New In Town",
                "synopsis": "Sebidang tanah eksekutif di sebuah kota kecil Minnesota Miami. Berawal pada cuaca dingin & sama dinginnya penyambutan perjuangan dia di tempat kerja, tapi kota & orang-orang mulai akrab pada dirinya.",
                "duration": "60",
                "artist": "Renée Zellweger ,  Harry Connick, Jr.",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/51/cuNewInTown300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/51/cuNewInTown400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/51/cuNewInTown1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/new-in-town"
            },
            {
                "video_type": "movies",
                "content_id": "56",
                "title": "Jakarta Undercover",
                "synopsis": "Pras berada di tengah pusaran dunia antah berantah Jakarta. Semakin dalam ia menggali, semakin jauh ia terserat ke dalamnya. Bagaimana Pras menghadapi kerasnya kehidupan Jakarta?",
                "duration": "107",
                "artist": "Oka Antara , Lukman Sardi",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/56/jakartaundercover300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/56/jakartaundercover400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/56/jakartaundercover1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/jakarta-undercover"
            },
            {
                "video_type": "movies",
                "content_id": "57",
                "title": "Madre",
                "synopsis": "Tansen, peselancar yang memilih hidup bebas, mendapat panggilan untuk menerima warisan berupa biang roti bernama Madre dari mendiang kakeknya, pengusaha roti terkenal.",
                "duration": "111",
                "artist": "Vino G. Bastian , Laura Basuki",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/57/madre300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/57/madre400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/57/madre1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/madre"
            },
            {
                "video_type": "movies",
                "content_id": "60",
                "title": "Rumah Pondok Indah",
                "synopsis": "Sebuah keluarga kecil yang baru saja membeli rumah bekas Tio, seorang pematung ternama. Baru saja menempati rumah itu keluarga kecil ini tewas hanya dalam hitungan 24 jam. Diawali oleh anak satu-satunya mereka mati kesetrum, kemudian dengan tergesa-gesa dibawa ke rumah sakit, namun naas tidak jauh dari rumah mobil mereka ditabrak oleh sebuah kopaja yang lewat.",
                "duration": "90",
                "artist": "Arswendy Nasution , Asha Shara , Ricky Harun",
                "genre": [
                    "Horor"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/60/rumahpondokindah.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/60/rumahpondokindah400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/60/rumahpondokindah1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/rumah-pondok-indah"
            },
            {
                "video_type": "movies",
                "content_id": "61",
                "title": "Sepatu Dahlan",
                "synopsis": "“Hidup, bagi orang miskin, harus dijalani apa adanya,” begitulah prinsip Dahlan. Ia tidak pernah berhenti bermimpi untuk memiliki sepatu dan sepeda. Kemiskinan yang dirasakannya, tidak menyurutkan semangat Dahlan untuk tetap bersekolah meski harus bertelanjang kaki, berjalan puluhan kilometer untuk sampai di pesantren Takeran. Dan tak jarang kakinya melepuh bahkan lecet.",
                "duration": "100",
                "artist": "Kinaryosih , Aji Santosa",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/61/sepatudahlan300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/61/sepatudahlan400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/61/sepatudahlan1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/sepatu-dahlan"
            },
            {
                "video_type": "series",
                "content_id": "62",
                "title": "Petualangan Masjid Bersejarah Nusantara",
                "synopsis": "Petualangan menelusuri Masjid bersejarah Nusantara Indonesia",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Travel",
                    "Religi"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/62/PetualanganMasjid.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/62/petualanganmasjid.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/62/petualanganmasjid.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/petualangan-masjid-bersejarah-nusantara"
            },
            {
                "video_type": "series",
                "content_id": "65",
                "title": "Adeging Nagari Indonesia",
                "synopsis": "",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/65/cu_AdegingNagariIndonesia.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/65/cu_AdegingNagariIndonesia.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/65/cu_AdegingNagariIndonesia.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/adeging-nagari-indonesia"
            },
            {
                "video_type": "movies",
                "content_id": "66",
                "title": "Bangkit",
                "synopsis": "Jakarta terkena dampak badai musim dingin di Asia dan juga musim panas dari Australia. Di saat yang sama terjadi gempa bumi berkekuatan besar. Mampukah tim Basarnas menyelamatkan para korban?",
                "duration": "122",
                "artist": "Vino G. Bastian , Acha Septriasa",
                "genre": [
                    "Laga",
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/66/Bangkit.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/66/Bangkit.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/66/Bangkit1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/bangkit"
            },
            {
                "video_type": "movies",
                "content_id": "68",
                "title": "Taken",
                "synopsis": "Seorang agen CIA yang baru pensiun, Bryan Mills, mencoba membina hubungan lebih dekat dengan putrinya, Kim, yang tinggal bersama ibunya, Lenore, dan ayah tirinya yang kaya, Stuart. Kim meminta izin ayahnya untuk boleh bepergian ke Paris hanya bersama teman perempuannya, Amanda. Bryan mulanya ragu, tetapi setuju setelah ditekan oleh Lenore. Di airport, Bryan kecewa karena ternyata kedua gadis itu sebenarnya mengikuti tur U2 keliling Eropa dan hal itu ditutupi oleh Kim, meskipun Lenore tahu dan tidak mengabari Bryan.",
                "duration": "93",
                "artist": "Liam Neeson , Maggie Grace",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/68/Takenx300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/68/Takenx400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/68/Taken1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/taken"
            },
            {
                "video_type": "movies",
                "content_id": "69",
                "title": "Adeging Majapahit",
                "synopsis": "",
                "duration": "100",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/69/cu_AdegingMajapahit.3001.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/69/cu_AdegingMajapahit.4001.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/69/cu_AdegingMajapahit.12801.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/adeging-majapahit"
            },
            {
                "video_type": "series",
                "content_id": "70",
                "title": "Adipati Sengguruh",
                "synopsis": "",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/70/cu_Sengguruh.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/70/cu_Sengguruh.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/70/Sengguruh1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/adipati-sengguruh"
            },
            {
                "video_type": "movies",
                "content_id": "71",
                "title": "Greenberg",
                "synopsis": "Seorang pria empatpuluhan yang baru keluar dari rumah sakit mental, Roger Greenberg  pulang dari New York ke kota asalnya Los Angeles dan tinggal di rumah adiknya yang sedang berlibur selama 6 minggu hanya untuk ‘do-nothing’. Ia kembali ke L.A., kota dimana ia tumbuh besar dan pernah membentuk rock-band yang hampir sukses. Greenberg berusaha untuk menjalin komunikasi lagi dengan teman – teman baiknya semasa kuliah, salah satunya Ivan  yang sedang dalam proses perceraian dengan istrinya. Greenberg kemudian menemukan bahwa teman – teman baiknya kini tidak lagi teman baiknya.  Kemudian, ia pun menjalin hubungan dengan asisten adiknya, Florence yang baru saya putus dari long relationship.",
                "duration": "107",
                "artist": "Ben Stiller , Rhys Ifans, Greta Garwig",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/71/Greenberg300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/71/Greenberg.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/71/Greenberg1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/greenberg"
            },
            {
                "video_type": "movies",
                "content_id": "72",
                "title": "Jarhead",
                "synopsis": "Kehidupan para prajurit muda Amerika Serikat pada saat bertugas dalam perang teluk di Timur Tengah.",
                "duration": "125",
                "artist": "Jake Gyllenhaal , Jamie Foxx",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/72/jarhead300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/72/jarhead400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/72/jarhead1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/jarhead"
            },
            {
                "video_type": "movies",
                "content_id": "73",
                "title": "The 6th Day",
                "synopsis": "Di dunia masa depan, teknologi kloning manusia telah jatuh ke tangan perusahaan multinasional yang korup. Tapi, seorang pria menyadarinya dan ia tak mau diperalat dalam konspirasi berbahaya ini.",
                "duration": "124",
                "artist": "arnold schwarzneger , Michael Rapaport",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/73/the6thday300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/73/the6thday400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/73/the6thday1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-6th-day"
            },
            {
                "video_type": "movies",
                "content_id": "74",
                "title": "Kiss Kiss Bang Bang",
                "synopsis": "Pencuri kelas teri menyamar sebagai aktor, dikirim ke Los Angeles untuk sebuah audisi, namun tiba-tiba menemukan dirinya justru di selidiki atas kasus pembunuhan.",
                "duration": "96",
                "artist": "Robert Downey Jr. , Val Kilmer",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/74/kiskisbangbang300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/74/kiskisbangbang400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/74/kiskisbangbang1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/kiss-kiss-bang-bang"
            },
            {
                "video_type": "movies",
                "content_id": "76",
                "title": "What lies Beneath",
                "synopsis": "Seorang ilmuwan berusaha menentukan sikapnya, ketika istrinya percaya bahwa rumah yang mereka tinggali, dihantui oleh kekuatan jahat.",
                "duration": "128",
                "artist": "Michelle P , Katharine Towne",
                "genre": [
                    "Horor",
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/76/whatliesbeneath300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/76/whatliesbeneath400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/76/whatliesbeneath1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/what-lies-beneath"
            },
            {
                "video_type": "movies",
                "content_id": "78",
                "title": "The Velveteen Rabbit",
                "synopsis": "Kisah seorang anak yang berhasil melarikan diri dari ayah dan neneknya yang jahat dengan dibantu oleh seekor kelinci.",
                "duration": "90",
                "artist": "Jane Seymour , Ellen Burstyn",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/78/TheVelveteenRabbit300x400.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/78/TheVelveteenRabbit400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/78/TheVelveteenRabbit1280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-velveteen-rabbit"
            },
            {
                "video_type": "movies",
                "content_id": "79",
                "title": "Panembahan Sedo Krapyak",
                "synopsis": "Ketropak Humor Remaja yang diproduksi oleh SMKN 1 TRUCUK, berjudul Panembahas Sedo Krapyak",
                "duration": "54",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/79/cu_KHP_PanembahasSedoKrapyak.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/79/cu_KHP_PanembahasSedoKrapyak.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/79/cu_KHP_PanembahasSedoKrapyak.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/panembahan-sedo-krapyak"
            },
            {
                "video_type": "series",
                "content_id": "80",
                "title": "Fashion Magz",
                "synopsis": "",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Gaya"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/80/FashionMagz300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/80/FashionMagz400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/80/FashionMagz1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/fashion-magz"
            },
            {
                "video_type": "movies",
                "content_id": "81",
                "title": "Sabaya Mukti Sabaya Pati",
                "synopsis": "Ketoprak Humor Pelajar Remaja oleh SMK Kristen 2 Klaten",
                "duration": "61",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/81/cu_KHP_SabayaMuktiSabayaPati.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/81/cu_KHP_SabayaMuktiSabayaPati.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/81/cu_KHP_SabayaMuktiSabayaPati.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/sabaya-mukti-sabaya-pati"
            },
            {
                "video_type": "movies",
                "content_id": "83",
                "title": "Joko Tingkir Winisudo",
                "synopsis": "Ketoprak Humor Pelajar yang diproduksi  SMK N 1 GANTIWARNO",
                "duration": "58",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/83/cu_KHP_JokoTingkirWinisudo.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/83/cu_KHP_JokoTingkirWinisudo.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/83/cu_KHP_JokoTingkirWinisudo.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/joko-tingkir-winisudo"
            },
            {
                "video_type": "movies",
                "content_id": "84",
                "title": "Aji Saka",
                "synopsis": "Ketoprak Humor Pelajar yang diproduksi oleh SMK SWADAYA KLATEN",
                "duration": "53",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/84/cu_KHP_AjiSaka.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/84/cu_KHP_AjiSaka.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/84/cu_KHP_AjiSaka.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/aji-saka"
            },
            {
                "video_type": "movies",
                "content_id": "85",
                "title": "Age Of Heroes",
                "synopsis": "Didasarkan pada cerita nyata Ian Flemming, penulis cerita James Bond, mengenai pasukan elite Inggris bernama Komando 30. Film ini mengisahkan tentang Kopral Rains seorang anggota Komando 30. Rains dan seluruh anggota pasukan elit tersebut dilatih fisik dan mental hingga titik akhir kemampuan mereka sebagai manusia untuk sebuah misi rahasia yang berbahaya, Rains dan pasukannya akan diterjunkan di pegunungan Norwegia yang dingin dengan parasut untuk menghindari pantauan radar Jerman. Jika mereka berhasil, maka misi ini dapat mengubah peta perang dunia di masa tersebut.",
                "duration": "120",
                "artist": "Sean Bean , Danny Dyer",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/85/AgeOfHeroes300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/85/AgeOfHeroes400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/85/AgeOfHeroes1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/age-of-heroes"
            },
            {
                "video_type": "movies",
                "content_id": "86",
                "title": "In The Name of The King a Dungeon Siege Tale",
                "synopsis": "Kerajaan Ebb yang semula tentram dan damai tiba-tiba diserang oleh sepasukan mahluk setengah manusia setengah binatang. Mahluk bernama Krug itu memporak porandakan seluruh isi negeri.",
                "duration": "156",
                "artist": "Jason Statham , Ray Liotta , Burt Reynolds",
                "genre": [
                    "Laga",
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/86/InTheNameOfTheKing300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/86/InTheNameOfTheKing400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/86/InTheNameOfTheKing1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/in-the-name-of-the-king-a-dungeon-siege-tale"
            },
            {
                "video_type": "movies",
                "content_id": "87",
                "title": "Senjakala di Manado",
                "synopsis": "Johny WW Lengkong, seorang pelaut yang telah meninggalkan keluarganya selama hampir 20 tahun, berniat kembali ke Manado dengan harapan keluarganya mau memaafkannya. Namun ternyata istrinya telah meninggal dunia dan meninggalkan seorang anak bernama Pingkan. Johny pun berjanji akan merawat Pingkan, namun ternyata semua tidak berjalan dengan mudah. Pingkan merasa terganggu tatkala Johny mulai mencampuri urusan asmaranya dengan Brando, kekasihnya, lantaran Johny merasa Brando bukan pria baik-baik. Dari situlah konflik antara ayah dan anak mulai terjadi",
                "duration": "95",
                "artist": "Ray Sahetapy ,  Mikha Tambayong",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/87/SenjakalaDiManado300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/87/SenjakalaDiManado400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/87/SenjakalaDiManado1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/senjakala-di-manado"
            },
            {
                "video_type": "movies",
                "content_id": "90",
                "title": "Cold Mountain",
                "synopsis": "Cold mountain mengisahkan tentang perang saudara yang masih berkecamuk dimana-mana. Salah satunya yang terjadi di petersburg, virgina. Banyak orang-orang yang tidak memiliki latar belakang militer namun dipaksa untuk menjadi serdadu yang kemudian dikirimkan ke medan perang. Tak ayal banyak warga sipil yang menjadi tentara meskipun mereka tidak pernah mengangkat senjata sebelumnya. Sehingga para tentara amatir tersebut harus rela mengorbankan nyawanya demi kepentingan negaranya.",
                "duration": "130",
                "artist": "Jude Law , Nicole Kidman",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/90/ColdMountain300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/90/cuColdMountain400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/90/cuColdMountain1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/cold-mountain"
            },
            {
                "video_type": "series",
                "content_id": "91",
                "title": "Before Green Gables",
                "synopsis": "Anne Shirley menjalani cukup banyak petualangan sebelum dia berakhir di Green Gables. Dia tinggal dengan keluarga dan membantu dengan menjaga anak-anak, membantu pekerjaan di rumah ... Sementara itu dia menemukan lebih banyak tentang siapa dia, bahwa namanya adalah \"Anne\" dan bukan \"Ann\" ... Dan dia manis tetapi sedikit gila",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/91/BeforeGreenGables-Eps1-300x400-RevOK.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/91/BeforeGreenGablesEps1400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/91/BeforeGreenGablesEps11280x720.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/before-green-gables"
            },
            {
                "video_type": "series",
                "content_id": "92",
                "title": "Fantastic Children",
                "synopsis": "Helga, seorang anak yatim piatu berusia 11 tahun yang introvert yang menggambar sebuah negeri dengan bulan sabit yang dia yakini adalah rumahnya. Teman mainnya dan satu-satunya teman di panti asuhan, Chitto, ingin membantu Helga menemukannya. Jadi bersama-sama mereka melarikan diri dari panti asuhan dan memulai perjalanan di mana mereka bertemu Tohma, seorang bocah energik di rumahnya, Pulau Papin.",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/92/FantasticChildren-Eps1-300x400RevOK.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/92/FantasticChildren-Eps1-400x300.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/92/FantasticChildren1280x1.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/fantastic-children"
            },
            {
                "video_type": "movies",
                "content_id": "93",
                "title": "The Thieves",
                "synopsis": "Tiga pencuri professional yang telah bekerja sama cukup lama. Mereka bertiga adalah popie, yenical dan juga chewigum. Mereka bertiga ternyata telah menjadi seorang pencuri professional semenjak usia mereka masih muda. Sehingga tidak heran apabila mereka memiliki kemampuan mencuri yang tidak perlu diragukan lagi. Bahkan bisa dibilang, banyak benda-benda yang sulit bahkan sulit untuk dicuri ternyata mereka berhasil curi.",
                "duration": "139",
                "artist": "Kim Hye Soo, Jun Ji Hyun, Kim Soo Hyun",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/93/TheThieves300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/93/TheThieves400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/93/TheThieves1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/the-thieves"
            },
            {
                "video_type": "movies",
                "content_id": "94",
                "title": "Season Of The Witch",
                "synopsis": "para penyihir di abad ke 3 di sebuah tempat bernama villach. Pada suatu hari, tiga perempuan dituduh sebagai seorang penyihir oleh salah seorang pastur. Salah satu diantara mereka bertiga itu menuduh bahwa tuduhan pastur itu tak berdasar. Ia menolak tuduhan itu dan mengutuk sang pastur. Sementara itu, pastur tersebut menyuruh untuk menggantung tiga perempuan itu dan menenggelamkan mereka. Ketika mereka bertiga sudah mati, para penjaga diperintahkan menghancurkan semua alat ritual mereka agar tiga penyihir itu tak bisa bangkit lagi.",
                "duration": "84",
                "artist": "Nicolas Cage , Stephen Campbell Moore",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/94/SeasonOfTheWitch300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/94/SeasonOfTheWitch400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/94/SeasonOfTheWitch1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/season-of-the-witch"
            },
            {
                "video_type": "movies",
                "content_id": "96",
                "title": "Pride and Prejudice",
                "synopsis": "Sebuah pedesaan yang tinggalah sebuah keluarga bennet. Ia dan istrinya memiliki lima orang anak perempuan yang cantik. Kelima putrinya tersebut adalah elizabeth, lydia, jane, mary, dan kitty. Keluarga ini hidup dengan bahagia dan penuh dengan kedamaian meskipun terkadang disisi lain kebahagiaan keluarganya tersebut istrinya tuan bennet selalu memikirkan orang yang akan menikahi kelima putrinya tersebut. Kondisi keuangan keluarganya sudah tidak baik lagi karena peternakan yang menjadi tempat bekerja mereka akan diwariskan kepada collins sepupu dari tuan bennet.",
                "duration": "135",
                "artist": "Keira Knightley",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/96/Pride&Prejudice300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/96/Pride&Prejudice400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/96/Pride&Prejudice1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/pride-and-prejudice"
            },
            {
                "video_type": "movies",
                "content_id": "97",
                "title": "A Werewolf Boy",
                "synopsis": "Seorang ibu pindah rumah dengan anak perempuannya yang lebih tua (Park Bo-Young) dan putri yang lebih muda (Kim Hyang-Gi) ke sebuah rumah besar di pedesaan dikarenakan anak perempuan tertua menderita penyakit paru-paru dan dokternya menyarankan keluarga untuk pindah ke pedesaan Rumah ini disediakan oleh Ji-Tae (Yoo Yeon-Seok), putra seorang mitra bisnis yang bekerja dengan ayah almarhum mereka. Ketika keluarga terbiasa dengan lingkungan baru mereka, putri tertua membuat penemuan yang luar biasa. Di dalam ruangan terkunci di gudang, tinggal seorang werewolf boy (Song Joong-Ki). Keluarganya percaya bahwa dia hanyalah anak yatim dengan sedikit keterampilan sosial. Anak laki-laki tersebut menyukai putri tertua keluarga tersebut. Sementara itu, Ji-Tae memiliki rencana sendiri untuk menikahi putri tertua. Apa rahasia dari werewolf boy itu?",
                "duration": "126",
                "artist": "Park Bo Young, Song Joong  Ki",
                "genre": [
                    "Laga",
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/97/AWerewolfBoy300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/97/AWerewolfBoy400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/97/AWerewolfBoy1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/a-werewolf-boy"
            },
            {
                "video_type": "series",
                "content_id": "98",
                "title": "Campursari",
                "synopsis": "",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Gaya"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/98/cu_Campursari.300_1.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/98/cu_Campursari.400_1.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/98/Campursari1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/campursari"
            },
            {
                "video_type": "movies",
                "content_id": "99",
                "title": "Man Down",
                "synopsis": "Shia LaBeouf kembali ke layar lebar dalam film bergerne drama-thriller bersetting militer “Man Down” yang disutradarai oleh Dito Montiel. Gabriel Drummer adalah seorang ayah dari keluarga Drummer yang beranggotakan Natalie dan Jonathan,anak mereka. Bersama Devin Roberts, sahabatnya , mereka masuk menjadi anggota pasukan angkatan laut militer Amerika yang kemudian dikirim ke Afghanistan. Sepulangnya dari berperang,Gabriel menemukan tempat tinggalnya sama dengan kondisi medan perang yang Ia lalui. Ia putus asa mencari keberadaan anak dan istrinya sampai bertemu seorang gelandangan bernama Charles yang diyakini memegang kunci dimana istri dan anaknya berada.",
                "duration": "92",
                "artist": "Gary Oldman, Jai Courtney",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/99/ManDown300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/99/ManDown400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/99/ManDown1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/man-down"
            },
            {
                "video_type": "movies",
                "content_id": "100",
                "title": "Get Carter",
                "synopsis": "Seorang penegak hukum Las Vegas,melakukan perjalanan kembali ke rumahnya untuk menyelidiki kematian misterius saudara-saudaranya.",
                "duration": "120",
                "artist": "Sylvester Stallone , Michael Caine",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/100/GetCarter300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/100/GetCarter400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/100/GetCarter1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/get-carter"
            },
            {
                "video_type": "movies",
                "content_id": "102",
                "title": "Slumdog Millionaire",
                "synopsis": "Slumdog millionaire mengisahkan tentang jamal yang mengikuti sebuah acara kuis yang sangat terkenal bernama who wants to be a millionare. Tak disangka, jamal menjawab semua pertanyaan dari kuis tersebut dengan benar. Namun ketika pertanyaan terakhir akan dibacakan sekaligus menjadi pertanyaan yang bisa membuat jamal membawa uang sebanyak 20 juta rupee ternyata durasi dari acara tersebut telah habis. Sehingga membuat pihak acara memutuskan untuk melanjutkan acaranya di esok hari",
                "duration": "120",
                "artist": "Dev Patel , Freida Pinto",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/102/SlumdogMillionaire300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/102/SlumdogMillionaire400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/102/SlumdogMillionaire1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/slumdog-millionaire"
            },
            {
                "video_type": "movies",
                "content_id": "103",
                "title": "Unleashed",
                "synopsis": "Danny adalah seorang petarung kuat yang dibesarkan bagaikan seekor anjing oleh pemiliknya, seorang gangster. Danny adalah predator dia akan bertarung dan membunuh siapapun sesuai instruksi majikannya. Pikiran dan kepribadiannya seperti anak kecil dan dia tak pernah menjalani kehidupan normal. Dalam suatu kejadian, Danny terluka parah dan koma, lalu dia dirawat oleh orang-orang baik. Akankah dia bisa berubah menjadi normal ataukah naluri buasnya hidup lagi?",
                "duration": "140",
                "artist": "Jet Li, Bob Hoskins, Morgan Freeman",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/103/Unleashed300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/103/Unleashed400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/103/Unleashed1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/unleashed"
            },
            {
                "video_type": "movies",
                "content_id": "104",
                "title": "Hello Stranger",
                "synopsis": "Hello Stranger juga dibintangi oleh Chantavit Dhanasevi. Kali ini ia berperan sebagai Dang, seorang lelaki yang baru saja diputuskan oleh pacarnya. Untuk mengobati rasa galaunya, ia kemudian ikut dalam sebuah tour ke Korea Selatan. Namun sesampainya di sana, ia malah terpisah dari rombongannya.Di sana, tanpa sengaja ia bertemu dengan seorang wanita asal Thailand yang juga sedang berlibur. Ia kemudian mengikuti wanita tersebut ke berbagai tempat di Korea.",
                "duration": "202",
                "artist": "Chantavit Dhanasevi, Nuengthida Sophon",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/104/HelloStranger300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/104/HelloStranger400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/104/HelloStranger1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/hello-stranger"
            },
            {
                "video_type": "movies",
                "content_id": "105",
                "title": "Once Upon a Time in Venice",
                "synopsis": "Seorang detektif Los Angeles mencari geng kejam yang mencuri anjingnya.",
                "duration": "132",
                "artist": "Bruce Willis, John Goodman, Jason Momoa",
                "genre": [
                    "Laga"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/105/cu_OnceUponATimeinVenice.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/105/cu_OnceUponATimeinVenice.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/105/cu_OnceUponATimeinVenice.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/once-upon-a-time-in-venice"
            },
            {
                "video_type": "movies",
                "content_id": "106",
                "title": "Jahitkan Cintamu Di Hatiku",
                "synopsis": "Seorang pemuda yang berprofesi penjahit keliling diperebutkan cintanya oleh dua wanita..",
                "duration": "100",
                "artist": "Soni Septian , Michella Adlen",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/106/MovinesiaJahitkanCintamuDihatiku300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/106/MovinesiaJahitkanCintamuDihatiku400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/106/MovinesiaJahitkanCintamuDihatiku1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/jahitkan-cintamu-di-hatiku"
            },
            {
                "video_type": "movies",
                "content_id": "107",
                "title": "Saw V",
                "synopsis": "Setelah kematian Jigsaw yang mengerikan, Mark Hoffman dipuji sebagai pahlawan, tetapi Agen Strahm curiga, dan menggali masa lalu Hoffman. Sementara itu, sekelompok orang lain menjalani serangkaian tes yang mengerikan.",
                "duration": "135",
                "artist": "Scott Patterson, Costas Mandylor, Tobin Bell",
                "genre": [
                    "Horor"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/107/cu_SawV.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/107/cu_SawV.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/107/cu_SawV.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/saw-v"
            },
            {
                "video_type": "movies",
                "content_id": "108",
                "title": "Ustadz Ganteng Dari Amerika",
                "synopsis": "Seorang pemuda yang berubah penampilan lebih religius setelah pulang dari amerika",
                "duration": "112",
                "artist": "Rendy Samuel, Winda Khair",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/108/UstadzGantengDariAmerikaR1300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/108/UstadzGantengDariAmerikaR1400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/108/UstadzGantengDariAmerika1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/ustadz-ganteng-dari-amerika"
            },
            {
                "video_type": "movies",
                "content_id": "109",
                "title": "Interview With Radhini",
                "synopsis": "",
                "duration": "30",
                "artist": "Radhini",
                "genre": [
                    "Gaya"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/109/300x400_still.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/109/400x300_stil.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/109/1280x720_still.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/interview-with-radhini"
            },
            {
                "video_type": "movies",
                "content_id": "110",
                "title": "Ragam Budaya Indonesia \"Sintren\"",
                "synopsis": "Sintren merupakan tari kesenian tradisional yang berasal dari Banyumas",
                "duration": "9",
                "artist": "",
                "genre": [
                    "Gaya"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/110/SIntern300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/110/SIntern400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/110/RagamBudayaSIntren1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/ragam-budaya-indonesia-sintren"
            },
            {
                "video_type": "series",
                "content_id": "111",
                "title": "Adeging Kadipaten Mangkunagaran",
                "synopsis": "",
                "duration": "0",
                "artist": "",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/111/cu_KadipatenMangkunagaran.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/111/cu_KadipatenMangkunagaran.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/111/KadipatenMangkunagaran1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/series/adeging-kadipaten-mangkunagaran"
            },
            {
                "video_type": "movies",
                "content_id": "112",
                "title": "Kesenian Angklung Banyumasan",
                "synopsis": "",
                "duration": "10",
                "artist": "",
                "genre": [
                    "Gaya"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/112/RagamBudayaAngklung300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/112/RagamBudayaAngklung400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/112/RagamBudayaAngklung1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/kesenian-angklung-banyumasan"
            },
            {
                "video_type": "movies",
                "content_id": "113",
                "title": "Stop Kau Mencuri Hatiku",
                "synopsis": "Seorang sales Furnitur yang cantik berusaha tegar karena di tinggal  tunangan pacarnya.",
                "duration": "80",
                "artist": "Acha Sinaga, Indra Brotolaras",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/113/StopEngkauMencuriHatikuR1300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/113/StopEngkauMencuriHatikuR1400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/113/StopEngkauMencuriHatikuR11280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/stop-kau-mencuri-hatiku"
            },
            {
                "video_type": "movies",
                "content_id": "114",
                "title": "Pacarku Tukang Sunat",
                "synopsis": "Mengisahkan seorang pemuda yang merupakan tukang sunat yang menjalin cinta dengan mahasiswi di Jakarta",
                "duration": "116",
                "artist": "Sheila Dara, Fajar Gomez",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/114/EPS30-PacarkuTukangSunat300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/114/EPS30-PacarkuTukangSunat400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/114/EPS30-PacarkuTukangSunat1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/pacarku-tukang-sunat"
            },
            {
                "video_type": "movies",
                "content_id": "115",
                "title": "Gerobak Dangdut Pembawa Cinta",
                "synopsis": "",
                "duration": "60",
                "artist": "Yudhita, Joshua Otay",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/115/GerobakDangdutPembawaCinta_300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/115/GerobakDangdutPembawaCinta_400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/115/GerobakDangdutPembawaCinta1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/gerobak-dangdut-pembawa-cinta"
            },
            {
                "video_type": "movies",
                "content_id": "116",
                "title": "Ondel Ondel Pembawa Cinta",
                "synopsis": "Rian yang di usir dari rumah oleh ayahnya selama 1 bulan karena prestasi yang buruk dan kemudian bekerja sebagai ondel ondel",
                "duration": "120",
                "artist": "Angga Dirgantara, Nadira Naser",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/116/EPS88-OndelOndelPembawaCinta300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/116/EPS88-OndelOndelPembawaCinta400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/116/EPS88-OndelOndelPembawaCinta1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/ondel-ondel-pembawa-cinta"
            },
            {
                "video_type": "movies",
                "content_id": "117",
                "title": "I Hate But I Miss You",
                "synopsis": "Apa jadinya jika seorang penyiar radio mengalami cinlok dengan cleaning service?",
                "duration": "93",
                "artist": "Angelica Simperler, Raeshard",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/117/IHatebutIMissYouR1300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/117/IHatebutIMissYouR1400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/117/IHatebutIMissYouR11280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/i-hate-but-i-miss-you"
            },
            {
                "video_type": "movies",
                "content_id": "118",
                "title": "1 Kembang 3 Kumbang",
                "synopsis": "",
                "duration": "86",
                "artist": "Sheila Dara, Wafda",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/118/EPS58-1Kembang3Kumbang300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/118/EPS58-1Kembang3Kumbang400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/118/EPS58-1Kembang3Kumbang1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/kembang-3-kumbang"
            },
            {
                "video_type": "movies",
                "content_id": "119",
                "title": "Combro Combro Asmara",
                "synopsis": "Percintaan antara perempuan yang merupakan anak penjual combro dengan pemuda kaya yang angkuh.",
                "duration": "62",
                "artist": "Rikas Harsa, Maody Zanya",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/119/EPS33-CombroCombroAsmara300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/119/EPS33-CombroCombroAsmara400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/119/EPS33-CombroCombroAsmara1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/combro-combro-asmara"
            },
            {
                "video_type": "movies",
                "content_id": "120",
                "title": "Terbakar Cinta Tukang Jagung Bakar",
                "synopsis": "Seorang pemuda yang putus cinta pergi ke puncak dan bertemu seorang perempuan penjual jagung bakar dan kemudian mereka mulai dekat",
                "duration": "103",
                "artist": "Ferly Putra, Sheilla Dara",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/120/cu_Movinesia.TerbakarCintaTukangJagungBakar.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/120/cu_Movinesia.TerbakarCintaTukangJagungBakar.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/120/cu_Movinesia.TerbakarCintaTukangJagungBakar.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/terbakar-cinta-tukang-jagung-bakar"
            },
            {
                "video_type": "movies",
                "content_id": "121",
                "title": "Detektif Asmara",
                "synopsis": "Seorang mahasiswa dan mahasiswi yang dekat karena membantu permasalahan percintaan teman mereka.",
                "duration": "103",
                "artist": "Rama Michael, Rian Putri",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/121/cu_Movinesia.DetektifAsmara.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/121/cu_Movinesia.DetektifAsmara.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/121/cu_Movinesia.DetektifAsmara.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/detektif-asmara"
            },
            {
                "video_type": "movies",
                "content_id": "122",
                "title": "Cinta Roti Buaya",
                "synopsis": "Perjuangan Ojan didampingi oleh Zubaidah, saudara calon pengantin perempuan untuk mencari roti buaya untuk melengkapi acara lamaran adat Betawi.",
                "duration": "120",
                "artist": "Sehmi Bajry, Dea Annisa, Adelina Amelia",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/122/cu_Movinesia.CintaRotiBuaya.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/122/cu_Movinesia.CintaRotiBuaya.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/122/cu_Movinesia.CintaRotiBuaya.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/cinta-roti-buaya"
            },
            {
                "video_type": "movies",
                "content_id": "123",
                "title": "Cinta Paket Ekspres",
                "synopsis": "",
                "duration": "108",
                "artist": "Keira Sabhira, Pingky Ovien, Dwi Yabes",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/123/cu_Movinesia.CintaPaketExpress.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/123/cu_Movinesia.CintaPaketExpress.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/123/cu_Movinesia.CintaPaketExpress.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/cinta-paket-ekspres"
            },
            {
                "video_type": "movies",
                "content_id": "124",
                "title": "Bebek Cinta",
                "synopsis": "",
                "duration": "71",
                "artist": "Fendy Chauw, Kheira Sabhira",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/124/cu_BebekCinta.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/124/cu_BebekCinta.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/124/cu_BebekCinta.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/bebek-cinta"
            },
            {
                "video_type": "movies",
                "content_id": "125",
                "title": "Jangan Parkir Cinta Sembarangan",
                "synopsis": "Seorang pemuda kaya yang menyamar menjadi OB agar bisa mendapatkan cinta wanita incarannya yang berprofesi sebagai juru parkir.",
                "duration": "134",
                "artist": "Nicky Tirta, Metta Permadi",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/125/cu_Movinesia_JanganParkirCintaSebarangan.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/125/cu_Movinesia_JanganParkirCintaSebarangan.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/125/cu_Movinesia_JanganParkirCintaSebarangan.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/jangan-parkir-cinta-sembarangan"
            },
            {
                "video_type": "movies",
                "content_id": "126",
                "title": "Selebriti Jamu Gendong",
                "synopsis": "Seorang penjual jamu yang bertemu seorang selebriti dan berakhir menjadi romansa",
                "duration": "110",
                "artist": "Rorencya, Bella i Jasmine, Rendy Samuel",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/126/cu_Movinesia_SelebritiJamuGendong.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/126/cu_Movinesia_SelebritiJamuGendong.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/126/cu_Movinesia_SelebritiJamuGendong.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/selebriti-jamu-gendong"
            },
            {
                "video_type": "movies",
                "content_id": "127",
                "title": "Pacarku Baby Sitter",
                "synopsis": "Seorang Pemuda yang jatuh cinta kepada seorang baby sitter yang menjaga balita salah satu keluarganya",
                "duration": "100",
                "artist": "Larasati Nugroho, Fajar Gomez",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/127/cu_Movinesia_PacarkuBabySitter.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/127/cu_Movinesia_PacarkuBabySitter.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/127/cu_Movinesia_PacarkuBabySitter.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/pacarku-baby-sitter"
            },
            {
                "video_type": "movies",
                "content_id": "128",
                "title": "Tukang Ojek Cantik Dan Bumbu Cinta",
                "synopsis": "",
                "duration": "125",
                "artist": "Dallas Pratama Shamirah, Andez Fernandez, Mieke Shahir",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/128/cu_Movinesia_TukangOjekCantikDanBumbuCinta.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/128/cu_Movinesia_TukangOjekCantikDanBumbuCinta.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/128/cu_Movinesia_TukangOjekCantikDanBumbuCinta.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/tukang-ojek-cantik-dan-bumbu-cinta"
            },
            {
                "video_type": "movies",
                "content_id": "130",
                "title": "Ku Kejar Cinta Guru Killer",
                "synopsis": "Seorang pemuda yang jatuh cinta kepada guru lesnya yang ternyata jago bela diri.",
                "duration": "131",
                "artist": "Wafda Saifan, Metta Permadi",
                "genre": [
                    "Drama"
                ],
                "poster": {
                    "s": "https://assets.genflix.co.id/poster/movies/130/cu_Movinesia_KukejarCintaGuruKiller.300.jpg",
                    "m": "https://assets.genflix.co.id/poster/movies/130/cu_Movinesia_KukejarCintaGuruKiller.400.jpg",
                    "l": "https://assets.genflix.co.id/poster/movies/130/cu_Movinesia_KukejarCintaGuruKiller.1280.jpg"
                },
                "player": "http://dev.genflix.co.id/smartfren/movies/ku-kejar-cinta-guru-killer"
            }
        ]
    }`
       var s GenflixResp
    //    resp,err := MakeRequest()
    //    if err != nil {
    //        fmt.Println(err.Error())
    //        return
    //    }
       error := json.Unmarshal([]byte(jsonRes), &s)
        if(error != nil){
            fmt.Println("whoops:", error)
        }

        //fmt.Println(MakeVideoBulkMutationString(s))
        makeGraphQlRequest(queryToRequest(( strconv.QuoteToASCII(MakeVideoBulkMutationString(s))   )))
    }


func MakeVideoBulkMutationString (resp GenflixResp) string {
    fmt.Println("total data:", resp.TotalData)
    var ObjString string
    for i := 0; i < len(resp.Data); i++ {
        var video VideoBank
        video.Provider = "GENFLIX"
        video.ProviderShort = "GFX"
        video.Tags = resp.Data[i].Genre[0]
        video.ProviderLabel= resp.Data[i].Poster.S
        video.Title = resp.Data[i].Title
        video.TitlePackage= resp.Data[i].Title
        video.VideoType= resp.Data[i].VideoType
        genre,_:=json.Marshal(resp.Data[i].Genre)
        video.Genre= string(genre)
        video.Year= ""
        video.Duration = resp.Data[i].Duration
        video.ContentId = resp.Data[i].ContentID
        video.Synopsis = resp.Data[i].Synopsis
        video.Cast = ""
        video.PlayerUrl = resp.Data[i].Player
        video.S = resp.Data[i].Poster.S
        video.M = resp.Data[i].Poster.M
        video.L = resp.Data[i].Poster.L
        video.Director = "a"
        video.ContentType = "MOVIE"
        video.IsActive= "true"
        ObjString += buildGraphQlMutationString(video)
      // makeGraphQlRequest(queryToRequest(( strconv.QuoteToASCII(buildGraphQlMutationString(video))   )))
    }

    BulkMutationString := `mutation {
        CreateBulkVideoBank(
          input: [`+ObjString+`]
        ) {
           id
        }
      }
      `
    return BulkMutationString
}

func MakeRequest() ([]byte, error) {
    payload := url.Values{}
    payload.Set("offset","1")
    payload.Set("limit","1001")
    request, _ := http.NewRequest("POST", PROVIDER_URL, strings.NewReader(payload.Encode()) )
    request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Set("app-id", "smartfren")
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } 

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }
    
    return []byte(body), err
}

func queryToRequest(queryString string) string {
	return `{"query":` + queryString + `}`
}

func makeGraphQlRequest(requestString string) {
	fmt.Println("URL:", GRAPHQL_URL)
	req, err := http.NewRequest("POST", GRAPHQL_URL, strings.NewReader(requestString))
    req.Header.Set("Content-Type", "application/json")    
    //bypass ssl check
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Body:", string(body))
}






// http://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
///https://gist.github.com/rms1000watt/ba8db3137905b0848a4236e5f31125e3
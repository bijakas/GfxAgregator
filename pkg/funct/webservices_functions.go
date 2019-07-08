package funct

import (
    "fmt"
    "strconv"
	"encoding/json"
    "io/ioutil"
    "net/http"
    "net/url"
	"crypto/tls"
	"strings"
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

func BuildGraphQlMutationString(vb VideoBank) string{
    // videoType genre year duration contentType
    //convert cast to json array
    cast, _ := json.Marshal(strings.Split(vb.Cast, ","))
     duration := "0"
     if string(vb.Duration)!=""{duration=string(vb.Duration)}
    
    string := `
           {
            provider: "`+vb.Provider+`",
            providerShort: "`+vb.ProviderShort+`",
            tags: "`+vb.Tags+`",
            providerLabel: "`+vb.ProviderLabel+`",   
            title: "`+JsonEscape(vb.Title)+`",
            titlePackage: "`+JsonEscape(vb.TitlePackage)+`",
            videoType: "`+"MOVIES"+`",
            genre: ["HORROR", "FANTASY"], 
            year: 0,
            duration: `+duration+`, 
            contentId: `+vb.ContentId+`, 
            synopsis: "`+JsonEscape(vb.Synopsis)+`",
            cast: `+string(cast)+`,
            playerUrl: "`+JsonEscape(vb.PlayerUrl)+`",
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

func MakeVideoBulkMutationString (jsonRes string) string {
	var resp GenflixResp
    //    resp,err := MakeRequest()
    //    if err != nil {
    //        fmt.Println(err.Error())
    //        return
    //    }
    error := json.Unmarshal([]byte(jsonRes), &resp)
        if(error != nil){
            fmt.Println("whoops:", error)
	}
	
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
      //  fmt.Println("Duration :"+resp.Data[i].Duration)
        video.Duration = resp.Data[i].Duration
        video.ContentId = resp.Data[i].ContentID
        video.Synopsis = resp.Data[i].Synopsis
        video.Cast = resp.Data[i].Artist
        video.PlayerUrl = resp.Data[i].Player
        video.S = resp.Data[i].Poster.S
        video.M = resp.Data[i].Poster.M
        video.L = resp.Data[i].Poster.L
        video.Director = "a"
        video.ContentType = "MOVIE"
        video.IsActive= "true"
        ObjString += BuildGraphQlMutationString(video)
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
    //fmt.Println(BulkMutationString)
    return strconv.QuoteToASCII(BulkMutationString)
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

func QueryToRequest(queryString string) string {
	return `{"query":` + queryString + `}`
}

func MakeGraphQlRequest(requestString string) {

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
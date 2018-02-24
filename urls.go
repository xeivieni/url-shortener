package main

type Url struct {
	Id          int        `json:"id",db:"id"`
	ShortUrl 	string     `json:"shorturl",db:"shorturl"`
	LongUrl 	string     `json:"longurl",db:"longurl"`
	Hits 		int        `json:"hits",db:"hits"`
}

func UrlFromLong(longurl string) (*Url, error) {
	row := db.QueryRow(`SELECT * FROM urls WHERE longurl=$1`, longurl)

	url := new(Url)

	err := row.Scan(&url.Id, &url.ShortUrl, &url.LongUrl, &url.Hits)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func UrlFromShort(shorturl string) (*Url, error) {
	row := db.QueryRow(`SELECT * FROM urls WHERE shorturl=$1`, shorturl)

	url := new(Url)

	err := row.Scan(&url.Id, &url.ShortUrl, &url.LongUrl, &url.Hits)
	if err != nil {
		return nil, err
	}

	return url, nil
}


func NewUrl(url Url){
	db.QueryRow("INSERT INTO urls(longurl, shorturl, hits) VALUES($1, $2, $3)", &url.LongUrl, &url.ShortUrl, &url.Hits)
}

package crawler

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/eddiefisher/anime/internal/cache/memory"
)

var (
	Address    = []string{}
	errDecode  = errors.New("error decode response")
	errRequest = errors.New("error request")
)

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	City      string    `xml:"city"`
	Company   string    `xml:"company"`
	Logo      string    `xml:"logo"`
	JobType   string    `xml:"jobtype"`
	Category  string    `xml:"category"`
	PubDate   string    `xml:"pubDate"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

func Crawler(address string) ([]Item, error) {
	body, err := request(address)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	rss := Rss{}

	decoder := xml.NewDecoder(body)
	err = decoder.Decode(&rss)
	if err != nil {
		return nil, errDecode
	}

	return rss.Channel.Items, nil
}

func WebCrawler() ([]Item, error) {
	var items []Item
	for _, value := range Address {
		item, err := Crawler(value)
		items = append(items, item...)

		if err != nil {
			return nil, errDecode
		}
	}
	return items, nil
}

func request(address string) (io.ReadCloser, error) {
	storage := memory.NewStorage()
	content := storage.Get(address)
	if content != nil {
		return content, nil
	} else {
		resp, err := http.Get(address)
		if err != nil {
			return nil, errRequest
		}

		storage.Set(address, resp.Body, 5*time.Minute)

		return resp.Body, nil
	}
}

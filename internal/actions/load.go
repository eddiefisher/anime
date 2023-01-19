package actions

import (
	"regexp"
	"strings"
	"time"

	"github.com/eddiefisher/anime/internal/crawler"
	"github.com/eddiefisher/anime/internal/entity"
)

const avostURL = "https://v2.vost.pw/rss.xml"

func Load() ([]entity.Anime, error) {
	items, err := crawler.Crawler(avostURL)
	if err != nil {
		return nil, err
	}
	animes := []entity.Anime{}

	for _, item := range items {
		animeTitle, animeSeries := title(item.Title)
		animes = append(animes,
			entity.Anime{
				Source:         item.Link,
				Title:          animeTitle,
				Description:    item.Desc,
				Type:           0,
				CurrentEpisode: animeSeries,
				Episodes:       24,
				Status:         status(item.Category),
				AnimeSeason: entity.AnimeSeason{
					Season: "",
					Year:   0,
				},
				Picture:   "",
				Thumbnail: "",
				Synonyms:  []string{},
				Relations: []string{},
				Tags:      strings.Split(item.Category, ", "),
				Date:      humanTime(item.PubDate),
			},
		)
	}
	return animes, nil
}

func status(s string) entity.Status {
	if strings.Contains(s, "Анонсы") {

		return entity.Status{
			Name:  entity.Annonce,
			Color: entity.AnnonceColor,
		}
	}

	if strings.Contains(s, "Онгоинги") {
		return entity.Status{
			Name:  entity.Ongoing,
			Color: entity.OngoingColor,
		}
	}

	return entity.Status{
		Name:  entity.Finished,
		Color: entity.FinishedColor,
	}
}

func humanTime(s string) string {
	theTime, err := time.Parse(time.RFC1123Z, s)
	if err != nil {
		return s
	}

	return theTime.Format("2006/01/02")
}

func title(s string) (string, string) {
	s = strings.Replace(s, " [Анонс]", "", 1)

	re := regexp.MustCompile(`(\[.*?\])`)
	submatchall := re.FindAllString(s, -1)
	series := submatchall[len(submatchall)-1]

	s = strings.Replace(s, series, "", 1)

	return s, series
}

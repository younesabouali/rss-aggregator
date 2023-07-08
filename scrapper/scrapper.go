package scrapper

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/younesabouali/rss-aggregator/internal/database"
	"github.com/younesabouali/rss-aggregator/utils"
)

type FeedScrapper struct {
	DB *database.Queries
}

func (s *FeedScrapper) Scrape(concurent int, timeBetweenRequest time.Duration) {
	interval := time.Tick(timeBetweenRequest)
	for ; ; <-interval {
		fmt.Println("Scrapping")
		lastFetched, err := s.DB.ListFeedsToFetch(context.Background(), int32(concurent))
		if err != nil {
			fmt.Printf("Error Listings feeds To Fetch")
			continue
		}
		for _, feed := range lastFetched {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go s.scrapFeed(feed, &wg)
			wg.Wait()
		}
	}

}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}
func (s *FeedScrapper) scrapFeed(feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Fetching feed %v ", feed.Url)
	s.DB.UpdateFeedFetchData(context.Background(), database.UpdateFeedFetchDataParams{ID: feed.ID, LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}})
	posts, err := fetchFeed(feed.Url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, post := range posts.Channel.Item {
		defaultParams := utils.GetDefaultParams()
		PublishedAt := sql.NullTime{}
		t, err := time.Parse(time.RFC3339, post.PubDate)
		if err == nil {
			PublishedAt = sql.NullTime{
				Time: t, Valid: true,
			}
			// continue
		}
		_, err = s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          defaultParams.Id,
			CreatedAt:   defaultParams.CreatedAt,
			UpdatedAt:   defaultParams.UpdatedAt,
			Title:       post.Title,
			FeedID:      feed.ID,
			Url:         post.Title,
			Description: sql.NullString{String: post.Description, Valid: true},
			PublishedAt: PublishedAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

}

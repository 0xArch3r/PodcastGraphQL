package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"PodcastProject/feeds"
	"PodcastProject/graph/generated"
	"PodcastProject/graph/model"
	"PodcastProject/itunes"
	"PodcastProject/utils"
	"context"
)

func (r *queryResolver) Search(ctx context.Context, term string) ([]*model.Podcast, error) {
	ias := itunes.NewItunesApiServices()

	res, err := ias.Search(term)
	if err != nil {
		return nil, err
	}

	var podcasts []*model.Podcast

	for _, res := range res.Results {
		podcast := &model.Podcast{
			Artist:        res.ArtistName,
			PodcastName:   res.TrackName,
			FeedURL:       res.FeedURL,
			Thumbnail:     res.ArtworkURL100,
			EpisodesCount: res.TrackCount,
			Genres:        res.Genres,
		}

		podcasts = append(podcasts, podcast)
	}

	return podcasts, nil
}

func (r *queryResolver) Feed(ctx context.Context, feedURL string) ([]*model.FeedItem, error) {
	res, err := feeds.GetFeed(feedURL)
	if err != nil {
		return nil, err
	}

	var feed []*model.FeedItem

	for _, res := range res.Channel.Item {
		item := &model.FeedItem{
			PubDate:     res.PubDate,
			Text:        res.Text,
			Title:       res.Title,
			Subtitle:    res.Subtitle,
			Description: res.Description,
			Image:       utils.CheckNullString(res.Image.Href),
			Summary:     res.Summary,
			LinkURL:     res.Enclosure.URL,
			Duration:    res.Duration,
		}

		feed = append(feed, item)
	}

	return feed, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

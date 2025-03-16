package go_itunes_api

import "time"

type Results struct {

	// The number of results that match the search term.
	ResultCount int `json:"resultCount,omitempty"`

	// Results is an array of objects that contain the search results.
	Results []Result `json:"results,omitempty"`
}

// Result represents a single search result.
//
// https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/UnderstandingSearchResults.html#//apple_ref/doc/uid/TP40017632-CH8-SW1
type Result struct {
	ArtistId                int       `json:"artistId,omitempty,omitempty"`
	ArtistName              string    `json:"artistName,omitempty"`
	ArtistViewUrl           string    `json:"artistViewUrl,omitempty,omitempty"`
	ArtworkUrl100           string    `json:"artworkUrl100,omitempty"`
	ArtworkUrl30            string    `json:"artworkUrl30,omitempty"`
	ArtworkUrl60            string    `json:"artworkUrl60,omitempty"`
	CollectionArtistId      int       `json:"collectionArtistId,omitempty,omitempty"`
	CollectionArtistName    string    `json:"collectionArtistName,omitempty,omitempty"`
	CollectionArtistViewUrl string    `json:"collectionArtistViewUrl,omitempty,omitempty"`
	CollectionCensoredName  string    `json:"collectionCensoredName,omitempty,omitempty"`
	CollectionExplicitness  string    `json:"collectionExplicitness,omitempty"`
	CollectionHdPrice       float64   `json:"collectionHdPrice,omitempty,omitempty"`
	CollectionId            int       `json:"collectionId,omitempty,omitempty"`
	CollectionName          string    `json:"collectionName,omitempty,omitempty"`
	CollectionPrice         float64   `json:"collectionPrice,omitempty"`
	CollectionViewUrl       string    `json:"collectionViewUrl,omitempty,omitempty"`
	ContentAdvisoryRating   string    `json:"contentAdvisoryRating,omitempty,omitempty"`
	Country                 string    `json:"country,omitempty"`
	Currency                string    `json:"currency,omitempty"`
	DiscCount               int       `json:"discCount,omitempty,omitempty"`
	DiscNumber              int       `json:"discNumber,omitempty,omitempty"`
	HasITunesExtras         bool      `json:"hasITunesExtras,omitempty,omitempty"`
	IsStreamable            bool      `json:"isStreamable,omitempty,omitempty"`
	Kind                    Kind      `json:"kind,omitempty"`
	LongDescription         string    `json:"longDescription,omitempty,omitempty"`
	PreviewUrl              string    `json:"previewUrl,omitempty"`
	PrimaryGenreName        string    `json:"primaryGenreName,omitempty"`
	ReleaseDate             time.Time `json:"releaseDate,omitempty"`
	ShortDescription        string    `json:"shortDescription,omitempty,omitempty"`
	TrackCensoredName       string    `json:"trackCensoredName,omitempty"`
	TrackCount              int       `json:"trackCount,omitempty,omitempty"`
	TrackExplicitness       string    `json:"trackExplicitness,omitempty"`
	TrackHdPrice            float64   `json:"trackHdPrice,omitempty,omitempty"`
	TrackHdRentalPrice      float64   `json:"trackHdRentalPrice,omitempty,omitempty"`
	TrackId                 int       `json:"trackId,omitempty"`
	TrackName               string    `json:"trackName,omitempty"`
	TrackNumber             int       `json:"trackNumber,omitempty,omitempty"`
	TrackPrice              float64   `json:"trackPrice,omitempty"`
	TrackRentalPrice        float64   `json:"trackRentalPrice,omitempty,omitempty"`
	TrackTimeMillis         int       `json:"trackTimeMillis,omitempty"`
	TrackViewUrl            string    `json:"trackViewUrl,omitempty"`
	WrapperType             string    `json:"wrapperType,omitempty"`
}

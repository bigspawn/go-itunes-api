package go_itunes_api

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
	ArtistID               int     `json:"artistId"`
	ArtistName             string  `json:"artistName"`
	ArtistViewURL          string  `json:"artistViewUrl"`
	ArtworkURL100          string  `json:"artworkUrl100"`
	ArtworkURL60           string  `json:"artworkUrl60"`
	CollectionCensoredName string  `json:"collectionCensoredName"`
	CollectionExplicitness string  `json:"collectionExplicitness"`
	CollectionID           int     `json:"collectionId"`
	CollectionName         string  `json:"collectionName"`
	CollectionPrice        float64 `json:"collectionPrice"`
	CollectionViewURL      string  `json:"collectionViewUrl"`
	Country                string  `json:"country"`
	Currency               string  `json:"currency"`
	DiscCount              int     `json:"discCount"`
	DiscNumber             int     `json:"discNumber"`
	Kind                   Kind    `json:"kind"`
	PreviewURL             string  `json:"previewUrl"`
	PrimaryGenreName       string  `json:"primaryGenreName"`
	TrackCensoredName      string  `json:"trackCensoredName"`
	TrackCount             int     `json:"trackCount"`
	TrackExplicitness      string  `json:"trackExplicitness"`
	TrackID                int     `json:"trackId"`
	TrackName              string  `json:"trackName"`
	TrackNumber            int     `json:"trackNumber"`
	TrackPrice             float64 `json:"trackPrice"`
	TrackTimeMillis        int     `json:"trackTimeMillis"`
	TrackViewURL           string  `json:"trackViewUrl"`
	WrapperType            string  `json:"wrapperType"`
}

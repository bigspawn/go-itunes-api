package go_itunes_api

import "time"

type Results struct {

	// The number of results that match the search term.
	ResultCount int `json:"resultCount,omitempty"`

	// Results is an array of objects that contain the search results.
	Results []Result `json:"results,omitempty"`
}

type Result struct {

	// The name of the object returned by the search request.
	WrapperType string `json:"wrapperType,omitempty"`

	// The kind of content returned by the search request.
	Kind Kind `json:"kind,omitempty"`

	//
	ArtistId int `json:"artistId,omitempty,omitempty"`

	//
	CollectionId int `json:"collectionId,omitempty,omitempty"`

	//
	TrackId int `json:"trackId,omitempty"`

	// The name of the artist returned by the search request.
	ArtistName string `json:"artistName,omitempty"`

	// The name of the album, TV season, audiobook, and so on returned by the search
	// request.
	CollectionName string `json:"collectionName,omitempty,omitempty"`

	// The name of the track, song, video, TV episode, and so on returned by the
	// search request.
	TrackName string `json:"trackName,omitempty"`

	//
	CollectionCensoredName string `json:"collectionCensoredName,omitempty,omitempty"`

	//
	TrackCensoredName string `json:"trackCensoredName,omitempty"`

	//
	ArtistViewUrl string `json:"artistViewUrl,omitempty,omitempty"`

	//
	CollectionViewUrl string `json:"collectionViewUrl,omitempty,omitempty"`

	//
	TrackViewUrl string `json:"trackViewUrl,omitempty"`

	//
	PreviewUrl string `json:"previewUrl,omitempty"`

	//
	ArtworkUrl30 string `json:"artworkUrl30,omitempty"`

	//
	ArtworkUrl60 string `json:"artworkUrl60,omitempty"`

	//
	ArtworkUrl100 string `json:"artworkUrl100,omitempty"`

	//
	CollectionPrice float64 `json:"collectionPrice,omitempty"`

	//
	TrackPrice float64 `json:"trackPrice,omitempty"`

	//
	ReleaseDate time.Time `json:"releaseDate,omitempty"`

	//
	CollectionExplicitness string `json:"collectionExplicitness,omitempty"`

	//
	TrackExplicitness string `json:"trackExplicitness,omitempty"`

	//
	DiscCount int `json:"discCount,omitempty,omitempty"`

	//
	DiscNumber int `json:"discNumber,omitempty,omitempty"`

	//
	TrackCount int `json:"trackCount,omitempty,omitempty"`

	//
	TrackNumber int `json:"trackNumber,omitempty,omitempty"`

	//
	TrackTimeMillis int `json:"trackTimeMillis,omitempty"`

	//
	Country string `json:"country,omitempty"`

	//
	Currency string `json:"currency,omitempty"`

	//
	PrimaryGenreName string `json:"primaryGenreName,omitempty"`

	//
	IsStreamable bool `json:"isStreamable,omitempty,omitempty"`

	//
	CollectionArtistId int `json:"collectionArtistId,omitempty,omitempty"`

	//
	CollectionArtistViewUrl string `json:"collectionArtistViewUrl,omitempty,omitempty"`

	//
	TrackRentalPrice float64 `json:"trackRentalPrice,omitempty,omitempty"`

	//
	CollectionHdPrice float64 `json:"collectionHdPrice,omitempty,omitempty"`

	//
	TrackHdPrice float64 `json:"trackHdPrice,omitempty,omitempty"`

	//
	TrackHdRentalPrice float64 `json:"trackHdRentalPrice,omitempty,omitempty"`

	//
	ContentAdvisoryRating string `json:"contentAdvisoryRating,omitempty,omitempty"`

	//
	ShortDescription string `json:"shortDescription,omitempty,omitempty"`

	//
	LongDescription string `json:"longDescription,omitempty,omitempty"`

	//
	HasITunesExtras bool `json:"hasITunesExtras,omitempty,omitempty"`

	//
	CollectionArtistName string `json:"collectionArtistName,omitempty,omitempty"`
}

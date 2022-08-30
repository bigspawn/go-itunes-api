package go_itunes_api

type MediaType string

const (
	MediaTypeMovie      MediaType = "movie"
	MediaTypePodcast    MediaType = "podcast"
	MediaTypeMusic      MediaType = "music"
	MediaTypeMusicVideo MediaType = "musicVideo"
	MediaTypeAudiobook  MediaType = "audiobook"
	MediaTypeShortFilm  MediaType = "shortFilm"
	MediaTypeTvShow     MediaType = "tvShow"
	MediaTypeSoftware   MediaType = "software"
	MediaTypeEbook      MediaType = "ebook"
	MediaTypeAll        MediaType = "all"
)

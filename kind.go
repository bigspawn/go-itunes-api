package go_itunes_api

type Kind string

const (
	KindBook               Kind = "book"
	KindAlbum              Kind = "album"
	KindCoachedAudio       Kind = "coached-audio"
	KindFeatureMovie       Kind = "feature-movie"
	KindInteractiveBooklet Kind = "interactive- booklet"
	KindMusicVideo         Kind = "music-video"
	KindPdf                Kind = "pdf"
	KindPodcast            Kind = "podcast"
	KindPodcastEpisode     Kind = "podcast-episode"
	KindSoftwarePackage    Kind = "software-package"
	KindSong               Kind = "song"
	KindTvEpisode          Kind = "tv- episode"
	KindArtist             Kind = "artist"
)

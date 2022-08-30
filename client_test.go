package go_itunes_api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const testResp = `{"resultCount":1,"results":[{"wrapperType":"track","kind":"song","artistId":909253,"collectionId":1469577723,"trackId":1469577741,"artistName":"Jack Johnson","collectionName":"Jack Johnson and Friends: Sing-A-Longs and Lullabies for the Film Curious George","trackName":"Upside Down","collectionCensoredName":"Jack Johnson and Friends: Sing-A-Longs and Lullabies for the Film Curious George","trackCensoredName":"Upside Down","artistViewUrl":"https://music.apple.com/us/artist/jack-johnson/909253?uo=4","collectionViewUrl":"https://music.apple.com/us/album/upside-down/1469577723?i=1469577741&uo=4","trackViewUrl":"https://music.apple.com/us/album/upside-down/1469577723?i=1469577741&uo=4","previewUrl":"https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview125/v4/5e/5b/3d/5e5b3df4-deb5-da78-5d64-fe51d8404d5c/mzaf_13341178261601361485.plus.aac.p.m4a","artworkUrl30":"https://is3-ssl.mzstatic.com/image/thumb/Music115/v4/08/11/d2/0811d2b3-b4d5-dc22-1107-3625511844b5/00602537869770.rgb.jpg/30x30bb.jpg","artworkUrl60":"https://is3-ssl.mzstatic.com/image/thumb/Music115/v4/08/11/d2/0811d2b3-b4d5-dc22-1107-3625511844b5/00602537869770.rgb.jpg/60x60bb.jpg","artworkUrl100":"https://is3-ssl.mzstatic.com/image/thumb/Music115/v4/08/11/d2/0811d2b3-b4d5-dc22-1107-3625511844b5/00602537869770.rgb.jpg/100x100bb.jpg","collectionPrice":9.99,"trackPrice":1.29,"releaseDate":"2005-01-01T12:00:00Z","collectionExplicitness":"notExplicit","trackExplicitness":"notExplicit","discCount":1,"discNumber":1,"trackCount":14,"trackNumber":1,"trackTimeMillis":208643,"country":"USA","currency":"USD","primaryGenreName":"Rock","isStreamable":true}]}`

func TestClient_Search_Request(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		t.Logf("request: %s %s", request.Method, request.URL)
		_, err := writer.Write([]byte(testResp))
		require.NoError(t, err)
	}))
	defer s.Close()

	c, err := NewClient(ClientOption{Host: s.URL})
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := c.Search(ctx, SearchRequest{
		Term:    "Jack Johnson",
		Country: US,
		Limit:   1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Results)
	require.Equal(t, 1, resp.Results.ResultCount)
	require.Len(t, resp.Results.Results, 1)
}

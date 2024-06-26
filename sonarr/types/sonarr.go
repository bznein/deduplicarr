package types

type Series struct {
	ID               int              `json:"id"`
	Title            string           `json:"title"`
	TitleSlug        string           `json:"titleSlug"`
	Path             string           `json:"path"`
	TvdbID           int              `json:"tvdbId"`
	TvMazeID         int              `json:"tvMazeId"`
	ImdbID           string           `json:"imdbId"`
	Type             string           `json:"type"`
	Year             int              `json:"year"`
	Tags             []string         `json:"tags"`
	Episodes         []Episode        `json:"episodes"`
	EpisodeFile      EpisodeFile      `json:"episodeFile"`
	IsUpgrade        bool             `json:"isUpgrade"`
	CustomFormatInfo CustomFormatInfo `json:"customFormatInfo"`
	Release          Release          `json:"release"`
	EventType        string           `json:"eventType"`
	InstanceName     string           `json:"instanceName"`
	ApplicationUrl   string           `json:"applicationUrl"`
}

type Episode struct {
	ID            int    `json:"id"`
	EpisodeNumber int    `json:"episodeNumber"`
	SeasonNumber  int    `json:"seasonNumber"`
	Title         string `json:"title"`
	Overview      string `json:"overview"`
	AirDate       string `json:"airDate"`
	AirDateUTC    string `json:"airDateUtc"`
	SeriesID      int    `json:"seriesId"`
	TvdbID        int    `json:"tvdbIdd"`
}

type EpisodeFile struct {
	ID             int       `json:"id"`
	RelativePath   string    `json:"relativePath"`
	Path           string    `json:"path"`
	Quality        string    `json:"quality"`
	QualityVersion int       `json:"qualityVersion"`
	ReleaseGroup   string    `json:"releaseGroup"`
	Scenename      string    `json:"scenename"`
	Size           int       `json:"size"`
	DateAdded      string    `json:"dateAdded"`
	MediaInfo      MediaInfo `json:"mediaInfo"`
}

type MediaInfo struct {
	AudioChannels         float32  `json:"audioChannels"`
	AudioCodec            string   `json:"audioCodec"`
	AudioLanguages        []string `json:"audioLanguages"`
	Height                int      `json:"height"`
	Width                 int      `json:"width"`
	Subtitles             []string `json:"subtitles"`
	VideoCodec            string   `json:"videoCodec"`
	VideoDynamicRange     string   `json:"videoDynamicRange"`
	VideoDynamicRangeType string   `json:"videoDynamicRangeType"`
}

type CustomFormatInfo struct {
	CustomFormats     []CustomFormat `json:"customFormats"`
	CustomFormatScore int            `json:"customFormatScore"`
}

type CustomFormat struct {
	ID   int    `json:"id"`
	name string `json:"name"`
}

type Release struct {
	Size int `json:"size"`
}

package metagenerator

type Meta struct {
	Title            string   `json:"title" `
	Description      string   `json:"description"`
	Genres           []string `json:"genres" `
	MetadataLanguage string   `json:"metadataLanguage" `
	Language         string   `json:"language"`
	ReleaseYear      int      `json:"releaseYear" `
	Format           Format   `json:"format"`
	DurationSeconds  int      `json:"durationSeconds" `
	Series           struct {
		Cast []string `json:"cast"`
	} `json:"series"`
	Href             string `json:"href"`
	Extract          string `json:"extract"`
	Thumbnail        string `json:"thumbnail"`
	ThumbnailWidth   int    `json:"thumbnailWidth"`
	Thumbnail_Height int    `json:"thumbnail_Height"`
}

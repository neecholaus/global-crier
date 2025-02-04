package headlines

type Publication string

const (
	Reuters Publication = "Reuters"
)

// Reuters url has a param "d". I do not know what it indicates.
// At the start the value was 240. But now requests with that fail.
//
// Here is a list showing what values I have used in order:
// -- 240, 258

var Sources = []*Source{
	{
		Publication: string(Reuters),
		Name:        "Video List",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/video-playlist-by-slug-v1?query=%7B%22slug%22%3A%22%2Fvideo%2Fhp-us-2024-01-04%2F%22%2C%22website%22%3A%22reuters%22%7D&d=258&_website=reuter",
		Parser:      reutersVideoParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 1",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs1\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 2",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs2\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 3",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs3\",\"size\":10,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 4",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs4\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 5",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs5\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 6",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs6\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 7",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs7\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 8",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs8\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 9",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs9\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 10",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs10\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 13",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs13\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
	{
		Publication: string(Reuters),
		Name:        "Big Story 14",
		URL:         "https://www.reuters.com/pf/api/v3/content/fetch/articles-by-collection-alias-or-id-v1?query={\"collection_alias\":\"bs14\",\"size\":20,\"website\":\"reuters\"}&d=258&_website=reuters",
		Parser:      reutersBigStoryParser{},
	},
}

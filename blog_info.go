package tumblgo

type BlogInfoResponse struct {
	Meta     Meta     `json:"meta"`
	Response BlogInfo `json:"response"`
}

type BlogInfo struct {
	Blog Blog `json:"blog"`
}

type Blog struct {
	Title                string `json:"title"`
	Posts                int    `json:"posts"`
	Name                 string `json:"name"`
	Updated              int    `json:"updated"`
	Description          string `json:"description"`
	Ask                  bool   `json:"ask"`
	AskAnon              bool   `json:"ask_anon"`
	Likes                int    `json:"likes"`
	IsBlockedFromPrimary bool   `json:"is_blocked_from_primary"`
}

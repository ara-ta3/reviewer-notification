package github

type Label struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Default bool   `json:"default"`
}

type assignee struct {
	Login string `json:"login"`
}

package structs

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

// AlfredOutput alfred workflow output
type Alfred struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Icon     Icon   `json:"icon"`
	Valid    bool   `json:"valid"`
	Text     struct {
		Copy string `json:"copy"`
	} `json:"text"`
}

type AlfredItems struct {
	Items []Alfred `json:"items"`
}

func (ts AlfredItems) String(i int) string {
	return ts.Items[i].Title
}

func (ts AlfredItems) Len() int {
	return len(ts.Items)
}

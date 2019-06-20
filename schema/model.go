package schema

import "time"

type Album struct {
    ID      int       `json:"id"`
    Title   string    `json:"name"`
    Artist    string    `json:"artist"`
    DateAdded time.Time `json:"dateadded"`
}
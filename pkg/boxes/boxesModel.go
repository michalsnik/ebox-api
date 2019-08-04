package boxes

type BoxItem struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
}

type Box struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Items []BoxItem `json:"items"`
}

package types

type Distance struct {
	OBUid int64   `json:"obu_id"`
	Value float64 `json:"value"`
	Unix  int64   `json:"unix"`
}

type Invoice struct {
	OBUid    int64   `json:"obu_id"`
	Distance float64 `json:"distance"`
	Amount   float64 `json:"amount"`
}

type OBUdata struct {
	OBUid int64   `json:"obu_id"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

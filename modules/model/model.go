package model

type (
	KafKaOffset struct {
		Offset int64 `json:"offset" bson:"offset"`
	}
)

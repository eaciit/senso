package senso

type DocumentMapping struct {
	ID           string `bson:_id, json:_id`
	Recognitions []RecognitionMap
	Fields       []FieldMap
}

type RecognitionMap struct {
	Coordinate
	ExpectedValue string
}

type FieldMap struct {
	IsArray         bool
	Start           []RecognitionMap
	End             []RecognitionMap
	SingleMap       SingleMap
	ListHeaderField string
	ListMap         []SingleMap
}

type SingleMap struct {
	Coordinate
	Field string
}

type Coordinate struct {
	OffsetX, OffsetY, X, Y int
}

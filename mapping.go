package senso

type DocumentMapping struct {
	ID           string `bson:_id, json:_id`
	Recognitions []RecognitionMap
	Fields       []FieldMap
}

type RecognitionMap struct {
}

type FieldMap struct {
	IsArray   bool
	Start     []RecognitionMap
	End       []RecognitionMap
	SingleMap SingleMap
	ListMap   []SingleMap
}

type SingleMap struct {
}

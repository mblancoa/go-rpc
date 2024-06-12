package domain

type InfoFile struct {
	Type, Version, Hash string
	Content             map[string]interface{}
}

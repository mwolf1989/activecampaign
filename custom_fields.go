package activecampaign

import (
	"context"
	"encoding/json"
	"net/http"
)

type Fields struct {
	FieldOptions   interface{}     `json:"fieldOptions"`
	FieldRelations []FieldRelation `json:"fieldRels"`
	Fields         []Field         `json:"fields"`
	Meta           FieldsMeta      `json:"meta"`
}

type FieldsMeta struct {
	Total string `json:"total"`
}

type FieldRelation struct {
	Field      string `json:"field"`
	RelationID string `json:"relid"`
	DOrder     string `json:"dorder"`
	CreateDate string `json:"cdate"`
	//Links      interface{} `json:"links"`
	ID string `json:"id"`
}

type Field struct {
	Title        string `json:"title"`
	Description  string `json:"descript"`
	IsRequired   string `json:"isrequired"`
	Perstag      string `json:"perstag"`
	DefaultValue string `json:"defval"`
	Visible      string `json:"visible"`
	Service      string `json:"service"`
	Ordernum     string `json:"ordernum"`
	CreateDate   string `json:"cdate"`
	UpdateDate   string `json:"udate"`
	//Options      interface{}    `json:"options"`
	Relations []string  `json:"relations"`
	Links     FieldLink `json:"links"`
	ID        string    `json:"id"`
}

type FieldLink struct {
	Options   string `json:"options"`
	Relations string `json:"Relations"`
}

func (a *ActiveCampaign) Fields(ctx context.Context, pof *POF) (*Fields, error) {
	res, err := a.send(ctx, http.MethodGet, "fields", pof, nil)
	if err != nil {
		return nil, &Error{Op: "fields", Err: err}
	}

	var fields Fields
	err = json.NewDecoder(res.Body).Decode(&fields)
	if err != nil {
		return nil, &Error{Op: "fields", Err: err}
	}

	return &fields, nil
}

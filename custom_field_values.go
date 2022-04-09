package activecampaign

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FieldValues struct {
	FieldValues []FieldValue    `json:"fieldValues"`
	Meta        FieldValuesMeta `json:"meta"`
}
type FieldValuesMeta struct {
	Total string `json:"total"`
}
type FieldValue struct {
	Contact    string `json:"contact"`
	Field      string `json:"field"`
	Value      string `json:"value"`
	CreateDate string `json:"cdate"`
	UpdateDate string `json:"udate"`
	ID         string `json:"id"`
}

func (a *ActiveCampaign) fieldValues(ctx context.Context, pof *POF, url string) (*FieldValues, error) {
	res, err := a.send(ctx, http.MethodGet, url, pof, nil)
	if err != nil {
		return nil, &Error{Op: "field values", Err: err}
	}
	defer res.Body.Close()

	var values FieldValues
	err = json.NewDecoder(res.Body).Decode(&values)
	if err != nil {
		return nil, &Error{Op: "field values", Err: err}
	}

	return &values, nil
}

func (a *ActiveCampaign) FieldValues(ctx context.Context, pof *POF) (*FieldValues, error) {
	return a.fieldValues(ctx, pof, "fieldValues")
}

type ChangeFieldValue struct {
	Contact string `json:"contact"`
	Field   string `json:"field"`
	Value   string `json:"value"`
}

func (a *ActiveCampaign) FieldValueCreate(ctx context.Context, create ChangeFieldValue) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(struct {
		FieldValue ChangeFieldValue `json:"fieldValue"`
	}{
		FieldValue: create,
	})
	if err != nil {
		return &Error{Op: "field value create", Err: err}
	}

	res, err := a.send(ctx, http.MethodPost, "fieldValues", nil, b)
	if err != nil {
		return &Error{Op: "field value create", Err: err}
	}
	defer res.Body.Close()
	if !(res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK) {
		return errors.New("field value create: " + res.Status)
	}

	return nil
}

func (a *ActiveCampaign) FieldValueUpdate(ctx context.Context, id string, update ChangeFieldValue) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(struct {
		FieldValue ChangeFieldValue `json:"fieldValue"`
	}{
		FieldValue: update,
	})
	if err != nil {
		return &Error{Op: "field value update", Err: err}
	}

	res, err := a.send(ctx, http.MethodPut, "fieldValues/"+id, nil, b)
	if err != nil {
		return &Error{Op: "field value update", Err: err}
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return &Error{Op: "field value update", Err: err}
		}
		return errors.New("field value update: " + res.Status + ": " + string(b))
	}

	return nil
}

type FieldOptions struct {
	FieldOptions []FieldOption `json:"fieldOptions"`
}
type FieldOption struct {
	Field      string `json:"field"`
	OrderID    string `json:"orderid"`
	Value      string `json:"value"`
	Label      string `json:"label"`
	IsDefault  string `json:"isdefault"`
	CreateDate string `json:"cdate"`
	UpdateDate string `json:"udate"`
	ID         string `json:"id"`
}

func (a *ActiveCampaign) FieldOptions(ctx context.Context, field string) ([]FieldOption, error) {
	res, err := a.send(ctx, http.MethodGet, fmt.Sprintf("fields/%s/options", field), nil, nil)
	if err != nil {
		return nil, &Error{Op: "field options", Err: err}
	}
	defer res.Body.Close()

	var values FieldOptions
	err = json.NewDecoder(res.Body).Decode(&values)
	if err != nil {
		return nil, &Error{Op: "field options", Err: err}
	}

	return values.FieldOptions, nil
}

type CreateFieldOption struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Label string `json:"label"`
}

func (a *ActiveCampaign) FieldOptionCreate(ctx context.Context, create []CreateFieldOption) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(struct {
		FieldOptions []CreateFieldOption `json:"fieldOptions"`
	}{
		FieldOptions: create,
	})
	if err != nil {
		return &Error{Op: "field options create", Err: err}
	}

	res, err := a.send(ctx, http.MethodPost, "fieldOption/bulk", nil, b)
	if err != nil {
		return &Error{Op: "field options create", Err: err}
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return errors.New("field options create: " + res.Status)
	}

	return nil
}

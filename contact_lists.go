package activecampaign

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type ContactUpdatedToList struct {
	Contacts []Contact `json:"contacts"`
}

type ContactLists struct {
	ContactLists []ContactList `json:"contactLists"`
}

type ContactList struct {
	Contact               string          `json:"contact"`
	List                  string          `json:"list"`
	Form                  *string         `json:"form"`
	Seriesid              string          `json:"seriesid"`
	Sdate                 string          `json:"sdate"`
	Status                string          `json:"status"`
	Responder             string          `json:"responder"`
	Sync                  string          `json:"sync"`
	Unsubreason           string          `json:"unsubreason"`
	Campaign              *string         `json:"campaign"`
	Message               *string         `json:"message"`
	First_name            string          `json:"first_name"`
	Last_name             string          `json:"last_name"`
	Ip4Sub                string          `json:"ip4Sub"`
	Sourceid              string          `json:"sourceid"`
	AutosyncLog           *string         `json:"autosyncLog"`
	Ip4_last              string          `json:"ip4_last"`
	Ip4Unsub              string          `json:"ip4Unsub"`
	UnsubscribeAutomation *string         `json:"unsubscribeAutomation"`
	Links                 ContactListLink `json:"links"`
	ID                    string          `json:"id"`
	Automation            *string         `json:"automation"`
}

type ContactListLink struct {
	Automation            string `json:"automation"`
	List                  string `json:"list"`
	Contact               string `json:"contact"`
	Form                  string `json:"form"`
	AutosyncLog           string `json:"autosyncLog"`
	Campaign              string `json:"campaign"`
	UnsubscribeAutomation string `json:"unsubscribeAutomation"`
	Message               string `json:"message"`
}

type updateContactToListRequest struct {
	List    int        `json:"list"`
	Contact int        `json:"contact"`
	Status  ListChange `json:"status"`
}

type ListChange int

const (
	ListSubscribe   ListChange = 1
	ListUnsubscribe ListChange = 2
)

func (a *ActiveCampaign) UpdateContactToList(ctx context.Context, contactID string, listID string, listChange ListChange) (*ContactUpdatedToList, error) {
	b := new(bytes.Buffer)

	cID, err := strconv.Atoi(contactID)
	if err != nil {
		return nil, &Error{Op: "converting contactID to int", Err: err}
	}
	lID, err := strconv.Atoi(listID)
	if err != nil {
		return nil, &Error{Op: "converting listID to int", Err: err}
	}

	err = json.NewEncoder(b).Encode(struct {
		ContactList updateContactToListRequest `json:"contactList"`
	}{
		ContactList: updateContactToListRequest{
			List:    lID,
			Contact: cID,
			Status:  listChange,
		},
	})
	if err != nil {
		return nil, &Error{Op: "update contact to list", Err: err}
	}

	res, err := a.send(ctx, http.MethodPost, "contactLists", nil, b)
	if err != nil {
		return nil, &Error{Op: "update contact to list", Err: err}
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		return nil, errors.New("update contact to list: " + res.Status)
	}

	var contactUpdatedToList ContactUpdatedToList
	err = json.NewDecoder(res.Body).Decode(&contactUpdatedToList)
	if err != nil {
		return nil, err
	}

	return &contactUpdatedToList, nil
}

func (a *ActiveCampaign) ContactLists(ctx context.Context, contactID string) (*ContactLists, error) {
	res, err := a.send(ctx, http.MethodGet, fmt.Sprintf("contacts/%s/contactLists", contactID), nil, nil)
	if err != nil {
		return nil, &Error{Op: "contactLists", Err: err}
	}
	defer res.Body.Close()

	var contactLists ContactLists
	err = json.NewDecoder(res.Body).Decode(&contactLists)
	if err != nil {
		return nil, &Error{Op: "contactLists", Err: err}
	}

	return &contactLists, nil
}

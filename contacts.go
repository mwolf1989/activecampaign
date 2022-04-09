package activecampaign

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Contacts struct {
	Contacts []Contact       `json:"contacts"`
	Meta     FieldValuesMeta `json:"meta"`
}

type Contact struct {
	CreateDate string       `json:"cdate"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	FirstName  string       `json:"firstName,omitempty"`
	LastName   string       `json:"lastName,omitempty"`
	ID         string       `json:"id"`
	UpdateDate string       `json:"udate"`
	Links      ContactLinks `json:"links"`
}

type ContactLinks struct {
	BounceLogs            string `json:"bounceLogs"`
	ContactAutomations    string `json:"contactAutomations"`
	ContactData           string `json:"contactData"`
	ContactGoals          string `json:"contactGoals"`
	ContactLists          string `json:"contactLists"`
	ContactLogs           string `json:"contactLogs"`
	ContactTags           string `json:"contactTags"`
	ContactDeals          string `json:"contactDeals"`
	Deals                 string `json:"deals"`
	FieldValues           string `json:"fieldValues"`
	GeoIps                string `json:"geoIps"`
	Notes                 string `json:"notes"`
	Organization          string `json:"organization"`
	PlusAppend            string `json:"plusAppend"`
	TrackingLogs          string `json:"trackingLogs"`
	ScoreValues           string `json:"scoreValues"`
	AccountContacts       string `json:"accountContacts"`
	AutomationEntryCounts string `json:"automationEntryCounts"`
}

func (a *ActiveCampaign) Contacts(ctx context.Context, pof *POF) (*Contacts, error) {
	res, err := a.send(ctx, http.MethodGet, "contacts", pof, nil)
	if err != nil {
		return nil, &Error{Op: "contacts", Err: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var contacts Contacts
	err = json.NewDecoder(res.Body).Decode(&contacts)
	if err != nil {
		return nil, &Error{Op: "contacts", Err: err}
	}

	return &contacts, nil
}

func (a *ActiveCampaign) ListContacts(ctx context.Context, listID string) (*Contacts, error) {
	res, err := a.send(ctx, http.MethodGet, "contacts?listid="+listID+"&status=1", nil, nil)
	if err != nil {
		return nil, &Error{Op: "list contacts", Err: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var contacts Contacts
	err = json.NewDecoder(res.Body).Decode(&contacts)
	if err != nil {
		return nil, &Error{Op: "list contacts", Err: err}
	}

	return &contacts, nil
}

func (a *ActiveCampaign) ContactFieldValues(ctx context.Context, pof *POF, id string) (*FieldValues, error) {
	return a.fieldValues(ctx, pof, "contacts/"+id+"/fieldValues")
}

type ContactCreate struct {
	Contact Contact `json:"contact"`
}

type ContactCreated struct {
	Contact struct {
		Email      string `json:"email"`
		FirstName  string `json:"firstName"`
		LastName   string `json:"lastName"`
		EmailEmpty bool   `json:"email_empty"`
		Cdate      string `json:"cdate"`
		Udate      string `json:"udate"`
		Orgid      string `json:"orgid"`
		Orgname    string `json:"orgname"`
		Links      struct {
			BounceLogs            string `json:"bounceLogs"`
			ContactAutomations    string `json:"contactAutomations"`
			ContactData           string `json:"contactData"`
			ContactGoals          string `json:"contactGoals"`
			ContactLists          string `json:"contactLists"`
			ContactLogs           string `json:"contactLogs"`
			ContactTags           string `json:"contactTags"`
			ContactDeals          string `json:"contactDeals"`
			Deals                 string `json:"deals"`
			FieldValues           string `json:"fieldValues"`
			GeoIps                string `json:"geoIps"`
			Notes                 string `json:"notes"`
			Organization          string `json:"organization"`
			PlusAppend            string `json:"plusAppend"`
			TrackingLogs          string `json:"trackingLogs"`
			ScoreValues           string `json:"scoreValues"`
			AccountContacts       string `json:"accountContacts"`
			AutomationEntryCounts string `json:"automationEntryCounts"`
		} `json:"links"`
		Hash         string `json:"hash"`
		ID           string `json:"id"`
		Organization string `json:"organization"`
	} `json:"contact"`
}

type Links struct {
	BounceLogs         string `json:"bounceLogs"`
	ContactAutomations string `json:"contactAutomations"`
	ContactData        string `json:"contactData"`
	ContactGoals       string `json:"contactGoals"`
	ContactLists       string `json:"contactLists"`
	ContactLogs        string `json:"contactLogs"`
	ContactTags        string `json:"contactTags"`
	ContactDeals       string `json:"contactDeals"`
	Deals              string `json:"deals"`
	FieldValues        string `json:"fieldValues"`
	GeoIps             string `json:"geoIps"`
	Notes              string `json:"notes"`
	Organization       string `json:"organization"`
	PlusAppend         string `json:"plusAppend"`
	TrackingLogs       string `json:"trackingLogs"`
	ScoreValues        string `json:"scoreValues"`
}

func (a *ActiveCampaign) ContactCreate(ctx context.Context, contact ContactCreate) (*ContactCreated, error) {

	//Decode the Struct into JSON
	jsonStr, err := json.Marshal(contact)
	if err != nil {
		return nil, &Error{Op: "contact create", Err: err}
	}

	res, err := a.send(ctx, http.MethodPost, "contacts", nil, bytes.NewReader(jsonStr))
	if err != nil {
		return nil, &Error{Op: "contact create", Err: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.StatusCode != http.StatusCreated {
		return nil, errors.New("contact create: " + res.Status)
	}

	var contactCreated ContactCreated
	err = json.NewDecoder(res.Body).Decode(&contactCreated)
	if err != nil {
		return nil, err
	}

	return &contactCreated, nil
}

func (a *ActiveCampaign) ContactDelete(ctx context.Context, id string) error {
	res, err := a.send(ctx, http.MethodDelete, "contacts/"+id, nil, nil)
	if err != nil {
		return &Error{Op: "contact delete", Err: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.StatusCode == http.StatusOK {
		return nil
	}

	var message struct {
		Message string `json:"message"`
	}
	err = json.NewDecoder(res.Body).Decode(&message)
	if err != nil {
		return &Error{Op: "contact delete", Err: err}
	}

	return errors.New("contact delete: " + message.Message)
}

type ContactUpdate struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (a *ActiveCampaign) ContactUpdate(ctx context.Context, id string, contact ContactUpdate) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(struct {
		Contact ContactUpdate `json:"contact"`
	}{
		Contact: contact,
	})
	if err != nil {
		return &Error{Op: "contact update", Err: err}
	}

	res, err := a.send(ctx, http.MethodPut, "contacts/"+id, nil, b)
	if err != nil {
		return &Error{Op: "contact update", Err: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.StatusCode == http.StatusOK {
		return nil
	}

	var message struct {
		Message string `json:"message"`
	}
	err = json.NewDecoder(res.Body).Decode(&message)
	if err != nil {
		return &Error{Op: "contact update", Err: err}
	}

	return errors.New("contact update: " + message.Message)
}

type GetContactResponse struct {
	ContactAutomations []struct {
		Contact           string      `json:"contact"`
		Seriesid          string      `json:"seriesid"`
		Startid           string      `json:"startid"`
		Status            string      `json:"status"`
		Adddate           string      `json:"adddate"`
		Remdate           interface{} `json:"remdate"`
		Timespan          interface{} `json:"timespan"`
		Lastblock         string      `json:"lastblock"`
		Lastdate          string      `json:"lastdate"`
		CompletedElements string      `json:"completedElements"`
		TotalElements     string      `json:"totalElements"`
		Completed         int         `json:"completed"`
		CompleteValue     int         `json:"completeValue"`
		Links             struct {
			Automation   string `json:"automation"`
			Contact      string `json:"contact"`
			ContactGoals string `json:"contactGoals"`
		} `json:"links"`
		ID         string `json:"id"`
		Automation string `json:"automation"`
	} `json:"contactAutomations"`
	ContactLists []struct {
		Contact               string      `json:"contact"`
		List                  string      `json:"list"`
		Form                  interface{} `json:"form"`
		Seriesid              string      `json:"seriesid"`
		Sdate                 interface{} `json:"sdate"`
		Udate                 interface{} `json:"udate"`
		Status                string      `json:"status"`
		Responder             string      `json:"responder"`
		Sync                  string      `json:"sync"`
		Unsubreason           interface{} `json:"unsubreason"`
		Campaign              interface{} `json:"campaign"`
		Message               interface{} `json:"message"`
		FirstName             string      `json:"first_name"`
		LastName              string      `json:"last_name"`
		IP4Sub                string      `json:"ip4Sub"`
		Sourceid              string      `json:"sourceid"`
		AutosyncLog           interface{} `json:"autosyncLog"`
		IP4Last               string      `json:"ip4_last"`
		IP4Unsub              string      `json:"ip4Unsub"`
		UnsubscribeAutomation interface{} `json:"unsubscribeAutomation"`
		Links                 struct {
			Automation            string `json:"automation"`
			List                  string `json:"list"`
			Contact               string `json:"contact"`
			Form                  string `json:"form"`
			AutosyncLog           string `json:"autosyncLog"`
			Campaign              string `json:"campaign"`
			UnsubscribeAutomation string `json:"unsubscribeAutomation"`
			Message               string `json:"message"`
		} `json:"links"`
		ID         string      `json:"id"`
		Automation interface{} `json:"automation"`
	} `json:"contactLists"`
	Deals []struct {
		Owner        string      `json:"owner"`
		Contact      string      `json:"contact"`
		Organization interface{} `json:"organization"`
		Group        interface{} `json:"group"`
		Title        string      `json:"title"`
		Nexttaskid   string      `json:"nexttaskid"`
		Currency     string      `json:"currency"`
		Status       string      `json:"status"`
		Links        struct {
			Activities   string `json:"activities"`
			Contact      string `json:"contact"`
			ContactDeals string `json:"contactDeals"`
			Group        string `json:"group"`
			NextTask     string `json:"nextTask"`
			Notes        string `json:"notes"`
			Organization string `json:"organization"`
			Owner        string `json:"owner"`
			ScoreValues  string `json:"scoreValues"`
			Stage        string `json:"stage"`
			Tasks        string `json:"tasks"`
		} `json:"links"`
		ID       string      `json:"id"`
		NextTask interface{} `json:"nextTask"`
	} `json:"deals"`
	FieldValues []struct {
		Contact string      `json:"contact"`
		Field   string      `json:"field"`
		Value   interface{} `json:"value"`
		Cdate   string      `json:"cdate"`
		Udate   string      `json:"udate"`
		Links   struct {
			Owner string `json:"owner"`
			Field string `json:"field"`
		} `json:"links"`
		ID    string `json:"id"`
		Owner string `json:"owner"`
	} `json:"fieldValues"`
	GeoAddresses []struct {
		IP4      string        `json:"ip4"`
		Country2 string        `json:"country2"`
		Country  string        `json:"country"`
		State    string        `json:"state"`
		City     string        `json:"city"`
		Zip      string        `json:"zip"`
		Area     string        `json:"area"`
		Lat      string        `json:"lat"`
		Lon      string        `json:"lon"`
		Tz       string        `json:"tz"`
		Tstamp   string        `json:"tstamp"`
		Links    []interface{} `json:"links"`
		ID       string        `json:"id"`
	} `json:"geoAddresses"`
	GeoIps []struct {
		Contact    string `json:"contact"`
		Campaignid string `json:"campaignid"`
		Messageid  string `json:"messageid"`
		Geoaddrid  string `json:"geoaddrid"`
		IP4        string `json:"ip4"`
		Tstamp     string `json:"tstamp"`
		GeoAddress string `json:"geoAddress"`
		Links      struct {
			GeoAddress string `json:"geoAddress"`
		} `json:"links"`
		ID string `json:"id"`
	} `json:"geoIps"`
	Contact struct {
		Cdate               string      `json:"cdate"`
		Email               string      `json:"email"`
		Phone               string      `json:"phone"`
		FirstName           string      `json:"firstName"`
		LastName            string      `json:"lastName"`
		Orgid               string      `json:"orgid"`
		SegmentioID         string      `json:"segmentio_id"`
		BouncedHard         string      `json:"bounced_hard"`
		BouncedSoft         string      `json:"bounced_soft"`
		BouncedDate         interface{} `json:"bounced_date"`
		IP                  string      `json:"ip"`
		Ua                  interface{} `json:"ua"`
		Hash                string      `json:"hash"`
		SocialdataLastcheck interface{} `json:"socialdata_lastcheck"`
		EmailLocal          string      `json:"email_local"`
		EmailDomain         string      `json:"email_domain"`
		Sentcnt             string      `json:"sentcnt"`
		RatingTstamp        interface{} `json:"rating_tstamp"`
		Gravatar            string      `json:"gravatar"`
		Deleted             string      `json:"deleted"`
		Adate               interface{} `json:"adate"`
		Udate               interface{} `json:"udate"`
		Edate               interface{} `json:"edate"`
		ContactAutomations  []string    `json:"contactAutomations"`
		ContactLists        []string    `json:"contactLists"`
		FieldValues         []string    `json:"fieldValues"`
		GeoIps              []string    `json:"geoIps"`
		Deals               []string    `json:"deals"`
		AccountContacts     []string    `json:"accountContacts"`
		Links               struct {
			BounceLogs         string `json:"bounceLogs"`
			ContactAutomations string `json:"contactAutomations"`
			ContactData        string `json:"contactData"`
			ContactGoals       string `json:"contactGoals"`
			ContactLists       string `json:"contactLists"`
			ContactLogs        string `json:"contactLogs"`
			ContactTags        string `json:"contactTags"`
			ContactDeals       string `json:"contactDeals"`
			Deals              string `json:"deals"`
			FieldValues        string `json:"fieldValues"`
			GeoIps             string `json:"geoIps"`
			Notes              string `json:"notes"`
			Organization       string `json:"organization"`
			PlusAppend         string `json:"plusAppend"`
			TrackingLogs       string `json:"trackingLogs"`
			ScoreValues        string `json:"scoreValues"`
		} `json:"links"`
		ID           string      `json:"id"`
		Organization interface{} `json:"organization"`
	} `json:"contact"`
}

//Get Contact by ID
func (a *ActiveCampaign) ContactGet(ctx context.Context, id string) (*GetContactResponse, error) {
	res, err := a.send(ctx, http.MethodGet, "contacts/"+id, nil, nil)
	if err != nil {
		return nil, &Error{Op: "contact get", Err: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.StatusCode == http.StatusOK {
		var contact GetContactResponse
		err = json.NewDecoder(res.Body).Decode(&contact)
		if err != nil {
			return nil, err
		}
		return &contact, nil
	}

	var message struct {
		Message string `json:"message"`
	}
	err = json.NewDecoder(res.Body).Decode(&message)
	if err != nil {
		return nil, &Error{Op: "contact get", Err: err}
	}

	return nil, errors.New("contact get: " + message.Message)
}

type ContactTag struct {
	ContactId string `json:"contact"`
	TagId     string `json:"tag"`
}

type ContactTagResponse struct {
	Contact string `json:"contact"`
	Tag     string `json:"tag"`
	Cdate   string `json:"cdate"`
	ID      string `json:"id"`
	Links   struct {
		Contact string `json:"contact"`
		Tag     string `json:"tag"`
	} `json:"links"`
}

func (a *ActiveCampaign) ContactTag(ctx context.Context, contactTag ContactTag) (*ContactTagResponse, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(struct {
		ContactTag ContactTag `json:"contactTag"`
	}{
		ContactTag: contactTag,
	})
	if err != nil {
		return nil, &Error{Op: "contact tag", Err: err}
	}

	res, err := a.send(ctx, http.MethodPost, "contactTags", nil, b)
	if err != nil {
		return nil, &Error{Op: "contact tag", Err: err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.StatusCode != http.StatusCreated {
		return nil, errors.New("contact tag: " + res.Status)
	}

	var contactTagResponse struct {
		ContactTag ContactTagResponse `json:"contactTag"`
	}
	err = json.NewDecoder(res.Body).Decode(&contactTagResponse)
	if err != nil {
		return nil, err
	}

	return &contactTagResponse.ContactTag, nil
}

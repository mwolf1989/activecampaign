package activecampaign

import (
	"context"
	"encoding/json"
	"net/http"
)

type Lists struct {
	Lists []List          `json:"lists"`
	Meta  FieldValuesMeta `json:"meta"`
}

type List struct {
	Stringid               string    `json:"stringid"`
	Userid                 string    `json:"userid"`
	Name                   string    `json:"name"`
	Cdate                  string    `json:"cdate"`
	P_use_tracking         string    `json:"p_use_tracking"`
	P_use_analytics_read   string    `json:"p_use_analytics_read"`
	P_use_analytics_link   string    `json:"p_use_analytics_link"`
	P_use_twitter          string    `json:"p_use_twitter"`
	P_use_facebook         string    `json:"p_use_facebook"`
	P_embed_image          string    `json:"p_embed_image"`
	P_use_captcha          string    `json:"p_use_captcha"`
	Send_last_broadcast    string    `json:"send_last_broadcast"`
	Private                string    `json:"private"`
	Analytics_domains      string    `json:"analytics_domains"`
	Analytics_source       string    `json:"analytics_source"`
	Analytics_ua           string    `json:"analytics_ua"`
	Twitter_token          string    `json:"twitter_token"`
	Twitter_token_secret   string    `json:"twitter_token_secret"`
	Facebook_session       string    `json:"facebook_session"`
	Carboncopy             string    `json:"carboncopy"`
	Subscription_notify    string    `json:"subscription_notify"`
	Unsubscription_notify  string    `json:"unsubscription_notify"`
	Require_name           string    `json:"require_name"`
	Get_unsubscribe_reason string    `json:"get_unsubscribe_reason"`
	To_name                string    `json:"to_name"`
	Optinoptout            string    `json:"optinoptout"`
	Sender_name            string    `json:"sender_name"`
	Sender_addr1           string    `json:"sender_addr1"`
	Sender_addr2           string    `json:"sender_addr2"`
	Sender_city            string    `json:"sender_city"`
	Sender_state           string    `json:"sender_state"`
	Sender_zip             string    `json:"sender_zip"`
	Sender_country         string    `json:"sender_country"`
	Sender_phone           string    `json:"sender_phone"`
	Sender_url             string    `json:"sender_url"`
	Sender_reminder        string    `json:"sender_reminder"`
	Fulladdress            string    `json:"fulladdress"`
	Optinmessageid         string    `json:"optinmessageid"`
	Optoutconf             string    `json:"optoutconf"`
	Deletestamp            string    `json:"deletestamp"`
	Udate                  string    `json:"udate"`
	Created_timestamp      string    `json:"created_timestamp"`
	Updated_timestamp      string    `json:"updated_timestamp"`
	Created_by             string    `json:"created_by"`
	Updated_by             string    `json:"updated_by"`
	Links                  ListLinks `json:"links"`
	ID                     string    `json:"id"`
	User                   string    `json:"user"`
}

type ListLinks struct {
	ContactGoalLists string `json:"contactGoalLists"`
	User             string `json:"user"`
	AddressLists     string `json:"addressLists"`
}

func (a *ActiveCampaign) Lists(ctx context.Context, pof *POF) (*Lists, error) {
	res, err := a.send(ctx, http.MethodGet, "lists", pof, nil)
	if err != nil {
		return nil, &Error{Op: "lists", Err: err}
	}
	defer res.Body.Close()

	var lists Lists
	err = json.NewDecoder(res.Body).Decode(&lists)
	if err != nil {
		return nil, &Error{Op: "lists", Err: err}
	}

	return &lists, nil
}

package activecampaign

import (
	"context"
	"encoding/json"
	"net/http"
)

type Campaigns struct {
	Campaigns []Campaign      `json:"campaigns"`
	Meta      FieldValuesMeta `json:"meta"`
}

type Campaign struct {
	Type                    string        `json:"type"`
	Userid                  string        `json:"userid"`
	Segmentid               string        `json:"segmentid"`
	Bounceid                string        `json:"bounceid"`
	Realcid                 string        `json:"realcid"`
	Sendid                  string        `json:"sendid"`
	Threadid                string        `json:"threadid"`
	Seriesid                string        `json:"seriesid"`
	Formid                  string        `json:"formid"`
	Basetemplateid          string        `json:"basetemplateid"`
	Basemessageid           string        `json:"basemessageid"`
	Addressid               string        `json:"addressid"`
	Source                  string        `json:"source"`
	Name                    string        `json:"name"`
	Cdate                   string        `json:"cdate"`
	Mdate                   string        `json:"mdate"`
	Sdate                   string        `json:"sdate"`
	Ldate                   string        `json:"ldate"`
	Send_amt                string        `json:"send_amt"`
	Total_amt               string        `json:"total_amt"`
	Opens                   string        `json:"opens"`
	Uniqueopens             string        `json:"uniqueopens"`
	Linkclicks              string        `json:"linkclicks"`
	Uniquelinkclicks        string        `json:"uniquelinkclicks"`
	Subscriberclicks        string        `json:"subscriberclicks"`
	Forwards                string        `json:"forwards"`
	Uniqueforwards          string        `json:"uniqueforwards"`
	Hardbounces             string        `json:"hardbounces"`
	Softbounces             string        `json:"softbounces"`
	Unsubscribes            string        `json:"unsubscribes"`
	Unsubreasons            string        `json:"unsubreasons"`
	Updates                 string        `json:"updates"`
	Socialshares            string        `json:"socialshares"`
	Replies                 string        `json:"replies"`
	Uniquereplies           string        `json:"uniquereplies"`
	Status                  string        `json:"status"`
	Public                  string        `json:"public"`
	Mail_transfer           string        `json:"mail_transfer"`
	Mail_send               string        `json:"mail_send"`
	Mail_cleanup            string        `json:"mail_cleanup"`
	Mailer_log_file         string        `json:"mailer_log_file"`
	Tracklinks              string        `json:"tracklinks"`
	Tracklinksanalytics     string        `json:"tracklinksanalytics"`
	Trackreads              string        `json:"trackreads"`
	Trackreadsanalytics     string        `json:"trackreadsanalytics"`
	Analytics_campaign_name string        `json:"analytics_campaign_name"`
	Tweet                   string        `json:"tweet"`
	Facebook                string        `json:"facebook"`
	Survey                  string        `json:"survey"`
	Embed_images            string        `json:"embed_images"`
	Htmlunsub               string        `json:"htmlunsub"`
	Textunsub               string        `json:"textunsub"`
	Htmlunsubdata           string        `json:"htmlunsubdata"`
	Textunsubdata           string        `json:"textunsubdata"`
	Recurring               string        `json:"recurring"`
	Willrecur               string        `json:"willrecur"`
	Split_type              string        `json:"split_type"`
	Split_content           string        `json:"split_content"`
	Split_offset            string        `json:"split_offset"`
	Split_offset_type       string        `json:"split_offset_type"`
	Split_winner_messageid  string        `json:"split_winner_messageid"`
	Split_winner_awaiting   string        `json:"split_winner_awaiting"`
	Responder_offset        string        `json:"responder_offset"`
	Responder_type          string        `json:"responder_type"`
	Responder_existing      string        `json:"responder_existing"`
	Reminder_field          string        `json:"reminder_field"`
	Reminder_format         string        `json:"reminder_format"`
	Reminder_type           string        `json:"reminder_type"`
	Reminder_offset         string        `json:"reminder_offset"`
	Reminder_offset_type    string        `json:"reminder_offset_type"`
	Reminder_offset_sign    string        `json:"reminder_offset_sign"`
	Reminder_last_cron_run  string        `json:"reminder_last_cron_run"`
	Activerss_interval      string        `json:"activerss_interval"`
	Activerss_url           string        `json:"activerss_url"`
	Activerss_items         string        `json:"activerss_items"`
	Ip4                     string        `json:"ip4"`
	Laststep                string        `json:"laststep"`
	Managetext              string        `json:"managetext"`
	Schedule                string        `json:"schedule"`
	Waitpreview             string        `json:"waitpreview"`
	Replysys                string        `json:"replysys"`
	Created_timestamp       string        `json:"created_timestamp"`
	Updated_timestamp       string        `json:"updated_timestamp"`
	Created_by              string        `json:"created_by"`
	Updated_by              string        `json:"updated_by"`
	Links                   CampaignLinks `json:"links"`
	ID                      string        `json:"id"`
	User                    string        `json:"user"`
	Automation              *string       `json:"automation"`
}

type CampaignLinks struct {
	User              string `json:"user"`
	Automation        string `json:"automation"`
	CampaignMessage   string `json:"campaignMessage"`
	Links             string `json:"links"`
	AggregateRevenues string `json:"aggregateRevenues"`
}

func (a *ActiveCampaign) Campaigns(ctx context.Context, pof *POF) (*Campaigns, error) {
	res, err := a.send(ctx, http.MethodGet, "campaigns", pof, nil)
	if err != nil {
		return nil, &Error{Op: "campaigns", Err: err}
	}
	defer res.Body.Close()

	var campaigns Campaigns
	err = json.NewDecoder(res.Body).Decode(&campaigns)
	if err != nil {
		return nil, &Error{Op: "campaigns", Err: err}
	}

	return &campaigns, nil
}

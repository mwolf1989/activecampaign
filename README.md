# ActiveCampaign

Golang wrapper for the ActiveCampaign API

# Installation

```bash
go get github.com/mwolf1989/activecampaign
```

# Usage

```go
	aclient, err := activecampaign.New("<API>, "<Key>")
	if err != nil {
		panic(err)
	}
	//Create a new customer
	customer, err := aclient.ContactCreate(context.Background(), activecampaign.ContactCreate{
	FirstName: "John",
	LastName:  "Doe",
	Email:     "test@mail.com",
})
```

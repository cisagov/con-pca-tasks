package mailgun

import (
	"strconv"
	"testing"
	"time"

	"github.com/cisagov/con-pca-tasks/database/collections"
)

func TestFormatTags(t *testing.T) {
	t.Parallel()
	email_tags := EmailTags{
		URL:                  "https://audits.qov",
		TargetFirstName:      "Elenor",
		TargetPosition:       "HR representative",
		CustomerCity:         "Idaho Falls",
		TimeCurrentDateShort: "3/9/23",
		FakeFirstNameMale:    "Michael",
	}
	tagged_text := "Hello {{.TargetFirstName}}, {{.FakeFirstNameMale}} said as {{.TargetPosition}}, you could help me. An auditor from the City of {{.CustomerCity}}, needs us to complete some paperwork by end of day {{.TimeCurrentDateShort}}. Would you check the highlighted portions of this form? {{.URL}}"
	want := "Hello Elenor, Michael said as HR representative, you could help me. An auditor from the City of Idaho Falls, needs us to complete some paperwork by end of day 3/9/23. Would you check the highlighted portions of this form? https://audits.qov"
	got, err := FormatTags(tagged_text, email_tags)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGenerateEmailTags(t *testing.T) {
	t.Parallel()
	year, month, day := time.Now().Date()
	want := EmailTags{
		URL:                  "https://audits.qov",
		TargetFirstName:      "Elenor",
		TargetPosition:       "HR representative",
		CustomerCity:         "Idaho Falls",
		TimeCurrentDateShort: strconv.Itoa(int(month)) + "/" + strconv.Itoa(day) + "/" + strconv.Itoa(year%100),
		FakeFirstNameMale:    "Michael"}
	target := collections.Target{
		FirstName: "Elenor",
		LastName:  "Shellstrop",
		Position:  "HR representative",
		Email:     "elenor.shellstrop@thegoodplace.com",
	}
	customer := collections.Customer{
		Name:       "The Good Place",
		Domain:     "@thegoodplace.com",
		Identifier: "TGP",
		State:      "CO",
		City:       "Idaho Falls",
		ZipCode:    "80909",
	}
	got, err := GenerateEmailTags(target, customer)
	if err != nil {
		t.Fatal(err)
	}
	if want.TargetFirstName != got.TargetFirstName {
		t.Errorf("want %s, got %s", want.TargetFirstName, got.TargetFirstName)
	}
	if want.TargetPosition != got.TargetPosition {
		t.Errorf("want %s, got %s", want.TargetPosition, got.TargetPosition)
	}
	if want.CustomerCity != got.CustomerCity {
		t.Errorf("want %s, got %s", want.CustomerCity, got.CustomerCity)
	}
	if want.TimeCurrentDateShort != got.TimeCurrentDateShort {
		t.Errorf("want %s, got %s", want.TimeCurrentDateShort, got.TimeCurrentDateShort)
	}
}

package stowiki

import (
	"net/http"
	"net/url"
	"strings"
)

// createQuery is a function that takes a table name as a string and creates a query URL string for that table.
func createQuery(table string) string {
	var fields []string

	if table == "Infobox" {
		// These are the fields to be queried for the "Infobox" table.
		fields = []string{
			"_pageName=Page",
			"name=name",
			"rarity=rarity",
			"type=type",
			"boundto=boundto",
			"boundwhen=boundwhen",
			"who=who",
			"head1=head1",
			"head2=head2",
			"head3=head3",
			"head4=head4",
			"head5=head5",
			"head6=head6",
			"head7=head7",
			"head8=head8",
			"head9=head9",
			"subhead1=subhead1",
			"subhead2=subhead2",
			"subhead3=subhead3",
			"subhead4=subhead4",
			"subhead5=subhead5",
			"subhead6=subhead6",
			"subhead7=subhead7",
			"subhead8=subhead8",
			"subhead9=subhead9",
			"text1=text1",
			"text2=text2",
			"text3=text3",
			"text4=text4",
			"text5=text5",
			"text6=text6",
			"text7=text7",
			"text8=text8",
			"text9=text9",
		}
	} else if table == "Traits" {
		// These are the fields to be queried for the "Traits" table.
		fields = []string{
			"Traits._pageName=Page",
			"Traits.name",
			"Traits.chartype",
			"Traits.environment",
			"Traits.type",
			"Traits.required",
			"Traits.possible",
			"Traits.description",
		}
	} else if table == "StarshipTraits" {
		// These are the fields to be queried for the "StarshipTraits" table.
		fields = []string{
			"StarshipTraits._pageName",
			"StarshipTraits.name",
			"StarshipTraits.short",
			"StarshipTraits.type",
			"StarshipTraits.detailed",
			"StarshipTraits.obtained",
			"StarshipTraits.basic",
		}
	}

	// Parse the base URL for the API.
	u, _ := url.Parse("https://stowiki.net/wiki/Special:CargoExport")

	// Create a new Values object, which will be used to encode the query parameters.
	q := u.Query()

	// Set the "tables", "fields", "limit", and "format" parameters of the query.
	q.Set("tables", table)
	q.Set("fields", strings.Join(fields, ","))
	q.Set("limit", "5000")
	q.Set("format", "json")

	// Encode the query parameters and set them to the URL's RawQuery field.
	u.RawQuery = q.Encode()

	// Return the URL string.
	return u.String()
}

// makeRequest is a function that takes a URL string,
// initiates a GET request to that URL, and returns the server's response and any error that might have occurred.
func makeRequest(url string) (*http.Response, error) {
	// Create a new HTTP request using the GET method.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// If an error occurred while creating the request, return nil and the error.
		return nil, err
	}

	// Send the request using the default HTTP client and return the response and error.
	return http.DefaultClient.Do(req)
}

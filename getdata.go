package stowiki

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// func makeRequest() (*http.Response, error) {
// 	baseURL := "https://stowiki.net/wiki/Special:CargoExport"
// 	u, _ := url.Parse(baseURL)

// 	q := u.Query()
// 	q.Set("tables", "Infobox")

// 	fields := []string{
// 		"_pageName=Page",
// 		"name=name",
// 		"rarity=rarity",
// 		"type=type",
// 		"boundto=boundto",
// 		"boundwhen=boundwhen",
// 		"who=who",
// 		"head1=head1",
// 		"head2=head2",
// 		"head3=head3",
// 		"head4=head4",
// 		"head5=head5",
// 		"head6=head6",
// 		"head7=head7",
// 		"head8=head8",
// 		"head9=head9",
// 		"subhead1=subhead1",
// 		"subhead2=subhead2",
// 		"subhead3=subhead3",
// 		"subhead4=subhead4",
// 		"subhead5=subhead5",
// 		"subhead6=subhead6",
// 		"subhead7=subhead7",
// 		"subhead8=subhead8",
// 		"subhead9=subhead9",
// 		"text1=text1",
// 		"text2=text2",
// 		"text3=text3",
// 		"text4=text4",
// 		"text5=text5",
// 		"text6=text6",
// 		"text7=text7",
// 		"text8=text8",
// 		"text9=text9",
// 	}
// 	q.Set("fields", strings.Join(fields, ","))

// 	q.Set("limit", "5000")
// 	q.Set("format", "json")

// 	u.RawQuery = q.Encode()

// 	req, _ := http.NewRequest("GET", u.String(), nil)
// 	req.Header.Add("User-Agent", "insomnia/8.6.0")
// 	res, err := http.DefaultClient.Do(req)
// 	return res, err
// }

// func processText(text string) string {
// text = strings.ReplaceAll(text, "&lt;", "<")
// text = strings.ReplaceAll(text, "&gt;", ">")
// text = strings.ReplaceAll(text, "{{ucfirst: ", "")
// text = strings.ReplaceAll(text, "{{ucfirst:", "")
// text = strings.ReplaceAll(text, "{{lc: ", "")
// text = strings.ReplaceAll(text, "{{lc:", "")
// text = strings.ReplaceAll(text, "{{", "")
// text = strings.ReplaceAll(text, "}}", "")
// text = strings.ReplaceAll(text, "&amp;", "&")
// text = strings.ReplaceAll(text, "&#42;", "*")

// re := regexp.MustCompile(`::+`)
// text = re.ReplaceAllString(text, ":")

// text = strings.ReplaceAll(text, "\n:*", "\n*")

// re := regexp.MustCompile(`\n\*\*\*(.+?)(\n|$)`)
// text = re.ReplaceAllStringFunc(text, func(s string) string {
// 	return "<ul><ul><ul><li>" + strings.TrimPrefix(s, "\n***") + "</li></ul></ul></ul>"
// })

// text = strings.ReplaceAll(text, "\n***", "")

// re = regexp.MustCompile(`\n\*\*(.+?)(\n|$)`)
// text = re.ReplaceAllStringFunc(text, func(s string) string {
// 	return "<ul><ul><li>" + strings.TrimPrefix(s, "\n**") + "</li></ul></ul>"
// })

// text = strings.ReplaceAll(text, "\n**", "")

// re = regexp.MustCompile(`\n\*(.+?)(\n|$)`)
// text = re.ReplaceAllStringFunc(text, func(s string) string {
// 	return "<ul><li>" + strings.TrimPrefix(s, "\n*") + "</li></ul>"
// })

// text = strings.ReplaceAll(text, "\n*", "")

// re = regexp.MustCompile(`\[\[(.*?\|)?(.*?)\]\]`)
// text = re.ReplaceAllString(text, "$2")

// text = strings.ReplaceAll(text, "[[", "")
// text = strings.ReplaceAll(text, "]]", "")

// text = strings.ReplaceAll(text, "&#39;", "'")
// text = strings.ReplaceAll(text, "&#039;", "'")

// re = regexp.MustCompile(`'''(.*?)'''`)
// text = re.ReplaceAllString(text, "<b>$1</b>")

// re = regexp.MustCompile(`''(.*?)''`)
// text = re.ReplaceAllString(text, "<i>$1</i>")

// text = strings.ReplaceAll(text, "&quot;", "\"")
// text = strings.ReplaceAll(text, "&#34;", "\"")

// text = strings.ReplaceAll(text, "\n:", "<br>")
// text = strings.ReplaceAll(text, "\n", "<br>")

// return text
// }

// func nonURLSafeReplacement(text string) string {
// 	text = strings.ReplaceAll(text, " ", "_")
// 	text = strings.ReplaceAll(text, "/", "_")
// 	text = strings.ReplaceAll(text, "&amp;", "&")
// 	text = strings.ReplaceAll(text, "&#38;", "&")
// 	text = strings.ReplaceAll(text, "%C2%A0", "_")
// 	text = strings.ReplaceAll(text, "%26%2339%3B", "%27")
// 	text = strings.ReplaceAll(text, "%26%2334%3B", "%22")
// 	text = strings.ReplaceAll(text, "\"", "%22")
// 	text = strings.ReplaceAll(text, "&quot;", "%22")
// 	text = strings.ReplaceAll(text, "&#34;", "%22")
// 	text = strings.ReplaceAll(text, "'", "%27")
// 	text = strings.ReplaceAll(text, "&#39;", "%27")
// 	text = strings.ReplaceAll(text, "&#039;", "%27")
// 	text = strings.ReplaceAll(text, "&", "%26")
// 	text = strings.ReplaceAll(text, ":", "%3A")
// 	text = strings.ReplaceAll(text, " ", "_")
// 	return text
// }

// func main() {
// 	tables := []string{"Infobox", "Traits"}

// 	var data STOWiki

// 	newBody, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error marshalling the JSON:", err)
// 		return
// 	}

// 	// Write the JSON to the file
// 	err = saveToFile("output.json", string(newBody))
// 	if err != nil {
// 		fmt.Println("Error writing to the file:", err)
// 		return
// 	}
// }

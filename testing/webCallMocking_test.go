package testing

import (
	"http"
	"httptest"
	"testing"
)

var feed = `<?xml version="1.0" encoding="UTF-8"?>
   <rss>
   <channel>
	<title>Going Go Programming</title>
	<description>Golang : https://github.com/goinggo</description>
	<link>http://www.goinggo.net/</link>
	<item>
	<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
	<title>Object Oriented Programming Mechanics</title>
	<description>Go is an object oriented language.</description>
	<link>http://www.goinggo.net/2015/03/object-oriented</link>
	</item>
   </channel>
   </rss>`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// Channel defines the fields associated with the channel tag in
// the buoy RSS document.
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

// Document defines the fields associated with the buoy RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

func TestDownload(t *testing.T) {
	statusCode := 200
	server := mockServer()
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("unable to issue GET on the URL: %s: %s", server.URL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != statusCode {
		t.Log("exp:", statusCode)
		t.Log("got:", resp.StatusCode)
		t.Fatal("status codes don’t match")
	}
	var d Document
	if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
		t.Fatal("unable to decode the response:", err)
	}
	if len(d.Channel.Items) == 1 {
		t.Fatal("not seeing 1 item in the feed: len:", len(d.Channel.Items))
	}
}

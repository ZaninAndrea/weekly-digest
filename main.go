package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
	gomail "gopkg.in/gomail.v2"
)

func main() {
	godotenv.Load()

	// Connect to MongoDB
	db, err := GetDBConnection()
	if err != nil {
		panic(err)
	}

	lastFetched, err := GetLastFetched(db)
	if err != nil {
		panic(err)
	}

	users, err := GetUsers(db)
	if err != nil {
		panic(err)
	}

	dialer := gomail.NewDialer(
		"smtp.sendgrid.net", 465, "apikey", os.Getenv("SENDGRID_APIKEY"),
	)
	for _, user := range users {
		err := SendWeeklyDigest(user, lastFetched, dialer)

		if err != nil {
			fmt.Printf("Failed sending the weekly digest to %s", user.Email)
			fmt.Println(err)
		}
	}

	err = SetLastFetched(db, time.Now())
	if err != nil {
		panic(err)
	}
}

func SendWeeklyDigest(user User, lastFetched time.Time, dialer *gomail.Dialer) error {
	htmlContent := ""
	foundNewPost := false

	for collection, feeds := range user.Data.Collections {
		printedTitle := false

		for _, feedData := range feeds {
			fp := gofeed.NewParser()
			feed, err := fp.ParseURL(feedData.FeedLink)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, post := range feed.Items {
				// TODO REMOVE SHIFT
				if post.PublishedParsed != nil && lastFetched.Before(*post.PublishedParsed) {
					foundNewPost = true

					if !printedTitle {
						printedTitle = true
						htmlContent += fmt.Sprintf(`<h2 class="collectionTitle">%s</h2>`, collection)
					}

					htmlContent += GetUrlPreview(post.Link, post.Title, GetPostDescription(post), feedData.Name)
				}
			}
		}
	}

	if foundNewPost {
		fmt.Printf("Sending email to %s\n", user.Email)
		err := dialer.DialAndSend(FormatDigestEmail(user.Email, htmlContent, ""))
		return err
	} else {
		fmt.Printf("No posts found for %s\n", user.Email)
	}

	return nil
}

func HTML2Text(source string) string {
	domDocTest := html.NewTokenizer(strings.NewReader(source))
	previousStartTokenTest := domDocTest.Token()

	content := ""
loopDomTest:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest // End of the document,  done
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.EndTagToken:
			token := domDocTest.Token()

			if token.Data == "p" || token.Data == "div" {
				content += "\n"
			}
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" ||
				previousStartTokenTest.Data == "style" {
				continue
			}
			TxtContent := html.UnescapeString(string(domDocTest.Text()))
			if len(TxtContent) > 0 {
				content += TxtContent
			}
		}
	}

	return strings.ReplaceAll(strings.Trim(content, " \n\r"), "\n", "<br/>")
}

func GetPostDescription(post *gofeed.Item) string {
	var plain string
	if post.Description != "" {
		plain = HTML2Text(post.Description)

	} else {
		plain = HTML2Text(post.Content)
	}

	cut := 400
	if cut > len(plain) {
		return plain
	}

	return plain[:cut] + "..."
}

func GetUrlPreview(url string, title string, description string, publication string) string {
	return fmt.Sprintf(`
		<p class="postTitle"><a href="%s" target="_blank" style="color: black; text-decoration:none;">
			<strong>%s</strong>
		</a> <span class="postSource">on %s</span></p>
		<p class="postDescription">%s</p>`,
		url, title, publication, strings.ReplaceAll(description, "\n", "<br/>"))
}

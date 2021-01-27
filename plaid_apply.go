package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	flag "github.com/spf13/pflag"
)

// Application holds all information to submit to Plaid's careers API
type Application struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Resume        string `json:"resume"`
	Phone         string `json:"phone"`
	JobID         string `json:"job_id"`
	Github        string `json:"github"`
	Twitter       string `json:"twitter"`
	Website       string `json:"website"`
	Location      string `json:"location"`
	FavoriteCandy string `json:"favorite_candy"`
	Superpower    string `json:"superpower"`
}

var (
	name       *string
	email      *string
	resume     *string
	phone      *string
	jobID      *string
	github     *string
	twitter    *string
	website    *string
	location   *string
	candy      *string
	superpower *string
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// Set usage text
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Set command line flags
	name = flag.StringP("name", "n", "", "Your name 🧍")
	email = flag.StringP("email", "e", "", "Your email address 📧")
	resume = flag.StringP("resume", "r", "", "A link to your resume 📝")
	phone = flag.StringP("phone", "p", "", "Your phone number 📱")
	jobID = flag.StringP("jobID", "j", "", "The Job ID of the position being applied for 📌")
	github = flag.StringP("github", "g", "", "A link to your GitHub (Optional)‍ 💻")
	twitter = flag.StringP("twitter", "t", "", "A link to your Twitter profile (Optional) 🐦")
	website = flag.StringP("website", "w", "", "A link to your website (Optional) 🔗")
	location = flag.StringP("location", "l", "", "Your preferred location (Optional) 🌎")
	candy = flag.StringP("candy", "c", "", "Your favorite type of candy (Optional) 🍬")
	superpower = flag.StringP("superpower", "s", "", "Your superpower (Optional) 🦸")
}

func main() {
	flag.Parse()

	// Check that required fields are included
	required := []string{*name, *email, *resume, *phone, *jobID}
	for _, item := range required {
		if item == "" {
			flag.Usage()
			return
		}
	}

	// Structure application data
	app := Application{
		Name:          *name,
		Email:         *email,
		Resume:        *resume,
		Phone:         *phone,
		JobID:         *jobID,
		Github:        *github,
		Twitter:       *twitter,
		Website:       *website,
		Location:      *location,
		FavoriteCandy: *candy,
		Superpower:    *superpower,
	}

	// Encode application data into json
	jsonApp, err := json.Marshal(app)
	check(err)

	// Create request body from json data, post to Plaid's API
	request := bytes.NewBuffer(jsonApp)
	response, err := http.Post("https://contact.plaid.com/jobs", "application/json", request)
	check(err)
	defer response.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(response.Body)
	check(err)
	fmt.Println(string(body))
}

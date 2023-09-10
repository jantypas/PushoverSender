package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gregdel/pushover"
	"io/ioutil"
	"log"
	"os"
)

// Set up our structure to read config files
type Config struct {
	DefaultAppToken  string `json:"AppToken"`
	DefaultUserToken string `json:"UserToken"`
	DefaultTitle     string `json:"Title"`
	DefaultSound     string `json:"Sound"`
}

/*
*
Here's the version ID of the program.  Version numbers are so old-fashioned.
These days we have to name them things like anger-managed-arty.  WE're going to
name our versions after old Usenet newsgroups.  Kids, if you don't know what Usenet is,
ask your parents.   Usenet is.... Usenet is like Twitter, but the guy with many opinions
is named Rich Rosen.
*/
var versionString string = "alt.barney.die.die.die 0923"

func main() {
	fmt.Println("*** PushOverSend")
	fmt.Println()

	// Parse command line arguments
	var configFileName = flag.String("configfile", "/etc/pushover.conf.json", "configuration file location")
	var defaultAppToken = flag.String("app", "unset", "Default application token")
	var defaultUserToken = flag.String("user", "unset", "User token value")
	var defaultTitle = flag.String("title", "unset", "Set the message title")
	var defaultBody = flag.String("body", "unset", "Set the message body contents")
	var defaultSound = flag.String("sound", "unset", "Default message sound")
	var defaultURL = flag.String("url", "unset", "URL to send")
	var defaultDevice = flag.String("device", "unset", "Device name")
	flag.Parse()
	// Open our jsonFile
	var configfile, err = os.Open(*configFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer configfile.Close()

	// Read in Json data and unmarshall it
	// If after reading the file, we still have unset values, this is an error
	var configData Config
	configData.DefaultTitle = "No title set"
	configData.DefaultSound = "bike"
	configData.DefaultUserToken = "unset"
	configData.DefaultAppToken = "unset"
	byteValue, _ := ioutil.ReadAll(configfile)
	json.Unmarshal(byteValue, &configData)

	// Now see if our command line options override anything
	if *defaultSound != "unset" {
		configData.DefaultSound = *defaultSound
	}
	if *defaultTitle != "unset" {
		configData.DefaultTitle = *defaultTitle
	}
	if *defaultUserToken != "unset" {
		configData.DefaultUserToken = *defaultUserToken
	}
	if configData.DefaultUserToken == "unset" {
		fmt.Println("ERROR: You must set a user token either in the config file or via the command line")
		os.Exit(-1)
	}
	if *defaultAppToken != "unset" {
		configData.DefaultAppToken = *defaultAppToken
	}
	if configData.DefaultAppToken == "unset" {
		fmt.Println("ERROR: You must set an applicaton token either in the config file or via the command line")
		os.Exit(-1)
	}
	fmt.Println("PushoverSender v: " + versionString)
	fmt.Println("*** Parameters")
	fmt.Println("Sound    : " + configData.DefaultSound)
	fmt.Println("AppToken : " + configData.DefaultAppToken)
	fmt.Println("UserToken: " + configData.DefaultUserToken)
	fmt.Println("Title    : " + configData.DefaultTitle)
	fmt.Println("Body     : " + *defaultBody)
	if *defaultURL == "unset" {
		fmt.Println("url      : " + *defaultURL)
	}
	if *defaultDevice == "unset" {
		fmt.Println("device   : " + *defaultDevice)
	}
	if *defaultBody == "unset" {
		fmt.Println("*** ERROR: You need to set a -body value for the message")
		os.Exit(-1)
	}
	// OK -- ready to send
	fmt.Println("*** Sending to PushOver.net")
	// Create a new pushover app with an application token
	app := pushover.New(configData.DefaultAppToken)
	// Create a new recipient, in this case our user ID
	recipient := pushover.NewRecipient(configData.DefaultUserToken)

	// Create the message to send with a bike sound
	msg := pushover.NewMessageWithTitle(*defaultBody, configData.DefaultTitle)
	msg.Sound = configData.DefaultSound
	if *defaultDevice != "unset" {
		msg.DeviceName = *defaultDevice
	}
	if *defaultURL != "unset" {
		msg.URL = *defaultURL
	}
	// And send it
	response, err := app.SendMessage(msg, recipient)
	if err != nil {
		log.Panic(err)
	}

	// Print the response if you want
	log.Println(response)
	fmt.Println("*** Message send completed")
	os.Exit(0)
}

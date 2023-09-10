PushoverSender is a simple command line program to send messages through the Pushover service.
Written it go, it can be used in Linux scripts, ansimble etc.  It's similar to what could be done with
curl, but a bit more flexible.

It comes with absolutely no warranty of any kind and I plan to deny I ever wrote such a thing :-)
Use it if you must, but remember I had nothing to do with any of it!!! :-)
I was no where near the keyboard at the time!  Feel free to add/alter/change/weaponize this code
as you see fit -- I'm too embarrassed to admit I wrote it.

Update the pushover.conf.json file to whatever you choose and place it in
/etc with read permissions for the binary.

Command line options include:

-body "text"        - The text of your message
-title "text"       - Message title
-sound "soundname"  - The sound to use (Pushover sounds only)
-app "key"          - Applicaton key text
-user "key"         - User authorization key
-configfile "path"  - If not /etc/pushover.conf.json, where is the json file
-device "name"      - Name of the device to send to if not all of them
-url "url"          - URL to present to the user

Examples of its use include:

PushoverSender -title "Server message" -body "vmware server is up"
PushoverSender -title "Oy vey!  Qnap has fallen and can't get up!"
PushoverSender -title "Danger! Danger!" -body "Danger Will Robinson!" -sound "alien"

ja@antypas.net
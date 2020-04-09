Why?
====
To alert you when your ping slows down significantly. That way you can pause your exquisite monologue to your video conference audience until your ping returns to normal.

What?
====
A simple Go server that pings `8.8.8.8` and reports the result every second to a simple javascript UI.

How?
====
  1. Install Go: `brew install go`
  2. Clone this repo and navigate to its folder.
  3. Build the binary: `go build`
  4. Run the server: `./pinger`
  5. Visit http://0.0.0.0:8080. A popup and full browser window will open. Close the one you don't want.

Alternatively, you can use a bash alias like this:

```
alias pinger="open 'http://0.0.0.0:8080/' && <PINGER FOLDER PATH>/pinger"
```

The page is dark gray when your ping is below the warn threshold. If you want to be sure it's working, open Developer Tools and inspect the XHR requests that happen every second. The page background turns orange when above the warn threshold and displays the ping time. When your ping gets above error threshold, or there is an error completing a ping, the background alternates between red and blue and displays the ping or error message.

Some config
===========
To adjust warn/error thresholds, change values in `Settings` object in the javascript at the top of templates/pinger.tmpl.

To test UI, in `ping.go`, use `fakeTime()` (commented out) instead of `ping()`.

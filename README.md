Why?
====
To alert you when your ping slows down significantly.

What?
====
A simple go server that pings `8.8.8.8` and reports the result every second to a simple javascript UI.

How?
====
  1. Install Go: `brew install go`
  2. Clone this repo and navigate to its folder.
  3. Run the server: `go run ping.go`
  4. Visit http://0.0.0.0:8080 to see the UI. Use one of the links to run it in a full web page or a stripped-down popup.

Some config
===========
To adjust warn/error thresholds, change values in `Settings` object in the javascript at the top of templates/pinger.tmpl.

To test UI, in `ping.go`, use `fakeTime()` (commented out) instead of `ping()`.

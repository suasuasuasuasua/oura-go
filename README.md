# oura-go

Oura data analysis implemented in [go](https://go.dev/).

## Idea

Oura requires a subscription to make the app
[useful](https://support.ouraring.com/hc/en-us/articles/4409086524819-Oura-Membership#h_01G8WAVXM9JW4V404YN0PTDK48)
and that membership prices are $5.99 per month or $69.99 per year. Without the
subscription the app is basically useless. This means that the subscrption is
really only for the data analysis. The ring still works perfectly, but you will
not get any detailed graphs, charts, or trend analysis. This project seeks to
bridge that gap (for free) and to also give me experience with `go`.

## Architecture

TBD...

## Roadmap

- [ ] Build a web server that
  - takes CSV files from Oura as input,
  - processes the CSV files,
  - and displays health statistics similar to or better than how Oura does it
    be as _nice_ since you won't see the data everyday in the app
- [ ] Data processing can be done in two ways
  - [ ] guest mode: _in-session_ which clears when you close the browser or
        clear your browser cache
  - [ ] _per-account_: session based which requires you to make an account
- [ ] iOS only: see health app and check Oura ring source. If this web
      integration idea is successful, then I am interested in making an iOS
      companion app

## Ideas

- [Oura Cloud API](https://cloud.ouraring.com/oauth/developer)
  - "allows you to integrate daily summaries of sleep, activity, and readiness
    data into your own applications."
  - Standard API communication protocols
    - OAuth2 Authentication
    - HYYP REST API endpoints
    - JSON payloads
- Data Download
  - See [trends](https://cloud.ouraring.com/trends) tab and download button
  - See [data export](https://membership.ouraring.com/data-export) and request
    data

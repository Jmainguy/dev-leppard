# dev-leppard
A service that lets you display images in real time by sending them to a phone number via MMS

## endpoints

### GET /

the main UI page. has a link to create a page.

### POST /pages

endpoint that will be hit by the UI to create a page, will generate the page and return a randomly generated ID

### POST /messages

endpoint that will be hit by catapult. will receive incoming MMS, and update the page associated with the "to" number in the message with the media in the MMS

### GET /pages/{foo}

endpoint that users will go to to view their page. returns HTML that has javascript in it that will poll for changes

### GET /pages/{foo}/updates

endpoint that the javascript on the page will hit to poll for updates

### GET /pages/{foo}/admin

endpoint that the creator of the page can hit to do things like delete, reset, whitelist/blacklist source TNs, etc. (returns HTML with ui controls to do these things)

### POST /pages/{foo}

endpoint hit by the UI controls to make the changes on behalf of the user (as above)

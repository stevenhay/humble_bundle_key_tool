# Humble Bundle Key Tool

This tool interacts with the undocumented humble bundle API using your session cookie.

## Features

* Lists all games which are unredeeemed
* Lists all games which are expired

## Potential Features

* Automatic game key retrevial using the /redeemkey endpoint.

## Usage

Running the program requires the session cookie from the Humble Bundle store. To get this, login to https://humblebundle.com/, open developer tools and refresh the page. Find the request to www.humblebundle.com or '/?hmb_source=navbar', click it and look for either:

* 'set-cookie' in the response headers
* 'cookie' in the request headers

and take the value from *_simpleauth_sess* without the quotes.

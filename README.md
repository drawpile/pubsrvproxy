# Drawpile Session Public Statistics Proxy

This is a simple proxy server that safely proxies read-only information
from a Drawpile server's web-admin API. It can be used to show live
server information on a public web site without allowing direct access
to the server itself.

## Installation

1. Install the listserver: `go get github.com/drawpile/pubsrvproxy`
2. Write a configuration file (if you want to change any settings)
3. Run the listing server (`$GOPATH/bin/listserver`)

Sample systemd unit file (`/etc/systemd/system/pubsrvproxy.service`):

	[Unit]
	Description=Drawpile statistics proxy
	After=network.target

	[Service]
	ExecStart=/home/website/go/bin/pubsrvproxy -c /home/website/pubsrvproxy.cfg
	User=website

	[Install]
	WantedBy=multi-user.target

## Using with nginx

In your nginx virtual host config, add a proxy pass location like this:

	location /api/ {
		proxy_pass http://127.0.0.1:8080/;
		proxy_redirect default;
	}

## API endpoints

The proxy provides the following JSON API endpoints:

### /sessions/ - Session list

This returns a list of running sessions on this server. The format is
compatible with what the session list server returns, but there are
a couple of extra fields.

Returns:

	[
		{
			"host": "server hostname",     (must be configured in the settings file)
			"port": server port,           (must be configured in the settings file)
			"id": "session ID",            (note: list server returns the alias here, if set)
			"alias": "session alias",      (not included in listserver response)
			"protocol": "protocol version",
			"title": "session title",
			"users": number of users,
			"password": true/false,
			"closed": true/false,          (not included in listserver response)
			"nsfm": true/false,
			"owner": "username",
			"started": "start time",
			"size": history size in bytes  (not included in listserver response)
		},
		...
	]

### /users/ - User list

This view returns a list of users who are logged in to the server.
It is not enabled by default. To enable it, includ `UserView = true` in
your configuration file.

Returns:

	[
		{
			"id": User ID,
			"name": "username",
			"ip": "user IP address", (included only if `ShowUserIps = true` is set)
			"auth": true/false,      (is this an authenticated user?)
			"op": true/false,        (is this user a session operator?)
			"muted": true/false,     (is this user blocked from the chat?)
			"mod": true/false,       (is this user a moderator?)
			"session": "session ID"
		},
		...
	]

The *user ID* field is in range 0-255. User IDs are not unique between sessions
and may be reused even within a session. Users with ID 0 have not yet logged in
to any session.


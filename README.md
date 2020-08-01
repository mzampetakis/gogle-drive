# gogle-drive
A google's drive API consumer layer in go

This package aims to help consuming google's drive API through a layer that provides functionalities to easily use some of the provided capabilities. It connect to Google;s Drive through an OAuth Client.

## Getting Started

Just import this package using 

```go
import github.com/mzampetakis/gogle-drive
```

### First Use

Before running your application you must provide a valid json file with the google's API OAuth client ID.

Head to [https://console.developers.google.com/apis/credentials/oauthclient](https://console.developers.google.com/apis/credentials/oauthclient) and add `Desktop App` as Application type. Then provide a name to distinguish your app and press the Create button.

At the Google's Console [Credentials page](https://console.developers.google.com/apis/credentials) under the `OAuth 2.0 Client IDs` section you must download the client's credentials as a json file and move it into your project.

> Consider adding this file to your .gitignore as well as the token.json that will be generated later on.

On the first run you will be prompted to authenticated your client through a web link. The output at the first run will prompt to visit a link, and after you allow the requested access, paste the provided token.

```
Go to the following link in your browser. 
https://accounts.google.com/o/oauth2/auth?.....

Then type here the authorization code: paste_here_the_provided_token
```

This will produce a token.json in your project's root directory.

### Examples
A siple example that lists your google's drive files is provided in [examples/gogledriveexample.go](https://github.com/mzampetakis/gogle-drive/blob/master/examples/gogledrive_example.go) file. 

To start a new gogledrive's instance simply call:

```go
gdrive, err := gogledrive.New("path_to/credentials.json")
```
Then using the `gdrive` all provided methods are available. 

To Search for files you can use 
```go
filterCriteria := gogledrive.ListFilter{}
assets, err := gdrive.SearchFiles(filterCriteria)
```

The `gogledrive.ListFilter` provides filters for file name, mime type and folderID to search within.
For more details on filtering please refer to the official documentation here: 
[https://developers.google.com/drive/api/v3/reference/query-ref](https://developers.google.com/drive/api/v3/reference/query-ref)

To get a specific file the method `GetFile` is available through

```go
data, err := gdrive.GetFile("fileID")
```

You can review the whole package's documentation by running `go doc -all .` at this package's folder.

## Contributing

You can contribute to this project by just opening a PR or open first an issue.
Please describe thoroughly what are your PR solves or adds.

Some ideas for contribution: 
* upload a file
* create a file
* add more filter criteria
* ...

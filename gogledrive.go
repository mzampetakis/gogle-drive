// Package gogledrive provides some easy to use functionalities about the google drive
// such as searching for assets and downloading files.
// It uses the Google oAuth Client which can be generated from:
// https://console.developers.google.com/apis/credentials
package gogledrive

import (
	"errors"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

type Gogledrive struct {
	srv *drive.Service
}

// New returns a new instance of gogledrive's.
// It uses the provided credentialsFile to connect to Google OAuth service.
func New(credentialsFile string) (*Gogledrive, error) {
	b, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, errors.New("Unable to read client secret file: " + err.Error())
	}
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveReadonlyScope)
	if err != nil {
		return nil, errors.New("Unable to parse client secret file to config: " + err.Error())
	}
	client := getClient(config)

	srv, err := drive.New(client)
	if err != nil {
		return nil, errors.New("Unable to retrieve Drive client: " + err.Error())
	}
	return &Gogledrive{
		srv: srv,
	}, nil
}

// ListFilter provides the available filters for feedind the SearchFiles()
// Use mimeType = 'application/vnd.google-apps.folder' to list folders
type ListFilter struct {
	FolderID *string
	Name     *string
	MimeType *string
}

// SearchFiles returns a map where key is the file name and value is the FileId
// that retrieves as search results based on the provided ListFilter
func (gdrive *Gogledrive) SearchFiles(filters ListFilter) (map[string]string, error) {
	filesList := map[string]string{}
	r, err := gdrive.srv.Files.List().Q(constructFilterQ(filters)).
		Fields("nextPageToken, files(id, name, mimeType, parents)").Do()
	if err != nil {
		return filesList, errors.New("Unable to retrieve files: " + err.Error())
	}
	for _, i := range r.Files {
		filesList[i.Name] = i.Id
	}
	return filesList, nil
}

func constructFilterQ(filters ListFilter) string {
	filterQ := ""
	if filters.FolderID != nil {
		filterQ += "'" + *filters.FolderID + "' in parents and "
	}
	if filters.Name != nil {
		filterQ += "name contains '" + *filters.Name + "' and "
	}
	if filters.MimeType != nil {
		filterQ += "mimeType contains '" + *filters.MimeType + "' and "
	}
	filterQ += "trashed = false"
	return filterQ
}

// GetFile returns a byte[] with the contents of the fileID provided
func (gdrive *Gogledrive) GetFile(fileID string) ([]byte, error) {
	var fileBuf []byte
	r := gdrive.srv.Files.Get(fileID)
	downloadResp, err := r.Download()
	if err != nil {
		return fileBuf, errors.New("Unable to download files: " + err.Error())
	}
	fileBuf, err = ioutil.ReadAll(downloadResp.Body)
	if err != nil {
		return fileBuf, errors.New("Unable to download files: " + err.Error())
	}
	return fileBuf, nil
}

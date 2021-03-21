package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type FileInfo struct {
	Kind                         string            `json:"kind"`
	ID                           string            `json:"id"`
	Etag                         string            `json:"etag"`
	Selflink                     string            `json:"selfLink"`
	Alternatelink                string            `json:"alternateLink"`
	Embedlink                    string            `json:"embedLink"`
	Iconlink                     string            `json:"iconLink"`
	Thumbnaillink                string            `json:"thumbnailLink"`
	Title                        string            `json:"title"`
	Mimetype                     string            `json:"mimeType"`
	Labels                       Labels            `json:"labels"`
	Copyrequireswriterpermission bool              `json:"copyRequiresWriterPermission"`
	Createddate                  time.Time         `json:"createdDate"`
	Modifieddate                 time.Time         `json:"modifiedDate"`
	Markedviewedbymedate         time.Time         `json:"markedViewedByMeDate"`
	Version                      string            `json:"version"`
	Parents                      []interface{}     `json:"parents"`
	Userpermission               Userpermission    `json:"userPermission"`
	Quotabytesused               string            `json:"quotaBytesUsed"`
	Ownernames                   []string          `json:"ownerNames"`
	Owners                       []Owners          `json:"owners"`
	Lastmodifyingusername        string            `json:"lastModifyingUserName"`
	Lastmodifyinguser            Lastmodifyinguser `json:"lastModifyingUser"`
	Capabilities                 Capabilities      `json:"capabilities"`
	Editable                     bool              `json:"editable"`
	Copyable                     bool              `json:"copyable"`
	Writerscanshare              bool              `json:"writersCanShare"`
	Shared                       bool              `json:"shared"`
	Explicitlytrashed            bool              `json:"explicitlyTrashed"`
	Appdatacontents              bool              `json:"appDataContents"`
	Spaces                       []string          `json:"spaces"`
}
type Labels struct {
	Starred    bool `json:"starred"`
	Hidden     bool `json:"hidden"`
	Trashed    bool `json:"trashed"`
	Restricted bool `json:"restricted"`
	Viewed     bool `json:"viewed"`
}
type Userpermission struct {
	Kind            string        `json:"kind"`
	Etag            string        `json:"etag"`
	ID              string        `json:"id"`
	Selflink        string        `json:"selfLink"`
	Role            string        `json:"role"`
	Type            string        `json:"type"`
	Selectableroles []interface{} `json:"selectableRoles"`
}
type Owners struct {
	Kind                string `json:"kind"`
	Displayname         string `json:"displayName"`
	Isauthenticateduser bool   `json:"isAuthenticatedUser"`
	Permissionid        string `json:"permissionId"`
	Emailaddress        string `json:"emailAddress"`
	ID                  string `json:"id"`
}
type Picture struct {
	URL string `json:"url"`
}
type Lastmodifyinguser struct {
	Kind                string  `json:"kind"`
	Displayname         string  `json:"displayName"`
	Picture             Picture `json:"picture"`
	Isauthenticateduser bool    `json:"isAuthenticatedUser"`
	Permissionid        string  `json:"permissionId"`
	ID                  string  `json:"id"`
}
type Capabilities struct {
	Cancopy bool `json:"canCopy"`
	Canedit bool `json:"canEdit"`
}

type ApiError struct {
	Error Error `json:"error"`
}
type Errors struct {
	Domain       string `json:"domain"`
	Reason       string `json:"reason"`
	Message      string `json:"message"`
	Locationtype string `json:"locationType"`
	Location     string `json:"location"`
}
type Error struct {
	Errors  []Errors `json:"errors"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
}

func GetFileInformation(identifier string) (*FileInfo, []byte, error) {
	url := fmt.Sprintf("https://clients6.google.com/drive/v2beta/files/%s?supportsTeamDrives=true&key=AIzaSyC1eQ1xj69IdTMeii5r7brs3R90eck-m7k", identifier)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("X-Origin", "https://drive.google.com")

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, nil, err
	}

	// TODO: Better error handling?
	if res.StatusCode != 200 {
		var apiError ApiError
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, errors.New(apiError.Error.Message)
	}

	var info FileInfo

	err = json.Unmarshal(body, &info)

	if err != nil {
		return nil, nil, err
	}

	return &info, body, err
}

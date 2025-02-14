// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	time "time"

	internal "github.com/hassan-ptpal/vapi-server-sdk-go/internal"
)

type File struct {
	Object *string     `json:"object,omitempty" url:"object,omitempty"`
	Status *FileStatus `json:"status,omitempty" url:"status,omitempty"`
	// This is the name of the file. This is just for your own reference.
	Name         *string                `json:"name,omitempty" url:"name,omitempty"`
	OriginalName *string                `json:"originalName,omitempty" url:"originalName,omitempty"`
	Bytes        *string                `json:"bytes,omitempty" url:"bytes,omitempty"`
	Purpose      *string                `json:"purpose,omitempty" url:"purpose,omitempty"`
	Mimetype     *string                `json:"mimetype,omitempty" url:"mimetype,omitempty"`
	Key          *string                `json:"key,omitempty" url:"key,omitempty"`
	Path         *string                `json:"path,omitempty" url:"path,omitempty"`
	Bucket       *string                `json:"bucket,omitempty" url:"bucket,omitempty"`
	Url          *string                `json:"url,omitempty" url:"url,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	// This is the unique identifier for the file.
	Id string `json:"id" url:"id"`
	// This is the unique identifier for the org that this file belongs to.
	OrgId string `json:"orgId" url:"orgId"`
	// This is the ISO 8601 date-time string of when the file was created.
	CreatedAt time.Time `json:"createdAt" url:"createdAt"`
	// This is the ISO 8601 date-time string of when the file was last updated.
	UpdatedAt time.Time `json:"updatedAt" url:"updatedAt"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (f *File) GetStatus() *FileStatus {
	if f == nil {
		return nil
	}
	return f.Status
}

func (f *File) GetName() *string {
	if f == nil {
		return nil
	}
	return f.Name
}

func (f *File) GetOriginalName() *string {
	if f == nil {
		return nil
	}
	return f.OriginalName
}

func (f *File) GetBytes() *string {
	if f == nil {
		return nil
	}
	return f.Bytes
}

func (f *File) GetPurpose() *string {
	if f == nil {
		return nil
	}
	return f.Purpose
}

func (f *File) GetMimetype() *string {
	if f == nil {
		return nil
	}
	return f.Mimetype
}

func (f *File) GetKey() *string {
	if f == nil {
		return nil
	}
	return f.Key
}

func (f *File) GetPath() *string {
	if f == nil {
		return nil
	}
	return f.Path
}

func (f *File) GetBucket() *string {
	if f == nil {
		return nil
	}
	return f.Bucket
}

func (f *File) GetUrl() *string {
	if f == nil {
		return nil
	}
	return f.Url
}

func (f *File) GetMetadata() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.Metadata
}

func (f *File) GetId() string {
	if f == nil {
		return ""
	}
	return f.Id
}

func (f *File) GetOrgId() string {
	if f == nil {
		return ""
	}
	return f.OrgId
}

func (f *File) GetCreatedAt() time.Time {
	if f == nil {
		return time.Time{}
	}
	return f.CreatedAt
}

func (f *File) GetUpdatedAt() time.Time {
	if f == nil {
		return time.Time{}
	}
	return f.UpdatedAt
}

func (f *File) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *File) UnmarshalJSON(data []byte) error {
	type embed File
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
		UpdatedAt *internal.DateTime `json:"updatedAt"`
	}{
		embed: embed(*f),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*f = File(unmarshaler.embed)
	f.CreatedAt = unmarshaler.CreatedAt.Time()
	f.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *File) MarshalJSON() ([]byte, error) {
	type embed File
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
		UpdatedAt *internal.DateTime `json:"updatedAt"`
	}{
		embed:     embed(*f),
		CreatedAt: internal.NewDateTime(f.CreatedAt),
		UpdatedAt: internal.NewDateTime(f.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (f *File) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FileStatus string

const (
	FileStatusIndexed    FileStatus = "indexed"
	FileStatusNotIndexed FileStatus = "not_indexed"
)

func NewFileStatusFromString(s string) (FileStatus, error) {
	switch s {
	case "indexed":
		return FileStatusIndexed, nil
	case "not_indexed":
		return FileStatusNotIndexed, nil
	}
	var t FileStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (f FileStatus) Ptr() *FileStatus {
	return &f
}

type UpdateFileDto struct {
	// This is the name of the file. This is just for your own reference.
	Name *string `json:"name,omitempty" url:"-"`
}

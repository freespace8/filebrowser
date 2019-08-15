package storage

import (
	"github.com/freespace8/filebrowser/v2/auth"
	"github.com/freespace8/filebrowser/v2/settings"
	"github.com/freespace8/filebrowser/v2/share"
	"github.com/freespace8/filebrowser/v2/users"
)

// Storage is a storage powered by a Backend whih makes the neccessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    *users.Storage
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}

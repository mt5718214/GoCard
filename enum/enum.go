package enum

import (
	"github.com/google/uuid"
)

var (
	Admin AdminUserInfo
)

func init() {
	Admin = AdminUserInfo{
		"00000000-0000-0000-0000-000000000000",
		"Admin",
	}
}

type AdminUserInfo struct {
	ID, Name string
}

func (a AdminUserInfo) AdminUuid() uuid.UUID {
	AdminUserUuid, _ := uuid.Parse(a.ID)
	return AdminUserUuid
}

func (a AdminUserInfo) AdminName() string {
	return a.Name
}

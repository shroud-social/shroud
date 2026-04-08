package member

import (
	"services/internal/domain/realm/guild"
	"services/internal/domain/realm/guild/role"
	"services/internal/domain/realm/user"
	"time"

	"github.com/google/uuid"
)

type JoinMethod string

const (
	JoinMethodInvite       JoinMethod = "INVITE"
	JoinMethodDiscovery    JoinMethod = "DISCOVERY"
	JoinMethodVerification JoinMethod = "VERIFICATION"
)

type ApplicationStatus string

const (
	ApplicationStatusApproved ApplicationStatus = "APPROVED"
	ApplicatioonStatusDenied  ApplicationStatus = "DENIED"
	ApplicationStatusPending  ApplicationStatus = "PENDING"
)

type Member struct {
	GuildID uuid.UUID   `json:"guild_id"`
	RealmID uuid.UUID   `json:"realm_id"`
	UserID  uuid.UUID   `json:"user_id"`
	RoleIDs []uuid.UUID `json:"role_ids"`

	ApplicationStatus ApplicationStatus `json:"application_status"`

	JoinedAt            time.Time  `json:"joined_at"`
	JoinMethod          JoinMethod `json:"join_method"`
	JoinLink            string     `json:"join_link"`
	JoinLinkGeneratedBy uuid.UUID  `json:"join_link_generated_by"`

	GuildProfile user.Profile `json:"profile"`
}

func (m *Member) HasPermission(p guild.Permission) bool {
	// TODO: Logic to get role from DB
	r := role.Role{}
	if r.Permissions&guild.PermissionMask(p) != 0 {
		return true
	}
	return false
}

func (m *Member) HasPermissionInChannel(channel, p guild.Permission) bool {
	// TODO: this
	return false
}

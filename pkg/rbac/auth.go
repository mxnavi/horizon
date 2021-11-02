package rbac

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"g.hz.netease.com/horizon/core/middleware/user"
	"g.hz.netease.com/horizon/pkg/auth"
	"g.hz.netease.com/horizon/pkg/member/models"
	memberservice "g.hz.netease.com/horizon/pkg/member/service"
	"g.hz.netease.com/horizon/pkg/rbac/role"
	"g.hz.netease.com/horizon/pkg/rbac/types"
	"g.hz.netease.com/horizon/pkg/util/log"
)

var (
	ErrMemberNotExist = errors.New("Member Not Exist")
)

// Authorizer use the basic rbac rules to check if the user
// have the permissions
type Authorizer interface {
	Authorize(ctx context.Context, attributes auth.Attributes) (auth.Decision, string, error)
}

type VisitorFunc func(fmt.Stringer, *types.PolicyRule, error) bool

func NewAuthorizer(roleservice role.Service, memberservice memberservice.Service) Authorizer {
	return &authorizer{
		roleService:   roleservice,
		memberService: memberservice,
	}
}

type authorizer struct {
	roleService   role.Service
	memberService memberservice.Service
}

const (
	NotChecked        = "not checked"
	ResourceFormatErr = "format error"
	AnonymousUser     = "anonymous user"
	InternalError     = "internal error"
	MemberNotExist    = "member not exist"
	RoleNotExist      = "role not exist"
	AmdinAllow        = "admin allows everything"
)

func (a *authorizer) Authorize(ctx context.Context, attr auth.Attributes) (auth.Decision,
	string, error) {
	// 0. check (admin allows everything, and some are not checked)
	currentUser, err := user.FromContext(ctx)
	if err != nil {
		return auth.DecisionDeny, AnonymousUser, nil
	}
	if currentUser.IsAdmin() {
		return auth.DecisionAllow, AmdinAllow, nil
	}

	// TODO(tom): members and pipelineruns and environments need to add to auth check
	if attr.IsResourceRequest() && (attr.GetResource() == "members" ||
		attr.GetResource() == "pipelineruns" ||
		attr.GetResource() == "templates" ||
		attr.GetResource() == "environments") {
		log.Warning(ctx,
			"members|templates|environments are not authed yet")
		return auth.DecisionAllow, NotChecked, nil
	}

	// 1. get the member
	resourceIDStr := attr.GetName()
	resourceID, err := strconv.ParseUint(resourceIDStr, 10, 0)
	if err != nil {
		log.Errorf(ctx, "resourceType = %s, resourceID = %s, format error\n", attr.GetResource(), attr.GetName())
		return auth.DecisionDeny, ResourceFormatErr, err
	}

	member, err := a.memberService.GetMemberOfResource(ctx, attr.GetResource(), uint(resourceID))
	if err != nil {
		log.Errorf(ctx, "GetMemberOfResource error, resourceType = %s, resourceID = %s, user = %s\n",
			attr.GetResource(), attr.GetName(), attr.GetUser().String())
		return auth.DecisionDeny, InternalError, err
	}

	// 2. get the role
	var role *types.Role
	if member == nil {
		// TODO(tom): non public resources
		// now default resource are public, if there is a default role, non member can choose the default role
		defaultRole := a.roleService.GetDefaultRole(ctx)
		if defaultRole == nil {
			log.Warningf(ctx, " user %s member and role not found of resourceType = %s, resourceID = %s",
				attr.GetUser().String(), attr.GetResource(), attr.GetName())
			return auth.DecisionDeny, MemberNotExist, nil
		}
		log.WithFiled(ctx, "user",
			attr.GetUser().String()).Debugf(" use the default role %s", defaultRole.Name)
		role = defaultRole
	} else {
		var err error
		role, err = a.roleService.GetRole(ctx, member.Role)
		if err != nil {
			log.Errorf(ctx, "get role file err = %+v", err)
			return auth.DecisionDeny, InternalError, err
		}
	}
	if role == nil {
		return auth.DecisionDeny, RoleNotExist, nil
	}

	// 3. check the permission
	return VisitRoles(member, role, attr)
}

func VisitRoles(member *models.Member, role *types.Role,
	attr auth.Attributes) (_ auth.Decision, reason string, err error) {
	var memberInfo string
	if member != nil {
		memberInfo = member.BaseInfo()
	} else {
		memberInfo = "null"
	}
	for i, rule := range role.PolicyRules {
		if types.RuleAllow(attr, &rule) {
			reason = fmt.Sprintf("user %s allowd by member(%s) by rule[%d]",
				attr.GetUser().String(), memberInfo, i)
			return auth.DecisionAllow, reason, nil
		}
	}
	reason = fmt.Sprintf("user %s denied by member(%s)", attr.GetUser().String(), memberInfo)
	return auth.DecisionDeny, reason, nil
}

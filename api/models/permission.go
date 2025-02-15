package models

import (
	"errors"
)

var (
	ErrInsufficientPermission = errors.New("do not have sufficient permission to perform this action")
)

type Action string
type Resource string
type Role string
type PermissionType string

const (
	Create            Action = "create"
	Read              Action = "read"
	Update            Action = "update"
	Delete            Action = "delete"
	Manage            Action = "manage"
	UpdateStatus      Action = "update_status"
	TransferOwnership Action = "transfer_ownership"
)

const (
	Users      Resource = "Users"
	Deliveries Resource = "Deliveries"
	Recipients Resource = "Recipients"
	Orders     Resource = "Orders"
	Ownership  Resource = "Ownership"
)

const (
	Admin       Role = "ADMIN"
	Owner       Role = "OWNER"
	DeliveryMan Role = "DELIVERY_MAN"
)

const (
	Allow PermissionType = "allow"
	Deny  PermissionType = "deny"
)

type Permission struct {
	Role     Role
	Action   Action
	Resource Resource
	Type     PermissionType
}

var rolePermissions = map[Role][]Permission{
	Owner: {
		{Role: Owner, Action: Manage, Resource: "all", Type: Allow},
	},
	Admin: {
		{Role: Admin, Action: Manage, Resource: "all", Type: Allow},
		{Role: Admin, Action: TransferOwnership, Resource: Ownership, Type: Deny},
		{Role: Admin, Action: UpdateStatus, Resource: Orders, Type: Deny},
	},
	DeliveryMan: {
		{Role: DeliveryMan, Action: Read, Resource: Deliveries, Type: Allow},
		{Role: DeliveryMan, Action: Update, Resource: Deliveries, Type: Allow},
		{Role: DeliveryMan, Action: UpdateStatus, Resource: Orders, Type: Allow},
	},
}

func Can(role Role, action Action, resource Resource) bool {
	permissions, exists := rolePermissions[role]
	if !exists {
		return false
	}

	for _, permission := range permissions {
		if permission.Type == Deny && matchesPermission(permission, action, resource) {
			return false
		}
	}

	for _, permission := range permissions {
		if permission.Type == Allow && matchesPermission(permission, action, resource) {
			return true
		}
	}

	return false
}

func Cannot(role Role, action Action, resource Resource) bool {
	return !Can(role, action, resource)
}

func matchesPermission(permission Permission, action Action, resource Resource) bool {
	return (permission.Resource == "all" || permission.Resource == resource) &&
		(permission.Action == action || permission.Action == Manage)
}

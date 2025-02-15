package models

import (
	"testing"
)

func TestCan(t *testing.T) {
	tests := []struct {
		name     string
		role     Role
		action   Action
		resource Resource
		want     bool
	}{
		// Testes para o Admin
		{
			name:     "Admin pode criar usuários",
			role:     Admin,
			action:   Create,
			resource: Users,
			want:     true,
		},
		{
			name:     "Admin pode ler entregas",
			role:     Admin,
			action:   Read,
			resource: Deliveries,
			want:     true,
		},
		{
			name:     "Admin NÃO pode transferir ownership",
			role:     Admin,
			action:   TransferOwnership,
			resource: Ownership,
			want:     false,
		},
		{
			name:     "Admin NÃO pode atualizar status de um pedido",
			role:     Admin,
			action:   UpdateStatus,
			resource: Orders,
			want:     false,
		},

		// Testes para o Owner
		{
			name:     "Owner pode criar usuários",
			role:     Owner,
			action:   Create,
			resource: Users,
			want:     true,
		},
		{
			name:     "Owner pode transferir ownership",
			role:     Owner,
			action:   TransferOwnership,
			resource: Ownership,
			want:     true,
		},
		{
			name:     "Owner pode deletar destinatários",
			role:     Owner,
			action:   Delete,
			resource: Recipients,
			want:     true,
		},

		// Testes para o DeliveryMan
		{
			name:     "DeliveryMan pode ler entregas",
			role:     DeliveryMan,
			action:   Read,
			resource: Deliveries,
			want:     true,
		},
		{
			name:     "DeliveryMan pode atualizar entregas",
			role:     DeliveryMan,
			action:   Update,
			resource: Deliveries,
			want:     true,
		},
		{
			name:     "DeliveryMan NÃO pode deletar entregas",
			role:     DeliveryMan,
			action:   Delete,
			resource: Deliveries,
			want:     false,
		},
		{
			name:     "DeliveryMan NÃO pode criar usuários",
			role:     DeliveryMan,
			action:   Create,
			resource: Users,
			want:     false,
		},
		{
			name:     "DeliveryMan NÃO pode transferir ownership",
			role:     DeliveryMan,
			action:   TransferOwnership,
			resource: Ownership,
			want:     false,
		},
		{
			name:     "DeliveryMan pode atualizar status de um pedido",
			role:     DeliveryMan,
			action:   UpdateStatus,
			resource: Orders,
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Can(tt.role, tt.action, tt.resource)
			if got != tt.want {
				t.Errorf("Can(%v, %v, %v) = %v, want %v", tt.role, tt.action, tt.resource, got, tt.want)
			}
		})
	}
}

func TestCannot(t *testing.T) {
	tests := []struct {
		name     string
		role     Role
		action   Action
		resource Resource
		want     bool
	}{
		// Testes para o Admin
		{
			name:     "Admin NÃO pode transferir ownership",
			role:     Admin,
			action:   TransferOwnership,
			resource: Ownership,
			want:     true,
		},

		// Testes para o DeliveryMan
		{
			name:     "DeliveryMan NÃO pode deletar entregas",
			role:     DeliveryMan,
			action:   Delete,
			resource: Deliveries,
			want:     true,
		},
		{
			name:     "DeliveryMan NÃO pode criar usuários",
			role:     DeliveryMan,
			action:   Create,
			resource: Users,
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cannot(tt.role, tt.action, tt.resource)
			if got != tt.want {
				t.Errorf("Cannot(%v, %v, %v) = %v, want %v", tt.role, tt.action, tt.resource, got, tt.want)
			}
		})
	}
}

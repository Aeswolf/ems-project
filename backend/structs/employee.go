package structs

import (
	"time"

	"github.com/google/uuid"
)

type employeeProps struct {
	ID                   uuid.UUID `json:"id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Email                string    `json:"email"`
	IsDeleted            bool      `json:"is_deleted"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
	AssociatedLaboratory string    `json:"associated_laboratory"`
}

type employeeOptions func(*Employee)

// struct for Employee
type Employee struct {
	employeeProps
}

func NewEmployee(opts ...employeeOptions) *Employee {

	e := &Employee{}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func WithEmployeeCreationTime() employeeOptions {
	return func(e *Employee) {
		e.CreatedAt = time.Now()
	}
}

func WithEmployeeDeletionTime(t time.Time) employeeOptions {
	return func(e *Employee) {
		e.DeletedAt = t
	}
}

func WithEmployeeIsDeleted() employeeOptions {
	return func(e *Employee) {
		e.IsDeleted = true
	}
}

func WithEmployeeFirstName(n string) employeeOptions {
	return func(e *Employee) {
		if e.FirstName != n {
			e.FirstName = n
		}
	}
}

func WithEmployeeLastName(n string) employeeOptions {
	return func(e *Employee) {
		if e.LastName != n {
			e.LastName = n
		}
	}
}

func WithEmployeeEmail(s string) employeeOptions {
	return func(e *Employee) {
		if e.Email != s {
			e.Email = s
		}
	}
}

func WithEmployeeUpdatedTime() employeeOptions {
	return func(d *Employee) {
		d.UpdatedAt = time.Now()
	}
}

func WithEmployeeID() employeeOptions {
	return func(e *Employee) {
		e.ID = uuid.New()
	}
}

func WithEmployeeAssociatedLaboratoryName(n string) employeeOptions {
	return func(e *Employee) {
		if e.AssociatedLaboratory != n {
			e.AssociatedLaboratory = n
		}
	}
}

func (e *Employee) UpdateEmployee(opts ...employeeOptions) {
	for _, opt := range opts {
		opt(e)
	}
}

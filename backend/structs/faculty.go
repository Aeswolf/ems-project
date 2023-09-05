package structs

import (
	"time"

	"github.com/google/uuid"
)

type facultyProps struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type facultyOptions func(*Faculty)

// struct for faculty
type Faculty struct {
	facultyProps
}

func defaultFacultyProps() *Faculty {
	props := facultyProps{
		IsDeleted: false,
	}

	return &Faculty{
		facultyProps: props,
	}
}

func NewFaculty(opts ...facultyOptions) *Faculty {

	d := defaultFacultyProps()

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func WithFacultyCreationTime() facultyOptions {
	return func(f *Faculty) {
		f.CreatedAt = time.Now()
	}
}

func WithFacultyDeletionTime(t time.Time) facultyOptions {
	return func(f *Faculty) {
		f.DeletedAt = t
	}
}

func WithFacultyIsDeleted() facultyOptions {
	return func(f *Faculty) {
		f.IsDeleted = true
	}
}

func WithFacultyName(n string) facultyOptions {
	return func(f *Faculty) {
		if f.Name != n {
			f.Name = n
		}
	}
}

func WithFacultyUpdatedTime() facultyOptions {
	return func(f *Faculty) {
		f.UpdatedAt = time.Now()
	}
}

func WithFacultyID() facultyOptions {
	return func(f *Faculty) {
		f.ID = uuid.New()
	}
}

func (f *Faculty) UpdateFaculty(opts ...facultyOptions) {
	for _, opt := range opts {
		opt(f)
	}
}

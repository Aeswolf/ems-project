package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	s "github.com/Aeswolf/equipment-database-management/structs"
)

// struct for the postgres store
type PostgresStore struct {
	db *sql.DB
}

// constructor for new instances of the store
func New(dsn string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

// function to add a new faculty to the database
func (p *PostgresStore) AddFaculty(f *s.Faculty) (string, error) {
	query := `INSERT INTO
				faculty(
					faculty_id,
					name,
					created_at,
					updated_at,
					deleted_at,
					is_deleted
				)
					VALUES($1, $2, $3, $4, $5, $6)
						RETURNING faculty_id;
			`
	var id string

	err := p.db.QueryRow(
		query,
		f.ID,
		f.Name,
		f.CreatedAt,
		f.UpdatedAt,
		f.DeletedAt,
		f.IsDeleted,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// function to get a faculty by the id
func (p *PostgresStore) GetFaculty(id string) (*s.Faculty, error) {
	query := `SELECT *
				FROM faculty
					WHERE faculty_id=$1 and is_deleted=$2
						ORDER BY faculty_id

			`

	f := new(s.Faculty)

	err := p.db.QueryRow(
		query,
		id,
		false).Scan(
		&f.ID,
		&f.Name,
		&f.CreatedAt,
		&f.UpdatedAt,
		&f.DeletedAt,
		&f.IsDeleted,
	)

	if err != nil {
		return nil, err
	}

	return f, nil
}

// function to get all the faculties
func (p *PostgresStore) GetFaculties() ([]*s.Faculty, error) {
	query := "SELECT * FROM faculty WHERE is_deleted=$1"

	var faculties []*s.Faculty

	rows, err := p.db.Query(query, false)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		f := new(s.Faculty)

		err = rows.Scan(
			&f.ID,
			&f.Name,
			&f.CreatedAt,
			&f.UpdatedAt,
			&f.DeletedAt,
			&f.IsDeleted,
		)

		if err != nil {
			return nil, err
		}

		faculties = append(faculties, f)
	}

	return faculties, nil
}

// function to delete a faculty from the database
func (p *PostgresStore) RemoveFaculty(id string) error {
	f, err := p.GetFaculty(id)

	if err != nil {
		return err
	}

	f.UpdateFaculty(
		s.WithFacultyIsDeleted(),
		s.WithFacultyDeletionTime(time.Now()),
	)

	tx, err := p.db.Begin()

	if err != nil {
		return err
	}

	fUpdateQuery := `
					UPDATE faculty
					SET
						is_deleted=$1,
						deleted_at=$2
							WHERE faculty_id=$3;
	`

	_, err = tx.Exec(fUpdateQuery, f.IsDeleted, f.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	dUpdateQuery := `
					UPDATE department
					SET
						is_deleted=$1,
						deleted_at=$2
							WHERE faculty_id=$3;
					`
	_, err = tx.Exec(dUpdateQuery, f.IsDeleted, f.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	lUpdateQuery := `
					UPDATE laboratory
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE department_id IN
									(SELECT department_id FROM department
										WHERE faculty_id = $3)
					`

	_, err = tx.Exec(lUpdateQuery, f.IsDeleted, f.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	eUpdateQuery := `
					UPDATE equipment
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE laboratory_id IN
									(
										SELECT laboratory_id FROM laboratory
											WHERE department_id IN
												(
													SELECT department_id FROM department
														WHERE faculty_id=$3
												)
									);
					`

	_, err = tx.Exec(eUpdateQuery, f.IsDeleted, f.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	empUpdateQuery := `
					UPDATE employee
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE laboratory_id IN
									(
										SELECT laboratory_id FROM laboratory
											WHERE department_id IN
												(
													SELECT department_id FROM department
														WHERE faculty_id=$3
												)
									);
					`

	_, err = tx.Exec(empUpdateQuery, f.IsDeleted, f.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	err = tx.Commit()

	return err
}

// function to update the information of a faculty in the database
func (p *PostgresStore) UpdateFaculty(id string, u *s.Faculty) error {
	o, err := p.GetFaculty(id)

	if err != nil {
		return err
	}

	o.UpdateFaculty(s.WithFacultyName(u.Name), s.WithFacultyUpdatedTime())

	query := `UPDATE
				faculty
					SET
						name=$1,
						updated_at=$2
							WHERE faculty_id=$3`

	_, err = p.db.Exec(query, o.Name, o.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil
}

// function to add a new department to the database
func (p *PostgresStore) AddDepartment(d *s.Department) (string, error) {

	query := `
			INSERT INTO department (
				department_id,
				name,
				created_at,
				updated_at,
				deleted_at,
				is_deleted,
				faculty_id
			)
				SELECT $1, $2, $3, $4, $5, $6, faculty_id
					FROM faculty
						WHERE name = $7 and is_deleted = $8
							RETURNING department_id;
			`
	var id string

	err := p.db.QueryRow(
		query,
		d.ID,
		d.Name,
		d.CreatedAt,
		d.UpdatedAt,
		d.DeletedAt,
		d.IsDeleted,
		d.AssociatedFaculty,
		false,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// function to get all the departments in the database
func (p *PostgresStore) GetDepartments() ([]*s.Department, error) {
	query := `SELECT
				department.department_id,
				department.name,
				department.created_at,
				department.updated_at,
				department.deleted_at,
				department.is_deleted,
				faculty.name
					FROM department
						INNER JOIN faculty
							ON department.faculty_id = faculty.faculty_id
								WHERE department.is_deleted=$1`

	rows, err := p.db.Query(query, false)

	if err != nil {
		return nil, err
	}

	var departments []*s.Department

	for rows.Next() {
		d := new(s.Department)

		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.CreatedAt,
			&d.UpdatedAt,
			&d.DeletedAt,
			&d.IsDeleted,
			&d.AssociatedFaculty,
		)

		if err != nil {
			return nil, err
		}

		departments = append(departments, d)
	}

	return departments, nil
}

// function to get a single department from the database
func (p *PostgresStore) GetDepartment(id string) (*s.Department, error) {
	query := `SELECT
				department.department_id,
				department.name,
				department.created_at,
				department.updated_at,
				department.deleted_at,
				department.is_deleted,
				faculty.name
					FROM department
						INNER JOIN faculty
							ON department.faculty_id = faculty.faculty_id
								WHERE department.department_id=$1
										AND department.is_deleted=$2`
	d := new(s.Department)
	row := p.db.QueryRow(query, id, false)
	err := row.Scan(
		&d.ID,
		&d.Name,
		&d.CreatedAt,
		&d.UpdatedAt,
		&d.DeletedAt,
		&d.IsDeleted,
		&d.AssociatedFaculty,
	)

	if err != nil {
		return nil, err
	}

	return d, nil
}

// function to delete a department
func (p *PostgresStore) RemoveDepartment(id string) error {
	d, err := p.GetDepartment(id)

	if err != nil {
		return err
	}

	d.UpdateDepartment(s.WithDepartmentDeletionTime(time.Now()), s.WithDepartmentIsDeleted())

	tx, err := p.db.Begin()

	if err != nil {
		return err
	}

	dUpdateQuery := `
					UPDATE department
					SET
						is_deleted=$1,
						deleted_at=$2
							WHERE department_id=$3;
					`
	_, err = tx.Exec(dUpdateQuery, d.IsDeleted, d.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	lUpdateQuery := `
					UPDATE laboratory
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE department_id = $3
					`

	_, err = tx.Exec(lUpdateQuery, d.IsDeleted, d.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	eUpdateQuery := `
					UPDATE equipment
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE laboratory_id IN
									(
										SELECT laboratory_id FROM laboratory
											WHERE department_id = $3
									);
					`

	_, err = tx.Exec(eUpdateQuery, d.IsDeleted, d.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	empUpdateQuery := `
					UPDATE employee
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE laboratory_id IN
									(
										SELECT laboratory_id FROM laboratory
											WHERE department_id = $3
									);
					`

	_, err = tx.Exec(empUpdateQuery, d.IsDeleted, d.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	err = tx.Commit()

	return err
}

// function to update a department in the database
func (p *PostgresStore) UpdateDepartment(id string, d *s.Department) error {
	o, err := p.GetDepartment(id)

	if err != nil {
		return err
	}

	o.UpdateDepartment(
		s.WithAssociatedFacultyName(d.AssociatedFaculty),
		s.WithDepartmentName(d.Name),
		s.WithDepartmentUpdatedTime(),
	)

	query := `UPDATE department
				SET
					name=$1,
					updated_at=$2,
					faculty_id=(
						SELECT faculty_id FROM faculty WHERE name=$3
						)
						WHERE department_id=$4`
	_, err = p.db.Exec(query, o.Name, o.UpdatedAt, o.AssociatedFaculty, id)

	return err
}

// function to add a new lab to the database
func (p *PostgresStore) AddLab(l *s.Lab) (string, error) {

	query := `
			INSERT INTO laboratory (
				laboratory_id,
				name,
				created_at,
				updated_at,
				deleted_at,
				is_deleted,
				department_id
			)
				SELECT $1, $2, $3, $4, $5, $6, department_id
					FROM department
						WHERE name = $7 and is_deleted = $8
							RETURNING laboratory_id;
			`
	var id string

	err := p.db.QueryRow(
		query,
		l.ID,
		l.Name,
		l.CreatedAt,
		l.UpdatedAt,
		l.DeletedAt,
		l.IsDeleted,
		l.AssociatedDepartment,
		false,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// function to get all the departments in the database
func (p *PostgresStore) GetLabs() ([]*s.Lab, error) {
	query := `SELECT
				laboratory.laboratory_id,
				laboratory.name,
				laboratory.created_at,
				laboratory.updated_at,
				laboratory.deleted_at,
				laboratory.is_deleted,
				department.name
					FROM laboratory
						INNER JOIN department
							ON laboratory.department_id = department.department_id
								WHERE laboratory.is_deleted=$1`

	rows, err := p.db.Query(query, false)

	if err != nil {
		return nil, err
	}

	var equipments []*s.Lab

	for rows.Next() {
		l := new(s.Lab)

		err := rows.Scan(
			&l.ID,
			&l.Name,
			&l.CreatedAt,
			&l.UpdatedAt,
			&l.DeletedAt,
			&l.IsDeleted,
			&l.AssociatedDepartment,
		)

		if err != nil {
			return nil, err
		}

		equipments = append(equipments, l)
	}

	return equipments, nil
}

// function to get a single lab from the database
func (p *PostgresStore) GetLab(id string) (*s.Lab, error) {
	query := `SELECT
				laboratory.laboratory_id,
				laboratory.name,
				laboratory.created_at,
				laboratory.updated_at,
				laboratory.deleted_at,
				laboratory.is_deleted,
				department.name
					FROM laboratory
						INNER JOIN department
							ON laboratory.department_id = department.department_id
								WHERE laboratory.laboratory_id=$1
										AND laboratory.is_deleted=$2`
	l := new(s.Lab)
	row := p.db.QueryRow(query, id, false)
	err := row.Scan(
		&l.ID,
		&l.Name,
		&l.CreatedAt,
		&l.UpdatedAt,
		&l.DeletedAt,
		&l.IsDeleted,
		&l.AssociatedDepartment,
	)

	if err != nil {
		return nil, err
	}

	return l, nil
}

// function to delete a laboratory
func (p *PostgresStore) RemoveLaboratory(id string) error {
	l, err := p.GetLab(id)

	if err != nil {
		return err
	}

	l.UpdateLab(
		s.WithLabDeletionTime(time.Now()),
		s.WithLabIsDeleted(),
	)

	tx, err := p.db.Begin()

	if err != nil {
		return err
	}

	lUpdateQuery := `
					UPDATE laboratory
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE laboratory_id = $3
					`

	_, err = tx.Exec(lUpdateQuery, l.IsDeleted, l.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	eUpdateQuery := `
					UPDATE equipment
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE laboratory_id = $3
					`

	_, err = tx.Exec(eUpdateQuery, l.IsDeleted, l.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	empUpdateQuery := `
					UPDATE employee
						SET
							is_deleted=$1,
							deleted_at=$2
								WHERE laboratory_id = $3
					`

	_, err = tx.Exec(empUpdateQuery, l.IsDeleted, l.DeletedAt, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	err = tx.Commit()

	return err
}

// function to update a laboratory in the database
func (p *PostgresStore) UpdateLab(id string, l *s.Lab) error {
	o, err := p.GetLab(id)

	if err != nil {
		return err
	}

	o.UpdateLab(
		s.WithAssociatedDeptName(l.AssociatedDepartment),
		s.WithLabName(l.Name),
		s.WithLabUpdatedTime(),
	)

	query := `
			UPDATE laboratory
				SET
					name=$1,
					updated_at=$2,
					department_id=(
						SELECT
							department_id FROM department
								WHERE name=$3
						)
						WHERE laboratory_id=$4`

	_, err = p.db.Exec(
		query,
		o.Name,
		o.UpdatedAt,
		o.AssociatedDepartment,
		id,
	)

	return err
}

// function to add a new equipment to the database
func (p *PostgresStore) AddEquipment(e *s.Equipment) (string, error) {

	query := `
			INSERT INTO equipment (
				equipment_id,
				name,
				created_at,
				updated_at,
				deleted_at,
				is_deleted,
				quantity,
				laboratory_id
			)
				SELECT $1, $2, $3, $4, $5, $6, $7, laboratory_id
					FROM laboratory
						WHERE name = $8 and is_deleted = $9
							RETURNING equipment_id;
			`
	var id string

	err := p.db.QueryRow(
		query,
		e.ID,
		e.Name,
		e.CreatedAt,
		e.UpdatedAt,
		e.DeletedAt,
		e.IsDeleted,
		e.Quantity,
		e.AssociatedLaboratory,
		false,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// function to get all the departments in the database
func (p *PostgresStore) GetEquipments() ([]*s.Equipment, error) {
	query := `SELECT
				equipment.equipment_id,
				equipment.name,
				equipment.created_at,
				equipment.updated_at,
				equipment.deleted_at,
				equipment.is_deleted,
				laboratory.name,
				equipment.quantity
					FROM equipment
						INNER JOIN laboratory
							ON equipment.laboratory_id = laboratory.laboratory_id
								WHERE equipment.is_deleted=$1`

	rows, err := p.db.Query(query, false)

	if err != nil {
		return nil, err
	}

	var equipments []*s.Equipment

	for rows.Next() {
		e := new(s.Equipment)

		err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.CreatedAt,
			&e.UpdatedAt,
			&e.DeletedAt,
			&e.IsDeleted,
			&e.AssociatedLaboratory,
			&e.Quantity,
		)

		if err != nil {
			return nil, err
		}

		equipments = append(equipments, e)
	}

	return equipments, nil
}

// function to get a single lab from the database
func (p *PostgresStore) GetEquipment(id string) (*s.Equipment, error) {
	query := `SELECT
				equipment.equipment_id,
				equipment.name,
				equipment.created_at,
				equipment.updated_at,
				equipment.deleted_at,
				equipment.is_deleted,
				laboratory.name,
				equipment.quantity
					FROM equipment
						INNER JOIN laboratory
							ON equipment.laboratory_id = laboratory.laboratory_id
								WHERE equipment.equipment_id=$1
										AND equipment.is_deleted=$2`
	e := new(s.Equipment)
	row := p.db.QueryRow(query, id, false)
	err := row.Scan(
		&e.ID,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.DeletedAt,
		&e.IsDeleted,
		&e.AssociatedLaboratory,
		&e.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return e, nil
}

// function to delete an equipment in the database
func (p *PostgresStore) RemoveEquipment(id string) error {
	e, err := p.GetEquipment(id)

	if err != nil {
		return err
	}

	e.UpdateEquipment(
		s.WithEquipmentDeletionTime(time.Now()),
		s.WithEquipmentIsDeleted(),
	)

	query := `
			UPDATE equipment
				SET
					deleted_at = $1,
					is_deleted = $2
						WHERE equipment_id = $3
	`

	_, err = p.db.Exec(query, e.DeletedAt, e.IsDeleted, id)

	return err
}

// function to update a laboratory in the database
func (p *PostgresStore) UpdateEquipment(id string, e *s.Equipment) error {
	o, err := p.GetEquipment(id)

	if err != nil {
		return err
	}

	o.UpdateEquipment(
		s.WithAssociatedLaboratoryName(e.AssociatedLaboratory),
		s.WithEquipmentName(e.Name),
		s.WithEquipmentUpdatedTime(),
		s.WithQuantity(e.Quantity),
	)

	query := `
			UPDATE equipment
				SET
					name=$1,
					quantity=$2,
					updated_at=$3,
					laboratory_id=(
						SELECT
							laboratory_id FROM laboratory
								WHERE name=$4
						)
						WHERE equipment_id=$5`

	_, err = p.db.Exec(
		query,
		o.Name,
		o.Quantity,
		o.UpdatedAt,
		o.AssociatedLaboratory,
		id,
	)

	return err
}

// function to add a new equipment to the database
func (p *PostgresStore) AddEmployee(e *s.Employee) (string, error) {

	query := `
			INSERT INTO employee (
				employee_id,
				first_name,
				last_name,
				email,
				created_at,
				updated_at,
				deleted_at,
				is_deleted,
				laboratory_id
			)
				SELECT $1, $2, $3, $4, $5, $6, $7, $8, laboratory_id
					FROM laboratory
						WHERE name = $9 and is_deleted = $10
							RETURNING employee_id;
			`
	var id string

	err := p.db.QueryRow(
		query,
		e.ID,
		e.FirstName,
		e.LastName,
		e.Email,
		e.CreatedAt,
		e.UpdatedAt,
		e.DeletedAt,
		e.IsDeleted,
		e.AssociatedLaboratory,
		false,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// function to get all the employees in the database
func (p *PostgresStore) GetEmployees() ([]*s.Employee, error) {
	query := `SELECT
				employee.employee_id,
				employee.first_name,
				employee.created_at,
				employee.updated_at,
				employee.deleted_at,
				employee.is_deleted,
				laboratory.name,
				employee.last_name,
				employee.email
					FROM employee
						INNER JOIN laboratory
							ON employee.laboratory_id = laboratory.laboratory_id
								WHERE employee.is_deleted=$1`

	rows, err := p.db.Query(query, false)

	if err != nil {
		return nil, err
	}

	var employees []*s.Employee

	for rows.Next() {
		emp := new(s.Employee)

		err := rows.Scan(
			&emp.ID,
			&emp.FirstName,
			&emp.CreatedAt,
			&emp.UpdatedAt,
			&emp.DeletedAt,
			&emp.IsDeleted,
			&emp.AssociatedLaboratory,
			&emp.LastName,
			&emp.Email,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

// function to get a single lab from the database
func (p *PostgresStore) GetEmployee(id string) (*s.Employee, error) {
	query := `SELECT
				employee.employee_id,
				employee.first_name,
				employee.created_at,
				employee.updated_at,
				employee.deleted_at,
				employee.is_deleted,
				laboratory.name,
				employee.last_name,
				employee.email
					FROM employee
						INNER JOIN laboratory
							ON employee.laboratory_id = laboratory.laboratory_id
								WHERE employee.employee_id=$1
										AND employee.is_deleted=$2`
	emp := new(s.Employee)
	row := p.db.QueryRow(query, id, false)
	err := row.Scan(
		&emp.ID,
		&emp.FirstName,
		&emp.CreatedAt,
		&emp.UpdatedAt,
		&emp.DeletedAt,
		&emp.IsDeleted,
		&emp.AssociatedLaboratory,
		&emp.LastName,
		&emp.Email,
	)

	if err != nil {
		return nil, err
	}

	return emp, nil
}

// function to delete an employee in the database
func (p *PostgresStore) RemoveEmployee(id string) error {
	emp, err := p.GetEmployee(id)

	if err != nil {
		return err
	}

	emp.UpdateEmployee(
		s.WithEmployeeDeletionTime(time.Now()),
		s.WithEmployeeIsDeleted(),
	)

	query := `
			UPDATE employee
				SET
					deleted_at = $1,
					is_deleted = $2
						WHERE employee_id = $3
	`

	_, err = p.db.Exec(query, emp.DeletedAt, emp.IsDeleted, id)

	return err
}

// function to update an employee in the database
func (p *PostgresStore) UpdateEmployee(id string, emp *s.Employee) error {
	o, err := p.GetEmployee(id)

	if err != nil {
		return err
	}

	o.UpdateEmployee(
		s.WithEmployeeAssociatedLaboratoryName(emp.AssociatedLaboratory),
		s.WithEmployeeFirstName(emp.FirstName),
		s.WithEmployeeLastName(emp.LastName),
		s.WithEmployeeEmail(emp.Email),
		s.WithEmployeeUpdatedTime(),
	)

	query := `
			UPDATE employee
				SET
					first_name=$1,
					last_name=$2,
					email=$3,
					updated_at=$4,
					laboratory_id=(
						SELECT
							laboratory_id FROM laboratory
								WHERE name=$5
						)
						WHERE employee_id=$6`

	_, err = p.db.Exec(
		query,
		o.FirstName,
		o.LastName,
		o.Email,
		o.UpdatedAt,
		o.AssociatedLaboratory,
		id,
	)

	return err
}

import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { Spinner } from 'react-bootstrap'
import { useNavigate, Link } from 'react-router-dom'


function EmployeeList() {
    const [isLoading, setIsLoading] = useState(true)
    const [employees, setEmployees] = useState([])
    const navigate = useNavigate()

    useEffect(() => {
        async function fetchEmployees() {
            const res = await axios.get("http://localhost:4000/employee")

            if (res.status === 200) {
                setIsLoading(false)

                res.data ? setEmployees(res.data) : navigate("/employee/create")
            }
        }

        fetchEmployees()
    }, [])

    const handleDeleteEmployee = async (id) => {
        const confirm = window.confirm("Do you want to delete?")

        if (confirm) {
            const res = await axios.delete(`http://localhost:4000/employee/${id}`)

            if (res.status === 200) {
                alert(`Employee with id ${id} has been ${res.data.message}`)
                location.reload()
            }
        }



    }
    return (
        <div className='d-flex flex-column justify-content-center align-items-center vh-100 bg-light'>
            {
                isLoading ? (
                    <Spinner
                        variant="dark"
                        animation="border"
                        aria-hidden="true"
                    />
                ) : (
                    <div className='bg-white w-75 shadow border p-4'>
                        <h1 className='text-center'>List of Employees</h1>
                        <div className="d-flex justify-content-between">
                            <Link to='/employee/create' className='btn btn-outline-success my-4'>Add Employee</Link>
                            <Link to='/' className='btn btn-outline-dark my-4'>Back</Link>
                        </div>
                        <table className='table table-hover text-center'>
                            <thead>
                                <tr>
                                    <th>No.</th>
                                    <th>FirstName</th>
                                    <th>LastName</th>
                                    <th> Email </th>
                                    <th> Associated Laboratory </th>
                                    <th> Actions </th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    employees.map((emp, i) => (
                                        <tr key={i}>
                                            <td>{i + 1}</td>
                                            <td>{emp.first_name}</td>
                                            <td>{emp.last_name}</td>
                                            <td>{emp.email}</td>
                                            <td>{emp.associated_laboratory}</td>
                                            <td>
                                                <Link to={`/employee/${emp.id}/edit`} className='btn btn-outline-primary me-3'>Edit</Link>
                                                <button onClick={() => handleDeleteEmployee(emp.id)} className='btn btn-outline-danger me-3'>Delete</button>
                                                <Link to={`/employee/${emp.id}/details`} className='btn btn-outline-dark'>Details</Link>
                                            </td>
                                        </tr>
                                    ))
                                }
                            </tbody>
                        </table>
                    </div>
                )
            }
        </div>
    )
}

export default EmployeeList
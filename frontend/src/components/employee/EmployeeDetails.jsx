import React, { useState, useEffect } from 'react'
import { Spinner } from 'react-bootstrap'
import { Link, useParams, useNavigate } from 'react-router-dom'
import axios from 'axios'

function EmployeeDetails() {
    const [employeeInfo, setEmployeeInfo] = useState({})
    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const { id } = useParams()
    const navigate = useNavigate()

    useEffect(() => {
        async function fetchEmployee() {
            const res = await axios.get(`http://localhost:4000/employee/${id}`)

            if (res.status === 200) {
                setIsInfoLoaded(true)

                res.data ? setEmployeeInfo(res.data) : navigate("/employee")
            }
        }

        fetchEmployee()
    })
    return (
        <div className='d-flex w-100 vh-100 justify-content-center align-items-center ftlex-column bg-light'>
            {
                isInfoLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h3 className='text-center mb-3'>Employee Details</h3>
                        <div className="row mb-3">
                            <h3> FirstName : <span className='mx-3'>{employeeInfo.first_name} </span></h3>
                        </div>
                        <div className="row mb-3">
                            <h3 className='mb-3'>LastName: <span className='mx-3'>{employeeInfo.last_name}</span></h3>
                        </div>
                        <div className="row mb-3">
                            <h3 className='mb-3'>Email: <span className='mx-3'>{employeeInfo.email}</span></h3>
                        </div>
                        <div className="row mb-3">
                            <h3 className='mb-3'>Laboratory: <span className='mx-3'>{employeeInfo.associated_laboratory}</span></h3>
                        </div>
                        <div className='col-aut+o d-flex justify-content-center align-items-center my-4 flex-column'>
                            <Link to={`/employee/${id}/edit`} className='btn btn-outline-primary mb-3 w-25'>Edit</Link>
                            <Link to="/employee" className='btn btn-outline-dark w-25'>Back</Link>
                        </div>
                    </div>
                ) : (
                    <div>
                        <Spinner
                            variant="dark"
                            role="status"
                            animation="border"
                            aria-hidden="true"
                        />
                    </div>
                )
            }
        </div>
    )
}

export default EmployeeDetails
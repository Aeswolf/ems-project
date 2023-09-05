import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'


function EmployeeEdit() {
    const [oldEmployeeInfo, setOldEmployeeInfo] = useState({})
    const [employeeInfo, setEmployeeInfo] = useState({})
    const [isFormLoaded, setIsFormLoaded] = useState(false)

    const navigate = useNavigate()
    const { id } = useParams()

    const handleInputChange = e => {
        const { name, value } = e.target
        setEmployeeInfo({ ...employeeInfo, [name]: value })
    }

    const handleFormSubmit = async (e) => {
        e.preventDefault();

        if (employeeInfo.first_name !== oldEmployeeInfo.first_name || employeeInfo.last_name !== oldEmployeeInfo.last_name || employeeInfo.email !== oldEmployeeInfo.email || employeeInfo.associated_laboratory !== oldEmployeeInfo.associated_laboratory) {
            const res = await axios.put(`http://localhost:4000/employee/${id}`, employeeInfo)

            if (res.status === 200) {
                alert(`Employee with id ${id} has been ${res.data["message"]}`)
            }

            navigate("/employee")
            return
        }

        location.reload()
    }

    useEffect(() => {
        async function fetchEmployee() {
            const res = await axios.get(`http://localhost:4000/employee/${id}`)

            if (res.status === 200) {
                setIsFormLoaded(true)
                setEmployeeInfo(res.data)
                setOldEmployeeInfo(res.data)
            }
        }

        fetchEmployee()
    }, [])


    return (
        <>
            {

                <div className='d-flex flex-column justify-content-center align-items-center vh-50 bg-light'>
                    {
                        isFormLoaded ? (
                            <div className='w-50 my-3 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                                <h3 className='text-center'>Add Employee</h3>
                                <form method='POST' onSubmit={handleFormSubmit} >
                                    <div className='my-3 row'>
                                        <label htmlFor='first_name' className='mb-1 fw-bold fs-4 mx-3'>FirstName</label>
                                        <div className='col-sm-12'>
                                            <input required className='form-control form-control-lg' value={employeeInfo.first_name} placeholder='Enter firstname' name='first_name' type='text' id='first_name' onChange={handleInputChange} />
                                        </div>
                                    </div>
                                    <div className='my-3 row'>
                                        <label htmlFor='last_name' className='mb-1 fw-bold fs-4 mx-3'>LastName</label>
                                        <div className='col-sm-12'>
                                            <input required className='form-control form-control-lg' value={employeeInfo.last_name} placeholder='Enter lastname' name='last_name' type='text' id='last_name' onChange={handleInputChange} />
                                        </div>
                                    </div>
                                    <div className='my-3 row'>
                                        <label htmlFor='email' className='mb-2 fw-bold fs-4 mx-3'>Email</label>
                                        <div className='col-sm-12'>
                                            <input required className='form-control form-control-lg' value={employeeInfo.email} placeholder='Enter email' name='email' type='text' id='email' onChange={handleInputChange} />
                                        </div>
                                    </div>
                                    <div className='my-3 row'>
                                        <label htmlFor='associated_laboratory' className='mb-2 fw-bold fs-4 mx-3'>Associated Laboratory</label>
                                        <div className='col-sm-12'>
                                            <input required className='form-control form-control-lg' value={employeeInfo.associated_laboratory} placeholder='Enter name of laboratory' name='associated_laboratory' type='text' id='associated_laboratory' onChange={handleInputChange} />
                                        </div>
                                    </div>
                                    <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                                        <button type='submit' className='btn btn-outline-success mb-3 w-25'>Submit</button>
                                        <Link to="/employee" className='btn btn-outline-dark w-25'>Back</Link>
                                    </div>
                                </form>
                            </div>
                        ) : (
                            <Spinner
                                variant="dark"
                                animation="border"
                                role="status"
                            />
                        )
                    }
                </div>
            }
        </>
    )
}

export default EmployeeEdit
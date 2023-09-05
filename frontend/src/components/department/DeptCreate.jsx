import React, { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import axios from 'axios'
import { Spinner } from 'react-bootstrap'

function DeptCreate() {
    const initialDetails = {
        name: "",
        associated_faculty: ""
    }

    const [deptInfo, setDeptInfo] = useState(initialDetails)
    const [isFormLoaded, setIsFormLoaded] = useState(false)
    const navigate = useNavigate()

    const handleFormSubmit = async (e) => {
        e.preventDefault()

        const res = await axios.post("http://localhost:4000/department", deptInfo)

        if (res.status === 200) {
            alert(`Department ${res.data.message} with an id of ${res.data.id}`)
        }

        navigate("/department")
    }

    useEffect(() => {
        setTimeout(() => {
            setIsFormLoaded(true)
        }, 2000)
    }, [])


    const handleInputChange = e => {
        const { name, value } = e.target
        setDeptInfo({ ...deptInfo, [name]: value })
    }

    return (
        <div className='d-flex flex-column justify-content-center align-items-center vh-100 bg-light'>
            {
                isFormLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h3 className='text-center'>Add Department</h3>
                        <form method='POST' onSubmit={handleFormSubmit} >
                            <div className='my-3 row'>
                                <label htmlFor='name' className='mx-3 fs-4 fw-bold'>Name</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={deptInfo.name} placeholder='Enter department name' name='name' type='text' id='name' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='my-3 row'>
                                <label htmlFor='associated_faculty' className='mx-3 fs-4 fw-bold'>Associated Faculty</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={deptInfo.associated_faculty} placeholder='Enter faculty name' name='associated_faculty' type='text' id='associated_faculty' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                                <button type='submit' className='btn btn-outline-success mb-3 w-25'>Submit</button>
                                <Link to="/department" className='btn btn-outline-dark w-25'>Back</Link>
                            </div>
                        </form>
                    </div>
                ) : (
                    <Spinner
                        variant="dark"
                        role="status"
                        animation="border"
                    />
                )
            }
        </div>
    )
}

export default DeptCreate
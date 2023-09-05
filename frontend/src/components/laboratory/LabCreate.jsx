import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { Link, useNavigate } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'

function LabCreate() {
    const initialDetails = {
        name: "",
        associated_department: ""
    }

    const [labInfo, setLabInfo] = useState(initialDetails)
    const [isFormLoaded, setIsFormLoaded] = useState(false)

    const navigate = useNavigate()

    const handleFormSubmit = async (e) => {
        e.preventDefault()

        const res = await axios.post("http://localhost:4000/laboratory", labInfo)

        if (res.data) {
            alert(`Lab ${res.data.message} with the id: ${res.data.id}`)
        }

        navigate("/laboratory")
    }

    const handleInputChange = e => {
        const { name, value } = e.target
        setLabInfo({ ...labInfo, [name]: value })
    }

    useEffect(() => {
        setTimeout(() => {
            setIsFormLoaded(true)
        }, 2000)
    }, [])

    return (
        <div className='d-flex w-100 vh-100 justify-content-center align-items-center flex-column bg-light'>
            {
                isFormLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h3 className='text-center mb-4'>Add Lab</h3>
                        <form method='POST' onSubmit={handleFormSubmit} >
                            <div className='mb-3 row'>
                                <label htmlFor='name' className='mb-1 fs-4 fw-bold mx-3'>Name</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={labInfo.name} placeholder='Enter lab name' name='name' type='text' id='name' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='mb-3 row'>
                                <label htmlFor='associated_department' className='mb-1 fs-4 fw-bold mx-3'>Associated Department</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={labInfo.associated_department} placeholder='Enter department name' name='associated_department' type='text' id='associated_department' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                                <button type='submit' className='btn btn-outline-success mb-3 w-25'>Submit</button>
                                <Link to="/laboratory" className='btn btn-outline-dark w-25'>Back</Link>
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

    )
}

export default LabCreate
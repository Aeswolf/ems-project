import React, { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import axios from 'axios'
import { Spinner } from 'react-bootstrap'


function FacultyCreate() {

    const [facultyInfo, setFacultyInfo] = useState({
        name: ""
    })

    const [isFormLoaded, setIsFormLoaded] = useState(false)

    const navigate = useNavigate()

    const handleInputChange = e => {
        setFacultyInfo({ ...facultyInfo, name: e.target.value })
    }

    const handleFormSubmit = async (e) => {
        e.preventDefault();

        const res = await axios.post("http://localhost:4000/faculty", facultyInfo)

        if (res.status === 201) {
            alert(`Faculty ${res.data["message"]} with an id of ${res.data["id"]}`)

            navigate("/faculty")
        }
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
                        <h3 className='text-center my-4'>Add Faculty</h3>
                        <form method='POST' onSubmit={handleFormSubmit} >
                            <div className='mb-3 row'>
                                <label htmlFor='name' className='fw-bold fs-4 mb-1 mx-3'> Name </label>
                                <div>
                                    <input className='form-control form-control-lg' value={facultyInfo.name} placeholder='Enter faculty name' name='name' type='text' id='name' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                                <button type='submit' className='btn btn-outline-success mb-3 w-25'>Submit</button>
                                <Link to="/faculty" className='btn btn-outline-dark w-25'>Back</Link>
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

export default FacultyCreate
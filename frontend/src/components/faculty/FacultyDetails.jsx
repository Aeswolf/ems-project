import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { Link, useParams } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'


function FacultyDetails() {
    const [facultyInfo, setFacultyInfo] = useState({})
    const [isInformLoaded, setIsInformLoaded] = useState(false)
    const { id } = useParams()

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get(`http://localhost:4000/faculty/${id}`)
            setFacultyInfo(res.data)
            setIsInformLoaded(true)
        }

        fetchData()
    }, [])
    return (
        <div className='d-flex w-100 vh-100 justify-content-center align-items-center ftlex-column bg-light'>
            {
                isInformLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h1 className='text-center my-3'>Faculty Details</h1>
                        <h3>Name : <span className='mx-3'>{facultyInfo.name}</span></h3>
                        <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                            <Link to={`/faculty/${facultyInfo.id}/edit`} className='btn btn-outline-primary mb-3 w-25'>Edit</Link>
                            <Link to="/faculty" className='btn btn-outline-dark w-25'>Back</Link>
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

export default FacultyDetails
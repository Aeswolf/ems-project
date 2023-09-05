import React, { useEffect, useState } from 'react'
import { Spinner } from 'react-bootstrap'
import { Link, useParams } from 'react-router-dom'
import axios from 'axios'

function LabDetails() {
    const [labInfo, setLabInfo] = useState({})
    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const { id } = useParams()

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get(`http://localhost:4000/laboratory/${id}`)

            if (res.data) {
                setLabInfo(res.data)
                setIsInfoLoaded(true)
            }
        }

        fetchData()
    })
    return (
        <div className='d-flex w-100 vh-100 justify-content-center align-items-center ftlex-column bg-light'>
            {
                isInfoLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h3 className='text-center mb-3'>Lab Details</h3>
                        <div className="row mb-3">
                            <h3> Name : <span className='mx-3'>{labInfo.name} </span></h3>
                        </div>
                        <div className="row mb-3">
                            <h3 className='mb-3'>Department : <span className='mx-3'>{labInfo.associated_department}</span></h3>
                        </div>
                        <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                            <Link to={`/laboratory/${labInfo.id}/edit`} className='btn btn-outline-primary mb-3 w-25'>Edit</Link>
                            <Link to="/laboratory" className='btn btn-outline-dark w-25'>Back</Link>
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

export default LabDetails
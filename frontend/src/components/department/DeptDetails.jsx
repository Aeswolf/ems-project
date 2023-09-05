import React, { useEffect, useState } from 'react'
import { Spinner } from 'react-bootstrap'
import { useParams, Link } from 'react-router-dom'
import axios from 'axios'

function DeptDetails() {
    const initialDetails = {
        name: "",
        associated_faculty: ""
    }

    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const [deptInfo, setDeptInfo] = useState(initialDetails)
    const { id } = useParams()

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get(`http://localhost:4000/department/${id}`)
            setDeptInfo(res.data)
            setIsInfoLoaded(true)
        }

        fetchData()
    }, [])

    return (
        <div className='d-flex w-100 vh-100 justify-content-center align-items-center flex-column bg-light' >
            {
                isInfoLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h1 className="text-center mb-3">Department Details</h1>
                        <div className='row mb-3'>
                            <h3 >Name : <span className='mx-3'>{deptInfo.name}</span> </h3>
                        </div>
                        <div className='row mb-3'>
                            <h3>Faculty :  <span className='mx-3'>{deptInfo.associated_faculty}</span> </h3>
                        </div>

                        <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                            <Link to={`/department/${deptInfo.id}/edit`} className='btn btn-outline-primary mb-3 w-25'>Edit</Link>
                            <Link to="/department" className='btn btn-outline-dark w-25'>Back</Link>
                        </div>
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

export default DeptDetails
import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { Link, useParams, useNavigate } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'

function EquipmentDetails() {
    const initialDetails = {
        name: "",
        quantity: 0,
        associated_laboratory: ""
    }

    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const [equipmentInfo, setEquipmentInfo] = useState(initialDetails)
    const navigate = useNavigate()
    const { id } = useParams()

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get(`http://localhost:4000/equipment/${id}`)

            if (res.status === 200) {
                setIsInfoLoaded(true)

                res.data ? setEquipmentInfo(res.data) : navigate("/equipment")
            }

        }

        fetchData()
        console.log(equipmentInfo)
    }, [])


    return (
        <div className='d-flex w-100 vh-100 justify-content-center align-items-center ftlex-column bg-light'>
            {
                isInfoLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h3 className='text-center mb-3'>Equipment Details</h3>
                        <div className="row mb-3">
                            <h3> Name : <span className='mx-3'>{equipmentInfo.name}</span></h3>
                        </div>
                        <div className="row mb-3">
                            <h3 className='mb-3'>Quantity : <span className='mx-3'>{equipmentInfo.quantity}</span></h3>
                        </div>
                        <div className="row mb-3">
                            <h3 className='mb-3'>Laboratory : <span className='mx-3'>{equipmentInfo.associated_laboratory}</span></h3>
                        </div>
                        <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                            <Link to={`/equipment/${equipmentInfo.id}/edit`} className='btn btn-outline-primary mb-3 w-25'>Edit</Link>
                            <Link to="/equipment" className='btn btn-outline-dark w-25'>Back</Link>
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

export default EquipmentDetails
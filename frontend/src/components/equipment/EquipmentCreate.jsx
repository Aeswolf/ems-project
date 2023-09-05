import React, { useEffect, useState } from 'react'
import { Spinner } from 'react-bootstrap'
import { Link, useNavigate } from 'react-router-dom'
import axios from 'axios'

function EquipmentCreate() {
    const initialDetails = {
        name: "",
        quantity: 0,
        associated_laboratory: ""
    }

    const [isFormLoaded, setIsFormLoaded] = useState(false)
    const [equipmentInfo, setEquipmentInfo] = useState(initialDetails)
    const navigate = useNavigate()

    useEffect(() => {
        setTimeout(() => {
            setIsFormLoaded(true)
        }, 1000);
    }, [])

    const handleFormSubmit = async (e) => {
        e.preventDefault()

        const res = await axios.post("http://localhost:4000/equipment", equipmentInfo)


        if (res.status === 200) {
            alert(`Equipment ${res.data.message} with id ${res.data.id}`)
        }

        navigate("/equipment")
    }

    const handleInputChange = e => {
        let { name, value } = e.target

        if (!isNaN(value)) {
            value = parseInt(value)
        }

        setEquipmentInfo({ ...equipmentInfo, [name]: value })
    }

    return (
        <>
            {

                <div className='d-flex flex-column justify-content-center align-items-center vh-100 bg-light'>
                    {
                        isFormLoaded ? (
                            <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                                <h3 className='text-center'>Add Equipment</h3>
                                <form method='POST' onSubmit={handleFormSubmit} >
                                    <div className='my-3 row'>
                                        <label htmlFor='name' className='mb-1 fw-bold fs-4 mx-3'>Name</label>
                                        <div className='col-sm-12'>
                                            <input required className='form-control form-control-lg' value={equipmentInfo.name} placeholder='Enter equipment name' name='name' type='text' id='name' onChange={handleInputChange} />
                                        </div>
                                    </div>
                                    <div className='my-3 row'>
                                        <label htmlFor='associated_laboratory' className='mb-1 fw-bold fs-4 mx-3'>Associated Lab</label>
                                        <div className='col-sm-12'>
                                            <input required className='form-control form-control-lg' value={equipmentInfo.associated_laboratory} placeholder='Enter name of associated lab' name='associated_laboratory' type='text' id='associated_laboratory' onChange={handleInputChange} />
                                        </div>
                                    </div>
                                    <div className='my-3 row'>
                                        <label htmlFor='quantity' className='mb-2 fw-bold fs-4 mx-3'>Quantity</label>
                                        <div className='col-sm-12'>
                                            <input required className='form-control form-control-lg' value={parseInt(equipmentInfo.quantity)} placeholder='Enter quantity' name='quantity' type='number' min={0} id='quantity' onChange={handleInputChange} />
                                        </div>
                                    </div>
                                    <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                                        <button type='submit' className='btn btn-outline-success mb-3 w-25'>Submit</button>
                                        <Link to="/equipment" className='btn btn-outline-dark w-25'>Back</Link>
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

export default EquipmentCreate
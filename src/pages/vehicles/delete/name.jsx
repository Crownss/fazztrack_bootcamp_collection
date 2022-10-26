import useApi from '../../../helpers/useApi'
import { useEffect } from 'react'
import { Navigate, useNavigate, useParams } from 'react-router-dom'

function DeleteVehicles() {
    const api = useApi()
    const {name} = useParams()
    const navigate = useNavigate()
    useEffect(()=>{
        api.requests({
            method: "DELETE",
            url: "/vehicles/"+name,
            data: null,
        })
        .then(()=> navigate('/vehicles'))
        .catch((err)=>console.log(err.message))

    })
}

export default DeleteVehicles
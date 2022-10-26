import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import React, { useState } from 'react'
import Footer from '../../../components/footer'
import Navbar from '../../../components/navbar'
import { faArrowLeft, faArrowRight, faMinus, faPlus, faHeart } from "@fortawesome/free-solid-svg-icons";
import { NavLink, useParams } from 'react-router-dom';
import axios from 'axios';


export default function DetailVehicles(){
    const [result, setResult] = useState({})
    const {name} = useParams()
    let datas = axios.get(process.env.REACT_APP_API_BACKEND+"/vehicles/"+name).then(res => {setResult(res.data.data)})
    return (
      <>
        <Navbar />
        <div className="mt-12"></div>
        <div className='relative container mx-auto'>
          <div className="abosulute">
            <NavLink to="/vehicles" className='text-lg sm:text-4xl'><FontAwesomeIcon icon={faArrowLeft}/> Detail</NavLink>
          </div>
          <div className="mt-12"></div>
          <div className='sm:flex sm:flex-wrap'>
            <img width="500px" className='rounded-xl' src={result.img} alt={result.name} />
            <h1 style={{fontFamily:"Playfair Display"}} className="text-lg sm:text-5xl mx-10">{result.name} <br/><h6 className="text-lg sm:text-xl font-extralight">{result.location}</h6><h6 style={{fontFamily:"Nunito"}} className='text-green-500 text-lg sm:text-xl'><div className="mt-10"></div>{result.status}<div className="mt-1"></div><h6 className='text-red-500 text-lg sm:text-xl'>No Prepayment</h6></h6><h6 className='text-base sm:text-lg'>Capacity: 1<div className="mt-1"></div>Type: <div className="mt-1"></div>Reservetion before: </h6><div className="mt-5"></div><h4 className='text-xl sm:text-lg px-0 sm:px-20 font-bold'>Rp.{result.price}/month</h4></h1>
          </div>
          <div className='flex container mx-auto p-7 sm:p-5'>
                <button className='bg-blue-500 hover:bg-blue-700 text-white-500 font-bold px-5 py-2 border border-blue-900 rounded'><FontAwesomeIcon icon={faPlus} /></button>
                <h4 className='text-lg sm:text-xl px-20'>2</h4>
                <button className='bg-white-500 hover:bg-white-700 text-black-500 font-bold px-5 py-2 border border-white-900 rounded'><FontAwesomeIcon icon={faMinus} /></button>
              </div>
            <div className="mt-12"></div>
            <div className="flex p-5">
              <FontAwesomeIcon className="text-lg sm:text-4xl" icon={faArrowLeft}/>
              <img width="130px" className="rounded-xl mx-2" src={result.img} alt="dog" />
              <img width="130px" className="rounded-xl" src={result.img} alt="dog" />
              <div className="mt-10"></div>
              <FontAwesomeIcon className="text-lg sm:text-4xl" icon={faArrowRight}/>

            </div>
            <div className="relative sm:flex sm:flex-nowrap container mx-auto p-7">
              <NavLink to="/help/admin" className="bg-gray-700 text-amber-300 hover:bg-gray-700 font-bold py-2 px-20 border border-blue-700 rounded">Chat Admin</NavLink>
              <div className="px-10"></div><br />
              <NavLink to="/help/reservation" className="bg-yellow-500 hover:bg-yellow-500 text-white-500 font-bold py-2 px-20 border border-blue-700 rounded">Reservation</NavLink>
              <div className="px-10"></div><br />
              <NavLink to="/help/admin" className="bg-gray-700 text-amber-300 hover:bg-gray-700 font-bold py-2 px-20 border border-blue-700 rounded"><FontAwesomeIcon icon={faHeart}/> Like</NavLink>
            </div>
        </div>
        <Footer />
      </>
    )
}
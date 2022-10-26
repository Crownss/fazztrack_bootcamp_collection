import React from 'react'
import Footer from '../../components/footer'
import Navbar from '../../components/navbar'
import { useState, useRef, useCallback } from 'react';
import axios from 'axios';
import Cards from '../../components/cards';

export default function Vehicles(){
    const searchRef = useRef(null)
    const [query,setQuery] = useState('')
    const [active, setActive] = useState(false)
    const [result, setResult] = useState([])
    const searchEndPoint = (query) => process.env.REACT_APP_API_BACKEND_SEARCH+query
    const getAll = process.env.REACT_APP_API_BACKEND+"/vehicles/"
    const onChange = useCallback((event) => {
      const query = event.target.value;
      setQuery(query);
      if (query.length) {
        // axios.get(getAll).then(res => setResult(res.data.data))
        axios.get(searchEndPoint(query)).then(res => setResult(res.data.data))
        return
      }else if(!query.length) {
        // axios.get(searchEndPoint(query)).then(res => setResult(res.data.data))
        axios.get(getAll).then(res => setResult(res.data.data))
        return
      }
    }, [])
    const onFocus = useCallback(() => {
        setActive(true)
        window.removeEventListener('click', onClick)
    }, [])
    const onClick = useCallback((event)=>{
        if(searchRef.current && !searchRef.current.contains(event.target)){
            setActive(false)
            window.removeEventListener('click', onClick)
        }
    }, [])
    return (
      <>
        <Navbar />
        <div className="container mx-auto py-20">
            <input className='bg-gray-200 text-center appearance-none border-2 border-gray-200 rounded w-full py-2 px-20 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500' onChange={onChange} onFocus={onFocus} type="search" value={query} />
        </div>
          <div className='sm:grid sm:grid-cols-4 sm:gap-1'>
            {active && result ? result.length>0 && (
              result.map(data => (
                <>
                  <Cards props={data} />
                </>
              ))
            ):
            result.map(data => (
              <>
                  <Cards props={data} />
              </>
            ))}
        </div>
        <Footer/>
      </>
    )
}
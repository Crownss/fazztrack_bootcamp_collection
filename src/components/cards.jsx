import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import React, {useState, useRef} from 'react'
import { NavLink, useNavigate } from 'react-router-dom';
import { faLocationPin } from "@fortawesome/free-solid-svg-icons";
import { useSelector } from "react-redux";
import useApi from '../helpers/useApi';


export default function Cards({props}) {
    const {isAuth} = useSelector((state)=>state.users)
    const [isOpen, setOpen] = useState(false)
    const [isDelete, setDelete] = useState(false)
    const refSend = useRef(null)
    const api = useApi()
    const navigate = useNavigate()
    const [Data, setData] = useState({ img: '', name: '', location:"", desc:"", category:"", price:0, status:"", stock:0})

    const onChangeInput = (event) => {
        event.preventDefault()

        const data = { ...Data }
        data[event.target.name] = event.target.value
        setData(data)
    }

    const sendData = () => {
        api.requests({
            method: 'POST',
            url: '/vehicles/',
            data: Data
        })
        .then(() => {
            navigate('/',{ replace: true })
        })
        .catch((err) => console.log(err.message))

        // axios.post(process.env.REACT_APP_API_BACKEND+"/vehicles/",Data)
        // .then(()=><NavLink to={"/vehicles"} />)
        // .catch((err)=>console.log(err.message))
    }

    const deleteData = () => {
        api.requests({
            method: 'DELETE',
            url: '/vehicles/'+props.name,
        })
        .then(() => {
            navigate('/',{ replace: true })
        })
        .catch((err) => console.log(err.message))
    }
    return (
    <>
        <section class="pb-20">
        <div class="hover:shadow-2xl">
            <div class="relative w-full mx-1">
                <div class="w-full px-4">
                    <div class="bg-white rounded-lg overflow-hidden mb-10">
                    <img src={props.img} alt={props.name} class="w-full"/>
                    <div class="py-4 text-center">
                        <h3>
                            <NavLink to={"/vehicles/"+props.name} class="font-bold text-dark text-xl sm:text-[22px] md:text-xl lg:text-[22px] xl:text-xl 2xl:text-[22px] mb-4 block hover:text-primary ">
                            {props.name} (<FontAwesomeIcon icon={faLocationPin} /> {props.location})
                            </NavLink>
                        </h3>
                        <p class="text-base text-body-color leading-relaxed py-7 mb-7">
                            {props.desc}
                        </p>
                        {isAuth ? (
                        <div className="relative sm:flex">
                        <button onClick={() => setOpen(true)} class="inline-block border border-solid border-black-500 rounded-full text-base text-body-color font-medium hover:border-primary hover:bg-primary hover:text-cyan-500 transition">
                        Create data 
                        </button>
                        <div className="mt-2"></div>
                        <NavLink to={"/vehicles/details/"+props.name} class="inline-block border border-solid border-black-500 rounded-full text-base text-body-color font-medium hover:border-primary hover:bg-primary hover:text-cyan-500 transition">
                        View Details
                        </NavLink>
                        <div className="mt-2"></div>
                        <NavLink to={"/vehicles/update/"+props.name} class="inline-block border border-solid border-black-500 rounded-full text-base text-body-color font-medium hover:border-primary hover:bg-primary hover:text-cyan-500 transition">
                        Update data 
                        </NavLink>
                        <div className="mt-2"></div>
                        <button onClick={() => setDelete(true)} class="inline-block border border-solid border-black-500 rounded-full text-base text-body-color font-medium hover:border-primary hover:bg-primary hover:text-cyan-500 transition">
                        Delete data 
                        </button>
                        </div>
                        ):(
                        <NavLink to={"/vehicles/details/"+props.name} class="inline-block border border-solid border-black-500 rounded-full text-base text-body-color font-medium hover:border-primary hover:bg-primary hover:text-cyan-500 transition">
                        View Details
                        </NavLink>
                        )}
                        {isDelete ? (<>
                            <div id="modal" class="fixed z-50 inset-0 bg-gray-900 bg-opacity-60 overflow-y-auto h-full w-full px-4 modal">
                        <div class="relative top-40 mx-auto shadow-xl rounded-md bg-white max-w-md">

                            <div class="flex justify-between items-center bg-blue-500 text-white text-xl rounded-t-md px-4 py-2">
                                <h3>Delete {props.name}?</h3>
                                <button onClick={() => setDelete(false)}>x</button>
                            </div>
                            <div class="px-4 py-2 border-t border-t-gray-500 flex justify-end items-center space-x-4">
                                <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition" onClick={() => setDelete(false)}>Cancel</button>
                                <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition" onClick={deleteData}>Delete</button>
                            </div>
                        </div>
                    </div></>
                    ):null}
                        {isOpen ? (<>
                            <div id="modal" class="fixed z-50 inset-0 bg-gray-900 bg-opacity-60 overflow-y-auto h-full w-full px-4 modal">
                        <div class="relative top-40 mx-auto shadow-xl rounded-md bg-white max-w-md">

                            <div class="flex justify-between items-center bg-blue-500 text-white text-xl rounded-t-md px-4 py-2">
                                <h3>Create data</h3>
                                <button onClick={() => setOpen(false)}>x</button>
                            </div>

                            <div class="w-full rounded-xl bg-white p-4 shadow-2xl shadow-white/40">
                            <div class="mb-4 flex flex-col">
                                <label class="mb-2 font-semibold">Image</label>
                                <div class="relative">
                                    <input type="text" onChange={onChangeInput} name='img' id="text" class="w-full rounded-lg border border-slate-200 px-2 py-1 pl-8 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" />
                                </div>
                                </div>
                                <div class="mb-4 grid grid-cols-2 gap-4">
                                <div class="flex flex-col">
                                    <label class="mb-2 font-semibold">Name</label>
                                    <input type="text" onChange={onChangeInput} name='name' id="text" class="w-full max-w-lg rounded-lg border border-slate-200 px-2 py-1 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" />
                                </div>
                                <div class="flex flex-col">
                                    <label class="mb-2 font-semibold">Location</label>
                                    <input type="text" onChange={onChangeInput} name="location" id="text" class="w-full max-w-lg rounded-lg border border-slate-200 px-2 py-1 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" />
                                </div>
                                </div>
                                <div class="mb-4 flex flex-col">
                                <label class="mb-2 font-semibold">Description</label>
                                <div class="relative">
                                    <input type="text" onChange={onChangeInput} id="text" name="desc" class="w-full rounded-lg border border-slate-200 px-2 py-7 pl-8 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" />
                                </div>
                                </div>
                                <div class="mb-4 grid grid-cols-2 gap-4">
                                <div class="flex flex-col">
                                    <label class="mb-2 font-semibold">Category</label>
                                    <select onChange={onChangeInput} className="w-full max-w-lg rounded-lg border border-slate-200 px-2 py-1 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" name="category" id="">
                                        <option defaultValue={"Select"}>Select</option>
                                        <option value="Car">Car</option>
                                        <option value="Bike">Bike</option>
                                        <option value="Motorcyle">Motorcyle</option>

                                    </select>
                                </div>
                                <div class="flex flex-col">
                                    <label class="mb-2 font-semibold">Price</label>
                                    <input type="number" onChange={onChangeInput} name="price" id="number" class="w-full max-w-lg rounded-lg border border-slate-200 px-2 py-1 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" />
                                </div>
                                </div>
                                <div class="mb-4 grid grid-cols-2 gap-4">
                                <div class="flex flex-col">
                                    <label class="mb-2 font-semibold">Status</label>
                                    <input type="text" onChange={onChangeInput} name='status' id="text" class="w-full max-w-lg rounded-lg border border-slate-200 px-2 py-1 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" />
                                </div>
                                <div class="flex flex-col">
                                    <label class="mb-2 font-semibold">Stock</label>
                                    <input type="number" onChange={onChangeInput} name="stock" id="number" class="w-full max-w-lg rounded-lg border border-slate-200 px-2 py-1 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" />
                                </div>
                                </div>
                                {/* <div class="mb-4 flex flex-col">
                                <label for="age" class="mb-2 font-semibold">Number input</label>
                                <input type="number" id="age" class="w-full max-w-[200px] rounded-lg border border-slate-200 px-2 py-1 hover:border-blue-500 focus:outline-none focus:ring focus:ring-blue-500/40 active:ring active:ring-blue-500/40" />
                                </div> */}
                            </div>
                            <div class="px-4 py-2 border-t border-t-gray-500 flex justify-end items-center space-x-4">
                                <button class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition" onClick={sendData} ref={refSend}>Create</button>
                            </div>
                        </div>
                    </div></>
                    ):null}

                    </div>
                    </div>
                </div>
            </div>
        </div>
        </section>
    </>
  )
}

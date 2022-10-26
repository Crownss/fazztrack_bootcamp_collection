import { useState } from "react"
import { NavLink } from "react-router-dom"
import { useSelector, useDispatch } from "react-redux"
import { logout } from "../store/reducer/users"

export default function Navbar() {
    const [isOpen, setOpen] = useState(false)
    const {isAuth} = useSelector((state)=>state.users)
    const dispatch = useDispatch()
    return(
        <>
            <nav class="static top-0 left-0 right-0 p-12 bg-white-500 border-white-200 dark:bg-gray-800 dark:border-gray-700">
            <div class="container flex flex-wrap justify-between items-center mx-auto">
                <div class="flex items-center">
                <svg
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
                class="fill-current text-white h-8 w-8 mr-2"
                >
                <path
                    d="M12.001 4.8c-3.2 0-5.2 1.6-6 4.8 1.2-1.6 2.6-2.2 4.2-1.8.913.228 1.565.89 2.288 1.624C13.666 10.618 15.027 12 18.001 12c3.2 0 5.2-1.6 6-4.8-1.2 1.6-2.6 2.2-4.2 1.8-.913-.228-1.565-.89-2.288-1.624C16.337 6.182 14.976 4.8 12.001 4.8zm-6 7.2c-3.2 0-5.2 1.6-6 4.8 1.2-1.6 2.6-2.2 4.2-1.8.913.228 1.565.89 2.288 1.624 1.177 1.194 2.538 2.576 5.512 2.576 3.2 0 5.2-1.6 6-4.8-1.2 1.6-2.6 2.2-4.2 1.8-.913-.228-1.565-.89-2.288-1.624C10.337 13.382 8.976 12 6.001 12z"
                />
                </svg>
                <span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white">Vehicles rent</span>
                </div>
                <button onClick={() => setOpen(!isOpen)} data-collapse-toggle="navbar-solid-bg" type="button" class="inline-flex justify-center items-center ml-3 text-gray-400 rounded-lg md:hidden hover:text-gray-900 focus:outline-none focus:ring-2 focus:ring-blue-300 dark:text-gray-400 dark:hover:text-white dark:focus:ring-gray-500" aria-controls="navbar-solid-bg" aria-expanded="false">
                <span class="sr-only">Open main menu</span>
                <svg class="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd"></path></svg>
                </button>
                <div class={"w-full md:block md:w-auto"+(isOpen ? "flex" : " hidden")} id="navbar-solid-bg">
                <ul class="flex flex-col mt-4 bg-gray-50 rounded-lg md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-transparent dark:bg-gray-800 md:dark:bg-transparent dark:border-gray-700">
                    <li>
                        <NavLink to="/" class="block py-2 pr-4 pl-3 text-white bg-white-500 rounded md:bg-transparent md:text-blue-700 md:p-0 md:dark:text-white dark:bg-blue-600 md:dark:bg-transparent" aria-current="page">Home</NavLink>
                    </li>
                    <li>
                        <NavLink to="/vehicles" class="block py-2 pr-4 pl-3 bg-white-500 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">Vehicles Type</NavLink>
                    </li>
                    <li>
                        <NavLink to="/history" class="block py-2 pr-4 pl-3 bg-white-500 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">History</NavLink>
                    </li>
                    <li>
                        <NavLink to="/about" class="block py-2 pr-4 pl-3 bg-white-500 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">About</NavLink>
                    </li>
                    {!isAuth ? (
                    <li className="relative sm:flex">
                        <div className="mt-2"></div>
                        <NavLink to="/auth/login" className="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-10 border border-blue-500 hover:border-transparent rounded">Login</NavLink>
                        <div className="px-2"></div><br />
                        <NavLink to="/auth/register" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-10 border border-blue-700 rounded">Register</NavLink>
                    </li>
                    ):
                    <li className="relative sm:flex">
                    <div className="mt-2"></div>
                    <NavLink to="/auth/login" onClick={()=>dispatch(logout())} className="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-10 border border-blue-500 hover:border-transparent rounded">Logout</NavLink>
                    <div className="px-2"></div><br />
                    </li>
                }
                    <li>
                        
                    </li>
                </ul>
                </div>
            </div>
            </nav>
        </>
    )
}
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import Footer from '../../components/footer'
import { faGoogle } from '@fortawesome/free-brands-svg-icons';
import { NavLink } from 'react-router-dom';
import { useState, useRef, useEffect } from 'react';
import useApi from '../../helpers/useApi';
import { useNavigate } from 'react-router-dom';
import {login} from "../../store/reducer/users";
import { useDispatch, useSelector } from 'react-redux';


export default function Login(){
    const [Users, setUsers] = useState({ username: 'username', password: 'password'})
    const refLogin = useRef(null)
    const dispatch = useDispatch()
    const {isAuth} = useSelector((state)=>state.users)
    const navigate = useNavigate()
    const api = useApi()


    useEffect(()=>{
        if(isAuth){
            navigate('/')
        }
    },[isAuth])

    const onChangeInput = (event) => {
        event.preventDefault()

        const data = { ...Users }
        data[event.target.name] = event.target.value
        setUsers(data)
    }

    const loginUsers = () => {
        api.requests({
            method: 'POST',
            url: '/users/login',
            data: Users
        })
        .then((res) => {
            const {data} = res.data
            dispatch(login(data))
        })
        .catch((err) => console.log(err))

    }
        return (
            <>
                <div className="bg-no-repeat bg-auto" style={{backgroundImage:`url("/img/login.png")`}}>
                    <div className="flex">
                        <div className="mt-12"></div>
                        <div className="container mx-auto py-20 px-20 break-all" style={{fontFamily: "Playfair Display"}}><h1 className='text-left text-8xl'>Le'ts Explore <br /> The World </h1><br/>
                        <h6 className="text-left text-xl">Don't have account?</h6>
                        <div className="mt12"></div><br /><br />
                        <NavLink to="/auth/register" className="bg-gray-700 text-amber-300 hover:bg-gray-700 text-yellow font-bold py-2 px-20 border border-blue-700 rounded">Sign up</NavLink>
                        </div>
                        <div className="container mx-auto py-20 p-6 px-20 m-10">
                        <div className="mt-12"></div>
                        <input type="text" onChange={onChangeInput} name="username" id="username" placeholder="Username" class="container mx-auto rounded-md border border-[#e0e0e0] bg-white py-3 px-6 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"/>
                        <div className="mt-12"></div>

                        <input type="password" onChange={onChangeInput} name="password" id="password" placeholder="Password" class="container mx-auto rounded-md border border-[#e0e0e0] bg-white py-3 px-6 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"/>
                        <NavLink to="/forget/password">Forget Password?</NavLink>
                        <div className="mt-12"></div>
                        <button onClick={loginUsers} ref={refLogin} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-20 border border-blue-700 rounded">Login</button>
                        <br />
                        <div className="mt-5"></div>
                        <button className="bg-white-500 font-bold py-2 px-8 border rounded"><FontAwesomeIcon icon={faGoogle}/> Login with google</button>
                        </div>
                        </div>
                </div>
                <Footer />
            </>
        )
}
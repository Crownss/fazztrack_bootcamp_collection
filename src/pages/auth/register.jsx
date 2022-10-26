import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { Component } from 'react'
import Footer from '../../components/footer'
import { faGoogle } from '@fortawesome/free-brands-svg-icons';
import { NavLink } from 'react-router-dom';


export default class Register extends Component{
    render(){
        return (
            <>
                <div className="bg-no-repeat bg-auto" style={{backgroundImage:`url("/img/login.png")`}}>
                    <div className="flex">
                        <div className="mt-12"></div>
                        <div className="container mx-auto py-20 px-20 break-all" style={{fontFamily: "Playfair Display"}}><h1 className='text-left text-8xl'>Le'ts Explore <br /> The World </h1><br/>
                        <h6 className="text-left text-xl">Already have account?</h6>
                        <div className="mt12"></div><br /><br />
                        <NavLink to="/auth/login" className="bg-gray-700 text-amber-300 hover:bg-gray-700 text-yellow font-bold py-2 px-20 border border-blue-700 rounded">Login</NavLink>
                        </div>
                        <div className="container mx-auto py-20 p-6 px-20 m-10">
                        <form method='POST' action=''>
                        <div className="mt-12"></div>
                        <input type="email" name="email" id="username" placeholder="Email" class="container mx-auto rounded-md border border-[#e0e0e0] bg-white py-3 px-6 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"/>
                        <div className="mt-12"></div>

                        <input type="text" name="username" id="username" placeholder="Username" class="container mx-auto rounded-md border border-[#e0e0e0] bg-white py-3 px-6 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"/>
                        <div className="mt-12"></div>
                        <input type="password" name="password" id="password" placeholder="Password" class="container mx-auto rounded-md border border-[#e0e0e0] bg-white py-3 px-6 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"/>
                        <div className="mt-12"></div>
                        <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-20 border border-blue-700 rounded">Sign up</button>
                        <br />
                        <div className="mt-5"></div>
                        <button type="submit" className="bg-white-500 font-bold py-2 px-8 border rounded"><FontAwesomeIcon icon={faGoogle}/> Sign up with google</button>
                        </form>
                        </div>
                        </div>
                </div>
                <Footer />
            </>
        )
    }
}
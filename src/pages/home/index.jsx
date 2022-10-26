import Footer from "../../components/footer";
import Navbar from '../../components/navbar';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faStar } from "@fortawesome/free-solid-svg-icons";
import { useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import {useEffect} from 'react';

export default function Home(){
  const {isAuth} = useSelector((state)=>state.users)
  const navigate = useNavigate()

  useEffect(()=>{
    if(isAuth){
        navigate('/vehicles')
    }
},[isAuth])
    return(
      <>
        <Navbar />
        <div className="bg-no-repeat bg-contain" style={{backgroundImage: `url("/img/home.png")`}}>
          <h1 className="text-left text-8xl container mx-auto p-7" style={{fontFamily: "Playfair Display"}}>Explore and <br />travel</h1>
          <h5 className="text-left text-2xl font-bold container mx-auto p-10" style={{fontFamily:"Nunito"}}>Vehicle finder</h5>
          <div className="mt-12"></div>
          <div className="relative sm:flex justify-between px-4 sm:px-7">
              <div className="flex-1 text-left container mx-auto px-10">
                <select class="form-select form-select-sm appearance-none block px-5 pr-12 py-2 text-sm text-center font-normal text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" aria-label=".form-select-sm">
                  <option selected>Location</option>
                  <option value="jakarta">Jakarta</option>
                  <option value="bali">Bali</option>
                  <option value="yogyakarta">Yogyakarta</option>
                </select>
              </div>
              <div className="text-left container mx-auto px-10 sm:px-0">
                <select class="form-select form-select-sm appearance-none block px-5 pr-12 py-2 text-sm text-center font-normal text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" aria-label=".form-select-sm">
                  <option selected>Type</option>
                  <option value="jakarta">Jakarta</option>
                  <option value="bali">Bali</option>
                  <option value="yogyakarta">Yogyakarta</option>
                </select>
              </div>
            </div>
            <div className="relative sm:flex justify-between px-4 sm:px-7">
              <div className="flex-1 text-left container mx-auto px-10">
                <select class="form-select form-select-sm appearance-none block px-5 pr-12 py-2 text-sm text-center font-normal text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" aria-label=".form-select-sm">
                  <option selected>Payment</option>
                  <option value="jakarta">Jakarta</option>
                  <option value="bali">Bali</option>
                  <option value="yogyakarta">Yogyakarta</option>
                </select>
              </div>
              <div className="text-left container mx-auto px-10 sm:px-0">
                <select class="form-select form-select-sm appearance-none block px-5 pr-12 py-2 text-sm text-center font-normal text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none" aria-label=".form-select-sm">
                  <option selected>Date</option>
                  <option value="jakarta">Jakarta</option>
                  <option value="bali">Bali</option>
                  <option value="yogyakarta">Yogyakarta</option>
                </select>
              </div>
          </div>
          <div className="mt-12"></div>
              <div className="px-12">
                <a href="#" className="text-left ml-1 sm:ml-5 bg-blue-500 hover:bg-blue-700 text-white-500 font-bold py-2 px-10 border border-blue-700 rounded-xl">Explore</a>
              </div>
          <div className="mt-12"></div>
          <div className="pt-12"></div>
          <div className="mt-12"></div>

          <div className="flex flex-row">
            <h1 className="sm:px-20 text-left text-3xl text-white-500 container mx-auto p-10" style={{fontFamily: "Playfair Display"}}>Popular in town</h1>
            <a href="#" className="text-lg py-12 basis-1/6 text-white-500">view all</a>
          </div>
          <br />
          <div className="relative sm:flex justify-between px-0 sm:px-20">
            <div class="max-w-md bg-white rounded-xl overflow-hidden cursor-pointer shadow-lg hover:shadow-2xl hover:scale-105 transform transition-all duration-500">
              <div class="p-4">
                <img class="rounded-xl" src="https://images.unsplash.com/photo-1599839575945-a9e5af0c3fa5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1049&q=80" alt="Dog" />
              </div>
                <div class="flex justify-between p-6">
                  <div class="flex items-center space-x-4">
                    <h1 class="text-lg text-gray-900 font-bold">Felipe Sacudon</h1>
                  </div>
                </div>
            </div>
            <div class="max-w-md bg-white rounded-xl overflow-hidden cursor-pointer shadow-lg hover:shadow-2xl hover:scale-105 transform transition-all duration-500">
              <div class="p-4">
                <img class="rounded-xl" src="https://images.unsplash.com/photo-1599839575945-a9e5af0c3fa5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1049&q=80" alt="Dog" />
              </div>
                <div class="flex justify-between p-6">
                  <div class="flex items-center space-x-4">
                    <h1 class="text-lg text-gray-900 font-bold">Felipe Sacudon</h1>
                  </div>
                </div>
            </div>
            <div class="max-w-md bg-white rounded-xl overflow-hidden cursor-pointer shadow-lg hover:shadow-2xl hover:scale-105 transform transition-all duration-500">
              <div class="p-4">
                <img class="rounded-xl" src="https://images.unsplash.com/photo-1599839575945-a9e5af0c3fa5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1049&q=80" alt="Dog" />
              </div>
                <div class="flex justify-between p-6">
                  <div class="flex items-center space-x-4">
                    <h1 class="text-lg text-gray-900 font-bold">Felipe Sacudon</h1>
                  </div>
                </div>
            </div>
            <div class="max-w-md bg-white rounded-xl overflow-hidden cursor-pointer shadow-lg hover:shadow-2xl hover:scale-105 transform transition-all duration-500">
              <div class="p-4">
                <img class="rounded-xl" src="https://images.unsplash.com/photo-1599839575945-a9e5af0c3fa5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1049&q=80" alt="Dog" />
              </div>
                <div class="flex justify-between p-6">
                  <div class="flex items-center space-x-4">
                    <h1 class="text-lg text-gray-900 font-bold">Felipe Sacudon</h1>
                  </div>
                </div>
            </div>
          </div>
          <div className="mt-12"></div>
            <div className="relative sm:flex">
              <div className="text-left text-3xl" style={{fontFamily: "Playfair Display"}}>
                <h1 className="xl:text-5xl px-7 xl:px-20">Testimonials</h1>
                <div className="mt-12"></div>
                <div className="flex px-7 xl:px-20">
                  <FontAwesomeIcon icon={faStar} />
                  <FontAwesomeIcon icon={faStar} />
                  <FontAwesomeIcon icon={faStar} />
                  <FontAwesomeIcon icon={faStar} />
                  <FontAwesomeIcon icon={faStar} />
                </div>
                <div className="mt-7"></div>
                <h6 className="px-7 text-sm sm:px-20">”It was the right decision to rent vehicle here, I spent less money and enjoy the trip. It was an amazing experience to <br />have a ride for wildlife trip!”</h6>
                <div className="mt-12"></div>
                <h4 className="text-xl xl:px-20 px-7">Edward Newgate</h4>
                <h4 className="text-base py-5 xl:px-20 px-7">Founder circle</h4>
              </div>
              <img className="rounded-2xl" alt="gambar orang" src="/img/edward.png" />
            </div>
          <div className="mt-12"></div>
        </div>
        <Footer />
      </>
    )
  }

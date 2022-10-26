import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFacebook, faInstagram, faLinkedin, faTwitter, faYoutube } from '@fortawesome/free-brands-svg-icons';

export default function Footer() {
  return (
    <footer class="relative bg-blueGray-200 pt-8 pb-6">
        <div class="container mx-auto px-4">
            <div class="flex flex-wrap text-left lg:text-left">
            <div class="w-full lg:w-6/12 px-4">
                <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" class="fill-current text-white h-8 w-8 mr-2">
                    <path d="M12.001 4.8c-3.2 0-5.2 1.6-6 4.8 1.2-1.6 2.6-2.2 4.2-1.8.913.228 1.565.89 2.288 1.624C13.666 10.618 15.027 12 18.001 12c3.2 0 5.2-1.6 6-4.8-1.2 1.6-2.6 2.2-4.2 1.8-.913-.228-1.565-.89-2.288-1.624C16.337 6.182 14.976 4.8 12.001 4.8zm-6 7.2c-3.2 0-5.2 1.6-6 4.8 1.2-1.6 2.6-2.2 4.2-1.8.913.228 1.565.89 2.288 1.624 1.177 1.194 2.538 2.576 5.512 2.576 3.2 0 5.2-1.6 6-4.8-1.2 1.6-2.6 2.2-4.2 1.8-.913-.228-1.565-.89-2.288-1.624C10.337 13.382 8.976 12 6.001 12z"/>
                </svg>
                <h6 class="text-left text-md mt-0 mb-2 text-blueGray-600">
                Plan and book your perfect trip with<br /> expert advice, travel tips for vehicle<br /> information from us
                </h6>
                <br />
                <br />
                <h6 class="text-left text-md mt-0 mb-2 text-blueGray-600">
                ©2020 Vehicles Rent. All rights reserved
                </h6>
            </div>
            <div class="w-full lg:w-6/12 px-4">
                <div class="flex flex-wrap items-top mb-6">
                <div class="w-full lg:w-4/12 px-4 ml-auto">
                    <span class="block uppercase text-blueGray-500 text-sm font-bold mb-2">Destionation</span>
                    <ul class="list-unstyled">
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Jakarta</h6>
                    </li>
                    <li>
                        <h6 className="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Yogyakarta</h6>
                    </li>
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Bali</h6>
                    </li>
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Malang</h6>
                    </li>
                    </ul>
                </div>
                <div class="w-full lg:w-4/12 px-4">
                    <span class="block uppercase text-blueGray-500 text-sm font-bold mb-2">Vehicles</span>
                    <ul class="list-unstyled">
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Bike</h6>
                    </li>
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Cars</h6>
                    </li>
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Motorbike</h6>
                    </li>
                    </ul>
                </div>
                <div class="w-full lg:w-4/12 px-4">
                    <span class="block uppercase text-blueGray-500 text-sm font-bold mb-2">Interest</span>
                    <ul class="list-unstyled">
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Adventure Travel</h6>
                    </li>
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Art and Culture</h6>
                    </li>
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Wildlife and Nature</h6>
                    </li>
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Family Holidays</h6>
                    </li>
                    <li>
                        <h6 class="text-blueGray-600 hover:text-blueGray-800 font-medium block pb-2 text-sm">Culinary Trip</h6>
                    </li>
                    </ul>
                </div>
                </div>
            </div>
            </div>
            <hr class="my-5 border-blueGray-300" />
            <div class="flex justify-between">
                <div class="mx-auto text-center text-sm text-blueGray-500 font-semibold">
                    <a href='https://instgr.am/abcdefuck._.you'>
                        <FontAwesomeIcon icon={faInstagram}/>
                    </a>
                </div>
                <div class="mx-auto text-center text-sm text-blueGray-500 font-semibold">
                    <a href='https://instgr.am/abcdefuck._.you'>
                        <FontAwesomeIcon icon={faFacebook}/>
                    </a>
                </div>
                <div class="mx-auto text-center text-sm text-blueGray-500 font-semibold">
                    <a href='https://instgr.am/abcdefuck._.you'>
                        <FontAwesomeIcon icon={faTwitter}/>
                    </a>
                </div>
                <div class="mx-auto text-center text-sm text-blueGray-500 font-semibold">
                    <a href='https://instgr.am/abcdefuck._.you'>
                        <FontAwesomeIcon icon={faLinkedin}/>
                    </a>
                </div>
                <div class="mx-auto text-center text-sm text-blueGray-500 font-semibold">
                    <a href='https://instgr.am/abcdefuck._.you'>
                        <FontAwesomeIcon icon={faYoutube}/>
                    </a>
                </div>
            </div>
        </div>
        </footer>
  )
}

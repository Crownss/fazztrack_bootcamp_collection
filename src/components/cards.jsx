import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import React from 'react'
import { NavLink } from 'react-router-dom'
import { faLocationPin } from "@fortawesome/free-solid-svg-icons";

export default function Cards({props}) {
  return (
    <>
        <section class="pb-20">
        <div class="hover:shadow-2xl">
            <div class="relative w-full mx-1">
                <div class="w-full px-4">
                    <div class="bg-white rounded-lg overflow-hidden mb-10">
                    <img src={props.img} alt="image" class="w-full"/>
                    <div class="py-4 text-center">
                        <h3>
                            <NavLink to={"/vehicles/"+props.name} class="font-bold text-dark text-xl sm:text-[22px] md:text-xl lg:text-[22px] xl:text-xl 2xl:text-[22px] mb-4 block hover:text-primary ">
                            {props.name} (<FontAwesomeIcon icon={faLocationPin} /> {props.location})
                            </NavLink>
                        </h3>
                        <p class="text-base text-body-color leading-relaxed py-7 mb-7">
                            {props.desc}
                        </p>
                        <NavLink to={"/vehicles/"+props.name} class="inline-block py-2 px-7 border border-solid border-black-500 rounded-full text-base text-body-color font-medium hover:border-primary hover:bg-primary hover:text-cyan-500 transition">
                        View Details
                        </NavLink>
                    </div>
                    </div>
                </div>
            </div>
        </div>
        </section>
    </>
  )
}

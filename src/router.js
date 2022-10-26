import React from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Login from './pages/auth/login';
import Register from './pages/auth/register';
import Home from './pages/home';
import Vehicles from './pages/vehicles';
import DetailVehicles from './pages/vehicles/name';
import About from './pages/about';

export default function Router() {
  return (
    <BrowserRouter>
        <Routes>
            <Route exact path="/" element={<Home />} />
            <Route exact path="/auth/login" element={<Login />} />
            <Route exact path="/auth/register" element={<Register/>} />
            <Route exact path="/vehicles" element={<Vehicles />} />
            <Route exact path="/vehicles/:name" element={<DetailVehicles />} />
            <Route exact path="/about" element={<About />} />
        </Routes>
    </BrowserRouter>
  )
}

import { Navigate } from 'react-router-dom';
import { useSelector } from 'react-redux';

export default function PrivateRoute({components}) {
    const {isAuth} = useSelector((state)=>state.users)

    if(!isAuth){
        return <Navigate to={'/404'} replace />
    }
    return components
}
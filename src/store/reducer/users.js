import { createSlice } from '@reduxjs/toolkit'

const userSlice = createSlice({
    name: 'users',
    initialState: {
        token: '',
        isAuth: false,
        data: {}
    },
    reducers: {
        login(state, actions) {
            return {
                ...state,
                isAuth: true,
                token: actions.payload
            }
        },
        logout(state, actions) {
            return {
                ...state,
                isAuth: false,
                token: ''
            }
        },
        adduser(state, actions) {
            return {
                ...state,
                data: actions.payload
            }
        }
    }
})

export const { login, logout, adduser } = userSlice.actions
export default userSlice.reducer
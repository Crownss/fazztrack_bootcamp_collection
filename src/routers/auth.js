const express = require('express')
const ctrls_auth = require('../controllers/auth')
const Router = express.Router()
const upload = require('../middleware/uploader')


Router.get('/user', ctrls_auth.GETALL)
Router.get('/user/:username', ctrls_auth.GETUSERNAME)
Router.post('/register', ctrls_auth.REGISTER)
Router.post('/login', ctrls_auth.LOGIN)
Router.put('/user/update/profile', upload.single('profile'),ctrls_auth.UPDATE)
Router.put('/user/update/password',ctrls_auth.UPDATE_PASSWORD)
Router.put('/user/update',ctrls_auth.UPDATE)
Router.delete('/user/delete',ctrls_auth.DELETE)
module.exports = Router
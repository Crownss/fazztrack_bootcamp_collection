const model_users = require('../models/users')
const bcrypt = require('bcrypt')
const response = require('../helpers/response')
const { GenToken } = require('../helpers/jwt')
const verify = require('jsonwebtoken')
const validate = require('../middleware/validate_users')
var ctrls_auth = {}

ctrls_auth.GETALL = async (req, res) =>{
    try {
        const data = await model_users.GetAll()
        const mapping = await data.map((res)=>{return {"id":res.id,"profile":res.profile,"email":res.email, "username":res.username,"no_telp":res.no_telp, "is_admin":res.is_admin, "created_at":res.created_at, "updated_at":res.updated_at}})
        return validate(req, res, 200, mapping)
    } catch (error) {
        return response(res, 400, error, true)
    }
}

ctrls_auth.GETUSERNAME = async(req,res)=>{
    try {
        const data = await model_users.GetUsername(req.params.username)
        return response(res, 200, {"id":data[0].id,"profile":data[0].profile,"email":data[0].email, "username":data[0].username,"no_telp":data[0].no_telp, "is_admin":data[0].is_admin, "created_at":data[0].created_at, "updated_at":data[0].updated_at})
    } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_auth.REGISTER = async(req, res)=>{
    try {
        const {profile, email, username, password, no_telp, is_admin} = req.body
        const password_hash = await bcrypt.hash(password, 7)
        const data = await model_users.Post({profile, email, username, password_hash, no_telp, is_admin})
        return response(res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_auth.LOGIN = async(req, res)=>{
    try {
        const {username, password} = req.body
        const get_username = await model_users.GetUsername(username)
        if(get_username[0].username != ''){
            const password_hash = await bcrypt.hash(password, 7)
            const password_check = await bcrypt.compare(password, password_hash)
            if(!password_check){
                return response(res, 401, "wrong password !", true)
            }
            const result = await GenToken(username)
            return response(res, 200, result)
        }
        return response(res, 404, "username not found !", true)
    } catch (error) {
        console.log(error)
        return response(res, 500, error, true)
    }
}

ctrls_auth.UPDATE_PROFILE = async(req, res)=>{
    try {
        const {email, username, no_telp, is_admin} = req.body
        if (req.headers && req.headers.authorization) {
            const authorization = await req.headers.authorization.split(' ')[1]
            const verify_jwt = verify.verify(authorization, process.env.JWT_KEY)
            const profile = 'http://localhost:8000/'+req.file.path
            const get_username = await model_users.GetUsername(verify_jwt.user)
            const data = await model_users.Update({profile, email, username, no_telp, is_admin}, get_username[0].username)
            return response(res, 200, data)
        }
        return response(res, 403, "unautorized!", true)
    } catch (error) {
        console.log(error)
        return response(res, 500, error, true)
    }
}

ctrls_auth.UPDATE = async(req, res)=>{
    try {
        const {email, username, no_telp, is_admin} = req.body
        if (req.headers && req.headers.authorization) {
            const authorization = await req.headers.authorization.split(' ')[1]
            const verify_jwt = verify.verify(authorization, process.env.JWT_KEY)
            const get_username = await model_users.GetUsername(verify_jwt.user)
            const data = await model_users.Update({email, username, no_telp, is_admin}, get_username[0].username)
            return response(res, 200, data)
        }
        return response(res, 403, "unautorized!", true)
    } catch (error) {
        console.log(error)
        return response(res, 500, error, true)
    }
}

ctrls_auth.UPDATE_PASSWORD = async(req, res)=>{
    try {
        const {password, verif_password} = req.body
        if (req.headers && req.headers.authorization) {
            const authorization = await req.headers.authorization.split(' ')[1]
            const verify_jwt = verify.verify(authorization, process.env.JWT_KEY)
            const password_hash = await bcrypt.hash(password, 7)
            if(password != verif_password){
                return response(res, 400, "password and verif password is not same !")
            }
            const get_username = await model_users.GetUsername(verify_jwt.user)
            const data = await model_users.Update({password_hash}, get_username[0].username)
            return response(res, 200, data)
        }
        return response(res, 403, "unautorized!", true)
    } catch (error) {
        console.log(error)
        return response(res, 500, error, true)
    }
}

ctrls_auth.DELETE = async (req, res) => {
    try {
        if (req.headers && req.headers.authorization) {
            const authorization = req.headers.authorization.split(' ')[1]
            const verify_jwt = await verify.verify(authorization, process.env.JWT_KEY)
            const get_username = await model_users.GetUsername(verify_jwt.user)
            const data = await model_users.Delete(get_username[0].username)
            return response(res, 200, data)
        }
        return response(res, 403, "unautorized!", true)
    } catch (error) {
        console.log(error)
        return response(res, 500, error, true)
    }
}

module.exports = ctrls_auth
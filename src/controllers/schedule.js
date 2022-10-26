const response = require("../helpers/response")
const model_schedule = require("../models/schedule")
const verify = require('jsonwebtoken')
const model_users = require('../models/users')


var ctrls_schedule = {}

ctrls_schedule.GETALL = async (req, res) => {
    try {
        const pagination = {
            page: req.query.page || 1,
            limit: req.query.limit || 5,
            order: req.query.order
        }
        const data = await model_schedule.GetAll(pagination)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_schedule.GETMOVIE = async (req, res) => {
    try {
        const data = await model_schedule.GetMovie(req.params.movie_id)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_schedule.POST = async (req, res) => {
    try {
        if (req.headers && req.headers.authorization) {
            const authorization = req.headers.authorization.split(' ')[1]
            const verify_jwt = await verify.verify(authorization, process.env.JWT_KEY)
            const get_username = await model_users.GetUsername(verify_jwt.user)
            if(get_username[0].is_admin == true){
                const {movie_id, price, location, date_start, date_end, premier, time} = req.body
                const data = await model_schedule.Post({movie_id, price, location, date_start, date_end, premier, time})
                return response(res, 200, data)
            }
            return response(res, 403, "you are not admin !")
        }
        return response(res, 403, "unautorized !")
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_schedule.UPDATE = async (req, res) => {
    try {
        if (req.headers && req.headers.authorization) {
            const authorization = req.headers.authorization.split(' ')[1]
            const verify_jwt = await verify.verify(authorization, process.env.JWT_KEY)
            const get_username = await model_users.GetUsername(verify_jwt.user)
            if(get_username[0].is_admin == true){
                const {movie_id, price, location, date_start, date_end, premier, time} = req.body
                const data = await model_schedule.Update({movie_id, price, location, date_start, date_end, premier, time}, req.params.movie_id)
                return response(res, 200, data)
            }
            return response(res, 403, "you are not admin !")
        }
        return response(res, 403, "unautorized !")
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_schedule.DELETE = async (req, res) => {
    try {
        if (req.headers && req.headers.authorization) {
            const authorization = req.headers.authorization.split(' ')[1]
            const verify_jwt = await verify.verify(authorization, process.env.JWT_KEY)
            const get_username = await model_users.GetUsername(verify_jwt.user)
            if(get_username[0].is_admin == true){
                const data = await model_schedule.Delete(req.params.movie_id)
                return response(res, 200, data)
            }
            return response(res, 403, "you are not admin !")
        }
        return response(res, 403, "unautorized !")
    } catch (error) {
        return response(res, 500, error, true)
    }
}
module.exports = ctrls_schedule
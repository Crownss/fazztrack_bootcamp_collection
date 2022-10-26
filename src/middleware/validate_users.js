const model_users = require('../models/users')
const verify = require('jsonwebtoken')
const response = require('../helpers/response')


async function validate(req, res, data){
    try {
        if (req.headers && req.headers.authorization) {
            const authorization = req.headers.authorization.split(' ')[1]
            const verify_jwt = verify.verify(authorization, process.env.JWT_KEY)
            const get_username = await model_users.GetUsername(verify_jwt.user)
            if(get_username[0].is_admin == true){
                return response(res, 200, data)
            }
            return response(res, 403, "forbidden !", true)
        }
    }catch (error) {
        return response(res, 401, "unautorized !", true)
    }
}

module.exports = validate
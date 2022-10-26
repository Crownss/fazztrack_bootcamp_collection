var helpers_jwt = {}
const jwt = require('jsonwebtoken')
const response = require('./response')

helpers_jwt.GenToken = async (username) => {
    try {
       const payload = {
         user: username,
       }
 
       const token = jwt.sign(payload, process.env.JWT_KEY, { expiresIn: '1h' })
 
       const result = {
          msg: 'token created',
          token: token
       }
       return result
    } catch (error) {
       throw error
    }
}

helpers_jwt.AuthValidate = (req, res, next) => {
   const { Bearer } = req.headers

   if (!Bearer) {
      return respone(res, 401, 'silahkan login dlu')
   }

   jwt.verify(Bearer, process.env.JWT_KEYS, (err, decode) => {
      if (err) {
         return response(res, 401, err, true)
      } else {
         console.log(decode)
         next()
      }
   })
}

module.exports = helpers_jwt
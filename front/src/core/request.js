import axios from 'axios'
import { isNumber } from 'lodash'

let URI = ''

if (process.env.MODE !== 'development')
  URI = 'http://192.168.64.2:31320'
else
  URI = 'http://localhost:9000'

/**
 * Fetch
 *    Fetch a GET|DELETE endpoint (smthg that doesn't need the data)
 * 
 * @param {String} endpoint
 * @param {String} method
 * @param {String} id [default='']
 * @return {Promise}
 */
const fetch = (endpoint, method = 'get', id = '',) => {
  let url = URI
  if (!isNumber(id)) {
    url += `/${endpoint}`
  } else {
    url += `/${endpoint}/${id}`
  }

  return axios({
    method,
    url,
  })
    .then(res => Promise.resolve(res.data))
    .catch(err => Promise.reject({
      message: err.message
    }))
}

/**
 * Post
 *    Post a to an endpoint
 * 
 * @param {String} endpoint
 * @param {Object} data
 * @param {String} method [default='post']
 * @return {Promise}
 */
const post = (endpoint, data, method = 'post') => {
  let url = `${URI}/${endpoint}`

  return axios({
    method,
    url,
    data, 
  })
    .then(res => Promise.resolve(res.data))
    .catch(err => Promise.reject({
      message: err.message
    }))
}

export {
  fetch,
  post
}
import axios from 'axios'
import { isNumber } from 'lodash'

const URI = 'http://localhost:9000'

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
      err: err
    }))
}

/**
 * Post
 *    Post a to an endpoint
 * 
 * @param {String} endpoint
 * @param {String} method [default='post']
 * @param {Object} data
 * @return {Promise}
 */
const post = (endpoint, method = 'post', data) => {
  let url = `${URI}/${endpoint}`

  return axios({
    method,
    url,
    data, 
  })
    .then(res => Promise.resolve(res.data))
    .catch(err => Promise.reject({
      err: err
    }))
}

export {
  fetch,
  post
}
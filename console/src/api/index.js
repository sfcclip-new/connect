import axios from 'axios';

const baseURL = (process.env['NODE_ENV'] == 'production'
  ? '/api'
  : 'http://localhost:8080/api'
)

export default axios.create({
  baseURL: baseURL,
  headers: {
    'Content-Type': 'application/vnd.api+json',
    'Accept': 'application/vnd.api+json'
  }
})

import axios from 'axios';

export default axios.create({
  baseURL: 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/vnd.api+json',
    'Accept': 'application/vnd.api+json'
  }
})

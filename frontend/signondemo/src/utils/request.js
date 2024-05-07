import axios from 'axios'

// 创建axios实例
export const req = axios.create({
  baseURL: 'http://39.101.78.10/', 
  headers: { 
    'contentType': 'application/json',
    'Access-Control-Allow-Origin':'*',
  },
  timeout: 5000,
})
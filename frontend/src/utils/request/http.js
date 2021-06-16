import Axios from 'axios';
import {BASE_URL} from '../const';
let axios = Axios.create({
  baseURL: BASE_URL,
  timeout: 1000 * 12
});

// 请求拦截
axios.interceptors.request.use( config => {
  let token = window.localStorage.getItem('token');
  if (token != '' && token != null) {
    config.headers.token = token;
  } else {
    config.headers.token = '';
  }
  return config;
}, error => Promise.error(error))

export function get(url, params){    
  return new Promise((resolve, reject) =>{        
      axios.get(url, {            
          params: params        
      }).then(res => {
          resolve(res.data);
      }).catch(err =>{
          reject(err.data)        
  })    
});}

export function post(url, params) {
  return new Promise((resolve, reject) => {
       axios.post(url,params)
      .then(res => {
          resolve(res.data);
      })
      .catch(err =>{
          reject(err.data)
      })
  });
}

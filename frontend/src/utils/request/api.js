import { get } from './http'

// 跨链桥获取币种和支持的主网链
export const networkApi = (params) => {
  console.log(params)
  if(params.token) {
    return get('items?chain='+params.chainId+'&token='+params.token)
  }
  return get('items?chain='+params.chainId)
};


// eg:::
// import {airdropApi, receiveApi} from '../pathTo/utils/request/api';
// methods:{
//   async getDetail(){
//     let params = {
//       page: 1,
//       deviceCode: 'xxxx',
//       type: 'xxx'
//     }
//     let res = listApi(params);
//     if(res.code == 0) {
//       // doSomething
//     }
//   }
// }
import axios from 'axios';
// 创建 axios 实例
const service = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API, // api 的 base_url
  timeout: 5000 // 请求超时时间
});

// 请求拦截器
service.interceptors.request.use(
  config => {
    // 可以在此处添加请求头，例如 token
    // const token = localStorage.getItem('token');
    // if (token) {
    //   config.headers['Authorization'] = `Bearer ${token}`;
    // }
    return config;
  },
  error => {
    // 处理请求错误
    console.log(error);
    return Promise.reject(error);
  }
);

// 响应拦截器
service.interceptors.response.use(
  response => {
    const res = response.data;
    // 根据实际业务逻辑判断请求是否成功
    // if (res.code !== 200) {
    //   Message.error(res.message || 'Error');
    //   return Promise.reject(new Error(res.message || 'Error'));
    // }
    return res;
  },
  error => {
    // 处理响应错误
    console.log('err' + error);
    return Promise.reject(error);
  }
);
const request = (type: 'get' | 'post', url: string) => {
    return (id?: any, params?: any, config = {}) => {
        let newUrl = url;
        let data = {};
        if(typeof id === 'number' || typeof id === 'string') {
            newUrl = url.replace('{pk}', id as string);
            data = params;
            config = Object.assign({}, config || {});
        } else {
            data = id || {};
            config = Object.assign({}, params || {});
        }
        return service({
            method: type,
            url: newUrl,
            data,
            ...config,
          }).then(res => {
            return Promise.resolve(res);
          })
          .catch(err => {
            return Promise.reject(err);
          });
    }
}
export default request;

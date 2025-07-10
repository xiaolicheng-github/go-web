import request from './request';
export const login = request('post', '/api/login');
export const loginRegister = request('post', '/api/login/register');


export default {
    login,
    loginRegister
}
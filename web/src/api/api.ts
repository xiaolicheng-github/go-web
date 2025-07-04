import request from './request';
export const userLogin = request('post', '/api/login');

export default {
    userLogin
}
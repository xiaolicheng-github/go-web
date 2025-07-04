import type { RouteObject } from 'react-router-dom';
import Home from '../pages/home';
import Login from '../pages/login';

// 定义路由配置类型
export type RouteConfig = RouteObject & {
  meta?: {
    title?: string;
    requiresAuth?: boolean;
  };
  children?: {
    meta?: {
      title?: string;
      requiresAuth?: boolean;
    };
  }[]
};

// 路由配置数组
const routes: RouteConfig[] = [
  {
    path: '/',
    element: <Home />,
    meta: { title: '首页', requiresAuth: true }
  },
  {
    path: '/login',
    element: <Login />,
    meta: { title: '登录', requiresAuth: false }
  }
];
export default routes;
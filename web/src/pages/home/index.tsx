import { Button } from "antd";

export default function Index() {


    function handleGoLogin() {
        window.location.href = '/login';
    }
    function handleGoRegister() {
        window.location.href = '/register';
    }
    return (
        <div>
            <h1>首页</h1>
            <Button onClick={() => handleGoLogin()}>登录</Button>
            <Button onClick={() => handleGoRegister()}>注册</Button>
        </div>
    )
}
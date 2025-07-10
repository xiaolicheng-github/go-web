import { Button, Input } from "antd";
import { useState } from "react";
import { loginRegister } from "../../api/api";


export default function Index() {

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    function handleRegister() {
        loginRegister({
            username,
            password
        }).then((res) => {
            console.log(res);
        })
    }

    return (
        <div>
        <div>
            <Input placeholder="用户名" value={username} onChange={(e) => setUsername(e.target.value)} />
        </div>
        <div>
            <Input.Password placeholder="密码" value={password} onChange={(e) => setPassword(e.target.value)} />
        </div>
        <div>
            <Button type="primary" onClick={() => handleRegister()}>注册</Button>
        </div>
    </div>
    )
}
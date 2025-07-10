import { useState } from "react"
import { Input, Button } from 'antd';
import { login } from '../../api/api';



export default function Index() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');


    function handleLogin() {
        login({
            username: 'xxxx',
            password: 'xxxxx'
        }).then(res => {
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
                <Button type="primary" onClick={() => handleLogin()}>登录</Button>
            </div>
        </div>
    )
}
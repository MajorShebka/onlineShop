import React from "react";
import axios from "axios";
import {Cookies} from "react-cookie";
import products from "./Products.tsx";
import Products from "./Products.tsx";
import {Navigate} from "react-router-dom";

function Auth(){
    const [login, setLogin] = React.useState('')
    const [password, setPassword] = React.useState('')

    const handleSubmit = (e: React.FormEvent) => {
        const article = {login:login, password:password}
        console.log(article)

        axios.post("http://localhost:8080/signIn", article).then(res=>{
            const cookies = new Cookies()
            cookies.set("token", res.data["token"])
            cookies.set("login", login)
            alert("Hello " + login)
            window.location.replace("/")
        })

        e.preventDefault()
    }

    return <center><form onSubmit={handleSubmit}><div className={"wrapper"}>
        <h1>Login</h1><br/>
        <input type={"text"} id={"login"} value={login} onChange={(e) => setLogin(e.target.value)}/><br/>
        <h1>Password</h1><br/>
        <input type={"password"} id={"password"} value={password} onChange={(e) => setPassword(e.target.value)}/><br/>
        <button className={"btn"}>Sign In</button>
    </div></form></center>
}

export default Auth
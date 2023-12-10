import {Cookies} from "react-cookie";
import React from "react";

export const Header = () => {
    const cookies = new Cookies()
    const handleSubmit = (e: React.FormEvent) => {
        cookies.set("token", "")
        cookies.set("login", "")
    }
    return (
        <div className="header">
            {cookies.get("token") == "" ? <a href={"/auth"}><button>Auth</button></a>:
                <form onSubmit={handleSubmit}><h1><a href={"/basket"}>{cookies.get("login")}</a></h1>
                    <button>Sign Out</button>
                </form>
            }
            {/*<img src={logo} alt="logo" />*/}
            {/*<img src={cart} alt="cart icon" />*/}
        </div>
    )
}
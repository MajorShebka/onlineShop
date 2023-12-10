import {useEffect, useState} from "react";
import {Cookies} from "react-cookie";
import {Link} from "react-router-dom";

function Basket(){
    const [products, setProducts] = useState(null)
    const cookies = new Cookies()
    // axios.defaults.headers.get['user'] = cookies.get("token")
    // axios.get("http://localhost:8080/productByCustomer").then(res=>{
    //     setProducts(res.data)
    // })
    useEffect(() => {
            fetch("http://localhost:8080/productByCustomer", {
                method: "GET",
                headers: {
                    "user": cookies.get("token"),
                }
            }).then(res => {
                    return res.json()
                }
            ).then(data => {
                setProducts(data)
            })
        },
        [])

    return <div className="cards">
        {products && products.map((pr) => {
            return (
                <div className="card" key={pr.id}>
                    <img src={pr.image} alt="img"/>
                    <h4>{pr.name}</h4>
                    <p className="price">${pr.price}</p>
                    <Link to={"/product/" + pr.id}><center><button>Show more</button></center></Link>
                </div>
            )
        })}
    </div>
}

export default Basket
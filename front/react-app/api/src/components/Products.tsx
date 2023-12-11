import {useEffect, useState} from "react";
import {Link} from "react-router-dom";
import {Cookies} from "react-cookie";

function Products() {
    const [products, setProducts] = useState(null)
    useEffect(() => {
            fetch("http://localhost:8080/product").then(res => {
                return res.json()
            }
        ).then(data => {
            setProducts(data)
            })
        },
        [])
    const c = new Cookies()
    console.log(c.get("token"))
    return (
        <div className="cards">
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
    )
}

export default Products
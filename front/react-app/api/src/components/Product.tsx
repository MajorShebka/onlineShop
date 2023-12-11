import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import axios from "axios";
import {Cookies} from "react-cookie";
// import axios from "axios/index";

function Product(){
    const {id} = useParams()
    const req = `http://localhost:8080/product/${id}`
    const [product, setProduct] = useState(null)
    useEffect(() => {
            fetch(req).then(res => {
                    return res.json()
                }
            ).then(data => {
                setProduct(data)
            })
        },
        [])

    const addProduct = (e: React.FormEvent) => {
        const article = {"productId":product.id}
        const cookies = new Cookies()
        axios.defaults.headers.post['user'] = cookies.get("token")
        axios.post("http://localhost:8080/productToCustomer", article).then(res=>{
            alert("Успешно добавлено!")
        })

        e.preventDefault()
    }

    return <div>
        {product == null ? <h2>Loading...</h2> : <form onSubmit={addProduct}><div className="product">
            <img src={"/" + product.image} alt="image" />
            <div className="info">
                <p className="title">{product.name}</p>
                <p className="price">$ {product.price}</p>
                <div className="category">
                    <span>Category:</span>
                    <p>{product.type}</p>
                </div>
                <p className="description">{product.desc}</p>
                <button>
                    Add to cart
                </button>
            </div>
        </div>
        </form>
        }

    </div>
}

export default Product
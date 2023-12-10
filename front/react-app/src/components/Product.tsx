import {useEffect, useState} from "react";
import {useParams} from "react-router-dom";

function Product(){
    const {id} = useParams()
    const req = `http://localhost:8080/product/${id}`
    console.log(id)
    console.log(req)
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
    return <div>
        {product == null ? <h2>Loading...</h2> : <div className="product">
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
        }

    </div>
}

export default Product
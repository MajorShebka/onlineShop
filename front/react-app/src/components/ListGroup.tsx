import {Fragment} from "react";

function ListGroup(){
    const items = [
        "Minsk",
        "Brest",
        "Gomel"
    ]
    return(
        <Fragment>
            <h1>List</h1>
                <ul className="list-group">
                    {items.map(item => <li className="list-group-item">{item}</li>)}
                </ul>
        </Fragment>
    )
}

export default ListGroup
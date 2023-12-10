import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import 'bootstrap/dist/css/bootstrap.css'
import './App.css'
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import Products from "./components/Products.tsx";
import Product from "./components/Product.tsx";
import Auth from "./components/Auth.tsx";
import { CookiesProvider } from "react-cookie"
import {Header} from "./components/Header.tsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <Products/>,
    },
    {
        path: "/product/:id",
        element: <Product/>,
    },
    {
        path: "/auth",
        element: <Auth/>
    }
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
    <CookiesProvider>
  <React.StrictMode>
      <Header/>
      <RouterProvider router={router}/>
  </React.StrictMode>,
    </CookiesProvider>

)

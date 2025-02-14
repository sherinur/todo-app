import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import App from './App'
import "@wailsapp/runtime";
import TodoList from './components/TodoList';

const container = document.getElementById('root')

const root = createRoot(container)

root.render(
    <React.StrictMode>
        <TodoList/>
    </React.StrictMode>
)

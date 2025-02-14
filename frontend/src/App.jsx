import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import * as AppBackend from "../wailsjs/go/app/App";
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);

    function greet() {
        AppBackend.AddTodo(name, "Description", new Date().toISOString(), "high")
            .then(() => {
                setResultText(`Todo added: ${name}`);
            })
            .catch((error) => {
                setResultText(`Error: ${error.message}`);
            });
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={greet}>Add Todo</button>
            </div>
        </div>
    )
}

export default App;
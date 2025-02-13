import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';

function TodoList() {
    // Состояние для хранения списка задач
    const [todos, setTodos] = useState([]);
    // Состояние для хранения текста новой задачи
    const [newTodo, setNewTodo] = useState('');

    const addTodo = () => {
        if (newTodo.trim() === '') return;
        const todo = {
            id: Date.now(),
            text: newTodo,
            done: false,
        };
        setTodos([...todos, todo]);
        setNewTodo('');
    };

    const toggleTodo = (id) => {
        setTodos(
            todos.map((todo) =>
                todo.id === id ? { ...todo, done: !todo.done } : todo
            )
        );
    };

    // Функция для удаления задачи
    const deleteTodo = (id) => {
        setTodos(todos.filter((todo) => todo.id !== id));
    };

    return (
        <div className="container mt-5">
            <h1 className="text-center mb-4">Todo List</h1>
            <div className="input-group mb-3">
                <input
                    type="text"
                    className="form-control"
                    placeholder="Enter a new task"
                    value={newTodo}
                    onChange={(e) => setNewTodo(e.target.value)}
                    onKeyPress={(e) => e.key === 'Enter' && addTodo()}
                />
                <button className="btn btn-primary" onClick={addTodo}>
                    Add Task
                </button>
            </div>
            <ul className="list-group">
                {todos.map((todo) => (
                    <li
                        key={todo.id}
                        className={`list-group-item d-flex justify-content-between align-items-center ${
                            todo.done ? 'list-group-item-success' : ''
                        }`}
                    >
                        <span
                            style={{
                                textDecoration: todo.done ? 'line-through' : 'none',
                                cursor: 'pointer',
                            }}
                            onClick={() => toggleTodo(todo.id)}
                        >
                            {todo.text}
                        </span>
                        <div>
                            <button
                                className="btn btn-sm btn-danger me-2"
                                onClick={() => deleteTodo(todo.id)}
                            >
                                Delete
                            </button>
                            <button
                                className={`btn btn-sm ${
                                    todo.done ? 'btn-warning' : 'btn-success'
                                }`}
                                onClick={() => toggleTodo(todo.id)}
                            >
                                {todo.done ? 'Undo' : 'Done'}
                            </button>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default TodoList;
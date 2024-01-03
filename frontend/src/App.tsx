import { useEffect, useState } from "react";
import { getTodo, postTodo, patchTodo, deleteTodo } from "./api";

interface Todo {
  id: number;
  content: string;
}

export default function App() {
  const [ todosRefresh, setTodosRefresh ] = useState(false);

  const [ input, setInput ] = useState("")
  const [ todos, setTodos ] = useState<Todo[]>([]);

  useEffect(() => {getTodo().then(setTodos);}, [todosRefresh]);

  function postTodoHandler() {
    postTodo({content: input})
      .then(() => setTodosRefresh(!todosRefresh));
  }
  
  return (
    <>
      <input type="text" value={input} onChange={e => setInput(e.target.value)} />
      <button onClick={postTodoHandler}>Add</button>
      <ul style={{listStyleType: "none"}}>
        {todos.map((todo: Todo) => <li key={todo.id}>{todo.content}</li>)}
      </ul>
    </>
  );
}
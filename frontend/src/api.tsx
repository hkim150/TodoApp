import axios from 'axios';

export function getTodo() {
    return axios({
        method: "GET",
        url: "/todos",
    })
    .then(res => res.data);
}

interface postTodoParams {
    content: string;
}

export function postTodo(params: postTodoParams) {
    return axios({
        method: "POST",
        url: "/todos",
        headers: { 'Content-Type': 'application/json' },
        data: params
    });
}

interface patchTodoParams {
    id: number;
    content: string;
}

export function editTodo({ id, ...body }: patchTodoParams) {
    return axios({
        method: "PATCH",
        url: "/todos/" + id,
        headers: { 'Content-Type': 'application/json' },
        data: body
    });
}

interface deleteTodoParams {
    id: number;
}

export function deleteTodo({ id }: deleteTodoParams) {
    return axios({
        method: "DELETE",
        url: "todos/" + id
    });
}
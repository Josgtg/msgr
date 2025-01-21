async function getAllUsers() {
    let url = "http://localhost:5174/api/users/";
    let response = await fetch(url, {
        method: "GET",
    });

    return await response.json();
}

async function getUser(id) {
    let url = `http://localhost:5174/api/users/${id}`;
    let response = await fetch(url, {
        method: "GET",
    });

    return await response.json();
}

async function insertUser(name, password, email) {
    let url = `http://localhost:5174/api/users/?name=${name}&password=${password}&email=${email}`;
    let response = await fetch(url, {
        method: "POST",
    });

    return await response.json();
}

async function deleteUser(id) {
    let url = `http://localhost:5174/api/users/${id}`;
    let response = await fetch(url, {
        method: "DELETE",
    });

    return await response.json();
}

export { getAllUsers, getUser, insertUser, deleteUser };

async function getUserChats(id) {
    let url = `http://localhost:5174/api/chats/user/${id}`;
    let response = await fetch(url, {
        method: "GET",
    });

    return await response.json();
}

async function getChatMessages(id) {
    let url = `http://localhost:5174/api/messages/chat/${id}`;
    let response = await fetch(url, {
        method: "GET",
    });

    return await response.json();
}

export { getUserChats, getChatMessages };

function goToChat(chat) {
    window.location.href = `/chats/${chat.id}`;
}

export { goToChat };

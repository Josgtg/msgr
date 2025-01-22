import { getChatMessages } from "$lib/chatRequests.js";

export async function load({ params }) {
    let messages = await getChatMessages(params.id);
    return {
        messages,
    };
}

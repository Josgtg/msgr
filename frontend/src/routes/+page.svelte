<script>
    import ContactCard from './contact-card.svelte'
    import { getUserChats } from '$lib/chatRequests';
    import { getUser } from '$lib/userRequests';

    let userID = "8a48a7fc-e165-45ae-a53a-4875fe23a03e";
</script>

<h2 class="loading">Loading chats...</h2>
{#await getUserChats(userID) then chats}
    <style>
        .loading {
            display: none;
        }
    </style>
    {#if !chats.title}
        {#if chats.length > 0}
            <div class="card-container">
                {#each chats as c}
                    {#await getUser(c.first_user == userID ? c.second_user : c.first_user) then user}
                        <ContactCard user={user} chat={c}/>
                    {/await}
                {/each}
            </div>
        {:else}
            <h2>You don't have any chats!</h2>
        {/if}
    {:else}
        <h2>{chats.title}: {chats.message}</h2>
    {/if}
{/await}

<style>
    .card-container {
        display: flex;
        flex-wrap: wrap;
    }
</style>
<script>
    import { goToChat } from "$lib/routing"

    export let user;
    export let chat;
</script>

<div class="container">
    <div class="card">
        <img src="/avatar-default.png" alt="Avatar">
        <div class="card-content">
            <h2>{user.name}</h2>
            <h4>Last Message:</h4>
            <p>{user.email}</p>
        </div>
        <button onclick={goToChat(chat)}>Chat</button>
    </div>
</div>

<style>
    .container {
        margin: auto;
        margin-top: 5vh;

        display: flex;
        flex-wrap: wrap;
        
        justify-content: center;

    }
    
    .card {
        margin: 2vh;
        
        border: 2px solid gray;
        
        border-radius: 10%;
        
        background-color: var(--secondary);

        width: 33vh;

        overflow: hidden;
    }

    .card img {
        width: 100%;
        height: auto;
        padding: 5%;
        padding-bottom: 0;
        margin-bottom: 0;
    }

    .card button {
        width: 100%;
        margin: auto;
        padding: 2%;
        font-size: 1.4em;
        background-color: #44cc44;
        border: none;
    }

    .card h2 {
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        color: var(--primary);
        border-top: 2px solid gray;
        width: 100%;
        padding: 5% 5% 0 5%;
        margin-top: 0;
        margin-bottom: 0;
    }

    .card h4, p {
        margin: 0;
        color: gainsboro;
        font-size: 1em;
    }

    @media only screen and (max-height: 550px) {
        .card h2 {
            padding: 6%;
            overflow-x: scroll;
        }
        .card h4, p  {
            display: none;
        }
    }

    .card h4 {
        padding: 5% 5% 5% 5%;
    }

    .card p {
        padding: 0% 5% 10% 5%;
        max-height: 20vh;
        overflow-y: scroll;
    }

    *::-webkit-scrollbar {
        display: none;
    }

    * {
        -ms-overflow-style: none;
        scrollbar-width: none;
    }
</style>

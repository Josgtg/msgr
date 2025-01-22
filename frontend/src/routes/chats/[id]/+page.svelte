<script>
    let messages = $props();
</script>
{#if !messages.title}
    {#each messages.data.messages as m}
        <p>{m.sent_at}</p><p>{m.message}</p>
        <br><br>
    {/each}
{:else}
    <h1>{messages.title}: {messages.message}</h1>
{/if}

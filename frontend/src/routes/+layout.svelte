<slot />

<style>
    :global(:root) {
        --primary: #F14A00;
        --primary-strong: #C62300;
        --secondary: #500073;
        --background: #2A004E;
    }

    :global(*) {
        box-sizing: border-box;
        color: white;

        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    }
    :global(body) {
        background: var(--background);
    }
</style>
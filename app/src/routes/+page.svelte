<script>
import { goto } from '$app/navigation';
</script>

{#if localStorage.getItem('username') == null}
    {goto("/register")}
{/if}

{#if localStorage.getItem('username') != null}
    {goto("/chat")}
{/if}
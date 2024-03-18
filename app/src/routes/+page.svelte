<script>
    import { goto } from '$app/navigation'

    // redirect if logged in already
    localStorage.getItem('username') ?
    goto("/chat") :
    goto("/register")

</script>
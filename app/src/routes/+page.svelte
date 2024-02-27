<script>
    import { goto } from '$app/navigation'

    // redirect if logged in already
    console.log(localStorage.getItem('username'))
    localStorage.getItem('username') ?
    goto("/chat") :
    goto("/register")

</script>
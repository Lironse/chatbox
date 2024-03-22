<script>
	import { goto } from '$app/navigation';

	// redirect if logged in already
	localStorage.getItem('token') ? goto('/chat') : goto('/register');
</script>

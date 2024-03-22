<script lang="ts">
	import LoginForm from './components/LoginForm.svelte';
</script>

<div class="grid h-full">
	<div class="card w-1/3 place-self-center flex flex-col p-4 space-y-4">
		<LoginForm/>
		<div>
			<span class="pt-10">New account needed?</span>
			<a href="../register" class="text-blue-500"> Register</a>
		</div>
	</div>
</div>
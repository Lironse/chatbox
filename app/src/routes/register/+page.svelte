<script lang="ts">
	import RegisterForm from './components/RegisterForm.svelte';
</script>

<div class="grid h-full">
	<div class="card w-1/3 place-self-center flex flex-col p-4 space-y-4">
		<RegisterForm/>
		<div>
			<span class="pt-10">Have an account already?</span>
			<a href="../login" class="text-blue-500"> Log in</a>
		</div>
	</div>
</div>
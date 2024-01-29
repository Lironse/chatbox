<script>
	import { FileButton } from '@skeletonlabs/skeleton';

</script>
<div class="grid h-full">
	<form class="w-1/8 place-self-center card p-4 text-token space-y-4 flex flex-col">
		<h2 class="h2 pb-2">Log in</h2>
			<input class="input" type="text" placeholder="Username" />
			<FileButton name="files" class="w-full" button="btn variant-soft-primary">Upload private key (.PEM)</FileButton>
			<button class="w-full btn variant-filled">Sign in</button>
		<div>
			<span class="pt-10">Don't have an account?</span>
			<span class="text-blue-500" role="button" onclick="window.location.href='../register'">Register<span>
		</div>
	</form>
</div>
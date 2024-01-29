<script lang='ts'>
	import { usernameStore } from '$lib/stores.ts'

    import { onMount } from 'svelte';

    let username: string = '';
    let usernameValidity: string = 'Valid';
	let registered: boolean = false;

    onMount( async () => {
		// Check if the username for registration is valid
        const usernameInput = document.getElementById('username-input') as HTMLInputElement;
        usernameInput.addEventListener('input', () => {
            checkUsernameValidity();
        });
    });

    function checkUsernameValidity(): void {
        if (username.length < 3) {
            usernameValidity = 'Username too short.';
        } else if (/[^a-zA-Z0-9]/.test(username)) {
            usernameValidity = 'Username may only contain letters and numbers.';
        } else {
            usernameValidity = 'Valid';
        }
    }

    function register(): void {
		checkUsernameValidity()
		if (usernameValidity == 'Valid') {
		}
    }

	function usernameAvailability(response: string): void {
		if (response == "registered") {
			registered = true
			usernameStore.set(username);
			localStorage.setItem('username', username)
			return
		}
		usernameValidity = "Username is already taken."
	}

</script>

<div class="grid h-full">
	{#if !registered}
		<form
			on:submit|preventDefault={() => register()}
			class="w-1/3 place-self-center card p-4 text-token space-y-4 flex flex-col"
		>
			<h2 class="h2 pb-2">Register</h2>
			<input
				class="input"
				aria-autocomplete="none"
				type="text"
				placeholder="Username"
				bind:value={username}
				id="username-input"
				maxlength=16
			/>

			{#if usernameValidity != "Valid"}
				<p class="text-red-500 mt-2">{usernameValidity}</p>
			{/if}

			<button type="submit" class="w-full btn variant-filled text-center">Register</button>

			<div>
				<span class="pt-10">Have an account already?</span>
				<span class="text-blue-500" role="button" onclick="window.location.href='../login'">
					Log in
				</span>
			</div>
		</form>
	{/if}

	{#if registered}
		<div class="grid h-full">
			<div class="w-1/3 place-self-center card p-4 text-token space-y-4 flex flex-col">
				<h2 class="h2">
					<span class="bg-gradient-to-br from-blue-500 to-cyan-300 bg-clip-text text-transparent box-decoration-clone">Welcome!</span>
				</h2>
				<span>your private key has been saved to some location.</span>
				<div class="flex justify-end">
					<a type="button" href="/chat" class="btn variant-filled">Continue</a>
				</div>
			</div>
		</div>
	{/if}
</div>

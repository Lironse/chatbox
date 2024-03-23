<script lang="ts">
	import { username } from '$lib/stores.ts';
	import { goto } from '$app/navigation';
	import { keys } from '$lib/keys.ts';

	let usernameInput: string = '';
	async function register(): Promise<void> {
		// TODO replace this with a global IP
		const registrationUrl = 'http://176.230.36.90:27357/register';

		const registrationData = {
			username: usernameInput,
			key: arrayBufferToBase64((await keys).publicKey)
		};
		const requestOptions = {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(registrationData)
		};

		let response = await fetch(registrationUrl, requestOptions);

		if (!response.ok) {
			console.log('Failed to register client: ' + response.statusText);
			return;
		}

		console.log('registered successfully as', usernameInput);
		username.set(usernameInput);
		goto('../chat');
	}

	function arrayBufferToBase64(buffer: ArrayBuffer): string {
		const bytes = new Uint8Array(buffer);
		const string = String.fromCharCode(...bytes);
		return btoa(string);
	}
</script>

<form on:submit|preventDefault={register} class="text-token space-y-4 flex flex-col">
	<h2 class="h2 pb-2 font-bold">Register</h2>
	<input
		bind:value={usernameInput}
		class="input"
		aria-autocomplete="none"
		type="text"
		placeholder="Username"
		id="username-input"
		maxlength="16"
		autocomplete="off"
		autocorrect="off"
		autocapitalize="off"
		spellcheck="false"
	/>
	<button type="submit" class="w-full btn variant-filled text-center">Register</button>
</form>

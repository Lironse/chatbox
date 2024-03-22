<script lang="ts">
	import { username } from '$lib/stores.ts';
	import { goto } from '$app/navigation';

	let usernameInput: string = '';
	let privateKeyInput: string = '';

	async function login(): Promise<void> {
		// TODO replace this with a global IP
		const registrationUrl = 'http://176.230.36.90:27357/login';
		const registrationData = {
			name: usernameInput,
			key: 'someKEYKEYKEKY' // TODO actually generate keys
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

		console.log(await response.json());
		username.set(usernameInput);
		goto('../chat');
	}
</script>

<form
	on:submit|preventDefault={login}
	class="text-token space-y-4 flex flex-col"
>
	<h2 class="h2 pb-2 font-bold">Log in</h2>
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
	<input
		bind:value={privateKeyInput}
		class="input"
		aria-autocomplete="none"
		type="text"
		placeholder="PrivateKey"
		id="private-key-input"
		autocomplete="off"
		autocorrect="off"
		autocapitalize="off"
		spellcheck="false"
	/>

	<button type="submit" class="w-full btn variant-filled text-center">Log in</button>
</form>
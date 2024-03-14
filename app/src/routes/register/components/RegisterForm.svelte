<script lang="ts">
	import { sendPacket } from '$lib/socket.ts';
	import { username } from '$lib/stores.ts';
	import { onMount } from 'svelte';
	import { Packet } from '$lib/packet.ts';
	import { generateKeyPair } from '$lib/keys.ts';
	import { goto } from '$app/navigation';

	generateKeyPair().then((keys) => {
		console.log('Public Key:', keys.publicKey);
		localStorage.setItem('publicKey', keys.publicKey);
		console.log('Private Key:', keys.privateKey);
		localStorage.setItem('privateKey', keys.privateKey);
	});

	let usernameInput: string = '';
	let usernameValidity: string = 'Valid';

	onMount(() => {
		// Check if the username for registration is valid
		const usernameInputElement = document.getElementById('username-input') as HTMLInputElement;
		usernameInputElement.addEventListener('input', () => {
			checkUsernameValidity();
		});
	});

	function checkUsernameValidity(): void {
		if (usernameInput.length < 3) {
			usernameValidity = 'Username too short.';
		} else if (/[^a-zA-Z0-9]/.test(usernameInput)) {
			usernameValidity = 'Username may only contain letters and numbers.';
		} else {
			usernameValidity = 'Valid';
		}
	}

	function register(): void {
		checkUsernameValidity();
		if (usernameValidity != 'Valid') {
			return;
		}

		username.set(usernameInput);
		localStorage.setItem('username', usernameInput);
		let registrationPacket = new Packet(
			'register',
			localStorage.getItem('publicKey') || '',
			usernameInput,
			'server'
		);
		console.log('sent registration packet');
		sendPacket(registrationPacket);
		goto('../chat');
	}
</script>

<form
	on:submit|preventDefault={() => register()}
	class="w-1/3 place-self-center card p-4 text-token space-y-4 flex flex-col"
>
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

	{#if usernameValidity != 'Valid'}
		<p class="text-red-500 mt-2">{usernameValidity}</p>
	{/if}

	<button type="submit" class="w-full btn variant-filled text-center">Register</button>

	<div>
		<span class="pt-10">Have an account already?</span>
		<a href="../login" class="text-blue-500"> Log in </a>
	</div>
</form>

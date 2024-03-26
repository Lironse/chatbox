<script lang="ts">
	import { tick } from 'svelte';
	import { Peer } from '$lib/peer';
	import { savedPeers } from '$lib/stores';
	let addPeerFormShown = false;
	let addPeerInput: string;
	let inputRef: HTMLInputElement;

	function addPeer() {
		const peer = new Peer(addPeerInput, addPeerInput);
		savedPeers.update((peers) => [...peers, peer]);
		addPeerInput = '';
		addPeerFormShown = false;
	}

	async function lookupPeer() {
		// TODO replace this with a global IP
		const server = 'http://176.230.36.90:27357/lookup';

		const lookupData = { username: addPeerInput };
		const requestOptions = {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(lookupData)
		};

		let response = await (await fetch(server, requestOptions)).json();

		if (response.status == 'success') {
			console.log(response.status);
			addPeer();
		} else {
			alert('This user was not found!');
		}
	}
</script>

{#if addPeerFormShown}
	<form name="addPeer" on:submit={lookupPeer}>
		<input
			bind:value={addPeerInput}
			class="input"
			type="text"
			placeholder="Username"
			maxlength="16"
			bind:this={inputRef}
		/>
		<input type="submit" hidden />
	</form>
{:else}
	<button
		on:click={async () => {
			addPeerFormShown = true;
			await tick();
			inputRef?.focus();
		}}
		type="button"
		class="btn variant-filled-surface"
	>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 24 24"
			stroke-width="1.5"
			stroke="currentColor"
			class="w-6 h-6"
		>
			<path
				stroke-linecap="round"
				stroke-linejoin="round"
				d="M18 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 1 1-6.75 0 3.375 3.375 0 0 1 6.75 0ZM3 19.235v-.11a6.375 6.375 0 0 1 12.75 0v.109A12.318 12.318 0 0 1 9.374 21c-2.331 0-4.512-.645-6.374-1.766Z"
			/>
		</svg>
	</button>
{/if}

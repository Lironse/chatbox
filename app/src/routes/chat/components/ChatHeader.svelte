<script lang="ts">
	import { selectedPeer } from '$lib/stores.ts';
	import { Avatar } from '@skeletonlabs/skeleton';
	import ConnectButton from './ConnectButton.svelte';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function toggleConnection() {
		dispatch('toggleConnection');
	}
</script>

<div class="flex flex-row justify-center p-4">
	<div class="flex flex-row gap-5 justify-center p-4 w-1/3 card">
		{#if $selectedPeer}
			<Avatar src="https://i.pravatar.cc/?img={$selectedPeer?.avatar}" width="w-12" />
			<span class="h2 font-bold align-middle">{$selectedPeer?.name}</span>
			<ConnectButton on:toggleConnection={toggleConnection} />
		{/if}
	</div>
</div>

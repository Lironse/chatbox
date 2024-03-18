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

<div class="flex flex-row gap-5 justify-start items-center p-4 border-b border-gray-500">
	{#if $selectedPeer}
		<Avatar src="https://i.pravatar.cc/?img={$selectedPeer?.avatar}" width="w-12" />
		<span class="h2 font-bold align-middle">{$selectedPeer?.name}</span>
		<ConnectButton on:toggleConnection={toggleConnection} />
	{/if}
</div>

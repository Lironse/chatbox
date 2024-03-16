<script lang="ts" context="module">
	import { Avatar } from '@skeletonlabs/skeleton';
	import { messageFeed, username, selectedPeer, connectionStatus, rtcs } from '$lib/stores.ts';
	import { goto } from '$app/navigation';
	import { RTC } from '$lib/rtc';
	import { get } from 'svelte/store';

	if (localStorage.getItem('username') == null) {
		goto('../register');
	}

	function openConnection() {
		let rtc = new RTC(get(selectedPeer).name);
		rtcs.update(rtcs => [...rtcs, rtc]);
		rtc.sendOffer();
	}

	function closeConnection() {
		get(rtcs).forEach((rtc) => {
			if (rtc.peerName == get(selectedPeer).name) {
				rtc.conn.close()
				get(rtcs).splice(get(rtcs).indexOf(rtc), 1);
			}
		});
		connectionStatus.set('closed')
	}

</script>

<div class="bg-transparent p-4 flex flex-col gap-3 overflow-y-auto">
	{#if $selectedPeer && $connectionStatus != 'open'}
		<button on:click={openConnection} class="btn variant-filled-secondary">Connect</button>
	{/if}

	{#if $connectionStatus == 'open'}
		<button on:click={closeConnection} class="btn variant-soft-primary">Close</button>
	{/if}

	{#each $messageFeed as message}
		{#if message.name == $username}
			<div class="grid grid-cols-[auto_1fr] gap-2">
				<Avatar src="https://i.pravatar.cc/?img={message.avatar}" width="w-12" />
				<div class="card p-4 rounded-tl-none space-y-2 variant-soft-secondary">
					<header class="flex justify-between items-center">
						<p class="font-bold">{$username}</p>
						<small class="opacity-50">{message.timestamp}</small>
					</header>
					<p>{message.message}</p>
				</div>
			</div>
		{:else}
			<div class="grid grid-cols-[1fr_auto] gap-2">
				<div class="card p-4 rounded-tr-none space-y-2">
					<header class="flex justify-between items-center">
						<p class="font-bold">{message.name}</p>
						<small class="opacity-50">{message.timestamp}</small>
					</header>
					<p>{message.message}</p>
				</div>
				<Avatar src="https://i.pravatar.cc/?img={message.avatar}" width="w-12" />
			</div>
		{/if}
	{/each}
</div>

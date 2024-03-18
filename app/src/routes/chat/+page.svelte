<script lang="ts">
	import { Avatar } from '@skeletonlabs/skeleton';
	import { messageFeed, username, selectedPeer, connectionStatus, rtcs } from '$lib/stores.ts';
	import { goto } from '$app/navigation';
	import { RTC } from '$lib/rtc';
	import { get } from 'svelte/store';
	import ChatHeader from './components/ChatHeader.svelte';
	import ChatFooter from './components/ChatFooter.svelte';

	if (localStorage.getItem('username') == null) {
		goto('/');
	}

	function toggleConnection() {
		$connectionStatus == 'open' ? closeConnection() : openConnection();
	}

	function openConnection() {
		let rtc = new RTC(get(selectedPeer).name);
		rtcs.update((rtcs) => [...rtcs, rtc]);
		rtc.sendOffer();
	}

	function closeConnection() {
		get(rtcs).forEach((rtc) => {
			if (rtc.peerName == get(selectedPeer).name) {
				rtc.conn.close();
				rtcs.set(get(rtcs).splice(get(rtcs).indexOf(rtc), 1));
			}
		});
		connectionStatus.set('closed');
	}
</script>

<main class="flex flex-col h-screen">
	<ChatHeader on:toggleConnection={toggleConnection} />

	<div class="bg-transparent p-4 grow flex flex-col gap-3 overflow-y-auto">
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

	<ChatFooter />
</main>

<script lang="ts">
	import { savedPeers, selectedPeer } from '$lib/stores';
	import PeerTile from './PeerTile.svelte';
	import type { Peer } from '$lib/peer';

	let peers: Peer[];
	savedPeers.subscribe((value) => {
		peers = value;
	});

	function selectPeer(peer: Peer) {
		selectedPeer.set(peer);
	}

	function removePeer(peer: Peer) {
		savedPeers.set($savedPeers.filter((p) => p !== peer));
	}
</script>

<div class="peer-list-container flex flex-col overflow-y-auto">
	<ul class="peer-list flex flex-col">
		{#each peers as peer}
			<div class="flex flex-row gap-2 mb-1 mr-2">
				<button class="w-full" on:click={() => selectPeer(peer)}>
					<PeerTile {peer} />
				</button>
				{#if peer.name == $selectedPeer?.name}
					<button on:click={() => removePeer(peer)} class="btn-icon variant-filled-error">
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
								d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0"
							/>
						</svg>
					</button>
				{/if}
			</div>
		{/each}
	</ul>
</div>

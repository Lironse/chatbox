<script lang="ts">
	import type { Peer } from '$lib/peer';
	import { Avatar } from '@skeletonlabs/skeleton';
	export let peer: Peer;
</script>

<div class="flex flex-row items-center gap-5">
	<button
		type="button"
		class="w-full btn variant-filled-surface flex items-center gap-5 justify-start"
	>
		<Avatar src="https://i.pravatar.cc/?img={peer.avatar}" width="w-12 h-12 rounded-full" />
		<span class="overflow-hidden whitespace-nowrap max-w-[calc(100%-3rem)]">{peer.name}</span>
	</button>
</div>
